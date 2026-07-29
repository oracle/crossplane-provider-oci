/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"container/list"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/v2/pkg/meta"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	upjetterraform "github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kruntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	clusterv1beta1 "github.com/oracle/provider-oci/apis/cluster/v1beta1"
	namespacedv1beta1 "github.com/oracle/provider-oci/apis/namespaced/v1beta1"
)

const (
	// error messages
	errNoProviderConfig           = "no providerConfigRef provided"
	errGetProviderConfig          = "cannot get referenced ProviderConfig"
	errTrackUsage                 = "cannot track ProviderConfig usage"
	errExtractCredentials         = "cannot extract credentials"
	errUnmarshalCredentials       = "cannot unmarshal oci credentials as JSON"
	errUnsupportedManaged         = "resource is not a managed"
	errUnsupportedProviderCfgKind = "unsupported providerConfigRef.kind"
)

const (
	credentialKeyTenancyOCID                     = "tenancy_ocid"
	credentialKeyUserOCID                        = "user_ocid"
	credentialKeyPrivateKey                      = "private_key"
	credentialKeyPrivateKeyPath                  = "private_key_path"
	credentialKeyFingerprint                     = "fingerprint"
	credentialKeyRegion                          = "region"
	credentialKeyAuth                            = "auth"
	credentialKeyConfigFileProfile               = "config_file_profile"
	credentialKeyWorkloadIdentityTokenPath       = "workload_identity_token_path"
	credentialKeyTokenExchangeDomainURL          = "token_exchange_domain_url"
	credentialKeyTokenExchangeAuth               = "token_exchange_auth"
	credentialKeyTokenExchangeClientID           = "token_exchange_client_id"
	credentialKeyTokenExchangeClientSecret       = "token_exchange_client_secret"
	credentialKeyTokenExchangeRequestedTokenType = "token_exchange_requested_token_type"
	credentialKeyTokenExchangeSubjectTokenType   = "token_exchange_subject_token_type"
	credentialKeyTokenExchangeResourceType       = "token_exchange_resource_type"
	credentialKeyTokenExchangeRPSTExpiration     = "token_exchange_rpst_exp"
	credentialKeyTokenExchangePublicKey          = "token_exchange_public_key"
)

type setupOptions struct {
	enableFrameworkProvider bool
	isSDKv2Resource         func(string) bool
}

// SetupOption customizes Terraform setup behavior.
type SetupOption func(*setupOptions)

// WithFrameworkProvider controls whether setup returns a Plugin Framework
// provider instance for framework-routed resources.
func WithFrameworkProvider(enabled bool) SetupOption {
	return func(o *setupOptions) {
		o.enableFrameworkProvider = enabled
	}
}

// WithSDKv2ResourcePredicate controls which Terraform resources receive
// in-process SDKv2 provider meta.
func WithSDKv2ResourcePredicate(predicate func(string) bool) SetupOption {
	return func(o *setupOptions) {
		o.isSDKv2Resource = predicate
	}
}

type terraformResourceTyper interface {
	GetTerraformResourceType() string
}

// TerraformSetupBuilder builds a terraform.SetupFn for in-process no-fork
// connectors. Build-time Terraform values are intentionally not required at
// runtime when all resources are routed through SDKv2 or Framework connectors.
func TerraformSetupBuilder(opts ...SetupOption) upjetterraform.SetupFn {
	options := setupOptions{}
	for _, opt := range opts {
		opt(&options)
	}

	return func(ctx context.Context, kube client.Client, mg resource.Managed) (upjetterraform.Setup, error) {
		ps := upjetterraform.Setup{}

		pcSpec, err := resolveProviderConfig(ctx, kube, mg)
		if err != nil {
			return ps, errors.Wrap(err, "cannot resolve provider config")
		}

		data, err := resource.CommonCredentialExtractor(ctx, pcSpec.Credentials.Source, kube, pcSpec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		ociCreds := map[string]string{}
		if err := json.Unmarshal(data, &ociCreds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		cfg := providerConfigurationFromCredentials(ociCreds)
		ps.Configuration = cfg

		if options.enableFrameworkProvider {
			setFrameworkProvider(&ps)
		}

		if !options.shouldConfigureSDKv2Provider(mg) {
			return ps, nil
		}

		uid, err := resolveProviderConfigIdentity(ctx, kube, mg)
		if err != nil {
			return ps, fmt.Errorf("cannot resolve ProviderConfig identity: %w", err)
		}
		if uid == "" {
			return ps, fmt.Errorf("ProviderConfig has empty UID")
		}

		providerMeta, err := getOrConfigureProviderMeta(ctx, uid, cfg)
		if err != nil {
			return ps, fmt.Errorf("cannot get or init OCI provider: %w", err)
		}
		ps.Meta = providerMeta
		ps.Scheduler = upjetterraform.NewNoOpProviderScheduler()

		return ps, nil
	}
}

func providerConfigurationFromCredentials(ociCreds map[string]string) map[string]any {
	config := map[string]any{
		credentialKeyTenancyOCID:                     ociCreds[credentialKeyTenancyOCID],
		credentialKeyUserOCID:                        ociCreds[credentialKeyUserOCID],
		credentialKeyPrivateKey:                      ociCreds[credentialKeyPrivateKey],
		credentialKeyPrivateKeyPath:                  ociCreds[credentialKeyPrivateKeyPath],
		credentialKeyFingerprint:                     ociCreds[credentialKeyFingerprint],
		credentialKeyRegion:                          ociCreds[credentialKeyRegion],
		credentialKeyAuth:                            ociCreds[credentialKeyAuth],
		credentialKeyConfigFileProfile:               ociCreds[credentialKeyConfigFileProfile],
		credentialKeyWorkloadIdentityTokenPath:       ociCreds[credentialKeyWorkloadIdentityTokenPath],
		credentialKeyTokenExchangeDomainURL:          ociCreds[credentialKeyTokenExchangeDomainURL],
		credentialKeyTokenExchangeAuth:               ociCreds[credentialKeyTokenExchangeAuth],
		credentialKeyTokenExchangeClientID:           ociCreds[credentialKeyTokenExchangeClientID],
		credentialKeyTokenExchangeClientSecret:       ociCreds[credentialKeyTokenExchangeClientSecret],
		credentialKeyTokenExchangeRequestedTokenType: ociCreds[credentialKeyTokenExchangeRequestedTokenType],
		credentialKeyTokenExchangeSubjectTokenType:   ociCreds[credentialKeyTokenExchangeSubjectTokenType],
		credentialKeyTokenExchangeResourceType:       ociCreds[credentialKeyTokenExchangeResourceType],
		credentialKeyTokenExchangeRPSTExpiration:     ociCreds[credentialKeyTokenExchangeRPSTExpiration],
		credentialKeyTokenExchangePublicKey:          ociCreds[credentialKeyTokenExchangePublicKey],
	}

	for key, value := range config {
		if value == "" {
			delete(config, key)
		}
	}

	return config
}

func (o setupOptions) shouldConfigureSDKv2Provider(mg resource.Managed) bool {
	if o.isSDKv2Resource == nil {
		return false
	}
	tr, ok := mg.(terraformResourceTyper)
	if !ok {
		return false
	}
	return o.isSDKv2Resource(tr.GetTerraformResourceType())
}

func providerConfigurationHash(cfg map[string]any) (string, error) {
	data, err := json.Marshal(cfg)
	if err != nil {
		return "", fmt.Errorf("cannot hash OCI provider configuration: %w", err)
	}
	return fmt.Sprintf("%x", sha256.Sum256(data)), nil
}

const defaultProviderMetaCacheSize = 32

type providerMetaCacheEntry struct {
	configHash string
	meta       any
	recency    *list.Element
}

type providerMetaCache struct {
	mu         sync.Mutex
	maxEntries int
	entries    map[string]*providerMetaCacheEntry
	recency    *list.List
}

func newProviderMetaCache(maxEntries int) *providerMetaCache {
	if maxEntries < 1 {
		maxEntries = 1
	}
	return &providerMetaCache{
		maxEntries: maxEntries,
		entries:    make(map[string]*providerMetaCacheEntry, maxEntries),
		recency:    list.New(),
	}
}

func (c *providerMetaCache) getOrCreate(uid, configHash string, create func() (any, error)) (any, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if entry, ok := c.entries[uid]; ok && entry.configHash == configHash {
		c.recency.MoveToFront(entry.recency)
		return entry.meta, nil
	}

	meta, err := create()
	if err != nil {
		return nil, err
	}

	if entry, ok := c.entries[uid]; ok {
		entry.configHash = configHash
		entry.meta = meta
		c.recency.MoveToFront(entry.recency)
		return meta, nil
	}

	entry := &providerMetaCacheEntry{
		configHash: configHash,
		meta:       meta,
		recency:    c.recency.PushFront(uid),
	}
	c.entries[uid] = entry
	c.evictOverflow()
	return meta, nil
}

func (c *providerMetaCache) evictOverflow() {
	for len(c.entries) > c.maxEntries {
		oldest := c.recency.Back()
		if oldest == nil {
			return
		}
		delete(c.entries, oldest.Value.(string))
		c.recency.Remove(oldest)
	}
}

func resolveProviderConfigIdentity(ctx context.Context, kube client.Client, mg resource.Managed) (string, error) {
	switch managed := mg.(type) {
	case resource.LegacyManaged:
		return resolveLegacyProviderConfigIdentity(ctx, kube, managed)
	case resource.ModernManaged:
		if isNamespacedModernManaged(managed) {
			return resolveNamespacedProviderConfigIdentity(ctx, kube, managed)
		}
		return resolveClusterProviderConfigIdentityForModernMR(ctx, kube, managed)
	default:
		return "", errors.New(errUnsupportedManaged)
	}
}

func resolveLegacyProviderConfigIdentity(ctx context.Context, kube client.Client, mg resource.LegacyManaged) (string, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return "", errors.New(errNoProviderConfig)
	}

	pc := &clusterv1beta1.ProviderConfig{}
	if err := kube.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
		return "", errors.Wrap(err, errGetProviderConfig)
	}
	return string(pc.GetUID()), nil
}

func resolveClusterProviderConfigIdentityForModernMR(ctx context.Context, kube client.Client, mg resource.ModernManaged) (string, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil || configRef.Name == "" {
		return "", errors.New(errNoProviderConfig)
	}

	kind := configRef.Kind
	if kind == "" {
		kind = clusterv1beta1.ProviderConfigGroupVersionKind.Kind
	}
	if kind != clusterv1beta1.ProviderConfigGroupVersionKind.Kind && kind != namespacedv1beta1.ClusterProviderConfigKind {
		return "", errors.Wrap(errors.New(kind), errUnsupportedProviderCfgKind)
	}

	pc := &clusterv1beta1.ProviderConfig{}
	if err := kube.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
		return "", errors.Wrap(err, errGetProviderConfig)
	}
	return string(pc.GetUID()), nil
}

func resolveNamespacedProviderConfigIdentity(ctx context.Context, kube client.Client, mg resource.ModernManaged) (string, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil || configRef.Name == "" {
		return "", errors.New(errNoProviderConfig)
	}

	kind := configRef.Kind
	if kind == "" {
		kind = namespacedv1beta1.ClusterProviderConfigKind
	}
	switch kind {
	case namespacedv1beta1.ProviderConfigKind, namespacedv1beta1.ClusterProviderConfigKind:
	default:
		return "", errors.Wrap(errors.New(kind), errUnsupportedProviderCfgKind)
	}

	pcRuntimeObj, err := kube.Scheme().New(namespacedv1beta1.SchemeGroupVersion.WithKind(kind))
	if err != nil {
		return "", errors.Wrap(err, errUnsupportedProviderCfgKind)
	}
	pcObj, ok := pcRuntimeObj.(client.Object)
	if !ok {
		return "", errors.New(errUnsupportedProviderCfgKind)
	}

	key := types.NamespacedName{Name: configRef.Name}
	if kind == namespacedv1beta1.ProviderConfigKind {
		key.Namespace = mg.GetNamespace()
	}
	if err := kube.Get(ctx, key, pcObj); err != nil {
		return "", errors.Wrap(err, errGetProviderConfig)
	}
	return string(pcObj.GetUID()), nil
}

func resolveProviderConfig(ctx context.Context, kube client.Client, mg resource.Managed) (*namespacedv1beta1.ProviderConfigSpec, error) {
	switch managed := mg.(type) {
	case resource.LegacyManaged:
		return resolveLegacyProviderConfig(ctx, kube, managed)
	case resource.ModernManaged:
		if isNamespacedModernManaged(managed) {
			return resolveNamespacedProviderConfig(ctx, kube, managed)
		}
		return resolveClusterProviderConfigForModernMR(ctx, kube, managed)
	default:
		return nil, errors.New(errUnsupportedManaged)
	}
}

func isNamespacedModernManaged(mg resource.ModernManaged) bool {
	if mg.GetNamespace() != "" {
		return true
	}

	group := mg.GetObjectKind().GroupVersionKind().Group
	return group == namespacedv1beta1.Group || strings.HasSuffix(group, "."+namespacedv1beta1.Group)
}

func resolveLegacyProviderConfig(ctx context.Context, kube client.Client, mg resource.LegacyManaged) (*namespacedv1beta1.ProviderConfigSpec, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}

	pc := &clusterv1beta1.ProviderConfig{}
	if err := kube.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	t := resource.NewLegacyProviderConfigUsageTracker(kube, &clusterv1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}

	return toSharedPCSpec(pc.Spec)
}

func resolveClusterProviderConfigForModernMR(ctx context.Context, kube client.Client, mg resource.ModernManaged) (*namespacedv1beta1.ProviderConfigSpec, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}
	if configRef.Name == "" {
		return nil, errors.New(errNoProviderConfig)
	}

	kind := configRef.Kind
	if kind == "" {
		kind = clusterv1beta1.ProviderConfigGroupVersionKind.Kind
	}
	if kind != clusterv1beta1.ProviderConfigGroupVersionKind.Kind && kind != namespacedv1beta1.ClusterProviderConfigKind {
		return nil, errors.Wrap(errors.New(kind), errUnsupportedProviderCfgKind)
	}

	pc := &clusterv1beta1.ProviderConfig{}
	if err := kube.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	if err := trackLegacyProviderConfigUsageForModernMR(ctx, kube, mg, configRef.Name); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}

	return toSharedPCSpec(pc.Spec)
}

func resolveNamespacedProviderConfig(ctx context.Context, kube client.Client, mg resource.ModernManaged) (*namespacedv1beta1.ProviderConfigSpec, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}
	if configRef.Name == "" {
		return nil, errors.New(errNoProviderConfig)
	}

	kind := configRef.Kind
	if kind == "" {
		kind = namespacedv1beta1.ClusterProviderConfigKind
	}
	switch kind {
	case namespacedv1beta1.ProviderConfigKind, namespacedv1beta1.ClusterProviderConfigKind:
	default:
		return nil, errors.Wrap(errors.New(kind), errUnsupportedProviderCfgKind)
	}

	if configRef.Kind != kind {
		mg.SetProviderConfigReference(&xpv1.ProviderConfigReference{Name: configRef.Name, Kind: kind})
	}

	pcRuntimeObj, err := kube.Scheme().New(namespacedv1beta1.SchemeGroupVersion.WithKind(kind))
	if err != nil {
		return nil, errors.Wrap(err, errUnsupportedProviderCfgKind)
	}
	pcObj, ok := pcRuntimeObj.(client.Object)
	if !ok {
		return nil, errors.New(errUnsupportedProviderCfgKind)
	}

	key := types.NamespacedName{Name: configRef.Name}
	if kind == namespacedv1beta1.ProviderConfigKind {
		key.Namespace = mg.GetNamespace()
	}
	if err := kube.Get(ctx, key, pcObj); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	var pcSpec namespacedv1beta1.ProviderConfigSpec
	switch pc := pcObj.(type) {
	case *namespacedv1beta1.ProviderConfig:
		pcSpec = pc.Spec
		if pcSpec.Credentials.SecretRef != nil {
			pcSpec.Credentials.SecretRef.Namespace = mg.GetNamespace()
		}
	case *namespacedv1beta1.ClusterProviderConfig:
		pcSpec = pc.Spec
	default:
		return nil, errors.New(errUnsupportedProviderCfgKind)
	}

	t := resource.NewProviderConfigUsageTracker(kube, &namespacedv1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}

	return &pcSpec, nil
}

func toSharedPCSpec(spec any) (*namespacedv1beta1.ProviderConfigSpec, error) {
	data, err := json.Marshal(spec)
	if err != nil {
		return nil, err
	}
	out := &namespacedv1beta1.ProviderConfigSpec{}
	if err := json.Unmarshal(data, out); err != nil {
		return nil, err
	}
	return out, nil
}

func trackLegacyProviderConfigUsageForModernMR(ctx context.Context, kube client.Client, mg resource.ModernManaged, providerConfigName string) error {
	pcu := &clusterv1beta1.ProviderConfigUsage{}
	gvk := mg.GetObjectKind().GroupVersionKind()

	pcu.SetName(string(mg.GetUID()))
	pcu.SetLabels(map[string]string{xpv1.LabelKeyProviderName: providerConfigName})
	pcu.SetOwnerReferences([]metav1.OwnerReference{meta.AsController(meta.TypedReferenceTo(mg, gvk))})
	pcu.SetProviderConfigReference(xpv1.Reference{Name: providerConfigName})
	pcu.SetResourceReference(xpv1.TypedReference{
		APIVersion: gvk.GroupVersion().String(),
		Kind:       gvk.Kind,
		Name:       mg.GetName(),
	})

	err := resource.NewAPIUpdatingApplicator(kube).Apply(ctx, pcu,
		resource.MustBeControllableBy(mg.GetUID()),
		resource.AllowUpdateIf(func(current, _ kruntime.Object) bool {
			return current.(*clusterv1beta1.ProviderConfigUsage).GetProviderConfigReference() != pcu.GetProviderConfigReference()
		}),
	)
	return errors.Wrap(resource.Ignore(resource.IsNotAllowed, err), errTrackUsage)
}

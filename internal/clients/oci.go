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
	"time"

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
	credentialKeyPrivateKeyPassword              = "private_key_password"
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
	providerMetaCacheSize   int
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

// WithProviderMetaCacheSize controls the maximum number of configured SDKv2
// provider instances retained by a service process.
func WithProviderMetaCacheSize(size int) SetupOption {
	return func(o *setupOptions) {
		o.providerMetaCacheSize = size
	}
}

func newSetupOptions(opts ...SetupOption) setupOptions {
	options := setupOptions{providerMetaCacheSize: defaultProviderMetaCacheSize}
	for _, opt := range opts {
		opt(&options)
	}
	return options
}

type terraformResourceTyper interface {
	GetTerraformResourceType() string
}

// TerraformSetupBuilder builds a terraform.SetupFn for in-process no-fork
// connectors. Build-time Terraform values are intentionally not required at
// runtime when all resources are routed through SDKv2 or Framework connectors.
func TerraformSetupBuilder(opts ...SetupOption) upjetterraform.SetupFn {
	options := newSetupOptions(opts...)
	providerMetaCache := newProviderMetaCache(options.providerMetaCacheSize)

	return func(ctx context.Context, kube client.Client, mg resource.Managed) (upjetterraform.Setup, error) {
		ps := upjetterraform.Setup{}

		pc, err := resolveProviderConfig(ctx, kube, mg)
		if err != nil {
			return ps, errors.Wrap(err, "cannot resolve provider config")
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.spec.Credentials.Source, kube, pc.spec.Credentials.CommonCredentialSelectors)
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

		if pc.uid == "" {
			return ps, fmt.Errorf("ProviderConfig has empty UID")
		}

		providerMeta, err := providerMetaCache.getOrConfigureProviderMeta(ctx, pc.uid, cfg)
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
		credentialKeyPrivateKeyPassword:              ociCreds[credentialKeyPrivateKeyPassword],
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

type providerMetaCacheCall struct {
	configHash string
	meta       any
	err        error
	done       chan struct{}
}

type providerMetaCache struct {
	mu         sync.Mutex
	maxEntries int
	entries    map[string]*providerMetaCacheEntry
	recency    *list.List
	inflight   map[string]*providerMetaCacheCall
	metrics    *providerMetaCacheMetrics
}

func newProviderMetaCache(maxEntries int) *providerMetaCache {
	return newProviderMetaCacheWithMetrics(maxEntries, defaultProviderMetaCacheMetrics)
}

func newProviderMetaCacheWithMetrics(maxEntries int, metrics *providerMetaCacheMetrics) *providerMetaCache {
	if maxEntries < 1 {
		maxEntries = 1
	}
	metrics.capacity.Set(float64(maxEntries))
	metrics.entries.Set(0)
	metrics.inflight.Set(0)
	metrics.waiting.Set(0)
	return &providerMetaCache{
		maxEntries: maxEntries,
		entries:    make(map[string]*providerMetaCacheEntry, maxEntries),
		recency:    list.New(),
		inflight:   make(map[string]*providerMetaCacheCall),
		metrics:    metrics,
	}
}

func (c *providerMetaCache) getOrCreate(ctx context.Context, uid, configHash string, create func() (any, error)) (any, error) {
	for {
		c.mu.Lock()
		if entry, ok := c.entries[uid]; ok && entry.configHash == configHash {
			c.recency.MoveToFront(entry.recency)
			c.metrics.hits.Inc()
			meta := entry.meta
			c.mu.Unlock()
			return meta, nil
		}

		if call, ok := c.inflight[uid]; ok {
			sameConfiguration := call.configHash == configHash
			c.mu.Unlock()
			c.metrics.waiting.Inc()
			select {
			case <-call.done:
			case <-ctx.Done():
				c.metrics.waiting.Dec()
				c.metrics.waits.WithLabelValues("canceled").Inc()
				return nil, ctx.Err()
			}
			c.metrics.waiting.Dec()
			if sameConfiguration && call.err != nil {
				c.metrics.waits.WithLabelValues("initialization_error").Inc()
				return nil, call.err
			}
			result := "completed"
			if !sameConfiguration {
				result = "configuration_changed"
			}
			c.metrics.waits.WithLabelValues(result).Inc()
			continue
		}

		missReason := "absent"
		if _, ok := c.entries[uid]; ok {
			missReason = "configuration_changed"
		}
		c.metrics.misses.WithLabelValues(missReason).Inc()
		call := &providerMetaCacheCall{configHash: configHash, done: make(chan struct{})}
		c.inflight[uid] = call
		c.metrics.inflight.Set(float64(len(c.inflight)))
		c.mu.Unlock()

		start := time.Now()
		call.meta, call.err = initializeProviderMeta(create)
		result := "success"
		if call.err != nil {
			result = "error"
		}
		c.metrics.initializations.WithLabelValues(result).Inc()
		c.metrics.initializationDuration.WithLabelValues(result).Observe(time.Since(start).Seconds())

		c.mu.Lock()
		if call.err == nil {
			if entry, ok := c.entries[uid]; ok {
				entry.configHash = configHash
				entry.meta = call.meta
				c.recency.MoveToFront(entry.recency)
			} else {
				c.entries[uid] = &providerMetaCacheEntry{
					configHash: configHash,
					meta:       call.meta,
					recency:    c.recency.PushFront(uid),
				}
			}
			c.evictOverflow()
			c.metrics.entries.Set(float64(len(c.entries)))
		}
		delete(c.inflight, uid)
		c.metrics.inflight.Set(float64(len(c.inflight)))
		close(call.done)
		c.mu.Unlock()
		return call.meta, call.err
	}
}

func initializeProviderMeta(create func() (any, error)) (meta any, err error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			meta = nil
			err = fmt.Errorf("provider metadata initialization panicked: %v", recovered)
		}
	}()
	return create()
}

func (c *providerMetaCache) evictOverflow() {
	for len(c.entries) > c.maxEntries {
		oldest := c.recency.Back()
		if oldest == nil {
			return
		}
		delete(c.entries, oldest.Value.(string))
		c.recency.Remove(oldest)
		c.metrics.evictions.Inc()
	}
}

type resolvedProviderConfig struct {
	spec *namespacedv1beta1.ProviderConfigSpec
	uid  string
}

func resolveProviderConfig(ctx context.Context, kube client.Client, mg resource.Managed) (resolvedProviderConfig, error) {
	switch managed := mg.(type) {
	case resource.LegacyManaged:
		return resolveLegacyProviderConfig(ctx, kube, managed)
	case resource.ModernManaged:
		if isNamespacedModernManaged(managed) {
			return resolveNamespacedProviderConfig(ctx, kube, managed)
		}
		return resolveClusterProviderConfigForModernMR(ctx, kube, managed)
	default:
		return resolvedProviderConfig{}, errors.New(errUnsupportedManaged)
	}
}

func isNamespacedModernManaged(mg resource.ModernManaged) bool {
	if mg.GetNamespace() != "" {
		return true
	}

	group := mg.GetObjectKind().GroupVersionKind().Group
	return group == namespacedv1beta1.Group || strings.HasSuffix(group, "."+namespacedv1beta1.Group)
}

func resolveLegacyProviderConfig(ctx context.Context, kube client.Client, mg resource.LegacyManaged) (resolvedProviderConfig, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return resolvedProviderConfig{}, errors.New(errNoProviderConfig)
	}

	pc := &clusterv1beta1.ProviderConfig{}
	if err := kube.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
		return resolvedProviderConfig{}, errors.Wrap(err, errGetProviderConfig)
	}

	t := resource.NewLegacyProviderConfigUsageTracker(kube, &clusterv1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return resolvedProviderConfig{}, errors.Wrap(err, errTrackUsage)
	}

	spec, err := toSharedPCSpec(pc.Spec)
	return resolvedProviderConfig{spec: spec, uid: string(pc.GetUID())}, err
}

func resolveClusterProviderConfigForModernMR(ctx context.Context, kube client.Client, mg resource.ModernManaged) (resolvedProviderConfig, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return resolvedProviderConfig{}, errors.New(errNoProviderConfig)
	}
	if configRef.Name == "" {
		return resolvedProviderConfig{}, errors.New(errNoProviderConfig)
	}

	kind := configRef.Kind
	if kind == "" {
		kind = clusterv1beta1.ProviderConfigGroupVersionKind.Kind
	}
	if kind != clusterv1beta1.ProviderConfigGroupVersionKind.Kind && kind != namespacedv1beta1.ClusterProviderConfigKind {
		return resolvedProviderConfig{}, errors.Wrap(errors.New(kind), errUnsupportedProviderCfgKind)
	}

	pc := &clusterv1beta1.ProviderConfig{}
	if err := kube.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
		return resolvedProviderConfig{}, errors.Wrap(err, errGetProviderConfig)
	}

	if err := trackLegacyProviderConfigUsageForModernMR(ctx, kube, mg, configRef.Name); err != nil {
		return resolvedProviderConfig{}, errors.Wrap(err, errTrackUsage)
	}

	spec, err := toSharedPCSpec(pc.Spec)
	return resolvedProviderConfig{spec: spec, uid: string(pc.GetUID())}, err
}

func resolveNamespacedProviderConfig(ctx context.Context, kube client.Client, mg resource.ModernManaged) (resolvedProviderConfig, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return resolvedProviderConfig{}, errors.New(errNoProviderConfig)
	}
	if configRef.Name == "" {
		return resolvedProviderConfig{}, errors.New(errNoProviderConfig)
	}

	kind := configRef.Kind
	if kind == "" {
		kind = namespacedv1beta1.ClusterProviderConfigKind
	}
	switch kind {
	case namespacedv1beta1.ProviderConfigKind, namespacedv1beta1.ClusterProviderConfigKind:
	default:
		return resolvedProviderConfig{}, errors.Wrap(errors.New(kind), errUnsupportedProviderCfgKind)
	}

	if configRef.Kind != kind {
		mg.SetProviderConfigReference(&xpv1.ProviderConfigReference{Name: configRef.Name, Kind: kind})
	}

	pcRuntimeObj, err := kube.Scheme().New(namespacedv1beta1.SchemeGroupVersion.WithKind(kind))
	if err != nil {
		return resolvedProviderConfig{}, errors.Wrap(err, errUnsupportedProviderCfgKind)
	}
	pcObj, ok := pcRuntimeObj.(client.Object)
	if !ok {
		return resolvedProviderConfig{}, errors.New(errUnsupportedProviderCfgKind)
	}

	key := types.NamespacedName{Name: configRef.Name}
	if kind == namespacedv1beta1.ProviderConfigKind {
		key.Namespace = mg.GetNamespace()
	}
	if err := kube.Get(ctx, key, pcObj); err != nil {
		return resolvedProviderConfig{}, errors.Wrap(err, errGetProviderConfig)
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
		return resolvedProviderConfig{}, errors.New(errUnsupportedProviderCfgKind)
	}

	t := resource.NewProviderConfigUsageTracker(kube, &namespacedv1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return resolvedProviderConfig{}, errors.Wrap(err, errTrackUsage)
	}

	return resolvedProviderConfig{spec: &pcSpec, uid: string(pcObj.GetUID())}, nil
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

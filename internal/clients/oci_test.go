/*
Copyright 2026 Oracle and/or its affiliates.
*/

package clients

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource/fake"
	"github.com/prometheus/client_golang/prometheus/testutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	fakeclient "sigs.k8s.io/controller-runtime/pkg/client/fake"

	namespacedv1beta1 "github.com/oracle/provider-oci/apis/namespaced/v1beta1"
)

type typedManaged struct {
	fake.Managed
	tfType string
}

type providerConfigGetCountingClient struct {
	client.Client
	gets atomic.Int32
}

func (c *providerConfigGetCountingClient) Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	if _, ok := obj.(*namespacedv1beta1.ProviderConfig); ok {
		c.gets.Add(1)
	}
	return c.Client.Get(ctx, key, obj, opts...)
}

func (m *typedManaged) GetTerraformResourceType() string {
	return m.tfType
}

func TestShouldConfigureSDKv2Provider(t *testing.T) {
	tests := map[string]struct {
		options setupOptions
		mg      *typedManaged
		want    bool
	}{
		"default skips SDKv2 provider setup": {
			options: setupOptions{},
			mg:      &typedManaged{tfType: "oci_budget_budget"},
			want:    false,
		},
		"predicate match enables SDKv2 provider setup": {
			options: setupOptions{isSDKv2Resource: func(name string) bool {
				return name == "oci_budget_budget"
			}},
			mg:   &typedManaged{tfType: "oci_budget_budget"},
			want: true,
		},
		"predicate miss skips SDKv2 provider setup": {
			options: setupOptions{isSDKv2Resource: func(name string) bool {
				return name == "oci_budget_budget"
			}},
			mg:   &typedManaged{tfType: "oci_objectstorage_bucket"},
			want: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.options.shouldConfigureSDKv2Provider(tc.mg)
			if got != tc.want {
				t.Fatalf("shouldConfigureSDKv2Provider() = %t, want %t", got, tc.want)
			}
		})
	}
}

func TestShouldConfigureSDKv2ProviderSkipsUnknownResourceType(t *testing.T) {
	options := setupOptions{isSDKv2Resource: func(string) bool {
		return true
	}}
	if options.shouldConfigureSDKv2Provider(&fake.Managed{}) {
		t.Fatal("shouldConfigureSDKv2Provider() = true, want false for managed resource without Terraform resource type")
	}
}

func TestSetupOptionsProviderMetaCacheSize(t *testing.T) {
	tests := map[string]struct {
		opts []SetupOption
		want int
	}{
		"default": {
			want: defaultProviderMetaCacheSize,
		},
		"custom": {
			opts: []SetupOption{WithProviderMetaCacheSize(7)},
			want: 7,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if got := newSetupOptions(tc.opts...).providerMetaCacheSize; got != tc.want {
				t.Fatalf("providerMetaCacheSize = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestResolveProviderConfigReturnsSpecAndIdentityFromOneFetch(t *testing.T) {
	scheme := runtime.NewScheme()
	if err := namespacedv1beta1.SchemeBuilder.AddToScheme(scheme); err != nil {
		t.Fatal(err)
	}
	pc := &namespacedv1beta1.ProviderConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "default",
			Namespace: "testing",
			UID:       types.UID("provider-config-uid"),
		},
		Spec: namespacedv1beta1.ProviderConfigSpec{
			Credentials: namespacedv1beta1.ProviderCredentials{Source: xpv1.CredentialsSourceInjectedIdentity},
		},
	}
	kube := &providerConfigGetCountingClient{Client: fakeclient.NewClientBuilder().WithScheme(scheme).WithObjects(pc).Build()}
	mg := &fake.ModernManaged{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "managed",
			Namespace: "testing",
			UID:       types.UID("managed-uid"),
		},
		TypedProviderConfigReferencer: fake.TypedProviderConfigReferencer{Ref: &xpv1.ProviderConfigReference{
			Name: "default",
			Kind: namespacedv1beta1.ProviderConfigKind,
		}},
	}

	resolved, err := resolveProviderConfig(t.Context(), kube, mg)
	if err != nil {
		t.Fatalf("resolveProviderConfig() error: %v", err)
	}
	if resolved.uid != "provider-config-uid" {
		t.Fatalf("resolved UID = %q, want provider-config-uid", resolved.uid)
	}
	if resolved.spec == nil || resolved.spec.Credentials.Source != xpv1.CredentialsSourceInjectedIdentity {
		t.Fatalf("resolved spec = %#v, want injected identity credentials", resolved.spec)
	}
	if got := kube.gets.Load(); got != 1 {
		t.Fatalf("ProviderConfig GET count = %d, want 1", got)
	}
}

func TestNewProviderMetaCacheUsesMinimumSize(t *testing.T) {
	for _, size := range []int{-1, 0} {
		t.Run(fmt.Sprintf("size_%d", size), func(t *testing.T) {
			if got := newProviderMetaCache(size).maxEntries; got != 1 {
				t.Fatalf("maxEntries = %d, want 1", got)
			}
		})
	}
}

func TestProviderConfigurationFromCredentialsIncludesOnlyNoForkSafeKeys(t *testing.T) {
	creds := map[string]string{
		"tenancy_ocid":         "tenancy",
		"user_ocid":            "user",
		"private_key":          "key",
		"private_key_password": "password",
		"private_key_path":     "path",
		"fingerprint":          "fingerprint",
		"region":               "us-ashburn-1",
		"auth":                 "api_key",
		"config_file_profile":  "DEFAULT",
		"disable_auto_retries": "true",
	}

	cfg := providerConfigurationFromCredentials(creds)
	wantKeys := map[string]string{
		"tenancy_ocid":         "tenancy",
		"user_ocid":            "user",
		"private_key":          "key",
		"private_key_password": "password",
		"private_key_path":     "path",
		"fingerprint":          "fingerprint",
		"region":               "us-ashburn-1",
		"auth":                 "api_key",
		"config_file_profile":  "DEFAULT",
	}

	if len(cfg) != len(wantKeys) {
		t.Fatalf("providerConfigurationFromCredentials() returned %d keys, want %d: %v", len(cfg), len(wantKeys), cfg)
	}
	for key, want := range wantKeys {
		if got := cfg[key]; got != want {
			t.Fatalf("providerConfigurationFromCredentials()[%q] = %v, want %q", key, got, want)
		}
	}

	unsafeProviderGlobalKeys := []string{
		"disable_auto_retries",
		"retry_duration_seconds",
		"retries_config_file",
		"ignore_defined_tags",
		"realm_specific_service_endpoint_template_enabled",
		"dual_stack_endpoint_enabled",
	}
	for _, key := range unsafeProviderGlobalKeys {
		if _, ok := cfg[key]; ok {
			t.Fatalf("providerConfigurationFromCredentials() included unsafe no-fork provider-global key %q", key)
		}
	}
}

func TestProviderConfigurationFromCredentialsSupportsAuthenticationModes(t *testing.T) {
	tests := map[string]map[string]string{
		"encrypted API key": {
			credentialKeyTenancyOCID:        "ocid1.tenancy.oc1..example",
			credentialKeyUserOCID:           "ocid1.user.oc1..example",
			credentialKeyPrivateKey:         "encrypted-key",
			credentialKeyPrivateKeyPassword: "password",
			credentialKeyFingerprint:        "fingerprint",
			credentialKeyAuth:               "api_key",
			credentialKeyRegion:             "us-ashburn-1",
		},
		"workload identity federation": {
			credentialKeyTenancyOCID:                     "ocid1.tenancy.oc1..example",
			credentialKeyAuth:                            "WorkloadIdentityFederation",
			credentialKeyRegion:                          "us-ashburn-1",
			credentialKeyWorkloadIdentityTokenPath:       "/var/run/secrets/tokens/oci",
			credentialKeyTokenExchangeDomainURL:          "https://idcs.example.com",
			credentialKeyTokenExchangeAuth:               "OAuthClientCredentials",
			credentialKeyTokenExchangeClientID:           "client-id",
			credentialKeyTokenExchangeClientSecret:       "client-secret",
			credentialKeyTokenExchangeRequestedTokenType: "urn:oci:token-type:oci-rpst",
			credentialKeyTokenExchangeSubjectTokenType:   "jwt",
			credentialKeyTokenExchangeResourceType:       "k8sworkload",
			credentialKeyTokenExchangeRPSTExpiration:     "3600",
			credentialKeyTokenExchangePublicKey:          "public-key",
		},
		"instance principal": {
			credentialKeyTenancyOCID: "ocid1.tenancy.oc1..example",
			credentialKeyAuth:        "InstancePrincipal",
			credentialKeyRegion:      "us-ashburn-1",
		},
		"OKE workload identity": {
			credentialKeyTenancyOCID: "ocid1.tenancy.oc1..example",
			credentialKeyAuth:        "OKEWorkloadIdentity",
			credentialKeyRegion:      "us-ashburn-1",
		},
	}

	for name, creds := range tests {
		t.Run(name, func(t *testing.T) {
			cfg := providerConfigurationFromCredentials(creds)
			if len(cfg) != len(creds) {
				t.Fatalf("configuration contains %d keys, want %d: %v", len(cfg), len(creds), cfg)
			}
			for key, want := range creds {
				if got, ok := cfg[key]; !ok || got != want {
					t.Fatalf("configuration[%q] = %v, %t, want %q, true", key, got, ok, want)
				}
			}
			if _, ok := cfg[credentialKeyTokenExchangeAuth]; name != "workload identity federation" && ok {
				t.Fatalf("configuration unexpectedly includes %q", credentialKeyTokenExchangeAuth)
			}
		})
	}
}

func TestProviderConfigurationHash(t *testing.T) {
	base := map[string]any{
		"region":       "us-ashburn-1",
		"tenancy_ocid": "tenancy-a",
	}
	reordered := map[string]any{
		"tenancy_ocid": "tenancy-a",
		"region":       "us-ashburn-1",
	}
	rotated := map[string]any{
		"region":       "us-ashburn-1",
		"tenancy_ocid": "tenancy-b",
	}

	baseHash, err := providerConfigurationHash(base)
	if err != nil {
		t.Fatalf("providerConfigurationHash(base) error: %v", err)
	}
	reorderedHash, err := providerConfigurationHash(reordered)
	if err != nil {
		t.Fatalf("providerConfigurationHash(reordered) error: %v", err)
	}
	rotatedHash, err := providerConfigurationHash(rotated)
	if err != nil {
		t.Fatalf("providerConfigurationHash(rotated) error: %v", err)
	}

	if baseHash != reorderedHash {
		t.Fatalf("providerConfigurationHash() changed with map order: %q != %q", baseHash, reorderedHash)
	}
	if baseHash == rotatedHash {
		t.Fatal("providerConfigurationHash() did not change after provider configuration changed")
	}
}

func TestProviderMetaCacheReusesUnchangedConfiguration(t *testing.T) {
	cache := newProviderMetaCache(2)
	var creates atomic.Int32
	create := func() (any, error) {
		return creates.Add(1), nil
	}

	first, err := cache.getOrCreate(t.Context(), "uid-a", "hash-a", create)
	if err != nil {
		t.Fatal(err)
	}
	second, err := cache.getOrCreate(t.Context(), "uid-a", "hash-a", create)
	if err != nil {
		t.Fatal(err)
	}
	if first != second {
		t.Fatalf("cache returned different metadata: %v != %v", first, second)
	}
	if got := creates.Load(); got != 1 {
		t.Fatalf("provider creation count = %d, want 1", got)
	}
}

func TestProviderMetaCacheReplacesChangedConfiguration(t *testing.T) {
	cache := newProviderMetaCache(2)
	var creates atomic.Int32
	create := func() (any, error) {
		return creates.Add(1), nil
	}

	first, err := cache.getOrCreate(t.Context(), "uid-a", "hash-a", create)
	if err != nil {
		t.Fatal(err)
	}
	second, err := cache.getOrCreate(t.Context(), "uid-a", "hash-b", create)
	if err != nil {
		t.Fatal(err)
	}
	if first == second {
		t.Fatalf("cache reused metadata after configuration change: %v", first)
	}
	if len(cache.entries) != 1 {
		t.Fatalf("cache entry count = %d, want 1", len(cache.entries))
	}
}

func TestProviderMetaCachePreservesPreviousEntryWhenCreationFails(t *testing.T) {
	cache := newProviderMetaCache(2)
	want := new(int)
	if _, err := cache.getOrCreate(t.Context(), "uid-a", "hash-a", func() (any, error) {
		return want, nil
	}); err != nil {
		t.Fatal(err)
	}

	wantErr := errors.New("configure failed")
	if _, err := cache.getOrCreate(t.Context(), "uid-a", "hash-b", func() (any, error) {
		return nil, wantErr
	}); !errors.Is(err, wantErr) {
		t.Fatalf("getOrCreate() error = %v, want %v", err, wantErr)
	}
	entry := cache.entries["uid-a"]
	if entry.configHash != "hash-a" || entry.meta != want {
		t.Fatalf("previous entry changed after failed creation: %#v", entry)
	}
}

func TestProviderMetaCacheRecoversFromInitializationPanic(t *testing.T) {
	cache := newProviderMetaCache(1)
	_, err := cache.getOrCreate(t.Context(), "uid-a", "hash-a", func() (any, error) {
		panic("constructor panic")
	})
	if err == nil || !strings.Contains(err.Error(), "provider metadata initialization panicked: constructor panic") {
		t.Fatalf("getOrCreate() error = %v, want recovered initialization panic", err)
	}
	if len(cache.inflight) != 0 {
		t.Fatalf("inflight call count = %d, want 0", len(cache.inflight))
	}

	meta, err := cache.getOrCreate(t.Context(), "uid-a", "hash-a", func() (any, error) {
		return "recovered", nil
	})
	if err != nil {
		t.Fatalf("retry after panic failed: %v", err)
	}
	if meta != "recovered" {
		t.Fatalf("retry metadata = %v, want recovered", meta)
	}
}

func TestProviderMetaCacheDoesNotShareFailureAcrossConfigurationChanges(t *testing.T) {
	cache := newProviderMetaCache(1)
	started := make(chan struct{})
	release := make(chan struct{})
	oldErr := errors.New("old configuration failed")
	oldDone := make(chan error, 1)
	go func() {
		_, err := cache.getOrCreate(t.Context(), "uid-a", "hash-old", func() (any, error) {
			close(started)
			<-release
			return nil, oldErr
		})
		oldDone <- err
	}()
	<-started

	newStarted := make(chan struct{})
	newDone := make(chan struct {
		meta any
		err  error
	}, 1)
	go func() {
		meta, err := cache.getOrCreate(t.Context(), "uid-a", "hash-new", func() (any, error) {
			close(newStarted)
			return "new metadata", nil
		})
		newDone <- struct {
			meta any
			err  error
		}{meta: meta, err: err}
	}()

	deadline := time.Now().Add(time.Second)
	for testutil.ToFloat64(cache.metrics.waiting) != 1 && time.Now().Before(deadline) {
		time.Sleep(time.Millisecond)
	}
	if got := testutil.ToFloat64(cache.metrics.waiting); got != 1 {
		t.Fatalf("waiting calls = %v, want 1", got)
	}
	select {
	case <-newStarted:
		t.Fatal("new configuration initialized before the prior call completed")
	default:
	}
	close(release)
	if err := <-oldDone; !errors.Is(err, oldErr) {
		t.Fatalf("old configuration error = %v, want %v", err, oldErr)
	}
	result := <-newDone
	if result.err != nil {
		t.Fatalf("new configuration inherited old error: %v", result.err)
	}
	if result.meta != "new metadata" {
		t.Fatalf("new configuration metadata = %v, want new metadata", result.meta)
	}
	if entry := cache.entries["uid-a"]; entry == nil || entry.configHash != "hash-new" || entry.meta != "new metadata" {
		t.Fatalf("cache entry = %#v, want new configuration metadata", entry)
	}
}

func TestProviderMetaCacheEvictsLeastRecentlyUsedEntry(t *testing.T) {
	cache := newProviderMetaCache(2)
	create := func(value string) func() (any, error) {
		return func() (any, error) { return value, nil }
	}

	if _, err := cache.getOrCreate(t.Context(), "uid-a", "hash-a", create("a")); err != nil {
		t.Fatal(err)
	}
	if _, err := cache.getOrCreate(t.Context(), "uid-b", "hash-b", create("b")); err != nil {
		t.Fatal(err)
	}
	if _, err := cache.getOrCreate(t.Context(), "uid-a", "hash-a", create("unused")); err != nil {
		t.Fatal(err)
	}
	if _, err := cache.getOrCreate(t.Context(), "uid-c", "hash-c", create("c")); err != nil {
		t.Fatal(err)
	}

	if _, ok := cache.entries["uid-b"]; ok {
		t.Fatal("least recently used entry uid-b was not evicted")
	}
	if len(cache.entries) != 2 {
		t.Fatalf("cache entry count = %d, want 2", len(cache.entries))
	}
}

func TestProviderMetaCacheCoalescesConcurrentCreation(t *testing.T) {
	cache := newProviderMetaCache(2)
	var creates atomic.Int32
	create := func() (any, error) {
		return creates.Add(1), nil
	}

	const callers = 16
	var wg sync.WaitGroup
	results := make(chan any, callers)
	for range callers {
		wg.Go(func() {
			meta, err := cache.getOrCreate(t.Context(), "uid-a", "hash-a", create)
			if err != nil {
				t.Errorf("getOrCreate() error: %v", err)
				return
			}
			results <- meta
		})
	}
	wg.Wait()
	close(results)

	for meta := range results {
		if meta != int32(1) {
			t.Fatalf("cache metadata = %v, want 1", meta)
		}
	}
	if got := creates.Load(); got != 1 {
		t.Fatalf("provider creation count = %d, want 1", got)
	}
}

func TestProviderMetaCacheConstructsDifferentProviderConfigsConcurrently(t *testing.T) {
	cache := newProviderMetaCache(2)
	started := make(chan string, 2)
	release := make(chan struct{})
	create := func(uid string) func() (any, error) {
		return func() (any, error) {
			started <- uid
			<-release
			return uid, nil
		}
	}

	results := make(chan string, 2)
	var wg sync.WaitGroup
	for _, uid := range []string{"uid-a", "uid-b"} {
		wg.Go(func() {
			meta, err := cache.getOrCreate(t.Context(), uid, "hash-"+uid, create(uid))
			if err != nil {
				t.Errorf("getOrCreate(%q) error: %v", uid, err)
				return
			}
			results <- meta.(string)
		})
	}

	for range 2 {
		select {
		case <-started:
		case <-time.After(time.Second):
			t.Fatal("provider configuration was serialized by the global cache lock")
		}
	}
	close(release)
	wg.Wait()
	close(results)

	got := map[string]bool{}
	for result := range results {
		got[result] = true
	}
	if !got["uid-a"] || !got["uid-b"] {
		t.Fatalf("results = %v, want both ProviderConfig values", got)
	}
}

func TestProviderMetaCacheReturnsMetadataForValidatedConfiguration(t *testing.T) {
	type providerMeta struct {
		configHash string
	}

	cache := newProviderMetaCache(1)
	const (
		callers    = 8
		iterations = 250
	)

	start := make(chan struct{})
	var wg sync.WaitGroup
	for caller := range callers {
		wg.Go(func() {
			<-start
			for iteration := range iterations {
				configHash := fmt.Sprintf("hash-%d", (caller+iteration)%2)
				meta, err := cache.getOrCreate(t.Context(), "uid-a", configHash, func() (any, error) {
					return &providerMeta{configHash: configHash}, nil
				})
				if err != nil {
					t.Errorf("getOrCreate(%q) error: %v", configHash, err)
					return
				}
				if got := meta.(*providerMeta).configHash; got != configHash {
					t.Errorf("getOrCreate(%q) returned metadata for %q", configHash, got)
					return
				}
			}
		})
	}
	close(start)
	wg.Wait()
}

func TestProviderMetaCacheWaiterHonorsContextCancellation(t *testing.T) {
	cache := newProviderMetaCache(1)
	started := make(chan struct{})
	release := make(chan struct{})
	creatorDone := make(chan error, 1)
	go func() {
		_, err := cache.getOrCreate(context.Background(), "uid-a", "hash-a", func() (any, error) {
			close(started)
			<-release
			return "meta", nil
		})
		creatorDone <- err
	}()
	<-started

	waiterCtx, cancel := context.WithCancel(t.Context())
	cancel()
	_, err := cache.getOrCreate(waiterCtx, "uid-a", "hash-a", func() (any, error) {
		t.Fatal("waiter unexpectedly configured duplicate provider metadata")
		return nil, nil
	})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("waiter error = %v, want context.Canceled", err)
	}

	close(release)
	if err := <-creatorDone; err != nil {
		t.Fatalf("creator returned error: %v", err)
	}
}

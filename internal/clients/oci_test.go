/*
Copyright 2026 Oracle and/or its affiliates.
*/

package clients

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource/fake"
)

type typedManaged struct {
	fake.Managed
	tfType string
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

func TestProviderConfigurationFromCredentialsIncludesOnlyNoForkSafeKeys(t *testing.T) {
	creds := map[string]string{
		"tenancy_ocid":         "tenancy",
		"user_ocid":            "user",
		"private_key":          "key",
		"private_key_path":     "path",
		"fingerprint":          "fingerprint",
		"region":               "us-ashburn-1",
		"auth":                 "api_key",
		"config_file_profile":  "DEFAULT",
		"disable_auto_retries": "true",
	}

	cfg := providerConfigurationFromCredentials(creds)
	wantKeys := map[string]string{
		"tenancy_ocid":        "tenancy",
		"user_ocid":           "user",
		"private_key":         "key",
		"private_key_path":    "path",
		"fingerprint":         "fingerprint",
		"region":              "us-ashburn-1",
		"auth":                "api_key",
		"config_file_profile": "DEFAULT",
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

	first, err := cache.getOrCreate("uid-a", "hash-a", create)
	if err != nil {
		t.Fatal(err)
	}
	second, err := cache.getOrCreate("uid-a", "hash-a", create)
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

	first, err := cache.getOrCreate("uid-a", "hash-a", create)
	if err != nil {
		t.Fatal(err)
	}
	second, err := cache.getOrCreate("uid-a", "hash-b", create)
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
	if _, err := cache.getOrCreate("uid-a", "hash-a", func() (any, error) {
		return want, nil
	}); err != nil {
		t.Fatal(err)
	}

	wantErr := errors.New("configure failed")
	if _, err := cache.getOrCreate("uid-a", "hash-b", func() (any, error) {
		return nil, wantErr
	}); !errors.Is(err, wantErr) {
		t.Fatalf("getOrCreate() error = %v, want %v", err, wantErr)
	}
	entry := cache.entries["uid-a"]
	if entry.configHash != "hash-a" || entry.meta != want {
		t.Fatalf("previous entry changed after failed creation: %#v", entry)
	}
}

func TestProviderMetaCacheEvictsLeastRecentlyUsedEntry(t *testing.T) {
	cache := newProviderMetaCache(2)
	create := func(value string) func() (any, error) {
		return func() (any, error) { return value, nil }
	}

	if _, err := cache.getOrCreate("uid-a", "hash-a", create("a")); err != nil {
		t.Fatal(err)
	}
	if _, err := cache.getOrCreate("uid-b", "hash-b", create("b")); err != nil {
		t.Fatal(err)
	}
	if _, err := cache.getOrCreate("uid-a", "hash-a", create("unused")); err != nil {
		t.Fatal(err)
	}
	if _, err := cache.getOrCreate("uid-c", "hash-c", create("c")); err != nil {
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
			meta, err := cache.getOrCreate("uid-a", "hash-a", create)
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

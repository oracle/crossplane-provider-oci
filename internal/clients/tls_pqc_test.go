/*
Copyright 2026 Oracle and/or its affiliates.
*/

package clients

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/objectstorage"
)

// Verify the actual embedded Terraform provider transport, rather than a
// separately constructed OCI SDK client: the provider supplies its own HTTP
// client and custom-CA configuration to the lazily initialized SDK clients.
func TestInProcessProviderTLSNegotiation(t *testing.T) {
	clearFileCredentialEnvironment(t)
	for _, name := range []string{
		"custom_cert_location", "accept_local_certs", "OCI_DEFAULT_CERTS_PATH",
		"OCI_DEFAULT_CLIENT_CERTS_PATH", "OCI_DEFAULT_CLIENT_CERTS_PRIVATE_KEY_PATH",
	} {
		t.Setenv(name, "")
	}
	configPath := filepath.Join(t.TempDir(), "config")
	if err := os.WriteFile(configPath, []byte("[DEFAULT]\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("OCI_CONFIG_FILE", configPath)
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	cfg := map[string]any{
		credentialKeyAuth:        "ApiKey",
		credentialKeyTenancyOCID: "ocid1.tenancy.oc1..pqc-test",
		credentialKeyUserOCID:    "ocid1.user.oc1..pqc-test",
		credentialKeyPrivateKey:  string(keyPEM),
		credentialKeyFingerprint: "00:11:22:33:44:55:66:77:88:99:aa:bb:cc:dd:ee:ff",
		credentialKeyRegion:      "us-ashburn-1",
	}
	for _, customBundle := range []bool{false, true} {
		name := "default TLS config"
		if customBundle {
			name = "custom CA bundle"
		}
		t.Run(name, func(t *testing.T) {
			for _, tc := range []struct {
				name       string
				minVersion uint16
				maxVersion uint16
				groups     []tls.CurveID
				wantGroup  tls.CurveID
			}{
				{"prefers hybrid TLS 1.3", tls.VersionTLS12, tls.VersionTLS13, []tls.CurveID{tls.X25519MLKEM768, tls.CurveP256}, tls.X25519MLKEM768},
				{"P-256 hybrid TLS 1.3", tls.VersionTLS13, tls.VersionTLS13, []tls.CurveID{tls.SecP256r1MLKEM768}, tls.SecP256r1MLKEM768},
				{"classical TLS 1.3 peer", tls.VersionTLS13, tls.VersionTLS13, []tls.CurveID{tls.CurveP256}, tls.CurveP256},
				{"TLS 1.2 peer", tls.VersionTLS12, tls.VersionTLS12, []tls.CurveID{tls.CurveP256}, tls.CurveP256},
			} {
				t.Run(tc.name, func(t *testing.T) {
					server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						if r.Header.Get("Authorization") == "" {
							http.Error(w, "expected a signed OCI request", http.StatusUnauthorized)
							return
						}
						w.Header().Set("Content-Type", "application/json")
						_, _ = w.Write([]byte(`"pqc-test"`))
					}))
					server.TLS = &tls.Config{MinVersion: tc.minVersion, MaxVersion: tc.maxVersion, CurvePreferences: tc.groups}
					server.StartTLS()
					defer server.Close()
					t.Setenv("custom_cert_location", "")
					if customBundle {
						bundlePath := filepath.Join(t.TempDir(), "ca-bundle.pem")
						bundle := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: server.Certificate().Raw})
						if err := os.WriteFile(bundlePath, bundle, 0o600); err != nil {
							t.Fatal(err)
						}
						t.Setenv("custom_cert_location", bundlePath)
					}
					meta, err := newProviderMetaCache(1).getOrConfigureProviderMeta(context.Background(), "pqc-test", cfg)
					if err != nil {
						t.Fatalf("configure embedded provider: %v", err)
					}
					clients, ok := meta.(interface {
						ObjectStorageClient() *objectstorage.ObjectStorageClient
					})
					if !ok {
						t.Fatalf("unexpected embedded provider meta %T", meta)
					}
					client := clients.ObjectStorageClient()
					httpClient, ok := client.HTTPClient.(*http.Client)
					if !ok {
						t.Fatalf("unexpected HTTP client %T", client.HTTPClient)
					}
					transport, ok := httpClient.Transport.(*http.Transport)
					if !ok || transport.TLSClientConfig == nil {
						t.Fatalf("unexpected TLS transport %T", httpClient.Transport)
					}
					defer httpClient.CloseIdleConnections()
					if !customBundle {
						// Trust only the loopback peer in place of system roots.
						transport.TLSClientConfig.RootCAs = x509.NewCertPool()
						transport.TLSClientConfig.RootCAs.AddCert(server.Certificate())
					}
					httpClient.Timeout = 5 * time.Second
					client.Host = server.URL
					noRetry := common.NoRetryPolicy()
					response, err := client.GetNamespace(context.Background(), objectstorage.GetNamespaceRequest{
						RequestMetadata: common.RequestMetadata{RetryPolicy: &noRetry},
					})
					if err != nil {
						t.Fatalf("embedded provider OCI TLS request: %v", err)
					}
					if response.Value == nil || *response.Value != "pqc-test" || response.RawResponse == nil || response.RawResponse.TLS == nil {
						t.Fatal("expected a parsed OCI response over TLS")
					}
					state := response.RawResponse.TLS
					if state.Version != tc.maxVersion || state.CurveID != tc.wantGroup {
						t.Fatalf("negotiated TLS %#x, group %s; want TLS %#x, group %s", state.Version, state.CurveID, tc.maxVersion, tc.wantGroup)
					}
					if len(state.VerifiedChains) == 0 || state.DidResume {
						t.Fatal("expected a fresh handshake with certificate verification enabled")
					}
					t.Logf("TLS %#x, negotiated group %s", state.Version, state.CurveID)
				})
			}
		})
	}
}

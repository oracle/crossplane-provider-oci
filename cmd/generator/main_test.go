package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateServiceControllerSetupPackages(t *testing.T) {
	root := t.TempDir()
	content := []byte("package controller\n")
	for _, scope := range []string{"cluster", "namespaced"} {
		controllerRoot := filepath.Join(root, "internal", "controller", scope)
		if err := os.MkdirAll(filepath.Join(controllerRoot, "identity"), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(controllerRoot, "zz_identity_setup.go"), content, 0o644); err != nil {
			t.Fatal(err)
		}
		// Config intentionally has no service-specific target. Monolith's
		// target is created by the generator because it aggregates all services.
		for _, service := range []string{"config", "monolith"} {
			if err := os.WriteFile(filepath.Join(controllerRoot, "zz_"+service+"_setup.go"), content, 0o644); err != nil {
				t.Fatal(err)
			}
		}
	}

	if err := generateServiceControllerSetupPackages(root); err != nil {
		t.Fatalf("generateServiceControllerSetupPackages() error = %v", err)
	}
	for _, scope := range []string{"cluster", "namespaced"} {
		for _, service := range []string{"identity", "monolith"} {
			got, err := os.ReadFile(filepath.Join(root, "internal", "controller", scope, service, "zz_setup.go"))
			if err != nil {
				t.Fatal(err)
			}
			if string(got) != string(content) {
				t.Fatalf("generated %s/%s setup = %q, want %q", scope, service, got, content)
			}
		}
	}
}

func TestGenerateServiceControllerSetupPackagesRejectsMissingServiceDirectory(t *testing.T) {
	root := t.TempDir()
	for _, scope := range []string{"cluster", "namespaced"} {
		controllerRoot := filepath.Join(root, "internal", "controller", scope)
		if err := os.MkdirAll(controllerRoot, 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(controllerRoot, "zz_identity_setup.go"), []byte("package controller\n"), 0o644); err != nil {
			t.Fatal(err)
		}
	}

	if err := generateServiceControllerSetupPackages(root); err == nil {
		t.Fatal("generateServiceControllerSetupPackages() error = nil, want missing service directory error")
	}
}

func TestAPIPackageForGroup(t *testing.T) {
	tests := map[string]struct {
		group string
		want  apiPackage
	}{
		"cluster": {
			group: "identity.oci.upbound.io",
			want:  apiPackage{scope: "cluster", service: "identity", version: "v1alpha1"},
		},
		"namespaced": {
			group: "identity.oci.m.upbound.io",
			want:  apiPackage{scope: "namespaced", service: "identity", version: "v1alpha1"},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := apiPackageForGroup(tt.group, "v1alpha1")
			if err != nil {
				t.Fatalf("apiPackageForGroup() error = %v", err)
			}
			if got != tt.want {
				t.Fatalf("apiPackageForGroup() = %#v, want %#v", got, tt.want)
			}
		})
	}
	if _, err := apiPackageForGroup("example.com", "v1alpha1"); err == nil {
		t.Fatal("apiPackageForGroup() error = nil, want unsupported group error")
	}
}

func TestWriteServiceAPIScheme(t *testing.T) {
	root := t.TempDir()
	packages := map[apiPackage]struct{}{
		{scope: "cluster", service: "identity", version: "v1alpha1"}: {},
		{scope: "namespaced", version: "v1beta1"}:                    {},
	}
	for pkg := range packages {
		dir := filepath.Join(root, "apis", pkg.scope)
		if pkg.service != "" {
			dir = filepath.Join(dir, pkg.service)
		}
		if err := os.MkdirAll(filepath.Join(dir, pkg.version), 0o755); err != nil {
			t.Fatal(err)
		}
	}

	if err := writeServiceAPIScheme(root, "identity", packages); err != nil {
		t.Fatalf("writeServiceAPIScheme() error = %v", err)
	}
	content, err := os.ReadFile(filepath.Join(root, "internal", "apis", "runtime", "identity", "zz_scheme.go"))
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"apis/cluster/identity/v1alpha1",
		"apis/namespaced/v1beta1",
		"SchemeBuilder.AddToScheme",
	} {
		if !strings.Contains(string(content), want) {
			t.Errorf("generated scheme does not contain %q", want)
		}
	}
}

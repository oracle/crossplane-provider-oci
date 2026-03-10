package workflowgenerator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunGeneratesWorkflow(t *testing.T) {
	root := t.TempDir()
	cfg := NewConfig(root, "v1alpha1", nil)

	writeTestFile(t, filepath.Join(cfg.InputDir, "network-v1alpha1.yaml"), `
apiVersion: argoproj.io/v1alpha1
kind: WorkflowTemplate
metadata:
  name: network-template
spec:
  entrypoint: network-tests
  arguments:
    parameters:
    - name: compartment_ocid
      value: foo
`)
	writeTestFile(t, filepath.Join(cfg.InputDir, "identity-v1alpha1.yaml"), `
apiVersion: argoproj.io/v1alpha1
kind: WorkflowTemplate
metadata:
  name: identity-template
spec:
  entrypoint: identity-tests
  arguments:
    parameters:
    - name: tenancy_ocid
`)

	if err := Run(cfg); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	outputPath := filepath.Join(cfg.OutputDir, "crossplane-provider-oci-v1alpha1.yaml")
	data, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("expected output file: %v", err)
	}

	output := string(data)
	if !strings.Contains(output, "name: run_identity_tests") {
		t.Fatalf("generated workflow missing identity run param: %s", output)
	}
	if !strings.Contains(output, "name: run_network_tests") {
		t.Fatalf("generated workflow missing network run param: %s", output)
	}
	if !strings.Contains(output, "name: compartment_ocid") || !strings.Contains(output, "name: tenancy_ocid") {
		t.Fatalf("generated workflow missing merged parameters: %s", output)
	}
}

func TestRunFiltersRequestedServices(t *testing.T) {
	root := t.TempDir()
	cfg := NewConfig(root, "v1alpha1", []string{"network"})

	writeTestFile(t, filepath.Join(cfg.InputDir, "network-v1alpha1.yaml"), `
apiVersion: argoproj.io/v1alpha1
kind: WorkflowTemplate
metadata:
  name: network-template
spec:
  entrypoint: network-tests
`)
	writeTestFile(t, filepath.Join(cfg.InputDir, "identity-v1alpha1.yaml"), `
apiVersion: argoproj.io/v1alpha1
kind: WorkflowTemplate
metadata:
  name: identity-template
spec:
  entrypoint: identity-tests
`)

	if err := Run(cfg); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	outputPath := filepath.Join(cfg.OutputDir, "crossplane-provider-oci-v1alpha1.yaml")
	data, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("expected output file: %v", err)
	}
	output := string(data)
	if strings.Contains(output, "identity-template") {
		t.Fatalf("generated workflow unexpectedly included filtered service: %s", output)
	}
	if !strings.Contains(output, "network-template") {
		t.Fatalf("generated workflow missing requested service: %s", output)
	}
}

func writeTestFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(strings.TrimSpace(content)+"\n"), 0o644); err != nil {
		t.Fatal(err)
	}
}

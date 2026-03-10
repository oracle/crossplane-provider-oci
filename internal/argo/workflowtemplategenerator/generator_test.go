package workflowtemplategenerator

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestResolvePrerequisite(t *testing.T) {
	pr, ok := resolvePrerequisite("sourceIdSelector", "instance", "ex")
	if !ok {
		t.Fatal("resolvePrerequisite() expected override match")
	}
	if pr.Kind != "image" || pr.SelectorId != "ex" {
		t.Fatalf("resolvePrerequisite() override = %+v", pr)
	}

	pr, ok = resolvePrerequisite("vaultIdSelector", "anything", "ex2")
	if !ok {
		t.Fatal("resolvePrerequisite() expected inferred match")
	}
	if pr.Kind != "vault" || pr.SelectorId != "ex2" {
		t.Fatalf("resolvePrerequisite() inferred = %+v", pr)
	}
}

func TestExtractMetadataFromForProvider(t *testing.T) {
	forProvider := map[string]interface{}{
		"sourceIdSelector": map[string]interface{}{
			"matchLabels": map[string]interface{}{exampleNameLabel: "img-a"},
		},
		"displayName": "${oci.instance.name}",
		"rules": []interface{}{
			map[string]interface{}{
				"subnetIdSelector": map[string]interface{}{
					"matchLabels": map[string]interface{}{exampleNameLabel: "sub-a"},
				},
			},
		},
	}

	prereqs, envVars := extractMetadataFromForProvider(forProvider, "instance")
	if len(prereqs) != 2 {
		t.Fatalf("extractMetadataFromForProvider() prereqs len = %d, want 2", len(prereqs))
	}
	if envVars["oci.instance.name"] != "oci_instance_name" {
		t.Fatalf("extractMetadataFromForProvider() env var mapping missing, got: %#v", envVars)
	}

	for i := 1; i < len(prereqs); i++ {
		prev := prereqs[i-1]
		cur := prereqs[i]
		if prev.Kind > cur.Kind || (prev.Kind == cur.Kind && prev.SelectorId > cur.SelectorId) {
			t.Fatalf("prerequisites are not deterministically sorted: %#v", prereqs)
		}
	}
}

func TestResolveEnvVarsDeterministicOrdering(t *testing.T) {
	envVars := map[string]string{
		"C": "third",
		"A": "first",
		"B": "second",
	}

	result1 := resolveEnvVars(envVars)
	result2 := resolveEnvVars(envVars)
	if result1 != result2 {
		t.Fatalf("resolveEnvVars() non-deterministic output: %q vs %q", result1, result2)
	}

	parts := strings.Split(result1, ",")
	if len(parts) != 3 {
		t.Fatalf("resolveEnvVars() expected 3 parts, got %d (%q)", len(parts), result1)
	}
	if !strings.HasPrefix(parts[0], "A=") || !strings.HasPrefix(parts[1], "B=") || !strings.HasPrefix(parts[2], "C=") {
		t.Fatalf("resolveEnvVars() not sorted by key: %q", result1)
	}
}

func TestIsResourceFilePresent(t *testing.T) {
	resourceFiles := []ResourceFile{{Kind: "subnet", SelectorId: "a"}}
	if !isResourceFilePresent(resourceFiles, Prerequisite{Kind: "subnet", SelectorId: "a"}) {
		t.Fatal("isResourceFilePresent() = false, want true")
	}
	if isResourceFilePresent(resourceFiles, Prerequisite{Kind: "subnet", SelectorId: "b"}) {
		t.Fatal("isResourceFilePresent() = true, want false")
	}
}

func TestSearchForResourceFileAndGetResourceFileName(t *testing.T) {
	root := t.TempDir()
	cfg := Config{
		RootDir:     root,
		Version:     "v1alpha1",
		ExamplesDir: filepath.Join(root, "examples"),
		OutputDir:   filepath.Join(root, "argo", "workflowtemplates", "generated-workflowtemplates"),
	}
	if err := os.MkdirAll(filepath.Join(cfg.ExamplesDir, "network", cfg.Version), 0o755); err != nil {
		t.Fatal(err)
	}

	resourcePath := filepath.Join(cfg.ExamplesDir, "network", cfg.Version, "subnet.yaml")
	writeTestFile(t, resourcePath, "kind: Subnet\nmetadata:\n  labels:\n    testing.upbound.io/example-name: ex-sub\n")

	g, err := NewGenerator(cfg)
	if err != nil {
		t.Fatal(err)
	}

	pr := Prerequisite{Kind: "subnet", SelectorId: "ex-sub"}
	relPath, err := g.searchForResourceFile(pr)
	if err != nil {
		t.Fatalf("searchForResourceFile() unexpected error: %v", err)
	}
	if relPath != filepath.Join("network", cfg.Version, "subnet.yaml") {
		t.Fatalf("searchForResourceFile() path = %q", relPath)
	}

	name, err := g.getResourceFileName(pr)
	if err != nil {
		t.Fatalf("getResourceFileName() unexpected error: %v", err)
	}
	if name != "subnet" {
		t.Fatalf("getResourceFileName() = %q, want subnet", name)
	}
}

func TestResolveDeleteParametersNormalizesTaskName(t *testing.T) {
	got := resolveDeleteParameters("my_resource", "Name")
	want := "{{tasks.create-my-resource.outputs.parameters.resourceName}}"
	if got != want {
		t.Fatalf("resolveDeleteParameters() = %q, want %q", got, want)
	}
}

func TestSearchForResourceFileCacheIsolation(t *testing.T) {
	root := t.TempDir()
	cfg := Config{
		RootDir:     root,
		Version:     "v1alpha1",
		ExamplesDir: filepath.Join(root, "examples"),
		OutputDir:   filepath.Join(root, "argo", "workflowtemplates", "generated-workflowtemplates"),
	}

	writeTestFile(t, filepath.Join(cfg.ExamplesDir, "svc-a", cfg.Version, "subnet.yaml"), "kind: Subnet\nmetadata:\n  labels:\n    testing.upbound.io/example-name: ex-sub\n")
	writeTestFile(t, filepath.Join(cfg.ExamplesDir, "svc-b", cfg.Version, "subnet.yaml"), "kind: Subnet\nmetadata:\n  labels:\n    testing.upbound.io/example-name: ex-sub\n")

	pr := Prerequisite{Kind: "subnet", SelectorId: "ex-sub"}

	gA, err := NewGenerator(cfg)
	if err != nil {
		t.Fatal(err)
	}
	gA.resourceKindToFileMapping[pr] = filepath.Join("svc-a", cfg.Version, "subnet.yaml")

	gB, err := NewGenerator(cfg)
	if err != nil {
		t.Fatal(err)
	}
	gB.resourceKindToFileMapping[pr] = filepath.Join("svc-b", cfg.Version, "subnet.yaml")

	pathA, err := gA.searchForResourceFile(pr)
	if err != nil {
		t.Fatalf("searchForResourceFile(gA) unexpected error: %v", err)
	}
	if pathA != filepath.Join("svc-a", cfg.Version, "subnet.yaml") {
		t.Fatalf("searchForResourceFile(gA) path = %q", pathA)
	}

	pathB, err := gB.searchForResourceFile(pr)
	if err != nil {
		t.Fatalf("searchForResourceFile(gB) unexpected error: %v", err)
	}
	if pathB != filepath.Join("svc-b", cfg.Version, "subnet.yaml") {
		t.Fatalf("searchForResourceFile(gB) path = %q", pathB)
	}
}

func TestProcessExamplesGeneratesTemplate(t *testing.T) {
	root := t.TempDir()
	cfg := Config{
		RootDir:     root,
		Version:     "v1alpha1",
		ExamplesDir: filepath.Join(root, "examples"),
		OutputDir:   filepath.Join(root, "argo", "workflowtemplates", "generated-workflowtemplates"),
	}

	writeTestFile(t, filepath.Join(cfg.ExamplesDir, "core", cfg.Version, "vcn.yaml"), "kind: VCN\nmetadata:\n  labels:\n    testing.upbound.io/example-name: ex-vcn\n")

	g, err := NewGenerator(cfg)
	if err != nil {
		t.Fatal(err)
	}
	if err := g.processExamples(); err != nil {
		t.Fatalf("processExamples() unexpected error: %v", err)
	}

	output := filepath.Join(cfg.OutputDir, "core-v1alpha1.yaml")
	if _, err := os.Stat(output); err != nil {
		t.Fatalf("expected generated template %s: %v", output, err)
	}
}

func TestProcessExamplesReturnsJoinedErrorOnServiceFailures(t *testing.T) {
	root := t.TempDir()
	cfg := Config{
		RootDir:     root,
		Version:     "v1alpha1",
		ExamplesDir: filepath.Join(root, "examples"),
		OutputDir:   filepath.Join(root, "argo", "workflowtemplates", "generated-workflowtemplates"),
	}

	writeTestFile(t, filepath.Join(cfg.ExamplesDir, "ok", cfg.Version, "vcn.yaml"), "kind: VCN\nmetadata:\n  labels:\n    testing.upbound.io/example-name: ex-vcn\n")
	writeTestFile(t, filepath.Join(cfg.ExamplesDir, "broken", cfg.Version, "bad.yaml"), "kind: [invalid\n")

	g, err := NewGenerator(cfg)
	if err != nil {
		t.Fatal(err)
	}

	err = g.processExamples()
	if err == nil {
		t.Fatal("processExamples() expected error, got nil")
	}
	if !strings.Contains(err.Error(), "broken") {
		t.Fatalf("processExamples() error = %q, want service name", err)
	}

	var joined interface{ Unwrap() []error }
	if !errors.As(err, &joined) {
		t.Fatalf("processExamples() expected joined error, got %T", err)
	}

	output := filepath.Join(cfg.OutputDir, "ok-v1alpha1.yaml")
	if _, statErr := os.Stat(output); statErr != nil {
		t.Fatalf("expected generated template for healthy service %s: %v", output, statErr)
	}
}

func TestRun(t *testing.T) {
	root := t.TempDir()
	cfg := NewConfig(root, "v1alpha1")

	writeTestFile(t, filepath.Join(cfg.ExamplesDir, "dns", "v1alpha1", "zone.yaml"), "kind: Zone\nmetadata:\n  labels:\n    testing.upbound.io/example-name: ex-zone\n")

	if err := Run(cfg); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	if _, err := os.Stat(filepath.Join(root, "argo", "workflowtemplates", "generated-workflowtemplates", "dns-v1alpha1.yaml")); err != nil {
		t.Fatalf("Run() expected output file: %v", err)
	}
}

func writeTestFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

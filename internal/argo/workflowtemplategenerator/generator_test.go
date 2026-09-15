package workflowtemplategenerator

import (
	"errors"
	"fmt"
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

func TestExtractMetadataFromForProviderDeduplicatesPrerequisites(t *testing.T) {
	selector := func() map[string]interface{} {
		return map[string]interface{}{
			"matchLabels": map[string]interface{}{exampleNameLabel: "sub-a"},
		}
	}
	forProvider := map[string]interface{}{
		"vnics": []interface{}{
			map[string]interface{}{"subnetIdSelector": selector()},
			map[string]interface{}{"subnetIdSelector": selector()},
		},
	}

	prereqs, _ := extractMetadataFromForProvider(forProvider, "nodepool")
	if len(prereqs) != 1 {
		t.Fatalf("extractMetadataFromForProvider() prereqs = %#v, want one unique prerequisite", prereqs)
	}
	want := Prerequisite{Kind: "subnet", SelectorId: "sub-a"}
	if prereqs[0] != want {
		t.Fatalf("extractMetadataFromForProvider() prerequisite = %+v, want %+v", prereqs[0], want)
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
	if err := os.MkdirAll(filepath.Join(cfg.ExamplesDir, scopeCluster, "network", cfg.Version), 0o755); err != nil {
		t.Fatal(err)
	}

	resourcePath := filepath.Join(cfg.ExamplesDir, scopeCluster, "network", cfg.Version, "subnet.yaml")
	writeTestFile(t, resourcePath, "kind: Subnet\nmetadata:\n  labels:\n    testing.upbound.io/example-name: ex-sub\n")

	g, err := NewGenerator(cfg)
	if err != nil {
		t.Fatal(err)
	}

	pr := Prerequisite{Kind: "subnet", SelectorId: "ex-sub"}
	relPath, err := g.searchForResourceFile(pr, "")
	if err != nil {
		t.Fatalf("searchForResourceFile() unexpected error: %v", err)
	}
	if relPath != filepath.Join(scopeCluster, "network", cfg.Version, "subnet.yaml") {
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

	writeTestFile(t, filepath.Join(cfg.ExamplesDir, scopeCluster, "svc-a", cfg.Version, "subnet.yaml"), "kind: Subnet\nmetadata:\n  labels:\n    testing.upbound.io/example-name: ex-sub\n")
	writeTestFile(t, filepath.Join(cfg.ExamplesDir, scopeNamespaced, "svc-b", cfg.Version, "subnet.yaml"), "kind: Subnet\nmetadata:\n  labels:\n    testing.upbound.io/example-name: ex-sub\n")

	pr := Prerequisite{Kind: "subnet", SelectorId: "ex-sub"}

	gA, err := NewGenerator(cfg)
	if err != nil {
		t.Fatal(err)
	}
	gA.resourceKindToFileMapping[pr] = map[string]string{
		scopeCluster: filepath.Join(scopeCluster, "svc-a", cfg.Version, "subnet.yaml"),
	}

	gB, err := NewGenerator(cfg)
	if err != nil {
		t.Fatal(err)
	}
	gB.resourceKindToFileMapping[pr] = map[string]string{
		scopeNamespaced: filepath.Join(scopeNamespaced, "svc-b", cfg.Version, "subnet.yaml"),
	}

	pathA, err := gA.searchForResourceFile(pr, "")
	if err != nil {
		t.Fatalf("searchForResourceFile(gA) unexpected error: %v", err)
	}
	if pathA != filepath.Join(scopeCluster, "svc-a", cfg.Version, "subnet.yaml") {
		t.Fatalf("searchForResourceFile(gA) path = %q", pathA)
	}

	pathB, err := gB.searchForResourceFile(pr, "")
	if err != nil {
		t.Fatalf("searchForResourceFile(gB) unexpected error: %v", err)
	}
	if pathB != filepath.Join(scopeNamespaced, "svc-b", cfg.Version, "subnet.yaml") {
		t.Fatalf("searchForResourceFile(gB) path = %q", pathB)
	}
}

func TestSearchForResourceFileCrossScope(t *testing.T) {
	root := t.TempDir()
	cfg := Config{
		RootDir:     root,
		Version:     "v1alpha1",
		ExamplesDir: filepath.Join(root, "examples"),
		OutputDir:   filepath.Join(root, "argo", "workflowtemplates", "generated-workflowtemplates"),
	}

	writeTestFile(t, filepath.Join(cfg.ExamplesDir, scopeNamespaced, "network", cfg.Version, "subnet.yaml"), "kind: Subnet\nmetadata:\n  labels:\n    testing.upbound.io/example-name: ex-sub\n")

	g, err := NewGenerator(cfg)
	if err != nil {
		t.Fatal(err)
	}

	pr := Prerequisite{Kind: "subnet", SelectorId: "ex-sub"}
	path, err := g.searchForResourceFile(pr, "")
	if err != nil {
		t.Fatalf("searchForResourceFile() unexpected error: %v", err)
	}
	if path != filepath.Join(scopeNamespaced, "network", cfg.Version, "subnet.yaml") {
		t.Fatalf("searchForResourceFile() cross-scope path = %q", path)
	}
}

func TestSearchForResourceFilePrefersRequestedScope(t *testing.T) {
	root := t.TempDir()
	cfg := Config{
		RootDir:     root,
		Version:     "v1alpha1",
		ExamplesDir: filepath.Join(root, "examples"),
		OutputDir:   filepath.Join(root, "argo", "workflowtemplates", "generated-workflowtemplates"),
	}

	pr := Prerequisite{Kind: "subnet", SelectorId: "ex-sub"}

	clusterRel := filepath.Join(scopeCluster, "network", cfg.Version, "subnet.yaml")
	namespacedRel := filepath.Join(scopeNamespaced, "network", cfg.Version, "subnet.yaml")

	writeTestFile(t, filepath.Join(cfg.ExamplesDir, clusterRel), "kind: Subnet\nmetadata:\n  labels:\n    testing.upbound.io/example-name: ex-sub\n")
	writeTestFile(t, filepath.Join(cfg.ExamplesDir, namespacedRel), "kind: Subnet\nmetadata:\n  labels:\n    testing.upbound.io/example-name: ex-sub\n")

	g, err := NewGenerator(cfg)
	if err != nil {
		t.Fatal(err)
	}

	g.resourceKindToFileMapping[pr] = map[string]string{
		scopeCluster: clusterRel,
	}

	path, err := g.searchForResourceFile(pr, scopeNamespaced)
	if err != nil {
		t.Fatalf("searchForResourceFile() unexpected error: %v", err)
	}
	if path != namespacedRel {
		t.Fatalf("searchForResourceFile() preferred scope path = %q, want %q", path, namespacedRel)
	}

	if stored := g.resourceKindToFileMapping[pr][scopeNamespaced]; stored != namespacedRel {
		t.Fatalf("expected namespaced mapping cached, got %q", stored)
	}
}

func TestCollectPrerequisitesPrefersCurrentScope(t *testing.T) {
	root := t.TempDir()
	cfg := Config{
		RootDir:     root,
		Version:     "v1alpha1",
		ExamplesDir: filepath.Join(root, "examples"),
		OutputDir:   filepath.Join(root, "argo", "workflowtemplates", "generated-workflowtemplates"),
	}

	clusterSubnet := filepath.Join(cfg.ExamplesDir, scopeCluster, "network", cfg.Version, "subnet.yaml")
	namespacedSubnet := filepath.Join(cfg.ExamplesDir, scopeNamespaced, "network", cfg.Version, "subnet.yaml")
	writeTestFile(t, clusterSubnet, "kind: Subnet\nmetadata:\n  labels:\n    testing.upbound.io/example-name: ex-subnet\n")
	writeTestFile(t, namespacedSubnet, "kind: Subnet\nmetadata:\n  labels:\n    testing.upbound.io/example-name: ex-subnet\n")

	namespacedSvcDir := filepath.Join(cfg.ExamplesDir, scopeNamespaced, "compute", cfg.Version)
	writeTestFile(t, filepath.Join(namespacedSvcDir, "instance.yaml"), ""+
		"kind: Instance\n"+
		"metadata:\n"+
		"  labels:\n"+
		"    testing.upbound.io/example-name: ex-instance\n"+
		"spec:\n"+
		"  forProvider:\n"+
		"    subnetIdSelector:\n"+
		"      matchLabels:\n"+
		"        testing.upbound.io/example-name: ex-subnet\n")

	g, err := NewGenerator(cfg)
	if err != nil {
		t.Fatal(err)
	}

	resourceFiles, err := g.getResourceFiles(namespacedSvcDir)
	if err != nil {
		t.Fatalf("getResourceFiles() error: %v", err)
	}
	if len(resourceFiles) != 1 {
		t.Fatalf("expected 1 resource file, got %d", len(resourceFiles))
	}

	data, err := g.buildTemplateData("compute", cfg.Version, scopeNamespaced, namespacedSvcDir, true)
	if err != nil {
		t.Fatalf("buildTemplateData() error: %v", err)
	}

	if len(data.PrerequisiteResourceFiles) != 1 {
		t.Fatalf("expected 1 prerequisite resource file, got %d", len(data.PrerequisiteResourceFiles))
	}

	if got, want := data.PrerequisiteResourceFiles[0].Path, filepath.Join(scopeNamespaced, "network", cfg.Version, "subnet.yaml"); got != want {
		t.Fatalf("prerequisite path = %q, want %q", got, want)
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

	writeTestFile(t, filepath.Join(cfg.ExamplesDir, scopeCluster, "core", cfg.Version, "vcn.yaml"), "kind: VCN\nmetadata:\n  labels:\n    testing.upbound.io/example-name: ex-vcn\n")

	g, err := NewGenerator(cfg)
	if err != nil {
		t.Fatal(err)
	}
	if err := g.processExamples(); err != nil {
		t.Fatalf("processExamples() unexpected error: %v", err)
	}

	output := filepath.Join(cfg.OutputDir, scopeCluster, "core-cluster-v1alpha1.yaml")
	if _, err := os.Stat(output); err != nil {
		t.Fatalf("expected generated template %s: %v", output, err)
	}
}

func TestProcessExamplesGeneratesNamespacedTemplateNamespaceInput(t *testing.T) {
	root := t.TempDir()
	cfg := Config{
		RootDir:     root,
		Version:     "v1alpha1",
		ExamplesDir: filepath.Join(root, "examples"),
		OutputDir:   filepath.Join(root, "argo", "workflowtemplates", "generated-workflowtemplates"),
	}

	writeTestFile(t, filepath.Join(cfg.ExamplesDir, scopeNamespaced, "core", cfg.Version, "vcn.yaml"), "kind: VCN\nmetadata:\n  labels:\n    testing.upbound.io/example-name: ex-vcn\n")

	g, err := NewGenerator(cfg)
	if err != nil {
		t.Fatal(err)
	}
	if err := g.processExamples(); err != nil {
		t.Fatalf("processExamples() unexpected error: %v", err)
	}

	output := filepath.Join(cfg.OutputDir, scopeNamespaced, "core-namespaced-v1alpha1.yaml")
	data, err := os.ReadFile(output)
	if err != nil {
		t.Fatalf("expected generated template %s: %v", output, err)
	}

	content := string(data)
	if !strings.Contains(content, "value: \"{{inputs.parameters.target_namespace}}\"") {
		t.Fatalf("namespaced template should pass namespace from entrypoint input, got:\n%s", content)
	}
	if !strings.Contains(content, "value: \"{{workflow.parameters.target_namespace}}\"") {
		t.Fatalf("namespaced template should default entrypoint input from workflow parameter for direct submits, got:\n%s", content)
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

	writeTestFile(t, filepath.Join(cfg.ExamplesDir, scopeCluster, "ok", cfg.Version, "vcn.yaml"), "kind: VCN\nmetadata:\n  labels:\n    testing.upbound.io/example-name: ex-vcn\n")
	writeTestFile(t, filepath.Join(cfg.ExamplesDir, scopeCluster, "broken", cfg.Version, "bad.yaml"), "kind: [invalid\n")

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

	output := filepath.Join(cfg.OutputDir, scopeCluster, "ok-cluster-v1alpha1.yaml")
	if _, statErr := os.Stat(output); statErr != nil {
		t.Fatalf("expected generated template for healthy service %s: %v", output, statErr)
	}
}

func TestRun(t *testing.T) {
	root := t.TempDir()
	cfg := NewConfig(root, "v1alpha1")

	writeTestFile(t, filepath.Join(cfg.ExamplesDir, scopeCluster, "dns", "v1alpha1", "zone.yaml"), "kind: Zone\nmetadata:\n  labels:\n    testing.upbound.io/example-name: ex-zone\n")

	if err := Run(cfg); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	if _, err := os.Stat(filepath.Join(root, "argo", "workflowtemplates", "generated-workflowtemplates", scopeCluster, "dns-cluster-v1alpha1.yaml")); err != nil {
		t.Fatalf("Run() expected output file: %v", err)
	}
}

func TestProcessExamplesReturnsErrorForUnknownRequestedService(t *testing.T) {
	root := t.TempDir()
	cfg := NewConfig(root, "v1alpha1")
	cfg.Services = []string{"doesnotexist"}

	writeTestFile(t, filepath.Join(cfg.ExamplesDir, scopeCluster, "dns", cfg.Version, "zone.yaml"), "kind: Zone\nmetadata:\n  labels:\n    testing.upbound.io/example-name: ex-zone\n")

	g, err := NewGenerator(cfg)
	if err != nil {
		t.Fatal(err)
	}

	err = g.processExamples()
	if err == nil {
		t.Fatal("processExamples() expected unknown service error, got nil")
	}
	if !strings.Contains(err.Error(), "unknown service(s): doesnotexist") {
		t.Fatalf("processExamples() error = %q, want unknown service error", err)
	}
}

func TestIntegrationApigatewayNamespacedReferences(t *testing.T) {
	if os.Getenv("WORKFLOW_GENERATOR_INTEGRATION") == "" {
		t.Skip("set WORKFLOW_GENERATOR_INTEGRATION=1 to run")
	}

	root, err := findRepoRootFromGeneratorTest()
	if err != nil {
		t.Skipf("repo root not available: %v", err)
	}

	if _, err := os.Stat(filepath.Join(root, "examples", scopeCluster)); err != nil {
		t.Skipf("cluster examples tree not available under %s: %v", root, err)
	}
	if _, err := os.Stat(filepath.Join(root, "examples", scopeNamespaced)); err != nil {
		t.Skipf("namespaced examples tree not available under %s: %v", root, err)
	}

	cfg := NewConfig(root, "v1alpha1")
	cfg.Services = []string{"apigateway"}
	outputDir := filepath.Join(os.TempDir(), "workflowtemplates-integration")
	if err := os.RemoveAll(outputDir); err != nil {
		t.Fatalf("RemoveAll() error: %v", err)
	}
	cfg.OutputDir = outputDir

	if err := Run(cfg); err != nil {
		t.Fatalf("Run() error: %v", err)
	}

	namespacedPath := filepath.Join(outputDir, scopeNamespaced, "apigateway-namespaced-v1alpha1.yaml")
	data, err := os.ReadFile(namespacedPath)
	if err != nil {
		t.Fatalf("ReadFile(%s) error: %v", namespacedPath, err)
	}

	content := string(data)
	if strings.Contains(content, "examples/cluster/identity") {
		t.Fatalf("namespaced template references cluster identity example:\n%s", namespacedPath)
	}
	if strings.Contains(content, "examples/cluster/networking") {
		t.Fatalf("namespaced template references cluster networking example:\n%s", namespacedPath)
	}

	t.Logf("apigateway templates generated under %s", outputDir)
}

func findRepoRootFromGeneratorTest() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for dir := wd; ; dir = filepath.Dir(dir) {
		if fileExistsFromGeneratorTest(filepath.Join(dir, "go.mod")) &&
			dirExistsFromGeneratorTest(filepath.Join(dir, "examples")) &&
			dirExistsFromGeneratorTest(filepath.Join(dir, "argo")) {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
	}

	return "", fmt.Errorf("unable to locate repository root from %s", wd)
}

func dirExistsFromGeneratorTest(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func fileExistsFromGeneratorTest(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
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

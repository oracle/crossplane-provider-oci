package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseVersionArg(t *testing.T) {
	version, err := parseVersionArg([]string{"cmd", "v1alpha1"})
	if err != nil {
		t.Fatalf("parseVersionArg() unexpected error: %v", err)
	}
	if version != "v1alpha1" {
		t.Fatalf("parseVersionArg() version = %q, want %q", version, "v1alpha1")
	}

	if _, err := parseVersionArg([]string{"cmd"}); err == nil {
		t.Fatal("parseVersionArg() expected error for missing version")
	}
}

func TestRun(t *testing.T) {
	origArgs := os.Args
	defer func() { os.Args = origArgs }()

	origWD, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = os.Chdir(origWD) }()

	root := t.TempDir()
	if err := os.Chdir(root); err != nil {
		t.Fatal(err)
	}

	os.Args = []string{"cmd", "v1alpha1"}
	if err := os.MkdirAll(filepath.Join(root, "examples", "dns", "v1alpha1"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "examples", "dns", "v1alpha1", "zone.yaml"), []byte("kind: Zone\nmetadata:\n  labels:\n    testing.upbound.io/example-name: ex-zone\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	if err := run(); err != nil {
		t.Fatalf("run() unexpected error: %v", err)
	}

	if _, err := os.Stat(filepath.Join(root, "argo", "workflowtemplates", "generated-workflowtemplates", "dns-v1alpha1.yaml")); err != nil {
		t.Fatalf("run() expected output file: %v", err)
	}
}

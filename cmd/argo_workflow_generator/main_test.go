package main

import "testing"

func TestParseArgs(t *testing.T) {
	version, services, err := parseArgs([]string{"cmd", "v1alpha1", "identity", "network", "identity"})
	if err != nil {
		t.Fatalf("parseArgs() unexpected error: %v", err)
	}
	if version != "v1alpha1" {
		t.Fatalf("parseArgs() version = %q, want %q", version, "v1alpha1")
	}
	if len(services) != 2 || services[0] != "identity" || services[1] != "network" {
		t.Fatalf("parseArgs() services = %#v", services)
	}
}

func TestParseArgsMissingVersion(t *testing.T) {
	if _, _, err := parseArgs([]string{"cmd"}); err == nil {
		t.Fatal("parseArgs() expected error for missing version")
	}
}

package noforkpatch

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCleanRestoresModuleFiles(t *testing.T) {
	root := t.TempDir()
	state := filepath.Join(root, ".work", "nofork", "state")
	providerDir := filepath.Join(root, ".work", "nofork", "terraform-provider-oci")
	gopath := filepath.Join(root, ".work", "nofork-gopath")

	mustWrite(t, filepath.Join(root, "go.mod"), "module broken\n")
	mustWrite(t, filepath.Join(root, "go.sum"), "broken sum\n")
	for _, source := range noForkSourceFiles {
		mustWrite(t, filepath.Join(root, source[1]), "package ignored\n")
	}
	mustWrite(t, filepath.Join(state, "go.mod"), "module restored\n")
	mustWrite(t, filepath.Join(state, "go.sum"), "restored sum\n")
	mustWrite(t, filepath.Join(providerDir, "README.md"), "patched provider\n")
	mustWrite(t, filepath.Join(gopath, "README.md"), "gopath\n")

	err := Clean(Options{
		RootDir:     root,
		StateDir:    state,
		ProviderDir: providerDir,
		GoPath:      gopath,
	})
	if err != nil {
		t.Fatalf("Clean returned error: %v", err)
	}

	gotMod, err := os.ReadFile(filepath.Join(root, "go.mod"))
	if err != nil {
		t.Fatal(err)
	}
	if string(gotMod) != "module restored\n" {
		t.Fatalf("go.mod was not restored, got %q", string(gotMod))
	}
	gotSum, err := os.ReadFile(filepath.Join(root, "go.sum"))
	if err != nil {
		t.Fatal(err)
	}
	if string(gotSum) != "restored sum\n" {
		t.Fatalf("go.sum was not restored, got %q", string(gotSum))
	}
	for _, path := range []string{state, providerDir, gopath} {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			t.Fatalf("expected %s to be removed, err=%v", path, err)
		}
	}
	for _, source := range noForkSourceFiles {
		if _, err := os.Stat(filepath.Join(root, source[1])); !os.IsNotExist(err) {
			t.Fatalf("expected no-fork source %s to be removed, err=%v", source[1], err)
		}
	}
}

func TestMaterializeNoForkSources(t *testing.T) {
	root := t.TempDir()
	for _, source := range noForkSourceFiles {
		mustWrite(t, filepath.Join(root, source[0]), "package ignored\n")
	}

	if err := materializeNoForkSources(Options{RootDir: root}); err != nil {
		t.Fatalf("materializeNoForkSources returned error: %v", err)
	}

	for _, source := range noForkSourceFiles {
		got, err := os.ReadFile(filepath.Join(root, source[1]))
		if err != nil {
			t.Fatal(err)
		}
		if string(got) != "package ignored\n" {
			t.Fatalf("materialized source %s = %q", source[1], string(got))
		}
	}
}

func TestApplyRestoresWorkspaceAfterTidyFailure(t *testing.T) {
	root := t.TempDir()
	providerDir := filepath.Join(root, ".work", "nofork", "terraform-provider-oci")
	stateDir := filepath.Join(root, ".work", "nofork", "state")
	mustWrite(t, filepath.Join(root, "go.mod"), "module original\n")
	mustWrite(t, filepath.Join(root, "go.sum"), "original sum\n")
	mustWrite(t, filepath.Join(root, "patch"), "patch")
	for _, source := range noForkSourceFiles {
		mustWrite(t, filepath.Join(root, source[0]), "package nofork\n")
	}

	err := apply(context.Background(), Options{
		RootDir:         root,
		ProviderVersion: "8.22.0",
		ProviderDir:     providerDir,
		StateDir:        stateDir,
		PatchFile:       filepath.Join(root, "patch"),
		RepoURL:         "https://example.com/terraform-provider-oci",
		ModulePath:      defaultModulePath,
		GoCache:         filepath.Join(root, ".cache", "go-build"),
		GoModCache:      filepath.Join(root, ".cache", "go-mod"),
		GoPath:          filepath.Join(root, ".work", "gopath"),
		Stdout:          io.Discard,
		Stderr:          io.Discard,
	}, failingTidyRunner{root: root})
	if !errors.Is(err, errTidy) {
		t.Fatalf("apply error = %v, want tidy failure", err)
	}

	for path, want := range map[string]string{
		"go.mod": "module original\n",
		"go.sum": "original sum\n",
	} {
		got, readErr := os.ReadFile(filepath.Join(root, path))
		if readErr != nil {
			t.Fatal(readErr)
		}
		if string(got) != want {
			t.Errorf("%s = %q, want %q", path, got, want)
		}
	}
	for _, source := range noForkSourceFiles {
		if _, statErr := os.Stat(filepath.Join(root, source[1])); !errors.Is(statErr, os.ErrNotExist) {
			t.Errorf("materialized source %s still exists, err=%v", source[1], statErr)
		}
	}
}

func TestValidateForPatchRequiresProviderVersion(t *testing.T) {
	opts := DefaultOptions(t.TempDir())
	opts.ProviderVersion = ""

	err := opts.validateForPatch()
	if err == nil {
		t.Fatal("expected missing provider-version error")
	}
}

func TestProviderTagAcceptsOptionalPrefix(t *testing.T) {
	for _, input := range []string{"8.12.0", "v8.12.0"} {
		if got := providerTag(input); got != "v8.12.0" {
			t.Fatalf("providerTag(%q) = %q", input, got)
		}
	}
}

func TestGoEnvAddsNoForkBuildTag(t *testing.T) {
	t.Setenv("GOFLAGS", "-mod=mod")

	env := goEnv(DefaultOptions(t.TempDir()))

	var got string
	for _, e := range env {
		if strings.HasPrefix(e, "GOFLAGS=") {
			got = e
			break
		}
	}
	if got != "GOFLAGS=-mod=mod -tags=nofork" {
		t.Fatalf("GOFLAGS entry = %q, want %q", got, "GOFLAGS=-mod=mod -tags=nofork")
	}
}

func mustWrite(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

var errTidy = errors.New("tidy failed")

type failingTidyRunner struct {
	root string
}

func (r failingTidyRunner) Run(_ context.Context, _ string, _ []string, name string, args ...string) error {
	if name == "git" {
		return os.MkdirAll(args[len(args)-1], 0o755)
	}
	if name != "go" {
		return nil
	}
	if len(args) >= 2 && args[0] == "mod" && args[1] == "edit" {
		if err := os.WriteFile(filepath.Join(r.root, "go.mod"), []byte("module rewritten\n"), 0o644); err != nil {
			return err
		}
		return os.WriteFile(filepath.Join(r.root, "go.sum"), []byte("rewritten sum\n"), 0o644)
	}
	if len(args) >= 2 && args[0] == "mod" && args[1] == "tidy" {
		return errTidy
	}
	return nil
}

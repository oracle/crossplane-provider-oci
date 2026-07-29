package noforkpatch

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	defaultRepoURL    = "https://github.com/oracle/terraform-provider-oci"
	defaultModulePath = "github.com/oracle/terraform-provider-oci"
	defaultPatchFile  = "build/nofork/patches/terraform-provider-oci.patch"
	noForkGoFlags     = "-tags=nofork"
)

var noForkSourceFiles = [][2]string{
	{"config/terraform_provider_nofork.go.tmpl", "config/terraform_provider_nofork.go"},
	{"internal/clients/oci_provider_nofork.go.tmpl", "internal/clients/oci_provider_nofork.go"},
}

// Options describes the workspace paths and upstream provider coordinates used
// by the no-fork patch flow.
type Options struct {
	RootDir         string
	ProviderVersion string
	ProviderDir     string
	StateDir        string
	PatchFile       string
	RepoURL         string
	ModulePath      string
	GoCache         string
	GoModCache      string
	GoPath          string
	Stdout          io.Writer
	Stderr          io.Writer
}

// DefaultOptions returns the production defaults used by the Makefile.
func DefaultOptions(rootDir string) Options {
	return Options{
		RootDir:         rootDir,
		ProviderVersion: os.Getenv("TERRAFORM_PROVIDER_VERSION"),
		ProviderDir:     filepath.Join(rootDir, ".work", "nofork", "terraform-provider-oci"),
		StateDir:        filepath.Join(rootDir, ".work", "nofork", "state"),
		PatchFile:       filepath.Join(rootDir, defaultPatchFile),
		RepoURL:         defaultRepoURL,
		ModulePath:      defaultModulePath,
		GoCache:         filepath.Join(rootDir, ".cache", "nofork-go-build"),
		GoModCache:      filepath.Join(rootDir, ".cache", "nofork-go-mod"),
		GoPath:          filepath.Join(rootDir, ".work", "nofork-gopath"),
		Stdout:          os.Stdout,
		Stderr:          os.Stderr,
	}
}

type runner interface {
	Run(ctx context.Context, dir string, env []string, name string, args ...string) error
}

type execRunner struct {
	stdout io.Writer
	stderr io.Writer
}

func (r execRunner) Run(ctx context.Context, dir string, env []string, name string, args ...string) error {
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), env...)
	cmd.Stdout = r.stdout
	cmd.Stderr = r.stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s %s: %w", name, strings.Join(args, " "), err)
	}
	return nil
}

// Apply clones the requested upstream provider version, applies the no-fork
// patch, and temporarily rewrites the workspace go.mod to use the patched clone.
func Apply(ctx context.Context, opts Options) error {
	return apply(ctx, normalizeOptions(opts), execRunner{stdout: opts.Stdout, stderr: opts.Stderr})
}

func apply(ctx context.Context, opts Options, r runner) error {
	if err := opts.validateForPatch(); err != nil {
		return err
	}
	if err := ensurePatchFile(opts.PatchFile); err != nil {
		return err
	}

	for _, dir := range []string{opts.GoCache, opts.GoModCache, opts.GoPath, filepath.Dir(opts.ProviderDir), opts.StateDir} {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("create directory %s: %w", dir, err)
		}
	}

	fmt.Fprintf(opts.Stdout, "==> Cloning terraform-provider-oci v%s\n", cleanVersion(opts.ProviderVersion))
	if err := os.RemoveAll(opts.ProviderDir); err != nil {
		return fmt.Errorf("remove provider directory %s: %w", opts.ProviderDir, err)
	}
	if err := r.Run(ctx, opts.RootDir, nil, "git", "clone", "--depth", "1", "--branch", providerTag(opts.ProviderVersion), opts.RepoURL, opts.ProviderDir); err != nil {
		return err
	}

	if err := dryRunPatch(ctx, opts, r); err != nil {
		return err
	}
	if err := applyPatch(ctx, opts, r); err != nil {
		return err
	}
	patched := false
	backedUp := false
	defer func() {
		if !patched {
			if backedUp {
				_ = restoreModuleFiles(opts)
			}
			_ = removeNoForkSources(opts)
		}
	}()
	if err := materializeNoForkSources(opts); err != nil {
		return err
	}

	if err := backupModuleFiles(opts); err != nil {
		return err
	}
	backedUp = true

	fmt.Fprintln(opts.Stdout, "==> Updating go.mod (require + replace)")
	env := goEnv(opts)
	if err := r.Run(ctx, opts.RootDir, env, "go", "mod", "edit",
		"-require", fmt.Sprintf("%s@%s+incompatible", opts.ModulePath, providerTag(opts.ProviderVersion)),
		"-replace", fmt.Sprintf("%s=%s", opts.ModulePath, opts.ProviderDir),
	); err != nil {
		return err
	}

	fmt.Fprintln(opts.Stdout, "==> Tidying go.mod/go.sum for patched provider dependency graph")
	if err := r.Run(ctx, opts.RootDir, env, "go", "mod", "tidy"); err != nil {
		return err
	}

	patched = true
	fmt.Fprintf(opts.Stdout, "==> Patch applied. go.mod/go.sum backups saved to %s.\n", opts.StateDir)
	return nil
}

// Validate checks that the no-fork patch applies to the requested upstream
// provider version and that the patched tree passes semantic safety checks.
func Validate(ctx context.Context, opts Options) error {
	return validate(ctx, normalizeOptions(opts), execRunner{stdout: opts.Stdout, stderr: opts.Stderr})
}

func validate(ctx context.Context, opts Options, r runner) error {
	if err := opts.validateForPatch(); err != nil {
		return err
	}
	if err := ensurePatchFile(opts.PatchFile); err != nil {
		return err
	}

	tmpDir, err := os.MkdirTemp("", "terraform-provider-oci-validate-*")
	if err != nil {
		return fmt.Errorf("create temporary validation directory: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	validationDir := filepath.Join(tmpDir, "terraform-provider-oci")
	fmt.Fprintf(opts.Stdout, "==> Validating patch against v%s\n", cleanVersion(opts.ProviderVersion))
	if err := r.Run(ctx, opts.RootDir, nil, "git", "clone", "--depth", "1", "--branch", providerTag(opts.ProviderVersion), opts.RepoURL, validationDir); err != nil {
		return err
	}

	validationOpts := opts
	validationOpts.ProviderDir = validationDir
	if err := dryRunPatch(ctx, validationOpts, r); err != nil {
		return err
	}
	if err := applyPatch(ctx, validationOpts, r); err != nil {
		return err
	}
	if err := ValidatePatchedTree(validationDir); err != nil {
		return err
	}

	fmt.Fprintf(opts.Stdout, "==> Patch validates cleanly against v%s\n", cleanVersion(opts.ProviderVersion))
	return nil
}

// Clean restores go.mod/go.sum from the patch state and removes transient
// no-fork build directories.
func Clean(opts Options) error {
	opts = normalizeOptions(opts)
	var errs []error
	if err := removeNoForkSources(opts); err != nil {
		errs = append(errs, fmt.Errorf("remove no-fork source: %w", err))
	}
	if err := restoreModuleFiles(opts); err != nil {
		errs = append(errs, err)
	}
	for _, dir := range []string{opts.StateDir, opts.ProviderDir, opts.GoPath} {
		if err := makeWritable(dir); err != nil {
			errs = append(errs, err)
		}
		if err := os.RemoveAll(dir); err != nil {
			errs = append(errs, fmt.Errorf("remove %s: %w", dir, err))
		}
	}
	return errors.Join(errs...)
}

func restoreModuleFiles(opts Options) error {
	var errs []error
	for _, name := range []string{"go.mod", "go.sum"} {
		src := filepath.Join(opts.StateDir, name)
		dst := filepath.Join(opts.RootDir, name)
		if _, err := os.Stat(src); err == nil {
			if err := copyFile(src, dst); err != nil {
				errs = append(errs, fmt.Errorf("restore %s: %w", name, err))
			}
		} else if !errors.Is(err, os.ErrNotExist) {
			errs = append(errs, fmt.Errorf("stat %s: %w", src, err))
		}
	}
	return errors.Join(errs...)
}

func dryRunPatch(ctx context.Context, opts Options, r runner) error {
	fmt.Fprintln(opts.Stdout, "==> Validating patch (dry-run)")
	return r.Run(ctx, opts.RootDir, nil, "patch", "--dry-run", "-p1", "-d", opts.ProviderDir, "-i", opts.PatchFile)
}

func applyPatch(ctx context.Context, opts Options, r runner) error {
	fmt.Fprintln(opts.Stdout, "==> Applying patch")
	return r.Run(ctx, opts.RootDir, nil, "patch", "-p1", "-d", opts.ProviderDir, "-i", opts.PatchFile)
}

func backupModuleFiles(opts Options) error {
	if err := os.RemoveAll(opts.StateDir); err != nil {
		return fmt.Errorf("reset state directory %s: %w", opts.StateDir, err)
	}
	if err := os.MkdirAll(opts.StateDir, 0o755); err != nil {
		return fmt.Errorf("create state directory %s: %w", opts.StateDir, err)
	}
	for _, name := range []string{"go.mod", "go.sum"} {
		if err := copyFile(filepath.Join(opts.RootDir, name), filepath.Join(opts.StateDir, name)); err != nil {
			return fmt.Errorf("backup %s: %w", name, err)
		}
	}
	return nil
}

func materializeNoForkSources(opts Options) error {
	for _, source := range noForkSourceFiles {
		if err := copyFile(filepath.Join(opts.RootDir, source[0]), filepath.Join(opts.RootDir, source[1])); err != nil {
			return fmt.Errorf("materialize no-fork source %s: %w", source[1], err)
		}
	}
	return nil
}

func removeNoForkSources(opts Options) error {
	var errs []error
	for _, source := range noForkSourceFiles {
		err := os.Remove(filepath.Join(opts.RootDir, source[1]))
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.OpenFile(dst, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}
	return out.Close()
}

func ensurePatchFile(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("patch file %s: %w", path, err)
	}
	if info.IsDir() {
		return fmt.Errorf("patch file %s is a directory", path)
	}
	return nil
}

func goEnv(opts Options) []string {
	return []string{
		"GOCACHE=" + opts.GoCache,
		"GOMODCACHE=" + opts.GoModCache,
		"GOPATH=" + opts.GoPath,
		"GOFLAGS=" + strings.TrimSpace(os.Getenv("GOFLAGS")+" "+noForkGoFlags),
	}
}

func (o Options) validateForPatch() error {
	var missing []string
	if o.RootDir == "" {
		missing = append(missing, "root-dir")
	}
	if o.ProviderVersion == "" {
		missing = append(missing, "provider-version")
	}
	if o.ProviderDir == "" {
		missing = append(missing, "provider-dir")
	}
	if o.StateDir == "" {
		missing = append(missing, "state-dir")
	}
	if o.PatchFile == "" {
		missing = append(missing, "patch-file")
	}
	if o.RepoURL == "" {
		missing = append(missing, "repo-url")
	}
	if o.ModulePath == "" {
		missing = append(missing, "module-path")
	}
	if len(missing) != 0 {
		return fmt.Errorf("missing required option(s): %s", strings.Join(missing, ", "))
	}
	return nil
}

func normalizeOptions(opts Options) Options {
	if opts.RootDir == "" {
		if wd, err := os.Getwd(); err == nil {
			opts.RootDir = wd
		}
	}
	opts.RootDir = absPath(opts.RootDir)
	if opts.ProviderVersion == "" {
		opts.ProviderVersion = os.Getenv("TERRAFORM_PROVIDER_VERSION")
	}
	if opts.ProviderDir == "" {
		opts.ProviderDir = filepath.Join(opts.RootDir, ".work", "nofork", "terraform-provider-oci")
	}
	if opts.StateDir == "" {
		opts.StateDir = filepath.Join(opts.RootDir, ".work", "nofork", "state")
	}
	if opts.PatchFile == "" {
		opts.PatchFile = filepath.Join(opts.RootDir, defaultPatchFile)
	}
	if opts.RepoURL == "" {
		opts.RepoURL = defaultRepoURL
	}
	if opts.ModulePath == "" {
		opts.ModulePath = defaultModulePath
	}
	if opts.GoCache == "" {
		opts.GoCache = filepath.Join(opts.RootDir, ".cache", "nofork-go-build")
	}
	if opts.GoModCache == "" {
		opts.GoModCache = filepath.Join(opts.RootDir, ".cache", "nofork-go-mod")
	}
	if opts.GoPath == "" {
		opts.GoPath = filepath.Join(opts.RootDir, ".work", "nofork-gopath")
	}
	if opts.Stdout == nil {
		opts.Stdout = io.Discard
	}
	if opts.Stderr == nil {
		opts.Stderr = io.Discard
	}
	opts.ProviderDir = absPath(opts.ProviderDir)
	opts.StateDir = absPath(opts.StateDir)
	opts.PatchFile = absPath(opts.PatchFile)
	opts.GoCache = absPath(opts.GoCache)
	opts.GoModCache = absPath(opts.GoModCache)
	opts.GoPath = absPath(opts.GoPath)
	return opts
}

func absPath(path string) string {
	if path == "" {
		return ""
	}
	abs, err := filepath.Abs(path)
	if err != nil {
		return path
	}
	return abs
}

func providerTag(version string) string {
	return "v" + cleanVersion(version)
}

func cleanVersion(version string) string {
	return strings.TrimPrefix(version, "v")
}

func makeWritable(root string) error {
	if _, err := os.Stat(root); errors.Is(err, os.ErrNotExist) {
		return nil
	}
	return filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return os.Chmod(path, 0o755)
		}
		return os.Chmod(path, 0o644)
	})
}

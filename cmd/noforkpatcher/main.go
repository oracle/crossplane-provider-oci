package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/oracle/provider-oci/internal/noforkpatch"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "noforkpatcher: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		usage()
		return fmt.Errorf("missing command")
	}
	root, err := os.Getwd()
	if err != nil {
		return err
	}
	switch args[0] {
	case "apply":
		opts, err := parseOptions(root, "apply", args[1:])
		if err != nil {
			return err
		}
		return noforkpatch.Apply(context.Background(), opts)
	case "validate":
		opts, err := parseOptions(root, "validate", args[1:])
		if err != nil {
			return err
		}
		return noforkpatch.Validate(context.Background(), opts)
	case "clean":
		opts, err := parseOptions(root, "clean", args[1:])
		if err != nil {
			return err
		}
		return noforkpatch.Clean(opts)
	case "-h", "--help", "help":
		usage()
		return nil
	default:
		usage()
		return fmt.Errorf("unknown command %q", args[0])
	}
}

func parseOptions(root, command string, args []string) (noforkpatch.Options, error) {
	opts := noforkpatch.DefaultOptions(root)
	fs := flag.NewFlagSet(command, flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	fs.StringVar(&opts.RootDir, "root-dir", opts.RootDir, "repository root directory")
	fs.StringVar(&opts.ProviderVersion, "provider-version", opts.ProviderVersion, "terraform-provider-oci version without leading v")
	fs.StringVar(&opts.ProviderDir, "provider-dir", opts.ProviderDir, "workspace-local patched terraform-provider-oci clone directory")
	fs.StringVar(&opts.StateDir, "state-dir", opts.StateDir, "directory used to back up go.mod and go.sum")
	fs.StringVar(&opts.PatchFile, "patch-file", opts.PatchFile, "patch file to apply to terraform-provider-oci")
	fs.StringVar(&opts.RepoURL, "repo-url", opts.RepoURL, "terraform-provider-oci git repository URL")
	fs.StringVar(&opts.ModulePath, "module-path", opts.ModulePath, "terraform-provider-oci Go module path")
	fs.StringVar(&opts.GoCache, "gocache", opts.GoCache, "GOCACHE used while tidying patched dependency graph")
	fs.StringVar(&opts.GoModCache, "gomodcache", opts.GoModCache, "GOMODCACHE used while tidying patched dependency graph")
	fs.StringVar(&opts.GoPath, "gopath", opts.GoPath, "GOPATH used while tidying patched dependency graph")
	if err := fs.Parse(args); err != nil {
		return noforkpatch.Options{}, err
	}
	opts.Stdout = os.Stdout
	opts.Stderr = os.Stderr
	return opts, nil
}

func usage() {
	fmt.Fprintln(os.Stderr, `Usage:
  go run ./cmd/noforkpatcher apply [flags]
  go run ./cmd/noforkpatcher validate [flags]
  go run ./cmd/noforkpatcher clean [flags]

Common flags:
  --provider-version  Terraform OCI provider version, e.g. 8.12.0
  --provider-dir      Patched provider clone directory
  --state-dir         Backup directory for go.mod and go.sum
  --patch-file        Patch file to apply
  --repo-url          Terraform provider OCI repository URL
  --gocache           Build cache used by go mod tidy
  --gomodcache        Module cache used by go mod tidy
  --gopath            GOPATH used by go mod tidy`)
}

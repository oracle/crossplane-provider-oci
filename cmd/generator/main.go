/*
Copyright 2021 Upbound Inc.
*/

package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/crossplane/upjet/pkg/pipeline"

	"github.com/oracle/provider-oci/config"
	"gopkg.in/alecthomas/kingpin.v2"
)

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

func copyDir(src, dst string) error {
	return filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		if d.IsDir() {
			return os.MkdirAll(target, 0755)
		}
		return copyFile(path, target)
	})
}

func main() {
	var (
		app          = kingpin.New(filepath.Base(os.Args[0]), "OCI provider code generator.").DefaultEnvars()
		rootDir      = app.Arg("rootDir", "Root directory of the project").Required().String()
		providerType = app.Flag("type", "Provider type (regular or self-contained)").Short('t').Default("regular").Enum("regular", "self-contained")
	)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	absRootDir, err := filepath.Abs(*rootDir)
	if err != nil {
		panic(fmt.Sprintf("cannot calculate the absolute path with %s", *rootDir))
	}

	if *providerType == "self-contained" {
		tmpRoot, err := os.MkdirTemp("", "oci-self-contained-*")
		if err != nil {
			panic(fmt.Sprintf("cannot create temp dir: %v", err))
		}
		defer os.RemoveAll(tmpRoot)
		// Copy hack directory to tmpRoot
		hackSrc := filepath.Join(absRootDir, "hack")
		hackDst := filepath.Join(tmpRoot, "hack")
		if err := copyDir(hackSrc, hackDst); err != nil {
			panic(fmt.Sprintf("copying hack dir: %v", err))
		}
		pipeline.Run(config.GetSelfContainedProvider(), tmpRoot)
		src := filepath.Join(tmpRoot, "examples-generated")
		dst := filepath.Join(absRootDir, "self-contained-examples-generated")
		_ = os.RemoveAll(dst)
		if err := copyDir(src, dst); err != nil {
			panic(fmt.Sprintf("copying examples: %v", err))
		}
		return
	}

	pipeline.Run(config.GetProvider(), absRootDir)
}

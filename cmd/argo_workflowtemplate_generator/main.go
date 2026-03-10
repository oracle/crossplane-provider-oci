package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/oracle/provider-oci/internal/argo/workflowtemplategenerator"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	version, err := parseVersionArg(os.Args)
	if err != nil {
		return err
	}

	rootDir, err := os.Getwd()
	if err != nil {
		return err
	}

	cfg := workflowtemplategenerator.NewConfig(rootDir, version)
	return workflowtemplategenerator.Run(cfg)
}

func parseVersionArg(args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("usage: go run main.go <version>")
	}

	version := strings.TrimSpace(args[1])
	if version == "" {
		return "", fmt.Errorf("version must not be empty")
	}

	return version, nil
}

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/oracle/provider-oci/internal/argo/workflowgenerator"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	version, serviceNames, err := parseArgs(os.Args)
	if err != nil {
		return err
	}

	rootDir, err := os.Getwd()
	if err != nil {
		return err
	}

	cfg := workflowgenerator.NewConfig(rootDir, version, serviceNames)
	return workflowgenerator.Run(cfg)
}

func parseArgs(args []string) (string, []string, error) {
	if len(args) < 2 {
		return "", nil, fmt.Errorf("usage: go run main.go <version> <optional_service_list>")
	}

	version := strings.TrimSpace(args[1])
	if version == "" {
		return "", nil, fmt.Errorf("version must not be empty")
	}

	serviceNames := make([]string, 0, len(args)-2)
	seen := make(map[string]struct{}, len(args)-2)
	for _, serviceName := range args[2:] {
		normalized := strings.TrimSpace(serviceName)
		if normalized == "" {
			continue
		}
		if _, exists := seen[normalized]; exists {
			continue
		}
		seen[normalized] = struct{}{}
		serviceNames = append(serviceNames, normalized)
	}

	return version, serviceNames, nil
}

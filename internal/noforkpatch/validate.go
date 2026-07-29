package noforkpatch

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type validationRule struct {
	name    string
	paths   []string
	pattern *regexp.Regexp
}

type requiredValidationRule struct {
	name    string
	path    string
	pattern *regexp.Regexp
}

// ValidatePatchedTree performs semantic checks that protect the no-fork runtime
// from upstream provider-global state hazards.
func ValidatePatchedTree(providerDir string) error {
	rules := []validationRule{
		{
			name: "unsafe ConfigureClientVar call site",
			paths: []string{
				"internal/provider",
				"internal/service",
				"internal/client",
			},
			pattern: regexp.MustCompile(`ConfigureClientVar[[:space:]]*\(`),
		},
		{
			name: "unsafe ConfigureClientVar assignment",
			paths: []string{
				"internal/provider",
				"internal/service",
				"internal/client",
			},
			pattern: regexp.MustCompile(`tf_client\.ConfigureClientVar[[:space:]]*=`),
		},
		{
			name: "service-level retry global mutation",
			paths: []string{
				"internal/service",
			},
			pattern: regexp.MustCompile(`tfresource\.(ShortRetryTime|LongRetryTime|ConfiguredRetryDuration)[[:space:]]*=`),
		},
		{
			name: "provider-global tfresource mutation",
			paths: []string{
				"internal/provider",
			},
			pattern: regexp.MustCompile(`^[[:space:]]*tf_resource\.(DefinedTagsToSuppress|RealmSpecificServiceEndpointTemplateEnabled|DualStackEndpointTemplateEnabled|ShortRetryTime|LongRetryTime|ConfiguredRetryDuration)[[:space:]]*=`),
		},
		{
			name: "provider-global retry config mutation",
			paths: []string{
				"internal/provider",
			},
			pattern: regexp.MustCompile(`^[[:space:]]*.*tf_resource\.SetRetriesConfig[[:space:]]*\(`),
		},
		{
			name: "provider-global delete wait mutation",
			paths: []string{
				"internal/provider",
			},
			pattern: regexp.MustCompile(`^[[:space:]]*AvoidWaitingForDeleteTarget[[:space:]]*=`),
		},
		{
			name: "runtime environment mutation",
			paths: []string{
				"internal/provider",
				"internal/service",
				"internal/client",
			},
			pattern: regexp.MustCompile(`^[[:space:]]*.*os\.Setenv[[:space:]]*\(`),
		},
		{
			name: "mutable SDKv2 provider singleton",
			paths: []string{
				"internal/provider",
			},
			pattern: regexp.MustCompile(`^[[:space:]]*var[[:space:]]+ociProvider[[:space:]]+\*schema\.Provider`),
		},
		{
			name: "mutable Terraform CLI version global",
			paths: []string{
				"internal/provider",
			},
			pattern: regexp.MustCompile(`^[[:space:]]*(var[[:space:]]+TerraformCLIVersion|TerraformCLIVersion[[:space:]]*=)`),
		},
		{
			name: "cached service client region mutation",
			paths: []string{
				"internal/service",
			},
			pattern: regexp.MustCompile(`s\.Client\.SetRegion[[:space:]]*\(`),
		},
	}

	var errs []error
	for _, rule := range rules {
		findings, err := findMatches(providerDir, rule)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		if len(findings) != 0 {
			errs = append(errs, fmt.Errorf("%s after patch:\n%s", rule.name, strings.Join(findings, "\n")))
		}
	}

	required := []requiredValidationRule{
		{
			name:    "instance-local SDKv2 provider factory",
			path:    "internal/provider/provider.go",
			pattern: regexp.MustCompile(`p[[:space:]]*:=[[:space:]]*&schema\.Provider`),
		},
		{
			name:    "instance-local Terraform version propagation",
			path:    "internal/provider/provider.go",
			pattern: regexp.MustCompile(`buildConfigureClientFn\(configProvider[^\n]*terraformVersion[[:space:]]+string`),
		},
		{
			name:    "one-time provider registration",
			path:    "internal/provider/provider.go",
			pattern: regexp.MustCompile(`providerRegistrationsOnce\.Do[[:space:]]*\(`),
		},
		{
			name:    "frozen provider registration",
			path:    "internal/provider/provider.go",
			pattern: regexp.MustCompile(`tf_resource\.FreezeRegistrations[[:space:]]*\(`),
		},
		{
			name:    "copied SDKv2 registration maps",
			path:    "internal/provider/provider.go",
			pattern: regexp.MustCompile(`return[[:space:]]+cloneSDKv2ResourceMap\(`),
		},
		{
			name:    "instance-local SDKv2 resource schemas",
			path:    "internal/provider/provider.go",
			pattern: regexp.MustCompile(`result\[name\][[:space:]]*=[[:space:]]*cloneSDKv2Resource\(resource\)`),
		},
		{
			name:    "Framework configuration error diagnostic",
			path:    "internal/provider/provider_framework.go",
			pattern: regexp.MustCompile(`resp\.Diagnostics\.AddError[[:space:]]*\(`),
		},
		{
			name:    "isolated MySQL backup copy client",
			path:    "internal/service/mysql/mysql_mysql_backup_resource.go",
			pattern: regexp.MustCompile(`copyClient\.CopyBackup[[:space:]]*\(`),
		},
		{
			name:    "ProviderConfig-specific MySQL client configuration",
			path:    "internal/service/mysql/helpers_mysql.go",
			pattern: regexp.MustCompile(`s\.ConfigureClient\(&dbBackupClient\.BaseClient\)`),
		},
	}
	for _, rule := range required {
		content, err := os.ReadFile(filepath.Join(providerDir, rule.path))
		if err != nil {
			errs = append(errs, fmt.Errorf("read required no-fork invariant %s: %w", rule.path, err))
			continue
		}
		if !rule.pattern.Match(content) {
			errs = append(errs, fmt.Errorf("required no-fork invariant %q is missing from %s", rule.name, rule.path))
		}
	}
	return errors.Join(errs...)
}

func findMatches(providerDir string, rule validationRule) ([]string, error) {
	var findings []string
	for _, rel := range rule.paths {
		root := filepath.Join(providerDir, rel)
		if _, err := os.Stat(root); errors.Is(err, os.ErrNotExist) {
			continue
		} else if err != nil {
			return nil, fmt.Errorf("stat %s: %w", root, err)
		}
		err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
				return nil
			}
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			lines := strings.Split(string(content), "\n")
			for i, line := range lines {
				if rule.pattern.MatchString(line) {
					displayPath, err := filepath.Rel(providerDir, path)
					if err != nil {
						displayPath = path
					}
					findings = append(findings, fmt.Sprintf("%s:%d:%s", displayPath, i+1, strings.TrimSpace(line)))
				}
			}
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("scan %s: %w", root, err)
		}
	}
	return findings, nil
}

package config

import (
	"regexp"
	"strings"

	"github.com/crossplane/upjet/pkg/config"
)

type References struct {
	TerraformName string
	Overwrite     bool
}

// StaticReferenceInjector uses a single patterns map for all rules
type StaticReferenceInjector struct {
	patterns map[string]References
}

// ReferenceInjector identifies resource prefixes from values assigned to fields, and adds as TerraformName
// Example: subnet_id = oci_core_subnet.test_subnet.id
// will be configured as "TerraformName: "oci_core_subnet"
// It identifies prefix and matches with cross resources in provider-metadata.yaml
// Problem: If any field is declared as var.compartment_id leads to failure since `var` prefix doesn't match with provider-metadata resources.
// Solution: Rule based StaticReferenceInjector to complement ReferenceInjector.

func NewStaticReferenceInjector() *StaticReferenceInjector {
	return &StaticReferenceInjector{
		patterns: map[string]References{
			// Identity Resources. Sample regex rule
			"^compartment_id$":   {TerraformName: "oci_identity_compartment"},
			"^compartment_ocid$": {TerraformName: "oci_identity_compartment"},
		},
	}
}

func (sr *StaticReferenceInjector) InjectReferences(resources map[string]*config.Resource) error {
	for _, r := range resources {
		for attr := range r.TerraformResource.Schema {
			if strings.Contains(attr, ".") || strings.Contains(attr, "[") {
				continue // skip nested
			}
			for pattern, target := range sr.patterns {
				re := regexp.MustCompile(pattern)
				if re.MatchString(attr) {
					if _, already := r.References[attr]; already && !target.Overwrite {
						continue // Do not overwrite unless specified
					}
					r.References[attr] = config.Reference{
						TerraformName: target.TerraformName,
					}
					break
				}
			}
		}
	}
	return nil
}

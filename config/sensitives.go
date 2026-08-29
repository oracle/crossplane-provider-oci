package config

import (
	"slices"
	"strings"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var autoSensitiveNameKeywords = []string{
	"secret",
	"password",
	"passphrase",
	"credential",
}

var autoSensitiveExcludedNameKeywords = []string{
	"_id",
	"_ocid",
	"_name",
}

// AutoSensitiveFieldConfiguration marks high-confidence computed-only string
// outputs as sensitive so Upjet emits them as connection details.
func AutoSensitiveFieldConfiguration() ujconfig.ResourceOption {
	return func(r *ujconfig.Resource) {
		if r == nil || r.TerraformResource == nil {
			return
		}
		markAutoSensitiveFields(r.TerraformResource.Schema)
	}
}

func markAutoSensitiveFields(fields map[string]*schema.Schema) {
	for name, sch := range fields {
		if sch == nil {
			continue
		}

		if sch.Computed && !sch.Optional && !sch.Required && sch.Type == schema.TypeString {
			name = strings.ToLower(name)
			if slices.ContainsFunc(autoSensitiveNameKeywords, func(keyword string) bool {
				return strings.Contains(name, keyword)
			}) && !slices.ContainsFunc(autoSensitiveExcludedNameKeywords, func(keyword string) bool {
				return strings.Contains(name, keyword)
			}) {
				sch.Sensitive = true
			}
		}
		if nested, ok := sch.Elem.(*schema.Resource); ok {
			markAutoSensitiveFields(nested.Schema)
		}
	}
}

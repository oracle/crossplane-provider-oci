package config

import (
	"testing"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestAutoSensitiveFieldConfiguration(t *testing.T) {
	tests := map[string]struct {
		fields        map[string]*schema.Schema
		field         string
		wantSensitive bool
	}{
		"marks computed-only secret_key": {
			fields: map[string]*schema.Schema{
				"secret_key": computedStringSchema(),
			},
			field:         "secret_key",
			wantSensitive: true,
		},
		"marks computed-only access_token": {
			fields: map[string]*schema.Schema{
				"access_token": computedStringSchema(),
			},
			field:         "access_token",
			wantSensitive: true,
		},
		"marks computed-only private_key": {
			fields: map[string]*schema.Schema{
				"private_key": computedStringSchema(),
			},
			field:         "private_key",
			wantSensitive: true,
		},
		"does not mark required secret_key": {
			fields: map[string]*schema.Schema{
				"secret_key": {Type: schema.TypeString, Required: true},
			},
			field: "secret_key",
		},
		"does not mark optional secret_key": {
			fields: map[string]*schema.Schema{
				"secret_key": {Type: schema.TypeString, Optional: true},
			},
			field: "secret_key",
		},
		"does not mark optional computed secret_key": {
			fields: map[string]*schema.Schema{
				"secret_key": {Type: schema.TypeString, Optional: true, Computed: true},
			},
			field: "secret_key",
		},
		"does not mark secret_id": {
			fields: map[string]*schema.Schema{
				"secret_id": computedStringSchema(),
			},
			field: "secret_id",
		},
		"does not mark public_key": {
			fields: map[string]*schema.Schema{
				"public_key": computedStringSchema(),
			},
			field: "public_key",
		},
		"does not mark correlation_token": {
			fields: map[string]*schema.Schema{
				"correlation_token": computedStringSchema(),
			},
			field: "correlation_token",
		},
		"does not mark last accepted request token": {
			fields: map[string]*schema.Schema{
				"last_accepted_request_token": computedStringSchema(),
			},
			field: "last_accepted_request_token",
		},
		"does not mark is_secret": {
			fields: map[string]*schema.Schema{
				"is_secret": computedStringSchema(),
			},
			field: "is_secret",
		},
		"preserves already-sensitive fields": {
			fields: map[string]*schema.Schema{
				"opaque_value": {Type: schema.TypeString, Computed: true, Sensitive: true},
			},
			field:         "opaque_value",
			wantSensitive: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			AutoSensitiveFieldConfiguration()(&ujconfig.Resource{
				TerraformResource: &schema.Resource{Schema: tc.fields},
			})

			if got := tc.fields[tc.field].Sensitive; got != tc.wantSensitive {
				t.Fatalf("Sensitive = %v, want %v", got, tc.wantSensitive)
			}
		})
	}
}

func TestAutoSensitiveFieldConfigurationNestedSchema(t *testing.T) {
	fields := map[string]*schema.Schema{
		"nested": {
			Type: schema.TypeList,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"client_secret": computedStringSchema(),
			}},
		},
	}

	AutoSensitiveFieldConfiguration()(&ujconfig.Resource{
		TerraformResource: &schema.Resource{Schema: fields},
	})

	nested := fields["nested"].Elem.(*schema.Resource)
	if !nested.Schema["client_secret"].Sensitive {
		t.Fatal("expected nested computed-only client_secret to be sensitive")
	}
}

func computedStringSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
}

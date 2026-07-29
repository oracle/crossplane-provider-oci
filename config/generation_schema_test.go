/*
 * Copyright (c) 2026 Oracle and/or its affiliates
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestApplyTerraformJSONGenerationSchemasPreservesCompleteShape(t *testing.T) {
	const resourceName = "test_resource"
	baseline := &schema.Resource{Schema: map[string]*schema.Schema{
		"count": {
			Type:        schema.TypeFloat,
			Optional:    true,
			Description: "JSON number",
		},
		"metadata": {
			Type:     schema.TypeMap,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"nested": {
			Type:     schema.TypeList,
			Computed: true,
			MaxItems: 1,
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"secret": {
					Type:      schema.TypeString,
					Sensitive: true,
					Computed:  true,
				},
			}},
		},
	}}
	live := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"count": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"metadata": {
				Type: schema.TypeMap,
				Elem: schema.TypeString,
			},
			"nested": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"different": {
						Type: schema.TypeBool,
					},
				}},
			},
			"sdk_only": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		CreateContext: func(_ context.Context, _ *schema.ResourceData, _ any) diag.Diagnostics {
			return nil
		},
	}
	provider := &schema.Provider{ResourcesMap: map[string]*schema.Resource{
		resourceName: live,
	}}

	applyGenerationResourceSchemas(provider, map[string]*schema.Resource{
		resourceName: baseline,
	})

	if !reflect.DeepEqual(live.Schema, baseline.Schema) {
		t.Fatalf("generation schema differs from Terraform JSON baseline")
	}
	if live.CreateContext == nil {
		t.Fatal("resource operation callbacks were not preserved")
	}
	live.Schema["count"].Description = "changed"
	if baseline.Schema["count"].Description == "changed" {
		t.Fatal("generation schema shares nested pointers with the Terraform JSON baseline")
	}
}

func TestCloneGenerationSchemaPreservesAllSchemaProperties(t *testing.T) {
	source := &schema.Schema{
		Type:          schema.TypeList,
		Optional:      true,
		Computed:      true,
		Sensitive:     true,
		ForceNew:      true,
		MinItems:      1,
		MaxItems:      2,
		Description:   "complete shape",
		ComputedWhen:  []string{"computed"},
		ConflictsWith: []string{"other"},
		ExactlyOneOf:  []string{"one"},
		AtLeastOneOf:  []string{"some"},
		RequiredWith:  []string{"required"},
		Elem: &schema.Resource{Schema: map[string]*schema.Schema{
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
		}},
	}

	cloned := cloneGenerationSchema(source)
	if !reflect.DeepEqual(cloned, source) {
		t.Fatal("cloned schema does not preserve the complete source shape")
	}
	cloned.Elem.(*schema.Resource).Schema["value"].Description = "changed"
	if source.Elem.(*schema.Resource).Schema["value"].Description == "changed" {
		t.Fatal("cloned nested resource shares schema pointers with source")
	}
	cloned.ConflictsWith[0] = "changed"
	if source.ConflictsWith[0] == "changed" {
		t.Fatal("cloned schema shares validation slices with source")
	}
}

/*
 * Copyright (c) 2026 Oracle and/or its affiliates
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"testing"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestConfigureSDKv2ScalarMapServerSideApplyMergeStrategies(t *testing.T) {
	r := &ujconfig.Resource{
		TerraformResource: &schema.Resource{Schema: map[string]*schema.Schema{
			"defined_tags": {
				Type: schema.TypeMap,
				Elem: schema.TypeString,
			},
			"nil_element_map": {
				Type: schema.TypeMap,
			},
			"legacy_map": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"sensitive_map": {
				Type:      schema.TypeMap,
				Elem:      schema.TypeString,
				Sensitive: true,
			},
			"object_map": {
				Type: schema.TypeMap,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"metadata": {
						Type: schema.TypeMap,
						Elem: schema.TypeString,
					},
				}},
			},
		}},
		ServerSideApplyMergeStrategies: ujconfig.ServerSideApplyMergeStrategies{
			"defined_tags": {
				MapMergeStrategy: ujconfig.MapTypeAtomic,
			},
		},
	}

	configureSDKv2ScalarMapServerSideApplyMergeStrategiesForResource(r)

	for _, field := range []string{"nil_element_map", "legacy_map", "object_map.metadata"} {
		strategy, ok := r.ServerSideApplyMergeStrategies[field]
		if !ok {
			t.Fatalf("%s strategy is missing", field)
		}
		if strategy.MapMergeStrategy != ujconfig.MapTypeGranular {
			t.Fatalf("%s map merge strategy = %q, want %q", field, strategy.MapMergeStrategy, ujconfig.MapTypeGranular)
		}
	}

	if got := r.ServerSideApplyMergeStrategies["defined_tags"].MapMergeStrategy; got != ujconfig.MapTypeAtomic {
		t.Fatalf("existing defined_tags strategy = %q, want %q", got, ujconfig.MapTypeAtomic)
	}
	for _, field := range []string{"sensitive_map"} {
		if _, ok := r.ServerSideApplyMergeStrategies[field]; ok {
			t.Fatalf("%s unexpectedly received an SDKv2 scalar map strategy", field)
		}
	}
}

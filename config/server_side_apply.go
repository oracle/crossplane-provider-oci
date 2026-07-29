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
	"strings"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// configureSDKv2ScalarMapServerSideApplyMergeStrategies restores the default
// granular server-side apply semantics for scalar maps sourced from an SDKv2
// provider.
//
// Upjet automatically adds granular markers when a scalar map element is
// represented by *schema.Schema. The embedded OCI SDKv2 provider represents
// equivalent scalar map elements as schema.ValueType instead. Configure the
// merge strategy explicitly so generated APIs and CRDs preserve the same
// ownership behavior regardless of that schema representation.
func configureSDKv2ScalarMapServerSideApplyMergeStrategies(p *ujconfig.Provider) {
	for _, r := range p.Resources {
		configureSDKv2ScalarMapServerSideApplyMergeStrategiesForResource(r)
	}
}

func configureSDKv2ScalarMapServerSideApplyMergeStrategiesForResource(r *ujconfig.Resource) {
	if r == nil || r.TerraformResource == nil {
		return
	}
	if r.ServerSideApplyMergeStrategies == nil {
		r.ServerSideApplyMergeStrategies = make(ujconfig.ServerSideApplyMergeStrategies)
	}
	walkSDKv2SchemaForScalarMaps(r.TerraformResource, "", r)
}

func walkSDKv2SchemaForScalarMaps(s *schema.Resource, prefix string, r *ujconfig.Resource) {
	for name, f := range s.Schema {
		if f == nil {
			continue
		}

		path := joinTerraformFieldPath(prefix, name)
		if isSDKv2ScalarMap(f) {
			if _, configured := r.ServerSideApplyMergeStrategies[path]; !configured {
				r.ServerSideApplyMergeStrategies[path] = ujconfig.MergeStrategy{
					MapMergeStrategy: ujconfig.MapTypeGranular,
				}
			}
		}

		if nested, ok := f.Elem.(*schema.Resource); ok {
			walkSDKv2SchemaForScalarMaps(nested, path, r)
		}
	}
}

func isSDKv2ScalarMap(f *schema.Schema) bool {
	if f.Sensitive || f.Type != schema.TypeMap {
		return false
	}

	// SDKv2 permits nil for a TypeMap element when the map holds strings.
	// The OCI provider uses that form in a few nested patch-operation schemas.
	if f.Elem == nil {
		return true
	}

	var element schema.ValueType
	switch typed := f.Elem.(type) {
	case schema.ValueType:
		element = typed
	case *schema.Schema:
		if typed == nil {
			return false
		}
		element = typed.Type
	default:
		return false
	}

	switch element {
	case schema.TypeString, schema.TypeBool, schema.TypeInt, schema.TypeFloat:
		return true
	default:
		return false
	}
}

func joinTerraformFieldPath(prefix, name string) string {
	if prefix == "" {
		return name
	}
	return strings.Join([]string{prefix, name}, ".")
}

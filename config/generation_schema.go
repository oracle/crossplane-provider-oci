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
	"fmt"

	conversiontfjson "github.com/crossplane/upjet/v2/pkg/types/conversion/tfjson"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// applyTerraformJSONGenerationSchemas preserves the public API and CRD schema
// while Upjet generates SDKv2 no-fork controllers. Upjet normally replaces its
// Terraform JSON-derived resource schemas with the live SDKv2 schemas for
// SDK-routed resources. The two representations differ in API-shaping details,
// including number and map types, field modes, sensitive fields, and nested
// field membership.
//
// This function must only receive a fresh provider instance used by the
// generator. Runtime provider instances must retain their native SDKv2 schemas.
// Resource operation callbacks and other runtime behavior remain on each live
// SDKv2 resource; only the schema used by Upjet's generation pipeline is
// replaced.
func applyTerraformJSONGenerationSchemas(provider *schema.Provider, rawProviderSchema []byte) {
	if provider == nil {
		panic("cannot apply Terraform JSON generation schemas to a nil provider")
	}

	applyGenerationResourceSchemas(provider, terraformJSONResourceMap(rawProviderSchema))
}

func applyGenerationResourceSchemas(provider *schema.Provider, generationResources map[string]*schema.Resource) {
	for name, generationResource := range generationResources {
		sdkResource := provider.ResourcesMap[name]
		if sdkResource == nil || generationResource == nil {
			continue
		}
		sdkResource.Schema = cloneGenerationSchemaMap(generationResource.Schema)
	}
}

func terraformJSONResourceMap(rawProviderSchema []byte) map[string]*schema.Resource {
	providerSchemas := tfjson.ProviderSchemas{}
	if err := providerSchemas.UnmarshalJSON(rawProviderSchema); err != nil {
		panic(fmt.Sprintf("cannot parse Terraform JSON generation schema: %v", err))
	}
	if len(providerSchemas.Schemas) != 1 {
		panic(fmt.Sprintf("expected one Terraform provider generation schema, got %d", len(providerSchemas.Schemas)))
	}
	for _, provider := range providerSchemas.Schemas {
		return conversiontfjson.GetV2ResourceMap(provider.ResourceSchemas)
	}
	return nil
}

func cloneGenerationSchemaMap(source map[string]*schema.Schema) map[string]*schema.Schema {
	if source == nil {
		return nil
	}
	result := make(map[string]*schema.Schema, len(source))
	for name, field := range source {
		result[name] = cloneGenerationSchema(field)
	}
	return result
}

func cloneGenerationSchema(source *schema.Schema) *schema.Schema {
	if source == nil {
		return nil
	}
	result := *source
	result.ComputedWhen = append([]string(nil), source.ComputedWhen...)
	result.ConflictsWith = append([]string(nil), source.ConflictsWith...)
	result.ExactlyOneOf = append([]string(nil), source.ExactlyOneOf...)
	result.AtLeastOneOf = append([]string(nil), source.AtLeastOneOf...)
	result.RequiredWith = append([]string(nil), source.RequiredWith...)
	switch element := source.Elem.(type) {
	case *schema.Schema:
		result.Elem = cloneGenerationSchema(element)
	case *schema.Resource:
		result.Elem = cloneGenerationResource(element)
	}
	return &result
}

func cloneGenerationResource(source *schema.Resource) *schema.Resource {
	if source == nil {
		return nil
	}
	result := *source
	result.Schema = cloneGenerationSchemaMap(source.Schema)
	return &result
}

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
	"slices"
	"testing"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestRuntimeTerraformResourceManifests(t *testing.T) {
	allResources := map[string]string{}
	for service, resources := range runtimeTerraformResources {
		if service == "monolith" {
			continue
		}
		if len(resources) == 0 {
			t.Errorf("service %q has an empty runtime resource manifest", service)
		}
		for _, resourceName := range resources {
			if previousService, ok := allResources[resourceName]; ok {
				t.Errorf("Terraform resource %q is assigned to both %q and %q", resourceName, previousService, service)
			}
			allResources[resourceName] = service
		}
	}

	monolithResources := runtimeTerraformResources["monolith"]
	if len(monolithResources) != len(allResources) {
		t.Fatalf("monolith manifest has %d resources, want %d", len(monolithResources), len(allResources))
	}
	for resourceName := range allResources {
		if !slices.Contains(monolithResources, resourceName) {
			t.Errorf("monolith manifest does not contain %q", resourceName)
		}
	}
}

func TestRuntimeProviderContainsOnlyServiceResources(t *testing.T) {
	const service = "events"
	want := runtimeTerraformResources[service]
	if len(want) == 0 {
		t.Fatalf("service %q has no generated Terraform resources", service)
	}

	provider := GetProviderForRuntime(service)
	if len(provider.Resources) != len(want) {
		t.Fatalf("runtime provider has %d resources, want %d", len(provider.Resources), len(want))
	}
	if len(provider.TerraformProvider.ResourcesMap) != len(want) {
		t.Fatalf("runtime Terraform provider has %d resources, want %d", len(provider.TerraformProvider.ResourcesMap), len(want))
	}
	if len(provider.TerraformProvider.DataSourcesMap) != 0 {
		t.Fatalf("runtime Terraform provider has %d data sources, want 0", len(provider.TerraformProvider.DataSourcesMap))
	}
	predicate := SDKv2ResourcePredicateForRuntime(service)
	for _, resourceName := range want {
		if provider.Resources[resourceName] == nil {
			t.Errorf("runtime provider does not contain %q", resourceName)
		}
		if !predicate(resourceName) {
			t.Errorf("runtime SDKv2 predicate rejected %q", resourceName)
		}
	}
	if predicate("oci_core_vcn") {
		t.Error("events runtime SDKv2 predicate accepted unrelated resource oci_core_vcn")
	}
}

func TestGeneratedAPICompatibilityDoesNotMutateRuntimeSchema(t *testing.T) {
	const resourceName = "oci_adm_vulnerability_audit"
	const integerField = "vulnerable_artifacts_count"

	runtimeResource := GetProvider().Resources[resourceName]
	if runtimeResource == nil {
		t.Fatalf("runtime resource %q is missing", resourceName)
	}
	if got := runtimeResource.TerraformResource.Schema[integerField].Type; got != schema.TypeInt {
		t.Fatalf("runtime %s.%s type = %s, want TypeInt", resourceName, integerField, got)
	}

	generationResource := GetProviderForGeneration().Resources[resourceName]
	if generationResource == nil {
		t.Fatalf("generation resource %q is missing", resourceName)
	}
	if got := generationResource.TerraformResource.Schema[integerField].Type; got != schema.TypeFloat {
		t.Fatalf("generation %s.%s type = %s, want TypeFloat", resourceName, integerField, got)
	}

	const mapResourceName = "oci_objectstorage_object"
	const mapField = "metadata"
	runtimeMap := GetProvider().Resources[mapResourceName].TerraformResource.Schema[mapField]
	if _, ok := runtimeMap.Elem.(schema.ValueType); !ok {
		t.Fatalf("runtime %s.%s element = %T, want schema.ValueType", mapResourceName, mapField, runtimeMap.Elem)
	}
	generationMap := GetProviderForGeneration().Resources[mapResourceName].TerraformResource.Schema[mapField]
	if _, ok := generationMap.Elem.(*schema.Schema); !ok {
		t.Fatalf("generation %s.%s element = %T, want *schema.Schema", mapResourceName, mapField, generationMap.Elem)
	}
}

func TestProviderConfigIncludesPreviouslySkippedNoForkResources(t *testing.T) {
	for name, provider := range map[string]func() *ujconfig.Provider{
		"cluster":    GetProvider,
		"namespaced": GetProviderNamespaced,
	} {
		t.Run(name, func(t *testing.T) {
			pc := provider()

			savedSearch := pc.Resources["oci_management_dashboard_management_saved_search"]
			if savedSearch == nil {
				t.Fatal("ManagementSavedSearch resource was not generated")
			}
			if !savedSearch.ShouldUseTerraformPluginSDKClient() {
				t.Fatal("ManagementSavedSearch is not routed through SDKv2 no-fork")
			}
			freeformTags := savedSearch.TerraformResource.Schema["freeform_tags"]
			if freeformTags == nil {
				t.Fatal("ManagementSavedSearch freeform_tags schema is missing")
			}
			freeformTagsElem, ok := freeformTags.Elem.(schema.ValueType)
			if !ok {
				t.Fatalf("ManagementSavedSearch freeform_tags Elem = %T, want schema.ValueType", freeformTags.Elem)
			}
			if freeformTagsElem != schema.TypeMap {
				t.Fatalf("ManagementSavedSearch freeform_tags Elem type = %s, want map", freeformTagsElem)
			}

			opensearchCluster := pc.Resources["oci_opensearch_opensearch_cluster"]
			if opensearchCluster == nil {
				t.Fatal("OpensearchCluster resource was not generated")
			}
			if !opensearchCluster.ShouldUseTerraformPluginSDKClient() {
				t.Fatal("OpensearchCluster is not routed through SDKv2 no-fork")
			}
			samlConfig := opensearchCluster.TerraformResource.Schema["security_saml_config"]
			if samlConfig == nil {
				t.Fatal("OpensearchCluster security_saml_config schema is missing")
			}
			if samlConfig.Sensitive {
				t.Fatal("OpensearchCluster security_saml_config is sensitive, want only sensitive leaf fields")
			}
			samlResource, ok := samlConfig.Elem.(*schema.Resource)
			if !ok {
				t.Fatalf("OpensearchCluster security_saml_config Elem = %T, want *schema.Resource", samlConfig.Elem)
			}
			idpMetadataContent := samlResource.Schema["idp_metadata_content"]
			if idpMetadataContent == nil {
				t.Fatal("OpensearchCluster security_saml_config.idp_metadata_content schema is missing")
			}
			if idpMetadataContent.Sensitive {
				t.Fatal("OpensearchCluster security_saml_config.idp_metadata_content is sensitive, want existing plain API field")
			}
		})
	}
}

func BenchmarkProviderConstruction(b *testing.B) {
	b.Run("full", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			_ = GetProvider()
		}
	})
	b.Run("events", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			_ = GetProviderForRuntime("events")
		}
	})
	b.Run("networking", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			_ = GetProviderForRuntime("networking")
		}
	})
}

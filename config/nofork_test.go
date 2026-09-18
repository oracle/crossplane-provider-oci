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

import "testing"

func TestFrameworkRoutingIsDormant(t *testing.T) {
	if len(terraformPluginFrameworkIncludeList) != 0 {
		t.Fatalf("framework routing length = %d, want 0", len(terraformPluginFrameworkIncludeList))
	}
	if HasFrameworkResources() {
		t.Fatal("HasFrameworkResources() = true, want false")
	}
}

func TestSDKv2RoutingIncludesAllSchemaResources(t *testing.T) {
	for _, resource := range terraformResourceNamesFromSchema() {
		if !IsSDKv2Resource(resource) {
			t.Fatalf("IsSDKv2Resource(%q) = false, want true", resource)
		}
	}
}

func TestSDKv2RoutingRejectsUnknownResource(t *testing.T) {
	if IsSDKv2Resource("oci_does_not_exist") {
		t.Fatal("IsSDKv2Resource() = true for an unknown resource")
	}
}

func TestSDKv2ResourceLookupDoesNotAllocate(t *testing.T) {
	const resource = "oci_core_vcn"
	if !IsSDKv2Resource(resource) {
		t.Fatalf("IsSDKv2Resource(%q) = false, want true", resource)
	}
	if got := testing.AllocsPerRun(100, func() {
		IsSDKv2Resource(resource)
	}); got != 0 {
		t.Fatalf("IsSDKv2Resource() allocations = %v, want 0", got)
	}
}

func BenchmarkIsSDKv2Resource(b *testing.B) {
	for b.Loop() {
		IsSDKv2Resource("oci_core_vcn")
	}
}

func TestPreviouslySkippedNoForkResourcesUseSDKv2(t *testing.T) {
	for _, resource := range []string{
		"oci_management_dashboard_management_saved_search",
		"oci_opensearch_opensearch_cluster",
	} {
		if !IsSDKv2Resource(resource) {
			t.Fatalf("IsSDKv2Resource(%q) = false, want true", resource)
		}
		if resourceMatches(resource, ProblematicResources()) {
			t.Fatalf("%q matches ProblematicResources(), want generated through SDKv2 no-fork", resource)
		}
	}
}

func TestSDKv2RoutingIncludesRepresentativeResources(t *testing.T) {
	for _, resource := range []string{
		"oci_apigateway_gateway",
		"oci_autoscaling_auto_scaling_configuration",
		"oci_bastion_bastion",
		"oci_budget_budget",
		"oci_file_storage_file_system",
		"oci_core_subnet",
		"oci_core_vcn",
		"oci_identity_tag_namespace",
		"oci_kms_key",
		"oci_load_balancer_load_balancer",
		"oci_mysql_mysql_db_system",
		"oci_network_load_balancer_network_load_balancer",
		"oci_objectstorage_bucket",
		"oci_psql_db_system",
		"oci_queue_queue",
		"oci_streaming_stream",
		"oci_vault_secret",
		"oci_waf_web_app_firewall",
		"oci_containerengine_cluster",
	} {
		if !IsSDKv2Resource(resource) {
			t.Fatalf("IsSDKv2Resource(%q) = false, want true", resource)
		}
	}
}

/*
 * Copyright (c) 2025 Oracle and/or its affiliates
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

// Package main generates GroupMap entries for all OCI Terraform resources
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

// TerraformSchema represents the structure of the Terraform provider schema
type TerraformSchema struct {
	ProviderSchemas map[string]ProviderSchema `json:"provider_schemas"`
}

type ProviderSchema struct {
	ResourceSchemas map[string]interface{} `json:"resource_schemas"`
}

// ServiceMapping represents the mapping rules for services
type ServiceMapping struct {
	Prefix      string
	ServiceName string
	Rules       []MappingRule
}

type MappingRule struct {
	Pattern string
	Service string
}

// Define comprehensive service mappings based on OCI service structure
var serviceMappings = []ServiceMapping{
	// Core services that need special splitting
	{
		Prefix:      "core",
		ServiceName: "", // Will be determined by rules
		Rules: []MappingRule{
			// Compute resources
			{Pattern: "instance", Service: "compute"},
			{Pattern: "image", Service: "compute"},
			{Pattern: "dedicated_vm", Service: "compute"},
			{Pattern: "console", Service: "compute"},
			{Pattern: "shape", Service: "compute"},
			{Pattern: "app_catalog", Service: "compute"},
			{Pattern: "cluster_network", Service: "compute"},
			{Pattern: "compute_", Service: "compute"},
			{Pattern: "instance_configuration", Service: "compute"},
			{Pattern: "instance_console", Service: "compute"},
			{Pattern: "instance_pool", Service: "compute"},
			
			// Networking resources
			{Pattern: "vcn", Service: "networking"},
			{Pattern: "subnet", Service: "networking"},
			{Pattern: "vnic", Service: "networking"},
			{Pattern: "dhcp", Service: "networking"},
			{Pattern: "vlan", Service: "networking"},
			{Pattern: "internet_gateway", Service: "networking"},
			{Pattern: "nat_gateway", Service: "networking"},
			{Pattern: "service_gateway", Service: "networking"},
			{Pattern: "network_security", Service: "networking"},
			{Pattern: "route_table", Service: "networking"},
			{Pattern: "security_list", Service: "networking"},
			{Pattern: "private_ip", Service: "networking"},
			{Pattern: "public_ip", Service: "networking"},
			{Pattern: "byoip", Service: "networking"},
			{Pattern: "local_peering", Service: "networking"},
			{Pattern: "remote_peering", Service: "networking"},
			{Pattern: "ipv6", Service: "networking"},
			
			// Block storage resources
			{Pattern: "volume", Service: "blockstorage"},
			{Pattern: "boot_volume", Service: "blockstorage"},
			
			// Network connectivity resources
			{Pattern: "drg", Service: "networkconnectivity"},
			{Pattern: "cross_connect", Service: "networkconnectivity"},
			{Pattern: "virtual_circuit", Service: "networkconnectivity"},
			{Pattern: "cpe", Service: "networkconnectivity"},
			{Pattern: "ipsec", Service: "networkconnectivity"},
			
			// Monitoring resources
			{Pattern: "capture_filter", Service: "monitoring"},
			{Pattern: "vtap", Service: "monitoring"},
		},
	},
	
	// Database services - largest service group
	{Prefix: "database", ServiceName: "database"},
	
	// AI and Machine Learning services
	{Prefix: "ai", ServiceName: "ailanguage"},
	{Prefix: "generative", ServiceName: "generativeai"},
	{Prefix: "datascience", ServiceName: "datascience"},
	
	// Analytics and Data services
	{Prefix: "analytics", ServiceName: "analytics"},
	{Prefix: "datacatalog", ServiceName: "datacatalog"},
	{Prefix: "dataflow", ServiceName: "dataflow"},
	{Prefix: "dataintegration", ServiceName: "dataintegration"},
	{Prefix: "data", ServiceName: "datasafe"}, // Data Safe service
	
	// Developer services
	{Prefix: "devops", ServiceName: "devops"},
	{Prefix: "apigateway", ServiceName: "apigateway"},
	{Prefix: "appmgmt", ServiceName: "applicationmanagement"},
	{Prefix: "apm", ServiceName: "apm"},
	
	// Infrastructure services
	{Prefix: "containerengine", ServiceName: "containerengine"},
	{Prefix: "functions", ServiceName: "functions"},
	{Prefix: "compute", ServiceName: "computemanagement"},
	{Prefix: "autoscaling", ServiceName: "autoscaling"},
	
	// Storage services
	{Prefix: "objectstorage", ServiceName: "objectstorage"},
	{Prefix: "file", ServiceName: "filestorage"},
	{Prefix: "filestorage", ServiceName: "filestorage"},
	
	// Security services
	{Prefix: "kms", ServiceName: "kms"},
	{Prefix: "vault", ServiceName: "vault"},
	{Prefix: "bastion", ServiceName: "bastion"},
	{Prefix: "certificates", ServiceName: "certificatesmanagement"},
	{Prefix: "security", ServiceName: "securityattribute"},
	{Prefix: "waf", ServiceName: "waf"},
	{Prefix: "waas", ServiceName: "waas"},
	
	// Networking services
	{Prefix: "dns", ServiceName: "dns"},
	{Prefix: "network", ServiceName: "networking"}, // Generic network resources
	{Prefix: "load", ServiceName: "loadbalancer"},
	{Prefix: "networkfirewall", ServiceName: "networkfirewall"},
	
	// Monitoring and Management
	{Prefix: "monitoring", ServiceName: "monitoring"},
	{Prefix: "logging", ServiceName: "logging"},
	{Prefix: "events", ServiceName: "events"},
	{Prefix: "ons", ServiceName: "ons"},
	{Prefix: "streaming", ServiceName: "streaming"},
	{Prefix: "health", ServiceName: "healthchecks"},
	
	// Identity and Access
	{Prefix: "identity", ServiceName: "identity"},
	
	// Other services
	{Prefix: "budget", ServiceName: "budget"},
	{Prefix: "limits", ServiceName: "limits"},
	{Prefix: "usage", ServiceName: "usageapi"},
	{Prefix: "metering", ServiceName: "meteringcomputation"},
	{Prefix: "audit", ServiceName: "audit"},
	{Prefix: "announcements", ServiceName: "announcements"},
	
	// Platform services
	{Prefix: "resourcemanager", ServiceName: "resourcemanager"},
	{Prefix: "stack", ServiceName: "stackmonitoring"},
	{Prefix: "opsi", ServiceName: "operatoraccesscontrol"},
	{Prefix: "optimizer", ServiceName: "cloudguard"},
	
	// Specialized services
	{Prefix: "blockchain", ServiceName: "blockchain"},
	{Prefix: "bds", ServiceName: "bigdataservice"},
	{Prefix: "mysql", ServiceName: "mysql"},
	{Prefix: "nosql", ServiceName: "nosql"},
	{Prefix: "opensearch", ServiceName: "opensearch"},
	{Prefix: "redis", ServiceName: "redis"},
	{Prefix: "psql", ServiceName: "psql"},
	
	// Email and Communication
	{Prefix: "email", ServiceName: "emaildataplane"},
	{Prefix: "queue", ServiceName: "queue"},
	
	// Media services
	{Prefix: "media", ServiceName: "mediaservices"},
	
	// Marketplace and Integration
	{Prefix: "marketplace", ServiceName: "marketplace"},
	{Prefix: "integration", ServiceName: "integration"},
	{Prefix: "oda", ServiceName: "digitalassistant"},
	{Prefix: "oce", ServiceName: "contentexperience"},
	
	// VMware services
	{Prefix: "ocvp", ServiceName: "ocvs"},
	
	// Disaster Recovery
	{Prefix: "disaster", ServiceName: "disasterrecovery"},
	
	// Java Management
	{Prefix: "jms", ServiceName: "jms"},
	
	// OS Management
	{Prefix: "os", ServiceName: "osmanagement"},
	
	// Visual Builder
	{Prefix: "visual", ServiceName: "visualbuilder"},
	
	// Vulnerability Scanning
	{Prefix: "vulnerability", ServiceName: "vulnerabilityscanning"},
	
	// Golden Gate
	{Prefix: "golden", ServiceName: "goldengate"},
	
	// Fleet Management
	{Prefix: "fleet", ServiceName: "fleetappsmanagement"},
	
	// Fusion Apps
	{Prefix: "fusion", ServiceName: "fusionapps"},
	
	// Cloud Guard
	{Prefix: "cloud", ServiceName: "cloudguard"},
	
	// Recovery Service
	{Prefix: "recovery", ServiceName: "recovery"},
	
	// Lustre File System
	{Prefix: "lustre", ServiceName: "lustre"},
	
	// Desktops
	{Prefix: "desktops", ServiceName: "desktops"},
	
	// Artifacts
	{Prefix: "artifacts", ServiceName: "artifacts"},
}

func main() {
	// Read the Terraform schema
	schemaFile := "config/schema.json"
	data, err := ioutil.ReadFile(schemaFile)
	if err != nil {
		log.Fatalf("Error reading schema file: %v", err)
	}

	var schema TerraformSchema
	if err := json.Unmarshal(data, &schema); err != nil {
		log.Fatalf("Error parsing schema: %v", err)
	}

	// Get OCI provider schema
	ociProvider, ok := schema.ProviderSchemas["registry.terraform.io/oracle/oci"]
	if !ok {
		log.Fatal("OCI provider schema not found")
	}

	// Collect all resources and organize by service
	serviceResources := make(map[string][]string)
	
	for resourceName := range ociProvider.ResourceSchemas {
		if !strings.HasPrefix(resourceName, "oci_") {
			continue
		}
		
		service := detectService(resourceName)
		serviceResources[service] = append(serviceResources[service], resourceName)
	}
	
	// Generate GroupMap entries
	generateGroupMapFile(serviceResources)
	
	// Generate statistics
	printStatistics(serviceResources)
}

func detectService(resourceName string) string {
	parts := strings.Split(resourceName, "_")
	if len(parts) < 2 {
		return "unknown"
	}
	
	prefix := parts[1]
	
	// Find matching service mapping
	for _, mapping := range serviceMappings {
		if mapping.Prefix == prefix {
			// If there are rules, apply them
			if len(mapping.Rules) > 0 {
				for _, rule := range mapping.Rules {
					if strings.Contains(resourceName, rule.Pattern) {
						return rule.Service
					}
				}
				// If no rule matched, use a default
				return "core" // Default for unmatched core resources
			}
			// Direct mapping
			if mapping.ServiceName != "" {
				return mapping.ServiceName
			}
		}
	}
	
	// Default: use the prefix as service name
	return prefix
}

func generateGroupMapFile(serviceResources map[string][]string) {
	outputFile := "config/groups_generated.go"
	
	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer file.Close()
	
	// Track kind names to detect and resolve duplicates
	kindTracker := make(map[string][]string) // kind -> []resourceName
	resourceKinds := make(map[string]string)  // resourceName -> kind
	
	// First pass: generate all kinds and track duplicates
	for _, service := range serviceResources {
		for _, resource := range service {
			kind := generateKindName(resource)
			kindTracker[kind] = append(kindTracker[kind], resource)
			resourceKinds[resource] = kind
		}
	}
	
	// Second pass: resolve duplicates by making them unique
	for kind, resources := range kindTracker {
		if len(resources) > 1 {
			// Multiple resources generate the same kind - make them unique
			for i, resource := range resources {
				// Use service name as suffix for uniqueness
				service := detectService(resource)
				if i == 0 {
					// Keep first one as is, or add service if it's too generic
					if kind == "Migration" || kind == "Instance" || kind == "Configuration" || 
					   kind == "PrivateEndpoint" || kind == "Resource" || kind == "Cluster" ||
					   kind == "Policy" || kind == "Domain" || kind == "DbSystem" || 
					   kind == "Backup" || kind == "Repository" || kind == "Project" ||
					   kind == "Deployment" || kind == "Connection" || kind == "Certificate" ||
					   kind == "AutoScalingConfiguration" || kind == "Application" {
						resourceKinds[resource] = strings.Title(service) + kind
					}
				} else {
					// Add service prefix to subsequent duplicates
					resourceKinds[resource] = strings.Title(service) + kind
				}
			}
		}
	}
	
	// Write file header
	fmt.Fprintln(file, `/*
 * Copyright (c) 2025 Oracle and/or its affiliates
 * 
 * GENERATED FILE - DO NOT EDIT
 * Generated by group_map_generator.go
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

// GeneratedGroupMap contains auto-generated service mappings for all OCI Terraform resources  
// This supplements the manual GroupMap with automatic discovery of all 853+ resources
var GeneratedGroupMap = map[string]GroupKindCalculator{`)
	
	// Sort services for consistent output
	services := make([]string, 0, len(serviceResources))
	for service := range serviceResources {
		services = append(services, service)
	}
	sort.Strings(services)
	
	// Write entries for each service
	for _, service := range services {
		resources := serviceResources[service]
		sort.Strings(resources)
		
		fmt.Fprintf(file, "\n\t// %s Service (%d resources)\n", strings.Title(service), len(resources))
		
		for _, resource := range resources {
			kind := resourceKinds[resource]
			fmt.Fprintf(file, "\t\"%s\": func(name string) (string, string) {\n", resource)
			fmt.Fprintf(file, "\t\treturn \"%s\", \"%s\"\n", service, kind)
			fmt.Fprintf(file, "\t},\n")
		}
	}
	
	fmt.Fprintln(file, "}")
	
	// Add a merge function
	fmt.Fprintln(file, `
// MergeGroupMaps merges the manual GroupMap with the generated one
// Manual mappings take precedence over generated ones
func MergeGroupMaps() map[string]GroupKindCalculator {
	merged := make(map[string]GroupKindCalculator)
	
	// First add all generated mappings
	for k, v := range GeneratedGroupMap {
		merged[k] = v
	}
	
	// Override with manual mappings (they take precedence)
	for k, v := range GroupMap {
		merged[k] = v
	}
	
	return merged
}`)
	
	fmt.Printf("Generated GroupMap file: %s\n", outputFile)
}

func deduplicateResourceName(resourceName string) string {
	// Handle known problematic patterns
	patterns := map[string]string{
		"autonomous_database_autonomous_database": "autonomous_database",
		"external_container_database_external_container": "external_container_database",
		"external_non_container_database_external_non_container": "external_non_container_database", 
		"external_pluggable_database_external_pluggable": "external_pluggable_database",
		"external_mysql_database_external_mysql": "external_mysql_database",
		"pluggable_database_pluggable_database": "pluggable_database",
		"managements_management": "management",
		"monitorings_management": "monitoring_management",
		"cloud_database_managements_management": "cloud_database_management",
		"cloud_stack_monitorings_management": "cloud_stack_monitoring_management",
		// Additional patterns for remaining problematic resources
		"external_container_dbm_features_managements": "external_container_dbm_features_management",
		"external_mysql_databases_managements": "external_mysql_databases_management", 
		"external_non_container_dbm_features_managements": "external_non_container_dbm_features_management",
		"external_pluggable_dbm_features_managements": "external_pluggable_dbm_features_management",
		"pluggable_database_dbm_features_managements": "pluggable_database_dbm_features_management",
	}
	
	result := resourceName
	for pattern, replacement := range patterns {
		result = strings.ReplaceAll(result, pattern, replacement)
	}
	
	// Generic deduplication: remove consecutive duplicate segments
	parts := strings.Split(result, "_")
	deduplicated := make([]string, 0, len(parts))
	
	for i, part := range parts {
		// Skip if this part is the same as the previous one
		if i > 0 && part == parts[i-1] && len(part) > 2 {
			continue
		}
		deduplicated = append(deduplicated, part)
	}
	
	result = strings.Join(deduplicated, "_")
	
	// Additional aggressive shortening for database management resources
	if strings.Contains(result, "database_management") {
		// Shorten common long patterns in database management
		result = strings.ReplaceAll(result, "_database_management_", "_dbmgmt_")
		result = strings.ReplaceAll(result, "_features_management", "_features_mgmt")
		result = strings.ReplaceAll(result, "_managements", "_mgmt")
		result = strings.ReplaceAll(result, "_management", "_mgmt")
		
		// Extra aggressive shortening for the longest resources
		result = strings.ReplaceAll(result, "_external_non_container_", "_ext_noncontainer_")
		result = strings.ReplaceAll(result, "_external_container_", "_ext_container_")
		result = strings.ReplaceAll(result, "_external_pluggable_", "_ext_pluggable_")
		result = strings.ReplaceAll(result, "_dbm_features_", "_dbm_feat_")
	}
	
	return result
}

func generateKindName(resourceName string) string {
	// First deduplicate the resource name
	deduplicatedName := deduplicateResourceName(resourceName)
	
	// Remove oci_ prefix and service prefix
	parts := strings.Split(deduplicatedName, "_")
	if len(parts) < 3 {
		return "Resource"
	}
	
	// Join remaining parts and convert to CamelCase
	kindParts := parts[2:]
	result := ""
	for _, part := range kindParts {
		result += strings.Title(part)
	}
	
	if result == "" {
		result = "Resource"
	}
	
	// Handle specific cases where shortening created duplicates
	// Use the original resource name to create unique suffixes
	switch {
	case resourceName == "oci_database_migration" && result == "Migration":
		result = "DatabaseMigration"
	case resourceName == "oci_database_migration_migration" && result == "Migration":
		result = "MigrationMigration"
	}
	
	return result
}

func printStatistics(serviceResources map[string][]string) {
	fmt.Println("\n=== Service Statistics ===")
	
	// Sort services by resource count
	type serviceStat struct {
		Name  string
		Count int
	}
	
	var stats []serviceStat
	totalResources := 0
	
	for service, resources := range serviceResources {
		stats = append(stats, serviceStat{Name: service, Count: len(resources)})
		totalResources += len(resources)
	}
	
	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Count > stats[j].Count
	})
	
	fmt.Printf("\nTotal Resources: %d\n", totalResources)
	fmt.Printf("Total Services: %d\n\n", len(serviceResources))
	
	fmt.Println("Top 20 Services by Resource Count:")
	for i, stat := range stats {
		if i >= 20 {
			break
		}
		fmt.Printf("%3d. %-30s: %3d resources\n", i+1, stat.Name, stat.Count)
	}
	
	// Show newly discovered services
	fmt.Println("\n=== Newly Discovered Services ===")
	existingServices := map[string]bool{
		"compute": true, "networking": true, "blockstorage": true,
		"networkconnectivity": true, "containerengine": true, "identity": true,
		"objectstorage": true, "kms": true, "artifacts": true, "ons": true,
		"networkloadbalancer": true, "dns": true, "monitoring": true,
		"healthchecks": true, "functions": true, "networkfirewall": true,
		"logging": true, "loadbalancer": true, "certificatesmanagement": true,
		"filestorage": true, "events": true, "streaming": true, "vault": true,
	}
	
	var newServices []string
	for service := range serviceResources {
		if !existingServices[service] {
			newServices = append(newServices, service)
		}
	}
	sort.Strings(newServices)
	
	for _, service := range newServices {
		fmt.Printf("- %s (%d resources)\n", service, len(serviceResources[service]))
	}
}
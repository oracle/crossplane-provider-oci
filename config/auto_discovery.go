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

package config

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/crossplane/upjet/pkg/config"
)

// ProblematicResources returns a list of regex patterns for resources that should be
// skipped during generation due to known issues or incompatibilities.
// These resources can be added to support later after resolving their specific issues.
func ProblematicResources() []string {
	return []string{
		// Skip data sources (not needed for managed resources)
		`.*_data_.*`,

		// Skip test resources (internal testing only)
		`.*_test.*`,

		// Skip deprecated resources
		`.*_deprecated.*`,

		// Known problematic resources that need special handling
		`oci_network_firewall_network_firewall_policy_service_list$`,     // Name collision: generates duplicate types
		`oci_network_firewall_network_firewall_policy_url_list$`,         // Similar potential naming conflict
		`oci_network_firewall_network_firewall_policy_application_list$`, // Similar potential naming conflict

		// Add more specific resources here as we discover generation issues
	}
}

// AutoExternalNameConfiguration provides automatic external name configuration
// for resources that don't have explicit configuration in ExternalNameConfigs.
// This ensures all discovered resources can be properly managed.
func AutoExternalNameConfiguration() config.ResourceOption {
	return func(r *config.Resource) {
		// Only apply if not already configured
		if r.ExternalName.GetExternalNameFn == nil {
			// Check if this resource has explicit configuration
			if _, ok := ExternalNameConfigs[r.Name]; !ok {
				// Use IdentifierFromProvider as default for all OCI resources
				// This is the most common pattern in OCI
				r.ExternalName = config.IdentifierFromProvider
				r.Version = "v1alpha1"
			}
		}
	}
}

// ServiceGroupDetector automatically determines the service group for a resource
// based on its name pattern. This is used when GroupMap doesn't have an explicit entry.
func ServiceGroupDetector(resourceName string) (group string, kind string) {
	// Extract the service prefix (e.g., "oci_database_*" -> "database")
	parts := strings.Split(resourceName, "_")
	if len(parts) < 2 {
		return "core", generateKindName(resourceName)
	}

	servicePrefix := parts[1]

	// Special handling for core resources that should be split
	if servicePrefix == "core" {
		return detectCoreServiceGroup(resourceName)
	}

	// Special handling for multi-word services
	switch servicePrefix {
	case "network":
		if len(parts) > 2 {
			switch parts[2] {
			case "firewall":
				return "networkfirewall", generateKindName(resourceName)
			case "load":
				return "networkloadbalancer", generateKindName(resourceName)
			default:
				return "networking", generateKindName(resourceName)
			}
		}
		return "networking", generateKindName(resourceName)

	case "load":
		return "loadbalancer", generateKindName(resourceName)

	case "file":
		return "filestorage", generateKindName(resourceName)

	case "health":
		return "healthchecks", generateKindName(resourceName)

	case "certificates":
		return "certificatesmanagement", generateKindName(resourceName)

	default:
		// Use the service prefix as the group
		return servicePrefix, generateKindName(resourceName)
	}
}

// detectCoreServiceGroup intelligently splits oci_core_* resources into logical services
func detectCoreServiceGroup(resourceName string) (group string, kind string) {
	// Compute resources
	if contains(resourceName, []string{"instance", "image", "dedicated_vm", "console", "shape", "app_catalog", "cluster_network", "compute_"}) {
		return "compute", generateKindName(resourceName)
	}

	// Networking resources
	if contains(resourceName, []string{"vcn", "subnet", "vnic", "dhcp", "vlan", "gateway", "security", "route", "ip", "peering"}) {
		return "networking", generateKindName(resourceName)
	}

	// Block storage resources
	if contains(resourceName, []string{"volume", "boot_volume"}) {
		return "blockstorage", generateKindName(resourceName)
	}

	// Network connectivity resources (DRG, IPSec, etc.)
	if contains(resourceName, []string{"drg", "cross_connect", "virtual_circuit", "cpe", "ipsec"}) {
		return "networkconnectivity", generateKindName(resourceName)
	}

	// Monitoring resources
	if contains(resourceName, []string{"capture_filter", "vtap"}) {
		return "monitoring", generateKindName(resourceName)
	}

	// Default to core for unmatched patterns
	return "core", generateKindName(resourceName)
}

// generateKindName converts a resource name to a proper Kind name
func generateKindName(resourceName string) string {
	// Remove oci_ prefix and service prefix
	parts := strings.Split(resourceName, "_")
	if len(parts) < 3 {
		return "Resource"
	}

	// Join remaining parts and convert to CamelCase
	kindParts := parts[2:]
	result := ""
	titleCaser := cases.Title(language.Und)
	for _, part := range kindParts {
		result += titleCaser.String(part)
	}

	if result == "" {
		result = "Resource"
	}

	return result
}

// contains checks if the resource name contains any of the given patterns
func contains(resourceName string, patterns []string) bool {
	for _, pattern := range patterns {
		if strings.Contains(resourceName, pattern) {
			return true
		}
	}
	return false
}

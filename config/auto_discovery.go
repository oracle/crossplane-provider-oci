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
		// OCI resources do not have a data section like AWS/GCP, so we do not need to specifically skip them.
		// Explicity using `*_data_*` skips oci_identity_data_plane_* resources hence commenting it out.
		// Skip data sources (not needed for managed resources)
		// `.*_data_.*`,

		// Skip test resources (internal testing only)
		`.*_test.*`,

		// Skip deprecated resources
		`.*_deprecated.*`,

		// Known problematic resources that need special handling
		`oci_network_firewall_network_firewall_policy_service_list$`,     // Name collision: generates duplicate types
		`oci_network_firewall_network_firewall_policy_url_list$`,         // Similar potential naming conflict
		`oci_network_firewall_network_firewall_policy_application_list$`, // Similar potential naming conflict
		`oci_load_balancer_backendset$`,                                  // Alias for oci_load_balancer_backend_set

		// Add more specific resources here as we discover generation issues
	}
}

// AutoExternalNameConfiguration provides automatic external name configuration
// for resources that don't have explicit configuration in ExternalNameConfigs.
// This ensures all discovered resources can be properly managed.
func AutoExternalNameConfiguration() config.ResourceOption {
	return func(r *config.Resource) {
		// Only apply if not already configured
		if r.ExternalName.DisableNameInitializer == false {
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

	// Fallback case: Resources with len(parts)<2 are grouped into core service
	if len(parts) < 2 {
		group = "core"
		return group, generateKindName(resourceName, group)
	}

	servicePrefix := parts[1]

	// Special handling for core resources that should be split
	if servicePrefix == "core" {
		group = detectCoreServiceGroup(resourceName)
		return group, generateKindName(resourceName, group)
	}

	// Special handling for multi-word services
	switch servicePrefix {
	case "network":
		if len(parts) > 2 {
			switch parts[2] {
			case "firewall":
				group = "networkfirewall"
			case "load":
				group = "networkloadbalancer"
			default:
				group = "networking"
			}
		} else {
			group = "networking"
		}

	case "load":
		group = "loadbalancer"

	case "file":
		group = "filestorage"

	case "health":
		group = "healthchecks"

	case "certificates":
		group = "certificatesmanagement"

	default:
		group = servicePrefix
	}

	return group, generateKindName(resourceName, group)
}

// detectCoreServiceGroup intelligently splits oci_core_* resources into logical services
func detectCoreServiceGroup(resourceName string) (group string) {
	// Compute resources
	if contains(resourceName, []string{"instance", "image", "dedicated_vm", "console", "shape", "app_catalog", "cluster_network", "compute_"}) {
		return "compute"
	}

	// Networking resources
	if contains(resourceName, []string{"vcn", "subnet", "vnic", "dhcp", "vlan", "gateway", "security", "route", "ip", "peering"}) {
		return "networking"
	}

	// Block storage resources
	if contains(resourceName, []string{"volume", "boot_volume"}) {
		return "blockstorage"
	}

	// Network connectivity resources (DRG, IPSec, etc.)
	if contains(resourceName, []string{"drg", "cross_connect", "virtual_circuit", "cpe", "ipsec"}) {
		return "networkconnectivity"
	}

	// Monitoring resources
	if contains(resourceName, []string{"capture_filter", "vtap"}) {
		return "monitoring"
	}

	// Default to core for unmatched patterns
	return "core"
}

// generateKindName converts a resource name to a Kind name by stripping the
// OCI prefix and all tokens that make up the logical service group.
//
//   - resourceName: full Terraform resource name, e.g. "oci_os_management_update_schedule".
//   - group: the service group returned by ServiceGroupDetector, e.g. "osmanagement".
//
// If the resource name does not have at least three underscore-separated parts,
// the function returns the original resourceName unchanged.
func generateKindName(resourceName, group string) string {
	parts := strings.Split(resourceName, "_")
	if len(parts) < 3 {
		// Not in expected oci_<service>_<resource> form; return as-is.
		return resourceName
	}

	tokens := parts[1:] // drop the first segment (usually "oci")
	if len(tokens) == 0 {
		return resourceName
	}

	// Find the shortest prefix of tokens whose concatenation matches the group
	// string (case-insensitive). That prefix represents the service name
	// portion to drop from the Kind.
	groupLower := strings.ToLower(group)
	prefixEnd := 0
	if groupLower != "" {
		for i := 1; i <= len(tokens); i++ {
			candidate := strings.ToLower(strings.Join(tokens[:i], ""))
			if candidate == groupLower {
				prefixEnd = i
				break
			}
		}
	}

	// If we couldn't match the group to a prefix, fall back to dropping
	// only the first token (service prefix) as in the original behavior.
	start := 1
	if prefixEnd > 0 {
		start = prefixEnd
	}
	if start >= len(tokens) {
		// No resource tokens left; return original resourceName defensively.
		return resourceName
	}

	kindTokens := tokens[start:]

	titleCaser := cases.Title(language.Und)
	var b strings.Builder
	for _, t := range kindTokens {
		if t == "" {
			continue
		}
		b.WriteString(titleCaser.String(t))
	}

	if b.Len() == 0 {
		return resourceName
	}

	return b.String()
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

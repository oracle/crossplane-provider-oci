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
		`oci_load_balancer$`,                                             // Alias for oci_load_balancer_load_balancer

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

	// Multi-word / colliding prefixes use parts[2] to disambiguate, defaulting to
	// the first known service for that prefix. Single-word services rely on the
	// default fallback (group = servicePrefix) and are not listed here.
	switch servicePrefix {
	case "ai":
		// Ai Data Platform, Ai Document, Ai Language, Ai Vision
		if len(parts) > 2 {
			switch parts[2] {
			case "data", "dataplatform", "data_platform":
				group = "aidataplatform"
			case "document":
				group = "aidocument"
			case "language":
				group = "ailanguage"
			case "vision":
				group = "aivision"
			default:
				// default to first known AI service
				group = "aidataplatform"
			}
		} else {
			group = "aidataplatform"
		}

	case "announcements":
		// "Announcements Service": "announcements_service" -> "announcementsservice"
		group = "announcementsservice"

	case "api":
		// "API Platform": "api_platform" -> "apiplatform"
		group = "apiplatform"

	case "apm":
		// Apm, Apm Config, Apm Synthetics, Apm Traces
		if len(parts) > 2 {
			switch parts[2] {
			case "config":
				group = "apmconfig"
			case "synthetics":
				group = "apmsynthetics"
			case "traces":
				group = "apmtraces"
			default:
				group = "apm"
			}
		} else {
			group = "apm"
		}

	case "capacity":
		// "Capacity Management": "capacity_management" -> "capacitymanagement"
		group = "capacitymanagement"

	case "certificates":
		// "Certificates Management": "certificates_management" -> "certificatesmanagement"
		group = "certificatesmanagement"

	case "cloud":
		// Cloud Bridge, Cloud Guard, Cloud Migrations
		if len(parts) > 2 {
			switch parts[2] {
			case "bridge":
				group = "cloudbridge"
			case "guard":
				group = "cloudguard"
			case "migrations":
				group = "cloudmigrations"
			default:
				group = "cloudbridge" // default to first service
			}
		} else {
			group = "cloudbridge"
		}

	case "cluster":
		// "Cluster Placement Groups": "cluster_placement_groups" -> "clusterplacementgroups"
		group = "clusterplacementgroups"

	case "compute":
		// "Compute Cloud At Customer": "compute_cloud_at_customer" -> "computecloudatcustomer"
		group = "computecloudatcustomer"

	case "container":
		// Container Engine, Container Instances
		if len(parts) > 2 {
			switch parts[2] {
			case "engine":
				group = "containerengine"
			case "instances":
				group = "containerinstances"
			default:
				group = "containerengine"
			}
		} else {
			group = "containerengine"
		}

	case "data":
		// Data Labeling Service: "data_labeling_service" -> "datalabelingservice"
		group = "datalabelingservice"

	case "delegate":
		// Delegate Access Control: "delegate_access_control" -> "delegateaccesscontrol"
		group = "delegateaccesscontrol"

	case "demand":
		// Demand Signal: "demand_signal" -> "demandsignal"
		group = "demandsignal"

	case "disaster":
		// Disaster Recovery: "disaster_recovery" -> "disasterrecovery"
		group = "disasterrecovery"

	case "file":
		// "File Storage": "file_storage" -> "filestorage"
		group = "filestorage"

	case "fleet":
		// Fleet Apps Management, Fleet Software Update
		if len(parts) > 2 {
			switch parts[2] {
			case "apps":
				group = "fleetappsmanagement"
			case "software":
				group = "fleetsoftwareupdate"
			default:
				group = "fleetappsmanagement"
			}
		} else {
			group = "fleetappsmanagement"
		}

	case "lustre":
		// "Lustre File Storage": "lustre_file_storage" -> "lustrefilestorage"
		group = "lustrefilestorage"

	case "managed":
		// "Managed Kafka": "managed_kafka" -> "managedkafka"
		group = "managedkafka"

	case "media":
		// "Media Services": "media_services" -> "mediaservices"
		group = "mediaservices"

	case "metering":
		// "Metering Computation": "metering_computation" -> "meteringcomputation"
		group = "meteringcomputation"

	case "generic":
		// "Generic Artifacts Content": "generic_artifacts_content" -> "genericartifactscontent"
		group = "genericartifactscontent"

	case "globally":
		// "Globally Distributed Database": "globally_distributed_database" -> "globallydistributeddatabase"
		group = "globallydistributeddatabase"

	case "golden":
		// "Golden Gate": "golden_gate" -> "goldengate"
		group = "goldengate"

	case "fusion":
		// "Fusion Apps": "fusion_apps" -> "fusionapps"
		group = "fusionapps"

	case "generative":
		// Generative AI, Generative Ai Agent
		group = "generativeai"

	case "health":
		// "Health Checks": "health_checks" -> "healthchecks"
		group = "healthchecks"

	case "identity":
		// Identity, Identity Data Plane, Identity Domains
		if len(parts) > 2 {
			switch parts[2] {
			case "data":
				group = "identitydataplane"
			case "domains":
				group = "identitydomains"
			default:
				group = "identity"
			}
		} else {
			group = "identity"
		}

	case "jms":
		// Jms, Jms Java Downloads, Jms Utils
		if len(parts) > 2 {
			switch parts[2] {
			case "java":
				group = "jmsjavadownloads"
			case "utils":
				group = "jmsutils"
			default:
				group = "jms"
			}
		} else {
			group = "jms"
		}

	case "license":
		// "License Manager": "license_manager" -> "licensemanager"
		group = "licensemanager"

	case "load":
		// "Load Balancer": "load_balancer" -> "loadbalancer"
		group = "loadbalancer"

	case "log":
		// "Log Analytics": "log_analytics" -> "loganalytics"
		group = "loganalytics"

	case "management":
		// Management Agent, Management Dashboard
		if len(parts) > 2 {
			switch parts[2] {
			case "agent":
				group = "managementagent"
			case "dashboard":
				group = "managementdashboard"
			default:
				group = "managementagent"
			}
		} else {
			group = "managementagent"
		}

	case "network":
		// Network Firewall, Network Load Balancer; default networking
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

	case "osp":
		// "Osp Gateway": "osp_gateway" -> "ospgateway"
		group = "ospgateway"

	case "operator":
		// "Operator Access Control": "operator_access_control" -> "operatoraccesscontrol"
		group = "operatoraccesscontrol"

	case "os":
		// "Os Management Hub": "os_management_hub" -> "osmanagementhub"
		group = "osmanagementhub"

	case "resource":
		// Resource Analytics, Resource Manager, Resource Scheduler
		if len(parts) > 2 {
			switch parts[2] {
			case "analytics":
				group = "resourceanalytics"
			case "manager":
				group = "resourcemanager"
			case "scheduler":
				group = "resourcescheduler"
			default:
				group = "resourcemanager"
			}
		} else {
			group = "resourcemanager"
		}

	case "security":
		// "Security Attribute": "security_attribute" -> "securityattribute"
		group = "securityattribute"

	case "service":
		// "Service Catalog": "service_catalog" -> "servicecatalog"
		group = "servicecatalog"

	case "stack":
		// "Stack Monitoring": "stack_monitoring" -> "stackmonitoring"
		group = "stackmonitoring"

	case "usage":
		// "Usage Proxy": "usage_proxy" -> "usageproxy"
		group = "usageproxy"

	case "vbs":
		// "Vbs Inst": "vbs_inst" -> "vbsinst"
		group = "vbsinst"

	case "visual":
		// "Visual Builder": "visual_builder" -> "visualbuilder"
		group = "visualbuilder"

	case "vn":
		// "Vn Monitoring": "vn_monitoring" -> "vnmonitoring"
		group = "vnmonitoring"

	case "vulnerability":
		// "Vulnerability Scanning": "vulnerability_scanning" -> "vulnerabilityscanning"
		group = "vulnerabilityscanning"

	default:
		// Fallback: use the raw service prefix
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

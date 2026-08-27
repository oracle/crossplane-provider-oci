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

	"github.com/crossplane/upjet/v2/pkg/config"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	maxKubernetesNameLength    = 63
	maxGeneratedKindNameLength = maxKubernetesNameLength - len("List")
)


// GroupKindCalculator returns the correct service group and kind for a resource
type GroupKindCalculator func(resource string) (group string, kind string)

// GroupMap remaps OCI Terraform resources to service-based groups
// This solves the API group vs service mismatch (e.g., oci_core_* → appropriate services)
var GroupMap = map[string]GroupKindCalculator{
	// Compute Service - Instance and compute-related resources from oci_core_*
	"oci_core_instance": func(name string) (string, string) {
		return "compute", "Instance"
	},
	"oci_core_image": func(name string) (string, string) {
		return "compute", "Image"
	},
	"oci_core_dedicated_vm_host": func(name string) (string, string) {
		return "compute", "DedicatedVmHost"
	},
	"oci_core_console_history": func(name string) (string, string) {
		return "compute", "ConsoleHistory"
	},
	"oci_core_instance_configuration": func(name string) (string, string) {
		return "compute", "InstanceConfiguration"
	},
	"oci_core_instance_console_connection": func(name string) (string, string) {
		return "compute", "InstanceConsoleConnection"
	},
	"oci_core_instance_pool": func(name string) (string, string) {
		return "compute", "InstancePool"
	},
	"oci_core_instance_pool_instance": func(name string) (string, string) {
		return "compute", "InstancePoolInstance"
	},
	"oci_core_cluster_network": func(name string) (string, string) {
		return "compute", "ClusterNetwork"
	},
	"oci_core_compute_capacity_reservation": func(name string) (string, string) {
		return "compute", "ComputeCapacityReservation"
	},
	"oci_core_compute_cluster": func(name string) (string, string) {
		return "compute", "ComputeCluster"
	},
	"oci_core_compute_image_capability_schema": func(name string) (string, string) {
		return "compute", "ComputeImageCapabilitySchema"
	},
	"oci_core_shape_management": func(name string) (string, string) {
		return "compute", "ShapeManagement"
	},
	"oci_core_app_catalog_listing_resource_version_agreement": func(name string) (string, string) {
		return "compute", "AppCatalogListingResourceVersionAgreement"
	},
	"oci_core_app_catalog_subscription": func(name string) (string, string) {
		return "compute", "AppCatalogSubscription"
	},

	// Networking Service - VCNs, subnets, gateways, and basic networking from oci_core_*
	"oci_core_vcn": func(name string) (string, string) {
		return "networking", "Vcn"
	},
	"oci_core_subnet": func(name string) (string, string) {
		return "networking", "Subnet"
	},
	"oci_core_vnic_attachment": func(name string) (string, string) {
		return "networking", "VnicAttachment"
	},
	"oci_core_dhcp_options": func(name string) (string, string) {
		return "networking", "DhcpOptions"
	},
	"oci_core_vlan": func(name string) (string, string) {
		return "networking", "Vlan"
	},
	"oci_core_internet_gateway": func(name string) (string, string) {
		return "networking", "InternetGateway"
	},
	"oci_core_nat_gateway": func(name string) (string, string) {
		return "networking", "NatGateway"
	},
	"oci_core_service_gateway": func(name string) (string, string) {
		return "networking", "ServiceGateway"
	},
	"oci_core_network_security_group": func(name string) (string, string) {
		return "networking", "NetworkSecurityGroup"
	},
	"oci_core_network_security_group_security_rule": func(name string) (string, string) {
		return "networking", "NetworkSecurityGroupSecurityRule"
	},
	"oci_core_route_table": func(name string) (string, string) {
		return "networking", "RouteTable"
	},
	"oci_core_security_list": func(name string) (string, string) {
		return "networking", "SecurityList"
	},
	"oci_core_private_ip": func(name string) (string, string) {
		return "networking", "PrivateIp"
	},
	"oci_core_public_ip": func(name string) (string, string) {
		return "networking", "PublicIp"
	},
	"oci_core_public_ip_pool": func(name string) (string, string) {
		return "networking", "PublicIpPool"
	},
	"oci_core_public_ip_pool_capacity": func(name string) (string, string) {
		return "networking", "PublicIpPoolCapacity"
	},
	"oci_core_byoip_range": func(name string) (string, string) {
		return "networking", "ByoipRange"
	},
	"oci_core_local_peering_gateway": func(name string) (string, string) {
		return "networking", "LocalPeeringGateway"
	},
	"oci_core_remote_peering_connection": func(name string) (string, string) {
		return "networking", "RemotePeeringConnection"
	},
	"oci_core_route_table_attachment": func(name string) (string, string) {
		return "networking", "RouteTableAttachment"
	},
	"oci_core_ipv6": func(name string) (string, string) {
		return "networking", "Ipv6"
	},

	// Block Storage Service - Volumes and storage from oci_core_*
	"oci_core_volume": func(name string) (string, string) {
		return "blockstorage", "Volume"
	},
	"oci_core_volume_backup": func(name string) (string, string) {
		return "blockstorage", "VolumeBackup"
	},
	"oci_core_volume_attachment": func(name string) (string, string) {
		return "blockstorage", "VolumeAttachment"
	},
	"oci_core_volume_backup_policy": func(name string) (string, string) {
		return "blockstorage", "VolumeBackupPolicy"
	},
	"oci_core_volume_backup_policy_assignment": func(name string) (string, string) {
		return "blockstorage", "VolumeBackupPolicyAssignment"
	},
	"oci_core_volume_group": func(name string) (string, string) {
		return "blockstorage", "VolumeGroup"
	},
	"oci_core_volume_group_backup": func(name string) (string, string) {
		return "blockstorage", "VolumeGroupBackup"
	},
	"oci_core_boot_volume": func(name string) (string, string) {
		return "blockstorage", "BootVolume"
	},
	"oci_core_boot_volume_backup": func(name string) (string, string) {
		return "blockstorage", "BootVolumeBackup"
	},

	// Network Connectivity Service - Advanced networking, DRG, cross-connects from oci_core_*
	"oci_core_drg": func(name string) (string, string) {
		return "networkconnectivity", "Drg"
	},
	"oci_core_drg_attachment": func(name string) (string, string) {
		return "networkconnectivity", "DrgAttachment"
	},
	"oci_core_drg_attachment_management": func(name string) (string, string) {
		return "networkconnectivity", "DrgAttachmentManagement"
	},
	"oci_core_drg_attachments_list": func(name string) (string, string) {
		return "networkconnectivity", "DrgAttachmentsList"
	},
	"oci_core_drg_route_distribution": func(name string) (string, string) {
		return "networkconnectivity", "DrgRouteDistribution"
	},
	"oci_core_drg_route_distribution_statement": func(name string) (string, string) {
		return "networkconnectivity", "DrgRouteDistributionStatement"
	},
	"oci_core_drg_route_table": func(name string) (string, string) {
		return "networkconnectivity", "DrgRouteTable"
	},
	"oci_core_default_drg_route_table": func(name string) (string, string) {
		return "networkconnectivity", "DefaultDrgRouteTable"
	},
	"oci_core_drg_route_table_route_rule": func(name string) (string, string) {
		return "networkconnectivity", "DrgRouteTableRouteRule"
	},
	"oci_core_cross_connect": func(name string) (string, string) {
		return "networkconnectivity", "CrossConnect"
	},
	"oci_core_cross_connect_group": func(name string) (string, string) {
		return "networkconnectivity", "CrossConnectGroup"
	},
	"oci_core_virtual_circuit": func(name string) (string, string) {
		return "networkconnectivity", "VirtualCircuit"
	},
	"oci_core_cpe": func(name string) (string, string) {
		return "networkconnectivity", "Cpe"
	},
	"oci_core_ipsec": func(name string) (string, string) {
		return "networkconnectivity", "Ipsec"
	},
	"oci_core_ipsec_connection_tunnel_management": func(name string) (string, string) {
		return "networkconnectivity", "IpsecConnectionTunnelManagement"
	},

	// Monitoring & Observability - Network monitoring from oci_core_*
	"oci_core_capture_filter": func(name string) (string, string) {
		return "monitoring", "CaptureFilter"
	},
	"oci_core_vtap": func(name string) (string, string) {
		return "monitoring", "Vtap"
	},

	// Identity Service - Keep existing identity prefix
	"oci_identity_compartment": func(name string) (string, string) {
		return "identity", "Compartment"
	},
	"oci_identity_group": func(name string) (string, string) {
		return "identity", "Group"
	},
	"oci_identity_policy": func(name string) (string, string) {
		return "identity", "Policy"
	},
	"oci_identity_tag": func(name string) (string, string) {
		return "identity", "Tag"
	},
	"oci_identity_tag_namespace": func(name string) (string, string) {
		return "identity", "TagNamespace"
	},
	"oci_identity_authentication_policy": func(name string) (string, string) {
		return "identity", "AuthenticationPolicy"
	},
	"oci_identity_identity_provider": func(name string) (string, string) {
		return "identity", "IdentityProvider"
	},

	// Container Engine Service
	"oci_containerengine_cluster": func(name string) (string, string) {
		return "containerengine", "Cluster"
	},
	"oci_containerengine_node_pool": func(name string) (string, string) {
		return "containerengine", "NodePool"
	},

	// Object Storage Service
	"oci_objectstorage_bucket": func(name string) (string, string) {
		return "objectstorage", "Bucket"
	},
	"oci_objectstorage_object": func(name string) (string, string) {
		return "objectstorage", "Object"
	},
	"oci_objectstorage_object_lifecycle_policy": func(name string) (string, string) {
		return "objectstorage", "ObjectLifecyclePolicy"
	},

	// Key Management Service
	"oci_kms_key": func(name string) (string, string) {
		return "kms", "Key"
	},
	"oci_kms_key_version": func(name string) (string, string) {
		return "kms", "KeyVersion"
	},
	"oci_kms_vault": func(name string) (string, string) {
		return "kms", "Vault"
	},

	// Artifacts Service
	"oci_artifacts_container_repository": func(name string) (string, string) {
		return "artifacts", "ContainerRepository"
	},
	"oci_artifacts_repository": func(name string) (string, string) {
		return "artifacts", "Repository"
	},
	"oci_artifacts_container_configuration": func(name string) (string, string) {
		return "artifacts", "ContainerConfiguration"
	},
	"oci_artifacts_generic_artifact": func(name string) (string, string) {
		return "artifacts", "GenericArtifact"
	},

	// Notification Service
	"oci_ons_notification_topic": func(name string) (string, string) {
		return "ons", "NotificationTopic"
	},
	"oci_ons_subscription": func(name string) (string, string) {
		return "ons", "Subscription"
	},

	// Network Load Balancer Service
	"oci_network_load_balancer_network_load_balancer": func(name string) (string, string) {
		return "networkloadbalancer", "NetworkLoadBalancer"
	},
	"oci_network_load_balancer_backend_set": func(name string) (string, string) {
		return "networkloadbalancer", "BackendSet"
	},
	"oci_network_load_balancer_backend": func(name string) (string, string) {
		return "networkloadbalancer", "Backend"
	},
	"oci_network_load_balancer_network_load_balancers_backend_sets_unified": func(name string) (string, string) {
		return "networkloadbalancer", "NetworkLoadBalancersBackendSetsUnified"
	},
	"oci_network_load_balancer_listener": func(name string) (string, string) {
		return "networkloadbalancer", "Listener"
	},

	// DNS Service
	"oci_dns_record": func(name string) (string, string) {
		return "dns", "Record"
	},
	"oci_dns_resolver": func(name string) (string, string) {
		return "dns", "Resolver"
	},
	"oci_dns_rrset": func(name string) (string, string) {
		return "dns", "Rrset"
	},
	"oci_dns_resolver_endpoint": func(name string) (string, string) {
		return "dns", "ResolverEndpoint"
	},
	"oci_dns_steering_policy": func(name string) (string, string) {
		return "dns", "SteeringPolicy"
	},
	"oci_dns_steering_policy_attachment": func(name string) (string, string) {
		return "dns", "SteeringPolicyAttachment"
	},
	"oci_dns_tsig_key": func(name string) (string, string) {
		return "dns", "TsigKey"
	},
	"oci_dns_zone": func(name string) (string, string) {
		return "dns", "Zone"
	},
	"oci_dns_view": func(name string) (string, string) {
		return "dns", "View"
	},

	// Monitoring Service
	"oci_monitoring_alarm_history_collection": func(name string) (string, string) {
		return "monitoring", "AlarmHistoryCollection"
	},
	"oci_monitoring_alarm": func(name string) (string, string) {
		return "monitoring", "Alarm"
	},

	// Health Checks Service
	"oci_health_checks_http_monitor": func(name string) (string, string) {
		return "healthchecks", "HttpMonitor"
	},
	"oci_health_checks_ping_monitor": func(name string) (string, string) {
		return "healthchecks", "PingMonitor"
	},

	// Functions Service
	"oci_functions_application": func(name string) (string, string) {
		return "functions", "Application"
	},
	"oci_functions_function": func(name string) (string, string) {
		return "functions", "Function"
	},
	"oci_functions_invoke_function": func(name string) (string, string) {
		return "functions", "InvokeFunction"
	},

	// Network Firewall Service
	"oci_network_firewall_network_firewall_policy": func(name string) (string, string) {
		return "networkfirewall", "NetworkFirewallPolicy"
	},
	"oci_network_firewall_network_firewall": func(name string) (string, string) {
		return "networkfirewall", "NetworkFirewall"
	},

	// Load Balancer Service
	"oci_load_balancer_load_balancer": func(name string) (string, string) {
		return "loadbalancer", "LoadBalancer"
	},
	"oci_load_balancer_hostname": func(name string) (string, string) {
		return "loadbalancer", "Hostname"
	},
	"oci_load_balancer_backend_set": func(name string) (string, string) {
		return "loadbalancer", "BackendSet"
	},
	"oci_load_balancer_rule_set": func(name string) (string, string) {
		return "loadbalancer", "RuleSet"
	},
	"oci_load_balancer_backend": func(name string) (string, string) {
		return "loadbalancer", "Backend"
	},
	"oci_load_balancer_certificate": func(name string) (string, string) {
		return "loadbalancer", "Certificate"
	},
	"oci_load_balancer_load_balancer_routing_policy": func(name string) (string, string) {
		return "loadbalancer", "LoadBalancerRoutingPolicy"
	},
	"oci_load_balancer_path_route_set": func(name string) (string, string) {
		return "loadbalancer", "PathRouteSet"
	},
	"oci_load_balancer_listener": func(name string) (string, string) {
		return "loadbalancer", "Listener"
	},
	"oci_load_balancer_ssl_cipher_suite": func(name string) (string, string) {
		return "loadbalancer", "SslCipherSuite"
	},

	// Certificates Management Service
	"oci_certificates_management_certificate_authority": func(name string) (string, string) {
		return "certificatesmanagement", "CertificateAuthority"
	},
	"oci_certificates_management_certificate": func(name string) (string, string) {
		return "certificatesmanagement", "Certificate"
	},
	"oci_certificates_management_ca_bundle": func(name string) (string, string) {
		return "certificatesmanagement", "CaBundle"
	},

	// MySQL Service
	"oci_mysql_backup":            func(name string) (string, string) { return "mysql", "MysqlBackup" },
	"oci_mysql_channel":           func(name string) (string, string) { return "mysql", "MysqlChannel" },
	"oci_mysql_configuration":     func(name string) (string, string) { return "mysql", "MysqlConfiguration" },
	"oci_mysql_db_system":         func(name string) (string, string) { return "mysql", "MysqlDbSystem" },
	"oci_mysql_heat_wave_cluster": func(name string) (string, string) { return "mysql", "MysqlHeatWaveCluster" },
	"oci_mysql_replica":           func(name string) (string, string) { return "mysql", "MysqlReplica" },

	// PostgreSQL Service
	"oci_psql_backup":        func(name string) (string, string) { return "psql", "PsqlBackup" },
	"oci_psql_configuration": func(name string) (string, string) { return "psql", "PsqlConfiguration" },
	"oci_psql_db_system":     func(name string) (string, string) { return "psql", "PsqlDbSystem" },
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
		// Data Labeling Service, Data Safe
		if len(parts) > 2 {
			switch parts[2] {
			case "labeling":
				group = "datalabelingservice"
			case "safe":
				group = "datasafe"
			default:
				group = "datasafe"
			}
		} else {
			group = "datasafe"
		}

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

	case "globally", "distributed":
		// OCI provider docs now publish distributed database resources under
		// "distributed_database", while older generated assets still use the
		// legacy "globally_distributed_database" prefix.
		group = "distributeddatabase"

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

	kindTokens := make([]string, 0, len(tokens[start:]))
	for _, token := range tokens[start:] {
		if token == "" {
			continue
		}
		kindTokens = append(kindTokens, token)
	}
	if len(kindTokens) == 0 {
		return resourceName
	}

	titleCaser := cases.Title(language.Und)
	var b strings.Builder
	for _, t := range kindTokens {
		b.WriteString(titleCaser.String(t))
	}

	if b.Len() == 0 {
		return resourceName
	}

	kindName := b.String()
	if len(kindName) > maxGeneratedKindNameLength {
		return removeLongestDuplicateSubstring(kindTokens)
	}
	return kindName
}
// removeLongestDuplicateSubstring shortens long generated kinds by removing the
// latest duplicated token window whose normalized form already appears earlier
// in the same name. This keeps the logic generic for overlong kinds without
// hardcoding resource-specific cases.
func removeLongestDuplicateSubstring(kindTokens []string) string {
	titleCaser := cases.Title(language.Und)
	render := func(tokens []string) string {
		var b strings.Builder
		for _, token := range tokens {
			b.WriteString(titleCaser.String(token))
		}
		return b.String()
	}
	tokens := append([]string(nil), kindTokens...)
	kindName := render(tokens)
	for len(kindName) > maxGeneratedKindNameLength {
		bestStart := -1
		bestEnd := -1
		bestLength := 0
		bestWindowSize := 0
		for start := 1; start < len(tokens); start++ {
			prefix := strings.ToLower(strings.Join(tokens[:start], ""))
			if prefix == "" {
				continue
			}
			var candidate strings.Builder
			for end := start; end < len(tokens); end++ {
				candidate.WriteString(strings.ToLower(tokens[end]))
				normalized := candidate.String()
				if normalized == "" || !strings.Contains(prefix, normalized) {
					continue
				}
				if start > 0 && end < len(tokens)-1 {
					left := strings.ToLower(tokens[start-1])
					right := strings.ToLower(tokens[end+1])
					if left != "" && right != "" && (strings.Contains(left, right) || strings.Contains(right, left)) {
						continue
					}
				}
				windowSize := end - start + 1
				if len(normalized) > bestLength ||
					(len(normalized) == bestLength && windowSize > bestWindowSize) ||
					(len(normalized) == bestLength && windowSize == bestWindowSize && start > bestStart) {
					bestStart = start
					bestEnd = end
					bestLength = len(normalized)
					bestWindowSize = windowSize
				}
			}
		}
		if bestStart == -1 {
			return kindName
		}
		tokens = append(tokens[:bestStart], tokens[bestEnd+1:]...)
		kindName = render(tokens)
	}
	return kindName
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

// GroupKindOverrides applies the GroupMap during provider configuration
func GroupKindOverrides() config.ResourceOption {
	return func(r *config.Resource) {
		if f, ok := GroupMap[r.Name]; ok {
			r.ShortGroup, r.Kind = f(r.Name)
		} else {
			// Fallback to automatic detection for any unmapped resources
			r.ShortGroup, r.Kind = ServiceGroupDetector(r.Name)
		}
	}
}

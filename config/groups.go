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
	"github.com/crossplane/upjet/pkg/config"
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

	// Logging Service
	"oci_logging_log_group": func(name string) (string, string) {
		return "logging", "LogGroup"
	},
	"oci_logging_log": func(name string) (string, string) {
		return "logging", "Log"
	},
	"oci_logging_log_saved_search": func(name string) (string, string) {
		return "logging", "LogSavedSearch"
	},
	"oci_logging_unified_agent_configuration": func(name string) (string, string) {
		return "logging", "UnifiedAgentConfiguration"
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

	// File Storage Service
	"oci_file_storage_file_system": func(name string) (string, string) {
		return "filestorage", "FileSystem"
	},
	"oci_file_storage_mount_target": func(name string) (string, string) {
		return "filestorage", "MountTarget"
	},
	"oci_file_storage_export_set": func(name string) (string, string) {
		return "filestorage", "ExportSet"
	},
	"oci_file_storage_export": func(name string) (string, string) {
		return "filestorage", "Export"
	},
	"oci_file_storage_replication": func(name string) (string, string) {
		return "filestorage", "Replication"
	},
	"oci_file_storage_snapshot": func(name string) (string, string) {
		return "filestorage", "Snapshot"
	},

	// Events Service
	"oci_events_rule": func(name string) (string, string) {
		return "events", "Rule"
	},

	// MySQL Service
	"oci_mysql_backup":            func(name string) (string, string) { return "mysql", "MysqlBackup" },
	"oci_mysql_channel":           func(name string) (string, string) { return "mysql", "MysqlChannel" },
	"oci_mysql_configuration":     func(name string) (string, string) { return "mysql", "MysqlConfiguration" },
	"oci_mysql_db_system":         func(name string) (string, string) { return "mysql", "MysqlDbSystem" },
	"oci_mysql_heat_wave_cluster": func(name string) (string, string) { return "mysql", "MysqlHeatWaveCluster" },
	"oci_mysql_replica":           func(name string) (string, string) { return "mysql", "MysqlReplica" },

	// Streaming Service
	"oci_streaming_stream": func(name string) (string, string) {
		return "streaming", "Stream"
	},
	"oci_streaming_stream_pool": func(name string) (string, string) {
		return "streaming", "StreamPool"
	},
	"oci_streaming_connect_harness": func(name string) (string, string) {
		return "streaming", "ConnectHarness"
	},

	// Vault Service
	"oci_vault_secret": func(name string) (string, string) {
		return "vault", "Secret"
	},
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

/*
 * Copyright (c) 2022, 2023 Oracle and/or its affiliates
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

package core

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("oci_core_compute_capacity_report", func(r *config.Resource) {
		// REQUIRED
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	p.AddResourceConfigurator("oci_core_compute_capacity_reservation", func(r *config.Resource) {
		// REQUIRED
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	p.AddResourceConfigurator("oci_core_compute_capacity_topology", func(r *config.Resource) {
		// REQUIRED
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	p.AddResourceConfigurator("oci_core_compute_cluster", func(r *config.Resource) {
		// REQUIRED
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	p.AddResourceConfigurator("oci_core_compute_gpu_memory_cluster", func(r *config.Resource) {
		// REQUIRED
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["compute_cluster_id"] = config.Reference{
			TerraformName: "oci_core_compute_cluster",
		}
		r.References["instance_configuration_id"] = config.Reference{
			TerraformName: "oci_core_instance_configuration",
		}
	})

	p.AddResourceConfigurator("oci_core_compute_host_group", func(r *config.Resource) {
		// REQUIRED
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	p.AddResourceConfigurator("oci_core_compute_image_capability_schema", func(r *config.Resource) {
		// REQUIRED
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["image_id"] = config.Reference{
			TerraformName: "oci_core_image",
		}
	})

	p.AddResourceConfigurator("oci_core_console_history", func(r *config.Resource) {
		// REQUIRED
		r.References["instance_id"] = config.Reference{
			TerraformName: "oci_core_instance",
		}
	})

	p.AddResourceConfigurator("oci_core_dedicated_vm_host", func(r *config.Resource) {
		// REQUIRED
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	p.AddResourceConfigurator("oci_core_image", func(r *config.Resource) {
		// REQUIRED
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["instance_id"] = config.Reference{
			TerraformName: "oci_core_instance",
		}
	})

	p.AddResourceConfigurator("oci_core_instance", func(r *config.Resource) {
		// REQUIRED
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}

		// OPTIONAL
		r.References["create_vnic_details.subnet_id"] = config.Reference{
			TerraformName: "oci_core_subnet",
		}
		r.References["create_vnic_details.nsg_ids"] = config.Reference{
			TerraformName: "oci_core_network_security_group",
		}
		r.References["create_vnic_details.vlan_id"] = config.Reference{
			TerraformName: "oci_core_vlan",
		}
		r.References["dedicated_vm_host_id"] = config.Reference{
			TerraformName: "oci_core_dedicated_vm_host",
		}
		r.References["source_details.source_id"] = config.Reference{
			TerraformName: "oci_core_image",
		}
	})

	p.AddResourceConfigurator("oci_core_instance_configuration", func(r *config.Resource) {
		// REQUIRED
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	p.AddResourceConfigurator("oci_core_instance_console_connection", func(r *config.Resource) {
		// REQUIRED
		r.References["instance_id"] = config.Reference{
			TerraformName: "oci_core_instance",
		}
	})

	p.AddResourceConfigurator("oci_core_instance_pool", func(r *config.Resource) {
		// REQUIRED
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["instance_configuration_id"] = config.Reference{
			TerraformName: "oci_core_instance_configuration",
		}
	})

	p.AddResourceConfigurator("oci_core_instance_pool_instance", func(r *config.Resource) {
		// REQUIRED
		r.References["instance_id"] = config.Reference{
			TerraformName: "oci_core_instance",
		}
	})

	p.AddResourceConfigurator("oci_core_shape_management", func(r *config.Resource) {
		// REQUIRED
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["source_details.source_id"] = config.Reference{
			TerraformName: "oci_core_image",
		}
	})

	p.AddResourceConfigurator("oci_core_vcn", func(r *config.Resource) {
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
			// SelectorFieldName: "compartmentIdSelector",
		}
	})

	p.AddResourceConfigurator("oci_core_subnet", func(r *config.Resource) {
		r.References["vcn_id"] = config.Reference{
			TerraformName: "oci_core_vcn",
			// RefFieldName:      "vcnIdRef",
			// SelectorFieldName: "vcnIdSelector",
		}
		r.References["route_table_id"] = config.Reference{
			TerraformName: "oci_core_route_table",
		}

		r.References["security_list_ids"] = config.Reference{
			TerraformName:     "oci_core_security_list", // how to add array
			RefFieldName:      "SecurityListIDRefs",
			SelectorFieldName: "SecurityListIDSelector",
		}
		r.References["dhcp_options_id"] = config.Reference{
			TerraformName: "oci_core_dhcp_options", // how to add array
		}

		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	p.AddResourceConfigurator("oci_core_internet_gateway", func(r *config.Resource) {
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}

		r.References["vcn_id"] = config.Reference{
			TerraformName: "oci_core_vcn",
		}
	})

	p.AddResourceConfigurator("oci_core_nat_gateway", func(r *config.Resource) {
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}

		r.References["vcn_id"] = config.Reference{
			TerraformName: "oci_core_vcn",
		}
	})

	p.AddResourceConfigurator("oci_core_service_gateway", func(r *config.Resource) {
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}

		r.References["vcn_id"] = config.Reference{
			TerraformName: "oci_core_vcn",
		}

		r.References["route_table_id"] = config.Reference{
			TerraformName: "oci_core_route_table",
		}
	})

	p.AddResourceConfigurator("oci_core_network_security_group", func(r *config.Resource) {
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}

		r.References["vcn_id"] = config.Reference{
			TerraformName: "oci_core_vcn",
		}

		// TODO: these need more modification to work (likely custom Selectors/Extractors?)
		// r.References["source"] = config.Reference{
		// 	Type: "NetworkSecurityGroup",
		// }

		// r.References["destination"] = config.Reference{
		// 	Type: "NetworkSecurityGroup",
		// }

	})

	p.AddResourceConfigurator("oci_core_network_security_group_security_rule", func(r *config.Resource) {
		r.References["network_security_group_id"] = config.Reference{
			TerraformName: "oci_core_network_security_group",
		}
	})

	p.AddResourceConfigurator("oci_core_route_table", func(r *config.Resource) {
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}

		r.References["vcn_id"] = config.Reference{
			TerraformName: "oci_core_vcn",
		}

		r.References["route_rules.network_entity_id"] = config.Reference{
			TerraformName: "oci_core_nat_gateway",
		}

		// r.References["route_rules.network_entity_id"] = config.Reference{
		// 	Type: "InternetGateway",
		// }
	})

	p.AddResourceConfigurator("oci_core_route_table_attachment", func(r *config.Resource) {
		r.References["route_table_id"] = config.Reference{
			TerraformName: "oci_core_route_table",
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName: "oci_core_subnet",
		}
	})

	p.AddResourceConfigurator("oci_core_security_list", func(r *config.Resource) {
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}

		r.References["vcn_id"] = config.Reference{
			TerraformName: "oci_core_vcn",
		}
	})

	p.AddResourceConfigurator("oci_core_dhcp_options", func(r *config.Resource) {
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}

		r.References["vcn_id"] = config.Reference{
			TerraformName: "oci_core_vcn",
		}
	})

	p.AddResourceConfigurator("oci_core_drg", func(r *config.Resource) {
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	p.AddResourceConfigurator("oci_core_drg_attachment", func(r *config.Resource) {
		r.References["drg_id"] = config.Reference{
			TerraformName: "oci_core_drg",
		}
		r.References["drg_route_table_id"] = config.Reference{
			TerraformName: "oci_core_drg_route_table",
		}
		r.References["network_details.route_table_id"] = config.Reference{
			TerraformName: "oci_core_route_table",
		}
		r.References["network_details.id"] = config.Reference{
			TerraformName: "oci_core_vcn",
		}
	})

	p.AddResourceConfigurator("oci_core_drg_attachment_management", func(r *config.Resource) {
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["network_id"] = config.Reference{
			TerraformName: "oci_core_vcn",
		}
		r.References["drg_id"] = config.Reference{
			TerraformName: "oci_core_drg",
		}
		r.References["drg_route_table_id"] = config.Reference{
			TerraformName: "oci_core_drg_route_table",
		}
	})

	p.AddResourceConfigurator("oci_core_drg_attachments_list", func(r *config.Resource) {
		r.References["drg_id"] = config.Reference{
			TerraformName: "oci_core_drg",
		}
	})

	p.AddResourceConfigurator("oci_core_drg_route_distribution", func(r *config.Resource) {
		r.References["drg_id"] = config.Reference{
			TerraformName: "oci_core_drg",
		}
	})

	p.AddResourceConfigurator("oci_core_drg_route_distribution_statement", func(r *config.Resource) {
		r.References["drg_route_distribution_id"] = config.Reference{
			TerraformName: "oci_core_drg_route_distribution",
		}
		r.References["drg_attachment_id"] = config.Reference{
			TerraformName: "oci_core_drg_attachment",
		}
	})

	p.AddResourceConfigurator("oci_core_drg_route_table", func(r *config.Resource) {
		r.References["drg_id"] = config.Reference{
			TerraformName: "oci_core_drg",
		}
		r.References["import_drg_route_distribution_id"] = config.Reference{
			TerraformName: "oci_core_drg_route_distribution",
		}
	})

	p.AddResourceConfigurator("oci_core_drg_route_table_route_rule", func(r *config.Resource) {
		r.References["drg_route_table_id"] = config.Reference{
			TerraformName: "oci_core_drg_route_table",
		}
		r.References["next_hop_drg_attachment_id"] = config.Reference{
			TerraformName: "oci_core_drg_attachment",
		}
	})

	p.AddResourceConfigurator("oci_core_private_ip", func(r *config.Resource) {
		r.References["vlan_id"] = config.Reference{
			TerraformName: "oci_core_vlan",
		}
		r.References["vnic_id"] = config.Reference{
			TerraformName: "oci_core_vnic_attachment",
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName: "oci_core_subnet",
		}
		r.References["route_table_id"] = config.Reference{
			TerraformName: "oci_core_route_table",
		}
	})

	p.AddResourceConfigurator("oci_core_public_ip", func(r *config.Resource) {
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["private_ip_id"] = config.Reference{
			TerraformName: "oci_core_private_ip",
		}
		r.References["public_ip_pool_id"] = config.Reference{
			TerraformName: "oci_core_public_ip_pool",
		}
	})

	p.AddResourceConfigurator("oci_core_public_ip_pool", func(r *config.Resource) {
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	p.AddResourceConfigurator("oci_core_public_ip_pool_capacity", func(r *config.Resource) {
		r.References["public_ip_pool_id"] = config.Reference{
			TerraformName: "oci_core_public_ip_pool",
		}
		// Data Source seems not supported at this moment for ByoipId
	})

	p.AddResourceConfigurator("oci_core_remote_peering_connection", func(r *config.Resource) {
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["drg_id"] = config.Reference{
			TerraformName: "oci_core_drg",
		}
		r.References["peer_drg_id"] = config.Reference{
			TerraformName: "oci_core_remote_peering_connection",
		}
	})

	p.AddResourceConfigurator("oci_core_cpe", func(r *config.Resource) {
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}

	})

	p.AddResourceConfigurator("oci_core_capture_filter", func(r *config.Resource) {
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	p.AddResourceConfigurator("oci_core_ipsec", func(r *config.Resource) {
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}

		r.References["cpe_id"] = config.Reference{
			TerraformName: "oci_core_cpe",
		}

		r.References["drg_id"] = config.Reference{
			TerraformName: "oci_core_drg",
		}
	})

	p.AddResourceConfigurator("oci_core_local_peering_gateway", func(r *config.Resource) {
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}

		r.References["vcn_id"] = config.Reference{
			TerraformName: "oci_core_vcn",
		}
	})

	p.AddResourceConfigurator("oci_core_byoip_range", func(r *config.Resource) {
		r.Kind = "ByoipRange"
		r.References["byoip_range_id"] = config.Reference{
			TerraformName: "oci_core_byoip_range",
		}
	})

	p.AddResourceConfigurator("oci_core_vlan", func(r *config.Resource) {
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["vcn_id"] = config.Reference{
			TerraformName: "oci_core_vcn",
		}
	})

	p.AddResourceConfigurator("oci_core_volume", func(r *config.Resource) {
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	p.AddResourceConfigurator("oci_core_volume_backup", func(r *config.Resource) {
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["volume_id"] = config.Reference{
			TerraformName: "oci_core_volume",
		}
	})

	p.AddResourceConfigurator("oci_core_volume_backup_policy", func(r *config.Resource) {
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	p.AddResourceConfigurator("oci_core_volume_backup_policy_assignment", func(r *config.Resource) {
		r.References["asset_id"] = config.Reference{
			TerraformName: "oci_core_volume",
		}
		r.References["policy_id"] = config.Reference{
			TerraformName: "oci_core_volume_backup_policy",
		}
	})

	p.AddResourceConfigurator("oci_core_volume_group_backup", func(r *config.Resource) {
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["volume_group_id"] = config.Reference{
			TerraformName: "oci_core_volume_group",
		}
	})

	p.AddResourceConfigurator("oci_core_volume_group", func(r *config.Resource) {
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["source_details.volume_ids"] = config.Reference{
			TerraformName: "oci_core_volume",
		}
	})

	p.AddResourceConfigurator("oci_core_volume_attachment", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			TerraformName: "oci_core_instance",
		}
		r.References["volume_id"] = config.Reference{
			TerraformName: "oci_core_volume",
		}
	})
	p.AddResourceConfigurator("oci_core_vnic_attachment", func(r *config.Resource) {
		r.References["instance_id"] = config.Reference{
			TerraformName: "oci_core_instance",
		}
	})

}

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

package containerengine

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("oci_containerengine_cluster", func(r *config.Resource) {

		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}

		r.References["vcn_id"] = config.Reference{
			TerraformName: "oci_core_vcn",
		}

		r.References["endpoint_config.nsg_ids"] = config.Reference{
			TerraformName: "oci_core_network_security_group",
		}

		r.References["endpoint_config.subnet_id"] = config.Reference{
			TerraformName: "oci_core_subnet",
		}

		r.References["options.service_lb_subnet_ids"] = config.Reference{
			TerraformName:     "oci_core_subnet",
			RefFieldName:      "ServiceLBSubnetIdsRef",
			SelectorFieldName: "ServiceLBSubnetIDSelector",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("oci_containerengine_node_pool", func(r *config.Resource) {

		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}

		r.References["cluster_id"] = config.Reference{
			TerraformName: "oci_containerengine_cluster",
		}

		r.References["cluster_id"] = config.Reference{
			TerraformName: "oci_containerengine_cluster",
		}

		r.References["node_config_details.nsg_ids"] = config.Reference{
			TerraformName: "oci_core_network_security_group",
		}

		r.References["node_config_details.placement_configs.subnet_id"] = config.Reference{
			TerraformName: "oci_core_subnet",
		}

		r.LateInitializer = config.LateInitializer{
			// These are ignored because they conflict with each other.
			// See the following for more details: https://github.com/crossplane/terrajet/issues/107
			IgnoredFields: []string{
				"subnet_ids",
				"node_image_name",
				"node_image_id",
			},
		}

		r.UseAsync = true
	})

	p.AddResourceConfigurator("oci_containerengine_addon", func(r *config.Resource) {

		// Required
		r.References["cluster_id"] = config.Reference{
			TerraformName: "oci_containerengine_cluster",
		}
	})

	p.AddResourceConfigurator("oci_containerengine_cluster_complete_credential_rotation_management", func(r *config.Resource) {

		// Required
		r.References["cluster_id"] = config.Reference{
			TerraformName: "oci_containerengine_cluster",
		}
	})

	p.AddResourceConfigurator("oci_containerengine_cluster_start_credential_rotation_management", func(r *config.Resource) {

		// Required
		r.References["cluster_id"] = config.Reference{
			TerraformName: "oci_containerengine_cluster",
		}
	})

	p.AddResourceConfigurator("oci_containerengine_cluster_workload_mapping", func(r *config.Resource) {

		// Required
		r.References["cluster_id"] = config.Reference{
			TerraformName: "oci_containerengine_cluster",
		}

		// Required
		r.References["mapped_compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}

	})

	p.AddResourceConfigurator("oci_containerengine_virtual_node_pool", func(r *config.Resource) {

		// Required
		r.References["cluster_id"] = config.Reference{
			TerraformName: "oci_containerengine_cluster",
		}

		// Required
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}

		// Required
		r.References["placement_configurations.subnet_id"] = config.Reference{
			TerraformName: "oci_core_subnet",
		}

		// Optional
		r.References["nsg_ids"] = config.Reference{
			TerraformName: "oci_core_network_security_group",
		}

		// Required
		r.References["pod_configuration.subnet_id"] = config.Reference{
			TerraformName: "oci_core_subnet",
		}
		// Optional
		r.References["pod_configuration.nsg_ids"] = config.Reference{
			TerraformName: "oci_core_network_security_group",
		}
		r.UseAsync = true
	})

}

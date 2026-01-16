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

package database

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("oci_database_autonomous_database", func(r *config.Resource) {
		// REQUIRED
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	p.AddResourceConfigurator("oci_database_autonomous_database_backup", func(r *config.Resource) {
		// REQUIRED
		r.References["autonomous_database_id"] = config.Reference{
			TerraformName: "oci_database_autonomous_database",
		}
	})

	p.AddResourceConfigurator("oci_database_autonomous_database_instance_wallet_management", func(r *config.Resource) {
		// REQUIRED
		r.References["autonomous_database_id"] = config.Reference{
			TerraformName: "oci_database_autonomous_database",
		}
	})

	p.AddResourceConfigurator("oci_database_autonomous_database_saas_admin_user", func(r *config.Resource) {
		// REQUIRED
		r.References["autonomous_database_id"] = config.Reference{
			TerraformName: "oci_database_autonomous_database",
		}
	})

	p.AddResourceConfigurator("oci_database_autonomous_database_software_image", func(r *config.Resource) {
		// REQUIRED
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	p.AddResourceConfigurator("oci_database_autonomous_database_wallet", func(r *config.Resource) {
		// REQUIRED
		r.References["autonomous_database_id"] = config.Reference{
			TerraformName: "oci_database_autonomous_database",
		}
	})

	// Please note that Autonomous Exadata Infrastructure is now end of life, and support from Crossplane will be deprecated soon
	p.AddResourceConfigurator("oci_database_autonomous_exadata_infrastructure", func(r *config.Resource) {
		// REQUIRED
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName: "oci_core_subnet",
		}
	})

	p.AddResourceConfigurator("oci_database_cloud_autonomous_vm_cluster", func(r *config.Resource) {
		// REQUIRED
		r.References["cloud_exadata_infrastructure_id"] = config.Reference{
			TerraformName: "oci_database_cloud_exadata_infrastructure",
		}
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName: "oci_core_subnet",
		}
	})

	p.AddResourceConfigurator("oci_database_cloud_exadata_infrastructure", func(r *config.Resource) {
		// REQUIRED
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	p.AddResourceConfigurator("oci_database_exadata_infrastructure", func(r *config.Resource) {
		// REQUIRED
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	p.AddResourceConfigurator("oci_database_exadata_infrastructure_configure_exascale_management", func(r *config.Resource) {
		// REQUIRED
		r.References["exadata_infrastructure_id"] = config.Reference{
			TerraformName: "oci_database_exadata_infrastructure",
		}
	})

	p.AddResourceConfigurator("oci_database_exadata_iorm_config", func(r *config.Resource) {
		// REQUIRED
		r.References["db_system_id"] = config.Reference{
			TerraformName: "oci_database_db_system",
		}
	})

}

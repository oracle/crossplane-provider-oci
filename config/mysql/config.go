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

package mysql

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// oci_mysql_mysql_db_system
	p.AddResourceConfigurator("oci_mysql_mysql_db_system", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.Version = "v1alpha1"

		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}

		r.References["subnet_id"] = config.Reference{
			TerraformName: "oci_core_subnet",
		}
	})

	// oci_mysql_replica
	p.AddResourceConfigurator("oci_mysql_replica", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.Version = "v1alpha1"

		r.References["db_system_id"] = config.Reference{
			TerraformName: "oci_mysql_mysql_db_system",
		}
	})

	// oci_mysql_mysql_configuration
	p.AddResourceConfigurator("oci_mysql_mysql_configuration", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.Version = "v1alpha1"

		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}

	})

	// oci_mysql_mysql_backup
	p.AddResourceConfigurator("oci_mysql_mysql_backup", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.Version = "v1alpha1"

		r.References["db_system_id"] = config.Reference{
			TerraformName: "oci_mysql_mysql_db_system",
		}
	})

	// oci_mysql_heat_wave_cluster
	p.AddResourceConfigurator("oci_mysql_heat_wave_cluster", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.Version = "v1alpha1"

		r.References["db_system_id"] = config.Reference{
			TerraformName: "oci_mysql_mysql_db_system",
		}
	})

	// oci_mysql_channel
	p.AddResourceConfigurator("oci_mysql_channel", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.Version = "v1alpha1"

		r.References["target.db_system_id"] = config.Reference{
			TerraformName: "oci_mysql_mysql_db_system",
		}

		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})
}

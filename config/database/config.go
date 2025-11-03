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
	// Database Autonomous Database resources
	p.AddResourceConfigurator("oci_database_autonomous_database", func(r *config.Resource) {
		r.ShortGroup = "database"
		r.Kind = "AutonomousDatabase"
		// Fix field mapping issue: prevent metadata.name from being sent to Terraform
		r.ExternalName = config.IdentifierFromProvider
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName: "oci_core_subnet",
		}
	})

	p.AddResourceConfigurator("oci_database_autonomous_database_backup", func(r *config.Resource) {
		r.ShortGroup = "database"
		r.Kind = "AutonomousDatabaseBackup"
		r.ExternalName = config.IdentifierFromProvider
		r.References["autonomous_database_id"] = config.Reference{
			TerraformName: "oci_database_autonomous_database",
		}
	})

	// Database System resources
	p.AddResourceConfigurator("oci_database_db_system", func(r *config.Resource) {
		r.ShortGroup = "database"
		r.Kind = "DbSystem"
		r.ExternalName = config.IdentifierFromProvider
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName: "oci_core_subnet",
		}
		r.References["backup_subnet_id"] = config.Reference{
			TerraformName: "oci_core_subnet",
		}
	})

	p.AddResourceConfigurator("oci_database_db_home", func(r *config.Resource) {
		r.ShortGroup = "database"
		r.Kind = "DbHome"
		r.References["db_system_id"] = config.Reference{
			TerraformName: "oci_database_db_system",
		}
	})

	p.AddResourceConfigurator("oci_database_database", func(r *config.Resource) {
		r.ShortGroup = "database"
		r.Kind = "Database"
		r.References["db_home_id"] = config.Reference{
			TerraformName: "oci_database_db_home",
		}
	})

	// Exadata resources
	p.AddResourceConfigurator("oci_database_exadata_infrastructure", func(r *config.Resource) {
		r.ShortGroup = "database"
		r.Kind = "ExadataInfrastructure"
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	p.AddResourceConfigurator("oci_database_vm_cluster", func(r *config.Resource) {
		r.ShortGroup = "database"
		r.Kind = "VmCluster"
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["exadata_infrastructure_id"] = config.Reference{
			TerraformName: "oci_database_exadata_infrastructure",
		}
	})

	p.AddResourceConfigurator("oci_database_vm_cluster_network", func(r *config.Resource) {
		r.ShortGroup = "database"
		r.Kind = "VmClusterNetwork"
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["exadata_infrastructure_id"] = config.Reference{
			TerraformName: "oci_database_exadata_infrastructure",
		}
	})

	// Database backup and maintenance
	p.AddResourceConfigurator("oci_database_backup", func(r *config.Resource) {
		r.ShortGroup = "database"
		r.Kind = "Backup"
		r.References["database_id"] = config.Reference{
			TerraformName: "oci_database_database",
		}
	})

	p.AddResourceConfigurator("oci_database_backup_destination", func(r *config.Resource) {
		r.ShortGroup = "database"
		r.Kind = "BackupDestination"
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	// Database software images
	p.AddResourceConfigurator("oci_database_database_software_image", func(r *config.Resource) {
		r.ShortGroup = "database"
		r.Kind = "DatabaseSoftwareImage"
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	// Database maintenance
	p.AddResourceConfigurator("oci_database_maintenance_run", func(r *config.Resource) {
		r.ShortGroup = "database"
		r.Kind = "MaintenanceRun"
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	// Database pluggable databases
	p.AddResourceConfigurator("oci_database_pluggable_database", func(r *config.Resource) {
		r.ShortGroup = "database"
		r.Kind = "PluggableDatabase"
		r.References["container_database_id"] = config.Reference{
			TerraformName: "oci_database_database",
		}
	})

	p.AddResourceConfigurator("oci_database_pluggable_databases_remote_clone", func(r *config.Resource) {
		r.ShortGroup = "database"
		r.Kind = "PluggableDatabasesRemoteClone"
		r.References["pluggable_database_id"] = config.Reference{
			TerraformName: "oci_database_pluggable_database",
		}
		r.References["target_container_database_id"] = config.Reference{
			TerraformName: "oci_database_database",
		}
	})

	// Autonomous Container Database
	p.AddResourceConfigurator("oci_database_autonomous_container_database", func(r *config.Resource) {
		r.ShortGroup = "database"
		r.Kind = "AutonomousContainerDatabase"
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["autonomous_exadata_infrastructure_id"] = config.Reference{
			TerraformName: "oci_database_autonomous_exadata_infrastructure",
		}
	})

	// Cloud Exadata Infrastructure
	p.AddResourceConfigurator("oci_database_cloud_exadata_infrastructure", func(r *config.Resource) {
		r.ShortGroup = "database"
		r.Kind = "CloudExadataInfrastructure"
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	// Cloud VM Cluster
	p.AddResourceConfigurator("oci_database_cloud_vm_cluster", func(r *config.Resource) {
		r.ShortGroup = "database"
		r.Kind = "CloudVmCluster"
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["cloud_exadata_infrastructure_id"] = config.Reference{
			TerraformName: "oci_database_cloud_exadata_infrastructure",
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName: "oci_core_subnet",
		}
		r.References["backup_subnet_id"] = config.Reference{
			TerraformName: "oci_core_subnet",
		}
	})

	// Database migration
	p.AddResourceConfigurator("oci_database_migration", func(r *config.Resource) {
		r.ShortGroup = "database"
		r.Kind = "Migration"
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	// Data Guard Association
	p.AddResourceConfigurator("oci_database_data_guard_association", func(r *config.Resource) {
		r.ShortGroup = "database"
		r.Kind = "DataGuardAssociation"
		r.References["database_id"] = config.Reference{
			TerraformName: "oci_database_database",
		}
	})

	// Key Store
	p.AddResourceConfigurator("oci_database_key_store", func(r *config.Resource) {
		r.ShortGroup = "database"
		r.Kind = "KeyStore"
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	// External Database resources
	p.AddResourceConfigurator("oci_database_external_database_connector", func(r *config.Resource) {
		r.ShortGroup = "database"
		r.Kind = "ExternalDatabaseConnector"
		r.References["external_database_id"] = config.Reference{
			TerraformName: "oci_database_external_non_container_database",
		}
	})

	p.AddResourceConfigurator("oci_database_external_non_container_database", func(r *config.Resource) {
		r.ShortGroup = "database"
		r.Kind = "ExternalNonContainerDatabase"
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	p.AddResourceConfigurator("oci_database_external_container_database", func(r *config.Resource) {
		r.ShortGroup = "database"
		r.Kind = "ExternalContainerDatabase"
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	p.AddResourceConfigurator("oci_database_external_pluggable_database", func(r *config.Resource) {
		r.ShortGroup = "database"
		r.Kind = "ExternalPluggableDatabase"
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["external_container_database_id"] = config.Reference{
			TerraformName: "oci_database_external_container_database",
		}
	})

	// Add more database resources as needed...
	// The database service has 122 resources total - this is a representative sample
}
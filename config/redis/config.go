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

package redis

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// oci_redis_oci_cache_user
	p.AddResourceConfigurator("oci_redis_oci_cache_user", func(r *config.Resource) {

		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	// oci_redis_oci_cache_user_get_redis_cluster
	p.AddResourceConfigurator("oci_redis_oci_cache_user_get_redis_cluster", func(r *config.Resource) {

		r.References["oci_cache_user_id"] = config.Reference{
			TerraformName: "oci_redis_oci_cache_user",
		}

		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	// oci_redis_redis_cluster
	p.AddResourceConfigurator("oci_redis_redis_cluster", func(r *config.Resource) {

		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}

		r.References["subnet_id"] = config.Reference{
			TerraformName: "oci_core_subnet",
		}
	})

	// oci_redis_redis_cluster_attach_oci_cache_user
	p.AddResourceConfigurator("oci_redis_redis_cluster_attach_oci_cache_user", func(r *config.Resource) {

		r.References["redis_cluster_id"] = config.Reference{
			TerraformName: "oci_redis_redis_cluster",
		}

		r.References["oci_cache_users"] = config.Reference{
			TerraformName: "oci_redis_oci_cache_user",
		}
	})

	// oci_redis_redis_cluster_create_identity_token
	p.AddResourceConfigurator("oci_redis_redis_cluster_create_identity_token", func(r *config.Resource) {

		r.References["redis_cluster_id"] = config.Reference{
			TerraformName: "oci_redis_redis_cluster",
		}
	})

	// oci_redis_redis_cluster_detach_oci_cache_user
	p.AddResourceConfigurator("oci_redis_redis_cluster_detach_oci_cache_user", func(r *config.Resource) {

		r.References["redis_cluster_id"] = config.Reference{
			TerraformName: "oci_redis_redis_cluster",
		}

		r.References["oci_cache_users"] = config.Reference{
			TerraformName: "oci_redis_oci_cache_user",
		}
	})

	// oci_redis_redis_cluster_get_oci_cache_user
	p.AddResourceConfigurator("oci_redis_redis_cluster_get_oci_cache_user", func(r *config.Resource) {

		r.References["redis_cluster_id"] = config.Reference{
			TerraformName: "oci_redis_redis_cluster",
		}

		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	// oci_redis_oci_cache_config_set
	p.AddResourceConfigurator("oci_redis_oci_cache_config_set", func(r *config.Resource) {

		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	// oci_redis_oci_cache_config_setlist_associated_oci_cache_cluster
	p.AddResourceConfigurator("oci_redis_oci_cache_config_setlist_associated_oci_cache_cluster", func(r *config.Resource) {

		r.References["oci_cache_config_set_id"] = config.Reference{
			TerraformName: "oci_redis_oci_cache_config_set",
		}
	})
}

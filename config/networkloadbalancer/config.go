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

package networkloadbalancer

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("oci_network_load_balancer_network_load_balancer", func(r *config.Resource) {
		// REQUIRED
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName: "oci_core_subnet",
		}
	})

	p.AddResourceConfigurator("oci_network_load_balancer_backend_set", func(r *config.Resource) {
		// REQUIRED
		r.References["network_load_balancer_id"] = config.Reference{
			TerraformName: "oci_network_load_balancer_network_load_balancer",
		}
	})

	p.AddResourceConfigurator("oci_network_load_balancer_backend", func(r *config.Resource) {
		// REQUIRED
		r.References["network_load_balancer_id"] = config.Reference{
			TerraformName: "oci_network_load_balancer_network_load_balancer",
		}
	})

	p.AddResourceConfigurator("oci_network_load_balancer_network_load_balancers_backend_sets_unified", func(r *config.Resource) {
		// REQUIRED
		r.References["network_load_balancer_id"] = config.Reference{
			TerraformName: "oci_network_load_balancer_network_load_balancer",
		}
	})

	p.AddResourceConfigurator("oci_network_load_balancer_listener", func(r *config.Resource) {
		// REQUIRED
		r.References["network_load_balancer_id"] = config.Reference{
			TerraformName: "oci_network_load_balancer_network_load_balancer",
		}
	})
}

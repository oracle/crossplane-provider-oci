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

package functions

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("oci_functions_application", func(r *config.Resource) {
		r.References["subnet_ids"] = config.Reference{
			TerraformName: "oci_core_subnet",
		}
		r.References["network_security_group_ids"] = config.Reference{
			TerraformName: "oci_core_network_security_group",
		}
	})
}

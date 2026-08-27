/*
 * Copyright (c) 2026 Oracle and/or its affiliates
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

package bds

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures BDS resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("oci_bds_bds_instance", configureBdsInstance)
}

func configureBdsInstance(r *config.Resource) {
	r.References["secret_id"] = config.Reference{TerraformName: "oci_vault_secret"}
}

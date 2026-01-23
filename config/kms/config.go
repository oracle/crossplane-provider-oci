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

package kms

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("oci_kms_ekms_private_endpoint", func(r *config.Resource) {
		// REQUIRED
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName: "oci_core_subnet",
		}
	})

	p.AddResourceConfigurator("oci_kms_encrypted_data", func(r *config.Resource) {
		// REQUIRED
		r.References["key_id"] = config.Reference{
			TerraformName: "oci_kms_key",
		}
		//r.References["crypto_endpoint"] = config.Reference{
		// // Need an extractor here to identify crypto endpoint
		//	TerraformName: "oci_kms_vault.crypto_endpoint",
		//}
	})

	p.AddResourceConfigurator("oci_kms_generated_key", func(r *config.Resource) {
		// REQUIRED
		r.References["key_id"] = config.Reference{
			TerraformName: "oci_kms_key",
		}
		//r.References["crypto_endpoint"] = config.Reference{
		// // Need an extractor here to identify crypto endpoint
		//	TerraformName: "oci_kms_vault.crypto_endpoint",
		//}
	})

	p.AddResourceConfigurator("oci_kms_key", func(r *config.Resource) {
		// REQUIRED
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
		//r.References["management_endpoint"] = config.Reference{
		// // Need an extractor here to identify management endpoint
		//	TerraformName: "oci_kms_vault.management_endpoint",
		//}
	})

	p.AddResourceConfigurator("oci_kms_key_version", func(r *config.Resource) {
		// REQUIRED
		r.References["key_id"] = config.Reference{
			TerraformName: "oci_kms_key",
		}
		//r.References["management_endpoint"] = config.Reference{
		// // Need an extractor here to identify management endpoint
		//	TerraformName: "oci_kms_vault.management_endpoint",
		//}
	})

	p.AddResourceConfigurator("oci_kms_sign", func(r *config.Resource) {
		// REQUIRED
		r.References["key_id"] = config.Reference{
			TerraformName: "oci_kms_key",
		}
		//r.References["crypto_endpoint"] = config.Reference{
		// // Need an extractor here to identify crypto endpoint
		//	TerraformName: "oci_kms_vault.crypto_endpoint",
		//}
	})

	p.AddResourceConfigurator("oci_kms_vault", func(r *config.Resource) {
		// REQUIRED
		r.References["compartment_id"] = config.Reference{
			TerraformName: "oci_identity_compartment",
		}
	})

	p.AddResourceConfigurator("oci_kms_vault_replication", func(r *config.Resource) {
		// REQUIRED
		r.References["vault_id"] = config.Reference{
			TerraformName: "oci_kms_vault",
		}
	})

	p.AddResourceConfigurator("oci_kms_verify", func(r *config.Resource) {
		// REQUIRED
		r.References["key_id"] = config.Reference{
			TerraformName: "oci_kms_key",
		}
		r.References["key_version_id"] = config.Reference{
			TerraformName: "oci_kms_key_version",
		}
		//r.References["signature"] = config.Reference{
		// // Need an extractor here to identify signature (base64 encoded binary data object)
		//	TerraformName: "oci_kms_sign.signature",
		//}
		//r.References["crypto_endpoint"] = config.Reference{
		// // Need an extractor here to identify crypto endpoint
		//	TerraformName: "oci_kms_vault.crypto_endpoint",
		//}
	})
}

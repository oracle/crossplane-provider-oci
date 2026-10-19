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

package dbmulticloud

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures DB Multicloud resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("oci_dbmulticloud_oracle_db_aws_key_pool", func(r *config.Resource) {
		r.OverrideFieldNames = map[string]string{
			"ReplicationMetadataReplicationDetailsInitParameters":                   "OracleDbAwsKeyPoolReplicationMetadataReplicationDetailsInitParameters",
			"ReplicationMetadataReplicationDetailsObservation":                      "OracleDbAwsKeyPoolReplicationMetadataReplicationDetailsObservation",
			"ReplicationMetadataReplicationDetailsParameters":                       "OracleDbAwsKeyPoolReplicationMetadataReplicationDetailsParameters",
			"OracleDbAzureVaultReplicationMetadataReplicationDetailsInitParameters": "ReplicationMetadataReplicationDetailsInitParameters",
			"OracleDbAzureVaultReplicationMetadataReplicationDetailsObservation":    "ReplicationMetadataReplicationDetailsObservation",
			"OracleDbAzureVaultReplicationMetadataReplicationDetailsParameters":     "ReplicationMetadataReplicationDetailsParameters",
		}
	})
}

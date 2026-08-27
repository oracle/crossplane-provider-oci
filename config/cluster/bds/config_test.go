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

import (
	"testing"

	"github.com/crossplane/upjet/v2/pkg/config"
)

func TestConfigureBdsInstance(t *testing.T) {
	r := &config.Resource{References: config.References{}}
	configureBdsInstance(r)

	if got := r.References["secret_id"].TerraformName; got != "oci_vault_secret" {
		t.Fatalf("secret_id TerraformName = %q, want %q", got, "oci_vault_secret")
	}
}

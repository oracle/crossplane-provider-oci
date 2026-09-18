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

package config

import (
	frameworkprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfoci "github.com/oracle/terraform-provider-oci/oci"
)

func terraformSDKProvider(resourceNames []string) *schema.Provider {
	if resourceNames == nil {
		return tfoci.Provider()
	}
	provider, err := tfoci.ProviderForResources(resourceNames...)
	if err != nil {
		panic(err)
	}
	return provider
}

func terraformFrameworkProvider() frameworkprovider.Provider {
	return tfoci.New()
}

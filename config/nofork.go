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
	"encoding/json"
	"regexp"
	"slices"
	"sort"
	"sync"
)

// terraformPluginFrameworkIncludeList is intentionally empty until OCI
// Terraform Plugin Framework resources are validated for no-fork routing.
var terraformPluginFrameworkIncludeList = []string{}

type terraformProviderSchema struct {
	ProviderSchemas map[string]struct {
		ResourceSchemas map[string]json.RawMessage `json:"resource_schemas"`
	} `json:"provider_schemas"`
}

type terraformResourceRouting struct {
	names          []string
	sdkV2Resources map[string]struct{}
	sdkV2Regexes   []string
}

var loadTerraformResourceRouting = sync.OnceValue(func() terraformResourceRouting {
	names := parseTerraformResourceNamesFromSchema()
	routing := terraformResourceRouting{
		names:          names,
		sdkV2Resources: make(map[string]struct{}, len(names)),
		sdkV2Regexes:   make([]string, 0, len(names)),
	}
	for _, resource := range names {
		routing.sdkV2Resources[resource] = struct{}{}
		routing.sdkV2Regexes = append(routing.sdkV2Regexes, exactResourceRegex(resource))
	}
	return routing
})

func terraformPluginSDKIncludeList() []string {
	return slices.Clone(loadTerraformResourceRouting().sdkV2Regexes)
}

func terraformPluginSDKResourcesForRuntime(service string) []string {
	resources, ok := runtimeTerraformResources[service]
	if !ok {
		panic("no Terraform runtime resource manifest for service " + service)
	}
	sdkResources := make([]string, 0, len(resources))
	for _, resource := range resources {
		if IsSDKv2Resource(resource) {
			sdkResources = append(sdkResources, resource)
		}
	}
	return sdkResources
}

func terraformPluginSDKIncludeListForResources(resources []string) []string {
	includeList := make([]string, 0, len(resources))
	for _, resource := range resources {
		includeList = append(includeList, exactResourceRegex(resource))
	}
	return includeList
}

// SDKv2ResourcePredicateForRuntime returns a predicate that accepts only the
// Terraform resources assigned to the specified service runtime.
func SDKv2ResourcePredicateForRuntime(service string) func(string) bool {
	resources := terraformPluginSDKResourcesForRuntime(service)
	resourceSet := make(map[string]struct{}, len(resources))
	for _, resource := range resources {
		resourceSet[resource] = struct{}{}
	}
	return func(terraformResourceType string) bool {
		_, ok := resourceSet[terraformResourceType]
		return ok
	}
}

// IsSDKv2Resource reports whether the Terraform resource type is routed through
// Upjet's in-process Terraform Plugin SDKv2 connector.
func IsSDKv2Resource(terraformResourceType string) bool {
	_, ok := loadTerraformResourceRouting().sdkV2Resources[terraformResourceType]
	return ok
}

// HasFrameworkResources reports whether any resources are routed through
// Upjet's in-process Terraform Plugin Framework connector.
func HasFrameworkResources() bool {
	return len(terraformPluginFrameworkIncludeList) != 0
}

func terraformResourceNamesFromSchema() []string {
	return slices.Clone(loadTerraformResourceRouting().names)
}

func parseTerraformResourceNamesFromSchema() []string {
	var schema terraformProviderSchema
	if err := json.Unmarshal([]byte(providerSchema), &schema); err != nil {
		panic(err)
	}

	var resources []string
	for _, provider := range schema.ProviderSchemas {
		for name := range provider.ResourceSchemas {
			resources = append(resources, name)
		}
	}
	sort.Strings(resources)
	return resources
}

func exactResourceRegex(resource string) string {
	return "^" + regexp.QuoteMeta(resource) + "$"
}

func resourceMatches(resource string, regexes []string) bool {
	for _, pattern := range regexes {
		if regexp.MustCompile(pattern).MatchString(resource) {
			return true
		}
	}
	return false
}

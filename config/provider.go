/*
 * Copyright (c) 2023 Oracle and/or its affiliates
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
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/registry/reference"

	"github.com/oracle/provider-oci/config/cluster"
	"github.com/oracle/provider-oci/config/namespaced"
	"github.com/oracle/provider-oci/hack"
)

const (
	resourcePrefix = "oci"
	modulePath     = "github.com/oracle/provider-oci"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// ProblematicResources returns a list of regex patterns for resources that should be
// skipped during generation due to known issues or incompatibilities.
// These resources can be added to support later after resolving their specific issues.
func ProblematicResources() []string {
	return []string{
		// OCI resources do not have a data section like AWS/GCP, so we do not need to specifically skip them.
		// Explicity using `*_data_*` skips oci_identity_data_plane_* resources hence commenting it out.
		// Skip data sources (not needed for managed resources)
		// `.*_data_.*`,

		// Skip test resources (internal testing only)
		`.*_test.*`,

		// Skip deprecated resources
		`.*_deprecated.*`,

		// Known problematic resources that need special handling
		`oci_network_firewall_network_firewall_policy_service_list$`,     // Name collision: generates duplicate types
		`oci_network_firewall_network_firewall_policy_url_list$`,         // Similar potential naming conflict
		`oci_network_firewall_network_firewall_policy_application_list$`, // Similar potential naming conflict
		`oci_load_balancer_backendset$`,                                  // Alias for oci_load_balancer_backend_set
		`oci_load_balancer$`,                                             // Alias for oci_load_balancer_load_balancer
		`oci_objectstorage_namespace_metadata$`,                          // Does not support import and manages tenancy-level S3/Swift defaults

		// Add more specific resources here as we discover generation issues
	}
}

func newProvider(rootGroup string, register func(*ujconfig.Provider), useTerraformJSONGenerationSchema bool) *ujconfig.Provider {
	sdkProvider := terraformSDKProvider()
	if useTerraformJSONGenerationSchema {
		applyTerraformJSONGenerationSchemas(sdkProvider, []byte(providerSchema))
	}

	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup(rootGroup),
		// Disable Upjet's Terraform CLI route. All generated OCI resources are
		// routed through in-process no-fork connectors.
		ujconfig.WithIncludeList(nil),
		ujconfig.WithTerraformProvider(sdkProvider),
		ujconfig.WithTerraformPluginSDKIncludeList(terraformPluginSDKIncludeList()),
		ujconfig.WithTerraformPluginFrameworkProvider(terraformFrameworkProvider()),
		ujconfig.WithTerraformPluginFrameworkIncludeList(terraformPluginFrameworkIncludeList),
		ujconfig.WithSkipList(ProblematicResources()),
		ujconfig.WithDefaultResourceOptions(
			GroupKindOverrides(),
			ExternalNameConfigurations(),
			AutoExternalNameConfiguration(), // Automatic external name for unconfigured resources

		),
		ujconfig.WithReferenceInjectors([]ujconfig.ReferenceInjector{
			reference.NewInjector(modulePath),
			NewStaticReferenceInjector(),
		}),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithMainTemplate(hack.MainTemplate),
	)

	register(pc)
	configureSDKv2ScalarMapServerSideApplyMergeStrategies(pc)
	pc.ConfigureResources()
	return pc
}

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	return newProvider("oci.upbound.io", func(pc *ujconfig.Provider) {
		for _, configure := range cluster.ProviderConfiguration {
			configure(pc)
		}
	}, false)
}

// GetProviderNamespaced returns namespaced provider configuration.
func GetProviderNamespaced() *ujconfig.Provider {
	return newProvider("oci.m.upbound.io", func(pc *ujconfig.Provider) {
		for _, configure := range namespaced.ProviderConfiguration {
			configure(pc)
		}
	}, false)
}

// GetProviderForGeneration returns the cluster-scoped provider configuration
// with compatibility transformations that preserve the existing public API and
// CRD shapes. Runtime controllers must use GetProvider so the embedded SDKv2
// connector retains the provider's native schema.
func GetProviderForGeneration() *ujconfig.Provider {
	return newProvider("oci.upbound.io", func(pc *ujconfig.Provider) {
		for _, configure := range cluster.ProviderConfiguration {
			configure(pc)
		}
	}, true)
}

// GetProviderNamespacedForGeneration returns the namespaced provider
// configuration with generation-only public API compatibility transformations.
func GetProviderNamespacedForGeneration() *ujconfig.Provider {
	return newProvider("oci.m.upbound.io", func(pc *ujconfig.Provider) {
		for _, configure := range namespaced.ProviderConfiguration {
			configure(pc)
		}
	}, true)
}

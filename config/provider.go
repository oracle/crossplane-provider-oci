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

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/oracle/provider-oci/config/artifacts"
	"github.com/oracle/provider-oci/config/certificatesmanagement"
	"github.com/oracle/provider-oci/config/containerengine"
	"github.com/oracle/provider-oci/config/core"
	"github.com/oracle/provider-oci/config/dns"
	"github.com/oracle/provider-oci/config/events"
	"github.com/oracle/provider-oci/config/filestorage"
	"github.com/oracle/provider-oci/config/functions"
	"github.com/oracle/provider-oci/config/healthchecks"
	"github.com/oracle/provider-oci/config/identity"
	"github.com/oracle/provider-oci/config/kms"
	"github.com/oracle/provider-oci/config/loadbalancer"
	"github.com/oracle/provider-oci/config/logging"
	"github.com/oracle/provider-oci/config/monitoring"
	"github.com/oracle/provider-oci/config/mysql"
	"github.com/oracle/provider-oci/config/networkfirewall"
	"github.com/oracle/provider-oci/config/networkloadbalancer"
	"github.com/oracle/provider-oci/config/objectstorage"
	"github.com/oracle/provider-oci/config/ons"
	"github.com/oracle/provider-oci/config/streaming"
	"github.com/oracle/provider-oci/config/vault"
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

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("oci.upbound.io"),
		ujconfig.WithIncludeList(append(ExternalNameConfigured(), "oci_mysql_.*")),
		ujconfig.WithDefaultResourceOptions(
			GroupKindOverrides(),
			ExternalNameConfigurations(),
			AutoExternalNameConfiguration(), // Automatic external name for unconfigured resources

		),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithMainTemplate(hack.MainTemplate),
	)

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		objectstorage.Configure,
		identity.Configure,
		core.Configure,
		kms.Configure,
		containerengine.Configure,
		artifacts.Configure,
		ons.Configure,
		networkloadbalancer.Configure,
		dns.Configure,
		healthchecks.Configure,
		functions.Configure,
		networkfirewall.Configure,
		logging.Configure,
		monitoring.Configure,
		loadbalancer.Configure,
		certificatesmanagement.Configure,
		filestorage.Configure,
		events.Configure,
		vault.Configure,
		streaming.Configure,
		mysql.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

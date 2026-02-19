/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	agent "github.com/oracle/provider-oci/internal/controller/cloudbridge/agent"
	agentdependency "github.com/oracle/provider-oci/internal/controller/cloudbridge/agentdependency"
	agentplugin "github.com/oracle/provider-oci/internal/controller/cloudbridge/agentplugin"
	asset "github.com/oracle/provider-oci/internal/controller/cloudbridge/asset"
	assetsource "github.com/oracle/provider-oci/internal/controller/cloudbridge/assetsource"
	discoveryschedule "github.com/oracle/provider-oci/internal/controller/cloudbridge/discoveryschedule"
	environment "github.com/oracle/provider-oci/internal/controller/cloudbridge/environment"
	inventory "github.com/oracle/provider-oci/internal/controller/cloudbridge/inventory"
)

// Setup_cloudbridge creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cloudbridge(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		agent.Setup,
		agentdependency.Setup,
		agentplugin.Setup,
		asset.Setup,
		assetsource.Setup,
		discoveryschedule.Setup,
		environment.Setup,
		inventory.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

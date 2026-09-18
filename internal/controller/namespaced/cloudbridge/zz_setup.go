/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	agent "github.com/oracle/provider-oci/internal/controller/namespaced/cloudbridge/agent"
	agentdependency "github.com/oracle/provider-oci/internal/controller/namespaced/cloudbridge/agentdependency"
	agentplugin "github.com/oracle/provider-oci/internal/controller/namespaced/cloudbridge/agentplugin"
	asset "github.com/oracle/provider-oci/internal/controller/namespaced/cloudbridge/asset"
	assetsource "github.com/oracle/provider-oci/internal/controller/namespaced/cloudbridge/assetsource"
	discoveryschedule "github.com/oracle/provider-oci/internal/controller/namespaced/cloudbridge/discoveryschedule"
	environment "github.com/oracle/provider-oci/internal/controller/namespaced/cloudbridge/environment"
	inventory "github.com/oracle/provider-oci/internal/controller/namespaced/cloudbridge/inventory"
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

// SetupGated_cloudbridge creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_cloudbridge(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		agent.SetupGated,
		agentdependency.SetupGated,
		agentplugin.SetupGated,
		asset.SetupGated,
		assetsource.SetupGated,
		discoveryschedule.SetupGated,
		environment.SetupGated,
		inventory.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

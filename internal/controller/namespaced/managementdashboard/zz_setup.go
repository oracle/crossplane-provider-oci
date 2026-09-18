/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	managementdashboardsimport "github.com/oracle/provider-oci/internal/controller/namespaced/managementdashboard/managementdashboardsimport"
	managementsavedsearch "github.com/oracle/provider-oci/internal/controller/namespaced/managementdashboard/managementsavedsearch"
)

// Setup_managementdashboard creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_managementdashboard(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		managementdashboardsimport.Setup,
		managementsavedsearch.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_managementdashboard creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_managementdashboard(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		managementdashboardsimport.SetupGated,
		managementsavedsearch.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

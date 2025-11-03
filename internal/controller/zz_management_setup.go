/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	agentmanagementagent "github.com/oracle/provider-oci/internal/controller/management/agentmanagementagent"
	agentmanagementagentinstallkey "github.com/oracle/provider-oci/internal/controller/management/agentmanagementagentinstallkey"
	agentnamedcredential "github.com/oracle/provider-oci/internal/controller/management/agentnamedcredential"
	dashboardmanagementdashboardsimport "github.com/oracle/provider-oci/internal/controller/management/dashboardmanagementdashboardsimport"
)

// Setup_management creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_management(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		agentmanagementagent.Setup,
		agentmanagementagentinstallkey.Setup,
		agentnamedcredential.Setup,
		dashboardmanagementdashboardsimport.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

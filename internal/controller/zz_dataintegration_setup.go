/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	workspace "github.com/oracle/provider-oci/internal/controller/dataintegration/workspace"
	workspaceapplication "github.com/oracle/provider-oci/internal/controller/dataintegration/workspaceapplication"
	workspaceapplicationpatch "github.com/oracle/provider-oci/internal/controller/dataintegration/workspaceapplicationpatch"
	workspaceapplicationschedule "github.com/oracle/provider-oci/internal/controller/dataintegration/workspaceapplicationschedule"
	workspaceapplicationtaskschedule "github.com/oracle/provider-oci/internal/controller/dataintegration/workspaceapplicationtaskschedule"
	workspaceexportrequest "github.com/oracle/provider-oci/internal/controller/dataintegration/workspaceexportrequest"
	workspacefolder "github.com/oracle/provider-oci/internal/controller/dataintegration/workspacefolder"
	workspaceimportrequest "github.com/oracle/provider-oci/internal/controller/dataintegration/workspaceimportrequest"
	workspaceproject "github.com/oracle/provider-oci/internal/controller/dataintegration/workspaceproject"
	workspacetask "github.com/oracle/provider-oci/internal/controller/dataintegration/workspacetask"
)

// Setup_dataintegration creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dataintegration(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		workspace.Setup,
		workspaceapplication.Setup,
		workspaceapplicationpatch.Setup,
		workspaceapplicationschedule.Setup,
		workspaceapplicationtaskschedule.Setup,
		workspaceexportrequest.Setup,
		workspacefolder.Setup,
		workspaceimportrequest.Setup,
		workspaceproject.Setup,
		workspacetask.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

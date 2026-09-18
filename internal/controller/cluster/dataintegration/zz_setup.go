/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	workspace "github.com/oracle/provider-oci/internal/controller/cluster/dataintegration/workspace"
	workspaceapplication "github.com/oracle/provider-oci/internal/controller/cluster/dataintegration/workspaceapplication"
	workspaceapplicationpatch "github.com/oracle/provider-oci/internal/controller/cluster/dataintegration/workspaceapplicationpatch"
	workspaceapplicationschedule "github.com/oracle/provider-oci/internal/controller/cluster/dataintegration/workspaceapplicationschedule"
	workspaceapplicationtaskschedule "github.com/oracle/provider-oci/internal/controller/cluster/dataintegration/workspaceapplicationtaskschedule"
	workspaceexportrequest "github.com/oracle/provider-oci/internal/controller/cluster/dataintegration/workspaceexportrequest"
	workspacefolder "github.com/oracle/provider-oci/internal/controller/cluster/dataintegration/workspacefolder"
	workspaceimportrequest "github.com/oracle/provider-oci/internal/controller/cluster/dataintegration/workspaceimportrequest"
	workspaceproject "github.com/oracle/provider-oci/internal/controller/cluster/dataintegration/workspaceproject"
	workspacetask "github.com/oracle/provider-oci/internal/controller/cluster/dataintegration/workspacetask"
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

// SetupGated_dataintegration creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_dataintegration(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		workspace.SetupGated,
		workspaceapplication.SetupGated,
		workspaceapplicationpatch.SetupGated,
		workspaceapplicationschedule.SetupGated,
		workspaceapplicationtaskschedule.SetupGated,
		workspaceexportrequest.SetupGated,
		workspacefolder.SetupGated,
		workspaceimportrequest.SetupGated,
		workspaceproject.SetupGated,
		workspacetask.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

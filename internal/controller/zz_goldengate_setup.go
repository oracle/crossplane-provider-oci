/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	gateconnection "github.com/oracle/provider-oci/internal/controller/goldengate/gateconnection"
	gateconnectionassignment "github.com/oracle/provider-oci/internal/controller/goldengate/gateconnectionassignment"
	gatedatabaseregistration "github.com/oracle/provider-oci/internal/controller/goldengate/gatedatabaseregistration"
	gatedeployment "github.com/oracle/provider-oci/internal/controller/goldengate/gatedeployment"
	gatedeploymentbackup "github.com/oracle/provider-oci/internal/controller/goldengate/gatedeploymentbackup"
	gatedeploymentcertificate "github.com/oracle/provider-oci/internal/controller/goldengate/gatedeploymentcertificate"
	gatepipeline "github.com/oracle/provider-oci/internal/controller/goldengate/gatepipeline"
)

// Setup_goldengate creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_goldengate(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		gateconnection.Setup,
		gateconnectionassignment.Setup,
		gatedatabaseregistration.Setup,
		gatedeployment.Setup,
		gatedeploymentbackup.Setup,
		gatedeploymentcertificate.Setup,
		gatepipeline.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

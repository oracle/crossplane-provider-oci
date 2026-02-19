/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	connection "github.com/oracle/provider-oci/internal/controller/goldengate/connection"
	connectionassignment "github.com/oracle/provider-oci/internal/controller/goldengate/connectionassignment"
	databaseregistration "github.com/oracle/provider-oci/internal/controller/goldengate/databaseregistration"
	deployment "github.com/oracle/provider-oci/internal/controller/goldengate/deployment"
	deploymentbackup "github.com/oracle/provider-oci/internal/controller/goldengate/deploymentbackup"
	deploymentcertificate "github.com/oracle/provider-oci/internal/controller/goldengate/deploymentcertificate"
	pipeline "github.com/oracle/provider-oci/internal/controller/goldengate/pipeline"
)

// Setup_goldengate creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_goldengate(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		connection.Setup,
		connectionassignment.Setup,
		databaseregistration.Setup,
		deployment.Setup,
		deploymentbackup.Setup,
		deploymentcertificate.Setup,
		pipeline.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

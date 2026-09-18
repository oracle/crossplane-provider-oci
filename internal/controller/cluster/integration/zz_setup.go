/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	integrationinstance "github.com/oracle/provider-oci/internal/controller/cluster/integration/integrationinstance"
	oraclemanagedcustomendpoint "github.com/oracle/provider-oci/internal/controller/cluster/integration/oraclemanagedcustomendpoint"
	privateendpointoutboundconnection "github.com/oracle/provider-oci/internal/controller/cluster/integration/privateendpointoutboundconnection"
)

// Setup_integration creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_integration(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		integrationinstance.Setup,
		oraclemanagedcustomendpoint.Setup,
		privateendpointoutboundconnection.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_integration creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_integration(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		integrationinstance.SetupGated,
		oraclemanagedcustomendpoint.SetupGated,
		privateendpointoutboundconnection.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

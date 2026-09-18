/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	application "github.com/oracle/provider-oci/internal/controller/namespaced/dataflow/application"
	invokerun "github.com/oracle/provider-oci/internal/controller/namespaced/dataflow/invokerun"
	pool "github.com/oracle/provider-oci/internal/controller/namespaced/dataflow/pool"
	privateendpoint "github.com/oracle/provider-oci/internal/controller/namespaced/dataflow/privateendpoint"
	runstatement "github.com/oracle/provider-oci/internal/controller/namespaced/dataflow/runstatement"
	sqlendpoint "github.com/oracle/provider-oci/internal/controller/namespaced/dataflow/sqlendpoint"
)

// Setup_dataflow creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dataflow(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		invokerun.Setup,
		pool.Setup,
		privateendpoint.Setup,
		runstatement.Setup,
		sqlendpoint.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_dataflow creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_dataflow(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.SetupGated,
		invokerun.SetupGated,
		pool.SetupGated,
		privateendpoint.SetupGated,
		runstatement.SetupGated,
		sqlendpoint.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

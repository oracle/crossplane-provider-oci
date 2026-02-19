/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	application "github.com/oracle/provider-oci/internal/controller/dataflow/application"
	invokerun "github.com/oracle/provider-oci/internal/controller/dataflow/invokerun"
	pool "github.com/oracle/provider-oci/internal/controller/dataflow/pool"
	privateendpoint "github.com/oracle/provider-oci/internal/controller/dataflow/privateendpoint"
	runstatement "github.com/oracle/provider-oci/internal/controller/dataflow/runstatement"
	sqlendpoint "github.com/oracle/provider-oci/internal/controller/dataflow/sqlendpoint"
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

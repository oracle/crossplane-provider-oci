/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	dataflowapplication "github.com/oracle/provider-oci/internal/controller/dataflow/dataflowapplication"
	dataflowprivateendpoint "github.com/oracle/provider-oci/internal/controller/dataflow/dataflowprivateendpoint"
	invokerun "github.com/oracle/provider-oci/internal/controller/dataflow/invokerun"
	pool "github.com/oracle/provider-oci/internal/controller/dataflow/pool"
	runstatement "github.com/oracle/provider-oci/internal/controller/dataflow/runstatement"
	sqlendpoint "github.com/oracle/provider-oci/internal/controller/dataflow/sqlendpoint"
)

// Setup_dataflow creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dataflow(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		dataflowapplication.Setup,
		dataflowprivateendpoint.Setup,
		invokerun.Setup,
		pool.Setup,
		runstatement.Setup,
		sqlendpoint.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

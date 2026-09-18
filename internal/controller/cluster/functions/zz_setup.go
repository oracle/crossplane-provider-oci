/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	application "github.com/oracle/provider-oci/internal/controller/cluster/functions/application"
	function "github.com/oracle/provider-oci/internal/controller/cluster/functions/function"
	invokefunction "github.com/oracle/provider-oci/internal/controller/cluster/functions/invokefunction"
)

// Setup_functions creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_functions(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		function.Setup,
		invokefunction.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_functions creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_functions(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.SetupGated,
		function.SetupGated,
		invokefunction.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

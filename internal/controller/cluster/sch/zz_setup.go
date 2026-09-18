/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	serviceconnector "github.com/oracle/provider-oci/internal/controller/cluster/sch/serviceconnector"
)

// Setup_sch creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_sch(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		serviceconnector.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_sch creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_sch(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		serviceconnector.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

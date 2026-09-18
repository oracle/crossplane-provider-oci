/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	cccinfrastructure "github.com/oracle/provider-oci/internal/controller/cluster/computecloudatcustomer/cccinfrastructure"
	cccupgradeschedule "github.com/oracle/provider-oci/internal/controller/cluster/computecloudatcustomer/cccupgradeschedule"
)

// Setup_computecloudatcustomer creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_computecloudatcustomer(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cccinfrastructure.Setup,
		cccupgradeschedule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_computecloudatcustomer creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_computecloudatcustomer(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cccinfrastructure.SetupGated,
		cccupgradeschedule.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

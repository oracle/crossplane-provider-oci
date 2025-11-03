/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	cloudatcustomercccinfrastructure "github.com/oracle/provider-oci/internal/controller/computemanagement/cloudatcustomercccinfrastructure"
	cloudatcustomercccupgradeschedule "github.com/oracle/provider-oci/internal/controller/computemanagement/cloudatcustomercccupgradeschedule"
)

// Setup_computemanagement creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_computemanagement(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cloudatcustomercccinfrastructure.Setup,
		cloudatcustomercccupgradeschedule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

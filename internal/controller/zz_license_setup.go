/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	managerconfiguration "github.com/oracle/provider-oci/internal/controller/license/managerconfiguration"
	managerlicenserecord "github.com/oracle/provider-oci/internal/controller/license/managerlicenserecord"
	managerproductlicense "github.com/oracle/provider-oci/internal/controller/license/managerproductlicense"
)

// Setup_license creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_license(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		managerconfiguration.Setup,
		managerlicenserecord.Setup,
		managerproductlicense.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

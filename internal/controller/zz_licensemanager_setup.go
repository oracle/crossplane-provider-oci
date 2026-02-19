/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	configuration "github.com/oracle/provider-oci/internal/controller/licensemanager/configuration"
	licenserecord "github.com/oracle/provider-oci/internal/controller/licensemanager/licenserecord"
	productlicense "github.com/oracle/provider-oci/internal/controller/licensemanager/productlicense"
)

// Setup_licensemanager creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_licensemanager(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		configuration.Setup,
		licenserecord.Setup,
		productlicense.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

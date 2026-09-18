/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	configuration "github.com/oracle/provider-oci/internal/controller/cluster/licensemanager/configuration"
	licenserecord "github.com/oracle/provider-oci/internal/controller/cluster/licensemanager/licenserecord"
	productlicense "github.com/oracle/provider-oci/internal/controller/cluster/licensemanager/productlicense"
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

// SetupGated_licensemanager creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_licensemanager(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		configuration.SetupGated,
		licenserecord.SetupGated,
		productlicense.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

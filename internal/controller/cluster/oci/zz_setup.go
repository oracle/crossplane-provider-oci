/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	productcataloginternaladminproduct "github.com/oracle/provider-oci/internal/controller/cluster/oci/productcataloginternaladminproduct"
	productcataloginternalproduct "github.com/oracle/provider-oci/internal/controller/cluster/oci/productcataloginternalproduct"
)

// Setup_oci creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_oci(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		productcataloginternaladminproduct.Setup,
		productcataloginternalproduct.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_oci creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_oci(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		productcataloginternaladminproduct.SetupGated,
		productcataloginternalproduct.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

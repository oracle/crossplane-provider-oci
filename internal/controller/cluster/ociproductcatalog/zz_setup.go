/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	internaladminproduct "github.com/oracle/provider-oci/internal/controller/cluster/ociproductcatalog/internaladminproduct"
	internalproduct "github.com/oracle/provider-oci/internal/controller/cluster/ociproductcatalog/internalproduct"
)

// Setup_ociproductcatalog creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_ociproductcatalog(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		internaladminproduct.Setup,
		internalproduct.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_ociproductcatalog creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_ociproductcatalog(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		internaladminproduct.SetupGated,
		internalproduct.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

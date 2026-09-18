/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	privateapplication "github.com/oracle/provider-oci/internal/controller/cluster/servicecatalog/privateapplication"
	servicecatalog "github.com/oracle/provider-oci/internal/controller/cluster/servicecatalog/servicecatalog"
	servicecatalogassociation "github.com/oracle/provider-oci/internal/controller/cluster/servicecatalog/servicecatalogassociation"
)

// Setup_servicecatalog creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_servicecatalog(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		privateapplication.Setup,
		servicecatalog.Setup,
		servicecatalogassociation.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_servicecatalog creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_servicecatalog(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		privateapplication.SetupGated,
		servicecatalog.SetupGated,
		servicecatalogassociation.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	catalogprivateapplication "github.com/oracle/provider-oci/internal/controller/service/catalogprivateapplication"
	catalogservicecatalog "github.com/oracle/provider-oci/internal/controller/service/catalogservicecatalog"
	catalogservicecatalogassociation "github.com/oracle/provider-oci/internal/controller/service/catalogservicecatalogassociation"
)

// Setup_service creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_service(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		catalogprivateapplication.Setup,
		catalogservicecatalog.Setup,
		catalogservicecatalogassociation.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

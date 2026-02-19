/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	catalog "github.com/oracle/provider-oci/internal/controller/datacatalog/catalog"
	catalogprivateendpoint "github.com/oracle/provider-oci/internal/controller/datacatalog/catalogprivateendpoint"
	connection "github.com/oracle/provider-oci/internal/controller/datacatalog/connection"
	dataasset "github.com/oracle/provider-oci/internal/controller/datacatalog/dataasset"
	metastore "github.com/oracle/provider-oci/internal/controller/datacatalog/metastore"
)

// Setup_datacatalog creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_datacatalog(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		catalog.Setup,
		catalogprivateendpoint.Setup,
		connection.Setup,
		dataasset.Setup,
		metastore.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	catalog "github.com/oracle/provider-oci/internal/controller/cluster/datacatalog/catalog"
	catalogprivateendpoint "github.com/oracle/provider-oci/internal/controller/cluster/datacatalog/catalogprivateendpoint"
	connection "github.com/oracle/provider-oci/internal/controller/cluster/datacatalog/connection"
	dataasset "github.com/oracle/provider-oci/internal/controller/cluster/datacatalog/dataasset"
	metastore "github.com/oracle/provider-oci/internal/controller/cluster/datacatalog/metastore"
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

// SetupGated_datacatalog creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_datacatalog(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		catalog.SetupGated,
		catalogprivateendpoint.SetupGated,
		connection.SetupGated,
		dataasset.SetupGated,
		metastore.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

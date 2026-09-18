/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	awrhub "github.com/oracle/provider-oci/internal/controller/cluster/opsi/awrhub"
	awrhubsource "github.com/oracle/provider-oci/internal/controller/cluster/opsi/awrhubsource"
	awrhubsourceawrhubsourcesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/opsi/awrhubsourceawrhubsourcesmanagement"
	chargebackplan "github.com/oracle/provider-oci/internal/controller/cluster/opsi/chargebackplan"
	databaseinsight "github.com/oracle/provider-oci/internal/controller/cluster/opsi/databaseinsight"
	enterprisemanagerbridge "github.com/oracle/provider-oci/internal/controller/cluster/opsi/enterprisemanagerbridge"
	exadatainsight "github.com/oracle/provider-oci/internal/controller/cluster/opsi/exadatainsight"
	hostinsight "github.com/oracle/provider-oci/internal/controller/cluster/opsi/hostinsight"
	newsreport "github.com/oracle/provider-oci/internal/controller/cluster/opsi/newsreport"
	operationsinsightsprivateendpoint "github.com/oracle/provider-oci/internal/controller/cluster/opsi/operationsinsightsprivateendpoint"
	operationsinsightswarehouse "github.com/oracle/provider-oci/internal/controller/cluster/opsi/operationsinsightswarehouse"
	operationsinsightswarehousedownloadwarehousewallet "github.com/oracle/provider-oci/internal/controller/cluster/opsi/operationsinsightswarehousedownloadwarehousewallet"
	operationsinsightswarehouserotatewarehousewallet "github.com/oracle/provider-oci/internal/controller/cluster/opsi/operationsinsightswarehouserotatewarehousewallet"
	operationsinsightswarehouseuser "github.com/oracle/provider-oci/internal/controller/cluster/opsi/operationsinsightswarehouseuser"
	opsiconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/opsi/opsiconfiguration"
)

// Setup_opsi creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_opsi(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		awrhub.Setup,
		awrhubsource.Setup,
		awrhubsourceawrhubsourcesmanagement.Setup,
		chargebackplan.Setup,
		databaseinsight.Setup,
		enterprisemanagerbridge.Setup,
		exadatainsight.Setup,
		hostinsight.Setup,
		newsreport.Setup,
		operationsinsightsprivateendpoint.Setup,
		operationsinsightswarehouse.Setup,
		operationsinsightswarehousedownloadwarehousewallet.Setup,
		operationsinsightswarehouserotatewarehousewallet.Setup,
		operationsinsightswarehouseuser.Setup,
		opsiconfiguration.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_opsi creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_opsi(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		awrhub.SetupGated,
		awrhubsource.SetupGated,
		awrhubsourceawrhubsourcesmanagement.SetupGated,
		chargebackplan.SetupGated,
		databaseinsight.SetupGated,
		enterprisemanagerbridge.SetupGated,
		exadatainsight.SetupGated,
		hostinsight.SetupGated,
		newsreport.SetupGated,
		operationsinsightsprivateendpoint.SetupGated,
		operationsinsightswarehouse.SetupGated,
		operationsinsightswarehousedownloadwarehousewallet.SetupGated,
		operationsinsightswarehouserotatewarehousewallet.SetupGated,
		operationsinsightswarehouseuser.SetupGated,
		opsiconfiguration.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

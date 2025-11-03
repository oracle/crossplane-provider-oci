/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	awrhub "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/awrhub"
	awrhubsource "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/awrhubsource"
	awrhubsourceawrhubsourcesmanagement "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/awrhubsourceawrhubsourcesmanagement"
	databaseinsight "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/databaseinsight"
	enterprisemanagerbridge "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/enterprisemanagerbridge"
	exadatainsight "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/exadatainsight"
	hostinsight "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/hostinsight"
	newsreport "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/newsreport"
	operationsinsightsprivateendpoint "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/operationsinsightsprivateendpoint"
	operationsinsightswarehouse "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/operationsinsightswarehouse"
	operationsinsightswarehousedownloadwarehousewallet "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/operationsinsightswarehousedownloadwarehousewallet"
	operationsinsightswarehouserotatewarehousewallet "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/operationsinsightswarehouserotatewarehousewallet"
	operationsinsightswarehouseuser "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/operationsinsightswarehouseuser"
	operatoraccesscontrolconfiguration "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/operatoraccesscontrolconfiguration"
)

// Setup_operatoraccesscontrol creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_operatoraccesscontrol(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		awrhub.Setup,
		awrhubsource.Setup,
		awrhubsourceawrhubsourcesmanagement.Setup,
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
		operatoraccesscontrolconfiguration.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

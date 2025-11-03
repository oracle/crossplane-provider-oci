/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	monitoringbaselineablemetric "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringbaselineablemetric"
	monitoringconfig "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringconfig"
	monitoringdiscoveryjob "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringdiscoveryjob"
	monitoringmaintenancewindow "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringmaintenancewindow"
	monitoringmaintenancewindowsretryfailedoperation "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringmaintenancewindowsretryfailedoperation"
	monitoringmaintenancewindowsstop "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringmaintenancewindowsstop"
	monitoringmetricextension "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringmetricextension"
	monitoringmetricextensionmetricextensionongivenresourcesmanagement "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringmetricextensionmetricextensionongivenresourcesmanagement"
	monitoringmonitoredresource "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringmonitoredresource"
	monitoringmonitoredresourcesassociatemonitoredresource "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringmonitoredresourcesassociatemonitoredresource"
	monitoringmonitoredresourceslistmember "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringmonitoredresourceslistmember"
	monitoringmonitoredresourcessearch "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringmonitoredresourcessearch"
	monitoringmonitoredresourcessearchassociation "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringmonitoredresourcessearchassociation"
	monitoringmonitoredresourcetask "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringmonitoredresourcetask"
	monitoringmonitoredresourcetype "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringmonitoredresourcetype"
	monitoringprocessset "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringprocessset"
	monitoringtemplate "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringtemplate"
	monitoringtemplatealarmcondition "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringtemplatealarmcondition"
	monitoringtemplatemonitoringtemplateongivenresourcesmanagement "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringtemplatemonitoringtemplateongivenresourcesmanagement"
)

// Setup_stackmonitoring creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_stackmonitoring(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		monitoringbaselineablemetric.Setup,
		monitoringconfig.Setup,
		monitoringdiscoveryjob.Setup,
		monitoringmaintenancewindow.Setup,
		monitoringmaintenancewindowsretryfailedoperation.Setup,
		monitoringmaintenancewindowsstop.Setup,
		monitoringmetricextension.Setup,
		monitoringmetricextensionmetricextensionongivenresourcesmanagement.Setup,
		monitoringmonitoredresource.Setup,
		monitoringmonitoredresourcesassociatemonitoredresource.Setup,
		monitoringmonitoredresourceslistmember.Setup,
		monitoringmonitoredresourcessearch.Setup,
		monitoringmonitoredresourcessearchassociation.Setup,
		monitoringmonitoredresourcetask.Setup,
		monitoringmonitoredresourcetype.Setup,
		monitoringprocessset.Setup,
		monitoringtemplate.Setup,
		monitoringtemplatealarmcondition.Setup,
		monitoringtemplatemonitoringtemplateongivenresourcesmanagement.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

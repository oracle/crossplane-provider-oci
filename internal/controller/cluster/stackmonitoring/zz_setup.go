/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	baselineablemetric "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/baselineablemetric"
	config "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/config"
	discoveryjob "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/discoveryjob"
	maintenancewindow "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/maintenancewindow"
	maintenancewindowsretryfailedoperation "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/maintenancewindowsretryfailedoperation"
	maintenancewindowsstop "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/maintenancewindowsstop"
	metricextension "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/metricextension"
	metricextensionmetricextensionongivenresourcesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/metricextensionmetricextensionongivenresourcesmanagement"
	monitoredresource "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/monitoredresource"
	monitoredresourcesassociatemonitoredresource "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/monitoredresourcesassociatemonitoredresource"
	monitoredresourceslistmember "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/monitoredresourceslistmember"
	monitoredresourcessearch "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/monitoredresourcessearch"
	monitoredresourcessearchassociation "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/monitoredresourcessearchassociation"
	monitoredresourcetask "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/monitoredresourcetask"
	monitoredresourcetype "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/monitoredresourcetype"
	monitoringtemplate "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/monitoringtemplate"
	monitoringtemplatealarmcondition "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/monitoringtemplatealarmcondition"
	monitoringtemplateongivenresourcesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/monitoringtemplateongivenresourcesmanagement"
	processset "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/processset"
)

// Setup_stackmonitoring creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_stackmonitoring(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		baselineablemetric.Setup,
		config.Setup,
		discoveryjob.Setup,
		maintenancewindow.Setup,
		maintenancewindowsretryfailedoperation.Setup,
		maintenancewindowsstop.Setup,
		metricextension.Setup,
		metricextensionmetricextensionongivenresourcesmanagement.Setup,
		monitoredresource.Setup,
		monitoredresourcesassociatemonitoredresource.Setup,
		monitoredresourceslistmember.Setup,
		monitoredresourcessearch.Setup,
		monitoredresourcessearchassociation.Setup,
		monitoredresourcetask.Setup,
		monitoredresourcetype.Setup,
		monitoringtemplate.Setup,
		monitoringtemplatealarmcondition.Setup,
		monitoringtemplateongivenresourcesmanagement.Setup,
		processset.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_stackmonitoring creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_stackmonitoring(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		baselineablemetric.SetupGated,
		config.SetupGated,
		discoveryjob.SetupGated,
		maintenancewindow.SetupGated,
		maintenancewindowsretryfailedoperation.SetupGated,
		maintenancewindowsstop.SetupGated,
		metricextension.SetupGated,
		metricextensionmetricextensionongivenresourcesmanagement.SetupGated,
		monitoredresource.SetupGated,
		monitoredresourcesassociatemonitoredresource.SetupGated,
		monitoredresourceslistmember.SetupGated,
		monitoredresourcessearch.SetupGated,
		monitoredresourcessearchassociation.SetupGated,
		monitoredresourcetask.SetupGated,
		monitoredresourcetype.SetupGated,
		monitoringtemplate.SetupGated,
		monitoringtemplatealarmcondition.SetupGated,
		monitoringtemplateongivenresourcesmanagement.SetupGated,
		processset.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

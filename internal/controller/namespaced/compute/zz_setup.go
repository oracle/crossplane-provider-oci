/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	appcataloglistingresourceversionagreement "github.com/oracle/provider-oci/internal/controller/namespaced/compute/appcataloglistingresourceversionagreement"
	appcatalogsubscription "github.com/oracle/provider-oci/internal/controller/namespaced/compute/appcatalogsubscription"
	clusternetwork "github.com/oracle/provider-oci/internal/controller/namespaced/compute/clusternetwork"
	computecapacityreport "github.com/oracle/provider-oci/internal/controller/namespaced/compute/computecapacityreport"
	computecapacityreservation "github.com/oracle/provider-oci/internal/controller/namespaced/compute/computecapacityreservation"
	computecapacitytopology "github.com/oracle/provider-oci/internal/controller/namespaced/compute/computecapacitytopology"
	computecluster "github.com/oracle/provider-oci/internal/controller/namespaced/compute/computecluster"
	computegpumemorycluster "github.com/oracle/provider-oci/internal/controller/namespaced/compute/computegpumemorycluster"
	computegpumemoryfabric "github.com/oracle/provider-oci/internal/controller/namespaced/compute/computegpumemoryfabric"
	computehost "github.com/oracle/provider-oci/internal/controller/namespaced/compute/computehost"
	computehostgroup "github.com/oracle/provider-oci/internal/controller/namespaced/compute/computehostgroup"
	computeimagecapabilityschema "github.com/oracle/provider-oci/internal/controller/namespaced/compute/computeimagecapabilityschema"
	consolehistory "github.com/oracle/provider-oci/internal/controller/namespaced/compute/consolehistory"
	dedicatedvmhost "github.com/oracle/provider-oci/internal/controller/namespaced/compute/dedicatedvmhost"
	image "github.com/oracle/provider-oci/internal/controller/namespaced/compute/image"
	instance "github.com/oracle/provider-oci/internal/controller/namespaced/compute/instance"
	instanceconfiguration "github.com/oracle/provider-oci/internal/controller/namespaced/compute/instanceconfiguration"
	instanceconsoleconnection "github.com/oracle/provider-oci/internal/controller/namespaced/compute/instanceconsoleconnection"
	instancemaintenanceevent "github.com/oracle/provider-oci/internal/controller/namespaced/compute/instancemaintenanceevent"
	instancepool "github.com/oracle/provider-oci/internal/controller/namespaced/compute/instancepool"
	instancepoolinstance "github.com/oracle/provider-oci/internal/controller/namespaced/compute/instancepoolinstance"
	shapemanagement "github.com/oracle/provider-oci/internal/controller/namespaced/compute/shapemanagement"
)

// Setup_compute creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_compute(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		appcataloglistingresourceversionagreement.Setup,
		appcatalogsubscription.Setup,
		clusternetwork.Setup,
		computecapacityreport.Setup,
		computecapacityreservation.Setup,
		computecapacitytopology.Setup,
		computecluster.Setup,
		computegpumemorycluster.Setup,
		computegpumemoryfabric.Setup,
		computehost.Setup,
		computehostgroup.Setup,
		computeimagecapabilityschema.Setup,
		consolehistory.Setup,
		dedicatedvmhost.Setup,
		image.Setup,
		instance.Setup,
		instanceconfiguration.Setup,
		instanceconsoleconnection.Setup,
		instancemaintenanceevent.Setup,
		instancepool.Setup,
		instancepoolinstance.Setup,
		shapemanagement.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_compute creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_compute(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		appcataloglistingresourceversionagreement.SetupGated,
		appcatalogsubscription.SetupGated,
		clusternetwork.SetupGated,
		computecapacityreport.SetupGated,
		computecapacityreservation.SetupGated,
		computecapacitytopology.SetupGated,
		computecluster.SetupGated,
		computegpumemorycluster.SetupGated,
		computegpumemoryfabric.SetupGated,
		computehost.SetupGated,
		computehostgroup.SetupGated,
		computeimagecapabilityschema.SetupGated,
		consolehistory.SetupGated,
		dedicatedvmhost.SetupGated,
		image.SetupGated,
		instance.SetupGated,
		instanceconfiguration.SetupGated,
		instanceconsoleconnection.SetupGated,
		instancemaintenanceevent.SetupGated,
		instancepool.SetupGated,
		instancepoolinstance.SetupGated,
		shapemanagement.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

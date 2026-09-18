/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	appcataloglistingresourceversionagreement "github.com/oracle/provider-oci/internal/controller/cluster/compute/appcataloglistingresourceversionagreement"
	appcatalogsubscription "github.com/oracle/provider-oci/internal/controller/cluster/compute/appcatalogsubscription"
	clusternetwork "github.com/oracle/provider-oci/internal/controller/cluster/compute/clusternetwork"
	computecapacityreport "github.com/oracle/provider-oci/internal/controller/cluster/compute/computecapacityreport"
	computecapacityreservation "github.com/oracle/provider-oci/internal/controller/cluster/compute/computecapacityreservation"
	computecapacitytopology "github.com/oracle/provider-oci/internal/controller/cluster/compute/computecapacitytopology"
	computecluster "github.com/oracle/provider-oci/internal/controller/cluster/compute/computecluster"
	computegpumemorycluster "github.com/oracle/provider-oci/internal/controller/cluster/compute/computegpumemorycluster"
	computegpumemoryfabric "github.com/oracle/provider-oci/internal/controller/cluster/compute/computegpumemoryfabric"
	computehost "github.com/oracle/provider-oci/internal/controller/cluster/compute/computehost"
	computehostgroup "github.com/oracle/provider-oci/internal/controller/cluster/compute/computehostgroup"
	computeimagecapabilityschema "github.com/oracle/provider-oci/internal/controller/cluster/compute/computeimagecapabilityschema"
	consolehistory "github.com/oracle/provider-oci/internal/controller/cluster/compute/consolehistory"
	dedicatedvmhost "github.com/oracle/provider-oci/internal/controller/cluster/compute/dedicatedvmhost"
	image "github.com/oracle/provider-oci/internal/controller/cluster/compute/image"
	instance "github.com/oracle/provider-oci/internal/controller/cluster/compute/instance"
	instanceconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/compute/instanceconfiguration"
	instanceconsoleconnection "github.com/oracle/provider-oci/internal/controller/cluster/compute/instanceconsoleconnection"
	instancemaintenanceevent "github.com/oracle/provider-oci/internal/controller/cluster/compute/instancemaintenanceevent"
	instancepool "github.com/oracle/provider-oci/internal/controller/cluster/compute/instancepool"
	instancepoolinstance "github.com/oracle/provider-oci/internal/controller/cluster/compute/instancepoolinstance"
	shapemanagement "github.com/oracle/provider-oci/internal/controller/cluster/compute/shapemanagement"
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

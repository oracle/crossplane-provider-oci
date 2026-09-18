/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	addon "github.com/oracle/provider-oci/internal/controller/namespaced/containerengine/addon"
	cluster "github.com/oracle/provider-oci/internal/controller/namespaced/containerengine/cluster"
	clustercompletecredentialrotationmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/containerengine/clustercompletecredentialrotationmanagement"
	clusterpublicapiendpointdecommissionmanager "github.com/oracle/provider-oci/internal/controller/namespaced/containerengine/clusterpublicapiendpointdecommissionmanager"
	clusterstartcredentialrotationmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/containerengine/clusterstartcredentialrotationmanagement"
	clusterworkloadmapping "github.com/oracle/provider-oci/internal/controller/namespaced/containerengine/clusterworkloadmapping"
	nodepool "github.com/oracle/provider-oci/internal/controller/namespaced/containerengine/nodepool"
	virtualnodepool "github.com/oracle/provider-oci/internal/controller/namespaced/containerengine/virtualnodepool"
)

// Setup_containerengine creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_containerengine(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		addon.Setup,
		cluster.Setup,
		clustercompletecredentialrotationmanagement.Setup,
		clusterpublicapiendpointdecommissionmanager.Setup,
		clusterstartcredentialrotationmanagement.Setup,
		clusterworkloadmapping.Setup,
		nodepool.Setup,
		virtualnodepool.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_containerengine creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_containerengine(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		addon.SetupGated,
		cluster.SetupGated,
		clustercompletecredentialrotationmanagement.SetupGated,
		clusterpublicapiendpointdecommissionmanager.SetupGated,
		clusterstartcredentialrotationmanagement.SetupGated,
		clusterworkloadmapping.SetupGated,
		nodepool.SetupGated,
		virtualnodepool.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

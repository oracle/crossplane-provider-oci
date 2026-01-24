/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	addon "github.com/oracle/provider-oci/internal/controller/containerengine/addon"
	cluster "github.com/oracle/provider-oci/internal/controller/containerengine/cluster"
	clustercompletecredentialrotationmanagement "github.com/oracle/provider-oci/internal/controller/containerengine/clustercompletecredentialrotationmanagement"
	clusterstartcredentialrotationmanagement "github.com/oracle/provider-oci/internal/controller/containerengine/clusterstartcredentialrotationmanagement"
	clusterworkloadmapping "github.com/oracle/provider-oci/internal/controller/containerengine/clusterworkloadmapping"
	nodepool "github.com/oracle/provider-oci/internal/controller/containerengine/nodepool"
	virtualnodepool "github.com/oracle/provider-oci/internal/controller/containerengine/virtualnodepool"
)

// Setup_containerengine creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_containerengine(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		addon.Setup,
		cluster.Setup,
		clustercompletecredentialrotationmanagement.Setup,
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

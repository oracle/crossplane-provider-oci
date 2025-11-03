/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	clusterpipeline "github.com/oracle/provider-oci/internal/controller/opensearch/clusterpipeline"
	opensearchcluster "github.com/oracle/provider-oci/internal/controller/opensearch/opensearchcluster"
)

// Setup_opensearch creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_opensearch(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		clusterpipeline.Setup,
		opensearchcluster.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

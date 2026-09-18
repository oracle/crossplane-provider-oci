/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	clusterplacementgroup "github.com/oracle/provider-oci/internal/controller/namespaced/clusterplacementgroups/clusterplacementgroup"
	healthdiagnosisstore "github.com/oracle/provider-oci/internal/controller/namespaced/clusterplacementgroups/healthdiagnosisstore"
)

// Setup_clusterplacementgroups creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_clusterplacementgroups(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		clusterplacementgroup.Setup,
		healthdiagnosisstore.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_clusterplacementgroups creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_clusterplacementgroups(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		clusterplacementgroup.SetupGated,
		healthdiagnosisstore.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

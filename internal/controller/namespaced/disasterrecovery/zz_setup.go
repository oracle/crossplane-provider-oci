/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	automaticdrconfiguration "github.com/oracle/provider-oci/internal/controller/namespaced/disasterrecovery/automaticdrconfiguration"
	drplan "github.com/oracle/provider-oci/internal/controller/namespaced/disasterrecovery/drplan"
	drplanexecution "github.com/oracle/provider-oci/internal/controller/namespaced/disasterrecovery/drplanexecution"
	drprotectiongroup "github.com/oracle/provider-oci/internal/controller/namespaced/disasterrecovery/drprotectiongroup"
)

// Setup_disasterrecovery creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_disasterrecovery(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		automaticdrconfiguration.Setup,
		drplan.Setup,
		drplanexecution.Setup,
		drprotectiongroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_disasterrecovery creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_disasterrecovery(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		automaticdrconfiguration.SetupGated,
		drplan.SetupGated,
		drplanexecution.SetupGated,
		drprotectiongroup.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

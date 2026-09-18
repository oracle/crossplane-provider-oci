/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	delegationcontrol "github.com/oracle/provider-oci/internal/controller/namespaced/delegateaccesscontrol/delegationcontrol"
	delegationsubscription "github.com/oracle/provider-oci/internal/controller/namespaced/delegateaccesscontrol/delegationsubscription"
)

// Setup_delegateaccesscontrol creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_delegateaccesscontrol(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		delegationcontrol.Setup,
		delegationsubscription.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_delegateaccesscontrol creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_delegateaccesscontrol(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		delegationcontrol.SetupGated,
		delegationsubscription.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	webappacceleration "github.com/oracle/provider-oci/internal/controller/cluster/waa/webappacceleration"
	webappaccelerationpolicy "github.com/oracle/provider-oci/internal/controller/cluster/waa/webappaccelerationpolicy"
)

// Setup_waa creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_waa(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		webappacceleration.Setup,
		webappaccelerationpolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_waa creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_waa(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		webappacceleration.SetupGated,
		webappaccelerationpolicy.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

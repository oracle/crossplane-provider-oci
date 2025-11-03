/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	webappacceleration "github.com/oracle/provider-oci/internal/controller/waa/webappacceleration"
	webappaccelerationpolicy "github.com/oracle/provider-oci/internal/controller/waa/webappaccelerationpolicy"
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

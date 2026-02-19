/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	protecteddatabase "github.com/oracle/provider-oci/internal/controller/recovery/protecteddatabase"
	protectionpolicy "github.com/oracle/provider-oci/internal/controller/recovery/protectionpolicy"
	recoveryservicesubnet "github.com/oracle/provider-oci/internal/controller/recovery/recoveryservicesubnet"
)

// Setup_recovery creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_recovery(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		protecteddatabase.Setup,
		protectionpolicy.Setup,
		recoveryservicesubnet.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

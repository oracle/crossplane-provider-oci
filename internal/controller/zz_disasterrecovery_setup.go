/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	recoverydrplan "github.com/oracle/provider-oci/internal/controller/disasterrecovery/recoverydrplan"
	recoverydrplanexecution "github.com/oracle/provider-oci/internal/controller/disasterrecovery/recoverydrplanexecution"
	recoverydrprotectiongroup "github.com/oracle/provider-oci/internal/controller/disasterrecovery/recoverydrprotectiongroup"
)

// Setup_disasterrecovery creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_disasterrecovery(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		recoverydrplan.Setup,
		recoverydrplanexecution.Setup,
		recoverydrprotectiongroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

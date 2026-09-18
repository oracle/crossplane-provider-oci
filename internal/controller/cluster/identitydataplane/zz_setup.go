/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	generatescopedaccesstoken "github.com/oracle/provider-oci/internal/controller/cluster/identitydataplane/generatescopedaccesstoken"
)

// Setup_identitydataplane creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_identitydataplane(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		generatescopedaccesstoken.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_identitydataplane creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_identitydataplane(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		generatescopedaccesstoken.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

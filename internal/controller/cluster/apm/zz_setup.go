/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	apmdomain "github.com/oracle/provider-oci/internal/controller/cluster/apm/apmdomain"
)

// Setup_apm creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_apm(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		apmdomain.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_apm creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_apm(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		apmdomain.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

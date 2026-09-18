/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	gdppipeline "github.com/oracle/provider-oci/internal/controller/namespaced/gdp/gdppipeline"
)

// Setup_gdp creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_gdp(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		gdppipeline.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_gdp creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_gdp(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		gdppipeline.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	config "github.com/oracle/provider-oci/internal/controller/apmconfig/config"
)

// Setup_apmconfig creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_apmconfig(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		config.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

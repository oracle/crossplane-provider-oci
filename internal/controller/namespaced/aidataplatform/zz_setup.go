/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	aidataplatform "github.com/oracle/provider-oci/internal/controller/namespaced/aidataplatform/aidataplatform"
)

// Setup_aidataplatform creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_aidataplatform(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		aidataplatform.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_aidataplatform creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_aidataplatform(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		aidataplatform.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

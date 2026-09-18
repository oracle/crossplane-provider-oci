/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	securityattribute "github.com/oracle/provider-oci/internal/controller/namespaced/securityattribute/securityattribute"
	securityattributenamespace "github.com/oracle/provider-oci/internal/controller/namespaced/securityattribute/securityattributenamespace"
)

// Setup_securityattribute creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_securityattribute(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		securityattribute.Setup,
		securityattributenamespace.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_securityattribute creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_securityattribute(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		securityattribute.SetupGated,
		securityattributenamespace.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

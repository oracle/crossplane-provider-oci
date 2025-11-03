/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	attributesecurityattribute "github.com/oracle/provider-oci/internal/controller/securityattribute/attributesecurityattribute"
	attributesecurityattributenamespace "github.com/oracle/provider-oci/internal/controller/securityattribute/attributesecurityattributenamespace"
)

// Setup_securityattribute creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_securityattribute(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		attributesecurityattribute.Setup,
		attributesecurityattributenamespace.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

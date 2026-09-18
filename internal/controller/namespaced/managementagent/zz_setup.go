/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	managementagent "github.com/oracle/provider-oci/internal/controller/namespaced/managementagent/managementagent"
	managementagentdatasource "github.com/oracle/provider-oci/internal/controller/namespaced/managementagent/managementagentdatasource"
	managementagentinstallkey "github.com/oracle/provider-oci/internal/controller/namespaced/managementagent/managementagentinstallkey"
	namedcredential "github.com/oracle/provider-oci/internal/controller/namespaced/managementagent/namedcredential"
)

// Setup_managementagent creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_managementagent(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		managementagent.Setup,
		managementagentdatasource.Setup,
		managementagentinstallkey.Setup,
		namedcredential.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_managementagent creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_managementagent(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		managementagent.SetupGated,
		managementagentdatasource.SetupGated,
		managementagentinstallkey.SetupGated,
		namedcredential.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

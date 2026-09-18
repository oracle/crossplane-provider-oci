/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	fsucollection "github.com/oracle/provider-oci/internal/controller/namespaced/fleetsoftwareupdate/fsucollection"
	fsucycle "github.com/oracle/provider-oci/internal/controller/namespaced/fleetsoftwareupdate/fsucycle"
	fsureadinesscheck "github.com/oracle/provider-oci/internal/controller/namespaced/fleetsoftwareupdate/fsureadinesscheck"
)

// Setup_fleetsoftwareupdate creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_fleetsoftwareupdate(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		fsucollection.Setup,
		fsucycle.Setup,
		fsureadinesscheck.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_fleetsoftwareupdate creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_fleetsoftwareupdate(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		fsucollection.SetupGated,
		fsucycle.SetupGated,
		fsureadinesscheck.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

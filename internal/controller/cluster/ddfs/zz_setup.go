/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	instance "github.com/oracle/provider-oci/internal/controller/cluster/ddfs/instance"
)

// Setup_ddfs creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_ddfs(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		instance.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_ddfs creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_ddfs(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		instance.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

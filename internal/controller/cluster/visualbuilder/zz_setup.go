/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	vbinstance "github.com/oracle/provider-oci/internal/controller/cluster/visualbuilder/vbinstance"
)

// Setup_visualbuilder creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_visualbuilder(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		vbinstance.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_visualbuilder creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_visualbuilder(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		vbinstance.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

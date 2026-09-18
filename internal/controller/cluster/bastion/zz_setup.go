/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	bastion "github.com/oracle/provider-oci/internal/controller/cluster/bastion/bastion"
	session "github.com/oracle/provider-oci/internal/controller/cluster/bastion/session"
)

// Setup_bastion creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_bastion(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bastion.Setup,
		session.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_bastion creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_bastion(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bastion.SetupGated,
		session.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

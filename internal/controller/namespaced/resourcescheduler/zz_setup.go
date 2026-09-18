/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	schedule "github.com/oracle/provider-oci/internal/controller/namespaced/resourcescheduler/schedule"
)

// Setup_resourcescheduler creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_resourcescheduler(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		schedule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_resourcescheduler creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_resourcescheduler(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		schedule.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

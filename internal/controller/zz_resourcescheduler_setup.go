/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	schedule "github.com/oracle/provider-oci/internal/controller/resourcescheduler/schedule"
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

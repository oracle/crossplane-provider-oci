/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	alarm "github.com/oracle/provider-oci/internal/controller/namespaced/monitoring/alarm"
	alarmsuppression "github.com/oracle/provider-oci/internal/controller/namespaced/monitoring/alarmsuppression"
	capturefilter "github.com/oracle/provider-oci/internal/controller/namespaced/monitoring/capturefilter"
	vtap "github.com/oracle/provider-oci/internal/controller/namespaced/monitoring/vtap"
)

// Setup_monitoring creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_monitoring(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alarm.Setup,
		alarmsuppression.Setup,
		capturefilter.Setup,
		vtap.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_monitoring creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_monitoring(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alarm.SetupGated,
		alarmsuppression.SetupGated,
		capturefilter.SetupGated,
		vtap.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

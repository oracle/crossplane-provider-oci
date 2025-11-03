/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	alarm "github.com/oracle/provider-oci/internal/controller/monitoring/alarm"
	alarmsuppression "github.com/oracle/provider-oci/internal/controller/monitoring/alarmsuppression"
	capturefilter "github.com/oracle/provider-oci/internal/controller/monitoring/capturefilter"
	vtap "github.com/oracle/provider-oci/internal/controller/monitoring/vtap"
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

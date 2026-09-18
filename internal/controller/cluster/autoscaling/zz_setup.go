/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	autoscalingconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/autoscaling/autoscalingconfiguration"
)

// Setup_autoscaling creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_autoscaling(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		autoscalingconfiguration.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_autoscaling creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_autoscaling(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		autoscalingconfiguration.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

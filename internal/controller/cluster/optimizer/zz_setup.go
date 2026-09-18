/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	enrollmentstatus "github.com/oracle/provider-oci/internal/controller/cluster/optimizer/enrollmentstatus"
	profile "github.com/oracle/provider-oci/internal/controller/cluster/optimizer/profile"
	recommendation "github.com/oracle/provider-oci/internal/controller/cluster/optimizer/recommendation"
	resourceaction "github.com/oracle/provider-oci/internal/controller/cluster/optimizer/resourceaction"
)

// Setup_optimizer creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_optimizer(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		enrollmentstatus.Setup,
		profile.Setup,
		recommendation.Setup,
		resourceaction.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_optimizer creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_optimizer(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		enrollmentstatus.SetupGated,
		profile.SetupGated,
		recommendation.SetupGated,
		resourceaction.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

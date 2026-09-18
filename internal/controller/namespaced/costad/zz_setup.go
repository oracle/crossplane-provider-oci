/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	costalertsubscription "github.com/oracle/provider-oci/internal/controller/namespaced/costad/costalertsubscription"
	costanomalyevent "github.com/oracle/provider-oci/internal/controller/namespaced/costad/costanomalyevent"
	costanomalymonitor "github.com/oracle/provider-oci/internal/controller/namespaced/costad/costanomalymonitor"
	costanomalymonitorcostanomalymonitorenabletogglesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/costad/costanomalymonitorcostanomalymonitorenabletogglesmanagement"
)

// Setup_costad creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_costad(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		costalertsubscription.Setup,
		costanomalyevent.Setup,
		costanomalymonitor.Setup,
		costanomalymonitorcostanomalymonitorenabletogglesmanagement.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_costad creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_costad(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		costalertsubscription.SetupGated,
		costanomalyevent.SetupGated,
		costanomalymonitor.SetupGated,
		costanomalymonitorcostanomalymonitorenabletogglesmanagement.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

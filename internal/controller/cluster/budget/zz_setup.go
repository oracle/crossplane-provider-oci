/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	alertrule "github.com/oracle/provider-oci/internal/controller/cluster/budget/alertrule"
	budget "github.com/oracle/provider-oci/internal/controller/cluster/budget/budget"
	costalertsubscription "github.com/oracle/provider-oci/internal/controller/cluster/budget/costalertsubscription"
	costanomalyevent "github.com/oracle/provider-oci/internal/controller/cluster/budget/costanomalyevent"
	costanomalymonitor "github.com/oracle/provider-oci/internal/controller/cluster/budget/costanomalymonitor"
	costanomalymonitorcostanomalymonitorenabletogglesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/budget/costanomalymonitorcostanomalymonitorenabletogglesmanagement"
)

// Setup_budget creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_budget(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alertrule.Setup,
		budget.Setup,
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

// SetupGated_budget creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_budget(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		alertrule.SetupGated,
		budget.SetupGated,
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

/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	internaloccmdemandsignal "github.com/oracle/provider-oci/internal/controller/cluster/capacitymanagement/internaloccmdemandsignal"
	internaloccmdemandsignaldelivery "github.com/oracle/provider-oci/internal/controller/cluster/capacitymanagement/internaloccmdemandsignaldelivery"
	occavailabilitycatalog "github.com/oracle/provider-oci/internal/controller/cluster/capacitymanagement/occavailabilitycatalog"
	occcapacityrequest "github.com/oracle/provider-oci/internal/controller/cluster/capacitymanagement/occcapacityrequest"
	occcustomergroup "github.com/oracle/provider-oci/internal/controller/cluster/capacitymanagement/occcustomergroup"
	occcustomergroupocccustomer "github.com/oracle/provider-oci/internal/controller/cluster/capacitymanagement/occcustomergroupocccustomer"
	occmdemandsignal "github.com/oracle/provider-oci/internal/controller/cluster/capacitymanagement/occmdemandsignal"
	occmdemandsignalitem "github.com/oracle/provider-oci/internal/controller/cluster/capacitymanagement/occmdemandsignalitem"
)

// Setup_capacitymanagement creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_capacitymanagement(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		internaloccmdemandsignal.Setup,
		internaloccmdemandsignaldelivery.Setup,
		occavailabilitycatalog.Setup,
		occcapacityrequest.Setup,
		occcustomergroup.Setup,
		occcustomergroupocccustomer.Setup,
		occmdemandsignal.Setup,
		occmdemandsignalitem.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_capacitymanagement creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_capacitymanagement(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		internaloccmdemandsignal.SetupGated,
		internaloccmdemandsignaldelivery.SetupGated,
		occavailabilitycatalog.SetupGated,
		occcapacityrequest.SetupGated,
		occcustomergroup.SetupGated,
		occcustomergroupocccustomer.SetupGated,
		occmdemandsignal.SetupGated,
		occmdemandsignalitem.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

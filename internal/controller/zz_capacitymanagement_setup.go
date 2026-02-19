/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	internaloccmdemandsignal "github.com/oracle/provider-oci/internal/controller/capacitymanagement/internaloccmdemandsignal"
	internaloccmdemandsignaldelivery "github.com/oracle/provider-oci/internal/controller/capacitymanagement/internaloccmdemandsignaldelivery"
	occavailabilitycatalog "github.com/oracle/provider-oci/internal/controller/capacitymanagement/occavailabilitycatalog"
	occcapacityrequest "github.com/oracle/provider-oci/internal/controller/capacitymanagement/occcapacityrequest"
	occcustomergroup "github.com/oracle/provider-oci/internal/controller/capacitymanagement/occcustomergroup"
	occcustomergroupocccustomer "github.com/oracle/provider-oci/internal/controller/capacitymanagement/occcustomergroupocccustomer"
	occmdemandsignal "github.com/oracle/provider-oci/internal/controller/capacitymanagement/occmdemandsignal"
	occmdemandsignalitem "github.com/oracle/provider-oci/internal/controller/capacitymanagement/occmdemandsignalitem"
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

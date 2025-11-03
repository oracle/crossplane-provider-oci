/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	managementinternaloccmdemandsignal "github.com/oracle/provider-oci/internal/controller/capacity/managementinternaloccmdemandsignal"
	managementinternaloccmdemandsignaldelivery "github.com/oracle/provider-oci/internal/controller/capacity/managementinternaloccmdemandsignaldelivery"
	managementoccavailabilitycatalog "github.com/oracle/provider-oci/internal/controller/capacity/managementoccavailabilitycatalog"
	managementocccapacityrequest "github.com/oracle/provider-oci/internal/controller/capacity/managementocccapacityrequest"
	managementocccustomergroup "github.com/oracle/provider-oci/internal/controller/capacity/managementocccustomergroup"
	managementocccustomergroupocccustomer "github.com/oracle/provider-oci/internal/controller/capacity/managementocccustomergroupocccustomer"
	managementoccmdemandsignal "github.com/oracle/provider-oci/internal/controller/capacity/managementoccmdemandsignal"
	managementoccmdemandsignalitem "github.com/oracle/provider-oci/internal/controller/capacity/managementoccmdemandsignalitem"
)

// Setup_capacity creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_capacity(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		managementinternaloccmdemandsignal.Setup,
		managementinternaloccmdemandsignaldelivery.Setup,
		managementoccavailabilitycatalog.Setup,
		managementocccapacityrequest.Setup,
		managementocccustomergroup.Setup,
		managementocccustomergroupocccustomer.Setup,
		managementoccmdemandsignal.Setup,
		managementoccmdemandsignalitem.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	actioncreatezonefromzonefile "github.com/oracle/provider-oci/internal/controller/cluster/dns/actioncreatezonefromzonefile"
	record "github.com/oracle/provider-oci/internal/controller/cluster/dns/record"
	resolver "github.com/oracle/provider-oci/internal/controller/cluster/dns/resolver"
	resolverendpoint "github.com/oracle/provider-oci/internal/controller/cluster/dns/resolverendpoint"
	rrset "github.com/oracle/provider-oci/internal/controller/cluster/dns/rrset"
	steeringpolicy "github.com/oracle/provider-oci/internal/controller/cluster/dns/steeringpolicy"
	steeringpolicyattachment "github.com/oracle/provider-oci/internal/controller/cluster/dns/steeringpolicyattachment"
	tsigkey "github.com/oracle/provider-oci/internal/controller/cluster/dns/tsigkey"
	view "github.com/oracle/provider-oci/internal/controller/cluster/dns/view"
	zone "github.com/oracle/provider-oci/internal/controller/cluster/dns/zone"
	zonepromotednsseckeyversion "github.com/oracle/provider-oci/internal/controller/cluster/dns/zonepromotednsseckeyversion"
	zonestagednsseckeyversion "github.com/oracle/provider-oci/internal/controller/cluster/dns/zonestagednsseckeyversion"
)

// Setup_dns creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dns(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		actioncreatezonefromzonefile.Setup,
		record.Setup,
		resolver.Setup,
		resolverendpoint.Setup,
		rrset.Setup,
		steeringpolicy.Setup,
		steeringpolicyattachment.Setup,
		tsigkey.Setup,
		view.Setup,
		zone.Setup,
		zonepromotednsseckeyversion.Setup,
		zonestagednsseckeyversion.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_dns creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_dns(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		actioncreatezonefromzonefile.SetupGated,
		record.SetupGated,
		resolver.SetupGated,
		resolverendpoint.SetupGated,
		rrset.SetupGated,
		steeringpolicy.SetupGated,
		steeringpolicyattachment.SetupGated,
		tsigkey.SetupGated,
		view.SetupGated,
		zone.SetupGated,
		zonepromotednsseckeyversion.SetupGated,
		zonestagednsseckeyversion.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

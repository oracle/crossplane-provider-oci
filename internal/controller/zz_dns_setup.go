/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	actioncreatezonefromzonefile "github.com/oracle/provider-oci/internal/controller/dns/actioncreatezonefromzonefile"
	record "github.com/oracle/provider-oci/internal/controller/dns/record"
	resolver "github.com/oracle/provider-oci/internal/controller/dns/resolver"
	resolverendpoint "github.com/oracle/provider-oci/internal/controller/dns/resolverendpoint"
	rrset "github.com/oracle/provider-oci/internal/controller/dns/rrset"
	steeringpolicy "github.com/oracle/provider-oci/internal/controller/dns/steeringpolicy"
	steeringpolicyattachment "github.com/oracle/provider-oci/internal/controller/dns/steeringpolicyattachment"
	tsigkey "github.com/oracle/provider-oci/internal/controller/dns/tsigkey"
	view "github.com/oracle/provider-oci/internal/controller/dns/view"
	zone "github.com/oracle/provider-oci/internal/controller/dns/zone"
	zonepromotednsseckeyversion "github.com/oracle/provider-oci/internal/controller/dns/zonepromotednsseckeyversion"
	zonestagednsseckeyversion "github.com/oracle/provider-oci/internal/controller/dns/zonestagednsseckeyversion"
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

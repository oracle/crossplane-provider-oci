/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	adhocquery "github.com/oracle/provider-oci/internal/controller/cluster/cloudguard/adhocquery"
	cloudguardconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/cloudguard/cloudguardconfiguration"
	datamaskrule "github.com/oracle/provider-oci/internal/controller/cluster/cloudguard/datamaskrule"
	datasource "github.com/oracle/provider-oci/internal/controller/cluster/cloudguard/datasource"
	detectorrecipe "github.com/oracle/provider-oci/internal/controller/cluster/cloudguard/detectorrecipe"
	managedlist "github.com/oracle/provider-oci/internal/controller/cluster/cloudguard/managedlist"
	responderrecipe "github.com/oracle/provider-oci/internal/controller/cluster/cloudguard/responderrecipe"
	savedquery "github.com/oracle/provider-oci/internal/controller/cluster/cloudguard/savedquery"
	securityrecipe "github.com/oracle/provider-oci/internal/controller/cluster/cloudguard/securityrecipe"
	securityzone "github.com/oracle/provider-oci/internal/controller/cluster/cloudguard/securityzone"
	target "github.com/oracle/provider-oci/internal/controller/cluster/cloudguard/target"
	wlpagent "github.com/oracle/provider-oci/internal/controller/cluster/cloudguard/wlpagent"
)

// Setup_cloudguard creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cloudguard(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		adhocquery.Setup,
		cloudguardconfiguration.Setup,
		datamaskrule.Setup,
		datasource.Setup,
		detectorrecipe.Setup,
		managedlist.Setup,
		responderrecipe.Setup,
		savedquery.Setup,
		securityrecipe.Setup,
		securityzone.Setup,
		target.Setup,
		wlpagent.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_cloudguard creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_cloudguard(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		adhocquery.SetupGated,
		cloudguardconfiguration.SetupGated,
		datamaskrule.SetupGated,
		datasource.SetupGated,
		detectorrecipe.SetupGated,
		managedlist.SetupGated,
		responderrecipe.SetupGated,
		savedquery.SetupGated,
		securityrecipe.SetupGated,
		securityzone.SetupGated,
		target.SetupGated,
		wlpagent.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	addresslist "github.com/oracle/provider-oci/internal/controller/cluster/waas/addresslist"
	certificate "github.com/oracle/provider-oci/internal/controller/cluster/waas/certificate"
	customprotectionrule "github.com/oracle/provider-oci/internal/controller/cluster/waas/customprotectionrule"
	httpredirect "github.com/oracle/provider-oci/internal/controller/cluster/waas/httpredirect"
	protectionrule "github.com/oracle/provider-oci/internal/controller/cluster/waas/protectionrule"
	purgecache "github.com/oracle/provider-oci/internal/controller/cluster/waas/purgecache"
	waaspolicy "github.com/oracle/provider-oci/internal/controller/cluster/waas/waaspolicy"
)

// Setup_waas creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_waas(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		addresslist.Setup,
		certificate.Setup,
		customprotectionrule.Setup,
		httpredirect.Setup,
		protectionrule.Setup,
		purgecache.Setup,
		waaspolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_waas creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_waas(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		addresslist.SetupGated,
		certificate.SetupGated,
		customprotectionrule.SetupGated,
		httpredirect.SetupGated,
		protectionrule.SetupGated,
		purgecache.SetupGated,
		waaspolicy.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

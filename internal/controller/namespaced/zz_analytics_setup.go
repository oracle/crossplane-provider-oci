/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	analyticsinstance "github.com/oracle/provider-oci/internal/controller/namespaced/analytics/analyticsinstance"
	analyticsinstanceprivateaccesschannel "github.com/oracle/provider-oci/internal/controller/namespaced/analytics/analyticsinstanceprivateaccesschannel"
	analyticsinstanceresourcegroup "github.com/oracle/provider-oci/internal/controller/namespaced/analytics/analyticsinstanceresourcegroup"
	analyticsinstancevanityurl "github.com/oracle/provider-oci/internal/controller/namespaced/analytics/analyticsinstancevanityurl"
)

// Setup_analytics creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_analytics(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		analyticsinstance.Setup,
		analyticsinstanceprivateaccesschannel.Setup,
		analyticsinstanceresourcegroup.Setup,
		analyticsinstancevanityurl.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_analytics creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_analytics(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		analyticsinstance.SetupGated,
		analyticsinstanceprivateaccesschannel.SetupGated,
		analyticsinstanceresourcegroup.SetupGated,
		analyticsinstancevanityurl.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	analyticsinstance "github.com/oracle/provider-oci/internal/controller/analytics/analyticsinstance"
	analyticsinstanceprivateaccesschannel "github.com/oracle/provider-oci/internal/controller/analytics/analyticsinstanceprivateaccesschannel"
	analyticsinstancevanityurl "github.com/oracle/provider-oci/internal/controller/analytics/analyticsinstancevanityurl"
)

// Setup_analytics creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_analytics(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		analyticsinstance.Setup,
		analyticsinstanceprivateaccesschannel.Setup,
		analyticsinstancevanityurl.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

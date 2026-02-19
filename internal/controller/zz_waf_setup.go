/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	networkaddresslist "github.com/oracle/provider-oci/internal/controller/waf/networkaddresslist"
	webappfirewall "github.com/oracle/provider-oci/internal/controller/waf/webappfirewall"
	webappfirewallpolicy "github.com/oracle/provider-oci/internal/controller/waf/webappfirewallpolicy"
)

// Setup_waf creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_waf(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		networkaddresslist.Setup,
		webappfirewall.Setup,
		webappfirewallpolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

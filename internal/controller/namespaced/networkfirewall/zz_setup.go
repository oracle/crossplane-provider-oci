/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	networkfirewall "github.com/oracle/provider-oci/internal/controller/namespaced/networkfirewall/networkfirewall"
	networkfirewallpolicy "github.com/oracle/provider-oci/internal/controller/namespaced/networkfirewall/networkfirewallpolicy"
	networkfirewallpolicyaddresslist "github.com/oracle/provider-oci/internal/controller/namespaced/networkfirewall/networkfirewallpolicyaddresslist"
	networkfirewallpolicyapplication "github.com/oracle/provider-oci/internal/controller/namespaced/networkfirewall/networkfirewallpolicyapplication"
	networkfirewallpolicyapplicationgroup "github.com/oracle/provider-oci/internal/controller/namespaced/networkfirewall/networkfirewallpolicyapplicationgroup"
	networkfirewallpolicydecryptionprofile "github.com/oracle/provider-oci/internal/controller/namespaced/networkfirewall/networkfirewallpolicydecryptionprofile"
	networkfirewallpolicydecryptionrule "github.com/oracle/provider-oci/internal/controller/namespaced/networkfirewall/networkfirewallpolicydecryptionrule"
	networkfirewallpolicymappedsecret "github.com/oracle/provider-oci/internal/controller/namespaced/networkfirewall/networkfirewallpolicymappedsecret"
	networkfirewallpolicynatrule "github.com/oracle/provider-oci/internal/controller/namespaced/networkfirewall/networkfirewallpolicynatrule"
	networkfirewallpolicysecurityrule "github.com/oracle/provider-oci/internal/controller/namespaced/networkfirewall/networkfirewallpolicysecurityrule"
	networkfirewallpolicyservice "github.com/oracle/provider-oci/internal/controller/namespaced/networkfirewall/networkfirewallpolicyservice"
	networkfirewallpolicytunnelinspectionrule "github.com/oracle/provider-oci/internal/controller/namespaced/networkfirewall/networkfirewallpolicytunnelinspectionrule"
)

// Setup_networkfirewall creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_networkfirewall(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		networkfirewall.Setup,
		networkfirewallpolicy.Setup,
		networkfirewallpolicyaddresslist.Setup,
		networkfirewallpolicyapplication.Setup,
		networkfirewallpolicyapplicationgroup.Setup,
		networkfirewallpolicydecryptionprofile.Setup,
		networkfirewallpolicydecryptionrule.Setup,
		networkfirewallpolicymappedsecret.Setup,
		networkfirewallpolicynatrule.Setup,
		networkfirewallpolicysecurityrule.Setup,
		networkfirewallpolicyservice.Setup,
		networkfirewallpolicytunnelinspectionrule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_networkfirewall creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_networkfirewall(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		networkfirewall.SetupGated,
		networkfirewallpolicy.SetupGated,
		networkfirewallpolicyaddresslist.SetupGated,
		networkfirewallpolicyapplication.SetupGated,
		networkfirewallpolicyapplicationgroup.SetupGated,
		networkfirewallpolicydecryptionprofile.SetupGated,
		networkfirewallpolicydecryptionrule.SetupGated,
		networkfirewallpolicymappedsecret.SetupGated,
		networkfirewallpolicynatrule.SetupGated,
		networkfirewallpolicysecurityrule.SetupGated,
		networkfirewallpolicyservice.SetupGated,
		networkfirewallpolicytunnelinspectionrule.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

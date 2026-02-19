/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	networkfirewall "github.com/oracle/provider-oci/internal/controller/networkfirewall/networkfirewall"
	networkfirewallpolicy "github.com/oracle/provider-oci/internal/controller/networkfirewall/networkfirewallpolicy"
	networkfirewallpolicyaddresslist "github.com/oracle/provider-oci/internal/controller/networkfirewall/networkfirewallpolicyaddresslist"
	networkfirewallpolicyapplication "github.com/oracle/provider-oci/internal/controller/networkfirewall/networkfirewallpolicyapplication"
	networkfirewallpolicyapplicationgroup "github.com/oracle/provider-oci/internal/controller/networkfirewall/networkfirewallpolicyapplicationgroup"
	networkfirewallpolicydecryptionprofile "github.com/oracle/provider-oci/internal/controller/networkfirewall/networkfirewallpolicydecryptionprofile"
	networkfirewallpolicydecryptionrule "github.com/oracle/provider-oci/internal/controller/networkfirewall/networkfirewallpolicydecryptionrule"
	networkfirewallpolicymappedsecret "github.com/oracle/provider-oci/internal/controller/networkfirewall/networkfirewallpolicymappedsecret"
	networkfirewallpolicynatrule "github.com/oracle/provider-oci/internal/controller/networkfirewall/networkfirewallpolicynatrule"
	networkfirewallpolicysecurityrule "github.com/oracle/provider-oci/internal/controller/networkfirewall/networkfirewallpolicysecurityrule"
	networkfirewallpolicyservice "github.com/oracle/provider-oci/internal/controller/networkfirewall/networkfirewallpolicyservice"
	networkfirewallpolicytunnelinspectionrule "github.com/oracle/provider-oci/internal/controller/networkfirewall/networkfirewallpolicytunnelinspectionrule"
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

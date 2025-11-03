/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	fleet "github.com/oracle/provider-oci/internal/controller/jms/fleet"
	fleetadvancedfeatureconfiguration "github.com/oracle/provider-oci/internal/controller/jms/fleetadvancedfeatureconfiguration"
	javadownloadsjavadownloadreport "github.com/oracle/provider-oci/internal/controller/jms/javadownloadsjavadownloadreport"
	javadownloadsjavadownloadtoken "github.com/oracle/provider-oci/internal/controller/jms/javadownloadsjavadownloadtoken"
	javadownloadsjavalicenseacceptancerecord "github.com/oracle/provider-oci/internal/controller/jms/javadownloadsjavalicenseacceptancerecord"
	plugin "github.com/oracle/provider-oci/internal/controller/jms/plugin"
)

// Setup_jms creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_jms(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		fleet.Setup,
		fleetadvancedfeatureconfiguration.Setup,
		javadownloadsjavadownloadreport.Setup,
		javadownloadsjavadownloadtoken.Setup,
		javadownloadsjavalicenseacceptancerecord.Setup,
		plugin.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

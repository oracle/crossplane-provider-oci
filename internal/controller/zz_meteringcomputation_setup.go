/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	customtable "github.com/oracle/provider-oci/internal/controller/meteringcomputation/customtable"
	query "github.com/oracle/provider-oci/internal/controller/meteringcomputation/query"
	schedule "github.com/oracle/provider-oci/internal/controller/meteringcomputation/schedule"
	usage "github.com/oracle/provider-oci/internal/controller/meteringcomputation/usage"
	usagecarbonemission "github.com/oracle/provider-oci/internal/controller/meteringcomputation/usagecarbonemission"
	usagecarbonemissionsquery "github.com/oracle/provider-oci/internal/controller/meteringcomputation/usagecarbonemissionsquery"
	usagestatementemailrecipientsgroup "github.com/oracle/provider-oci/internal/controller/meteringcomputation/usagestatementemailrecipientsgroup"
)

// Setup_meteringcomputation creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_meteringcomputation(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		customtable.Setup,
		query.Setup,
		schedule.Setup,
		usage.Setup,
		usagecarbonemission.Setup,
		usagecarbonemissionsquery.Setup,
		usagestatementemailrecipientsgroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

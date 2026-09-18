/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	customtable "github.com/oracle/provider-oci/internal/controller/cluster/meteringcomputation/customtable"
	query "github.com/oracle/provider-oci/internal/controller/cluster/meteringcomputation/query"
	schedule "github.com/oracle/provider-oci/internal/controller/cluster/meteringcomputation/schedule"
	usage "github.com/oracle/provider-oci/internal/controller/cluster/meteringcomputation/usage"
	usagecarbonemission "github.com/oracle/provider-oci/internal/controller/cluster/meteringcomputation/usagecarbonemission"
	usagecarbonemissionsquery "github.com/oracle/provider-oci/internal/controller/cluster/meteringcomputation/usagecarbonemissionsquery"
	usagestatementemailrecipientsgroup "github.com/oracle/provider-oci/internal/controller/cluster/meteringcomputation/usagestatementemailrecipientsgroup"
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

// SetupGated_meteringcomputation creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_meteringcomputation(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		customtable.SetupGated,
		query.SetupGated,
		schedule.SetupGated,
		usage.SetupGated,
		usagecarbonemission.SetupGated,
		usagecarbonemissionsquery.SetupGated,
		usagestatementemailrecipientsgroup.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	computationcustomtable "github.com/oracle/provider-oci/internal/controller/meteringcomputation/computationcustomtable"
	computationquery "github.com/oracle/provider-oci/internal/controller/meteringcomputation/computationquery"
	computationschedule "github.com/oracle/provider-oci/internal/controller/meteringcomputation/computationschedule"
	computationusage "github.com/oracle/provider-oci/internal/controller/meteringcomputation/computationusage"
	computationusagecarbonemission "github.com/oracle/provider-oci/internal/controller/meteringcomputation/computationusagecarbonemission"
	computationusagecarbonemissionsquery "github.com/oracle/provider-oci/internal/controller/meteringcomputation/computationusagecarbonemissionsquery"
	computationusagestatementemailrecipientsgroup "github.com/oracle/provider-oci/internal/controller/meteringcomputation/computationusagestatementemailrecipientsgroup"
)

// Setup_meteringcomputation creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_meteringcomputation(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		computationcustomtable.Setup,
		computationquery.Setup,
		computationschedule.Setup,
		computationusage.Setup,
		computationusagecarbonemission.Setup,
		computationusagecarbonemissionsquery.Setup,
		computationusagestatementemailrecipientsgroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	dedicatedvantagepoint "github.com/oracle/provider-oci/internal/controller/cluster/apmsynthetics/dedicatedvantagepoint"
	monitor "github.com/oracle/provider-oci/internal/controller/cluster/apmsynthetics/monitor"
	onpremisevantagepoint "github.com/oracle/provider-oci/internal/controller/cluster/apmsynthetics/onpremisevantagepoint"
	onpremisevantagepointworker "github.com/oracle/provider-oci/internal/controller/cluster/apmsynthetics/onpremisevantagepointworker"
	script "github.com/oracle/provider-oci/internal/controller/cluster/apmsynthetics/script"
)

// Setup_apmsynthetics creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_apmsynthetics(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		dedicatedvantagepoint.Setup,
		monitor.Setup,
		onpremisevantagepoint.Setup,
		onpremisevantagepointworker.Setup,
		script.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_apmsynthetics creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_apmsynthetics(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		dedicatedvantagepoint.SetupGated,
		monitor.SetupGated,
		onpremisevantagepoint.SetupGated,
		onpremisevantagepointworker.SetupGated,
		script.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

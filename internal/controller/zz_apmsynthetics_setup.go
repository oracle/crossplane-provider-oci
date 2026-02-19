/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	dedicatedvantagepoint "github.com/oracle/provider-oci/internal/controller/apmsynthetics/dedicatedvantagepoint"
	monitor "github.com/oracle/provider-oci/internal/controller/apmsynthetics/monitor"
	onpremisevantagepoint "github.com/oracle/provider-oci/internal/controller/apmsynthetics/onpremisevantagepoint"
	onpremisevantagepointworker "github.com/oracle/provider-oci/internal/controller/apmsynthetics/onpremisevantagepointworker"
	script "github.com/oracle/provider-oci/internal/controller/apmsynthetics/script"
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

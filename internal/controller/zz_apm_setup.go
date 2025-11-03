/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	apmdomain "github.com/oracle/provider-oci/internal/controller/apm/apmdomain"
	config "github.com/oracle/provider-oci/internal/controller/apm/config"
	syntheticsdedicatedvantagepoint "github.com/oracle/provider-oci/internal/controller/apm/syntheticsdedicatedvantagepoint"
	syntheticsmonitor "github.com/oracle/provider-oci/internal/controller/apm/syntheticsmonitor"
	syntheticsonpremisevantagepoint "github.com/oracle/provider-oci/internal/controller/apm/syntheticsonpremisevantagepoint"
	syntheticsonpremisevantagepointworker "github.com/oracle/provider-oci/internal/controller/apm/syntheticsonpremisevantagepointworker"
	syntheticsscript "github.com/oracle/provider-oci/internal/controller/apm/syntheticsscript"
	tracesscheduledquery "github.com/oracle/provider-oci/internal/controller/apm/tracesscheduledquery"
)

// Setup_apm creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_apm(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		apmdomain.Setup,
		config.Setup,
		syntheticsdedicatedvantagepoint.Setup,
		syntheticsmonitor.Setup,
		syntheticsonpremisevantagepoint.Setup,
		syntheticsonpremisevantagepointworker.Setup,
		syntheticsscript.Setup,
		tracesscheduledquery.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

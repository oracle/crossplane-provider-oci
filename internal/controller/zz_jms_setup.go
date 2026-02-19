/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	fleet "github.com/oracle/provider-oci/internal/controller/jms/fleet"
	fleetadvancedfeatureconfiguration "github.com/oracle/provider-oci/internal/controller/jms/fleetadvancedfeatureconfiguration"
	fleetagentconfiguration "github.com/oracle/provider-oci/internal/controller/jms/fleetagentconfiguration"
	jmsplugin "github.com/oracle/provider-oci/internal/controller/jms/jmsplugin"
	taskschedule "github.com/oracle/provider-oci/internal/controller/jms/taskschedule"
)

// Setup_jms creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_jms(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		fleet.Setup,
		fleetadvancedfeatureconfiguration.Setup,
		fleetagentconfiguration.Setup,
		jmsplugin.Setup,
		taskschedule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	fleet "github.com/oracle/provider-oci/internal/controller/cluster/jms/fleet"
	fleetadvancedfeatureconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/jms/fleetadvancedfeatureconfiguration"
	fleetagentconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/jms/fleetagentconfiguration"
	jmsplugin "github.com/oracle/provider-oci/internal/controller/cluster/jms/jmsplugin"
	taskschedule "github.com/oracle/provider-oci/internal/controller/cluster/jms/taskschedule"
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

// SetupGated_jms creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_jms(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		fleet.SetupGated,
		fleetadvancedfeatureconfiguration.SetupGated,
		fleetagentconfiguration.SetupGated,
		jmsplugin.SetupGated,
		taskschedule.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

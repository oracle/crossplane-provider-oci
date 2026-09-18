/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	catalogitem "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/catalogitem"
	compliancepolicyrule "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/compliancepolicyrule"
	fleet "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/fleet"
	fleetcredential "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/fleetcredential"
	fleetproperty "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/fleetproperty"
	fleetresource "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/fleetresource"
	maintenancewindow "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/maintenancewindow"
	onboarding "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/onboarding"
	patch "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/patch"
	platformconfiguration "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/platformconfiguration"
	property "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/property"
	provision "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/provision"
	runbook "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/runbook"
	runbookversion "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/runbookversion"
	schedulerdefinition "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/schedulerdefinition"
	taskrecord "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/taskrecord"
)

// Setup_fleetappsmanagement creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_fleetappsmanagement(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		catalogitem.Setup,
		compliancepolicyrule.Setup,
		fleet.Setup,
		fleetcredential.Setup,
		fleetproperty.Setup,
		fleetresource.Setup,
		maintenancewindow.Setup,
		onboarding.Setup,
		patch.Setup,
		platformconfiguration.Setup,
		property.Setup,
		provision.Setup,
		runbook.Setup,
		runbookversion.Setup,
		schedulerdefinition.Setup,
		taskrecord.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_fleetappsmanagement creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_fleetappsmanagement(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		catalogitem.SetupGated,
		compliancepolicyrule.SetupGated,
		fleet.SetupGated,
		fleetcredential.SetupGated,
		fleetproperty.SetupGated,
		fleetresource.SetupGated,
		maintenancewindow.SetupGated,
		onboarding.SetupGated,
		patch.SetupGated,
		platformconfiguration.SetupGated,
		property.SetupGated,
		provision.SetupGated,
		runbook.SetupGated,
		runbookversion.SetupGated,
		schedulerdefinition.SetupGated,
		taskrecord.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

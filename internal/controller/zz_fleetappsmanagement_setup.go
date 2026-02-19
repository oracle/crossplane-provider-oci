/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	catalogitem "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/catalogitem"
	compliancepolicyrule "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/compliancepolicyrule"
	fleet "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/fleet"
	fleetcredential "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/fleetcredential"
	fleetproperty "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/fleetproperty"
	fleetresource "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/fleetresource"
	maintenancewindow "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/maintenancewindow"
	onboarding "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/onboarding"
	patch "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/patch"
	platformconfiguration "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/platformconfiguration"
	property "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/property"
	provision "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/provision"
	runbook "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/runbook"
	runbookversion "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/runbookversion"
	schedulerdefinition "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/schedulerdefinition"
	taskrecord "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/taskrecord"
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

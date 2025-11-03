/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	appsmanagementcatalogitem "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/appsmanagementcatalogitem"
	appsmanagementcompliancepolicyrule "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/appsmanagementcompliancepolicyrule"
	appsmanagementfleet "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/appsmanagementfleet"
	appsmanagementfleetcredential "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/appsmanagementfleetcredential"
	appsmanagementfleetproperty "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/appsmanagementfleetproperty"
	appsmanagementfleetresource "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/appsmanagementfleetresource"
	appsmanagementmaintenancewindow "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/appsmanagementmaintenancewindow"
	appsmanagementonboarding "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/appsmanagementonboarding"
	appsmanagementpatch "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/appsmanagementpatch"
	appsmanagementplatformconfiguration "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/appsmanagementplatformconfiguration"
	appsmanagementproperty "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/appsmanagementproperty"
	appsmanagementprovision "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/appsmanagementprovision"
	appsmanagementrunbook "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/appsmanagementrunbook"
	appsmanagementrunbookversion "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/appsmanagementrunbookversion"
	appsmanagementschedulerdefinition "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/appsmanagementschedulerdefinition"
	appsmanagementtaskrecord "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/appsmanagementtaskrecord"
	softwareupdatefsucollection "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/softwareupdatefsucollection"
	softwareupdatefsucycle "github.com/oracle/provider-oci/internal/controller/fleetappsmanagement/softwareupdatefsucycle"
)

// Setup_fleetappsmanagement creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_fleetappsmanagement(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		appsmanagementcatalogitem.Setup,
		appsmanagementcompliancepolicyrule.Setup,
		appsmanagementfleet.Setup,
		appsmanagementfleetcredential.Setup,
		appsmanagementfleetproperty.Setup,
		appsmanagementfleetresource.Setup,
		appsmanagementmaintenancewindow.Setup,
		appsmanagementonboarding.Setup,
		appsmanagementpatch.Setup,
		appsmanagementplatformconfiguration.Setup,
		appsmanagementproperty.Setup,
		appsmanagementprovision.Setup,
		appsmanagementrunbook.Setup,
		appsmanagementrunbookversion.Setup,
		appsmanagementschedulerdefinition.Setup,
		appsmanagementtaskrecord.Setup,
		softwareupdatefsucollection.Setup,
		softwareupdatefsucycle.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

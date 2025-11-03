/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	bridgeagent "github.com/oracle/provider-oci/internal/controller/cloudguard/bridgeagent"
	bridgeagentdependency "github.com/oracle/provider-oci/internal/controller/cloudguard/bridgeagentdependency"
	bridgeagentplugin "github.com/oracle/provider-oci/internal/controller/cloudguard/bridgeagentplugin"
	bridgeasset "github.com/oracle/provider-oci/internal/controller/cloudguard/bridgeasset"
	bridgeassetsource "github.com/oracle/provider-oci/internal/controller/cloudguard/bridgeassetsource"
	bridgediscoveryschedule "github.com/oracle/provider-oci/internal/controller/cloudguard/bridgediscoveryschedule"
	bridgeenvironment "github.com/oracle/provider-oci/internal/controller/cloudguard/bridgeenvironment"
	bridgeinventory "github.com/oracle/provider-oci/internal/controller/cloudguard/bridgeinventory"
	enrollmentstatus "github.com/oracle/provider-oci/internal/controller/cloudguard/enrollmentstatus"
	guardadhocquery "github.com/oracle/provider-oci/internal/controller/cloudguard/guardadhocquery"
	guardcloudguardconfiguration "github.com/oracle/provider-oci/internal/controller/cloudguard/guardcloudguardconfiguration"
	guarddetectorrecipe "github.com/oracle/provider-oci/internal/controller/cloudguard/guarddetectorrecipe"
	guardmanagedlist "github.com/oracle/provider-oci/internal/controller/cloudguard/guardmanagedlist"
	guardresponderrecipe "github.com/oracle/provider-oci/internal/controller/cloudguard/guardresponderrecipe"
	guardsavedquery "github.com/oracle/provider-oci/internal/controller/cloudguard/guardsavedquery"
	guardsecurityrecipe "github.com/oracle/provider-oci/internal/controller/cloudguard/guardsecurityrecipe"
	guardsecurityzone "github.com/oracle/provider-oci/internal/controller/cloudguard/guardsecurityzone"
	guardtarget "github.com/oracle/provider-oci/internal/controller/cloudguard/guardtarget"
	guardwlpagent "github.com/oracle/provider-oci/internal/controller/cloudguard/guardwlpagent"
	migrationsmigration "github.com/oracle/provider-oci/internal/controller/cloudguard/migrationsmigration"
	migrationsmigrationasset "github.com/oracle/provider-oci/internal/controller/cloudguard/migrationsmigrationasset"
	migrationsmigrationplan "github.com/oracle/provider-oci/internal/controller/cloudguard/migrationsmigrationplan"
	migrationsreplicationschedule "github.com/oracle/provider-oci/internal/controller/cloudguard/migrationsreplicationschedule"
	migrationstargetasset "github.com/oracle/provider-oci/internal/controller/cloudguard/migrationstargetasset"
	profile "github.com/oracle/provider-oci/internal/controller/cloudguard/profile"
	recommendation "github.com/oracle/provider-oci/internal/controller/cloudguard/recommendation"
	resourceaction "github.com/oracle/provider-oci/internal/controller/cloudguard/resourceaction"
)

// Setup_cloudguard creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cloudguard(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bridgeagent.Setup,
		bridgeagentdependency.Setup,
		bridgeagentplugin.Setup,
		bridgeasset.Setup,
		bridgeassetsource.Setup,
		bridgediscoveryschedule.Setup,
		bridgeenvironment.Setup,
		bridgeinventory.Setup,
		enrollmentstatus.Setup,
		guardadhocquery.Setup,
		guardcloudguardconfiguration.Setup,
		guarddetectorrecipe.Setup,
		guardmanagedlist.Setup,
		guardresponderrecipe.Setup,
		guardsavedquery.Setup,
		guardsecurityrecipe.Setup,
		guardsecurityzone.Setup,
		guardtarget.Setup,
		guardwlpagent.Setup,
		migrationsmigration.Setup,
		migrationsmigrationasset.Setup,
		migrationsmigrationplan.Setup,
		migrationsreplicationschedule.Setup,
		migrationstargetasset.Setup,
		profile.Setup,
		recommendation.Setup,
		resourceaction.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

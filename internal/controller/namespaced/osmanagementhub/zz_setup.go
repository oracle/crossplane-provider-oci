/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	dynamicset "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/dynamicset"
	dynamicsetinstallpackagesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/dynamicsetinstallpackagesmanagement"
	dynamicsetrebootmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/dynamicsetrebootmanagement"
	dynamicsetremovepackagesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/dynamicsetremovepackagesmanagement"
	dynamicsetupdatepackagesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/dynamicsetupdatepackagesmanagement"
	event "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/event"
	lifecycleenvironment "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/lifecycleenvironment"
	lifecyclestageattachmanagedinstancesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/lifecyclestageattachmanagedinstancesmanagement"
	lifecyclestagedetachmanagedinstancesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/lifecyclestagedetachmanagedinstancesmanagement"
	lifecyclestagepromotesoftwaresourcemanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/lifecyclestagepromotesoftwaresourcemanagement"
	lifecyclestagerebootmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/lifecyclestagerebootmanagement"
	managedinstance "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managedinstance"
	managedinstanceattachprofilemanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managedinstanceattachprofilemanagement"
	managedinstanceattachsoftwaresourcesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managedinstanceattachsoftwaresourcesmanagement"
	managedinstancedetachprofilemanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managedinstancedetachprofilemanagement"
	managedinstancedetachsoftwaresourcesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managedinstancedetachsoftwaresourcesmanagement"
	managedinstancegroup "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managedinstancegroup"
	managedinstancegroupattachmanagedinstancesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managedinstancegroupattachmanagedinstancesmanagement"
	managedinstancegroupattachsoftwaresourcesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managedinstancegroupattachsoftwaresourcesmanagement"
	managedinstancegroupdetachmanagedinstancesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managedinstancegroupdetachmanagedinstancesmanagement"
	managedinstancegroupdetachsoftwaresourcesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managedinstancegroupdetachsoftwaresourcesmanagement"
	managedinstancegroupinstallpackagesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managedinstancegroupinstallpackagesmanagement"
	managedinstancegroupinstallwindowsupdatesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managedinstancegroupinstallwindowsupdatesmanagement"
	managedinstancegroupmanagemodulestreamsmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managedinstancegroupmanagemodulestreamsmanagement"
	managedinstancegrouprebootmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managedinstancegrouprebootmanagement"
	managedinstancegroupremovepackagesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managedinstancegroupremovepackagesmanagement"
	managedinstancegroupupdateallpackagesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managedinstancegroupupdateallpackagesmanagement"
	managedinstanceinstallpackagesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managedinstanceinstallpackagesmanagement"
	managedinstanceinstallsnapsmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managedinstanceinstallsnapsmanagement"
	managedinstanceinstallwindowsupdatesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managedinstanceinstallwindowsupdatesmanagement"
	managedinstancerebootmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managedinstancerebootmanagement"
	managedinstancerefreshsoftwaremanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managedinstancerefreshsoftwaremanagement"
	managedinstanceremovepackagesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managedinstanceremovepackagesmanagement"
	managedinstanceremovesnapsmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managedinstanceremovesnapsmanagement"
	managedinstancesinstallwindowsupdatesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managedinstancesinstallwindowsupdatesmanagement"
	managedinstancesupdatepackagesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managedinstancesupdatepackagesmanagement"
	managedinstanceswitchsnapchannelmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managedinstanceswitchsnapchannelmanagement"
	managedinstanceupdatepackagesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managedinstanceupdatepackagesmanagement"
	managementstation "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managementstation"
	managementstationassociatemanagedinstancesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managementstationassociatemanagedinstancesmanagement"
	managementstationmirrorsynchronizemanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managementstationmirrorsynchronizemanagement"
	managementstationrefreshmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managementstationrefreshmanagement"
	managementstationsynchronizemirrorsmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/managementstationsynchronizemirrorsmanagement"
	profile "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/profile"
	profileattachlifecyclestagemanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/profileattachlifecyclestagemanagement"
	profileattachmanagedinstancegroupmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/profileattachmanagedinstancegroupmanagement"
	profileattachmanagementstationmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/profileattachmanagementstationmanagement"
	profileattachsoftwaresourcesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/profileattachsoftwaresourcesmanagement"
	profiledetachmanagementstationmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/profiledetachmanagementstationmanagement"
	profiledetachsoftwaresourcesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/profiledetachsoftwaresourcesmanagement"
	scheduledjob "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/scheduledjob"
	softwaresource "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/softwaresource"
	softwaresourceaddpackagesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/softwaresourceaddpackagesmanagement"
	softwaresourcechangeavailabilitymanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/softwaresourcechangeavailabilitymanagement"
	softwaresourcegeneratemetadatamanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/softwaresourcegeneratemetadatamanagement"
	softwaresourcemanifest "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/softwaresourcemanifest"
	softwaresourceremovepackagesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/softwaresourceremovepackagesmanagement"
	softwaresourcereplacepackagesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/softwaresourcereplacepackagesmanagement"
	workrequestrerunmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/workrequestrerunmanagement"
)

// Setup_osmanagementhub creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_osmanagementhub(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		dynamicset.Setup,
		dynamicsetinstallpackagesmanagement.Setup,
		dynamicsetrebootmanagement.Setup,
		dynamicsetremovepackagesmanagement.Setup,
		dynamicsetupdatepackagesmanagement.Setup,
		event.Setup,
		lifecycleenvironment.Setup,
		lifecyclestageattachmanagedinstancesmanagement.Setup,
		lifecyclestagedetachmanagedinstancesmanagement.Setup,
		lifecyclestagepromotesoftwaresourcemanagement.Setup,
		lifecyclestagerebootmanagement.Setup,
		managedinstance.Setup,
		managedinstanceattachprofilemanagement.Setup,
		managedinstanceattachsoftwaresourcesmanagement.Setup,
		managedinstancedetachprofilemanagement.Setup,
		managedinstancedetachsoftwaresourcesmanagement.Setup,
		managedinstancegroup.Setup,
		managedinstancegroupattachmanagedinstancesmanagement.Setup,
		managedinstancegroupattachsoftwaresourcesmanagement.Setup,
		managedinstancegroupdetachmanagedinstancesmanagement.Setup,
		managedinstancegroupdetachsoftwaresourcesmanagement.Setup,
		managedinstancegroupinstallpackagesmanagement.Setup,
		managedinstancegroupinstallwindowsupdatesmanagement.Setup,
		managedinstancegroupmanagemodulestreamsmanagement.Setup,
		managedinstancegrouprebootmanagement.Setup,
		managedinstancegroupremovepackagesmanagement.Setup,
		managedinstancegroupupdateallpackagesmanagement.Setup,
		managedinstanceinstallpackagesmanagement.Setup,
		managedinstanceinstallsnapsmanagement.Setup,
		managedinstanceinstallwindowsupdatesmanagement.Setup,
		managedinstancerebootmanagement.Setup,
		managedinstancerefreshsoftwaremanagement.Setup,
		managedinstanceremovepackagesmanagement.Setup,
		managedinstanceremovesnapsmanagement.Setup,
		managedinstancesinstallwindowsupdatesmanagement.Setup,
		managedinstancesupdatepackagesmanagement.Setup,
		managedinstanceswitchsnapchannelmanagement.Setup,
		managedinstanceupdatepackagesmanagement.Setup,
		managementstation.Setup,
		managementstationassociatemanagedinstancesmanagement.Setup,
		managementstationmirrorsynchronizemanagement.Setup,
		managementstationrefreshmanagement.Setup,
		managementstationsynchronizemirrorsmanagement.Setup,
		profile.Setup,
		profileattachlifecyclestagemanagement.Setup,
		profileattachmanagedinstancegroupmanagement.Setup,
		profileattachmanagementstationmanagement.Setup,
		profileattachsoftwaresourcesmanagement.Setup,
		profiledetachmanagementstationmanagement.Setup,
		profiledetachsoftwaresourcesmanagement.Setup,
		scheduledjob.Setup,
		softwaresource.Setup,
		softwaresourceaddpackagesmanagement.Setup,
		softwaresourcechangeavailabilitymanagement.Setup,
		softwaresourcegeneratemetadatamanagement.Setup,
		softwaresourcemanifest.Setup,
		softwaresourceremovepackagesmanagement.Setup,
		softwaresourcereplacepackagesmanagement.Setup,
		workrequestrerunmanagement.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_osmanagementhub creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_osmanagementhub(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		dynamicset.SetupGated,
		dynamicsetinstallpackagesmanagement.SetupGated,
		dynamicsetrebootmanagement.SetupGated,
		dynamicsetremovepackagesmanagement.SetupGated,
		dynamicsetupdatepackagesmanagement.SetupGated,
		event.SetupGated,
		lifecycleenvironment.SetupGated,
		lifecyclestageattachmanagedinstancesmanagement.SetupGated,
		lifecyclestagedetachmanagedinstancesmanagement.SetupGated,
		lifecyclestagepromotesoftwaresourcemanagement.SetupGated,
		lifecyclestagerebootmanagement.SetupGated,
		managedinstance.SetupGated,
		managedinstanceattachprofilemanagement.SetupGated,
		managedinstanceattachsoftwaresourcesmanagement.SetupGated,
		managedinstancedetachprofilemanagement.SetupGated,
		managedinstancedetachsoftwaresourcesmanagement.SetupGated,
		managedinstancegroup.SetupGated,
		managedinstancegroupattachmanagedinstancesmanagement.SetupGated,
		managedinstancegroupattachsoftwaresourcesmanagement.SetupGated,
		managedinstancegroupdetachmanagedinstancesmanagement.SetupGated,
		managedinstancegroupdetachsoftwaresourcesmanagement.SetupGated,
		managedinstancegroupinstallpackagesmanagement.SetupGated,
		managedinstancegroupinstallwindowsupdatesmanagement.SetupGated,
		managedinstancegroupmanagemodulestreamsmanagement.SetupGated,
		managedinstancegrouprebootmanagement.SetupGated,
		managedinstancegroupremovepackagesmanagement.SetupGated,
		managedinstancegroupupdateallpackagesmanagement.SetupGated,
		managedinstanceinstallpackagesmanagement.SetupGated,
		managedinstanceinstallsnapsmanagement.SetupGated,
		managedinstanceinstallwindowsupdatesmanagement.SetupGated,
		managedinstancerebootmanagement.SetupGated,
		managedinstancerefreshsoftwaremanagement.SetupGated,
		managedinstanceremovepackagesmanagement.SetupGated,
		managedinstanceremovesnapsmanagement.SetupGated,
		managedinstancesinstallwindowsupdatesmanagement.SetupGated,
		managedinstancesupdatepackagesmanagement.SetupGated,
		managedinstanceswitchsnapchannelmanagement.SetupGated,
		managedinstanceupdatepackagesmanagement.SetupGated,
		managementstation.SetupGated,
		managementstationassociatemanagedinstancesmanagement.SetupGated,
		managementstationmirrorsynchronizemanagement.SetupGated,
		managementstationrefreshmanagement.SetupGated,
		managementstationsynchronizemirrorsmanagement.SetupGated,
		profile.SetupGated,
		profileattachlifecyclestagemanagement.SetupGated,
		profileattachmanagedinstancegroupmanagement.SetupGated,
		profileattachmanagementstationmanagement.SetupGated,
		profileattachsoftwaresourcesmanagement.SetupGated,
		profiledetachmanagementstationmanagement.SetupGated,
		profiledetachsoftwaresourcesmanagement.SetupGated,
		scheduledjob.SetupGated,
		softwaresource.SetupGated,
		softwaresourceaddpackagesmanagement.SetupGated,
		softwaresourcechangeavailabilitymanagement.SetupGated,
		softwaresourcegeneratemetadatamanagement.SetupGated,
		softwaresourcemanifest.SetupGated,
		softwaresourceremovepackagesmanagement.SetupGated,
		softwaresourcereplacepackagesmanagement.SetupGated,
		workrequestrerunmanagement.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

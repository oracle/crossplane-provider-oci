/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	event "github.com/oracle/provider-oci/internal/controller/osmanagementhub/event"
	lifecycleenvironment "github.com/oracle/provider-oci/internal/controller/osmanagementhub/lifecycleenvironment"
	lifecyclestageattachmanagedinstancesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/lifecyclestageattachmanagedinstancesmanagement"
	lifecyclestagedetachmanagedinstancesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/lifecyclestagedetachmanagedinstancesmanagement"
	lifecyclestagepromotesoftwaresourcemanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/lifecyclestagepromotesoftwaresourcemanagement"
	lifecyclestagerebootmanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/lifecyclestagerebootmanagement"
	managedinstance "github.com/oracle/provider-oci/internal/controller/osmanagementhub/managedinstance"
	managedinstanceattachprofilemanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/managedinstanceattachprofilemanagement"
	managedinstancedetachprofilemanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/managedinstancedetachprofilemanagement"
	managedinstancegroup "github.com/oracle/provider-oci/internal/controller/osmanagementhub/managedinstancegroup"
	managedinstancegroupattachmanagedinstancesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/managedinstancegroupattachmanagedinstancesmanagement"
	managedinstancegroupattachsoftwaresourcesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/managedinstancegroupattachsoftwaresourcesmanagement"
	managedinstancegroupdetachmanagedinstancesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/managedinstancegroupdetachmanagedinstancesmanagement"
	managedinstancegroupdetachsoftwaresourcesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/managedinstancegroupdetachsoftwaresourcesmanagement"
	managedinstancegroupinstallpackagesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/managedinstancegroupinstallpackagesmanagement"
	managedinstancegroupinstallwindowsupdatesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/managedinstancegroupinstallwindowsupdatesmanagement"
	managedinstancegroupmanagemodulestreamsmanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/managedinstancegroupmanagemodulestreamsmanagement"
	managedinstancegrouprebootmanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/managedinstancegrouprebootmanagement"
	managedinstancegroupremovepackagesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/managedinstancegroupremovepackagesmanagement"
	managedinstancegroupupdateallpackagesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/managedinstancegroupupdateallpackagesmanagement"
	managedinstanceinstallwindowsupdatesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/managedinstanceinstallwindowsupdatesmanagement"
	managedinstancerebootmanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/managedinstancerebootmanagement"
	managedinstanceupdatepackagesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/managedinstanceupdatepackagesmanagement"
	managementstation "github.com/oracle/provider-oci/internal/controller/osmanagementhub/managementstation"
	managementstationassociatemanagedinstancesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/managementstationassociatemanagedinstancesmanagement"
	managementstationmirrorsynchronizemanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/managementstationmirrorsynchronizemanagement"
	managementstationrefreshmanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/managementstationrefreshmanagement"
	managementstationsynchronizemirrorsmanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/managementstationsynchronizemirrorsmanagement"
	profile "github.com/oracle/provider-oci/internal/controller/osmanagementhub/profile"
	profileattachlifecyclestagemanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/profileattachlifecyclestagemanagement"
	profileattachmanagedinstancegroupmanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/profileattachmanagedinstancegroupmanagement"
	profileattachmanagementstationmanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/profileattachmanagementstationmanagement"
	profileattachsoftwaresourcesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/profileattachsoftwaresourcesmanagement"
	profiledetachsoftwaresourcesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/profiledetachsoftwaresourcesmanagement"
	scheduledjob "github.com/oracle/provider-oci/internal/controller/osmanagementhub/scheduledjob"
	softwaresource "github.com/oracle/provider-oci/internal/controller/osmanagementhub/softwaresource"
	softwaresourceaddpackagesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/softwaresourceaddpackagesmanagement"
	softwaresourcechangeavailabilitymanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/softwaresourcechangeavailabilitymanagement"
	softwaresourcegeneratemetadatamanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/softwaresourcegeneratemetadatamanagement"
	softwaresourcemanifest "github.com/oracle/provider-oci/internal/controller/osmanagementhub/softwaresourcemanifest"
	softwaresourceremovepackagesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/softwaresourceremovepackagesmanagement"
	softwaresourcereplacepackagesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/softwaresourcereplacepackagesmanagement"
	workrequestrerunmanagement "github.com/oracle/provider-oci/internal/controller/osmanagementhub/workrequestrerunmanagement"
)

// Setup_osmanagementhub creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_osmanagementhub(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		event.Setup,
		lifecycleenvironment.Setup,
		lifecyclestageattachmanagedinstancesmanagement.Setup,
		lifecyclestagedetachmanagedinstancesmanagement.Setup,
		lifecyclestagepromotesoftwaresourcemanagement.Setup,
		lifecyclestagerebootmanagement.Setup,
		managedinstance.Setup,
		managedinstanceattachprofilemanagement.Setup,
		managedinstancedetachprofilemanagement.Setup,
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
		managedinstanceinstallwindowsupdatesmanagement.Setup,
		managedinstancerebootmanagement.Setup,
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

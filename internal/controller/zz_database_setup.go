/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	applicationvip "github.com/oracle/provider-oci/internal/controller/database/applicationvip"
	autonomouscontainerdatabase "github.com/oracle/provider-oci/internal/controller/database/autonomouscontainerdatabase"
	autonomouscontainerdatabaseaddstandby "github.com/oracle/provider-oci/internal/controller/database/autonomouscontainerdatabaseaddstandby"
	autonomouscontainerdatabasedataguardassociation "github.com/oracle/provider-oci/internal/controller/database/autonomouscontainerdatabasedataguardassociation"
	autonomouscontainerdatabasedataguardassociationoperation "github.com/oracle/provider-oci/internal/controller/database/autonomouscontainerdatabasedataguardassociationoperation"
	autonomouscontainerdatabasedataguardrolechange "github.com/oracle/provider-oci/internal/controller/database/autonomouscontainerdatabasedataguardrolechange"
	autonomouscontainerdatabasesnapshotstandby "github.com/oracle/provider-oci/internal/controller/database/autonomouscontainerdatabasesnapshotstandby"
	autonomousdatabase "github.com/oracle/provider-oci/internal/controller/database/autonomousdatabase"
	autonomousdatabasebackup "github.com/oracle/provider-oci/internal/controller/database/autonomousdatabasebackup"
	autonomousdatabaseinstancewalletmanagement "github.com/oracle/provider-oci/internal/controller/database/autonomousdatabaseinstancewalletmanagement"
	autonomousdatabaseregionalwalletmanagement "github.com/oracle/provider-oci/internal/controller/database/autonomousdatabaseregionalwalletmanagement"
	autonomousdatabasesaasadminuser "github.com/oracle/provider-oci/internal/controller/database/autonomousdatabasesaasadminuser"
	autonomousdatabasesoftwareimage "github.com/oracle/provider-oci/internal/controller/database/autonomousdatabasesoftwareimage"
	autonomousdatabasewallet "github.com/oracle/provider-oci/internal/controller/database/autonomousdatabasewallet"
	autonomousexadatainfrastructure "github.com/oracle/provider-oci/internal/controller/database/autonomousexadatainfrastructure"
	autonomousvmcluster "github.com/oracle/provider-oci/internal/controller/database/autonomousvmcluster"
	autonomousvmclusterordscertificatemanagement "github.com/oracle/provider-oci/internal/controller/database/autonomousvmclusterordscertificatemanagement"
	autonomousvmclustersslcertificatemanagement "github.com/oracle/provider-oci/internal/controller/database/autonomousvmclustersslcertificatemanagement"
	backup "github.com/oracle/provider-oci/internal/controller/database/backup"
	backupcancelmanagement "github.com/oracle/provider-oci/internal/controller/database/backupcancelmanagement"
	backupdestination "github.com/oracle/provider-oci/internal/controller/database/backupdestination"
	cloudautonomousvmcluster "github.com/oracle/provider-oci/internal/controller/database/cloudautonomousvmcluster"
	clouddatabasemanagement "github.com/oracle/provider-oci/internal/controller/database/clouddatabasemanagement"
	cloudexadatainfrastructure "github.com/oracle/provider-oci/internal/controller/database/cloudexadatainfrastructure"
	cloudexadatainfrastructureconfigureexascalemanagement "github.com/oracle/provider-oci/internal/controller/database/cloudexadatainfrastructureconfigureexascalemanagement"
	cloudvmcluster "github.com/oracle/provider-oci/internal/controller/database/cloudvmcluster"
	cloudvmclusteriormconfig "github.com/oracle/provider-oci/internal/controller/database/cloudvmclusteriormconfig"
	database "github.com/oracle/provider-oci/internal/controller/database/database"
	databasesnapshotstandby "github.com/oracle/provider-oci/internal/controller/database/databasesnapshotstandby"
	databasesoftwareimage "github.com/oracle/provider-oci/internal/controller/database/databasesoftwareimage"
	databaseupgrade "github.com/oracle/provider-oci/internal/controller/database/databaseupgrade"
	dataguardassociation "github.com/oracle/provider-oci/internal/controller/database/dataguardassociation"
	dbhome "github.com/oracle/provider-oci/internal/controller/database/dbhome"
	dbnode "github.com/oracle/provider-oci/internal/controller/database/dbnode"
	dbnodeconsoleconnection "github.com/oracle/provider-oci/internal/controller/database/dbnodeconsoleconnection"
	dbnodeconsolehistory "github.com/oracle/provider-oci/internal/controller/database/dbnodeconsolehistory"
	dbnodesnapshot "github.com/oracle/provider-oci/internal/controller/database/dbnodesnapshot"
	dbnodesnapshotmanagement "github.com/oracle/provider-oci/internal/controller/database/dbnodesnapshotmanagement"
	dbsystem "github.com/oracle/provider-oci/internal/controller/database/dbsystem"
	dbsystemsupgrade "github.com/oracle/provider-oci/internal/controller/database/dbsystemsupgrade"
	exadatainfrastructure "github.com/oracle/provider-oci/internal/controller/database/exadatainfrastructure"
	exadatainfrastructurecompute "github.com/oracle/provider-oci/internal/controller/database/exadatainfrastructurecompute"
	exadatainfrastructureconfigureexascalemanagement "github.com/oracle/provider-oci/internal/controller/database/exadatainfrastructureconfigureexascalemanagement"
	exadatainfrastructurestorage "github.com/oracle/provider-oci/internal/controller/database/exadatainfrastructurestorage"
	exadataiormconfig "github.com/oracle/provider-oci/internal/controller/database/exadataiormconfig"
	exadbvmcluster "github.com/oracle/provider-oci/internal/controller/database/exadbvmcluster"
	exascaledbstoragevault "github.com/oracle/provider-oci/internal/controller/database/exascaledbstoragevault"
	executionaction "github.com/oracle/provider-oci/internal/controller/database/executionaction"
	executionwindow "github.com/oracle/provider-oci/internal/controller/database/executionwindow"
	externalcontainerdatabase "github.com/oracle/provider-oci/internal/controller/database/externalcontainerdatabase"
	externalcontainerdatabasemanagement "github.com/oracle/provider-oci/internal/controller/database/externalcontainerdatabasemanagement"
	externalcontainerdatabasesstackmonitoring "github.com/oracle/provider-oci/internal/controller/database/externalcontainerdatabasesstackmonitoring"
	externaldatabaseconnector "github.com/oracle/provider-oci/internal/controller/database/externaldatabaseconnector"
	externalnoncontainerdatabase "github.com/oracle/provider-oci/internal/controller/database/externalnoncontainerdatabase"
	externalnoncontainerdatabasemanagement "github.com/oracle/provider-oci/internal/controller/database/externalnoncontainerdatabasemanagement"
	externalnoncontainerdatabaseoperationsinsightsmanagement "github.com/oracle/provider-oci/internal/controller/database/externalnoncontainerdatabaseoperationsinsightsmanagement"
	externalnoncontainerdatabasesstackmonitoring "github.com/oracle/provider-oci/internal/controller/database/externalnoncontainerdatabasesstackmonitoring"
	externalpluggabledatabase "github.com/oracle/provider-oci/internal/controller/database/externalpluggabledatabase"
	externalpluggabledatabasemanagement "github.com/oracle/provider-oci/internal/controller/database/externalpluggabledatabasemanagement"
	externalpluggabledatabaseoperationsinsightsmanagement "github.com/oracle/provider-oci/internal/controller/database/externalpluggabledatabaseoperationsinsightsmanagement"
	externalpluggabledatabasesstackmonitoring "github.com/oracle/provider-oci/internal/controller/database/externalpluggabledatabasesstackmonitoring"
	keystore "github.com/oracle/provider-oci/internal/controller/database/keystore"
	maintenancerun "github.com/oracle/provider-oci/internal/controller/database/maintenancerun"
	managementautonomousdatabaseautonomousdatabasedbmfeaturesmanagement "github.com/oracle/provider-oci/internal/controller/database/managementautonomousdatabaseautonomousdatabasedbmfeaturesmanagement"
	managementcloudasm "github.com/oracle/provider-oci/internal/controller/database/managementcloudasm"
	managementcloudasminstance "github.com/oracle/provider-oci/internal/controller/database/managementcloudasminstance"
	managementcloudcluster "github.com/oracle/provider-oci/internal/controller/database/managementcloudcluster"
	managementcloudclusterinstance "github.com/oracle/provider-oci/internal/controller/database/managementcloudclusterinstance"
	managementclouddbhome "github.com/oracle/provider-oci/internal/controller/database/managementclouddbhome"
	managementclouddbnode "github.com/oracle/provider-oci/internal/controller/database/managementclouddbnode"
	managementclouddbsystem "github.com/oracle/provider-oci/internal/controller/database/managementclouddbsystem"
	managementclouddbsystemclouddatabasemanagementsmanagement "github.com/oracle/provider-oci/internal/controller/database/managementclouddbsystemclouddatabasemanagementsmanagement"
	managementclouddbsystemcloudstackmonitoringsmanagement "github.com/oracle/provider-oci/internal/controller/database/managementclouddbsystemcloudstackmonitoringsmanagement"
	managementclouddbsystemconnector "github.com/oracle/provider-oci/internal/controller/database/managementclouddbsystemconnector"
	managementclouddbsystemdiscovery "github.com/oracle/provider-oci/internal/controller/database/managementclouddbsystemdiscovery"
	managementcloudlistener "github.com/oracle/provider-oci/internal/controller/database/managementcloudlistener"
	managementdatabasedbmfeaturesmanagement "github.com/oracle/provider-oci/internal/controller/database/managementdatabasedbmfeaturesmanagement"
	managementdbmanagementprivateendpoint "github.com/oracle/provider-oci/internal/controller/database/managementdbmanagementprivateendpoint"
	managementexternalasm "github.com/oracle/provider-oci/internal/controller/database/managementexternalasm"
	managementexternalasminstance "github.com/oracle/provider-oci/internal/controller/database/managementexternalasminstance"
	managementexternalcluster "github.com/oracle/provider-oci/internal/controller/database/managementexternalcluster"
	managementexternalclusterinstance "github.com/oracle/provider-oci/internal/controller/database/managementexternalclusterinstance"
	managementexternalcontainerdatabaseexternalcontainerdbmfeaturesmanagement "github.com/oracle/provider-oci/internal/controller/database/managementexternalcontainerdatabaseexternalcontainerdbmfeaturesmanagement"
	managementexternaldbhome "github.com/oracle/provider-oci/internal/controller/database/managementexternaldbhome"
	managementexternaldbnode "github.com/oracle/provider-oci/internal/controller/database/managementexternaldbnode"
	managementexternaldbsystem "github.com/oracle/provider-oci/internal/controller/database/managementexternaldbsystem"
	managementexternaldbsystemconnector "github.com/oracle/provider-oci/internal/controller/database/managementexternaldbsystemconnector"
	managementexternaldbsystemdatabasemanagementsmanagement "github.com/oracle/provider-oci/internal/controller/database/managementexternaldbsystemdatabasemanagementsmanagement"
	managementexternaldbsystemdiscovery "github.com/oracle/provider-oci/internal/controller/database/managementexternaldbsystemdiscovery"
	managementexternaldbsystemstackmonitoringsmanagement "github.com/oracle/provider-oci/internal/controller/database/managementexternaldbsystemstackmonitoringsmanagement"
	managementexternalexadatainfrastructure "github.com/oracle/provider-oci/internal/controller/database/managementexternalexadatainfrastructure"
	managementexternalexadatainfrastructureexadatamanagement "github.com/oracle/provider-oci/internal/controller/database/managementexternalexadatainfrastructureexadatamanagement"
	managementexternalexadatastorageconnector "github.com/oracle/provider-oci/internal/controller/database/managementexternalexadatastorageconnector"
	managementexternalexadatastoragegrid "github.com/oracle/provider-oci/internal/controller/database/managementexternalexadatastoragegrid"
	managementexternalexadatastorageserver "github.com/oracle/provider-oci/internal/controller/database/managementexternalexadatastorageserver"
	managementexternallistener "github.com/oracle/provider-oci/internal/controller/database/managementexternallistener"
	managementexternalmysqldatabase "github.com/oracle/provider-oci/internal/controller/database/managementexternalmysqldatabase"
	managementexternalmysqldatabaseconnector "github.com/oracle/provider-oci/internal/controller/database/managementexternalmysqldatabaseconnector"
	managementexternalmysqldatabaseexternalmysqldatabasesmanagement "github.com/oracle/provider-oci/internal/controller/database/managementexternalmysqldatabaseexternalmysqldatabasesmanagement"
	managementexternalnoncontainerdatabaseexternalnoncontainerdbmfeaturesmanagement "github.com/oracle/provider-oci/internal/controller/database/managementexternalnoncontainerdatabaseexternalnoncontainerdbmfeaturesmanagement"
	managementexternalpluggabledatabaseexternalpluggabledbmfeaturesmanagement "github.com/oracle/provider-oci/internal/controller/database/managementexternalpluggabledatabaseexternalpluggabledbmfeaturesmanagement"
	managementmanageddatabase "github.com/oracle/provider-oci/internal/controller/database/managementmanageddatabase"
	managementmanageddatabasegroup "github.com/oracle/provider-oci/internal/controller/database/managementmanageddatabasegroup"
	managementmanageddatabaseschangedatabaseparameter "github.com/oracle/provider-oci/internal/controller/database/managementmanageddatabaseschangedatabaseparameter"
	managementmanageddatabasesresetdatabaseparameter "github.com/oracle/provider-oci/internal/controller/database/managementmanageddatabasesresetdatabaseparameter"
	managementnamedcredential "github.com/oracle/provider-oci/internal/controller/database/managementnamedcredential"
	managementpluggabledatabasepluggabledatabasedbmfeaturesmanagement "github.com/oracle/provider-oci/internal/controller/database/managementpluggabledatabasepluggabledatabasedbmfeaturesmanagement"
	migration "github.com/oracle/provider-oci/internal/controller/database/migration"
	migrationconnection "github.com/oracle/provider-oci/internal/controller/database/migrationconnection"
	migrationjob "github.com/oracle/provider-oci/internal/controller/database/migrationjob"
	migrationmigration "github.com/oracle/provider-oci/internal/controller/database/migrationmigration"
	oneoffpatch "github.com/oracle/provider-oci/internal/controller/database/oneoffpatch"
	pluggabledatabase "github.com/oracle/provider-oci/internal/controller/database/pluggabledatabase"
	pluggabledatabasepluggabledatabasemanagementsmanagement "github.com/oracle/provider-oci/internal/controller/database/pluggabledatabasepluggabledatabasemanagementsmanagement"
	pluggabledatabaseslocalclone "github.com/oracle/provider-oci/internal/controller/database/pluggabledatabaseslocalclone"
	pluggabledatabasesnapshot "github.com/oracle/provider-oci/internal/controller/database/pluggabledatabasesnapshot"
	pluggabledatabasesremoteclone "github.com/oracle/provider-oci/internal/controller/database/pluggabledatabasesremoteclone"
	scheduledaction "github.com/oracle/provider-oci/internal/controller/database/scheduledaction"
	schedulingplan "github.com/oracle/provider-oci/internal/controller/database/schedulingplan"
	schedulingpolicy "github.com/oracle/provider-oci/internal/controller/database/schedulingpolicy"
	schedulingpolicyschedulingwindow "github.com/oracle/provider-oci/internal/controller/database/schedulingpolicyschedulingwindow"
	toolsdatabasetoolsconnection "github.com/oracle/provider-oci/internal/controller/database/toolsdatabasetoolsconnection"
	toolsdatabasetoolsidentity "github.com/oracle/provider-oci/internal/controller/database/toolsdatabasetoolsidentity"
	toolsdatabasetoolsprivateendpoint "github.com/oracle/provider-oci/internal/controller/database/toolsdatabasetoolsprivateendpoint"
	vmcluster "github.com/oracle/provider-oci/internal/controller/database/vmcluster"
	vmclusteraddvirtualmachine "github.com/oracle/provider-oci/internal/controller/database/vmclusteraddvirtualmachine"
	vmclusternetwork "github.com/oracle/provider-oci/internal/controller/database/vmclusternetwork"
	vmclusterremovevirtualmachine "github.com/oracle/provider-oci/internal/controller/database/vmclusterremovevirtualmachine"
)

// Setup_database creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_database(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		applicationvip.Setup,
		autonomouscontainerdatabase.Setup,
		autonomouscontainerdatabaseaddstandby.Setup,
		autonomouscontainerdatabasedataguardassociation.Setup,
		autonomouscontainerdatabasedataguardassociationoperation.Setup,
		autonomouscontainerdatabasedataguardrolechange.Setup,
		autonomouscontainerdatabasesnapshotstandby.Setup,
		autonomousdatabase.Setup,
		autonomousdatabasebackup.Setup,
		autonomousdatabaseinstancewalletmanagement.Setup,
		autonomousdatabaseregionalwalletmanagement.Setup,
		autonomousdatabasesaasadminuser.Setup,
		autonomousdatabasesoftwareimage.Setup,
		autonomousdatabasewallet.Setup,
		autonomousexadatainfrastructure.Setup,
		autonomousvmcluster.Setup,
		autonomousvmclusterordscertificatemanagement.Setup,
		autonomousvmclustersslcertificatemanagement.Setup,
		backup.Setup,
		backupcancelmanagement.Setup,
		backupdestination.Setup,
		cloudautonomousvmcluster.Setup,
		clouddatabasemanagement.Setup,
		cloudexadatainfrastructure.Setup,
		cloudexadatainfrastructureconfigureexascalemanagement.Setup,
		cloudvmcluster.Setup,
		cloudvmclusteriormconfig.Setup,
		database.Setup,
		databasesnapshotstandby.Setup,
		databasesoftwareimage.Setup,
		databaseupgrade.Setup,
		dataguardassociation.Setup,
		dbhome.Setup,
		dbnode.Setup,
		dbnodeconsoleconnection.Setup,
		dbnodeconsolehistory.Setup,
		dbnodesnapshot.Setup,
		dbnodesnapshotmanagement.Setup,
		dbsystem.Setup,
		dbsystemsupgrade.Setup,
		exadatainfrastructure.Setup,
		exadatainfrastructurecompute.Setup,
		exadatainfrastructureconfigureexascalemanagement.Setup,
		exadatainfrastructurestorage.Setup,
		exadataiormconfig.Setup,
		exadbvmcluster.Setup,
		exascaledbstoragevault.Setup,
		executionaction.Setup,
		executionwindow.Setup,
		externalcontainerdatabase.Setup,
		externalcontainerdatabasemanagement.Setup,
		externalcontainerdatabasesstackmonitoring.Setup,
		externaldatabaseconnector.Setup,
		externalnoncontainerdatabase.Setup,
		externalnoncontainerdatabasemanagement.Setup,
		externalnoncontainerdatabaseoperationsinsightsmanagement.Setup,
		externalnoncontainerdatabasesstackmonitoring.Setup,
		externalpluggabledatabase.Setup,
		externalpluggabledatabasemanagement.Setup,
		externalpluggabledatabaseoperationsinsightsmanagement.Setup,
		externalpluggabledatabasesstackmonitoring.Setup,
		keystore.Setup,
		maintenancerun.Setup,
		managementautonomousdatabaseautonomousdatabasedbmfeaturesmanagement.Setup,
		managementcloudasm.Setup,
		managementcloudasminstance.Setup,
		managementcloudcluster.Setup,
		managementcloudclusterinstance.Setup,
		managementclouddbhome.Setup,
		managementclouddbnode.Setup,
		managementclouddbsystem.Setup,
		managementclouddbsystemclouddatabasemanagementsmanagement.Setup,
		managementclouddbsystemcloudstackmonitoringsmanagement.Setup,
		managementclouddbsystemconnector.Setup,
		managementclouddbsystemdiscovery.Setup,
		managementcloudlistener.Setup,
		managementdatabasedbmfeaturesmanagement.Setup,
		managementdbmanagementprivateendpoint.Setup,
		managementexternalasm.Setup,
		managementexternalasminstance.Setup,
		managementexternalcluster.Setup,
		managementexternalclusterinstance.Setup,
		managementexternalcontainerdatabaseexternalcontainerdbmfeaturesmanagement.Setup,
		managementexternaldbhome.Setup,
		managementexternaldbnode.Setup,
		managementexternaldbsystem.Setup,
		managementexternaldbsystemconnector.Setup,
		managementexternaldbsystemdatabasemanagementsmanagement.Setup,
		managementexternaldbsystemdiscovery.Setup,
		managementexternaldbsystemstackmonitoringsmanagement.Setup,
		managementexternalexadatainfrastructure.Setup,
		managementexternalexadatainfrastructureexadatamanagement.Setup,
		managementexternalexadatastorageconnector.Setup,
		managementexternalexadatastoragegrid.Setup,
		managementexternalexadatastorageserver.Setup,
		managementexternallistener.Setup,
		managementexternalmysqldatabase.Setup,
		managementexternalmysqldatabaseconnector.Setup,
		managementexternalmysqldatabaseexternalmysqldatabasesmanagement.Setup,
		managementexternalnoncontainerdatabaseexternalnoncontainerdbmfeaturesmanagement.Setup,
		managementexternalpluggabledatabaseexternalpluggabledbmfeaturesmanagement.Setup,
		managementmanageddatabase.Setup,
		managementmanageddatabasegroup.Setup,
		managementmanageddatabaseschangedatabaseparameter.Setup,
		managementmanageddatabasesresetdatabaseparameter.Setup,
		managementnamedcredential.Setup,
		managementpluggabledatabasepluggabledatabasedbmfeaturesmanagement.Setup,
		migration.Setup,
		migrationconnection.Setup,
		migrationjob.Setup,
		migrationmigration.Setup,
		oneoffpatch.Setup,
		pluggabledatabase.Setup,
		pluggabledatabasepluggabledatabasemanagementsmanagement.Setup,
		pluggabledatabaseslocalclone.Setup,
		pluggabledatabasesnapshot.Setup,
		pluggabledatabasesremoteclone.Setup,
		scheduledaction.Setup,
		schedulingplan.Setup,
		schedulingpolicy.Setup,
		schedulingpolicyschedulingwindow.Setup,
		toolsdatabasetoolsconnection.Setup,
		toolsdatabasetoolsidentity.Setup,
		toolsdatabasetoolsprivateendpoint.Setup,
		vmcluster.Setup,
		vmclusteraddvirtualmachine.Setup,
		vmclusternetwork.Setup,
		vmclusterremovevirtualmachine.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

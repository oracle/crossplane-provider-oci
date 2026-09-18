/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	advancedclusterfilesystem "github.com/oracle/provider-oci/internal/controller/cluster/database/advancedclusterfilesystem"
	advancedclusterfilesystemmount "github.com/oracle/provider-oci/internal/controller/cluster/database/advancedclusterfilesystemmount"
	advancedclusterfilesystemunmount "github.com/oracle/provider-oci/internal/controller/cluster/database/advancedclusterfilesystemunmount"
	applicationvip "github.com/oracle/provider-oci/internal/controller/cluster/database/applicationvip"
	autonomouscontainerdatabase "github.com/oracle/provider-oci/internal/controller/cluster/database/autonomouscontainerdatabase"
	autonomouscontainerdatabaseaddstandby "github.com/oracle/provider-oci/internal/controller/cluster/database/autonomouscontainerdatabaseaddstandby"
	autonomouscontainerdatabasedataguardassociation "github.com/oracle/provider-oci/internal/controller/cluster/database/autonomouscontainerdatabasedataguardassociation"
	autonomouscontainerdatabasedataguardassociationoperation "github.com/oracle/provider-oci/internal/controller/cluster/database/autonomouscontainerdatabasedataguardassociationoperation"
	autonomouscontainerdatabasedataguardrolechange "github.com/oracle/provider-oci/internal/controller/cluster/database/autonomouscontainerdatabasedataguardrolechange"
	autonomouscontainerdatabasesnapshotstandby "github.com/oracle/provider-oci/internal/controller/cluster/database/autonomouscontainerdatabasesnapshotstandby"
	autonomousdatabase "github.com/oracle/provider-oci/internal/controller/cluster/database/autonomousdatabase"
	autonomousdatabasebackup "github.com/oracle/provider-oci/internal/controller/cluster/database/autonomousdatabasebackup"
	autonomousdatabaseinstancewalletmanagement "github.com/oracle/provider-oci/internal/controller/cluster/database/autonomousdatabaseinstancewalletmanagement"
	autonomousdatabaseregionalwalletmanagement "github.com/oracle/provider-oci/internal/controller/cluster/database/autonomousdatabaseregionalwalletmanagement"
	autonomousdatabasesaasadminuser "github.com/oracle/provider-oci/internal/controller/cluster/database/autonomousdatabasesaasadminuser"
	autonomousdatabasesoftwareimage "github.com/oracle/provider-oci/internal/controller/cluster/database/autonomousdatabasesoftwareimage"
	autonomousdatabasewallet "github.com/oracle/provider-oci/internal/controller/cluster/database/autonomousdatabasewallet"
	autonomousexadatainfrastructure "github.com/oracle/provider-oci/internal/controller/cluster/database/autonomousexadatainfrastructure"
	autonomousvmcluster "github.com/oracle/provider-oci/internal/controller/cluster/database/autonomousvmcluster"
	autonomousvmclusterordscertificatemanagement "github.com/oracle/provider-oci/internal/controller/cluster/database/autonomousvmclusterordscertificatemanagement"
	autonomousvmclustersslcertificatemanagement "github.com/oracle/provider-oci/internal/controller/cluster/database/autonomousvmclustersslcertificatemanagement"
	backup "github.com/oracle/provider-oci/internal/controller/cluster/database/backup"
	backupcancelmanagement "github.com/oracle/provider-oci/internal/controller/cluster/database/backupcancelmanagement"
	backupdestination "github.com/oracle/provider-oci/internal/controller/cluster/database/backupdestination"
	cloudautonomousvmcluster "github.com/oracle/provider-oci/internal/controller/cluster/database/cloudautonomousvmcluster"
	clouddatabasemanagement "github.com/oracle/provider-oci/internal/controller/cluster/database/clouddatabasemanagement"
	cloudexadatainfrastructure "github.com/oracle/provider-oci/internal/controller/cluster/database/cloudexadatainfrastructure"
	cloudexadatainfrastructureconfigureexascalemanagement "github.com/oracle/provider-oci/internal/controller/cluster/database/cloudexadatainfrastructureconfigureexascalemanagement"
	cloudvmcluster "github.com/oracle/provider-oci/internal/controller/cluster/database/cloudvmcluster"
	cloudvmclusteriormconfig "github.com/oracle/provider-oci/internal/controller/cluster/database/cloudvmclusteriormconfig"
	database "github.com/oracle/provider-oci/internal/controller/cluster/database/database"
	databasesnapshotstandby "github.com/oracle/provider-oci/internal/controller/cluster/database/databasesnapshotstandby"
	databasesoftwareimage "github.com/oracle/provider-oci/internal/controller/cluster/database/databasesoftwareimage"
	databasesoftwareschedulemanagement "github.com/oracle/provider-oci/internal/controller/cluster/database/databasesoftwareschedulemanagement"
	databaseupgrade "github.com/oracle/provider-oci/internal/controller/cluster/database/databaseupgrade"
	dataguardassociation "github.com/oracle/provider-oci/internal/controller/cluster/database/dataguardassociation"
	datapatch "github.com/oracle/provider-oci/internal/controller/cluster/database/datapatch"
	dbhome "github.com/oracle/provider-oci/internal/controller/cluster/database/dbhome"
	dbnode "github.com/oracle/provider-oci/internal/controller/cluster/database/dbnode"
	dbnodeconsoleconnection "github.com/oracle/provider-oci/internal/controller/cluster/database/dbnodeconsoleconnection"
	dbnodeconsolehistory "github.com/oracle/provider-oci/internal/controller/cluster/database/dbnodeconsolehistory"
	dbnodesnapshot "github.com/oracle/provider-oci/internal/controller/cluster/database/dbnodesnapshot"
	dbnodesnapshotmanagement "github.com/oracle/provider-oci/internal/controller/cluster/database/dbnodesnapshotmanagement"
	dbsystem "github.com/oracle/provider-oci/internal/controller/cluster/database/dbsystem"
	dbsystemsupgrade "github.com/oracle/provider-oci/internal/controller/cluster/database/dbsystemsupgrade"
	exadatainfrastructure "github.com/oracle/provider-oci/internal/controller/cluster/database/exadatainfrastructure"
	exadatainfrastructurecompute "github.com/oracle/provider-oci/internal/controller/cluster/database/exadatainfrastructurecompute"
	exadatainfrastructureconfigureexascalemanagement "github.com/oracle/provider-oci/internal/controller/cluster/database/exadatainfrastructureconfigureexascalemanagement"
	exadatainfrastructurestorage "github.com/oracle/provider-oci/internal/controller/cluster/database/exadatainfrastructurestorage"
	exadataiormconfig "github.com/oracle/provider-oci/internal/controller/cluster/database/exadataiormconfig"
	exadbvmcluster "github.com/oracle/provider-oci/internal/controller/cluster/database/exadbvmcluster"
	exascaledbstoragevault "github.com/oracle/provider-oci/internal/controller/cluster/database/exascaledbstoragevault"
	executionaction "github.com/oracle/provider-oci/internal/controller/cluster/database/executionaction"
	executionwindow "github.com/oracle/provider-oci/internal/controller/cluster/database/executionwindow"
	externalcontainerdatabase "github.com/oracle/provider-oci/internal/controller/cluster/database/externalcontainerdatabase"
	externalcontainerdatabasemanagement "github.com/oracle/provider-oci/internal/controller/cluster/database/externalcontainerdatabasemanagement"
	externalcontainerdatabasesstackmonitoring "github.com/oracle/provider-oci/internal/controller/cluster/database/externalcontainerdatabasesstackmonitoring"
	externaldatabaseconnector "github.com/oracle/provider-oci/internal/controller/cluster/database/externaldatabaseconnector"
	externalnoncontainerdatabase "github.com/oracle/provider-oci/internal/controller/cluster/database/externalnoncontainerdatabase"
	externalnoncontainerdatabasemanagement "github.com/oracle/provider-oci/internal/controller/cluster/database/externalnoncontainerdatabasemanagement"
	externalnoncontainerdatabaseoperationsinsightsmanagement "github.com/oracle/provider-oci/internal/controller/cluster/database/externalnoncontainerdatabaseoperationsinsightsmanagement"
	externalnoncontainerdatabasesstackmonitoring "github.com/oracle/provider-oci/internal/controller/cluster/database/externalnoncontainerdatabasesstackmonitoring"
	externalpluggabledatabase "github.com/oracle/provider-oci/internal/controller/cluster/database/externalpluggabledatabase"
	externalpluggabledatabasemanagement "github.com/oracle/provider-oci/internal/controller/cluster/database/externalpluggabledatabasemanagement"
	externalpluggabledatabaseoperationsinsightsmanagement "github.com/oracle/provider-oci/internal/controller/cluster/database/externalpluggabledatabaseoperationsinsightsmanagement"
	externalpluggabledatabasesstackmonitoring "github.com/oracle/provider-oci/internal/controller/cluster/database/externalpluggabledatabasesstackmonitoring"
	keystore "github.com/oracle/provider-oci/internal/controller/cluster/database/keystore"
	maintenancerun "github.com/oracle/provider-oci/internal/controller/cluster/database/maintenancerun"
	managementautonomousdatabasedbmfeaturesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/database/managementautonomousdatabasedbmfeaturesmanagement"
	managementcloudasm "github.com/oracle/provider-oci/internal/controller/cluster/database/managementcloudasm"
	managementcloudasminstance "github.com/oracle/provider-oci/internal/controller/cluster/database/managementcloudasminstance"
	managementcloudcluster "github.com/oracle/provider-oci/internal/controller/cluster/database/managementcloudcluster"
	managementcloudclusterinstance "github.com/oracle/provider-oci/internal/controller/cluster/database/managementcloudclusterinstance"
	managementclouddbhome "github.com/oracle/provider-oci/internal/controller/cluster/database/managementclouddbhome"
	managementclouddbnode "github.com/oracle/provider-oci/internal/controller/cluster/database/managementclouddbnode"
	managementclouddbsystem "github.com/oracle/provider-oci/internal/controller/cluster/database/managementclouddbsystem"
	managementclouddbsystemclouddatabasemanagementsmanagement "github.com/oracle/provider-oci/internal/controller/cluster/database/managementclouddbsystemclouddatabasemanagementsmanagement"
	managementclouddbsystemcloudstackmonitoringsmanagement "github.com/oracle/provider-oci/internal/controller/cluster/database/managementclouddbsystemcloudstackmonitoringsmanagement"
	managementclouddbsystemconnector "github.com/oracle/provider-oci/internal/controller/cluster/database/managementclouddbsystemconnector"
	managementclouddbsystemdiscovery "github.com/oracle/provider-oci/internal/controller/cluster/database/managementclouddbsystemdiscovery"
	managementcloudexadatainfrastructure "github.com/oracle/provider-oci/internal/controller/cluster/database/managementcloudexadatainfrastructure"
	managementcloudexadatainfrastructuremanagedexadata "github.com/oracle/provider-oci/internal/controller/cluster/database/managementcloudexadatainfrastructuremanagedexadata"
	managementcloudexadatastorageconnector "github.com/oracle/provider-oci/internal/controller/cluster/database/managementcloudexadatastorageconnector"
	managementcloudexadatastoragegrid "github.com/oracle/provider-oci/internal/controller/cluster/database/managementcloudexadatastoragegrid"
	managementcloudexadatastorageserver "github.com/oracle/provider-oci/internal/controller/cluster/database/managementcloudexadatastorageserver"
	managementcloudlistener "github.com/oracle/provider-oci/internal/controller/cluster/database/managementcloudlistener"
	managementdatabasedbmfeaturesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/database/managementdatabasedbmfeaturesmanagement"
	managementdbmanagementprivateendpoint "github.com/oracle/provider-oci/internal/controller/cluster/database/managementdbmanagementprivateendpoint"
	managementexternalasm "github.com/oracle/provider-oci/internal/controller/cluster/database/managementexternalasm"
	managementexternalasminstance "github.com/oracle/provider-oci/internal/controller/cluster/database/managementexternalasminstance"
	managementexternalcluster "github.com/oracle/provider-oci/internal/controller/cluster/database/managementexternalcluster"
	managementexternalclusterinstance "github.com/oracle/provider-oci/internal/controller/cluster/database/managementexternalclusterinstance"
	managementexternalcontainerdatabasedbmfeaturesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/database/managementexternalcontainerdatabasedbmfeaturesmanagement"
	managementexternaldbhome "github.com/oracle/provider-oci/internal/controller/cluster/database/managementexternaldbhome"
	managementexternaldbnode "github.com/oracle/provider-oci/internal/controller/cluster/database/managementexternaldbnode"
	managementexternaldbsystem "github.com/oracle/provider-oci/internal/controller/cluster/database/managementexternaldbsystem"
	managementexternaldbsystemconnector "github.com/oracle/provider-oci/internal/controller/cluster/database/managementexternaldbsystemconnector"
	managementexternaldbsystemdatabasemanagementsmanagement "github.com/oracle/provider-oci/internal/controller/cluster/database/managementexternaldbsystemdatabasemanagementsmanagement"
	managementexternaldbsystemdiscovery "github.com/oracle/provider-oci/internal/controller/cluster/database/managementexternaldbsystemdiscovery"
	managementexternaldbsystemstackmonitoringsmanagement "github.com/oracle/provider-oci/internal/controller/cluster/database/managementexternaldbsystemstackmonitoringsmanagement"
	managementexternalexadatainfrastructure "github.com/oracle/provider-oci/internal/controller/cluster/database/managementexternalexadatainfrastructure"
	managementexternalexadatainfrastructureexadatamanagement "github.com/oracle/provider-oci/internal/controller/cluster/database/managementexternalexadatainfrastructureexadatamanagement"
	managementexternalexadatastorageconnector "github.com/oracle/provider-oci/internal/controller/cluster/database/managementexternalexadatastorageconnector"
	managementexternalexadatastoragegrid "github.com/oracle/provider-oci/internal/controller/cluster/database/managementexternalexadatastoragegrid"
	managementexternalexadatastorageserver "github.com/oracle/provider-oci/internal/controller/cluster/database/managementexternalexadatastorageserver"
	managementexternallistener "github.com/oracle/provider-oci/internal/controller/cluster/database/managementexternallistener"
	managementexternalmysqldatabase "github.com/oracle/provider-oci/internal/controller/cluster/database/managementexternalmysqldatabase"
	managementexternalmysqldatabaseconnector "github.com/oracle/provider-oci/internal/controller/cluster/database/managementexternalmysqldatabaseconnector"
	managementexternalmysqldatabaseexternalmysqldatabases "github.com/oracle/provider-oci/internal/controller/cluster/database/managementexternalmysqldatabaseexternalmysqldatabases"
	managementexternalnoncontainerdatabasedbmfeaturesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/database/managementexternalnoncontainerdatabasedbmfeaturesmanagement"
	managementexternalpluggabledatabasedbmfeaturesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/database/managementexternalpluggabledatabasedbmfeaturesmanagement"
	managementmanageddatabase "github.com/oracle/provider-oci/internal/controller/cluster/database/managementmanageddatabase"
	managementmanageddatabasegroup "github.com/oracle/provider-oci/internal/controller/cluster/database/managementmanageddatabasegroup"
	managementmanageddatabaseschangedatabaseparameter "github.com/oracle/provider-oci/internal/controller/cluster/database/managementmanageddatabaseschangedatabaseparameter"
	managementmanageddatabasesresetdatabaseparameter "github.com/oracle/provider-oci/internal/controller/cluster/database/managementmanageddatabasesresetdatabaseparameter"
	managementnamedcredential "github.com/oracle/provider-oci/internal/controller/cluster/database/managementnamedcredential"
	managementpluggabledatabasedbmfeaturesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/database/managementpluggabledatabasedbmfeaturesmanagement"
	migration "github.com/oracle/provider-oci/internal/controller/cluster/database/migration"
	migrationassessment "github.com/oracle/provider-oci/internal/controller/cluster/database/migrationassessment"
	migrationassessmentassessoraction "github.com/oracle/provider-oci/internal/controller/cluster/database/migrationassessmentassessoraction"
	migrationconnection "github.com/oracle/provider-oci/internal/controller/cluster/database/migrationconnection"
	migrationjob "github.com/oracle/provider-oci/internal/controller/cluster/database/migrationjob"
	migrationjobadvisorreportcheck "github.com/oracle/provider-oci/internal/controller/cluster/database/migrationjobadvisorreportcheck"
	migrationmigration "github.com/oracle/provider-oci/internal/controller/cluster/database/migrationmigration"
	oneoffpatch "github.com/oracle/provider-oci/internal/controller/cluster/database/oneoffpatch"
	pluggabledatabase "github.com/oracle/provider-oci/internal/controller/cluster/database/pluggabledatabase"
	pluggabledatabasepluggabledatabasemanagementsmanagement "github.com/oracle/provider-oci/internal/controller/cluster/database/pluggabledatabasepluggabledatabasemanagementsmanagement"
	pluggabledatabaseslocalclone "github.com/oracle/provider-oci/internal/controller/cluster/database/pluggabledatabaseslocalclone"
	pluggabledatabasesnapshot "github.com/oracle/provider-oci/internal/controller/cluster/database/pluggabledatabasesnapshot"
	pluggabledatabasesremoteclone "github.com/oracle/provider-oci/internal/controller/cluster/database/pluggabledatabasesremoteclone"
	scheduledaction "github.com/oracle/provider-oci/internal/controller/cluster/database/scheduledaction"
	schedulingplan "github.com/oracle/provider-oci/internal/controller/cluster/database/schedulingplan"
	schedulingpolicy "github.com/oracle/provider-oci/internal/controller/cluster/database/schedulingpolicy"
	schedulingpolicyschedulingwindow "github.com/oracle/provider-oci/internal/controller/cluster/database/schedulingpolicyschedulingwindow"
	toolsdatabasetoolsconnection "github.com/oracle/provider-oci/internal/controller/cluster/database/toolsdatabasetoolsconnection"
	toolsdatabasetoolsdatabaseapigatewayconfig "github.com/oracle/provider-oci/internal/controller/cluster/database/toolsdatabasetoolsdatabaseapigatewayconfig"
	toolsdatabasetoolsidentity "github.com/oracle/provider-oci/internal/controller/cluster/database/toolsdatabasetoolsidentity"
	toolsdatabasetoolsmcpserver "github.com/oracle/provider-oci/internal/controller/cluster/database/toolsdatabasetoolsmcpserver"
	toolsdatabasetoolsmcptoolset "github.com/oracle/provider-oci/internal/controller/cluster/database/toolsdatabasetoolsmcptoolset"
	toolsdatabasetoolsprivateendpoint "github.com/oracle/provider-oci/internal/controller/cluster/database/toolsdatabasetoolsprivateendpoint"
	toolsdatabasetoolssqlreport "github.com/oracle/provider-oci/internal/controller/cluster/database/toolsdatabasetoolssqlreport"
	toolsruntimedatabasetoolsapigatewayconfigpoolapispec "github.com/oracle/provider-oci/internal/controller/cluster/database/toolsruntimedatabasetoolsapigatewayconfigpoolapispec"
	toolsruntimedatabasetoolsapigatewayconfigpoolautoapispec "github.com/oracle/provider-oci/internal/controller/cluster/database/toolsruntimedatabasetoolsapigatewayconfigpoolautoapispec"
	toolsruntimedatabasetoolsconnectioncredential "github.com/oracle/provider-oci/internal/controller/cluster/database/toolsruntimedatabasetoolsconnectioncredential"
	toolsruntimedatabasetoolsconnectioncredentialexecutegrantee "github.com/oracle/provider-oci/internal/controller/cluster/database/toolsruntimedatabasetoolsconnectioncredentialexecutegrantee"
	toolsruntimedatabasetoolsconnectioncredentialpublicsynonym "github.com/oracle/provider-oci/internal/controller/cluster/database/toolsruntimedatabasetoolsconnectioncredentialpublicsynonym"
	toolsruntimedatabasetoolsconnectionpropertyset "github.com/oracle/provider-oci/internal/controller/cluster/database/toolsruntimedatabasetoolsconnectionpropertyset"
	toolsruntimedatabasetoolsdatabaseapigatewayconfigglobal "github.com/oracle/provider-oci/internal/controller/cluster/database/toolsruntimedatabasetoolsdatabaseapigatewayconfigglobal"
	toolsruntimedatabasetoolsdatabaseapigatewayconfigpool "github.com/oracle/provider-oci/internal/controller/cluster/database/toolsruntimedatabasetoolsdatabaseapigatewayconfigpool"
	vmcluster "github.com/oracle/provider-oci/internal/controller/cluster/database/vmcluster"
	vmclusteraddvirtualmachine "github.com/oracle/provider-oci/internal/controller/cluster/database/vmclusteraddvirtualmachine"
	vmclusternetwork "github.com/oracle/provider-oci/internal/controller/cluster/database/vmclusternetwork"
	vmclusterremovevirtualmachine "github.com/oracle/provider-oci/internal/controller/cluster/database/vmclusterremovevirtualmachine"
)

// Setup_database creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_database(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		advancedclusterfilesystem.Setup,
		advancedclusterfilesystemmount.Setup,
		advancedclusterfilesystemunmount.Setup,
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
		databasesoftwareschedulemanagement.Setup,
		databaseupgrade.Setup,
		dataguardassociation.Setup,
		datapatch.Setup,
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
		managementautonomousdatabasedbmfeaturesmanagement.Setup,
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
		managementcloudexadatainfrastructure.Setup,
		managementcloudexadatainfrastructuremanagedexadata.Setup,
		managementcloudexadatastorageconnector.Setup,
		managementcloudexadatastoragegrid.Setup,
		managementcloudexadatastorageserver.Setup,
		managementcloudlistener.Setup,
		managementdatabasedbmfeaturesmanagement.Setup,
		managementdbmanagementprivateendpoint.Setup,
		managementexternalasm.Setup,
		managementexternalasminstance.Setup,
		managementexternalcluster.Setup,
		managementexternalclusterinstance.Setup,
		managementexternalcontainerdatabasedbmfeaturesmanagement.Setup,
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
		managementexternalmysqldatabaseexternalmysqldatabases.Setup,
		managementexternalnoncontainerdatabasedbmfeaturesmanagement.Setup,
		managementexternalpluggabledatabasedbmfeaturesmanagement.Setup,
		managementmanageddatabase.Setup,
		managementmanageddatabasegroup.Setup,
		managementmanageddatabaseschangedatabaseparameter.Setup,
		managementmanageddatabasesresetdatabaseparameter.Setup,
		managementnamedcredential.Setup,
		managementpluggabledatabasedbmfeaturesmanagement.Setup,
		migration.Setup,
		migrationassessment.Setup,
		migrationassessmentassessoraction.Setup,
		migrationconnection.Setup,
		migrationjob.Setup,
		migrationjobadvisorreportcheck.Setup,
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
		toolsdatabasetoolsdatabaseapigatewayconfig.Setup,
		toolsdatabasetoolsidentity.Setup,
		toolsdatabasetoolsmcpserver.Setup,
		toolsdatabasetoolsmcptoolset.Setup,
		toolsdatabasetoolsprivateendpoint.Setup,
		toolsdatabasetoolssqlreport.Setup,
		toolsruntimedatabasetoolsapigatewayconfigpoolapispec.Setup,
		toolsruntimedatabasetoolsapigatewayconfigpoolautoapispec.Setup,
		toolsruntimedatabasetoolsconnectioncredential.Setup,
		toolsruntimedatabasetoolsconnectioncredentialexecutegrantee.Setup,
		toolsruntimedatabasetoolsconnectioncredentialpublicsynonym.Setup,
		toolsruntimedatabasetoolsconnectionpropertyset.Setup,
		toolsruntimedatabasetoolsdatabaseapigatewayconfigglobal.Setup,
		toolsruntimedatabasetoolsdatabaseapigatewayconfigpool.Setup,
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

// SetupGated_database creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_database(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		advancedclusterfilesystem.SetupGated,
		advancedclusterfilesystemmount.SetupGated,
		advancedclusterfilesystemunmount.SetupGated,
		applicationvip.SetupGated,
		autonomouscontainerdatabase.SetupGated,
		autonomouscontainerdatabaseaddstandby.SetupGated,
		autonomouscontainerdatabasedataguardassociation.SetupGated,
		autonomouscontainerdatabasedataguardassociationoperation.SetupGated,
		autonomouscontainerdatabasedataguardrolechange.SetupGated,
		autonomouscontainerdatabasesnapshotstandby.SetupGated,
		autonomousdatabase.SetupGated,
		autonomousdatabasebackup.SetupGated,
		autonomousdatabaseinstancewalletmanagement.SetupGated,
		autonomousdatabaseregionalwalletmanagement.SetupGated,
		autonomousdatabasesaasadminuser.SetupGated,
		autonomousdatabasesoftwareimage.SetupGated,
		autonomousdatabasewallet.SetupGated,
		autonomousexadatainfrastructure.SetupGated,
		autonomousvmcluster.SetupGated,
		autonomousvmclusterordscertificatemanagement.SetupGated,
		autonomousvmclustersslcertificatemanagement.SetupGated,
		backup.SetupGated,
		backupcancelmanagement.SetupGated,
		backupdestination.SetupGated,
		cloudautonomousvmcluster.SetupGated,
		clouddatabasemanagement.SetupGated,
		cloudexadatainfrastructure.SetupGated,
		cloudexadatainfrastructureconfigureexascalemanagement.SetupGated,
		cloudvmcluster.SetupGated,
		cloudvmclusteriormconfig.SetupGated,
		database.SetupGated,
		databasesnapshotstandby.SetupGated,
		databasesoftwareimage.SetupGated,
		databasesoftwareschedulemanagement.SetupGated,
		databaseupgrade.SetupGated,
		dataguardassociation.SetupGated,
		datapatch.SetupGated,
		dbhome.SetupGated,
		dbnode.SetupGated,
		dbnodeconsoleconnection.SetupGated,
		dbnodeconsolehistory.SetupGated,
		dbnodesnapshot.SetupGated,
		dbnodesnapshotmanagement.SetupGated,
		dbsystem.SetupGated,
		dbsystemsupgrade.SetupGated,
		exadatainfrastructure.SetupGated,
		exadatainfrastructurecompute.SetupGated,
		exadatainfrastructureconfigureexascalemanagement.SetupGated,
		exadatainfrastructurestorage.SetupGated,
		exadataiormconfig.SetupGated,
		exadbvmcluster.SetupGated,
		exascaledbstoragevault.SetupGated,
		executionaction.SetupGated,
		executionwindow.SetupGated,
		externalcontainerdatabase.SetupGated,
		externalcontainerdatabasemanagement.SetupGated,
		externalcontainerdatabasesstackmonitoring.SetupGated,
		externaldatabaseconnector.SetupGated,
		externalnoncontainerdatabase.SetupGated,
		externalnoncontainerdatabasemanagement.SetupGated,
		externalnoncontainerdatabaseoperationsinsightsmanagement.SetupGated,
		externalnoncontainerdatabasesstackmonitoring.SetupGated,
		externalpluggabledatabase.SetupGated,
		externalpluggabledatabasemanagement.SetupGated,
		externalpluggabledatabaseoperationsinsightsmanagement.SetupGated,
		externalpluggabledatabasesstackmonitoring.SetupGated,
		keystore.SetupGated,
		maintenancerun.SetupGated,
		managementautonomousdatabasedbmfeaturesmanagement.SetupGated,
		managementcloudasm.SetupGated,
		managementcloudasminstance.SetupGated,
		managementcloudcluster.SetupGated,
		managementcloudclusterinstance.SetupGated,
		managementclouddbhome.SetupGated,
		managementclouddbnode.SetupGated,
		managementclouddbsystem.SetupGated,
		managementclouddbsystemclouddatabasemanagementsmanagement.SetupGated,
		managementclouddbsystemcloudstackmonitoringsmanagement.SetupGated,
		managementclouddbsystemconnector.SetupGated,
		managementclouddbsystemdiscovery.SetupGated,
		managementcloudexadatainfrastructure.SetupGated,
		managementcloudexadatainfrastructuremanagedexadata.SetupGated,
		managementcloudexadatastorageconnector.SetupGated,
		managementcloudexadatastoragegrid.SetupGated,
		managementcloudexadatastorageserver.SetupGated,
		managementcloudlistener.SetupGated,
		managementdatabasedbmfeaturesmanagement.SetupGated,
		managementdbmanagementprivateendpoint.SetupGated,
		managementexternalasm.SetupGated,
		managementexternalasminstance.SetupGated,
		managementexternalcluster.SetupGated,
		managementexternalclusterinstance.SetupGated,
		managementexternalcontainerdatabasedbmfeaturesmanagement.SetupGated,
		managementexternaldbhome.SetupGated,
		managementexternaldbnode.SetupGated,
		managementexternaldbsystem.SetupGated,
		managementexternaldbsystemconnector.SetupGated,
		managementexternaldbsystemdatabasemanagementsmanagement.SetupGated,
		managementexternaldbsystemdiscovery.SetupGated,
		managementexternaldbsystemstackmonitoringsmanagement.SetupGated,
		managementexternalexadatainfrastructure.SetupGated,
		managementexternalexadatainfrastructureexadatamanagement.SetupGated,
		managementexternalexadatastorageconnector.SetupGated,
		managementexternalexadatastoragegrid.SetupGated,
		managementexternalexadatastorageserver.SetupGated,
		managementexternallistener.SetupGated,
		managementexternalmysqldatabase.SetupGated,
		managementexternalmysqldatabaseconnector.SetupGated,
		managementexternalmysqldatabaseexternalmysqldatabases.SetupGated,
		managementexternalnoncontainerdatabasedbmfeaturesmanagement.SetupGated,
		managementexternalpluggabledatabasedbmfeaturesmanagement.SetupGated,
		managementmanageddatabase.SetupGated,
		managementmanageddatabasegroup.SetupGated,
		managementmanageddatabaseschangedatabaseparameter.SetupGated,
		managementmanageddatabasesresetdatabaseparameter.SetupGated,
		managementnamedcredential.SetupGated,
		managementpluggabledatabasedbmfeaturesmanagement.SetupGated,
		migration.SetupGated,
		migrationassessment.SetupGated,
		migrationassessmentassessoraction.SetupGated,
		migrationconnection.SetupGated,
		migrationjob.SetupGated,
		migrationjobadvisorreportcheck.SetupGated,
		migrationmigration.SetupGated,
		oneoffpatch.SetupGated,
		pluggabledatabase.SetupGated,
		pluggabledatabasepluggabledatabasemanagementsmanagement.SetupGated,
		pluggabledatabaseslocalclone.SetupGated,
		pluggabledatabasesnapshot.SetupGated,
		pluggabledatabasesremoteclone.SetupGated,
		scheduledaction.SetupGated,
		schedulingplan.SetupGated,
		schedulingpolicy.SetupGated,
		schedulingpolicyschedulingwindow.SetupGated,
		toolsdatabasetoolsconnection.SetupGated,
		toolsdatabasetoolsdatabaseapigatewayconfig.SetupGated,
		toolsdatabasetoolsidentity.SetupGated,
		toolsdatabasetoolsmcpserver.SetupGated,
		toolsdatabasetoolsmcptoolset.SetupGated,
		toolsdatabasetoolsprivateendpoint.SetupGated,
		toolsdatabasetoolssqlreport.SetupGated,
		toolsruntimedatabasetoolsapigatewayconfigpoolapispec.SetupGated,
		toolsruntimedatabasetoolsapigatewayconfigpoolautoapispec.SetupGated,
		toolsruntimedatabasetoolsconnectioncredential.SetupGated,
		toolsruntimedatabasetoolsconnectioncredentialexecutegrantee.SetupGated,
		toolsruntimedatabasetoolsconnectioncredentialpublicsynonym.SetupGated,
		toolsruntimedatabasetoolsconnectionpropertyset.SetupGated,
		toolsruntimedatabasetoolsdatabaseapigatewayconfigglobal.SetupGated,
		toolsruntimedatabasetoolsdatabaseapigatewayconfigpool.SetupGated,
		vmcluster.SetupGated,
		vmclusteraddvirtualmachine.SetupGated,
		vmclusternetwork.SetupGated,
		vmclusterremovevirtualmachine.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

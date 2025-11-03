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
	autonomousdatabasedbmfeatmgmt "github.com/oracle/provider-oci/internal/controller/database/autonomousdatabasedbmfeatmgmt"
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
	cloudasm "github.com/oracle/provider-oci/internal/controller/database/cloudasm"
	cloudasminstance "github.com/oracle/provider-oci/internal/controller/database/cloudasminstance"
	cloudautonomousvmcluster "github.com/oracle/provider-oci/internal/controller/database/cloudautonomousvmcluster"
	cloudcluster "github.com/oracle/provider-oci/internal/controller/database/cloudcluster"
	cloudclusterinstance "github.com/oracle/provider-oci/internal/controller/database/cloudclusterinstance"
	clouddatabasemgmt "github.com/oracle/provider-oci/internal/controller/database/clouddatabasemgmt"
	clouddbhome "github.com/oracle/provider-oci/internal/controller/database/clouddbhome"
	clouddbnode "github.com/oracle/provider-oci/internal/controller/database/clouddbnode"
	clouddbsystem "github.com/oracle/provider-oci/internal/controller/database/clouddbsystem"
	clouddbsystemclouddatabasemgmt "github.com/oracle/provider-oci/internal/controller/database/clouddbsystemclouddatabasemgmt"
	clouddbsystemcloudstackmonitoringmgmt "github.com/oracle/provider-oci/internal/controller/database/clouddbsystemcloudstackmonitoringmgmt"
	clouddbsystemconnector "github.com/oracle/provider-oci/internal/controller/database/clouddbsystemconnector"
	clouddbsystemdiscovery "github.com/oracle/provider-oci/internal/controller/database/clouddbsystemdiscovery"
	cloudexadatainfrastructure "github.com/oracle/provider-oci/internal/controller/database/cloudexadatainfrastructure"
	cloudlistener "github.com/oracle/provider-oci/internal/controller/database/cloudlistener"
	cloudvmcluster "github.com/oracle/provider-oci/internal/controller/database/cloudvmcluster"
	cloudvmclusteriormconfig "github.com/oracle/provider-oci/internal/controller/database/cloudvmclusteriormconfig"
	database "github.com/oracle/provider-oci/internal/controller/database/database"
	databasedbmfeatmgmt "github.com/oracle/provider-oci/internal/controller/database/databasedbmfeatmgmt"
	databasesoftwareimage "github.com/oracle/provider-oci/internal/controller/database/databasesoftwareimage"
	dbhome "github.com/oracle/provider-oci/internal/controller/database/dbhome"
	dbmgmtprivateendpoint "github.com/oracle/provider-oci/internal/controller/database/dbmgmtprivateendpoint"
	dbnode "github.com/oracle/provider-oci/internal/controller/database/dbnode"
	dbnodeconsoleconnection "github.com/oracle/provider-oci/internal/controller/database/dbnodeconsoleconnection"
	dbnodeconsolehistory "github.com/oracle/provider-oci/internal/controller/database/dbnodeconsolehistory"
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
	extcontainerdatabasemgmt "github.com/oracle/provider-oci/internal/controller/database/extcontainerdatabasemgmt"
	externalasm "github.com/oracle/provider-oci/internal/controller/database/externalasm"
	externalasminstance "github.com/oracle/provider-oci/internal/controller/database/externalasminstance"
	externalcluster "github.com/oracle/provider-oci/internal/controller/database/externalcluster"
	externalclusterinstance "github.com/oracle/provider-oci/internal/controller/database/externalclusterinstance"
	externalcontainerdatabase "github.com/oracle/provider-oci/internal/controller/database/externalcontainerdatabase"
	externalcontainerdatabaseextcontainerdbmfeatmgmt "github.com/oracle/provider-oci/internal/controller/database/externalcontainerdatabaseextcontainerdbmfeatmgmt"
	externalcontainerdatabasesstackmonitoring "github.com/oracle/provider-oci/internal/controller/database/externalcontainerdatabasesstackmonitoring"
	externaldatabaseconnector "github.com/oracle/provider-oci/internal/controller/database/externaldatabaseconnector"
	externaldbhome "github.com/oracle/provider-oci/internal/controller/database/externaldbhome"
	externaldbnode "github.com/oracle/provider-oci/internal/controller/database/externaldbnode"
	externaldbsystem "github.com/oracle/provider-oci/internal/controller/database/externaldbsystem"
	externaldbsystemconnector "github.com/oracle/provider-oci/internal/controller/database/externaldbsystemconnector"
	externaldbsystemdatabasemgmt "github.com/oracle/provider-oci/internal/controller/database/externaldbsystemdatabasemgmt"
	externaldbsystemdiscovery "github.com/oracle/provider-oci/internal/controller/database/externaldbsystemdiscovery"
	externaldbsystemstackmonitoringmgmt "github.com/oracle/provider-oci/internal/controller/database/externaldbsystemstackmonitoringmgmt"
	externalexadatainfrastructure "github.com/oracle/provider-oci/internal/controller/database/externalexadatainfrastructure"
	externalexadatainfrastructureexadatamgmt "github.com/oracle/provider-oci/internal/controller/database/externalexadatainfrastructureexadatamgmt"
	externalexadatastorageconnector "github.com/oracle/provider-oci/internal/controller/database/externalexadatastorageconnector"
	externalexadatastoragegrid "github.com/oracle/provider-oci/internal/controller/database/externalexadatastoragegrid"
	externalexadatastorageserver "github.com/oracle/provider-oci/internal/controller/database/externalexadatastorageserver"
	externallistener "github.com/oracle/provider-oci/internal/controller/database/externallistener"
	externalmysqldatabase "github.com/oracle/provider-oci/internal/controller/database/externalmysqldatabase"
	externalmysqldatabaseconnector "github.com/oracle/provider-oci/internal/controller/database/externalmysqldatabaseconnector"
	externalmysqldatabaseexternalmysqldatabasesmgmt "github.com/oracle/provider-oci/internal/controller/database/externalmysqldatabaseexternalmysqldatabasesmgmt"
	externalnoncontainerdatabase "github.com/oracle/provider-oci/internal/controller/database/externalnoncontainerdatabase"
	externalnoncontainerdatabaseextnoncontainerdbmfeatmgmt "github.com/oracle/provider-oci/internal/controller/database/externalnoncontainerdatabaseextnoncontainerdbmfeatmgmt"
	externalnoncontainerdatabaseoperationsinsightsmanagement "github.com/oracle/provider-oci/internal/controller/database/externalnoncontainerdatabaseoperationsinsightsmanagement"
	externalnoncontainerdatabasesstackmonitoring "github.com/oracle/provider-oci/internal/controller/database/externalnoncontainerdatabasesstackmonitoring"
	externalpluggabledatabase "github.com/oracle/provider-oci/internal/controller/database/externalpluggabledatabase"
	externalpluggabledatabaseextpluggabledbmfeatmgmt "github.com/oracle/provider-oci/internal/controller/database/externalpluggabledatabaseextpluggabledbmfeatmgmt"
	externalpluggabledatabaseoperationsinsightsmanagement "github.com/oracle/provider-oci/internal/controller/database/externalpluggabledatabaseoperationsinsightsmanagement"
	externalpluggabledatabasesstackmonitoring "github.com/oracle/provider-oci/internal/controller/database/externalpluggabledatabasesstackmonitoring"
	extnoncontainerdatabasemgmt "github.com/oracle/provider-oci/internal/controller/database/extnoncontainerdatabasemgmt"
	extpluggabledatabasemgmt "github.com/oracle/provider-oci/internal/controller/database/extpluggabledatabasemgmt"
	keystore "github.com/oracle/provider-oci/internal/controller/database/keystore"
	maintenancerun "github.com/oracle/provider-oci/internal/controller/database/maintenancerun"
	manageddatabase "github.com/oracle/provider-oci/internal/controller/database/manageddatabase"
	manageddatabasegroup "github.com/oracle/provider-oci/internal/controller/database/manageddatabasegroup"
	manageddatabaseschangedatabaseparameter "github.com/oracle/provider-oci/internal/controller/database/manageddatabaseschangedatabaseparameter"
	manageddatabasesresetdatabaseparameter "github.com/oracle/provider-oci/internal/controller/database/manageddatabasesresetdatabaseparameter"
	migration "github.com/oracle/provider-oci/internal/controller/database/migration"
	migrationconnection "github.com/oracle/provider-oci/internal/controller/database/migrationconnection"
	migrationjob "github.com/oracle/provider-oci/internal/controller/database/migrationjob"
	migrationmigration "github.com/oracle/provider-oci/internal/controller/database/migrationmigration"
	namedcredential "github.com/oracle/provider-oci/internal/controller/database/namedcredential"
	oneoffpatch "github.com/oracle/provider-oci/internal/controller/database/oneoffpatch"
	pluggabledatabase "github.com/oracle/provider-oci/internal/controller/database/pluggabledatabase"
	pluggabledatabasepluggabledatabasedbmfeatmgmt "github.com/oracle/provider-oci/internal/controller/database/pluggabledatabasepluggabledatabasedbmfeatmgmt"
	pluggabledatabasepluggabledatabasemanagement "github.com/oracle/provider-oci/internal/controller/database/pluggabledatabasepluggabledatabasemanagement"
	pluggabledatabaseslocalclone "github.com/oracle/provider-oci/internal/controller/database/pluggabledatabaseslocalclone"
	pluggabledatabasesremoteclone "github.com/oracle/provider-oci/internal/controller/database/pluggabledatabasesremoteclone"
	scheduledaction "github.com/oracle/provider-oci/internal/controller/database/scheduledaction"
	schedulingplan "github.com/oracle/provider-oci/internal/controller/database/schedulingplan"
	schedulingpolicy "github.com/oracle/provider-oci/internal/controller/database/schedulingpolicy"
	schedulingpolicyschedulingwindow "github.com/oracle/provider-oci/internal/controller/database/schedulingpolicyschedulingwindow"
	toolsdatabasetoolsconnection "github.com/oracle/provider-oci/internal/controller/database/toolsdatabasetoolsconnection"
	toolsdatabasetoolsprivateendpoint "github.com/oracle/provider-oci/internal/controller/database/toolsdatabasetoolsprivateendpoint"
	upgrade "github.com/oracle/provider-oci/internal/controller/database/upgrade"
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
		autonomousdatabasedbmfeatmgmt.Setup,
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
		cloudasm.Setup,
		cloudasminstance.Setup,
		cloudautonomousvmcluster.Setup,
		cloudcluster.Setup,
		cloudclusterinstance.Setup,
		clouddatabasemgmt.Setup,
		clouddbhome.Setup,
		clouddbnode.Setup,
		clouddbsystem.Setup,
		clouddbsystemclouddatabasemgmt.Setup,
		clouddbsystemcloudstackmonitoringmgmt.Setup,
		clouddbsystemconnector.Setup,
		clouddbsystemdiscovery.Setup,
		cloudexadatainfrastructure.Setup,
		cloudlistener.Setup,
		cloudvmcluster.Setup,
		cloudvmclusteriormconfig.Setup,
		database.Setup,
		databasedbmfeatmgmt.Setup,
		databasesoftwareimage.Setup,
		dbhome.Setup,
		dbmgmtprivateendpoint.Setup,
		dbnode.Setup,
		dbnodeconsoleconnection.Setup,
		dbnodeconsolehistory.Setup,
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
		extcontainerdatabasemgmt.Setup,
		externalasm.Setup,
		externalasminstance.Setup,
		externalcluster.Setup,
		externalclusterinstance.Setup,
		externalcontainerdatabase.Setup,
		externalcontainerdatabaseextcontainerdbmfeatmgmt.Setup,
		externalcontainerdatabasesstackmonitoring.Setup,
		externaldatabaseconnector.Setup,
		externaldbhome.Setup,
		externaldbnode.Setup,
		externaldbsystem.Setup,
		externaldbsystemconnector.Setup,
		externaldbsystemdatabasemgmt.Setup,
		externaldbsystemdiscovery.Setup,
		externaldbsystemstackmonitoringmgmt.Setup,
		externalexadatainfrastructure.Setup,
		externalexadatainfrastructureexadatamgmt.Setup,
		externalexadatastorageconnector.Setup,
		externalexadatastoragegrid.Setup,
		externalexadatastorageserver.Setup,
		externallistener.Setup,
		externalmysqldatabase.Setup,
		externalmysqldatabaseconnector.Setup,
		externalmysqldatabaseexternalmysqldatabasesmgmt.Setup,
		externalnoncontainerdatabase.Setup,
		externalnoncontainerdatabaseextnoncontainerdbmfeatmgmt.Setup,
		externalnoncontainerdatabaseoperationsinsightsmanagement.Setup,
		externalnoncontainerdatabasesstackmonitoring.Setup,
		externalpluggabledatabase.Setup,
		externalpluggabledatabaseextpluggabledbmfeatmgmt.Setup,
		externalpluggabledatabaseoperationsinsightsmanagement.Setup,
		externalpluggabledatabasesstackmonitoring.Setup,
		extnoncontainerdatabasemgmt.Setup,
		extpluggabledatabasemgmt.Setup,
		keystore.Setup,
		maintenancerun.Setup,
		manageddatabase.Setup,
		manageddatabasegroup.Setup,
		manageddatabaseschangedatabaseparameter.Setup,
		manageddatabasesresetdatabaseparameter.Setup,
		migration.Setup,
		migrationconnection.Setup,
		migrationjob.Setup,
		migrationmigration.Setup,
		namedcredential.Setup,
		oneoffpatch.Setup,
		pluggabledatabase.Setup,
		pluggabledatabasepluggabledatabasedbmfeatmgmt.Setup,
		pluggabledatabasepluggabledatabasemanagement.Setup,
		pluggabledatabaseslocalclone.Setup,
		pluggabledatabasesremoteclone.Setup,
		scheduledaction.Setup,
		schedulingplan.Setup,
		schedulingpolicy.Setup,
		schedulingpolicyschedulingwindow.Setup,
		toolsdatabasetoolsconnection.Setup,
		toolsdatabasetoolsprivateendpoint.Setup,
		upgrade.Setup,
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

/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	knowledgebase "github.com/oracle/provider-oci/internal/controller/cluster/adm/knowledgebase"
	remediationrecipe "github.com/oracle/provider-oci/internal/controller/cluster/adm/remediationrecipe"
	remediationrun "github.com/oracle/provider-oci/internal/controller/cluster/adm/remediationrun"
	vulnerabilityaudit "github.com/oracle/provider-oci/internal/controller/cluster/adm/vulnerabilityaudit"
	aidataplatform "github.com/oracle/provider-oci/internal/controller/cluster/aidataplatform/aidataplatform"
	model "github.com/oracle/provider-oci/internal/controller/cluster/aidocument/model"
	processorjob "github.com/oracle/provider-oci/internal/controller/cluster/aidocument/processorjob"
	project "github.com/oracle/provider-oci/internal/controller/cluster/aidocument/project"
	endpoint "github.com/oracle/provider-oci/internal/controller/cluster/ailanguage/endpoint"
	job "github.com/oracle/provider-oci/internal/controller/cluster/ailanguage/job"
	modelailanguage "github.com/oracle/provider-oci/internal/controller/cluster/ailanguage/model"
	projectailanguage "github.com/oracle/provider-oci/internal/controller/cluster/ailanguage/project"
	modelaivision "github.com/oracle/provider-oci/internal/controller/cluster/aivision/model"
	projectaivision "github.com/oracle/provider-oci/internal/controller/cluster/aivision/project"
	streamgroup "github.com/oracle/provider-oci/internal/controller/cluster/aivision/streamgroup"
	streamjob "github.com/oracle/provider-oci/internal/controller/cluster/aivision/streamjob"
	streamsource "github.com/oracle/provider-oci/internal/controller/cluster/aivision/streamsource"
	visionprivateendpoint "github.com/oracle/provider-oci/internal/controller/cluster/aivision/visionprivateendpoint"
	analyticsinstance "github.com/oracle/provider-oci/internal/controller/cluster/analytics/analyticsinstance"
	analyticsinstanceprivateaccesschannel "github.com/oracle/provider-oci/internal/controller/cluster/analytics/analyticsinstanceprivateaccesschannel"
	analyticsinstanceresourcegroup "github.com/oracle/provider-oci/internal/controller/cluster/analytics/analyticsinstanceresourcegroup"
	analyticsinstancevanityurl "github.com/oracle/provider-oci/internal/controller/cluster/analytics/analyticsinstancevanityurl"
	announcementsubscription "github.com/oracle/provider-oci/internal/controller/cluster/announcementsservice/announcementsubscription"
	announcementsubscriptionsactionschangecompartment "github.com/oracle/provider-oci/internal/controller/cluster/announcementsservice/announcementsubscriptionsactionschangecompartment"
	announcementsubscriptionsfiltergroup "github.com/oracle/provider-oci/internal/controller/cluster/announcementsservice/announcementsubscriptionsfiltergroup"
	privilegedapicontrol "github.com/oracle/provider-oci/internal/controller/cluster/apiaccesscontrol/privilegedapicontrol"
	privilegedapirequest "github.com/oracle/provider-oci/internal/controller/cluster/apiaccesscontrol/privilegedapirequest"
	api "github.com/oracle/provider-oci/internal/controller/cluster/apigateway/api"
	certificate "github.com/oracle/provider-oci/internal/controller/cluster/apigateway/certificate"
	deployment "github.com/oracle/provider-oci/internal/controller/cluster/apigateway/deployment"
	gateway "github.com/oracle/provider-oci/internal/controller/cluster/apigateway/gateway"
	subscriber "github.com/oracle/provider-oci/internal/controller/cluster/apigateway/subscriber"
	usageplan "github.com/oracle/provider-oci/internal/controller/cluster/apigateway/usageplan"
	apiplatforminstance "github.com/oracle/provider-oci/internal/controller/cluster/apiplatform/apiplatforminstance"
	apmdomain "github.com/oracle/provider-oci/internal/controller/cluster/apm/apmdomain"
	config "github.com/oracle/provider-oci/internal/controller/cluster/apmconfig/config"
	datafile "github.com/oracle/provider-oci/internal/controller/cluster/apmconfig/datafile"
	dedicatedvantagepoint "github.com/oracle/provider-oci/internal/controller/cluster/apmsynthetics/dedicatedvantagepoint"
	monitor "github.com/oracle/provider-oci/internal/controller/cluster/apmsynthetics/monitor"
	onpremisevantagepoint "github.com/oracle/provider-oci/internal/controller/cluster/apmsynthetics/onpremisevantagepoint"
	onpremisevantagepointworker "github.com/oracle/provider-oci/internal/controller/cluster/apmsynthetics/onpremisevantagepointworker"
	script "github.com/oracle/provider-oci/internal/controller/cluster/apmsynthetics/script"
	scheduledquery "github.com/oracle/provider-oci/internal/controller/cluster/apmtraces/scheduledquery"
	controlmonitorpluginmanagement "github.com/oracle/provider-oci/internal/controller/cluster/appmgmt/controlmonitorpluginmanagement"
	containerconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/artifacts/containerconfiguration"
	containerimagesignature "github.com/oracle/provider-oci/internal/controller/cluster/artifacts/containerimagesignature"
	containerrepository "github.com/oracle/provider-oci/internal/controller/cluster/artifacts/containerrepository"
	genericartifact "github.com/oracle/provider-oci/internal/controller/cluster/artifacts/genericartifact"
	repository "github.com/oracle/provider-oci/internal/controller/cluster/artifacts/repository"
	configuration "github.com/oracle/provider-oci/internal/controller/cluster/audit/configuration"
	autoscalingconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/autoscaling/autoscalingconfiguration"
	bastion "github.com/oracle/provider-oci/internal/controller/cluster/bastion/bastion"
	session "github.com/oracle/provider-oci/internal/controller/cluster/bastion/session"
	batchcontext "github.com/oracle/provider-oci/internal/controller/cluster/batch/batchcontext"
	batchjobpool "github.com/oracle/provider-oci/internal/controller/cluster/batch/batchjobpool"
	batchtaskenvironment "github.com/oracle/provider-oci/internal/controller/cluster/batch/batchtaskenvironment"
	batchtaskprofile "github.com/oracle/provider-oci/internal/controller/cluster/batch/batchtaskprofile"
	autoscalingconfigurationbds "github.com/oracle/provider-oci/internal/controller/cluster/bds/autoscalingconfiguration"
	bdscapacityreport "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdscapacityreport"
	bdscapacityreservation "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdscapacityreservation"
	bdsclusteradminpasswordresetaction "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsclusteradminpasswordresetaction"
	bdsinstance "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstance"
	bdsinstanceapikey "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstanceapikey"
	bdsinstancebdscapacityreservationconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstancebdscapacityreservationconfiguration"
	bdsinstancebdscertificateconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstancebdscertificateconfiguration"
	bdsinstanceexecutebootstrapscriptaction "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstanceexecutebootstrapscriptaction"
	bdsinstanceidentityconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstanceidentityconfiguration"
	bdsinstancemetastoreconfig "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstancemetastoreconfig"
	bdsinstancenodebackup "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstancenodebackup"
	bdsinstancenodebackupconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstancenodebackupconfiguration"
	bdsinstancenodereplaceconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstancenodereplaceconfiguration"
	bdsinstanceoperationcertificatemanagementsmanagement "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstanceoperationcertificatemanagementsmanagement"
	bdsinstanceospatchaction "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstanceospatchaction"
	bdsinstancepatchaction "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstancepatchaction"
	bdsinstancereplacenodeaction "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstancereplacenodeaction"
	bdsinstanceresourceprincipalconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstanceresourceprincipalconfiguration"
	bdsinstancesoftwareupdateaction "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstancesoftwareupdateaction"
	blockchainplatform "github.com/oracle/provider-oci/internal/controller/cluster/blockchain/blockchainplatform"
	osn "github.com/oracle/provider-oci/internal/controller/cluster/blockchain/osn"
	peer "github.com/oracle/provider-oci/internal/controller/cluster/blockchain/peer"
	bootvolume "github.com/oracle/provider-oci/internal/controller/cluster/blockstorage/bootvolume"
	bootvolumebackup "github.com/oracle/provider-oci/internal/controller/cluster/blockstorage/bootvolumebackup"
	volume "github.com/oracle/provider-oci/internal/controller/cluster/blockstorage/volume"
	volumeattachment "github.com/oracle/provider-oci/internal/controller/cluster/blockstorage/volumeattachment"
	volumebackup "github.com/oracle/provider-oci/internal/controller/cluster/blockstorage/volumebackup"
	volumebackuppolicy "github.com/oracle/provider-oci/internal/controller/cluster/blockstorage/volumebackuppolicy"
	volumebackuppolicyassignment "github.com/oracle/provider-oci/internal/controller/cluster/blockstorage/volumebackuppolicyassignment"
	volumegroup "github.com/oracle/provider-oci/internal/controller/cluster/blockstorage/volumegroup"
	volumegroupbackup "github.com/oracle/provider-oci/internal/controller/cluster/blockstorage/volumegroupbackup"
	alertrule "github.com/oracle/provider-oci/internal/controller/cluster/budget/alertrule"
	budget "github.com/oracle/provider-oci/internal/controller/cluster/budget/budget"
	costalertsubscription "github.com/oracle/provider-oci/internal/controller/cluster/budget/costalertsubscription"
	costanomalyevent "github.com/oracle/provider-oci/internal/controller/cluster/budget/costanomalyevent"
	costanomalymonitor "github.com/oracle/provider-oci/internal/controller/cluster/budget/costanomalymonitor"
	costanomalymonitorcostanomalymonitorenabletogglesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/budget/costanomalymonitorcostanomalymonitorenabletogglesmanagement"
	internaloccmdemandsignal "github.com/oracle/provider-oci/internal/controller/cluster/capacitymanagement/internaloccmdemandsignal"
	internaloccmdemandsignaldelivery "github.com/oracle/provider-oci/internal/controller/cluster/capacitymanagement/internaloccmdemandsignaldelivery"
	occavailabilitycatalog "github.com/oracle/provider-oci/internal/controller/cluster/capacitymanagement/occavailabilitycatalog"
	occcapacityrequest "github.com/oracle/provider-oci/internal/controller/cluster/capacitymanagement/occcapacityrequest"
	occcustomergroup "github.com/oracle/provider-oci/internal/controller/cluster/capacitymanagement/occcustomergroup"
	occcustomergroupocccustomer "github.com/oracle/provider-oci/internal/controller/cluster/capacitymanagement/occcustomergroupocccustomer"
	occmdemandsignal "github.com/oracle/provider-oci/internal/controller/cluster/capacitymanagement/occmdemandsignal"
	occmdemandsignalitem "github.com/oracle/provider-oci/internal/controller/cluster/capacitymanagement/occmdemandsignalitem"
	cabundle "github.com/oracle/provider-oci/internal/controller/cluster/certificatesmanagement/cabundle"
	certificatecertificatesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/certificatesmanagement/certificate"
	certificateauthority "github.com/oracle/provider-oci/internal/controller/cluster/certificatesmanagement/certificateauthority"
	agent "github.com/oracle/provider-oci/internal/controller/cluster/cloudbridge/agent"
	agentdependency "github.com/oracle/provider-oci/internal/controller/cluster/cloudbridge/agentdependency"
	agentplugin "github.com/oracle/provider-oci/internal/controller/cluster/cloudbridge/agentplugin"
	asset "github.com/oracle/provider-oci/internal/controller/cluster/cloudbridge/asset"
	assetsource "github.com/oracle/provider-oci/internal/controller/cluster/cloudbridge/assetsource"
	discoveryschedule "github.com/oracle/provider-oci/internal/controller/cluster/cloudbridge/discoveryschedule"
	environment "github.com/oracle/provider-oci/internal/controller/cluster/cloudbridge/environment"
	inventory "github.com/oracle/provider-oci/internal/controller/cluster/cloudbridge/inventory"
	adhocquery "github.com/oracle/provider-oci/internal/controller/cluster/cloudguard/adhocquery"
	cloudguardconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/cloudguard/cloudguardconfiguration"
	datamaskrule "github.com/oracle/provider-oci/internal/controller/cluster/cloudguard/datamaskrule"
	datasource "github.com/oracle/provider-oci/internal/controller/cluster/cloudguard/datasource"
	detectorrecipe "github.com/oracle/provider-oci/internal/controller/cluster/cloudguard/detectorrecipe"
	managedlist "github.com/oracle/provider-oci/internal/controller/cluster/cloudguard/managedlist"
	responderrecipe "github.com/oracle/provider-oci/internal/controller/cluster/cloudguard/responderrecipe"
	savedquery "github.com/oracle/provider-oci/internal/controller/cluster/cloudguard/savedquery"
	securityrecipe "github.com/oracle/provider-oci/internal/controller/cluster/cloudguard/securityrecipe"
	securityzone "github.com/oracle/provider-oci/internal/controller/cluster/cloudguard/securityzone"
	target "github.com/oracle/provider-oci/internal/controller/cluster/cloudguard/target"
	wlpagent "github.com/oracle/provider-oci/internal/controller/cluster/cloudguard/wlpagent"
	migration "github.com/oracle/provider-oci/internal/controller/cluster/cloudmigrations/migration"
	migrationasset "github.com/oracle/provider-oci/internal/controller/cluster/cloudmigrations/migrationasset"
	migrationplan "github.com/oracle/provider-oci/internal/controller/cluster/cloudmigrations/migrationplan"
	replicationschedule "github.com/oracle/provider-oci/internal/controller/cluster/cloudmigrations/replicationschedule"
	targetasset "github.com/oracle/provider-oci/internal/controller/cluster/cloudmigrations/targetasset"
	clusterplacementgroup "github.com/oracle/provider-oci/internal/controller/cluster/clusterplacementgroups/clusterplacementgroup"
	appcataloglistingresourceversionagreement "github.com/oracle/provider-oci/internal/controller/cluster/compute/appcataloglistingresourceversionagreement"
	appcatalogsubscription "github.com/oracle/provider-oci/internal/controller/cluster/compute/appcatalogsubscription"
	clusternetwork "github.com/oracle/provider-oci/internal/controller/cluster/compute/clusternetwork"
	computecapacityreport "github.com/oracle/provider-oci/internal/controller/cluster/compute/computecapacityreport"
	computecapacityreservation "github.com/oracle/provider-oci/internal/controller/cluster/compute/computecapacityreservation"
	computecapacitytopology "github.com/oracle/provider-oci/internal/controller/cluster/compute/computecapacitytopology"
	computecluster "github.com/oracle/provider-oci/internal/controller/cluster/compute/computecluster"
	computegpumemorycluster "github.com/oracle/provider-oci/internal/controller/cluster/compute/computegpumemorycluster"
	computegpumemoryfabric "github.com/oracle/provider-oci/internal/controller/cluster/compute/computegpumemoryfabric"
	computehost "github.com/oracle/provider-oci/internal/controller/cluster/compute/computehost"
	computehostgroup "github.com/oracle/provider-oci/internal/controller/cluster/compute/computehostgroup"
	computeimagecapabilityschema "github.com/oracle/provider-oci/internal/controller/cluster/compute/computeimagecapabilityschema"
	consolehistory "github.com/oracle/provider-oci/internal/controller/cluster/compute/consolehistory"
	dedicatedvmhost "github.com/oracle/provider-oci/internal/controller/cluster/compute/dedicatedvmhost"
	image "github.com/oracle/provider-oci/internal/controller/cluster/compute/image"
	instance "github.com/oracle/provider-oci/internal/controller/cluster/compute/instance"
	instanceconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/compute/instanceconfiguration"
	instanceconsoleconnection "github.com/oracle/provider-oci/internal/controller/cluster/compute/instanceconsoleconnection"
	instancemaintenanceevent "github.com/oracle/provider-oci/internal/controller/cluster/compute/instancemaintenanceevent"
	instancepool "github.com/oracle/provider-oci/internal/controller/cluster/compute/instancepool"
	instancepoolinstance "github.com/oracle/provider-oci/internal/controller/cluster/compute/instancepoolinstance"
	shapemanagement "github.com/oracle/provider-oci/internal/controller/cluster/compute/shapemanagement"
	cccinfrastructure "github.com/oracle/provider-oci/internal/controller/cluster/computecloudatcustomer/cccinfrastructure"
	cccupgradeschedule "github.com/oracle/provider-oci/internal/controller/cluster/computecloudatcustomer/cccupgradeschedule"
	addon "github.com/oracle/provider-oci/internal/controller/cluster/containerengine/addon"
	cluster "github.com/oracle/provider-oci/internal/controller/cluster/containerengine/cluster"
	clustercompletecredentialrotationmanagement "github.com/oracle/provider-oci/internal/controller/cluster/containerengine/clustercompletecredentialrotationmanagement"
	clusterpublicapiendpointdecommissionmanager "github.com/oracle/provider-oci/internal/controller/cluster/containerengine/clusterpublicapiendpointdecommissionmanager"
	clusterstartcredentialrotationmanagement "github.com/oracle/provider-oci/internal/controller/cluster/containerengine/clusterstartcredentialrotationmanagement"
	clusterworkloadmapping "github.com/oracle/provider-oci/internal/controller/cluster/containerengine/clusterworkloadmapping"
	nodepool "github.com/oracle/provider-oci/internal/controller/cluster/containerengine/nodepool"
	virtualnodepool "github.com/oracle/provider-oci/internal/controller/cluster/containerengine/virtualnodepool"
	containerinstance "github.com/oracle/provider-oci/internal/controller/cluster/containerinstances/containerinstance"
	byoasn "github.com/oracle/provider-oci/internal/controller/cluster/core/byoasn"
	listingresourceversionagreement "github.com/oracle/provider-oci/internal/controller/cluster/core/listingresourceversionagreement"
	virtualnetwork "github.com/oracle/provider-oci/internal/controller/cluster/core/virtualnetwork"
	costalertsubscriptioncostad "github.com/oracle/provider-oci/internal/controller/cluster/costad/costalertsubscription"
	costanomalyeventcostad "github.com/oracle/provider-oci/internal/controller/cluster/costad/costanomalyevent"
	costanomalymonitorcostad "github.com/oracle/provider-oci/internal/controller/cluster/costad/costanomalymonitor"
	costanomalymonitorcostanomalymonitorenabletogglesmanagementcostad "github.com/oracle/provider-oci/internal/controller/cluster/costad/costanomalymonitorcostanomalymonitorenabletogglesmanagement"
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
	migrationdatabase "github.com/oracle/provider-oci/internal/controller/cluster/database/migration"
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
	catalog "github.com/oracle/provider-oci/internal/controller/cluster/datacatalog/catalog"
	catalogprivateendpoint "github.com/oracle/provider-oci/internal/controller/cluster/datacatalog/catalogprivateendpoint"
	connection "github.com/oracle/provider-oci/internal/controller/cluster/datacatalog/connection"
	dataasset "github.com/oracle/provider-oci/internal/controller/cluster/datacatalog/dataasset"
	metastore "github.com/oracle/provider-oci/internal/controller/cluster/datacatalog/metastore"
	infrastructure "github.com/oracle/provider-oci/internal/controller/cluster/datacc/infrastructure"
	vmclusternetworkdatacc "github.com/oracle/provider-oci/internal/controller/cluster/datacc/vmclusternetwork"
	vminstance "github.com/oracle/provider-oci/internal/controller/cluster/datacc/vminstance"
	application "github.com/oracle/provider-oci/internal/controller/cluster/dataflow/application"
	invokerun "github.com/oracle/provider-oci/internal/controller/cluster/dataflow/invokerun"
	pool "github.com/oracle/provider-oci/internal/controller/cluster/dataflow/pool"
	privateendpoint "github.com/oracle/provider-oci/internal/controller/cluster/dataflow/privateendpoint"
	runstatement "github.com/oracle/provider-oci/internal/controller/cluster/dataflow/runstatement"
	sqlendpoint "github.com/oracle/provider-oci/internal/controller/cluster/dataflow/sqlendpoint"
	workspace "github.com/oracle/provider-oci/internal/controller/cluster/dataintegration/workspace"
	workspaceapplication "github.com/oracle/provider-oci/internal/controller/cluster/dataintegration/workspaceapplication"
	workspaceapplicationpatch "github.com/oracle/provider-oci/internal/controller/cluster/dataintegration/workspaceapplicationpatch"
	workspaceapplicationschedule "github.com/oracle/provider-oci/internal/controller/cluster/dataintegration/workspaceapplicationschedule"
	workspaceapplicationtaskschedule "github.com/oracle/provider-oci/internal/controller/cluster/dataintegration/workspaceapplicationtaskschedule"
	workspaceexportrequest "github.com/oracle/provider-oci/internal/controller/cluster/dataintegration/workspaceexportrequest"
	workspacefolder "github.com/oracle/provider-oci/internal/controller/cluster/dataintegration/workspacefolder"
	workspaceimportrequest "github.com/oracle/provider-oci/internal/controller/cluster/dataintegration/workspaceimportrequest"
	workspaceproject "github.com/oracle/provider-oci/internal/controller/cluster/dataintegration/workspaceproject"
	workspacetask "github.com/oracle/provider-oci/internal/controller/cluster/dataintegration/workspacetask"
	dataset "github.com/oracle/provider-oci/internal/controller/cluster/datalabelingservice/dataset"
	addsdmcolumns "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/addsdmcolumns"
	alert "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/alert"
	alertpolicy "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/alertpolicy"
	alertpolicyrule "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/alertpolicyrule"
	attributeset "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/attributeset"
	auditarchiveretrieval "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/auditarchiveretrieval"
	auditpolicy "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/auditpolicy"
	auditpolicymanagement "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/auditpolicymanagement"
	auditprofile "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/auditprofile"
	auditprofilemanagement "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/auditprofilemanagement"
	audittrail "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/audittrail"
	audittrailmanagement "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/audittrailmanagement"
	calculateauditvolumeavailable "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/calculateauditvolumeavailable"
	calculateauditvolumecollected "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/calculateauditvolumecollected"
	comparesecurityassessment "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/comparesecurityassessment"
	compareuserassessment "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/compareuserassessment"
	databasesecurityconfig "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/databasesecurityconfig"
	databasesecurityconfigmanagement "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/databasesecurityconfigmanagement"
	datasafeconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/datasafeconfiguration"
	datasafeprivateendpoint "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/datasafeprivateendpoint"
	discoveryjob "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/discoveryjob"
	discoveryjobsresult "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/discoveryjobsresult"
	generateonpremconnectorconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/generateonpremconnectorconfiguration"
	librarymaskingformat "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/librarymaskingformat"
	maskdata "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/maskdata"
	maskingpoliciesapplydifferencetomaskingcolumns "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/maskingpoliciesapplydifferencetomaskingcolumns"
	maskingpoliciesmaskingcolumn "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/maskingpoliciesmaskingcolumn"
	maskingpolicy "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/maskingpolicy"
	maskingpolicyhealthreportmanagement "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/maskingpolicyhealthreportmanagement"
	maskingreportmanagement "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/maskingreportmanagement"
	onpremconnector "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/onpremconnector"
	report "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/report"
	reportdefinition "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/reportdefinition"
	sdmmaskingpolicydifference "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/sdmmaskingpolicydifference"
	securityassessment "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/securityassessment"
	securityassessmentcheck "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/securityassessmentcheck"
	securityassessmentfinding "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/securityassessmentfinding"
	securitypolicy "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/securitypolicy"
	securitypolicyconfig "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/securitypolicyconfig"
	securitypolicydeployment "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/securitypolicydeployment"
	securitypolicydeploymentmanagement "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/securitypolicydeploymentmanagement"
	securitypolicymanagement "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/securitypolicymanagement"
	sensitivedatamodel "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/sensitivedatamodel"
	sensitivedatamodelreferentialrelation "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/sensitivedatamodelreferentialrelation"
	sensitivedatamodelsapplydiscoveryjobresults "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/sensitivedatamodelsapplydiscoveryjobresults"
	sensitivedatamodelssensitivecolumn "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/sensitivedatamodelssensitivecolumn"
	sensitivetype "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/sensitivetype"
	sensitivetypegroup "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/sensitivetypegroup"
	sensitivetypegroupgroupedsensitivetype "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/sensitivetypegroupgroupedsensitivetype"
	sensitivetypesexport "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/sensitivetypesexport"
	setsecurityassessmentbaseline "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/setsecurityassessmentbaseline"
	setsecurityassessmentbaselinemanagement "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/setsecurityassessmentbaselinemanagement"
	setuserassessmentbaseline "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/setuserassessmentbaseline"
	setuserassessmentbaselinemanagement "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/setuserassessmentbaselinemanagement"
	sqlcollection "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/sqlcollection"
	sqlfirewallpolicy "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/sqlfirewallpolicy"
	sqlfirewallpolicymanagement "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/sqlfirewallpolicymanagement"
	targetalertpolicyassociation "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/targetalertpolicyassociation"
	targetdatabase "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/targetdatabase"
	targetdatabasegroup "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/targetdatabasegroup"
	targetdatabasepeertargetdatabase "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/targetdatabasepeertargetdatabase"
	unifiedauditpolicy "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/unifiedauditpolicy"
	unifiedauditpolicydefinition "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/unifiedauditpolicydefinition"
	unsetsecurityassessmentbaseline "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/unsetsecurityassessmentbaseline"
	unsetsecurityassessmentbaselinemanagement "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/unsetsecurityassessmentbaselinemanagement"
	unsetuserassessmentbaseline "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/unsetuserassessmentbaseline"
	unsetuserassessmentbaselinemanagement "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/unsetuserassessmentbaselinemanagement"
	userassessment "github.com/oracle/provider-oci/internal/controller/cluster/datasafe/userassessment"
	computetarget "github.com/oracle/provider-oci/internal/controller/cluster/datascience/computetarget"
	jobdatascience "github.com/oracle/provider-oci/internal/controller/cluster/datascience/job"
	jobrun "github.com/oracle/provider-oci/internal/controller/cluster/datascience/jobrun"
	mlapplication "github.com/oracle/provider-oci/internal/controller/cluster/datascience/mlapplication"
	mlapplicationimplementation "github.com/oracle/provider-oci/internal/controller/cluster/datascience/mlapplicationimplementation"
	mlapplicationinstance "github.com/oracle/provider-oci/internal/controller/cluster/datascience/mlapplicationinstance"
	modeldatascience "github.com/oracle/provider-oci/internal/controller/cluster/datascience/model"
	modelartifactexport "github.com/oracle/provider-oci/internal/controller/cluster/datascience/modelartifactexport"
	modelartifactimport "github.com/oracle/provider-oci/internal/controller/cluster/datascience/modelartifactimport"
	modelcustommetadataartifact "github.com/oracle/provider-oci/internal/controller/cluster/datascience/modelcustommetadataartifact"
	modeldefinedmetadataartifact "github.com/oracle/provider-oci/internal/controller/cluster/datascience/modeldefinedmetadataartifact"
	modeldeployment "github.com/oracle/provider-oci/internal/controller/cluster/datascience/modeldeployment"
	modelgroup "github.com/oracle/provider-oci/internal/controller/cluster/datascience/modelgroup"
	modelgroupartifact "github.com/oracle/provider-oci/internal/controller/cluster/datascience/modelgroupartifact"
	modelgroupversionhistory "github.com/oracle/provider-oci/internal/controller/cluster/datascience/modelgroupversionhistory"
	modelprovenance "github.com/oracle/provider-oci/internal/controller/cluster/datascience/modelprovenance"
	modelversionset "github.com/oracle/provider-oci/internal/controller/cluster/datascience/modelversionset"
	notebooksession "github.com/oracle/provider-oci/internal/controller/cluster/datascience/notebooksession"
	pipeline "github.com/oracle/provider-oci/internal/controller/cluster/datascience/pipeline"
	pipelinerun "github.com/oracle/provider-oci/internal/controller/cluster/datascience/pipelinerun"
	privateendpointdatascience "github.com/oracle/provider-oci/internal/controller/cluster/datascience/privateendpoint"
	projectdatascience "github.com/oracle/provider-oci/internal/controller/cluster/datascience/project"
	schedule "github.com/oracle/provider-oci/internal/controller/cluster/datascience/schedule"
	vulnerabilityscan "github.com/oracle/provider-oci/internal/controller/cluster/dblm/vulnerabilityscan"
	multicloudresourcediscovery "github.com/oracle/provider-oci/internal/controller/cluster/dbmulticloud/multicloudresourcediscovery"
	oracledbawsidentityconnector "github.com/oracle/provider-oci/internal/controller/cluster/dbmulticloud/oracledbawsidentityconnector"
	oracledbawskey "github.com/oracle/provider-oci/internal/controller/cluster/dbmulticloud/oracledbawskey"
	oracledbazureblobcontainer "github.com/oracle/provider-oci/internal/controller/cluster/dbmulticloud/oracledbazureblobcontainer"
	oracledbazureblobmount "github.com/oracle/provider-oci/internal/controller/cluster/dbmulticloud/oracledbazureblobmount"
	oracledbazureconnector "github.com/oracle/provider-oci/internal/controller/cluster/dbmulticloud/oracledbazureconnector"
	oracledbazurevault "github.com/oracle/provider-oci/internal/controller/cluster/dbmulticloud/oracledbazurevault"
	oracledbazurevaultassociation "github.com/oracle/provider-oci/internal/controller/cluster/dbmulticloud/oracledbazurevaultassociation"
	oracledbgcpidentityconnector "github.com/oracle/provider-oci/internal/controller/cluster/dbmulticloud/oracledbgcpidentityconnector"
	oracledbgcpkeyring "github.com/oracle/provider-oci/internal/controller/cluster/dbmulticloud/oracledbgcpkeyring"
	instanceddfs "github.com/oracle/provider-oci/internal/controller/cluster/ddfs/instance"
	delegationcontrol "github.com/oracle/provider-oci/internal/controller/cluster/delegateaccesscontrol/delegationcontrol"
	delegationsubscription "github.com/oracle/provider-oci/internal/controller/cluster/delegateaccesscontrol/delegationsubscription"
	occdemandsignal "github.com/oracle/provider-oci/internal/controller/cluster/demandsignal/occdemandsignal"
	occmetricalarm "github.com/oracle/provider-oci/internal/controller/cluster/demandsignal/occmetricalarm"
	desktoppool "github.com/oracle/provider-oci/internal/controller/cluster/desktops/desktoppool"
	buildpipeline "github.com/oracle/provider-oci/internal/controller/cluster/devops/buildpipeline"
	buildpipelinestage "github.com/oracle/provider-oci/internal/controller/cluster/devops/buildpipelinestage"
	buildrun "github.com/oracle/provider-oci/internal/controller/cluster/devops/buildrun"
	connectiondevops "github.com/oracle/provider-oci/internal/controller/cluster/devops/connection"
	deployartifact "github.com/oracle/provider-oci/internal/controller/cluster/devops/deployartifact"
	deployenvironment "github.com/oracle/provider-oci/internal/controller/cluster/devops/deployenvironment"
	deploymentdevops "github.com/oracle/provider-oci/internal/controller/cluster/devops/deployment"
	deploypipeline "github.com/oracle/provider-oci/internal/controller/cluster/devops/deploypipeline"
	deploystage "github.com/oracle/provider-oci/internal/controller/cluster/devops/deploystage"
	projectdevops "github.com/oracle/provider-oci/internal/controller/cluster/devops/project"
	projectrepositorysetting "github.com/oracle/provider-oci/internal/controller/cluster/devops/projectrepositorysetting"
	repositorydevops "github.com/oracle/provider-oci/internal/controller/cluster/devops/repository"
	repositorymirror "github.com/oracle/provider-oci/internal/controller/cluster/devops/repositorymirror"
	repositoryprotectedbranchmanagement "github.com/oracle/provider-oci/internal/controller/cluster/devops/repositoryprotectedbranchmanagement"
	repositoryref "github.com/oracle/provider-oci/internal/controller/cluster/devops/repositoryref"
	repositorysetting "github.com/oracle/provider-oci/internal/controller/cluster/devops/repositorysetting"
	trigger "github.com/oracle/provider-oci/internal/controller/cluster/devops/trigger"
	stack "github.com/oracle/provider-oci/internal/controller/cluster/dif/stack"
	automaticdrconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/disasterrecovery/automaticdrconfiguration"
	drplan "github.com/oracle/provider-oci/internal/controller/cluster/disasterrecovery/drplan"
	drplanexecution "github.com/oracle/provider-oci/internal/controller/cluster/disasterrecovery/drplanexecution"
	drprotectiongroup "github.com/oracle/provider-oci/internal/controller/cluster/disasterrecovery/drprotectiongroup"
	distributedautonomousdatabase "github.com/oracle/provider-oci/internal/controller/cluster/distributeddatabase/distributedautonomousdatabase"
	distributeddatabase "github.com/oracle/provider-oci/internal/controller/cluster/distributeddatabase/distributeddatabase"
	distributeddatabaseprivateendpoint "github.com/oracle/provider-oci/internal/controller/cluster/distributeddatabase/distributeddatabaseprivateendpoint"
	actioncreatezonefromzonefile "github.com/oracle/provider-oci/internal/controller/cluster/dns/actioncreatezonefromzonefile"
	record "github.com/oracle/provider-oci/internal/controller/cluster/dns/record"
	resolver "github.com/oracle/provider-oci/internal/controller/cluster/dns/resolver"
	resolverendpoint "github.com/oracle/provider-oci/internal/controller/cluster/dns/resolverendpoint"
	rrset "github.com/oracle/provider-oci/internal/controller/cluster/dns/rrset"
	steeringpolicy "github.com/oracle/provider-oci/internal/controller/cluster/dns/steeringpolicy"
	steeringpolicyattachment "github.com/oracle/provider-oci/internal/controller/cluster/dns/steeringpolicyattachment"
	tsigkey "github.com/oracle/provider-oci/internal/controller/cluster/dns/tsigkey"
	view "github.com/oracle/provider-oci/internal/controller/cluster/dns/view"
	zone "github.com/oracle/provider-oci/internal/controller/cluster/dns/zone"
	zonepromotednsseckeyversion "github.com/oracle/provider-oci/internal/controller/cluster/dns/zonepromotednsseckeyversion"
	zonestagednsseckeyversion "github.com/oracle/provider-oci/internal/controller/cluster/dns/zonestagednsseckeyversion"
	dkim "github.com/oracle/provider-oci/internal/controller/cluster/email/dkim"
	emaildomain "github.com/oracle/provider-oci/internal/controller/cluster/email/emaildomain"
	emailippool "github.com/oracle/provider-oci/internal/controller/cluster/email/emailippool"
	emailreturnpath "github.com/oracle/provider-oci/internal/controller/cluster/email/emailreturnpath"
	sender "github.com/oracle/provider-oci/internal/controller/cluster/email/sender"
	suppression "github.com/oracle/provider-oci/internal/controller/cluster/email/suppression"
	rule "github.com/oracle/provider-oci/internal/controller/cluster/events/rule"
	export "github.com/oracle/provider-oci/internal/controller/cluster/filestorage/export"
	exportset "github.com/oracle/provider-oci/internal/controller/cluster/filestorage/exportset"
	filesystem "github.com/oracle/provider-oci/internal/controller/cluster/filestorage/filesystem"
	filesystemquotarule "github.com/oracle/provider-oci/internal/controller/cluster/filestorage/filesystemquotarule"
	filesystemsnapshotpolicy "github.com/oracle/provider-oci/internal/controller/cluster/filestorage/filesystemsnapshotpolicy"
	mounttarget "github.com/oracle/provider-oci/internal/controller/cluster/filestorage/mounttarget"
	outboundconnector "github.com/oracle/provider-oci/internal/controller/cluster/filestorage/outboundconnector"
	replication "github.com/oracle/provider-oci/internal/controller/cluster/filestorage/replication"
	snapshot "github.com/oracle/provider-oci/internal/controller/cluster/filestorage/snapshot"
	catalogitem "github.com/oracle/provider-oci/internal/controller/cluster/fleetappsmanagement/catalogitem"
	compliancepolicyrule "github.com/oracle/provider-oci/internal/controller/cluster/fleetappsmanagement/compliancepolicyrule"
	fleet "github.com/oracle/provider-oci/internal/controller/cluster/fleetappsmanagement/fleet"
	fleetcredential "github.com/oracle/provider-oci/internal/controller/cluster/fleetappsmanagement/fleetcredential"
	fleetproperty "github.com/oracle/provider-oci/internal/controller/cluster/fleetappsmanagement/fleetproperty"
	fleetresource "github.com/oracle/provider-oci/internal/controller/cluster/fleetappsmanagement/fleetresource"
	maintenancewindow "github.com/oracle/provider-oci/internal/controller/cluster/fleetappsmanagement/maintenancewindow"
	onboarding "github.com/oracle/provider-oci/internal/controller/cluster/fleetappsmanagement/onboarding"
	patch "github.com/oracle/provider-oci/internal/controller/cluster/fleetappsmanagement/patch"
	platformconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/fleetappsmanagement/platformconfiguration"
	property "github.com/oracle/provider-oci/internal/controller/cluster/fleetappsmanagement/property"
	provision "github.com/oracle/provider-oci/internal/controller/cluster/fleetappsmanagement/provision"
	runbook "github.com/oracle/provider-oci/internal/controller/cluster/fleetappsmanagement/runbook"
	runbookversion "github.com/oracle/provider-oci/internal/controller/cluster/fleetappsmanagement/runbookversion"
	schedulerdefinition "github.com/oracle/provider-oci/internal/controller/cluster/fleetappsmanagement/schedulerdefinition"
	taskrecord "github.com/oracle/provider-oci/internal/controller/cluster/fleetappsmanagement/taskrecord"
	fsucollection "github.com/oracle/provider-oci/internal/controller/cluster/fleetsoftwareupdate/fsucollection"
	fsucycle "github.com/oracle/provider-oci/internal/controller/cluster/fleetsoftwareupdate/fsucycle"
	fsureadinesscheck "github.com/oracle/provider-oci/internal/controller/cluster/fleetsoftwareupdate/fsureadinesscheck"
	applicationfunctions "github.com/oracle/provider-oci/internal/controller/cluster/functions/application"
	function "github.com/oracle/provider-oci/internal/controller/cluster/functions/function"
	invokefunction "github.com/oracle/provider-oci/internal/controller/cluster/functions/invokefunction"
	fusionenvironment "github.com/oracle/provider-oci/internal/controller/cluster/fusionapps/fusionenvironment"
	fusionenvironmentadminuser "github.com/oracle/provider-oci/internal/controller/cluster/fusionapps/fusionenvironmentadminuser"
	fusionenvironmentdatamaskingactivity "github.com/oracle/provider-oci/internal/controller/cluster/fusionapps/fusionenvironmentdatamaskingactivity"
	fusionenvironmentfamily "github.com/oracle/provider-oci/internal/controller/cluster/fusionapps/fusionenvironmentfamily"
	fusionenvironmentrefreshactivity "github.com/oracle/provider-oci/internal/controller/cluster/fusionapps/fusionenvironmentrefreshactivity"
	fusionenvironmentserviceattachment "github.com/oracle/provider-oci/internal/controller/cluster/fusionapps/fusionenvironmentserviceattachment"
	gdppipeline "github.com/oracle/provider-oci/internal/controller/cluster/gdp/gdppipeline"
	agentagent "github.com/oracle/provider-oci/internal/controller/cluster/generativeai/agentagent"
	agentagentendpoint "github.com/oracle/provider-oci/internal/controller/cluster/generativeai/agentagentendpoint"
	agentdataingestionjob "github.com/oracle/provider-oci/internal/controller/cluster/generativeai/agentdataingestionjob"
	agentdatasource "github.com/oracle/provider-oci/internal/controller/cluster/generativeai/agentdatasource"
	agentknowledgebase "github.com/oracle/provider-oci/internal/controller/cluster/generativeai/agentknowledgebase"
	agentprovisionedcapacity "github.com/oracle/provider-oci/internal/controller/cluster/generativeai/agentprovisionedcapacity"
	agenttool "github.com/oracle/provider-oci/internal/controller/cluster/generativeai/agenttool"
	dedicatedaicluster "github.com/oracle/provider-oci/internal/controller/cluster/generativeai/dedicatedaicluster"
	endpointgenerativeai "github.com/oracle/provider-oci/internal/controller/cluster/generativeai/endpoint"
	generativeaiprivateendpoint "github.com/oracle/provider-oci/internal/controller/cluster/generativeai/generativeaiprivateendpoint"
	hostedapplication "github.com/oracle/provider-oci/internal/controller/cluster/generativeai/hostedapplication"
	hostedapplicationiam "github.com/oracle/provider-oci/internal/controller/cluster/generativeai/hostedapplicationiam"
	hostedapplicationstorage "github.com/oracle/provider-oci/internal/controller/cluster/generativeai/hostedapplicationstorage"
	hosteddeployment "github.com/oracle/provider-oci/internal/controller/cluster/generativeai/hosteddeployment"
	importedmodel "github.com/oracle/provider-oci/internal/controller/cluster/generativeai/importedmodel"
	modelgenerativeai "github.com/oracle/provider-oci/internal/controller/cluster/generativeai/model"
	projectgenerativeai "github.com/oracle/provider-oci/internal/controller/cluster/generativeai/project"
	semanticstore "github.com/oracle/provider-oci/internal/controller/cluster/generativeai/semanticstore"
	artifactbypath "github.com/oracle/provider-oci/internal/controller/cluster/genericartifactscontent/artifactbypath"
	connectiongoldengate "github.com/oracle/provider-oci/internal/controller/cluster/goldengate/connection"
	connectionassignment "github.com/oracle/provider-oci/internal/controller/cluster/goldengate/connectionassignment"
	databaseregistration "github.com/oracle/provider-oci/internal/controller/cluster/goldengate/databaseregistration"
	deploymentgoldengate "github.com/oracle/provider-oci/internal/controller/cluster/goldengate/deployment"
	deploymentbackup "github.com/oracle/provider-oci/internal/controller/cluster/goldengate/deploymentbackup"
	deploymentcertificate "github.com/oracle/provider-oci/internal/controller/cluster/goldengate/deploymentcertificate"
	pipelinegoldengate "github.com/oracle/provider-oci/internal/controller/cluster/goldengate/pipeline"
	httpmonitor "github.com/oracle/provider-oci/internal/controller/cluster/healthchecks/httpmonitor"
	httpprobe "github.com/oracle/provider-oci/internal/controller/cluster/healthchecks/httpprobe"
	pingmonitor "github.com/oracle/provider-oci/internal/controller/cluster/healthchecks/pingmonitor"
	pingprobe "github.com/oracle/provider-oci/internal/controller/cluster/healthchecks/pingprobe"
	apikey "github.com/oracle/provider-oci/internal/controller/cluster/identity/apikey"
	authenticationpolicy "github.com/oracle/provider-oci/internal/controller/cluster/identity/authenticationpolicy"
	authtoken "github.com/oracle/provider-oci/internal/controller/cluster/identity/authtoken"
	compartment "github.com/oracle/provider-oci/internal/controller/cluster/identity/compartment"
	customersecretkey "github.com/oracle/provider-oci/internal/controller/cluster/identity/customersecretkey"
	dbcredential "github.com/oracle/provider-oci/internal/controller/cluster/identity/dbcredential"
	domain "github.com/oracle/provider-oci/internal/controller/cluster/identity/domain"
	domainreplicationtoregion "github.com/oracle/provider-oci/internal/controller/cluster/identity/domainreplicationtoregion"
	dynamicgroup "github.com/oracle/provider-oci/internal/controller/cluster/identity/dynamicgroup"
	group "github.com/oracle/provider-oci/internal/controller/cluster/identity/group"
	identityprovider "github.com/oracle/provider-oci/internal/controller/cluster/identity/identityprovider"
	idpgroupmapping "github.com/oracle/provider-oci/internal/controller/cluster/identity/idpgroupmapping"
	importstandardtagsmanagement "github.com/oracle/provider-oci/internal/controller/cluster/identity/importstandardtagsmanagement"
	networksource "github.com/oracle/provider-oci/internal/controller/cluster/identity/networksource"
	policy "github.com/oracle/provider-oci/internal/controller/cluster/identity/policy"
	smtpcredential "github.com/oracle/provider-oci/internal/controller/cluster/identity/smtpcredential"
	tag "github.com/oracle/provider-oci/internal/controller/cluster/identity/tag"
	tagdefault "github.com/oracle/provider-oci/internal/controller/cluster/identity/tagdefault"
	tagnamespace "github.com/oracle/provider-oci/internal/controller/cluster/identity/tagnamespace"
	uipassword "github.com/oracle/provider-oci/internal/controller/cluster/identity/uipassword"
	user "github.com/oracle/provider-oci/internal/controller/cluster/identity/user"
	usercapabilitiesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/identity/usercapabilitiesmanagement"
	usergroupmembership "github.com/oracle/provider-oci/internal/controller/cluster/identity/usergroupmembership"
	generatescopedaccesstoken "github.com/oracle/provider-oci/internal/controller/cluster/identitydataplane/generatescopedaccesstoken"
	accountrecoverysetting "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/accountrecoverysetting"
	apikeyidentitydomains "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/apikey"
	app "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/app"
	approle "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/approle"
	approvalworkflow "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/approvalworkflow"
	approvalworkflowassignment "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/approvalworkflowassignment"
	approvalworkflowstep "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/approvalworkflowstep"
	authenticationfactorsetting "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/authenticationfactorsetting"
	authtokenidentitydomains "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/authtoken"
	cloudgate "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/cloudgate"
	cloudgatemapping "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/cloudgatemapping"
	cloudgateserver "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/cloudgateserver"
	condition "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/condition"
	customersecretkeyidentitydomains "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/customersecretkey"
	dynamicresourcegroup "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/dynamicresourcegroup"
	grant "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/grant"
	groupidentitydomains "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/group"
	identityproofingprovider "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/identityproofingprovider"
	identityproofingprovidertemplate "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/identityproofingprovidertemplate"
	identitypropagationtrust "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/identitypropagationtrust"
	identityprovideridentitydomains "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/identityprovider"
	identitysetting "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/identitysetting"
	kmsisetting "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/kmsisetting"
	mappedattribute "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/mappedattribute"
	myapikey "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/myapikey"
	myauthtoken "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/myauthtoken"
	mycustomersecretkey "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/mycustomersecretkey"
	myoauth2clientcredential "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/myoauth2clientcredential"
	myrequest "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/myrequest"
	mysmtpcredential "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/mysmtpcredential"
	mysupportaccount "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/mysupportaccount"
	myuserdbcredential "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/myuserdbcredential"
	networkperimeter "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/networkperimeter"
	notificationsetting "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/notificationsetting"
	oauth2clientcredential "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/oauth2clientcredential"
	oauthclientcertificate "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/oauthclientcertificate"
	oauthpartnercertificate "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/oauthpartnercertificate"
	passwordpolicy "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/passwordpolicy"
	policyidentitydomains "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/policy"
	ruleidentitydomains "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/rule"
	securityquestion "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/securityquestion"
	securityquestionsetting "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/securityquestionsetting"
	selfregistrationprofile "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/selfregistrationprofile"
	setting "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/setting"
	smtpcredentialidentitydomains "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/smtpcredential"
	socialidentityprovider "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/socialidentityprovider"
	useridentitydomains "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/user"
	userdbcredential "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/userdbcredential"
	integrationinstance "github.com/oracle/provider-oci/internal/controller/cluster/integration/integrationinstance"
	oraclemanagedcustomendpoint "github.com/oracle/provider-oci/internal/controller/cluster/integration/oraclemanagedcustomendpoint"
	privateendpointoutboundconnection "github.com/oracle/provider-oci/internal/controller/cluster/integration/privateendpointoutboundconnection"
	digitaltwinadapter "github.com/oracle/provider-oci/internal/controller/cluster/iot/digitaltwinadapter"
	digitaltwininstance "github.com/oracle/provider-oci/internal/controller/cluster/iot/digitaltwininstance"
	digitaltwininstanceinvokerawcommand "github.com/oracle/provider-oci/internal/controller/cluster/iot/digitaltwininstanceinvokerawcommand"
	digitaltwinmodel "github.com/oracle/provider-oci/internal/controller/cluster/iot/digitaltwinmodel"
	digitaltwinrelationship "github.com/oracle/provider-oci/internal/controller/cluster/iot/digitaltwinrelationship"
	iotdomain "github.com/oracle/provider-oci/internal/controller/cluster/iot/iotdomain"
	iotdomainchangedataretentionperiod "github.com/oracle/provider-oci/internal/controller/cluster/iot/iotdomainchangedataretentionperiod"
	iotdomainconfiguredataaccess "github.com/oracle/provider-oci/internal/controller/cluster/iot/iotdomainconfiguredataaccess"
	iotdomaingroup "github.com/oracle/provider-oci/internal/controller/cluster/iot/iotdomaingroup"
	iotdomaingroupconfiguredataaccess "github.com/oracle/provider-oci/internal/controller/cluster/iot/iotdomaingroupconfiguredataaccess"
	fleetjms "github.com/oracle/provider-oci/internal/controller/cluster/jms/fleet"
	fleetadvancedfeatureconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/jms/fleetadvancedfeatureconfiguration"
	fleetagentconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/jms/fleetagentconfiguration"
	jmsplugin "github.com/oracle/provider-oci/internal/controller/cluster/jms/jmsplugin"
	taskschedule "github.com/oracle/provider-oci/internal/controller/cluster/jms/taskschedule"
	javadownloadreport "github.com/oracle/provider-oci/internal/controller/cluster/jmsjavadownloads/javadownloadreport"
	javadownloadtoken "github.com/oracle/provider-oci/internal/controller/cluster/jmsjavadownloads/javadownloadtoken"
	javalicenseacceptancerecord "github.com/oracle/provider-oci/internal/controller/cluster/jmsjavadownloads/javalicenseacceptancerecord"
	analyzeapplicationsconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/jmsutils/analyzeapplicationsconfiguration"
	subscriptionacknowledgmentconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/jmsutils/subscriptionacknowledgmentconfiguration"
	ekmsprivateendpoint "github.com/oracle/provider-oci/internal/controller/cluster/kms/ekmsprivateendpoint"
	encrypteddata "github.com/oracle/provider-oci/internal/controller/cluster/kms/encrypteddata"
	generatedkey "github.com/oracle/provider-oci/internal/controller/cluster/kms/generatedkey"
	key "github.com/oracle/provider-oci/internal/controller/cluster/kms/key"
	keyversion "github.com/oracle/provider-oci/internal/controller/cluster/kms/keyversion"
	sign "github.com/oracle/provider-oci/internal/controller/cluster/kms/sign"
	vault "github.com/oracle/provider-oci/internal/controller/cluster/kms/vault"
	vaultreplication "github.com/oracle/provider-oci/internal/controller/cluster/kms/vaultreplication"
	verify "github.com/oracle/provider-oci/internal/controller/cluster/kms/verify"
	configurationlicensemanager "github.com/oracle/provider-oci/internal/controller/cluster/licensemanager/configuration"
	licenserecord "github.com/oracle/provider-oci/internal/controller/cluster/licensemanager/licenserecord"
	productlicense "github.com/oracle/provider-oci/internal/controller/cluster/licensemanager/productlicense"
	quota "github.com/oracle/provider-oci/internal/controller/cluster/limits/quota"
	backend "github.com/oracle/provider-oci/internal/controller/cluster/loadbalancer/backend"
	backendset "github.com/oracle/provider-oci/internal/controller/cluster/loadbalancer/backendset"
	certificateloadbalancer "github.com/oracle/provider-oci/internal/controller/cluster/loadbalancer/certificate"
	lbhostname "github.com/oracle/provider-oci/internal/controller/cluster/loadbalancer/lbhostname"
	listener "github.com/oracle/provider-oci/internal/controller/cluster/loadbalancer/listener"
	loadbalancer "github.com/oracle/provider-oci/internal/controller/cluster/loadbalancer/loadbalancer"
	pathrouteset "github.com/oracle/provider-oci/internal/controller/cluster/loadbalancer/pathrouteset"
	routingpolicy "github.com/oracle/provider-oci/internal/controller/cluster/loadbalancer/routingpolicy"
	ruleset "github.com/oracle/provider-oci/internal/controller/cluster/loadbalancer/ruleset"
	sslciphersuite "github.com/oracle/provider-oci/internal/controller/cluster/loadbalancer/sslciphersuite"
	loganalyticsentity "github.com/oracle/provider-oci/internal/controller/cluster/loganalytics/loganalyticsentity"
	loganalyticsentityassociationsadd "github.com/oracle/provider-oci/internal/controller/cluster/loganalytics/loganalyticsentityassociationsadd"
	loganalyticsentityassociationsremove "github.com/oracle/provider-oci/internal/controller/cluster/loganalytics/loganalyticsentityassociationsremove"
	loganalyticsentitytype "github.com/oracle/provider-oci/internal/controller/cluster/loganalytics/loganalyticsentitytype"
	loganalyticsimportcustomcontent "github.com/oracle/provider-oci/internal/controller/cluster/loganalytics/loganalyticsimportcustomcontent"
	loganalyticsloggroup "github.com/oracle/provider-oci/internal/controller/cluster/loganalytics/loganalyticsloggroup"
	loganalyticsobjectcollectionrule "github.com/oracle/provider-oci/internal/controller/cluster/loganalytics/loganalyticsobjectcollectionrule"
	loganalyticspreferencesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/loganalytics/loganalyticspreferencesmanagement"
	loganalyticsresourcecategoriesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/loganalytics/loganalyticsresourcecategoriesmanagement"
	loganalyticsunprocesseddatabucketmanagement "github.com/oracle/provider-oci/internal/controller/cluster/loganalytics/loganalyticsunprocesseddatabucketmanagement"
	namespace "github.com/oracle/provider-oci/internal/controller/cluster/loganalytics/namespace"
	namespaceassociation "github.com/oracle/provider-oci/internal/controller/cluster/loganalytics/namespaceassociation"
	namespaceingesttimerule "github.com/oracle/provider-oci/internal/controller/cluster/loganalytics/namespaceingesttimerule"
	namespaceingesttimerulesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/loganalytics/namespaceingesttimerulesmanagement"
	namespacelookup "github.com/oracle/provider-oci/internal/controller/cluster/loganalytics/namespacelookup"
	namespacelookupsappenddatamanagement "github.com/oracle/provider-oci/internal/controller/cluster/loganalytics/namespacelookupsappenddatamanagement"
	namespacelookupsupdatedatamanagement "github.com/oracle/provider-oci/internal/controller/cluster/loganalytics/namespacelookupsupdatedatamanagement"
	namespacescheduledtask "github.com/oracle/provider-oci/internal/controller/cluster/loganalytics/namespacescheduledtask"
	namespacestoragearchivalconfig "github.com/oracle/provider-oci/internal/controller/cluster/loganalytics/namespacestoragearchivalconfig"
	namespacestorageenabledisablearchiving "github.com/oracle/provider-oci/internal/controller/cluster/loganalytics/namespacestorageenabledisablearchiving"
	log "github.com/oracle/provider-oci/internal/controller/cluster/logging/log"
	loggroup "github.com/oracle/provider-oci/internal/controller/cluster/logging/loggroup"
	logsavedsearch "github.com/oracle/provider-oci/internal/controller/cluster/logging/logsavedsearch"
	unifiedagentconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/logging/unifiedagentconfiguration"
	lustrefilesystem "github.com/oracle/provider-oci/internal/controller/cluster/lustrefilestorage/lustrefilesystem"
	objectstoragelink "github.com/oracle/provider-oci/internal/controller/cluster/lustrefilestorage/objectstoragelink"
	kafkacluster "github.com/oracle/provider-oci/internal/controller/cluster/managedkafka/kafkacluster"
	kafkaclusteraddon "github.com/oracle/provider-oci/internal/controller/cluster/managedkafka/kafkaclusteraddon"
	kafkaclusterconfig "github.com/oracle/provider-oci/internal/controller/cluster/managedkafka/kafkaclusterconfig"
	kafkaclustersuperusersmanagement "github.com/oracle/provider-oci/internal/controller/cluster/managedkafka/kafkaclustersuperusersmanagement"
	managementagent "github.com/oracle/provider-oci/internal/controller/cluster/managementagent/managementagent"
	managementagentdatasource "github.com/oracle/provider-oci/internal/controller/cluster/managementagent/managementagentdatasource"
	managementagentinstallkey "github.com/oracle/provider-oci/internal/controller/cluster/managementagent/managementagentinstallkey"
	namedcredential "github.com/oracle/provider-oci/internal/controller/cluster/managementagent/namedcredential"
	managementdashboardsimport "github.com/oracle/provider-oci/internal/controller/cluster/managementdashboard/managementdashboardsimport"
	managementsavedsearch "github.com/oracle/provider-oci/internal/controller/cluster/managementdashboard/managementsavedsearch"
	acceptedagreement "github.com/oracle/provider-oci/internal/controller/cluster/marketplace/acceptedagreement"
	listingpackageagreement "github.com/oracle/provider-oci/internal/controller/cluster/marketplace/listingpackageagreement"
	marketplaceexternalattestedmetadata "github.com/oracle/provider-oci/internal/controller/cluster/marketplace/marketplaceexternalattestedmetadata"
	publication "github.com/oracle/provider-oci/internal/controller/cluster/marketplace/publication"
	mediaasset "github.com/oracle/provider-oci/internal/controller/cluster/mediaservices/mediaasset"
	mediaworkflow "github.com/oracle/provider-oci/internal/controller/cluster/mediaservices/mediaworkflow"
	mediaworkflowconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/mediaservices/mediaworkflowconfiguration"
	mediaworkflowjob "github.com/oracle/provider-oci/internal/controller/cluster/mediaservices/mediaworkflowjob"
	streamcdnconfig "github.com/oracle/provider-oci/internal/controller/cluster/mediaservices/streamcdnconfig"
	streamdistributionchannel "github.com/oracle/provider-oci/internal/controller/cluster/mediaservices/streamdistributionchannel"
	streampackagingconfig "github.com/oracle/provider-oci/internal/controller/cluster/mediaservices/streampackagingconfig"
	customtable "github.com/oracle/provider-oci/internal/controller/cluster/meteringcomputation/customtable"
	query "github.com/oracle/provider-oci/internal/controller/cluster/meteringcomputation/query"
	schedulemeteringcomputation "github.com/oracle/provider-oci/internal/controller/cluster/meteringcomputation/schedule"
	usage "github.com/oracle/provider-oci/internal/controller/cluster/meteringcomputation/usage"
	usagecarbonemission "github.com/oracle/provider-oci/internal/controller/cluster/meteringcomputation/usagecarbonemission"
	usagecarbonemissionsquery "github.com/oracle/provider-oci/internal/controller/cluster/meteringcomputation/usagecarbonemissionsquery"
	usagestatementemailrecipientsgroup "github.com/oracle/provider-oci/internal/controller/cluster/meteringcomputation/usagestatementemailrecipientsgroup"
	alarm "github.com/oracle/provider-oci/internal/controller/cluster/monitoring/alarm"
	alarmsuppression "github.com/oracle/provider-oci/internal/controller/cluster/monitoring/alarmsuppression"
	capturefilter "github.com/oracle/provider-oci/internal/controller/cluster/monitoring/capturefilter"
	vtap "github.com/oracle/provider-oci/internal/controller/cluster/monitoring/vtap"
	bluegreendeployment "github.com/oracle/provider-oci/internal/controller/cluster/mysql/bluegreendeployment"
	mysqlbackup "github.com/oracle/provider-oci/internal/controller/cluster/mysql/mysqlbackup"
	mysqlchannel "github.com/oracle/provider-oci/internal/controller/cluster/mysql/mysqlchannel"
	mysqlconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/mysql/mysqlconfiguration"
	mysqldbsystem "github.com/oracle/provider-oci/internal/controller/cluster/mysql/mysqldbsystem"
	mysqlheatwavecluster "github.com/oracle/provider-oci/internal/controller/cluster/mysql/mysqlheatwavecluster"
	mysqlreplica "github.com/oracle/provider-oci/internal/controller/cluster/mysql/mysqlreplica"
	cpe "github.com/oracle/provider-oci/internal/controller/cluster/networkconnectivity/cpe"
	crossconnect "github.com/oracle/provider-oci/internal/controller/cluster/networkconnectivity/crossconnect"
	crossconnectgroup "github.com/oracle/provider-oci/internal/controller/cluster/networkconnectivity/crossconnectgroup"
	drg "github.com/oracle/provider-oci/internal/controller/cluster/networkconnectivity/drg"
	drgattachment "github.com/oracle/provider-oci/internal/controller/cluster/networkconnectivity/drgattachment"
	drgattachmentmanagement "github.com/oracle/provider-oci/internal/controller/cluster/networkconnectivity/drgattachmentmanagement"
	drgattachmentslist "github.com/oracle/provider-oci/internal/controller/cluster/networkconnectivity/drgattachmentslist"
	drgroutedistribution "github.com/oracle/provider-oci/internal/controller/cluster/networkconnectivity/drgroutedistribution"
	drgroutedistributionstatement "github.com/oracle/provider-oci/internal/controller/cluster/networkconnectivity/drgroutedistributionstatement"
	drgroutetable "github.com/oracle/provider-oci/internal/controller/cluster/networkconnectivity/drgroutetable"
	drgroutetablerouterule "github.com/oracle/provider-oci/internal/controller/cluster/networkconnectivity/drgroutetablerouterule"
	ipsec "github.com/oracle/provider-oci/internal/controller/cluster/networkconnectivity/ipsec"
	ipsecconnectiontunnelmanagement "github.com/oracle/provider-oci/internal/controller/cluster/networkconnectivity/ipsecconnectiontunnelmanagement"
	virtualcircuit "github.com/oracle/provider-oci/internal/controller/cluster/networkconnectivity/virtualcircuit"
	networkfirewall "github.com/oracle/provider-oci/internal/controller/cluster/networkfirewall/networkfirewall"
	networkfirewallpolicy "github.com/oracle/provider-oci/internal/controller/cluster/networkfirewall/networkfirewallpolicy"
	networkfirewallpolicyaddresslist "github.com/oracle/provider-oci/internal/controller/cluster/networkfirewall/networkfirewallpolicyaddresslist"
	networkfirewallpolicyapplication "github.com/oracle/provider-oci/internal/controller/cluster/networkfirewall/networkfirewallpolicyapplication"
	networkfirewallpolicyapplicationgroup "github.com/oracle/provider-oci/internal/controller/cluster/networkfirewall/networkfirewallpolicyapplicationgroup"
	networkfirewallpolicydecryptionprofile "github.com/oracle/provider-oci/internal/controller/cluster/networkfirewall/networkfirewallpolicydecryptionprofile"
	networkfirewallpolicydecryptionrule "github.com/oracle/provider-oci/internal/controller/cluster/networkfirewall/networkfirewallpolicydecryptionrule"
	networkfirewallpolicymappedsecret "github.com/oracle/provider-oci/internal/controller/cluster/networkfirewall/networkfirewallpolicymappedsecret"
	networkfirewallpolicynatrule "github.com/oracle/provider-oci/internal/controller/cluster/networkfirewall/networkfirewallpolicynatrule"
	networkfirewallpolicysecurityrule "github.com/oracle/provider-oci/internal/controller/cluster/networkfirewall/networkfirewallpolicysecurityrule"
	networkfirewallpolicyservice "github.com/oracle/provider-oci/internal/controller/cluster/networkfirewall/networkfirewallpolicyservice"
	networkfirewallpolicytunnelinspectionrule "github.com/oracle/provider-oci/internal/controller/cluster/networkfirewall/networkfirewallpolicytunnelinspectionrule"
	defaultdhcpoptions "github.com/oracle/provider-oci/internal/controller/cluster/networking/defaultdhcpoptions"
	defaultdrgroutetable "github.com/oracle/provider-oci/internal/controller/cluster/networking/defaultdrgroutetable"
	defaultroutetable "github.com/oracle/provider-oci/internal/controller/cluster/networking/defaultroutetable"
	defaultsecuritylist "github.com/oracle/provider-oci/internal/controller/cluster/networking/defaultsecuritylist"
	dhcpoptions "github.com/oracle/provider-oci/internal/controller/cluster/networking/dhcpoptions"
	internetgateway "github.com/oracle/provider-oci/internal/controller/cluster/networking/internetgateway"
	ipv6 "github.com/oracle/provider-oci/internal/controller/cluster/networking/ipv6"
	localpeeringgateway "github.com/oracle/provider-oci/internal/controller/cluster/networking/localpeeringgateway"
	natgateway "github.com/oracle/provider-oci/internal/controller/cluster/networking/natgateway"
	networksecuritygroup "github.com/oracle/provider-oci/internal/controller/cluster/networking/networksecuritygroup"
	networksecuritygroupsecurityrule "github.com/oracle/provider-oci/internal/controller/cluster/networking/networksecuritygroupsecurityrule"
	privateip "github.com/oracle/provider-oci/internal/controller/cluster/networking/privateip"
	publicip "github.com/oracle/provider-oci/internal/controller/cluster/networking/publicip"
	publicippool "github.com/oracle/provider-oci/internal/controller/cluster/networking/publicippool"
	publicippoolcapacity "github.com/oracle/provider-oci/internal/controller/cluster/networking/publicippoolcapacity"
	remotepeeringconnection "github.com/oracle/provider-oci/internal/controller/cluster/networking/remotepeeringconnection"
	routetable "github.com/oracle/provider-oci/internal/controller/cluster/networking/routetable"
	routetableattachment "github.com/oracle/provider-oci/internal/controller/cluster/networking/routetableattachment"
	securitylist "github.com/oracle/provider-oci/internal/controller/cluster/networking/securitylist"
	servicegateway "github.com/oracle/provider-oci/internal/controller/cluster/networking/servicegateway"
	subnet "github.com/oracle/provider-oci/internal/controller/cluster/networking/subnet"
	vcn "github.com/oracle/provider-oci/internal/controller/cluster/networking/vcn"
	vlan "github.com/oracle/provider-oci/internal/controller/cluster/networking/vlan"
	vnicattachment "github.com/oracle/provider-oci/internal/controller/cluster/networking/vnicattachment"
	backendnetworkloadbalancer "github.com/oracle/provider-oci/internal/controller/cluster/networkloadbalancer/backend"
	backendsetnetworkloadbalancer "github.com/oracle/provider-oci/internal/controller/cluster/networkloadbalancer/backendset"
	listenernetworkloadbalancer "github.com/oracle/provider-oci/internal/controller/cluster/networkloadbalancer/listener"
	networkloadbalancer "github.com/oracle/provider-oci/internal/controller/cluster/networkloadbalancer/networkloadbalancer"
	networkloadbalancersbackendsetsunified "github.com/oracle/provider-oci/internal/controller/cluster/networkloadbalancer/networkloadbalancersbackendsetsunified"
	configurationnosql "github.com/oracle/provider-oci/internal/controller/cluster/nosql/configuration"
	index "github.com/oracle/provider-oci/internal/controller/cluster/nosql/index"
	table "github.com/oracle/provider-oci/internal/controller/cluster/nosql/table"
	tablereplica "github.com/oracle/provider-oci/internal/controller/cluster/nosql/tablereplica"
	bucket "github.com/oracle/provider-oci/internal/controller/cluster/objectstorage/bucket"
	object "github.com/oracle/provider-oci/internal/controller/cluster/objectstorage/object"
	objectlifecyclepolicy "github.com/oracle/provider-oci/internal/controller/cluster/objectstorage/objectlifecyclepolicy"
	preauthrequest "github.com/oracle/provider-oci/internal/controller/cluster/objectstorage/preauthrequest"
	privateendpointobjectstorage "github.com/oracle/provider-oci/internal/controller/cluster/objectstorage/privateendpoint"
	replicationpolicy "github.com/oracle/provider-oci/internal/controller/cluster/objectstorage/replicationpolicy"
	oceinstance "github.com/oracle/provider-oci/internal/controller/cluster/oce/oceinstance"
	byol "github.com/oracle/provider-oci/internal/controller/cluster/ocvp/byol"
	byolallocation "github.com/oracle/provider-oci/internal/controller/cluster/ocvp/byolallocation"
	clusterocvp "github.com/oracle/provider-oci/internal/controller/cluster/ocvp/cluster"
	datastore "github.com/oracle/provider-oci/internal/controller/cluster/ocvp/datastore"
	datastorecluster "github.com/oracle/provider-oci/internal/controller/cluster/ocvp/datastorecluster"
	esxihost "github.com/oracle/provider-oci/internal/controller/cluster/ocvp/esxihost"
	managementappliance "github.com/oracle/provider-oci/internal/controller/cluster/ocvp/managementappliance"
	sddc "github.com/oracle/provider-oci/internal/controller/cluster/ocvp/sddc"
	odainstance "github.com/oracle/provider-oci/internal/controller/cluster/oda/odainstance"
	odaprivateendpoint "github.com/oracle/provider-oci/internal/controller/cluster/oda/odaprivateendpoint"
	odaprivateendpointattachment "github.com/oracle/provider-oci/internal/controller/cluster/oda/odaprivateendpointattachment"
	odaprivateendpointscanproxy "github.com/oracle/provider-oci/internal/controller/cluster/oda/odaprivateendpointscanproxy"
	notificationtopic "github.com/oracle/provider-oci/internal/controller/cluster/ons/notificationtopic"
	subscription "github.com/oracle/provider-oci/internal/controller/cluster/ons/subscription"
	opainstance "github.com/oracle/provider-oci/internal/controller/cluster/opa/opainstance"
	opensearchcluster "github.com/oracle/provider-oci/internal/controller/cluster/opensearch/opensearchcluster"
	opensearchclusterpipeline "github.com/oracle/provider-oci/internal/controller/cluster/opensearch/opensearchclusterpipeline"
	operatorcontrol "github.com/oracle/provider-oci/internal/controller/cluster/operatoraccesscontrol/operatorcontrol"
	operatorcontrolassignment "github.com/oracle/provider-oci/internal/controller/cluster/operatoraccesscontrol/operatorcontrolassignment"
	awrhub "github.com/oracle/provider-oci/internal/controller/cluster/opsi/awrhub"
	awrhubsource "github.com/oracle/provider-oci/internal/controller/cluster/opsi/awrhubsource"
	awrhubsourceawrhubsourcesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/opsi/awrhubsourceawrhubsourcesmanagement"
	chargebackplan "github.com/oracle/provider-oci/internal/controller/cluster/opsi/chargebackplan"
	databaseinsight "github.com/oracle/provider-oci/internal/controller/cluster/opsi/databaseinsight"
	enterprisemanagerbridge "github.com/oracle/provider-oci/internal/controller/cluster/opsi/enterprisemanagerbridge"
	exadatainsight "github.com/oracle/provider-oci/internal/controller/cluster/opsi/exadatainsight"
	hostinsight "github.com/oracle/provider-oci/internal/controller/cluster/opsi/hostinsight"
	newsreport "github.com/oracle/provider-oci/internal/controller/cluster/opsi/newsreport"
	operationsinsightsprivateendpoint "github.com/oracle/provider-oci/internal/controller/cluster/opsi/operationsinsightsprivateendpoint"
	operationsinsightswarehouse "github.com/oracle/provider-oci/internal/controller/cluster/opsi/operationsinsightswarehouse"
	operationsinsightswarehousedownloadwarehousewallet "github.com/oracle/provider-oci/internal/controller/cluster/opsi/operationsinsightswarehousedownloadwarehousewallet"
	operationsinsightswarehouserotatewarehousewallet "github.com/oracle/provider-oci/internal/controller/cluster/opsi/operationsinsightswarehouserotatewarehousewallet"
	operationsinsightswarehouseuser "github.com/oracle/provider-oci/internal/controller/cluster/opsi/operationsinsightswarehouseuser"
	opsiconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/opsi/opsiconfiguration"
	enrollmentstatus "github.com/oracle/provider-oci/internal/controller/cluster/optimizer/enrollmentstatus"
	profile "github.com/oracle/provider-oci/internal/controller/cluster/optimizer/profile"
	recommendation "github.com/oracle/provider-oci/internal/controller/cluster/optimizer/recommendation"
	resourceaction "github.com/oracle/provider-oci/internal/controller/cluster/optimizer/resourceaction"
	dynamicset "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/dynamicset"
	dynamicsetinstallpackagesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/dynamicsetinstallpackagesmanagement"
	dynamicsetrebootmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/dynamicsetrebootmanagement"
	dynamicsetremovepackagesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/dynamicsetremovepackagesmanagement"
	dynamicsetupdatepackagesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/dynamicsetupdatepackagesmanagement"
	event "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/event"
	lifecycleenvironment "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/lifecycleenvironment"
	lifecyclestageattachmanagedinstancesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/lifecyclestageattachmanagedinstancesmanagement"
	lifecyclestagedetachmanagedinstancesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/lifecyclestagedetachmanagedinstancesmanagement"
	lifecyclestagepromotesoftwaresourcemanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/lifecyclestagepromotesoftwaresourcemanagement"
	lifecyclestagerebootmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/lifecyclestagerebootmanagement"
	managedinstance "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managedinstance"
	managedinstanceattachprofilemanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managedinstanceattachprofilemanagement"
	managedinstanceattachsoftwaresourcesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managedinstanceattachsoftwaresourcesmanagement"
	managedinstancedetachprofilemanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managedinstancedetachprofilemanagement"
	managedinstancedetachsoftwaresourcesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managedinstancedetachsoftwaresourcesmanagement"
	managedinstancegroup "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managedinstancegroup"
	managedinstancegroupattachmanagedinstancesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managedinstancegroupattachmanagedinstancesmanagement"
	managedinstancegroupattachsoftwaresourcesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managedinstancegroupattachsoftwaresourcesmanagement"
	managedinstancegroupdetachmanagedinstancesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managedinstancegroupdetachmanagedinstancesmanagement"
	managedinstancegroupdetachsoftwaresourcesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managedinstancegroupdetachsoftwaresourcesmanagement"
	managedinstancegroupinstallpackagesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managedinstancegroupinstallpackagesmanagement"
	managedinstancegroupinstallwindowsupdatesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managedinstancegroupinstallwindowsupdatesmanagement"
	managedinstancegroupmanagemodulestreamsmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managedinstancegroupmanagemodulestreamsmanagement"
	managedinstancegrouprebootmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managedinstancegrouprebootmanagement"
	managedinstancegroupremovepackagesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managedinstancegroupremovepackagesmanagement"
	managedinstancegroupupdateallpackagesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managedinstancegroupupdateallpackagesmanagement"
	managedinstanceinstallpackagesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managedinstanceinstallpackagesmanagement"
	managedinstanceinstallsnapsmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managedinstanceinstallsnapsmanagement"
	managedinstanceinstallwindowsupdatesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managedinstanceinstallwindowsupdatesmanagement"
	managedinstancerebootmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managedinstancerebootmanagement"
	managedinstancerefreshsoftwaremanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managedinstancerefreshsoftwaremanagement"
	managedinstanceremovepackagesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managedinstanceremovepackagesmanagement"
	managedinstanceremovesnapsmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managedinstanceremovesnapsmanagement"
	managedinstancesinstallwindowsupdatesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managedinstancesinstallwindowsupdatesmanagement"
	managedinstancesupdatepackagesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managedinstancesupdatepackagesmanagement"
	managedinstanceswitchsnapchannelmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managedinstanceswitchsnapchannelmanagement"
	managedinstanceupdatepackagesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managedinstanceupdatepackagesmanagement"
	managementstation "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managementstation"
	managementstationassociatemanagedinstancesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managementstationassociatemanagedinstancesmanagement"
	managementstationmirrorsynchronizemanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managementstationmirrorsynchronizemanagement"
	managementstationrefreshmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managementstationrefreshmanagement"
	managementstationsynchronizemirrorsmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/managementstationsynchronizemirrorsmanagement"
	profileosmanagementhub "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/profile"
	profileattachlifecyclestagemanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/profileattachlifecyclestagemanagement"
	profileattachmanagedinstancegroupmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/profileattachmanagedinstancegroupmanagement"
	profileattachmanagementstationmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/profileattachmanagementstationmanagement"
	profileattachsoftwaresourcesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/profileattachsoftwaresourcesmanagement"
	profiledetachmanagementstationmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/profiledetachmanagementstationmanagement"
	profiledetachsoftwaresourcesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/profiledetachsoftwaresourcesmanagement"
	scheduledjob "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/scheduledjob"
	softwaresource "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/softwaresource"
	softwaresourceaddpackagesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/softwaresourceaddpackagesmanagement"
	softwaresourcechangeavailabilitymanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/softwaresourcechangeavailabilitymanagement"
	softwaresourcegeneratemetadatamanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/softwaresourcegeneratemetadatamanagement"
	softwaresourcemanifest "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/softwaresourcemanifest"
	softwaresourceremovepackagesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/softwaresourceremovepackagesmanagement"
	softwaresourcereplacepackagesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/softwaresourcereplacepackagesmanagement"
	workrequestrerunmanagement "github.com/oracle/provider-oci/internal/controller/cluster/osmanagementhub/workrequestrerunmanagement"
	addressactionverification "github.com/oracle/provider-oci/internal/controller/cluster/ospgateway/addressactionverification"
	subscriptionospgateway "github.com/oracle/provider-oci/internal/controller/cluster/ospgateway/subscription"
	providerconfig "github.com/oracle/provider-oci/internal/controller/cluster/providerconfig"
	privateserviceaccess "github.com/oracle/provider-oci/internal/controller/cluster/psa/privateserviceaccess"
	psqlbackup "github.com/oracle/provider-oci/internal/controller/cluster/psql/psqlbackup"
	psqlconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/psql/psqlconfiguration"
	psqldbsystem "github.com/oracle/provider-oci/internal/controller/cluster/psql/psqldbsystem"
	consumergroup "github.com/oracle/provider-oci/internal/controller/cluster/queue/consumergroup"
	queue "github.com/oracle/provider-oci/internal/controller/cluster/queue/queue"
	protecteddatabase "github.com/oracle/provider-oci/internal/controller/cluster/recovery/protecteddatabase"
	protectionpolicy "github.com/oracle/provider-oci/internal/controller/cluster/recovery/protectionpolicy"
	recoveryservicesubnet "github.com/oracle/provider-oci/internal/controller/cluster/recovery/recoveryservicesubnet"
	ocicachebackup "github.com/oracle/provider-oci/internal/controller/cluster/redis/ocicachebackup"
	ocicachebackupexporttoobjectstorage "github.com/oracle/provider-oci/internal/controller/cluster/redis/ocicachebackupexporttoobjectstorage"
	ocicacheconfigset "github.com/oracle/provider-oci/internal/controller/cluster/redis/ocicacheconfigset"
	ocicacheconfigsetlistassociatedocicachecluster "github.com/oracle/provider-oci/internal/controller/cluster/redis/ocicacheconfigsetlistassociatedocicachecluster"
	ocicacheuser "github.com/oracle/provider-oci/internal/controller/cluster/redis/ocicacheuser"
	ocicacheusergetrediscluster "github.com/oracle/provider-oci/internal/controller/cluster/redis/ocicacheusergetrediscluster"
	rediscluster "github.com/oracle/provider-oci/internal/controller/cluster/redis/rediscluster"
	redisclusterattachocicacheuser "github.com/oracle/provider-oci/internal/controller/cluster/redis/redisclusterattachocicacheuser"
	redisclustercreateidentitytoken "github.com/oracle/provider-oci/internal/controller/cluster/redis/redisclustercreateidentitytoken"
	redisclusterdetachocicacheuser "github.com/oracle/provider-oci/internal/controller/cluster/redis/redisclusterdetachocicacheuser"
	redisclustergetocicacheuser "github.com/oracle/provider-oci/internal/controller/cluster/redis/redisclustergetocicacheuser"
	monitoredregion "github.com/oracle/provider-oci/internal/controller/cluster/resourceanalytics/monitoredregion"
	resourceanalyticsinstance "github.com/oracle/provider-oci/internal/controller/cluster/resourceanalytics/resourceanalyticsinstance"
	resourceanalyticsinstanceoacmanagement "github.com/oracle/provider-oci/internal/controller/cluster/resourceanalytics/resourceanalyticsinstanceoacmanagement"
	tenancyattachment "github.com/oracle/provider-oci/internal/controller/cluster/resourceanalytics/tenancyattachment"
	privateendpointresourcemanager "github.com/oracle/provider-oci/internal/controller/cluster/resourcemanager/privateendpoint"
	scheduleresourcescheduler "github.com/oracle/provider-oci/internal/controller/cluster/resourcescheduler/schedule"
	serviceconnector "github.com/oracle/provider-oci/internal/controller/cluster/sch/serviceconnector"
	securityattribute "github.com/oracle/provider-oci/internal/controller/cluster/securityattribute/securityattribute"
	securityattributenamespace "github.com/oracle/provider-oci/internal/controller/cluster/securityattribute/securityattributenamespace"
	subscriptionself "github.com/oracle/provider-oci/internal/controller/cluster/self/subscription"
	privateapplication "github.com/oracle/provider-oci/internal/controller/cluster/servicecatalog/privateapplication"
	servicecatalog "github.com/oracle/provider-oci/internal/controller/cluster/servicecatalog/servicecatalog"
	servicecatalogassociation "github.com/oracle/provider-oci/internal/controller/cluster/servicecatalog/servicecatalogassociation"
	baselineablemetric "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/baselineablemetric"
	configstackmonitoring "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/config"
	discoveryjobstackmonitoring "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/discoveryjob"
	maintenancewindowstackmonitoring "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/maintenancewindow"
	maintenancewindowsretryfailedoperation "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/maintenancewindowsretryfailedoperation"
	maintenancewindowsstop "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/maintenancewindowsstop"
	metricextension "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/metricextension"
	metricextensionmetricextensionongivenresourcesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/metricextensionmetricextensionongivenresourcesmanagement"
	monitoredresource "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/monitoredresource"
	monitoredresourcesassociatemonitoredresource "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/monitoredresourcesassociatemonitoredresource"
	monitoredresourceslistmember "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/monitoredresourceslistmember"
	monitoredresourcessearch "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/monitoredresourcessearch"
	monitoredresourcessearchassociation "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/monitoredresourcessearchassociation"
	monitoredresourcetask "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/monitoredresourcetask"
	monitoredresourcetype "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/monitoredresourcetype"
	monitoringtemplate "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/monitoringtemplate"
	monitoringtemplatealarmcondition "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/monitoringtemplatealarmcondition"
	monitoringtemplateongivenresourcesmanagement "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/monitoringtemplateongivenresourcesmanagement"
	processset "github.com/oracle/provider-oci/internal/controller/cluster/stackmonitoring/processset"
	connectharness "github.com/oracle/provider-oci/internal/controller/cluster/streaming/connectharness"
	stream "github.com/oracle/provider-oci/internal/controller/cluster/streaming/stream"
	streampool "github.com/oracle/provider-oci/internal/controller/cluster/streaming/streampool"
	subscriptionmapping "github.com/oracle/provider-oci/internal/controller/cluster/tenantmanagercontrolplane/subscriptionmapping"
	subscriptionredeemableuser "github.com/oracle/provider-oci/internal/controller/cluster/usageproxy/subscriptionredeemableuser"
	secret "github.com/oracle/provider-oci/internal/controller/cluster/vault/secret"
	vbsinstance "github.com/oracle/provider-oci/internal/controller/cluster/vbsinst/vbsinstance"
	vbinstance "github.com/oracle/provider-oci/internal/controller/cluster/visualbuilder/vbinstance"
	pathanalysi "github.com/oracle/provider-oci/internal/controller/cluster/vnmonitoring/pathanalysi"
	containerscanrecipe "github.com/oracle/provider-oci/internal/controller/cluster/vulnerabilityscanning/containerscanrecipe"
	containerscantarget "github.com/oracle/provider-oci/internal/controller/cluster/vulnerabilityscanning/containerscantarget"
	hostscanrecipe "github.com/oracle/provider-oci/internal/controller/cluster/vulnerabilityscanning/hostscanrecipe"
	hostscantarget "github.com/oracle/provider-oci/internal/controller/cluster/vulnerabilityscanning/hostscantarget"
	webappacceleration "github.com/oracle/provider-oci/internal/controller/cluster/waa/webappacceleration"
	webappaccelerationpolicy "github.com/oracle/provider-oci/internal/controller/cluster/waa/webappaccelerationpolicy"
	addresslist "github.com/oracle/provider-oci/internal/controller/cluster/waas/addresslist"
	certificatewaas "github.com/oracle/provider-oci/internal/controller/cluster/waas/certificate"
	customprotectionrule "github.com/oracle/provider-oci/internal/controller/cluster/waas/customprotectionrule"
	httpredirect "github.com/oracle/provider-oci/internal/controller/cluster/waas/httpredirect"
	protectionrule "github.com/oracle/provider-oci/internal/controller/cluster/waas/protectionrule"
	purgecache "github.com/oracle/provider-oci/internal/controller/cluster/waas/purgecache"
	waaspolicy "github.com/oracle/provider-oci/internal/controller/cluster/waas/waaspolicy"
	networkaddresslist "github.com/oracle/provider-oci/internal/controller/cluster/waf/networkaddresslist"
	webappfirewall "github.com/oracle/provider-oci/internal/controller/cluster/waf/webappfirewall"
	webappfirewallpolicy "github.com/oracle/provider-oci/internal/controller/cluster/waf/webappfirewallpolicy"
	configurationzpr "github.com/oracle/provider-oci/internal/controller/cluster/zpr/configuration"
	zprpolicy "github.com/oracle/provider-oci/internal/controller/cluster/zpr/zprpolicy"
)

// Setup_monolith creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_monolith(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		knowledgebase.Setup,
		remediationrecipe.Setup,
		remediationrun.Setup,
		vulnerabilityaudit.Setup,
		aidataplatform.Setup,
		model.Setup,
		processorjob.Setup,
		project.Setup,
		endpoint.Setup,
		job.Setup,
		modelailanguage.Setup,
		projectailanguage.Setup,
		modelaivision.Setup,
		projectaivision.Setup,
		streamgroup.Setup,
		streamjob.Setup,
		streamsource.Setup,
		visionprivateendpoint.Setup,
		analyticsinstance.Setup,
		analyticsinstanceprivateaccesschannel.Setup,
		analyticsinstanceresourcegroup.Setup,
		analyticsinstancevanityurl.Setup,
		announcementsubscription.Setup,
		announcementsubscriptionsactionschangecompartment.Setup,
		announcementsubscriptionsfiltergroup.Setup,
		privilegedapicontrol.Setup,
		privilegedapirequest.Setup,
		api.Setup,
		certificate.Setup,
		deployment.Setup,
		gateway.Setup,
		subscriber.Setup,
		usageplan.Setup,
		apiplatforminstance.Setup,
		apmdomain.Setup,
		config.Setup,
		datafile.Setup,
		dedicatedvantagepoint.Setup,
		monitor.Setup,
		onpremisevantagepoint.Setup,
		onpremisevantagepointworker.Setup,
		script.Setup,
		scheduledquery.Setup,
		controlmonitorpluginmanagement.Setup,
		containerconfiguration.Setup,
		containerimagesignature.Setup,
		containerrepository.Setup,
		genericartifact.Setup,
		repository.Setup,
		configuration.Setup,
		autoscalingconfiguration.Setup,
		bastion.Setup,
		session.Setup,
		batchcontext.Setup,
		batchjobpool.Setup,
		batchtaskenvironment.Setup,
		batchtaskprofile.Setup,
		autoscalingconfigurationbds.Setup,
		bdscapacityreport.Setup,
		bdscapacityreservation.Setup,
		bdsclusteradminpasswordresetaction.Setup,
		bdsinstance.Setup,
		bdsinstanceapikey.Setup,
		bdsinstancebdscapacityreservationconfiguration.Setup,
		bdsinstancebdscertificateconfiguration.Setup,
		bdsinstanceexecutebootstrapscriptaction.Setup,
		bdsinstanceidentityconfiguration.Setup,
		bdsinstancemetastoreconfig.Setup,
		bdsinstancenodebackup.Setup,
		bdsinstancenodebackupconfiguration.Setup,
		bdsinstancenodereplaceconfiguration.Setup,
		bdsinstanceoperationcertificatemanagementsmanagement.Setup,
		bdsinstanceospatchaction.Setup,
		bdsinstancepatchaction.Setup,
		bdsinstancereplacenodeaction.Setup,
		bdsinstanceresourceprincipalconfiguration.Setup,
		bdsinstancesoftwareupdateaction.Setup,
		blockchainplatform.Setup,
		osn.Setup,
		peer.Setup,
		bootvolume.Setup,
		bootvolumebackup.Setup,
		volume.Setup,
		volumeattachment.Setup,
		volumebackup.Setup,
		volumebackuppolicy.Setup,
		volumebackuppolicyassignment.Setup,
		volumegroup.Setup,
		volumegroupbackup.Setup,
		alertrule.Setup,
		budget.Setup,
		costalertsubscription.Setup,
		costanomalyevent.Setup,
		costanomalymonitor.Setup,
		costanomalymonitorcostanomalymonitorenabletogglesmanagement.Setup,
		internaloccmdemandsignal.Setup,
		internaloccmdemandsignaldelivery.Setup,
		occavailabilitycatalog.Setup,
		occcapacityrequest.Setup,
		occcustomergroup.Setup,
		occcustomergroupocccustomer.Setup,
		occmdemandsignal.Setup,
		occmdemandsignalitem.Setup,
		cabundle.Setup,
		certificatecertificatesmanagement.Setup,
		certificateauthority.Setup,
		agent.Setup,
		agentdependency.Setup,
		agentplugin.Setup,
		asset.Setup,
		assetsource.Setup,
		discoveryschedule.Setup,
		environment.Setup,
		inventory.Setup,
		adhocquery.Setup,
		cloudguardconfiguration.Setup,
		datamaskrule.Setup,
		datasource.Setup,
		detectorrecipe.Setup,
		managedlist.Setup,
		responderrecipe.Setup,
		savedquery.Setup,
		securityrecipe.Setup,
		securityzone.Setup,
		target.Setup,
		wlpagent.Setup,
		migration.Setup,
		migrationasset.Setup,
		migrationplan.Setup,
		replicationschedule.Setup,
		targetasset.Setup,
		clusterplacementgroup.Setup,
		appcataloglistingresourceversionagreement.Setup,
		appcatalogsubscription.Setup,
		clusternetwork.Setup,
		computecapacityreport.Setup,
		computecapacityreservation.Setup,
		computecapacitytopology.Setup,
		computecluster.Setup,
		computegpumemorycluster.Setup,
		computegpumemoryfabric.Setup,
		computehost.Setup,
		computehostgroup.Setup,
		computeimagecapabilityschema.Setup,
		consolehistory.Setup,
		dedicatedvmhost.Setup,
		image.Setup,
		instance.Setup,
		instanceconfiguration.Setup,
		instanceconsoleconnection.Setup,
		instancemaintenanceevent.Setup,
		instancepool.Setup,
		instancepoolinstance.Setup,
		shapemanagement.Setup,
		cccinfrastructure.Setup,
		cccupgradeschedule.Setup,
		addon.Setup,
		cluster.Setup,
		clustercompletecredentialrotationmanagement.Setup,
		clusterpublicapiendpointdecommissionmanager.Setup,
		clusterstartcredentialrotationmanagement.Setup,
		clusterworkloadmapping.Setup,
		nodepool.Setup,
		virtualnodepool.Setup,
		containerinstance.Setup,
		byoasn.Setup,
		listingresourceversionagreement.Setup,
		virtualnetwork.Setup,
		costalertsubscriptioncostad.Setup,
		costanomalyeventcostad.Setup,
		costanomalymonitorcostad.Setup,
		costanomalymonitorcostanomalymonitorenabletogglesmanagementcostad.Setup,
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
		migrationdatabase.Setup,
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
		catalog.Setup,
		catalogprivateendpoint.Setup,
		connection.Setup,
		dataasset.Setup,
		metastore.Setup,
		infrastructure.Setup,
		vmclusternetworkdatacc.Setup,
		vminstance.Setup,
		application.Setup,
		invokerun.Setup,
		pool.Setup,
		privateendpoint.Setup,
		runstatement.Setup,
		sqlendpoint.Setup,
		workspace.Setup,
		workspaceapplication.Setup,
		workspaceapplicationpatch.Setup,
		workspaceapplicationschedule.Setup,
		workspaceapplicationtaskschedule.Setup,
		workspaceexportrequest.Setup,
		workspacefolder.Setup,
		workspaceimportrequest.Setup,
		workspaceproject.Setup,
		workspacetask.Setup,
		dataset.Setup,
		addsdmcolumns.Setup,
		alert.Setup,
		alertpolicy.Setup,
		alertpolicyrule.Setup,
		attributeset.Setup,
		auditarchiveretrieval.Setup,
		auditpolicy.Setup,
		auditpolicymanagement.Setup,
		auditprofile.Setup,
		auditprofilemanagement.Setup,
		audittrail.Setup,
		audittrailmanagement.Setup,
		calculateauditvolumeavailable.Setup,
		calculateauditvolumecollected.Setup,
		comparesecurityassessment.Setup,
		compareuserassessment.Setup,
		databasesecurityconfig.Setup,
		databasesecurityconfigmanagement.Setup,
		datasafeconfiguration.Setup,
		datasafeprivateendpoint.Setup,
		discoveryjob.Setup,
		discoveryjobsresult.Setup,
		generateonpremconnectorconfiguration.Setup,
		librarymaskingformat.Setup,
		maskdata.Setup,
		maskingpoliciesapplydifferencetomaskingcolumns.Setup,
		maskingpoliciesmaskingcolumn.Setup,
		maskingpolicy.Setup,
		maskingpolicyhealthreportmanagement.Setup,
		maskingreportmanagement.Setup,
		onpremconnector.Setup,
		report.Setup,
		reportdefinition.Setup,
		sdmmaskingpolicydifference.Setup,
		securityassessment.Setup,
		securityassessmentcheck.Setup,
		securityassessmentfinding.Setup,
		securitypolicy.Setup,
		securitypolicyconfig.Setup,
		securitypolicydeployment.Setup,
		securitypolicydeploymentmanagement.Setup,
		securitypolicymanagement.Setup,
		sensitivedatamodel.Setup,
		sensitivedatamodelreferentialrelation.Setup,
		sensitivedatamodelsapplydiscoveryjobresults.Setup,
		sensitivedatamodelssensitivecolumn.Setup,
		sensitivetype.Setup,
		sensitivetypegroup.Setup,
		sensitivetypegroupgroupedsensitivetype.Setup,
		sensitivetypesexport.Setup,
		setsecurityassessmentbaseline.Setup,
		setsecurityassessmentbaselinemanagement.Setup,
		setuserassessmentbaseline.Setup,
		setuserassessmentbaselinemanagement.Setup,
		sqlcollection.Setup,
		sqlfirewallpolicy.Setup,
		sqlfirewallpolicymanagement.Setup,
		targetalertpolicyassociation.Setup,
		targetdatabase.Setup,
		targetdatabasegroup.Setup,
		targetdatabasepeertargetdatabase.Setup,
		unifiedauditpolicy.Setup,
		unifiedauditpolicydefinition.Setup,
		unsetsecurityassessmentbaseline.Setup,
		unsetsecurityassessmentbaselinemanagement.Setup,
		unsetuserassessmentbaseline.Setup,
		unsetuserassessmentbaselinemanagement.Setup,
		userassessment.Setup,
		computetarget.Setup,
		jobdatascience.Setup,
		jobrun.Setup,
		mlapplication.Setup,
		mlapplicationimplementation.Setup,
		mlapplicationinstance.Setup,
		modeldatascience.Setup,
		modelartifactexport.Setup,
		modelartifactimport.Setup,
		modelcustommetadataartifact.Setup,
		modeldefinedmetadataartifact.Setup,
		modeldeployment.Setup,
		modelgroup.Setup,
		modelgroupartifact.Setup,
		modelgroupversionhistory.Setup,
		modelprovenance.Setup,
		modelversionset.Setup,
		notebooksession.Setup,
		pipeline.Setup,
		pipelinerun.Setup,
		privateendpointdatascience.Setup,
		projectdatascience.Setup,
		schedule.Setup,
		vulnerabilityscan.Setup,
		multicloudresourcediscovery.Setup,
		oracledbawsidentityconnector.Setup,
		oracledbawskey.Setup,
		oracledbazureblobcontainer.Setup,
		oracledbazureblobmount.Setup,
		oracledbazureconnector.Setup,
		oracledbazurevault.Setup,
		oracledbazurevaultassociation.Setup,
		oracledbgcpidentityconnector.Setup,
		oracledbgcpkeyring.Setup,
		instanceddfs.Setup,
		delegationcontrol.Setup,
		delegationsubscription.Setup,
		occdemandsignal.Setup,
		occmetricalarm.Setup,
		desktoppool.Setup,
		buildpipeline.Setup,
		buildpipelinestage.Setup,
		buildrun.Setup,
		connectiondevops.Setup,
		deployartifact.Setup,
		deployenvironment.Setup,
		deploymentdevops.Setup,
		deploypipeline.Setup,
		deploystage.Setup,
		projectdevops.Setup,
		projectrepositorysetting.Setup,
		repositorydevops.Setup,
		repositorymirror.Setup,
		repositoryprotectedbranchmanagement.Setup,
		repositoryref.Setup,
		repositorysetting.Setup,
		trigger.Setup,
		stack.Setup,
		automaticdrconfiguration.Setup,
		drplan.Setup,
		drplanexecution.Setup,
		drprotectiongroup.Setup,
		distributedautonomousdatabase.Setup,
		distributeddatabase.Setup,
		distributeddatabaseprivateendpoint.Setup,
		actioncreatezonefromzonefile.Setup,
		record.Setup,
		resolver.Setup,
		resolverendpoint.Setup,
		rrset.Setup,
		steeringpolicy.Setup,
		steeringpolicyattachment.Setup,
		tsigkey.Setup,
		view.Setup,
		zone.Setup,
		zonepromotednsseckeyversion.Setup,
		zonestagednsseckeyversion.Setup,
		dkim.Setup,
		emaildomain.Setup,
		emailippool.Setup,
		emailreturnpath.Setup,
		sender.Setup,
		suppression.Setup,
		rule.Setup,
		export.Setup,
		exportset.Setup,
		filesystem.Setup,
		filesystemquotarule.Setup,
		filesystemsnapshotpolicy.Setup,
		mounttarget.Setup,
		outboundconnector.Setup,
		replication.Setup,
		snapshot.Setup,
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
		fsucollection.Setup,
		fsucycle.Setup,
		fsureadinesscheck.Setup,
		applicationfunctions.Setup,
		function.Setup,
		invokefunction.Setup,
		fusionenvironment.Setup,
		fusionenvironmentadminuser.Setup,
		fusionenvironmentdatamaskingactivity.Setup,
		fusionenvironmentfamily.Setup,
		fusionenvironmentrefreshactivity.Setup,
		fusionenvironmentserviceattachment.Setup,
		gdppipeline.Setup,
		agentagent.Setup,
		agentagentendpoint.Setup,
		agentdataingestionjob.Setup,
		agentdatasource.Setup,
		agentknowledgebase.Setup,
		agentprovisionedcapacity.Setup,
		agenttool.Setup,
		dedicatedaicluster.Setup,
		endpointgenerativeai.Setup,
		generativeaiprivateendpoint.Setup,
		hostedapplication.Setup,
		hostedapplicationiam.Setup,
		hostedapplicationstorage.Setup,
		hosteddeployment.Setup,
		importedmodel.Setup,
		modelgenerativeai.Setup,
		projectgenerativeai.Setup,
		semanticstore.Setup,
		artifactbypath.Setup,
		connectiongoldengate.Setup,
		connectionassignment.Setup,
		databaseregistration.Setup,
		deploymentgoldengate.Setup,
		deploymentbackup.Setup,
		deploymentcertificate.Setup,
		pipelinegoldengate.Setup,
		httpmonitor.Setup,
		httpprobe.Setup,
		pingmonitor.Setup,
		pingprobe.Setup,
		apikey.Setup,
		authenticationpolicy.Setup,
		authtoken.Setup,
		compartment.Setup,
		customersecretkey.Setup,
		dbcredential.Setup,
		domain.Setup,
		domainreplicationtoregion.Setup,
		dynamicgroup.Setup,
		group.Setup,
		identityprovider.Setup,
		idpgroupmapping.Setup,
		importstandardtagsmanagement.Setup,
		networksource.Setup,
		policy.Setup,
		smtpcredential.Setup,
		tag.Setup,
		tagdefault.Setup,
		tagnamespace.Setup,
		uipassword.Setup,
		user.Setup,
		usercapabilitiesmanagement.Setup,
		usergroupmembership.Setup,
		generatescopedaccesstoken.Setup,
		accountrecoverysetting.Setup,
		apikeyidentitydomains.Setup,
		app.Setup,
		approle.Setup,
		approvalworkflow.Setup,
		approvalworkflowassignment.Setup,
		approvalworkflowstep.Setup,
		authenticationfactorsetting.Setup,
		authtokenidentitydomains.Setup,
		cloudgate.Setup,
		cloudgatemapping.Setup,
		cloudgateserver.Setup,
		condition.Setup,
		customersecretkeyidentitydomains.Setup,
		dynamicresourcegroup.Setup,
		grant.Setup,
		groupidentitydomains.Setup,
		identityproofingprovider.Setup,
		identityproofingprovidertemplate.Setup,
		identitypropagationtrust.Setup,
		identityprovideridentitydomains.Setup,
		identitysetting.Setup,
		kmsisetting.Setup,
		mappedattribute.Setup,
		myapikey.Setup,
		myauthtoken.Setup,
		mycustomersecretkey.Setup,
		myoauth2clientcredential.Setup,
		myrequest.Setup,
		mysmtpcredential.Setup,
		mysupportaccount.Setup,
		myuserdbcredential.Setup,
		networkperimeter.Setup,
		notificationsetting.Setup,
		oauth2clientcredential.Setup,
		oauthclientcertificate.Setup,
		oauthpartnercertificate.Setup,
		passwordpolicy.Setup,
		policyidentitydomains.Setup,
		ruleidentitydomains.Setup,
		securityquestion.Setup,
		securityquestionsetting.Setup,
		selfregistrationprofile.Setup,
		setting.Setup,
		smtpcredentialidentitydomains.Setup,
		socialidentityprovider.Setup,
		useridentitydomains.Setup,
		userdbcredential.Setup,
		integrationinstance.Setup,
		oraclemanagedcustomendpoint.Setup,
		privateendpointoutboundconnection.Setup,
		digitaltwinadapter.Setup,
		digitaltwininstance.Setup,
		digitaltwininstanceinvokerawcommand.Setup,
		digitaltwinmodel.Setup,
		digitaltwinrelationship.Setup,
		iotdomain.Setup,
		iotdomainchangedataretentionperiod.Setup,
		iotdomainconfiguredataaccess.Setup,
		iotdomaingroup.Setup,
		iotdomaingroupconfiguredataaccess.Setup,
		fleetjms.Setup,
		fleetadvancedfeatureconfiguration.Setup,
		fleetagentconfiguration.Setup,
		jmsplugin.Setup,
		taskschedule.Setup,
		javadownloadreport.Setup,
		javadownloadtoken.Setup,
		javalicenseacceptancerecord.Setup,
		analyzeapplicationsconfiguration.Setup,
		subscriptionacknowledgmentconfiguration.Setup,
		ekmsprivateendpoint.Setup,
		encrypteddata.Setup,
		generatedkey.Setup,
		key.Setup,
		keyversion.Setup,
		sign.Setup,
		vault.Setup,
		vaultreplication.Setup,
		verify.Setup,
		configurationlicensemanager.Setup,
		licenserecord.Setup,
		productlicense.Setup,
		quota.Setup,
		backend.Setup,
		backendset.Setup,
		certificateloadbalancer.Setup,
		lbhostname.Setup,
		listener.Setup,
		loadbalancer.Setup,
		pathrouteset.Setup,
		routingpolicy.Setup,
		ruleset.Setup,
		sslciphersuite.Setup,
		loganalyticsentity.Setup,
		loganalyticsentityassociationsadd.Setup,
		loganalyticsentityassociationsremove.Setup,
		loganalyticsentitytype.Setup,
		loganalyticsimportcustomcontent.Setup,
		loganalyticsloggroup.Setup,
		loganalyticsobjectcollectionrule.Setup,
		loganalyticspreferencesmanagement.Setup,
		loganalyticsresourcecategoriesmanagement.Setup,
		loganalyticsunprocesseddatabucketmanagement.Setup,
		namespace.Setup,
		namespaceassociation.Setup,
		namespaceingesttimerule.Setup,
		namespaceingesttimerulesmanagement.Setup,
		namespacelookup.Setup,
		namespacelookupsappenddatamanagement.Setup,
		namespacelookupsupdatedatamanagement.Setup,
		namespacescheduledtask.Setup,
		namespacestoragearchivalconfig.Setup,
		namespacestorageenabledisablearchiving.Setup,
		log.Setup,
		loggroup.Setup,
		logsavedsearch.Setup,
		unifiedagentconfiguration.Setup,
		lustrefilesystem.Setup,
		objectstoragelink.Setup,
		kafkacluster.Setup,
		kafkaclusteraddon.Setup,
		kafkaclusterconfig.Setup,
		kafkaclustersuperusersmanagement.Setup,
		managementagent.Setup,
		managementagentdatasource.Setup,
		managementagentinstallkey.Setup,
		namedcredential.Setup,
		managementdashboardsimport.Setup,
		managementsavedsearch.Setup,
		acceptedagreement.Setup,
		listingpackageagreement.Setup,
		marketplaceexternalattestedmetadata.Setup,
		publication.Setup,
		mediaasset.Setup,
		mediaworkflow.Setup,
		mediaworkflowconfiguration.Setup,
		mediaworkflowjob.Setup,
		streamcdnconfig.Setup,
		streamdistributionchannel.Setup,
		streampackagingconfig.Setup,
		customtable.Setup,
		query.Setup,
		schedulemeteringcomputation.Setup,
		usage.Setup,
		usagecarbonemission.Setup,
		usagecarbonemissionsquery.Setup,
		usagestatementemailrecipientsgroup.Setup,
		alarm.Setup,
		alarmsuppression.Setup,
		capturefilter.Setup,
		vtap.Setup,
		bluegreendeployment.Setup,
		mysqlbackup.Setup,
		mysqlchannel.Setup,
		mysqlconfiguration.Setup,
		mysqldbsystem.Setup,
		mysqlheatwavecluster.Setup,
		mysqlreplica.Setup,
		cpe.Setup,
		crossconnect.Setup,
		crossconnectgroup.Setup,
		drg.Setup,
		drgattachment.Setup,
		drgattachmentmanagement.Setup,
		drgattachmentslist.Setup,
		drgroutedistribution.Setup,
		drgroutedistributionstatement.Setup,
		drgroutetable.Setup,
		drgroutetablerouterule.Setup,
		ipsec.Setup,
		ipsecconnectiontunnelmanagement.Setup,
		virtualcircuit.Setup,
		networkfirewall.Setup,
		networkfirewallpolicy.Setup,
		networkfirewallpolicyaddresslist.Setup,
		networkfirewallpolicyapplication.Setup,
		networkfirewallpolicyapplicationgroup.Setup,
		networkfirewallpolicydecryptionprofile.Setup,
		networkfirewallpolicydecryptionrule.Setup,
		networkfirewallpolicymappedsecret.Setup,
		networkfirewallpolicynatrule.Setup,
		networkfirewallpolicysecurityrule.Setup,
		networkfirewallpolicyservice.Setup,
		networkfirewallpolicytunnelinspectionrule.Setup,
		defaultdhcpoptions.Setup,
		defaultdrgroutetable.Setup,
		defaultroutetable.Setup,
		defaultsecuritylist.Setup,
		dhcpoptions.Setup,
		internetgateway.Setup,
		ipv6.Setup,
		localpeeringgateway.Setup,
		natgateway.Setup,
		networksecuritygroup.Setup,
		networksecuritygroupsecurityrule.Setup,
		privateip.Setup,
		publicip.Setup,
		publicippool.Setup,
		publicippoolcapacity.Setup,
		remotepeeringconnection.Setup,
		routetable.Setup,
		routetableattachment.Setup,
		securitylist.Setup,
		servicegateway.Setup,
		subnet.Setup,
		vcn.Setup,
		vlan.Setup,
		vnicattachment.Setup,
		backendnetworkloadbalancer.Setup,
		backendsetnetworkloadbalancer.Setup,
		listenernetworkloadbalancer.Setup,
		networkloadbalancer.Setup,
		networkloadbalancersbackendsetsunified.Setup,
		configurationnosql.Setup,
		index.Setup,
		table.Setup,
		tablereplica.Setup,
		bucket.Setup,
		object.Setup,
		objectlifecyclepolicy.Setup,
		preauthrequest.Setup,
		privateendpointobjectstorage.Setup,
		replicationpolicy.Setup,
		oceinstance.Setup,
		byol.Setup,
		byolallocation.Setup,
		clusterocvp.Setup,
		datastore.Setup,
		datastorecluster.Setup,
		esxihost.Setup,
		managementappliance.Setup,
		sddc.Setup,
		odainstance.Setup,
		odaprivateendpoint.Setup,
		odaprivateendpointattachment.Setup,
		odaprivateendpointscanproxy.Setup,
		notificationtopic.Setup,
		subscription.Setup,
		opainstance.Setup,
		opensearchcluster.Setup,
		opensearchclusterpipeline.Setup,
		operatorcontrol.Setup,
		operatorcontrolassignment.Setup,
		awrhub.Setup,
		awrhubsource.Setup,
		awrhubsourceawrhubsourcesmanagement.Setup,
		chargebackplan.Setup,
		databaseinsight.Setup,
		enterprisemanagerbridge.Setup,
		exadatainsight.Setup,
		hostinsight.Setup,
		newsreport.Setup,
		operationsinsightsprivateendpoint.Setup,
		operationsinsightswarehouse.Setup,
		operationsinsightswarehousedownloadwarehousewallet.Setup,
		operationsinsightswarehouserotatewarehousewallet.Setup,
		operationsinsightswarehouseuser.Setup,
		opsiconfiguration.Setup,
		enrollmentstatus.Setup,
		profile.Setup,
		recommendation.Setup,
		resourceaction.Setup,
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
		profileosmanagementhub.Setup,
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
		addressactionverification.Setup,
		subscriptionospgateway.Setup,
		providerconfig.Setup,
		privateserviceaccess.Setup,
		psqlbackup.Setup,
		psqlconfiguration.Setup,
		psqldbsystem.Setup,
		consumergroup.Setup,
		queue.Setup,
		protecteddatabase.Setup,
		protectionpolicy.Setup,
		recoveryservicesubnet.Setup,
		ocicachebackup.Setup,
		ocicachebackupexporttoobjectstorage.Setup,
		ocicacheconfigset.Setup,
		ocicacheconfigsetlistassociatedocicachecluster.Setup,
		ocicacheuser.Setup,
		ocicacheusergetrediscluster.Setup,
		rediscluster.Setup,
		redisclusterattachocicacheuser.Setup,
		redisclustercreateidentitytoken.Setup,
		redisclusterdetachocicacheuser.Setup,
		redisclustergetocicacheuser.Setup,
		monitoredregion.Setup,
		resourceanalyticsinstance.Setup,
		resourceanalyticsinstanceoacmanagement.Setup,
		tenancyattachment.Setup,
		privateendpointresourcemanager.Setup,
		scheduleresourcescheduler.Setup,
		serviceconnector.Setup,
		securityattribute.Setup,
		securityattributenamespace.Setup,
		subscriptionself.Setup,
		privateapplication.Setup,
		servicecatalog.Setup,
		servicecatalogassociation.Setup,
		baselineablemetric.Setup,
		configstackmonitoring.Setup,
		discoveryjobstackmonitoring.Setup,
		maintenancewindowstackmonitoring.Setup,
		maintenancewindowsretryfailedoperation.Setup,
		maintenancewindowsstop.Setup,
		metricextension.Setup,
		metricextensionmetricextensionongivenresourcesmanagement.Setup,
		monitoredresource.Setup,
		monitoredresourcesassociatemonitoredresource.Setup,
		monitoredresourceslistmember.Setup,
		monitoredresourcessearch.Setup,
		monitoredresourcessearchassociation.Setup,
		monitoredresourcetask.Setup,
		monitoredresourcetype.Setup,
		monitoringtemplate.Setup,
		monitoringtemplatealarmcondition.Setup,
		monitoringtemplateongivenresourcesmanagement.Setup,
		processset.Setup,
		connectharness.Setup,
		stream.Setup,
		streampool.Setup,
		subscriptionmapping.Setup,
		subscriptionredeemableuser.Setup,
		secret.Setup,
		vbsinstance.Setup,
		vbinstance.Setup,
		pathanalysi.Setup,
		containerscanrecipe.Setup,
		containerscantarget.Setup,
		hostscanrecipe.Setup,
		hostscantarget.Setup,
		webappacceleration.Setup,
		webappaccelerationpolicy.Setup,
		addresslist.Setup,
		certificatewaas.Setup,
		customprotectionrule.Setup,
		httpredirect.Setup,
		protectionrule.Setup,
		purgecache.Setup,
		waaspolicy.Setup,
		networkaddresslist.Setup,
		webappfirewall.Setup,
		webappfirewallpolicy.Setup,
		configurationzpr.Setup,
		zprpolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_monolith creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_monolith(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		knowledgebase.SetupGated,
		remediationrecipe.SetupGated,
		remediationrun.SetupGated,
		vulnerabilityaudit.SetupGated,
		aidataplatform.SetupGated,
		model.SetupGated,
		processorjob.SetupGated,
		project.SetupGated,
		endpoint.SetupGated,
		job.SetupGated,
		modelailanguage.SetupGated,
		projectailanguage.SetupGated,
		modelaivision.SetupGated,
		projectaivision.SetupGated,
		streamgroup.SetupGated,
		streamjob.SetupGated,
		streamsource.SetupGated,
		visionprivateendpoint.SetupGated,
		analyticsinstance.SetupGated,
		analyticsinstanceprivateaccesschannel.SetupGated,
		analyticsinstanceresourcegroup.SetupGated,
		analyticsinstancevanityurl.SetupGated,
		announcementsubscription.SetupGated,
		announcementsubscriptionsactionschangecompartment.SetupGated,
		announcementsubscriptionsfiltergroup.SetupGated,
		privilegedapicontrol.SetupGated,
		privilegedapirequest.SetupGated,
		api.SetupGated,
		certificate.SetupGated,
		deployment.SetupGated,
		gateway.SetupGated,
		subscriber.SetupGated,
		usageplan.SetupGated,
		apiplatforminstance.SetupGated,
		apmdomain.SetupGated,
		config.SetupGated,
		datafile.SetupGated,
		dedicatedvantagepoint.SetupGated,
		monitor.SetupGated,
		onpremisevantagepoint.SetupGated,
		onpremisevantagepointworker.SetupGated,
		script.SetupGated,
		scheduledquery.SetupGated,
		controlmonitorpluginmanagement.SetupGated,
		containerconfiguration.SetupGated,
		containerimagesignature.SetupGated,
		containerrepository.SetupGated,
		genericartifact.SetupGated,
		repository.SetupGated,
		configuration.SetupGated,
		autoscalingconfiguration.SetupGated,
		bastion.SetupGated,
		session.SetupGated,
		batchcontext.SetupGated,
		batchjobpool.SetupGated,
		batchtaskenvironment.SetupGated,
		batchtaskprofile.SetupGated,
		autoscalingconfigurationbds.SetupGated,
		bdscapacityreport.SetupGated,
		bdscapacityreservation.SetupGated,
		bdsclusteradminpasswordresetaction.SetupGated,
		bdsinstance.SetupGated,
		bdsinstanceapikey.SetupGated,
		bdsinstancebdscapacityreservationconfiguration.SetupGated,
		bdsinstancebdscertificateconfiguration.SetupGated,
		bdsinstanceexecutebootstrapscriptaction.SetupGated,
		bdsinstanceidentityconfiguration.SetupGated,
		bdsinstancemetastoreconfig.SetupGated,
		bdsinstancenodebackup.SetupGated,
		bdsinstancenodebackupconfiguration.SetupGated,
		bdsinstancenodereplaceconfiguration.SetupGated,
		bdsinstanceoperationcertificatemanagementsmanagement.SetupGated,
		bdsinstanceospatchaction.SetupGated,
		bdsinstancepatchaction.SetupGated,
		bdsinstancereplacenodeaction.SetupGated,
		bdsinstanceresourceprincipalconfiguration.SetupGated,
		bdsinstancesoftwareupdateaction.SetupGated,
		blockchainplatform.SetupGated,
		osn.SetupGated,
		peer.SetupGated,
		bootvolume.SetupGated,
		bootvolumebackup.SetupGated,
		volume.SetupGated,
		volumeattachment.SetupGated,
		volumebackup.SetupGated,
		volumebackuppolicy.SetupGated,
		volumebackuppolicyassignment.SetupGated,
		volumegroup.SetupGated,
		volumegroupbackup.SetupGated,
		alertrule.SetupGated,
		budget.SetupGated,
		costalertsubscription.SetupGated,
		costanomalyevent.SetupGated,
		costanomalymonitor.SetupGated,
		costanomalymonitorcostanomalymonitorenabletogglesmanagement.SetupGated,
		internaloccmdemandsignal.SetupGated,
		internaloccmdemandsignaldelivery.SetupGated,
		occavailabilitycatalog.SetupGated,
		occcapacityrequest.SetupGated,
		occcustomergroup.SetupGated,
		occcustomergroupocccustomer.SetupGated,
		occmdemandsignal.SetupGated,
		occmdemandsignalitem.SetupGated,
		cabundle.SetupGated,
		certificatecertificatesmanagement.SetupGated,
		certificateauthority.SetupGated,
		agent.SetupGated,
		agentdependency.SetupGated,
		agentplugin.SetupGated,
		asset.SetupGated,
		assetsource.SetupGated,
		discoveryschedule.SetupGated,
		environment.SetupGated,
		inventory.SetupGated,
		adhocquery.SetupGated,
		cloudguardconfiguration.SetupGated,
		datamaskrule.SetupGated,
		datasource.SetupGated,
		detectorrecipe.SetupGated,
		managedlist.SetupGated,
		responderrecipe.SetupGated,
		savedquery.SetupGated,
		securityrecipe.SetupGated,
		securityzone.SetupGated,
		target.SetupGated,
		wlpagent.SetupGated,
		migration.SetupGated,
		migrationasset.SetupGated,
		migrationplan.SetupGated,
		replicationschedule.SetupGated,
		targetasset.SetupGated,
		clusterplacementgroup.SetupGated,
		appcataloglistingresourceversionagreement.SetupGated,
		appcatalogsubscription.SetupGated,
		clusternetwork.SetupGated,
		computecapacityreport.SetupGated,
		computecapacityreservation.SetupGated,
		computecapacitytopology.SetupGated,
		computecluster.SetupGated,
		computegpumemorycluster.SetupGated,
		computegpumemoryfabric.SetupGated,
		computehost.SetupGated,
		computehostgroup.SetupGated,
		computeimagecapabilityschema.SetupGated,
		consolehistory.SetupGated,
		dedicatedvmhost.SetupGated,
		image.SetupGated,
		instance.SetupGated,
		instanceconfiguration.SetupGated,
		instanceconsoleconnection.SetupGated,
		instancemaintenanceevent.SetupGated,
		instancepool.SetupGated,
		instancepoolinstance.SetupGated,
		shapemanagement.SetupGated,
		cccinfrastructure.SetupGated,
		cccupgradeschedule.SetupGated,
		addon.SetupGated,
		cluster.SetupGated,
		clustercompletecredentialrotationmanagement.SetupGated,
		clusterpublicapiendpointdecommissionmanager.SetupGated,
		clusterstartcredentialrotationmanagement.SetupGated,
		clusterworkloadmapping.SetupGated,
		nodepool.SetupGated,
		virtualnodepool.SetupGated,
		containerinstance.SetupGated,
		byoasn.SetupGated,
		listingresourceversionagreement.SetupGated,
		virtualnetwork.SetupGated,
		costalertsubscriptioncostad.SetupGated,
		costanomalyeventcostad.SetupGated,
		costanomalymonitorcostad.SetupGated,
		costanomalymonitorcostanomalymonitorenabletogglesmanagementcostad.SetupGated,
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
		migrationdatabase.SetupGated,
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
		catalog.SetupGated,
		catalogprivateendpoint.SetupGated,
		connection.SetupGated,
		dataasset.SetupGated,
		metastore.SetupGated,
		infrastructure.SetupGated,
		vmclusternetworkdatacc.SetupGated,
		vminstance.SetupGated,
		application.SetupGated,
		invokerun.SetupGated,
		pool.SetupGated,
		privateendpoint.SetupGated,
		runstatement.SetupGated,
		sqlendpoint.SetupGated,
		workspace.SetupGated,
		workspaceapplication.SetupGated,
		workspaceapplicationpatch.SetupGated,
		workspaceapplicationschedule.SetupGated,
		workspaceapplicationtaskschedule.SetupGated,
		workspaceexportrequest.SetupGated,
		workspacefolder.SetupGated,
		workspaceimportrequest.SetupGated,
		workspaceproject.SetupGated,
		workspacetask.SetupGated,
		dataset.SetupGated,
		addsdmcolumns.SetupGated,
		alert.SetupGated,
		alertpolicy.SetupGated,
		alertpolicyrule.SetupGated,
		attributeset.SetupGated,
		auditarchiveretrieval.SetupGated,
		auditpolicy.SetupGated,
		auditpolicymanagement.SetupGated,
		auditprofile.SetupGated,
		auditprofilemanagement.SetupGated,
		audittrail.SetupGated,
		audittrailmanagement.SetupGated,
		calculateauditvolumeavailable.SetupGated,
		calculateauditvolumecollected.SetupGated,
		comparesecurityassessment.SetupGated,
		compareuserassessment.SetupGated,
		databasesecurityconfig.SetupGated,
		databasesecurityconfigmanagement.SetupGated,
		datasafeconfiguration.SetupGated,
		datasafeprivateendpoint.SetupGated,
		discoveryjob.SetupGated,
		discoveryjobsresult.SetupGated,
		generateonpremconnectorconfiguration.SetupGated,
		librarymaskingformat.SetupGated,
		maskdata.SetupGated,
		maskingpoliciesapplydifferencetomaskingcolumns.SetupGated,
		maskingpoliciesmaskingcolumn.SetupGated,
		maskingpolicy.SetupGated,
		maskingpolicyhealthreportmanagement.SetupGated,
		maskingreportmanagement.SetupGated,
		onpremconnector.SetupGated,
		report.SetupGated,
		reportdefinition.SetupGated,
		sdmmaskingpolicydifference.SetupGated,
		securityassessment.SetupGated,
		securityassessmentcheck.SetupGated,
		securityassessmentfinding.SetupGated,
		securitypolicy.SetupGated,
		securitypolicyconfig.SetupGated,
		securitypolicydeployment.SetupGated,
		securitypolicydeploymentmanagement.SetupGated,
		securitypolicymanagement.SetupGated,
		sensitivedatamodel.SetupGated,
		sensitivedatamodelreferentialrelation.SetupGated,
		sensitivedatamodelsapplydiscoveryjobresults.SetupGated,
		sensitivedatamodelssensitivecolumn.SetupGated,
		sensitivetype.SetupGated,
		sensitivetypegroup.SetupGated,
		sensitivetypegroupgroupedsensitivetype.SetupGated,
		sensitivetypesexport.SetupGated,
		setsecurityassessmentbaseline.SetupGated,
		setsecurityassessmentbaselinemanagement.SetupGated,
		setuserassessmentbaseline.SetupGated,
		setuserassessmentbaselinemanagement.SetupGated,
		sqlcollection.SetupGated,
		sqlfirewallpolicy.SetupGated,
		sqlfirewallpolicymanagement.SetupGated,
		targetalertpolicyassociation.SetupGated,
		targetdatabase.SetupGated,
		targetdatabasegroup.SetupGated,
		targetdatabasepeertargetdatabase.SetupGated,
		unifiedauditpolicy.SetupGated,
		unifiedauditpolicydefinition.SetupGated,
		unsetsecurityassessmentbaseline.SetupGated,
		unsetsecurityassessmentbaselinemanagement.SetupGated,
		unsetuserassessmentbaseline.SetupGated,
		unsetuserassessmentbaselinemanagement.SetupGated,
		userassessment.SetupGated,
		computetarget.SetupGated,
		jobdatascience.SetupGated,
		jobrun.SetupGated,
		mlapplication.SetupGated,
		mlapplicationimplementation.SetupGated,
		mlapplicationinstance.SetupGated,
		modeldatascience.SetupGated,
		modelartifactexport.SetupGated,
		modelartifactimport.SetupGated,
		modelcustommetadataartifact.SetupGated,
		modeldefinedmetadataartifact.SetupGated,
		modeldeployment.SetupGated,
		modelgroup.SetupGated,
		modelgroupartifact.SetupGated,
		modelgroupversionhistory.SetupGated,
		modelprovenance.SetupGated,
		modelversionset.SetupGated,
		notebooksession.SetupGated,
		pipeline.SetupGated,
		pipelinerun.SetupGated,
		privateendpointdatascience.SetupGated,
		projectdatascience.SetupGated,
		schedule.SetupGated,
		vulnerabilityscan.SetupGated,
		multicloudresourcediscovery.SetupGated,
		oracledbawsidentityconnector.SetupGated,
		oracledbawskey.SetupGated,
		oracledbazureblobcontainer.SetupGated,
		oracledbazureblobmount.SetupGated,
		oracledbazureconnector.SetupGated,
		oracledbazurevault.SetupGated,
		oracledbazurevaultassociation.SetupGated,
		oracledbgcpidentityconnector.SetupGated,
		oracledbgcpkeyring.SetupGated,
		instanceddfs.SetupGated,
		delegationcontrol.SetupGated,
		delegationsubscription.SetupGated,
		occdemandsignal.SetupGated,
		occmetricalarm.SetupGated,
		desktoppool.SetupGated,
		buildpipeline.SetupGated,
		buildpipelinestage.SetupGated,
		buildrun.SetupGated,
		connectiondevops.SetupGated,
		deployartifact.SetupGated,
		deployenvironment.SetupGated,
		deploymentdevops.SetupGated,
		deploypipeline.SetupGated,
		deploystage.SetupGated,
		projectdevops.SetupGated,
		projectrepositorysetting.SetupGated,
		repositorydevops.SetupGated,
		repositorymirror.SetupGated,
		repositoryprotectedbranchmanagement.SetupGated,
		repositoryref.SetupGated,
		repositorysetting.SetupGated,
		trigger.SetupGated,
		stack.SetupGated,
		automaticdrconfiguration.SetupGated,
		drplan.SetupGated,
		drplanexecution.SetupGated,
		drprotectiongroup.SetupGated,
		distributedautonomousdatabase.SetupGated,
		distributeddatabase.SetupGated,
		distributeddatabaseprivateendpoint.SetupGated,
		actioncreatezonefromzonefile.SetupGated,
		record.SetupGated,
		resolver.SetupGated,
		resolverendpoint.SetupGated,
		rrset.SetupGated,
		steeringpolicy.SetupGated,
		steeringpolicyattachment.SetupGated,
		tsigkey.SetupGated,
		view.SetupGated,
		zone.SetupGated,
		zonepromotednsseckeyversion.SetupGated,
		zonestagednsseckeyversion.SetupGated,
		dkim.SetupGated,
		emaildomain.SetupGated,
		emailippool.SetupGated,
		emailreturnpath.SetupGated,
		sender.SetupGated,
		suppression.SetupGated,
		rule.SetupGated,
		export.SetupGated,
		exportset.SetupGated,
		filesystem.SetupGated,
		filesystemquotarule.SetupGated,
		filesystemsnapshotpolicy.SetupGated,
		mounttarget.SetupGated,
		outboundconnector.SetupGated,
		replication.SetupGated,
		snapshot.SetupGated,
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
		fsucollection.SetupGated,
		fsucycle.SetupGated,
		fsureadinesscheck.SetupGated,
		applicationfunctions.SetupGated,
		function.SetupGated,
		invokefunction.SetupGated,
		fusionenvironment.SetupGated,
		fusionenvironmentadminuser.SetupGated,
		fusionenvironmentdatamaskingactivity.SetupGated,
		fusionenvironmentfamily.SetupGated,
		fusionenvironmentrefreshactivity.SetupGated,
		fusionenvironmentserviceattachment.SetupGated,
		gdppipeline.SetupGated,
		agentagent.SetupGated,
		agentagentendpoint.SetupGated,
		agentdataingestionjob.SetupGated,
		agentdatasource.SetupGated,
		agentknowledgebase.SetupGated,
		agentprovisionedcapacity.SetupGated,
		agenttool.SetupGated,
		dedicatedaicluster.SetupGated,
		endpointgenerativeai.SetupGated,
		generativeaiprivateendpoint.SetupGated,
		hostedapplication.SetupGated,
		hostedapplicationiam.SetupGated,
		hostedapplicationstorage.SetupGated,
		hosteddeployment.SetupGated,
		importedmodel.SetupGated,
		modelgenerativeai.SetupGated,
		projectgenerativeai.SetupGated,
		semanticstore.SetupGated,
		artifactbypath.SetupGated,
		connectiongoldengate.SetupGated,
		connectionassignment.SetupGated,
		databaseregistration.SetupGated,
		deploymentgoldengate.SetupGated,
		deploymentbackup.SetupGated,
		deploymentcertificate.SetupGated,
		pipelinegoldengate.SetupGated,
		httpmonitor.SetupGated,
		httpprobe.SetupGated,
		pingmonitor.SetupGated,
		pingprobe.SetupGated,
		apikey.SetupGated,
		authenticationpolicy.SetupGated,
		authtoken.SetupGated,
		compartment.SetupGated,
		customersecretkey.SetupGated,
		dbcredential.SetupGated,
		domain.SetupGated,
		domainreplicationtoregion.SetupGated,
		dynamicgroup.SetupGated,
		group.SetupGated,
		identityprovider.SetupGated,
		idpgroupmapping.SetupGated,
		importstandardtagsmanagement.SetupGated,
		networksource.SetupGated,
		policy.SetupGated,
		smtpcredential.SetupGated,
		tag.SetupGated,
		tagdefault.SetupGated,
		tagnamespace.SetupGated,
		uipassword.SetupGated,
		user.SetupGated,
		usercapabilitiesmanagement.SetupGated,
		usergroupmembership.SetupGated,
		generatescopedaccesstoken.SetupGated,
		accountrecoverysetting.SetupGated,
		apikeyidentitydomains.SetupGated,
		app.SetupGated,
		approle.SetupGated,
		approvalworkflow.SetupGated,
		approvalworkflowassignment.SetupGated,
		approvalworkflowstep.SetupGated,
		authenticationfactorsetting.SetupGated,
		authtokenidentitydomains.SetupGated,
		cloudgate.SetupGated,
		cloudgatemapping.SetupGated,
		cloudgateserver.SetupGated,
		condition.SetupGated,
		customersecretkeyidentitydomains.SetupGated,
		dynamicresourcegroup.SetupGated,
		grant.SetupGated,
		groupidentitydomains.SetupGated,
		identityproofingprovider.SetupGated,
		identityproofingprovidertemplate.SetupGated,
		identitypropagationtrust.SetupGated,
		identityprovideridentitydomains.SetupGated,
		identitysetting.SetupGated,
		kmsisetting.SetupGated,
		mappedattribute.SetupGated,
		myapikey.SetupGated,
		myauthtoken.SetupGated,
		mycustomersecretkey.SetupGated,
		myoauth2clientcredential.SetupGated,
		myrequest.SetupGated,
		mysmtpcredential.SetupGated,
		mysupportaccount.SetupGated,
		myuserdbcredential.SetupGated,
		networkperimeter.SetupGated,
		notificationsetting.SetupGated,
		oauth2clientcredential.SetupGated,
		oauthclientcertificate.SetupGated,
		oauthpartnercertificate.SetupGated,
		passwordpolicy.SetupGated,
		policyidentitydomains.SetupGated,
		ruleidentitydomains.SetupGated,
		securityquestion.SetupGated,
		securityquestionsetting.SetupGated,
		selfregistrationprofile.SetupGated,
		setting.SetupGated,
		smtpcredentialidentitydomains.SetupGated,
		socialidentityprovider.SetupGated,
		useridentitydomains.SetupGated,
		userdbcredential.SetupGated,
		integrationinstance.SetupGated,
		oraclemanagedcustomendpoint.SetupGated,
		privateendpointoutboundconnection.SetupGated,
		digitaltwinadapter.SetupGated,
		digitaltwininstance.SetupGated,
		digitaltwininstanceinvokerawcommand.SetupGated,
		digitaltwinmodel.SetupGated,
		digitaltwinrelationship.SetupGated,
		iotdomain.SetupGated,
		iotdomainchangedataretentionperiod.SetupGated,
		iotdomainconfiguredataaccess.SetupGated,
		iotdomaingroup.SetupGated,
		iotdomaingroupconfiguredataaccess.SetupGated,
		fleetjms.SetupGated,
		fleetadvancedfeatureconfiguration.SetupGated,
		fleetagentconfiguration.SetupGated,
		jmsplugin.SetupGated,
		taskschedule.SetupGated,
		javadownloadreport.SetupGated,
		javadownloadtoken.SetupGated,
		javalicenseacceptancerecord.SetupGated,
		analyzeapplicationsconfiguration.SetupGated,
		subscriptionacknowledgmentconfiguration.SetupGated,
		ekmsprivateendpoint.SetupGated,
		encrypteddata.SetupGated,
		generatedkey.SetupGated,
		key.SetupGated,
		keyversion.SetupGated,
		sign.SetupGated,
		vault.SetupGated,
		vaultreplication.SetupGated,
		verify.SetupGated,
		configurationlicensemanager.SetupGated,
		licenserecord.SetupGated,
		productlicense.SetupGated,
		quota.SetupGated,
		backend.SetupGated,
		backendset.SetupGated,
		certificateloadbalancer.SetupGated,
		lbhostname.SetupGated,
		listener.SetupGated,
		loadbalancer.SetupGated,
		pathrouteset.SetupGated,
		routingpolicy.SetupGated,
		ruleset.SetupGated,
		sslciphersuite.SetupGated,
		loganalyticsentity.SetupGated,
		loganalyticsentityassociationsadd.SetupGated,
		loganalyticsentityassociationsremove.SetupGated,
		loganalyticsentitytype.SetupGated,
		loganalyticsimportcustomcontent.SetupGated,
		loganalyticsloggroup.SetupGated,
		loganalyticsobjectcollectionrule.SetupGated,
		loganalyticspreferencesmanagement.SetupGated,
		loganalyticsresourcecategoriesmanagement.SetupGated,
		loganalyticsunprocesseddatabucketmanagement.SetupGated,
		namespace.SetupGated,
		namespaceassociation.SetupGated,
		namespaceingesttimerule.SetupGated,
		namespaceingesttimerulesmanagement.SetupGated,
		namespacelookup.SetupGated,
		namespacelookupsappenddatamanagement.SetupGated,
		namespacelookupsupdatedatamanagement.SetupGated,
		namespacescheduledtask.SetupGated,
		namespacestoragearchivalconfig.SetupGated,
		namespacestorageenabledisablearchiving.SetupGated,
		log.SetupGated,
		loggroup.SetupGated,
		logsavedsearch.SetupGated,
		unifiedagentconfiguration.SetupGated,
		lustrefilesystem.SetupGated,
		objectstoragelink.SetupGated,
		kafkacluster.SetupGated,
		kafkaclusteraddon.SetupGated,
		kafkaclusterconfig.SetupGated,
		kafkaclustersuperusersmanagement.SetupGated,
		managementagent.SetupGated,
		managementagentdatasource.SetupGated,
		managementagentinstallkey.SetupGated,
		namedcredential.SetupGated,
		managementdashboardsimport.SetupGated,
		managementsavedsearch.SetupGated,
		acceptedagreement.SetupGated,
		listingpackageagreement.SetupGated,
		marketplaceexternalattestedmetadata.SetupGated,
		publication.SetupGated,
		mediaasset.SetupGated,
		mediaworkflow.SetupGated,
		mediaworkflowconfiguration.SetupGated,
		mediaworkflowjob.SetupGated,
		streamcdnconfig.SetupGated,
		streamdistributionchannel.SetupGated,
		streampackagingconfig.SetupGated,
		customtable.SetupGated,
		query.SetupGated,
		schedulemeteringcomputation.SetupGated,
		usage.SetupGated,
		usagecarbonemission.SetupGated,
		usagecarbonemissionsquery.SetupGated,
		usagestatementemailrecipientsgroup.SetupGated,
		alarm.SetupGated,
		alarmsuppression.SetupGated,
		capturefilter.SetupGated,
		vtap.SetupGated,
		bluegreendeployment.SetupGated,
		mysqlbackup.SetupGated,
		mysqlchannel.SetupGated,
		mysqlconfiguration.SetupGated,
		mysqldbsystem.SetupGated,
		mysqlheatwavecluster.SetupGated,
		mysqlreplica.SetupGated,
		cpe.SetupGated,
		crossconnect.SetupGated,
		crossconnectgroup.SetupGated,
		drg.SetupGated,
		drgattachment.SetupGated,
		drgattachmentmanagement.SetupGated,
		drgattachmentslist.SetupGated,
		drgroutedistribution.SetupGated,
		drgroutedistributionstatement.SetupGated,
		drgroutetable.SetupGated,
		drgroutetablerouterule.SetupGated,
		ipsec.SetupGated,
		ipsecconnectiontunnelmanagement.SetupGated,
		virtualcircuit.SetupGated,
		networkfirewall.SetupGated,
		networkfirewallpolicy.SetupGated,
		networkfirewallpolicyaddresslist.SetupGated,
		networkfirewallpolicyapplication.SetupGated,
		networkfirewallpolicyapplicationgroup.SetupGated,
		networkfirewallpolicydecryptionprofile.SetupGated,
		networkfirewallpolicydecryptionrule.SetupGated,
		networkfirewallpolicymappedsecret.SetupGated,
		networkfirewallpolicynatrule.SetupGated,
		networkfirewallpolicysecurityrule.SetupGated,
		networkfirewallpolicyservice.SetupGated,
		networkfirewallpolicytunnelinspectionrule.SetupGated,
		defaultdhcpoptions.SetupGated,
		defaultdrgroutetable.SetupGated,
		defaultroutetable.SetupGated,
		defaultsecuritylist.SetupGated,
		dhcpoptions.SetupGated,
		internetgateway.SetupGated,
		ipv6.SetupGated,
		localpeeringgateway.SetupGated,
		natgateway.SetupGated,
		networksecuritygroup.SetupGated,
		networksecuritygroupsecurityrule.SetupGated,
		privateip.SetupGated,
		publicip.SetupGated,
		publicippool.SetupGated,
		publicippoolcapacity.SetupGated,
		remotepeeringconnection.SetupGated,
		routetable.SetupGated,
		routetableattachment.SetupGated,
		securitylist.SetupGated,
		servicegateway.SetupGated,
		subnet.SetupGated,
		vcn.SetupGated,
		vlan.SetupGated,
		vnicattachment.SetupGated,
		backendnetworkloadbalancer.SetupGated,
		backendsetnetworkloadbalancer.SetupGated,
		listenernetworkloadbalancer.SetupGated,
		networkloadbalancer.SetupGated,
		networkloadbalancersbackendsetsunified.SetupGated,
		configurationnosql.SetupGated,
		index.SetupGated,
		table.SetupGated,
		tablereplica.SetupGated,
		bucket.SetupGated,
		object.SetupGated,
		objectlifecyclepolicy.SetupGated,
		preauthrequest.SetupGated,
		privateendpointobjectstorage.SetupGated,
		replicationpolicy.SetupGated,
		oceinstance.SetupGated,
		byol.SetupGated,
		byolallocation.SetupGated,
		clusterocvp.SetupGated,
		datastore.SetupGated,
		datastorecluster.SetupGated,
		esxihost.SetupGated,
		managementappliance.SetupGated,
		sddc.SetupGated,
		odainstance.SetupGated,
		odaprivateendpoint.SetupGated,
		odaprivateendpointattachment.SetupGated,
		odaprivateendpointscanproxy.SetupGated,
		notificationtopic.SetupGated,
		subscription.SetupGated,
		opainstance.SetupGated,
		opensearchcluster.SetupGated,
		opensearchclusterpipeline.SetupGated,
		operatorcontrol.SetupGated,
		operatorcontrolassignment.SetupGated,
		awrhub.SetupGated,
		awrhubsource.SetupGated,
		awrhubsourceawrhubsourcesmanagement.SetupGated,
		chargebackplan.SetupGated,
		databaseinsight.SetupGated,
		enterprisemanagerbridge.SetupGated,
		exadatainsight.SetupGated,
		hostinsight.SetupGated,
		newsreport.SetupGated,
		operationsinsightsprivateendpoint.SetupGated,
		operationsinsightswarehouse.SetupGated,
		operationsinsightswarehousedownloadwarehousewallet.SetupGated,
		operationsinsightswarehouserotatewarehousewallet.SetupGated,
		operationsinsightswarehouseuser.SetupGated,
		opsiconfiguration.SetupGated,
		enrollmentstatus.SetupGated,
		profile.SetupGated,
		recommendation.SetupGated,
		resourceaction.SetupGated,
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
		profileosmanagementhub.SetupGated,
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
		addressactionverification.SetupGated,
		subscriptionospgateway.SetupGated,
		providerconfig.SetupGated,
		privateserviceaccess.SetupGated,
		psqlbackup.SetupGated,
		psqlconfiguration.SetupGated,
		psqldbsystem.SetupGated,
		consumergroup.SetupGated,
		queue.SetupGated,
		protecteddatabase.SetupGated,
		protectionpolicy.SetupGated,
		recoveryservicesubnet.SetupGated,
		ocicachebackup.SetupGated,
		ocicachebackupexporttoobjectstorage.SetupGated,
		ocicacheconfigset.SetupGated,
		ocicacheconfigsetlistassociatedocicachecluster.SetupGated,
		ocicacheuser.SetupGated,
		ocicacheusergetrediscluster.SetupGated,
		rediscluster.SetupGated,
		redisclusterattachocicacheuser.SetupGated,
		redisclustercreateidentitytoken.SetupGated,
		redisclusterdetachocicacheuser.SetupGated,
		redisclustergetocicacheuser.SetupGated,
		monitoredregion.SetupGated,
		resourceanalyticsinstance.SetupGated,
		resourceanalyticsinstanceoacmanagement.SetupGated,
		tenancyattachment.SetupGated,
		privateendpointresourcemanager.SetupGated,
		scheduleresourcescheduler.SetupGated,
		serviceconnector.SetupGated,
		securityattribute.SetupGated,
		securityattributenamespace.SetupGated,
		subscriptionself.SetupGated,
		privateapplication.SetupGated,
		servicecatalog.SetupGated,
		servicecatalogassociation.SetupGated,
		baselineablemetric.SetupGated,
		configstackmonitoring.SetupGated,
		discoveryjobstackmonitoring.SetupGated,
		maintenancewindowstackmonitoring.SetupGated,
		maintenancewindowsretryfailedoperation.SetupGated,
		maintenancewindowsstop.SetupGated,
		metricextension.SetupGated,
		metricextensionmetricextensionongivenresourcesmanagement.SetupGated,
		monitoredresource.SetupGated,
		monitoredresourcesassociatemonitoredresource.SetupGated,
		monitoredresourceslistmember.SetupGated,
		monitoredresourcessearch.SetupGated,
		monitoredresourcessearchassociation.SetupGated,
		monitoredresourcetask.SetupGated,
		monitoredresourcetype.SetupGated,
		monitoringtemplate.SetupGated,
		monitoringtemplatealarmcondition.SetupGated,
		monitoringtemplateongivenresourcesmanagement.SetupGated,
		processset.SetupGated,
		connectharness.SetupGated,
		stream.SetupGated,
		streampool.SetupGated,
		subscriptionmapping.SetupGated,
		subscriptionredeemableuser.SetupGated,
		secret.SetupGated,
		vbsinstance.SetupGated,
		vbinstance.SetupGated,
		pathanalysi.SetupGated,
		containerscanrecipe.SetupGated,
		containerscantarget.SetupGated,
		hostscanrecipe.SetupGated,
		hostscantarget.SetupGated,
		webappacceleration.SetupGated,
		webappaccelerationpolicy.SetupGated,
		addresslist.SetupGated,
		certificatewaas.SetupGated,
		customprotectionrule.SetupGated,
		httpredirect.SetupGated,
		protectionrule.SetupGated,
		purgecache.SetupGated,
		waaspolicy.SetupGated,
		networkaddresslist.SetupGated,
		webappfirewall.SetupGated,
		webappfirewallpolicy.SetupGated,
		configurationzpr.SetupGated,
		zprpolicy.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

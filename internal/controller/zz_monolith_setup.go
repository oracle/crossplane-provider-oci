/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	knowledgebase "github.com/oracle/provider-oci/internal/controller/adm/knowledgebase"
	remediationrecipe "github.com/oracle/provider-oci/internal/controller/adm/remediationrecipe"
	remediationrun "github.com/oracle/provider-oci/internal/controller/adm/remediationrun"
	vulnerabilityaudit "github.com/oracle/provider-oci/internal/controller/adm/vulnerabilityaudit"
	aidataplatform "github.com/oracle/provider-oci/internal/controller/aidataplatform/aidataplatform"
	model "github.com/oracle/provider-oci/internal/controller/aidocument/model"
	processorjob "github.com/oracle/provider-oci/internal/controller/aidocument/processorjob"
	project "github.com/oracle/provider-oci/internal/controller/aidocument/project"
	endpoint "github.com/oracle/provider-oci/internal/controller/ailanguage/endpoint"
	job "github.com/oracle/provider-oci/internal/controller/ailanguage/job"
	modelailanguage "github.com/oracle/provider-oci/internal/controller/ailanguage/model"
	projectailanguage "github.com/oracle/provider-oci/internal/controller/ailanguage/project"
	modelaivision "github.com/oracle/provider-oci/internal/controller/aivision/model"
	projectaivision "github.com/oracle/provider-oci/internal/controller/aivision/project"
	streamgroup "github.com/oracle/provider-oci/internal/controller/aivision/streamgroup"
	streamjob "github.com/oracle/provider-oci/internal/controller/aivision/streamjob"
	streamsource "github.com/oracle/provider-oci/internal/controller/aivision/streamsource"
	visionprivateendpoint "github.com/oracle/provider-oci/internal/controller/aivision/visionprivateendpoint"
	analyticsinstance "github.com/oracle/provider-oci/internal/controller/analytics/analyticsinstance"
	analyticsinstanceprivateaccesschannel "github.com/oracle/provider-oci/internal/controller/analytics/analyticsinstanceprivateaccesschannel"
	analyticsinstancevanityurl "github.com/oracle/provider-oci/internal/controller/analytics/analyticsinstancevanityurl"
	announcementsubscription "github.com/oracle/provider-oci/internal/controller/announcementsservice/announcementsubscription"
	announcementsubscriptionsactionschangecompartment "github.com/oracle/provider-oci/internal/controller/announcementsservice/announcementsubscriptionsactionschangecompartment"
	announcementsubscriptionsfiltergroup "github.com/oracle/provider-oci/internal/controller/announcementsservice/announcementsubscriptionsfiltergroup"
	privilegedapicontrol "github.com/oracle/provider-oci/internal/controller/apiaccesscontrol/privilegedapicontrol"
	privilegedapirequest "github.com/oracle/provider-oci/internal/controller/apiaccesscontrol/privilegedapirequest"
	api "github.com/oracle/provider-oci/internal/controller/apigateway/api"
	certificate "github.com/oracle/provider-oci/internal/controller/apigateway/certificate"
	deployment "github.com/oracle/provider-oci/internal/controller/apigateway/deployment"
	gateway "github.com/oracle/provider-oci/internal/controller/apigateway/gateway"
	subscriber "github.com/oracle/provider-oci/internal/controller/apigateway/subscriber"
	usageplan "github.com/oracle/provider-oci/internal/controller/apigateway/usageplan"
	apiplatforminstance "github.com/oracle/provider-oci/internal/controller/apiplatform/apiplatforminstance"
	apmdomain "github.com/oracle/provider-oci/internal/controller/apm/apmdomain"
	config "github.com/oracle/provider-oci/internal/controller/apmconfig/config"
	dedicatedvantagepoint "github.com/oracle/provider-oci/internal/controller/apmsynthetics/dedicatedvantagepoint"
	monitor "github.com/oracle/provider-oci/internal/controller/apmsynthetics/monitor"
	onpremisevantagepoint "github.com/oracle/provider-oci/internal/controller/apmsynthetics/onpremisevantagepoint"
	onpremisevantagepointworker "github.com/oracle/provider-oci/internal/controller/apmsynthetics/onpremisevantagepointworker"
	script "github.com/oracle/provider-oci/internal/controller/apmsynthetics/script"
	scheduledquery "github.com/oracle/provider-oci/internal/controller/apmtraces/scheduledquery"
	controlmonitorpluginmanagement "github.com/oracle/provider-oci/internal/controller/appmgmt/controlmonitorpluginmanagement"
	containerconfiguration "github.com/oracle/provider-oci/internal/controller/artifacts/containerconfiguration"
	containerimagesignature "github.com/oracle/provider-oci/internal/controller/artifacts/containerimagesignature"
	containerrepository "github.com/oracle/provider-oci/internal/controller/artifacts/containerrepository"
	genericartifact "github.com/oracle/provider-oci/internal/controller/artifacts/genericartifact"
	repository "github.com/oracle/provider-oci/internal/controller/artifacts/repository"
	configuration "github.com/oracle/provider-oci/internal/controller/audit/configuration"
	autoscalingconfiguration "github.com/oracle/provider-oci/internal/controller/autoscaling/autoscalingconfiguration"
	bastion "github.com/oracle/provider-oci/internal/controller/bastion/bastion"
	session "github.com/oracle/provider-oci/internal/controller/bastion/session"
	autoscalingconfigurationbds "github.com/oracle/provider-oci/internal/controller/bds/autoscalingconfiguration"
	bdscapacityreport "github.com/oracle/provider-oci/internal/controller/bds/bdscapacityreport"
	bdsinstance "github.com/oracle/provider-oci/internal/controller/bds/bdsinstance"
	bdsinstanceapikey "github.com/oracle/provider-oci/internal/controller/bds/bdsinstanceapikey"
	bdsinstanceidentityconfiguration "github.com/oracle/provider-oci/internal/controller/bds/bdsinstanceidentityconfiguration"
	bdsinstancemetastoreconfig "github.com/oracle/provider-oci/internal/controller/bds/bdsinstancemetastoreconfig"
	bdsinstancenodebackup "github.com/oracle/provider-oci/internal/controller/bds/bdsinstancenodebackup"
	bdsinstancenodebackupconfiguration "github.com/oracle/provider-oci/internal/controller/bds/bdsinstancenodebackupconfiguration"
	bdsinstancenodereplaceconfiguration "github.com/oracle/provider-oci/internal/controller/bds/bdsinstancenodereplaceconfiguration"
	bdsinstanceoperationcertificatemanagementsmanagement "github.com/oracle/provider-oci/internal/controller/bds/bdsinstanceoperationcertificatemanagementsmanagement"
	bdsinstanceospatchaction "github.com/oracle/provider-oci/internal/controller/bds/bdsinstanceospatchaction"
	bdsinstancepatchaction "github.com/oracle/provider-oci/internal/controller/bds/bdsinstancepatchaction"
	bdsinstancereplacenodeaction "github.com/oracle/provider-oci/internal/controller/bds/bdsinstancereplacenodeaction"
	bdsinstanceresourceprincipalconfiguration "github.com/oracle/provider-oci/internal/controller/bds/bdsinstanceresourceprincipalconfiguration"
	bdsinstancesoftwareupdateaction "github.com/oracle/provider-oci/internal/controller/bds/bdsinstancesoftwareupdateaction"
	blockchainplatform "github.com/oracle/provider-oci/internal/controller/blockchain/blockchainplatform"
	osn "github.com/oracle/provider-oci/internal/controller/blockchain/osn"
	peer "github.com/oracle/provider-oci/internal/controller/blockchain/peer"
	bootvolume "github.com/oracle/provider-oci/internal/controller/blockstorage/bootvolume"
	bootvolumebackup "github.com/oracle/provider-oci/internal/controller/blockstorage/bootvolumebackup"
	volume "github.com/oracle/provider-oci/internal/controller/blockstorage/volume"
	volumeattachment "github.com/oracle/provider-oci/internal/controller/blockstorage/volumeattachment"
	volumebackup "github.com/oracle/provider-oci/internal/controller/blockstorage/volumebackup"
	volumebackuppolicy "github.com/oracle/provider-oci/internal/controller/blockstorage/volumebackuppolicy"
	volumebackuppolicyassignment "github.com/oracle/provider-oci/internal/controller/blockstorage/volumebackuppolicyassignment"
	volumegroup "github.com/oracle/provider-oci/internal/controller/blockstorage/volumegroup"
	volumegroupbackup "github.com/oracle/provider-oci/internal/controller/blockstorage/volumegroupbackup"
	alertrule "github.com/oracle/provider-oci/internal/controller/budget/alertrule"
	budget "github.com/oracle/provider-oci/internal/controller/budget/budget"
	internaloccmdemandsignal "github.com/oracle/provider-oci/internal/controller/capacitymanagement/internaloccmdemandsignal"
	internaloccmdemandsignaldelivery "github.com/oracle/provider-oci/internal/controller/capacitymanagement/internaloccmdemandsignaldelivery"
	occavailabilitycatalog "github.com/oracle/provider-oci/internal/controller/capacitymanagement/occavailabilitycatalog"
	occcapacityrequest "github.com/oracle/provider-oci/internal/controller/capacitymanagement/occcapacityrequest"
	occcustomergroup "github.com/oracle/provider-oci/internal/controller/capacitymanagement/occcustomergroup"
	occcustomergroupocccustomer "github.com/oracle/provider-oci/internal/controller/capacitymanagement/occcustomergroupocccustomer"
	occmdemandsignal "github.com/oracle/provider-oci/internal/controller/capacitymanagement/occmdemandsignal"
	occmdemandsignalitem "github.com/oracle/provider-oci/internal/controller/capacitymanagement/occmdemandsignalitem"
	cabundle "github.com/oracle/provider-oci/internal/controller/certificatesmanagement/cabundle"
	certificatecertificatesmanagement "github.com/oracle/provider-oci/internal/controller/certificatesmanagement/certificate"
	certificateauthority "github.com/oracle/provider-oci/internal/controller/certificatesmanagement/certificateauthority"
	agent "github.com/oracle/provider-oci/internal/controller/cloudbridge/agent"
	agentdependency "github.com/oracle/provider-oci/internal/controller/cloudbridge/agentdependency"
	agentplugin "github.com/oracle/provider-oci/internal/controller/cloudbridge/agentplugin"
	asset "github.com/oracle/provider-oci/internal/controller/cloudbridge/asset"
	assetsource "github.com/oracle/provider-oci/internal/controller/cloudbridge/assetsource"
	discoveryschedule "github.com/oracle/provider-oci/internal/controller/cloudbridge/discoveryschedule"
	environment "github.com/oracle/provider-oci/internal/controller/cloudbridge/environment"
	inventory "github.com/oracle/provider-oci/internal/controller/cloudbridge/inventory"
	adhocquery "github.com/oracle/provider-oci/internal/controller/cloudguard/adhocquery"
	cloudguardconfiguration "github.com/oracle/provider-oci/internal/controller/cloudguard/cloudguardconfiguration"
	datamaskrule "github.com/oracle/provider-oci/internal/controller/cloudguard/datamaskrule"
	datasource "github.com/oracle/provider-oci/internal/controller/cloudguard/datasource"
	detectorrecipe "github.com/oracle/provider-oci/internal/controller/cloudguard/detectorrecipe"
	managedlist "github.com/oracle/provider-oci/internal/controller/cloudguard/managedlist"
	responderrecipe "github.com/oracle/provider-oci/internal/controller/cloudguard/responderrecipe"
	savedquery "github.com/oracle/provider-oci/internal/controller/cloudguard/savedquery"
	securityrecipe "github.com/oracle/provider-oci/internal/controller/cloudguard/securityrecipe"
	securityzone "github.com/oracle/provider-oci/internal/controller/cloudguard/securityzone"
	target "github.com/oracle/provider-oci/internal/controller/cloudguard/target"
	wlpagent "github.com/oracle/provider-oci/internal/controller/cloudguard/wlpagent"
	migration "github.com/oracle/provider-oci/internal/controller/cloudmigrations/migration"
	migrationasset "github.com/oracle/provider-oci/internal/controller/cloudmigrations/migrationasset"
	migrationplan "github.com/oracle/provider-oci/internal/controller/cloudmigrations/migrationplan"
	replicationschedule "github.com/oracle/provider-oci/internal/controller/cloudmigrations/replicationschedule"
	targetasset "github.com/oracle/provider-oci/internal/controller/cloudmigrations/targetasset"
	clusterplacementgroup "github.com/oracle/provider-oci/internal/controller/clusterplacementgroups/clusterplacementgroup"
	appcataloglistingresourceversionagreement "github.com/oracle/provider-oci/internal/controller/compute/appcataloglistingresourceversionagreement"
	appcatalogsubscription "github.com/oracle/provider-oci/internal/controller/compute/appcatalogsubscription"
	clusternetwork "github.com/oracle/provider-oci/internal/controller/compute/clusternetwork"
	computecapacityreport "github.com/oracle/provider-oci/internal/controller/compute/computecapacityreport"
	computecapacityreservation "github.com/oracle/provider-oci/internal/controller/compute/computecapacityreservation"
	computecapacitytopology "github.com/oracle/provider-oci/internal/controller/compute/computecapacitytopology"
	computecluster "github.com/oracle/provider-oci/internal/controller/compute/computecluster"
	computegpumemorycluster "github.com/oracle/provider-oci/internal/controller/compute/computegpumemorycluster"
	computegpumemoryfabric "github.com/oracle/provider-oci/internal/controller/compute/computegpumemoryfabric"
	computehost "github.com/oracle/provider-oci/internal/controller/compute/computehost"
	computehostgroup "github.com/oracle/provider-oci/internal/controller/compute/computehostgroup"
	computeimagecapabilityschema "github.com/oracle/provider-oci/internal/controller/compute/computeimagecapabilityschema"
	consolehistory "github.com/oracle/provider-oci/internal/controller/compute/consolehistory"
	dedicatedvmhost "github.com/oracle/provider-oci/internal/controller/compute/dedicatedvmhost"
	image "github.com/oracle/provider-oci/internal/controller/compute/image"
	instance "github.com/oracle/provider-oci/internal/controller/compute/instance"
	instanceconfiguration "github.com/oracle/provider-oci/internal/controller/compute/instanceconfiguration"
	instanceconsoleconnection "github.com/oracle/provider-oci/internal/controller/compute/instanceconsoleconnection"
	instancemaintenanceevent "github.com/oracle/provider-oci/internal/controller/compute/instancemaintenanceevent"
	instancepool "github.com/oracle/provider-oci/internal/controller/compute/instancepool"
	instancepoolinstance "github.com/oracle/provider-oci/internal/controller/compute/instancepoolinstance"
	shapemanagement "github.com/oracle/provider-oci/internal/controller/compute/shapemanagement"
	cccinfrastructure "github.com/oracle/provider-oci/internal/controller/computecloudatcustomer/cccinfrastructure"
	cccupgradeschedule "github.com/oracle/provider-oci/internal/controller/computecloudatcustomer/cccupgradeschedule"
	addon "github.com/oracle/provider-oci/internal/controller/containerengine/addon"
	cluster "github.com/oracle/provider-oci/internal/controller/containerengine/cluster"
	clustercompletecredentialrotationmanagement "github.com/oracle/provider-oci/internal/controller/containerengine/clustercompletecredentialrotationmanagement"
	clusterstartcredentialrotationmanagement "github.com/oracle/provider-oci/internal/controller/containerengine/clusterstartcredentialrotationmanagement"
	clusterworkloadmapping "github.com/oracle/provider-oci/internal/controller/containerengine/clusterworkloadmapping"
	nodepool "github.com/oracle/provider-oci/internal/controller/containerengine/nodepool"
	virtualnodepool "github.com/oracle/provider-oci/internal/controller/containerengine/virtualnodepool"
	containerinstance "github.com/oracle/provider-oci/internal/controller/containerinstances/containerinstance"
	byoasn "github.com/oracle/provider-oci/internal/controller/core/byoasn"
	listingresourceversionagreement "github.com/oracle/provider-oci/internal/controller/core/listingresourceversionagreement"
	virtualnetwork "github.com/oracle/provider-oci/internal/controller/core/virtualnetwork"
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
	migrationdatabase "github.com/oracle/provider-oci/internal/controller/database/migration"
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
	catalog "github.com/oracle/provider-oci/internal/controller/datacatalog/catalog"
	catalogprivateendpoint "github.com/oracle/provider-oci/internal/controller/datacatalog/catalogprivateendpoint"
	connection "github.com/oracle/provider-oci/internal/controller/datacatalog/connection"
	dataasset "github.com/oracle/provider-oci/internal/controller/datacatalog/dataasset"
	metastore "github.com/oracle/provider-oci/internal/controller/datacatalog/metastore"
	application "github.com/oracle/provider-oci/internal/controller/dataflow/application"
	invokerun "github.com/oracle/provider-oci/internal/controller/dataflow/invokerun"
	pool "github.com/oracle/provider-oci/internal/controller/dataflow/pool"
	privateendpoint "github.com/oracle/provider-oci/internal/controller/dataflow/privateendpoint"
	runstatement "github.com/oracle/provider-oci/internal/controller/dataflow/runstatement"
	sqlendpoint "github.com/oracle/provider-oci/internal/controller/dataflow/sqlendpoint"
	workspace "github.com/oracle/provider-oci/internal/controller/dataintegration/workspace"
	workspaceapplication "github.com/oracle/provider-oci/internal/controller/dataintegration/workspaceapplication"
	workspaceapplicationpatch "github.com/oracle/provider-oci/internal/controller/dataintegration/workspaceapplicationpatch"
	workspaceapplicationschedule "github.com/oracle/provider-oci/internal/controller/dataintegration/workspaceapplicationschedule"
	workspaceapplicationtaskschedule "github.com/oracle/provider-oci/internal/controller/dataintegration/workspaceapplicationtaskschedule"
	workspaceexportrequest "github.com/oracle/provider-oci/internal/controller/dataintegration/workspaceexportrequest"
	workspacefolder "github.com/oracle/provider-oci/internal/controller/dataintegration/workspacefolder"
	workspaceimportrequest "github.com/oracle/provider-oci/internal/controller/dataintegration/workspaceimportrequest"
	workspaceproject "github.com/oracle/provider-oci/internal/controller/dataintegration/workspaceproject"
	workspacetask "github.com/oracle/provider-oci/internal/controller/dataintegration/workspacetask"
	dataset "github.com/oracle/provider-oci/internal/controller/datalabelingservice/dataset"
	safeaddsdmcolumns "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safeaddsdmcolumns"
	safealert "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safealert"
	safealertpolicy "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safealertpolicy"
	safealertpolicyrule "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safealertpolicyrule"
	safeattributeset "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safeattributeset"
	safeauditarchiveretrieval "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safeauditarchiveretrieval"
	safeauditpolicy "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safeauditpolicy"
	safeauditpolicymanagement "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safeauditpolicymanagement"
	safeauditprofile "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safeauditprofile"
	safeauditprofilemanagement "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safeauditprofilemanagement"
	safeaudittrail "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safeaudittrail"
	safeaudittrailmanagement "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safeaudittrailmanagement"
	safecalculateauditvolumeavailable "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safecalculateauditvolumeavailable"
	safecalculateauditvolumecollected "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safecalculateauditvolumecollected"
	safecomparesecurityassessment "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safecomparesecurityassessment"
	safecompareuserassessment "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safecompareuserassessment"
	safedatabasesecurityconfig "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safedatabasesecurityconfig"
	safedatabasesecurityconfigmanagement "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safedatabasesecurityconfigmanagement"
	safedatasafeconfiguration "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safedatasafeconfiguration"
	safedatasafeprivateendpoint "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safedatasafeprivateendpoint"
	safediscoveryjob "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safediscoveryjob"
	safediscoveryjobsresult "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safediscoveryjobsresult"
	safegenerateonpremconnectorconfiguration "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safegenerateonpremconnectorconfiguration"
	safelibrarymaskingformat "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safelibrarymaskingformat"
	safemaskdata "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safemaskdata"
	safemaskingpoliciesapplydifferencetomaskingcolumns "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safemaskingpoliciesapplydifferencetomaskingcolumns"
	safemaskingpoliciesmaskingcolumn "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safemaskingpoliciesmaskingcolumn"
	safemaskingpolicy "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safemaskingpolicy"
	safemaskingpolicyhealthreportmanagement "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safemaskingpolicyhealthreportmanagement"
	safemaskingreportmanagement "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safemaskingreportmanagement"
	safeonpremconnector "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safeonpremconnector"
	safereport "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safereport"
	safereportdefinition "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safereportdefinition"
	safesdmmaskingpolicydifference "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safesdmmaskingpolicydifference"
	safesecurityassessment "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safesecurityassessment"
	safesecurityassessmentcheck "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safesecurityassessmentcheck"
	safesecurityassessmentfinding "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safesecurityassessmentfinding"
	safesecuritypolicy "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safesecuritypolicy"
	safesecuritypolicyconfig "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safesecuritypolicyconfig"
	safesecuritypolicydeployment "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safesecuritypolicydeployment"
	safesecuritypolicydeploymentmanagement "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safesecuritypolicydeploymentmanagement"
	safesecuritypolicymanagement "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safesecuritypolicymanagement"
	safesensitivedatamodel "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safesensitivedatamodel"
	safesensitivedatamodelreferentialrelation "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safesensitivedatamodelreferentialrelation"
	safesensitivedatamodelsapplydiscoveryjobresults "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safesensitivedatamodelsapplydiscoveryjobresults"
	safesensitivedatamodelssensitivecolumn "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safesensitivedatamodelssensitivecolumn"
	safesensitivetype "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safesensitivetype"
	safesensitivetypegroup "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safesensitivetypegroup"
	safesensitivetypegroupgroupedsensitivetype "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safesensitivetypegroupgroupedsensitivetype"
	safesensitivetypesexport "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safesensitivetypesexport"
	safesetsecurityassessmentbaseline "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safesetsecurityassessmentbaseline"
	safesetsecurityassessmentbaselinemanagement "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safesetsecurityassessmentbaselinemanagement"
	safesetuserassessmentbaseline "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safesetuserassessmentbaseline"
	safesetuserassessmentbaselinemanagement "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safesetuserassessmentbaselinemanagement"
	safesqlcollection "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safesqlcollection"
	safesqlfirewallpolicy "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safesqlfirewallpolicy"
	safesqlfirewallpolicymanagement "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safesqlfirewallpolicymanagement"
	safetargetalertpolicyassociation "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safetargetalertpolicyassociation"
	safetargetdatabase "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safetargetdatabase"
	safetargetdatabasegroup "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safetargetdatabasegroup"
	safetargetdatabasepeertargetdatabase "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safetargetdatabasepeertargetdatabase"
	safeunifiedauditpolicy "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safeunifiedauditpolicy"
	safeunifiedauditpolicydefinition "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safeunifiedauditpolicydefinition"
	safeunsetsecurityassessmentbaseline "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safeunsetsecurityassessmentbaseline"
	safeunsetsecurityassessmentbaselinemanagement "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safeunsetsecurityassessmentbaselinemanagement"
	safeunsetuserassessmentbaseline "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safeunsetuserassessmentbaseline"
	safeunsetuserassessmentbaselinemanagement "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safeunsetuserassessmentbaselinemanagement"
	safeuserassessment "github.com/oracle/provider-oci/internal/controller/datalabelingservice/safeuserassessment"
	jobdatascience "github.com/oracle/provider-oci/internal/controller/datascience/job"
	jobrun "github.com/oracle/provider-oci/internal/controller/datascience/jobrun"
	mlapplication "github.com/oracle/provider-oci/internal/controller/datascience/mlapplication"
	mlapplicationimplementation "github.com/oracle/provider-oci/internal/controller/datascience/mlapplicationimplementation"
	mlapplicationinstance "github.com/oracle/provider-oci/internal/controller/datascience/mlapplicationinstance"
	modeldatascience "github.com/oracle/provider-oci/internal/controller/datascience/model"
	modelartifactexport "github.com/oracle/provider-oci/internal/controller/datascience/modelartifactexport"
	modelartifactimport "github.com/oracle/provider-oci/internal/controller/datascience/modelartifactimport"
	modelcustommetadataartifact "github.com/oracle/provider-oci/internal/controller/datascience/modelcustommetadataartifact"
	modeldefinedmetadataartifact "github.com/oracle/provider-oci/internal/controller/datascience/modeldefinedmetadataartifact"
	modeldeployment "github.com/oracle/provider-oci/internal/controller/datascience/modeldeployment"
	modelgroup "github.com/oracle/provider-oci/internal/controller/datascience/modelgroup"
	modelgroupartifact "github.com/oracle/provider-oci/internal/controller/datascience/modelgroupartifact"
	modelgroupversionhistory "github.com/oracle/provider-oci/internal/controller/datascience/modelgroupversionhistory"
	modelprovenance "github.com/oracle/provider-oci/internal/controller/datascience/modelprovenance"
	modelversionset "github.com/oracle/provider-oci/internal/controller/datascience/modelversionset"
	notebooksession "github.com/oracle/provider-oci/internal/controller/datascience/notebooksession"
	pipeline "github.com/oracle/provider-oci/internal/controller/datascience/pipeline"
	pipelinerun "github.com/oracle/provider-oci/internal/controller/datascience/pipelinerun"
	privateendpointdatascience "github.com/oracle/provider-oci/internal/controller/datascience/privateendpoint"
	projectdatascience "github.com/oracle/provider-oci/internal/controller/datascience/project"
	schedule "github.com/oracle/provider-oci/internal/controller/datascience/schedule"
	vulnerabilityscan "github.com/oracle/provider-oci/internal/controller/dblm/vulnerabilityscan"
	multicloudresourcediscovery "github.com/oracle/provider-oci/internal/controller/dbmulticloud/multicloudresourcediscovery"
	oracledbawsidentityconnector "github.com/oracle/provider-oci/internal/controller/dbmulticloud/oracledbawsidentityconnector"
	oracledbawskey "github.com/oracle/provider-oci/internal/controller/dbmulticloud/oracledbawskey"
	oracledbazureblobcontainer "github.com/oracle/provider-oci/internal/controller/dbmulticloud/oracledbazureblobcontainer"
	oracledbazureblobmount "github.com/oracle/provider-oci/internal/controller/dbmulticloud/oracledbazureblobmount"
	oracledbazureconnector "github.com/oracle/provider-oci/internal/controller/dbmulticloud/oracledbazureconnector"
	oracledbazurevault "github.com/oracle/provider-oci/internal/controller/dbmulticloud/oracledbazurevault"
	oracledbazurevaultassociation "github.com/oracle/provider-oci/internal/controller/dbmulticloud/oracledbazurevaultassociation"
	oracledbgcpidentityconnector "github.com/oracle/provider-oci/internal/controller/dbmulticloud/oracledbgcpidentityconnector"
	oracledbgcpkeyring "github.com/oracle/provider-oci/internal/controller/dbmulticloud/oracledbgcpkeyring"
	delegationcontrol "github.com/oracle/provider-oci/internal/controller/delegateaccesscontrol/delegationcontrol"
	delegationsubscription "github.com/oracle/provider-oci/internal/controller/delegateaccesscontrol/delegationsubscription"
	occdemandsignal "github.com/oracle/provider-oci/internal/controller/demandsignal/occdemandsignal"
	desktoppool "github.com/oracle/provider-oci/internal/controller/desktops/desktoppool"
	buildpipeline "github.com/oracle/provider-oci/internal/controller/devops/buildpipeline"
	buildpipelinestage "github.com/oracle/provider-oci/internal/controller/devops/buildpipelinestage"
	buildrun "github.com/oracle/provider-oci/internal/controller/devops/buildrun"
	connectiondevops "github.com/oracle/provider-oci/internal/controller/devops/connection"
	deployartifact "github.com/oracle/provider-oci/internal/controller/devops/deployartifact"
	deployenvironment "github.com/oracle/provider-oci/internal/controller/devops/deployenvironment"
	deploymentdevops "github.com/oracle/provider-oci/internal/controller/devops/deployment"
	deploypipeline "github.com/oracle/provider-oci/internal/controller/devops/deploypipeline"
	deploystage "github.com/oracle/provider-oci/internal/controller/devops/deploystage"
	projectdevops "github.com/oracle/provider-oci/internal/controller/devops/project"
	projectrepositorysetting "github.com/oracle/provider-oci/internal/controller/devops/projectrepositorysetting"
	repositorydevops "github.com/oracle/provider-oci/internal/controller/devops/repository"
	repositorymirror "github.com/oracle/provider-oci/internal/controller/devops/repositorymirror"
	repositoryprotectedbranchmanagement "github.com/oracle/provider-oci/internal/controller/devops/repositoryprotectedbranchmanagement"
	repositoryref "github.com/oracle/provider-oci/internal/controller/devops/repositoryref"
	repositorysetting "github.com/oracle/provider-oci/internal/controller/devops/repositorysetting"
	trigger "github.com/oracle/provider-oci/internal/controller/devops/trigger"
	stack "github.com/oracle/provider-oci/internal/controller/dif/stack"
	automaticdrconfiguration "github.com/oracle/provider-oci/internal/controller/disasterrecovery/automaticdrconfiguration"
	drplan "github.com/oracle/provider-oci/internal/controller/disasterrecovery/drplan"
	drplanexecution "github.com/oracle/provider-oci/internal/controller/disasterrecovery/drplanexecution"
	drprotectiongroup "github.com/oracle/provider-oci/internal/controller/disasterrecovery/drprotectiongroup"
	actioncreatezonefromzonefile "github.com/oracle/provider-oci/internal/controller/dns/actioncreatezonefromzonefile"
	record "github.com/oracle/provider-oci/internal/controller/dns/record"
	resolver "github.com/oracle/provider-oci/internal/controller/dns/resolver"
	resolverendpoint "github.com/oracle/provider-oci/internal/controller/dns/resolverendpoint"
	rrset "github.com/oracle/provider-oci/internal/controller/dns/rrset"
	steeringpolicy "github.com/oracle/provider-oci/internal/controller/dns/steeringpolicy"
	steeringpolicyattachment "github.com/oracle/provider-oci/internal/controller/dns/steeringpolicyattachment"
	tsigkey "github.com/oracle/provider-oci/internal/controller/dns/tsigkey"
	view "github.com/oracle/provider-oci/internal/controller/dns/view"
	zone "github.com/oracle/provider-oci/internal/controller/dns/zone"
	zonepromotednsseckeyversion "github.com/oracle/provider-oci/internal/controller/dns/zonepromotednsseckeyversion"
	zonestagednsseckeyversion "github.com/oracle/provider-oci/internal/controller/dns/zonestagednsseckeyversion"
	dkim "github.com/oracle/provider-oci/internal/controller/email/dkim"
	emaildomain "github.com/oracle/provider-oci/internal/controller/email/emaildomain"
	emailreturnpath "github.com/oracle/provider-oci/internal/controller/email/emailreturnpath"
	sender "github.com/oracle/provider-oci/internal/controller/email/sender"
	suppression "github.com/oracle/provider-oci/internal/controller/email/suppression"
	rule "github.com/oracle/provider-oci/internal/controller/events/rule"
	export "github.com/oracle/provider-oci/internal/controller/filestorage/export"
	exportset "github.com/oracle/provider-oci/internal/controller/filestorage/exportset"
	filesystem "github.com/oracle/provider-oci/internal/controller/filestorage/filesystem"
	filesystemquotarule "github.com/oracle/provider-oci/internal/controller/filestorage/filesystemquotarule"
	filesystemsnapshotpolicy "github.com/oracle/provider-oci/internal/controller/filestorage/filesystemsnapshotpolicy"
	mounttarget "github.com/oracle/provider-oci/internal/controller/filestorage/mounttarget"
	outboundconnector "github.com/oracle/provider-oci/internal/controller/filestorage/outboundconnector"
	replication "github.com/oracle/provider-oci/internal/controller/filestorage/replication"
	snapshot "github.com/oracle/provider-oci/internal/controller/filestorage/snapshot"
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
	fsucollection "github.com/oracle/provider-oci/internal/controller/fleetsoftwareupdate/fsucollection"
	fsucycle "github.com/oracle/provider-oci/internal/controller/fleetsoftwareupdate/fsucycle"
	applicationfunctions "github.com/oracle/provider-oci/internal/controller/functions/application"
	function "github.com/oracle/provider-oci/internal/controller/functions/function"
	invokefunction "github.com/oracle/provider-oci/internal/controller/functions/invokefunction"
	fusionenvironment "github.com/oracle/provider-oci/internal/controller/fusionapps/fusionenvironment"
	fusionenvironmentadminuser "github.com/oracle/provider-oci/internal/controller/fusionapps/fusionenvironmentadminuser"
	fusionenvironmentdatamaskingactivity "github.com/oracle/provider-oci/internal/controller/fusionapps/fusionenvironmentdatamaskingactivity"
	fusionenvironmentfamily "github.com/oracle/provider-oci/internal/controller/fusionapps/fusionenvironmentfamily"
	fusionenvironmentrefreshactivity "github.com/oracle/provider-oci/internal/controller/fusionapps/fusionenvironmentrefreshactivity"
	fusionenvironmentserviceattachment "github.com/oracle/provider-oci/internal/controller/fusionapps/fusionenvironmentserviceattachment"
	agentagent "github.com/oracle/provider-oci/internal/controller/generativeai/agentagent"
	agentagentendpoint "github.com/oracle/provider-oci/internal/controller/generativeai/agentagentendpoint"
	agentdataingestionjob "github.com/oracle/provider-oci/internal/controller/generativeai/agentdataingestionjob"
	agentdatasource "github.com/oracle/provider-oci/internal/controller/generativeai/agentdatasource"
	agentknowledgebase "github.com/oracle/provider-oci/internal/controller/generativeai/agentknowledgebase"
	agenttool "github.com/oracle/provider-oci/internal/controller/generativeai/agenttool"
	dedicatedaicluster "github.com/oracle/provider-oci/internal/controller/generativeai/dedicatedaicluster"
	endpointgenerativeai "github.com/oracle/provider-oci/internal/controller/generativeai/endpoint"
	generativeaiprivateendpoint "github.com/oracle/provider-oci/internal/controller/generativeai/generativeaiprivateendpoint"
	modelgenerativeai "github.com/oracle/provider-oci/internal/controller/generativeai/model"
	artifactbypath "github.com/oracle/provider-oci/internal/controller/genericartifactscontent/artifactbypath"
	privateendpointgloballydistributeddatabase "github.com/oracle/provider-oci/internal/controller/globallydistributeddatabase/privateendpoint"
	shardeddatabase "github.com/oracle/provider-oci/internal/controller/globallydistributeddatabase/shardeddatabase"
	connectiongoldengate "github.com/oracle/provider-oci/internal/controller/goldengate/connection"
	connectionassignment "github.com/oracle/provider-oci/internal/controller/goldengate/connectionassignment"
	databaseregistration "github.com/oracle/provider-oci/internal/controller/goldengate/databaseregistration"
	deploymentgoldengate "github.com/oracle/provider-oci/internal/controller/goldengate/deployment"
	deploymentbackup "github.com/oracle/provider-oci/internal/controller/goldengate/deploymentbackup"
	deploymentcertificate "github.com/oracle/provider-oci/internal/controller/goldengate/deploymentcertificate"
	pipelinegoldengate "github.com/oracle/provider-oci/internal/controller/goldengate/pipeline"
	httpmonitor "github.com/oracle/provider-oci/internal/controller/healthchecks/httpmonitor"
	httpprobe "github.com/oracle/provider-oci/internal/controller/healthchecks/httpprobe"
	pingmonitor "github.com/oracle/provider-oci/internal/controller/healthchecks/pingmonitor"
	pingprobe "github.com/oracle/provider-oci/internal/controller/healthchecks/pingprobe"
	apikey "github.com/oracle/provider-oci/internal/controller/identity/apikey"
	authenticationpolicy "github.com/oracle/provider-oci/internal/controller/identity/authenticationpolicy"
	authtoken "github.com/oracle/provider-oci/internal/controller/identity/authtoken"
	compartment "github.com/oracle/provider-oci/internal/controller/identity/compartment"
	customersecretkey "github.com/oracle/provider-oci/internal/controller/identity/customersecretkey"
	dbcredential "github.com/oracle/provider-oci/internal/controller/identity/dbcredential"
	domain "github.com/oracle/provider-oci/internal/controller/identity/domain"
	domainreplicationtoregion "github.com/oracle/provider-oci/internal/controller/identity/domainreplicationtoregion"
	dynamicgroup "github.com/oracle/provider-oci/internal/controller/identity/dynamicgroup"
	group "github.com/oracle/provider-oci/internal/controller/identity/group"
	identityprovider "github.com/oracle/provider-oci/internal/controller/identity/identityprovider"
	idpgroupmapping "github.com/oracle/provider-oci/internal/controller/identity/idpgroupmapping"
	importstandardtagsmanagement "github.com/oracle/provider-oci/internal/controller/identity/importstandardtagsmanagement"
	networksource "github.com/oracle/provider-oci/internal/controller/identity/networksource"
	policy "github.com/oracle/provider-oci/internal/controller/identity/policy"
	smtpcredential "github.com/oracle/provider-oci/internal/controller/identity/smtpcredential"
	tag "github.com/oracle/provider-oci/internal/controller/identity/tag"
	tagdefault "github.com/oracle/provider-oci/internal/controller/identity/tagdefault"
	tagnamespace "github.com/oracle/provider-oci/internal/controller/identity/tagnamespace"
	uipassword "github.com/oracle/provider-oci/internal/controller/identity/uipassword"
	user "github.com/oracle/provider-oci/internal/controller/identity/user"
	usercapabilitiesmanagement "github.com/oracle/provider-oci/internal/controller/identity/usercapabilitiesmanagement"
	usergroupmembership "github.com/oracle/provider-oci/internal/controller/identity/usergroupmembership"
	generatescopedaccesstoken "github.com/oracle/provider-oci/internal/controller/identitydataplane/generatescopedaccesstoken"
	accountrecoverysetting "github.com/oracle/provider-oci/internal/controller/identitydomains/accountrecoverysetting"
	apikeyidentitydomains "github.com/oracle/provider-oci/internal/controller/identitydomains/apikey"
	app "github.com/oracle/provider-oci/internal/controller/identitydomains/app"
	approle "github.com/oracle/provider-oci/internal/controller/identitydomains/approle"
	approvalworkflow "github.com/oracle/provider-oci/internal/controller/identitydomains/approvalworkflow"
	approvalworkflowassignment "github.com/oracle/provider-oci/internal/controller/identitydomains/approvalworkflowassignment"
	approvalworkflowstep "github.com/oracle/provider-oci/internal/controller/identitydomains/approvalworkflowstep"
	authenticationfactorsetting "github.com/oracle/provider-oci/internal/controller/identitydomains/authenticationfactorsetting"
	authtokenidentitydomains "github.com/oracle/provider-oci/internal/controller/identitydomains/authtoken"
	cloudgate "github.com/oracle/provider-oci/internal/controller/identitydomains/cloudgate"
	cloudgatemapping "github.com/oracle/provider-oci/internal/controller/identitydomains/cloudgatemapping"
	cloudgateserver "github.com/oracle/provider-oci/internal/controller/identitydomains/cloudgateserver"
	condition "github.com/oracle/provider-oci/internal/controller/identitydomains/condition"
	customersecretkeyidentitydomains "github.com/oracle/provider-oci/internal/controller/identitydomains/customersecretkey"
	dynamicresourcegroup "github.com/oracle/provider-oci/internal/controller/identitydomains/dynamicresourcegroup"
	grant "github.com/oracle/provider-oci/internal/controller/identitydomains/grant"
	groupidentitydomains "github.com/oracle/provider-oci/internal/controller/identitydomains/group"
	identitypropagationtrust "github.com/oracle/provider-oci/internal/controller/identitydomains/identitypropagationtrust"
	identityprovideridentitydomains "github.com/oracle/provider-oci/internal/controller/identitydomains/identityprovider"
	identitysetting "github.com/oracle/provider-oci/internal/controller/identitydomains/identitysetting"
	kmsisetting "github.com/oracle/provider-oci/internal/controller/identitydomains/kmsisetting"
	myapikey "github.com/oracle/provider-oci/internal/controller/identitydomains/myapikey"
	myauthtoken "github.com/oracle/provider-oci/internal/controller/identitydomains/myauthtoken"
	mycustomersecretkey "github.com/oracle/provider-oci/internal/controller/identitydomains/mycustomersecretkey"
	myoauth2clientcredential "github.com/oracle/provider-oci/internal/controller/identitydomains/myoauth2clientcredential"
	myrequest "github.com/oracle/provider-oci/internal/controller/identitydomains/myrequest"
	mysmtpcredential "github.com/oracle/provider-oci/internal/controller/identitydomains/mysmtpcredential"
	mysupportaccount "github.com/oracle/provider-oci/internal/controller/identitydomains/mysupportaccount"
	myuserdbcredential "github.com/oracle/provider-oci/internal/controller/identitydomains/myuserdbcredential"
	networkperimeter "github.com/oracle/provider-oci/internal/controller/identitydomains/networkperimeter"
	notificationsetting "github.com/oracle/provider-oci/internal/controller/identitydomains/notificationsetting"
	oauth2clientcredential "github.com/oracle/provider-oci/internal/controller/identitydomains/oauth2clientcredential"
	oauthclientcertificate "github.com/oracle/provider-oci/internal/controller/identitydomains/oauthclientcertificate"
	oauthpartnercertificate "github.com/oracle/provider-oci/internal/controller/identitydomains/oauthpartnercertificate"
	passwordpolicy "github.com/oracle/provider-oci/internal/controller/identitydomains/passwordpolicy"
	policyidentitydomains "github.com/oracle/provider-oci/internal/controller/identitydomains/policy"
	ruleidentitydomains "github.com/oracle/provider-oci/internal/controller/identitydomains/rule"
	securityquestion "github.com/oracle/provider-oci/internal/controller/identitydomains/securityquestion"
	securityquestionsetting "github.com/oracle/provider-oci/internal/controller/identitydomains/securityquestionsetting"
	selfregistrationprofile "github.com/oracle/provider-oci/internal/controller/identitydomains/selfregistrationprofile"
	setting "github.com/oracle/provider-oci/internal/controller/identitydomains/setting"
	smtpcredentialidentitydomains "github.com/oracle/provider-oci/internal/controller/identitydomains/smtpcredential"
	socialidentityprovider "github.com/oracle/provider-oci/internal/controller/identitydomains/socialidentityprovider"
	useridentitydomains "github.com/oracle/provider-oci/internal/controller/identitydomains/user"
	userdbcredential "github.com/oracle/provider-oci/internal/controller/identitydomains/userdbcredential"
	integrationinstance "github.com/oracle/provider-oci/internal/controller/integration/integrationinstance"
	oraclemanagedcustomendpoint "github.com/oracle/provider-oci/internal/controller/integration/oraclemanagedcustomendpoint"
	privateendpointoutboundconnection "github.com/oracle/provider-oci/internal/controller/integration/privateendpointoutboundconnection"
	digitaltwinadapter "github.com/oracle/provider-oci/internal/controller/iot/digitaltwinadapter"
	digitaltwininstance "github.com/oracle/provider-oci/internal/controller/iot/digitaltwininstance"
	digitaltwininstanceinvokerawcommand "github.com/oracle/provider-oci/internal/controller/iot/digitaltwininstanceinvokerawcommand"
	digitaltwinmodel "github.com/oracle/provider-oci/internal/controller/iot/digitaltwinmodel"
	digitaltwinrelationship "github.com/oracle/provider-oci/internal/controller/iot/digitaltwinrelationship"
	iotdomain "github.com/oracle/provider-oci/internal/controller/iot/iotdomain"
	iotdomainchangedataretentionperiod "github.com/oracle/provider-oci/internal/controller/iot/iotdomainchangedataretentionperiod"
	iotdomainconfiguredataaccess "github.com/oracle/provider-oci/internal/controller/iot/iotdomainconfiguredataaccess"
	iotdomaingroup "github.com/oracle/provider-oci/internal/controller/iot/iotdomaingroup"
	iotdomaingroupconfiguredataaccess "github.com/oracle/provider-oci/internal/controller/iot/iotdomaingroupconfiguredataaccess"
	fleetjms "github.com/oracle/provider-oci/internal/controller/jms/fleet"
	fleetadvancedfeatureconfiguration "github.com/oracle/provider-oci/internal/controller/jms/fleetadvancedfeatureconfiguration"
	fleetagentconfiguration "github.com/oracle/provider-oci/internal/controller/jms/fleetagentconfiguration"
	jmsplugin "github.com/oracle/provider-oci/internal/controller/jms/jmsplugin"
	taskschedule "github.com/oracle/provider-oci/internal/controller/jms/taskschedule"
	javadownloadreport "github.com/oracle/provider-oci/internal/controller/jmsjavadownloads/javadownloadreport"
	javadownloadtoken "github.com/oracle/provider-oci/internal/controller/jmsjavadownloads/javadownloadtoken"
	javalicenseacceptancerecord "github.com/oracle/provider-oci/internal/controller/jmsjavadownloads/javalicenseacceptancerecord"
	analyzeapplicationsconfiguration "github.com/oracle/provider-oci/internal/controller/jmsutils/analyzeapplicationsconfiguration"
	subscriptionacknowledgmentconfiguration "github.com/oracle/provider-oci/internal/controller/jmsutils/subscriptionacknowledgmentconfiguration"
	ekmsprivateendpoint "github.com/oracle/provider-oci/internal/controller/kms/ekmsprivateendpoint"
	encrypteddata "github.com/oracle/provider-oci/internal/controller/kms/encrypteddata"
	generatedkey "github.com/oracle/provider-oci/internal/controller/kms/generatedkey"
	key "github.com/oracle/provider-oci/internal/controller/kms/key"
	keyversion "github.com/oracle/provider-oci/internal/controller/kms/keyversion"
	sign "github.com/oracle/provider-oci/internal/controller/kms/sign"
	vault "github.com/oracle/provider-oci/internal/controller/kms/vault"
	vaultreplication "github.com/oracle/provider-oci/internal/controller/kms/vaultreplication"
	verify "github.com/oracle/provider-oci/internal/controller/kms/verify"
	configurationlicensemanager "github.com/oracle/provider-oci/internal/controller/licensemanager/configuration"
	licenserecord "github.com/oracle/provider-oci/internal/controller/licensemanager/licenserecord"
	productlicense "github.com/oracle/provider-oci/internal/controller/licensemanager/productlicense"
	quota "github.com/oracle/provider-oci/internal/controller/limits/quota"
	backend "github.com/oracle/provider-oci/internal/controller/loadbalancer/backend"
	backendset "github.com/oracle/provider-oci/internal/controller/loadbalancer/backendset"
	certificateloadbalancer "github.com/oracle/provider-oci/internal/controller/loadbalancer/certificate"
	lbhostname "github.com/oracle/provider-oci/internal/controller/loadbalancer/lbhostname"
	listener "github.com/oracle/provider-oci/internal/controller/loadbalancer/listener"
	loadbalancer "github.com/oracle/provider-oci/internal/controller/loadbalancer/loadbalancer"
	pathrouteset "github.com/oracle/provider-oci/internal/controller/loadbalancer/pathrouteset"
	routingpolicy "github.com/oracle/provider-oci/internal/controller/loadbalancer/routingpolicy"
	ruleset "github.com/oracle/provider-oci/internal/controller/loadbalancer/ruleset"
	sslciphersuite "github.com/oracle/provider-oci/internal/controller/loadbalancer/sslciphersuite"
	loganalyticsentity "github.com/oracle/provider-oci/internal/controller/loganalytics/loganalyticsentity"
	loganalyticsentitytype "github.com/oracle/provider-oci/internal/controller/loganalytics/loganalyticsentitytype"
	loganalyticsimportcustomcontent "github.com/oracle/provider-oci/internal/controller/loganalytics/loganalyticsimportcustomcontent"
	loganalyticsloggroup "github.com/oracle/provider-oci/internal/controller/loganalytics/loganalyticsloggroup"
	loganalyticsobjectcollectionrule "github.com/oracle/provider-oci/internal/controller/loganalytics/loganalyticsobjectcollectionrule"
	loganalyticspreferencesmanagement "github.com/oracle/provider-oci/internal/controller/loganalytics/loganalyticspreferencesmanagement"
	loganalyticsresourcecategoriesmanagement "github.com/oracle/provider-oci/internal/controller/loganalytics/loganalyticsresourcecategoriesmanagement"
	loganalyticsunprocesseddatabucketmanagement "github.com/oracle/provider-oci/internal/controller/loganalytics/loganalyticsunprocesseddatabucketmanagement"
	namespace "github.com/oracle/provider-oci/internal/controller/loganalytics/namespace"
	namespaceingesttimerule "github.com/oracle/provider-oci/internal/controller/loganalytics/namespaceingesttimerule"
	namespaceingesttimerulesmanagement "github.com/oracle/provider-oci/internal/controller/loganalytics/namespaceingesttimerulesmanagement"
	namespacelookup "github.com/oracle/provider-oci/internal/controller/loganalytics/namespacelookup"
	namespacelookupsappenddatamanagement "github.com/oracle/provider-oci/internal/controller/loganalytics/namespacelookupsappenddatamanagement"
	namespacelookupsupdatedatamanagement "github.com/oracle/provider-oci/internal/controller/loganalytics/namespacelookupsupdatedatamanagement"
	namespacescheduledtask "github.com/oracle/provider-oci/internal/controller/loganalytics/namespacescheduledtask"
	namespacestoragearchivalconfig "github.com/oracle/provider-oci/internal/controller/loganalytics/namespacestoragearchivalconfig"
	namespacestorageenabledisablearchiving "github.com/oracle/provider-oci/internal/controller/loganalytics/namespacestorageenabledisablearchiving"
	log "github.com/oracle/provider-oci/internal/controller/logging/log"
	loggroup "github.com/oracle/provider-oci/internal/controller/logging/loggroup"
	logsavedsearch "github.com/oracle/provider-oci/internal/controller/logging/logsavedsearch"
	unifiedagentconfiguration "github.com/oracle/provider-oci/internal/controller/logging/unifiedagentconfiguration"
	lustrefilesystem "github.com/oracle/provider-oci/internal/controller/lustrefilestorage/lustrefilesystem"
	objectstoragelink "github.com/oracle/provider-oci/internal/controller/lustrefilestorage/objectstoragelink"
	kafkacluster "github.com/oracle/provider-oci/internal/controller/managedkafka/kafkacluster"
	kafkaclusterconfig "github.com/oracle/provider-oci/internal/controller/managedkafka/kafkaclusterconfig"
	kafkaclustersuperusersmanagement "github.com/oracle/provider-oci/internal/controller/managedkafka/kafkaclustersuperusersmanagement"
	managementagent "github.com/oracle/provider-oci/internal/controller/managementagent/managementagent"
	managementagentdatasource "github.com/oracle/provider-oci/internal/controller/managementagent/managementagentdatasource"
	managementagentinstallkey "github.com/oracle/provider-oci/internal/controller/managementagent/managementagentinstallkey"
	namedcredential "github.com/oracle/provider-oci/internal/controller/managementagent/namedcredential"
	managementdashboardsimport "github.com/oracle/provider-oci/internal/controller/managementdashboard/managementdashboardsimport"
	acceptedagreement "github.com/oracle/provider-oci/internal/controller/marketplace/acceptedagreement"
	listingpackageagreement "github.com/oracle/provider-oci/internal/controller/marketplace/listingpackageagreement"
	marketplaceexternalattestedmetadata "github.com/oracle/provider-oci/internal/controller/marketplace/marketplaceexternalattestedmetadata"
	publication "github.com/oracle/provider-oci/internal/controller/marketplace/publication"
	mediaasset "github.com/oracle/provider-oci/internal/controller/mediaservices/mediaasset"
	mediaworkflow "github.com/oracle/provider-oci/internal/controller/mediaservices/mediaworkflow"
	mediaworkflowconfiguration "github.com/oracle/provider-oci/internal/controller/mediaservices/mediaworkflowconfiguration"
	mediaworkflowjob "github.com/oracle/provider-oci/internal/controller/mediaservices/mediaworkflowjob"
	streamcdnconfig "github.com/oracle/provider-oci/internal/controller/mediaservices/streamcdnconfig"
	streamdistributionchannel "github.com/oracle/provider-oci/internal/controller/mediaservices/streamdistributionchannel"
	streampackagingconfig "github.com/oracle/provider-oci/internal/controller/mediaservices/streampackagingconfig"
	customtable "github.com/oracle/provider-oci/internal/controller/meteringcomputation/customtable"
	query "github.com/oracle/provider-oci/internal/controller/meteringcomputation/query"
	schedulemeteringcomputation "github.com/oracle/provider-oci/internal/controller/meteringcomputation/schedule"
	usage "github.com/oracle/provider-oci/internal/controller/meteringcomputation/usage"
	usagecarbonemission "github.com/oracle/provider-oci/internal/controller/meteringcomputation/usagecarbonemission"
	usagecarbonemissionsquery "github.com/oracle/provider-oci/internal/controller/meteringcomputation/usagecarbonemissionsquery"
	usagestatementemailrecipientsgroup "github.com/oracle/provider-oci/internal/controller/meteringcomputation/usagestatementemailrecipientsgroup"
	alarm "github.com/oracle/provider-oci/internal/controller/monitoring/alarm"
	alarmsuppression "github.com/oracle/provider-oci/internal/controller/monitoring/alarmsuppression"
	capturefilter "github.com/oracle/provider-oci/internal/controller/monitoring/capturefilter"
	vtap "github.com/oracle/provider-oci/internal/controller/monitoring/vtap"
	mysqlbackup "github.com/oracle/provider-oci/internal/controller/mysql/mysqlbackup"
	mysqlchannel "github.com/oracle/provider-oci/internal/controller/mysql/mysqlchannel"
	mysqlconfiguration "github.com/oracle/provider-oci/internal/controller/mysql/mysqlconfiguration"
	mysqldbsystem "github.com/oracle/provider-oci/internal/controller/mysql/mysqldbsystem"
	mysqlheatwavecluster "github.com/oracle/provider-oci/internal/controller/mysql/mysqlheatwavecluster"
	mysqlreplica "github.com/oracle/provider-oci/internal/controller/mysql/mysqlreplica"
	cpe "github.com/oracle/provider-oci/internal/controller/networkconnectivity/cpe"
	crossconnect "github.com/oracle/provider-oci/internal/controller/networkconnectivity/crossconnect"
	crossconnectgroup "github.com/oracle/provider-oci/internal/controller/networkconnectivity/crossconnectgroup"
	drg "github.com/oracle/provider-oci/internal/controller/networkconnectivity/drg"
	drgattachment "github.com/oracle/provider-oci/internal/controller/networkconnectivity/drgattachment"
	drgattachmentmanagement "github.com/oracle/provider-oci/internal/controller/networkconnectivity/drgattachmentmanagement"
	drgattachmentslist "github.com/oracle/provider-oci/internal/controller/networkconnectivity/drgattachmentslist"
	drgroutedistribution "github.com/oracle/provider-oci/internal/controller/networkconnectivity/drgroutedistribution"
	drgroutedistributionstatement "github.com/oracle/provider-oci/internal/controller/networkconnectivity/drgroutedistributionstatement"
	drgroutetable "github.com/oracle/provider-oci/internal/controller/networkconnectivity/drgroutetable"
	drgroutetablerouterule "github.com/oracle/provider-oci/internal/controller/networkconnectivity/drgroutetablerouterule"
	ipsec "github.com/oracle/provider-oci/internal/controller/networkconnectivity/ipsec"
	ipsecconnectiontunnelmanagement "github.com/oracle/provider-oci/internal/controller/networkconnectivity/ipsecconnectiontunnelmanagement"
	virtualcircuit "github.com/oracle/provider-oci/internal/controller/networkconnectivity/virtualcircuit"
	networkfirewall "github.com/oracle/provider-oci/internal/controller/networkfirewall/networkfirewall"
	networkfirewallpolicy "github.com/oracle/provider-oci/internal/controller/networkfirewall/networkfirewallpolicy"
	networkfirewallpolicyaddresslist "github.com/oracle/provider-oci/internal/controller/networkfirewall/networkfirewallpolicyaddresslist"
	networkfirewallpolicyapplication "github.com/oracle/provider-oci/internal/controller/networkfirewall/networkfirewallpolicyapplication"
	networkfirewallpolicyapplicationgroup "github.com/oracle/provider-oci/internal/controller/networkfirewall/networkfirewallpolicyapplicationgroup"
	networkfirewallpolicydecryptionprofile "github.com/oracle/provider-oci/internal/controller/networkfirewall/networkfirewallpolicydecryptionprofile"
	networkfirewallpolicydecryptionrule "github.com/oracle/provider-oci/internal/controller/networkfirewall/networkfirewallpolicydecryptionrule"
	networkfirewallpolicymappedsecret "github.com/oracle/provider-oci/internal/controller/networkfirewall/networkfirewallpolicymappedsecret"
	networkfirewallpolicynatrule "github.com/oracle/provider-oci/internal/controller/networkfirewall/networkfirewallpolicynatrule"
	networkfirewallpolicysecurityrule "github.com/oracle/provider-oci/internal/controller/networkfirewall/networkfirewallpolicysecurityrule"
	networkfirewallpolicyservice "github.com/oracle/provider-oci/internal/controller/networkfirewall/networkfirewallpolicyservice"
	networkfirewallpolicytunnelinspectionrule "github.com/oracle/provider-oci/internal/controller/networkfirewall/networkfirewallpolicytunnelinspectionrule"
	defaultdhcpoptions "github.com/oracle/provider-oci/internal/controller/networking/defaultdhcpoptions"
	defaultroutetable "github.com/oracle/provider-oci/internal/controller/networking/defaultroutetable"
	defaultsecuritylist "github.com/oracle/provider-oci/internal/controller/networking/defaultsecuritylist"
	dhcpoptions "github.com/oracle/provider-oci/internal/controller/networking/dhcpoptions"
	internetgateway "github.com/oracle/provider-oci/internal/controller/networking/internetgateway"
	ipv6 "github.com/oracle/provider-oci/internal/controller/networking/ipv6"
	localpeeringgateway "github.com/oracle/provider-oci/internal/controller/networking/localpeeringgateway"
	natgateway "github.com/oracle/provider-oci/internal/controller/networking/natgateway"
	networksecuritygroup "github.com/oracle/provider-oci/internal/controller/networking/networksecuritygroup"
	networksecuritygroupsecurityrule "github.com/oracle/provider-oci/internal/controller/networking/networksecuritygroupsecurityrule"
	privateip "github.com/oracle/provider-oci/internal/controller/networking/privateip"
	publicip "github.com/oracle/provider-oci/internal/controller/networking/publicip"
	publicippool "github.com/oracle/provider-oci/internal/controller/networking/publicippool"
	publicippoolcapacity "github.com/oracle/provider-oci/internal/controller/networking/publicippoolcapacity"
	remotepeeringconnection "github.com/oracle/provider-oci/internal/controller/networking/remotepeeringconnection"
	routetable "github.com/oracle/provider-oci/internal/controller/networking/routetable"
	routetableattachment "github.com/oracle/provider-oci/internal/controller/networking/routetableattachment"
	securitylist "github.com/oracle/provider-oci/internal/controller/networking/securitylist"
	servicegateway "github.com/oracle/provider-oci/internal/controller/networking/servicegateway"
	subnet "github.com/oracle/provider-oci/internal/controller/networking/subnet"
	vcn "github.com/oracle/provider-oci/internal/controller/networking/vcn"
	vlan "github.com/oracle/provider-oci/internal/controller/networking/vlan"
	vnicattachment "github.com/oracle/provider-oci/internal/controller/networking/vnicattachment"
	backendnetworkloadbalancer "github.com/oracle/provider-oci/internal/controller/networkloadbalancer/backend"
	backendsetnetworkloadbalancer "github.com/oracle/provider-oci/internal/controller/networkloadbalancer/backendset"
	listenernetworkloadbalancer "github.com/oracle/provider-oci/internal/controller/networkloadbalancer/listener"
	networkloadbalancer "github.com/oracle/provider-oci/internal/controller/networkloadbalancer/networkloadbalancer"
	networkloadbalancersbackendsetsunified "github.com/oracle/provider-oci/internal/controller/networkloadbalancer/networkloadbalancersbackendsetsunified"
	configurationnosql "github.com/oracle/provider-oci/internal/controller/nosql/configuration"
	index "github.com/oracle/provider-oci/internal/controller/nosql/index"
	table "github.com/oracle/provider-oci/internal/controller/nosql/table"
	tablereplica "github.com/oracle/provider-oci/internal/controller/nosql/tablereplica"
	bucket "github.com/oracle/provider-oci/internal/controller/objectstorage/bucket"
	namespacemetadata "github.com/oracle/provider-oci/internal/controller/objectstorage/namespacemetadata"
	object "github.com/oracle/provider-oci/internal/controller/objectstorage/object"
	objectlifecyclepolicy "github.com/oracle/provider-oci/internal/controller/objectstorage/objectlifecyclepolicy"
	preauthrequest "github.com/oracle/provider-oci/internal/controller/objectstorage/preauthrequest"
	privateendpointobjectstorage "github.com/oracle/provider-oci/internal/controller/objectstorage/privateendpoint"
	replicationpolicy "github.com/oracle/provider-oci/internal/controller/objectstorage/replicationpolicy"
	oceinstance "github.com/oracle/provider-oci/internal/controller/oce/oceinstance"
	clusterocvp "github.com/oracle/provider-oci/internal/controller/ocvp/cluster"
	datastore "github.com/oracle/provider-oci/internal/controller/ocvp/datastore"
	datastorecluster "github.com/oracle/provider-oci/internal/controller/ocvp/datastorecluster"
	esxihost "github.com/oracle/provider-oci/internal/controller/ocvp/esxihost"
	sddc "github.com/oracle/provider-oci/internal/controller/ocvp/sddc"
	odainstance "github.com/oracle/provider-oci/internal/controller/oda/odainstance"
	odaprivateendpoint "github.com/oracle/provider-oci/internal/controller/oda/odaprivateendpoint"
	odaprivateendpointattachment "github.com/oracle/provider-oci/internal/controller/oda/odaprivateendpointattachment"
	odaprivateendpointscanproxy "github.com/oracle/provider-oci/internal/controller/oda/odaprivateendpointscanproxy"
	notificationtopic "github.com/oracle/provider-oci/internal/controller/ons/notificationtopic"
	subscription "github.com/oracle/provider-oci/internal/controller/ons/subscription"
	opainstance "github.com/oracle/provider-oci/internal/controller/opa/opainstance"
	opensearchcluster "github.com/oracle/provider-oci/internal/controller/opensearch/opensearchcluster"
	opensearchclusterpipeline "github.com/oracle/provider-oci/internal/controller/opensearch/opensearchclusterpipeline"
	operatorcontrol "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/operatorcontrol"
	operatorcontrolassignment "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/operatorcontrolassignment"
	awrhub "github.com/oracle/provider-oci/internal/controller/opsi/awrhub"
	awrhubsource "github.com/oracle/provider-oci/internal/controller/opsi/awrhubsource"
	awrhubsourceawrhubsourcesmanagement "github.com/oracle/provider-oci/internal/controller/opsi/awrhubsourceawrhubsourcesmanagement"
	databaseinsight "github.com/oracle/provider-oci/internal/controller/opsi/databaseinsight"
	enterprisemanagerbridge "github.com/oracle/provider-oci/internal/controller/opsi/enterprisemanagerbridge"
	exadatainsight "github.com/oracle/provider-oci/internal/controller/opsi/exadatainsight"
	hostinsight "github.com/oracle/provider-oci/internal/controller/opsi/hostinsight"
	newsreport "github.com/oracle/provider-oci/internal/controller/opsi/newsreport"
	operationsinsightsprivateendpoint "github.com/oracle/provider-oci/internal/controller/opsi/operationsinsightsprivateendpoint"
	operationsinsightswarehouse "github.com/oracle/provider-oci/internal/controller/opsi/operationsinsightswarehouse"
	operationsinsightswarehousedownloadwarehousewallet "github.com/oracle/provider-oci/internal/controller/opsi/operationsinsightswarehousedownloadwarehousewallet"
	operationsinsightswarehouserotatewarehousewallet "github.com/oracle/provider-oci/internal/controller/opsi/operationsinsightswarehouserotatewarehousewallet"
	operationsinsightswarehouseuser "github.com/oracle/provider-oci/internal/controller/opsi/operationsinsightswarehouseuser"
	opsiconfiguration "github.com/oracle/provider-oci/internal/controller/opsi/opsiconfiguration"
	enrollmentstatus "github.com/oracle/provider-oci/internal/controller/optimizer/enrollmentstatus"
	profile "github.com/oracle/provider-oci/internal/controller/optimizer/profile"
	recommendation "github.com/oracle/provider-oci/internal/controller/optimizer/recommendation"
	resourceaction "github.com/oracle/provider-oci/internal/controller/optimizer/resourceaction"
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
	profileosmanagementhub "github.com/oracle/provider-oci/internal/controller/osmanagementhub/profile"
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
	addressactionverification "github.com/oracle/provider-oci/internal/controller/ospgateway/addressactionverification"
	subscriptionospgateway "github.com/oracle/provider-oci/internal/controller/ospgateway/subscription"
	providerconfig "github.com/oracle/provider-oci/internal/controller/providerconfig"
	privateserviceaccess "github.com/oracle/provider-oci/internal/controller/psa/privateserviceaccess"
	psqlbackup "github.com/oracle/provider-oci/internal/controller/psql/psqlbackup"
	psqlconfiguration "github.com/oracle/provider-oci/internal/controller/psql/psqlconfiguration"
	psqldbsystem "github.com/oracle/provider-oci/internal/controller/psql/psqldbsystem"
	queue "github.com/oracle/provider-oci/internal/controller/queue/queue"
	protecteddatabase "github.com/oracle/provider-oci/internal/controller/recovery/protecteddatabase"
	protectionpolicy "github.com/oracle/provider-oci/internal/controller/recovery/protectionpolicy"
	recoveryservicesubnet "github.com/oracle/provider-oci/internal/controller/recovery/recoveryservicesubnet"
	ocicacheconfigset "github.com/oracle/provider-oci/internal/controller/redis/ocicacheconfigset"
	ocicacheconfigsetlistassociatedocicachecluster "github.com/oracle/provider-oci/internal/controller/redis/ocicacheconfigsetlistassociatedocicachecluster"
	ocicacheuser "github.com/oracle/provider-oci/internal/controller/redis/ocicacheuser"
	ocicacheusergetrediscluster "github.com/oracle/provider-oci/internal/controller/redis/ocicacheusergetrediscluster"
	rediscluster "github.com/oracle/provider-oci/internal/controller/redis/rediscluster"
	redisclusterattachocicacheuser "github.com/oracle/provider-oci/internal/controller/redis/redisclusterattachocicacheuser"
	redisclustercreateidentitytoken "github.com/oracle/provider-oci/internal/controller/redis/redisclustercreateidentitytoken"
	redisclusterdetachocicacheuser "github.com/oracle/provider-oci/internal/controller/redis/redisclusterdetachocicacheuser"
	redisclustergetocicacheuser "github.com/oracle/provider-oci/internal/controller/redis/redisclustergetocicacheuser"
	monitoredregion "github.com/oracle/provider-oci/internal/controller/resourceanalytics/monitoredregion"
	resourceanalyticsinstance "github.com/oracle/provider-oci/internal/controller/resourceanalytics/resourceanalyticsinstance"
	resourceanalyticsinstanceoacmanagement "github.com/oracle/provider-oci/internal/controller/resourceanalytics/resourceanalyticsinstanceoacmanagement"
	tenancyattachment "github.com/oracle/provider-oci/internal/controller/resourceanalytics/tenancyattachment"
	privateendpointresourcemanager "github.com/oracle/provider-oci/internal/controller/resourcemanager/privateendpoint"
	scheduleresourcescheduler "github.com/oracle/provider-oci/internal/controller/resourcescheduler/schedule"
	serviceconnector "github.com/oracle/provider-oci/internal/controller/sch/serviceconnector"
	securityattribute "github.com/oracle/provider-oci/internal/controller/securityattribute/securityattribute"
	securityattributenamespace "github.com/oracle/provider-oci/internal/controller/securityattribute/securityattributenamespace"
	privateapplication "github.com/oracle/provider-oci/internal/controller/servicecatalog/privateapplication"
	servicecatalog "github.com/oracle/provider-oci/internal/controller/servicecatalog/servicecatalog"
	servicecatalogassociation "github.com/oracle/provider-oci/internal/controller/servicecatalog/servicecatalogassociation"
	baselineablemetric "github.com/oracle/provider-oci/internal/controller/stackmonitoring/baselineablemetric"
	configstackmonitoring "github.com/oracle/provider-oci/internal/controller/stackmonitoring/config"
	discoveryjob "github.com/oracle/provider-oci/internal/controller/stackmonitoring/discoveryjob"
	maintenancewindowstackmonitoring "github.com/oracle/provider-oci/internal/controller/stackmonitoring/maintenancewindow"
	maintenancewindowsretryfailedoperation "github.com/oracle/provider-oci/internal/controller/stackmonitoring/maintenancewindowsretryfailedoperation"
	maintenancewindowsstop "github.com/oracle/provider-oci/internal/controller/stackmonitoring/maintenancewindowsstop"
	metricextension "github.com/oracle/provider-oci/internal/controller/stackmonitoring/metricextension"
	metricextensionmetricextensionongivenresourcesmanagement "github.com/oracle/provider-oci/internal/controller/stackmonitoring/metricextensionmetricextensionongivenresourcesmanagement"
	monitoredresource "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoredresource"
	monitoredresourcesassociatemonitoredresource "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoredresourcesassociatemonitoredresource"
	monitoredresourceslistmember "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoredresourceslistmember"
	monitoredresourcessearch "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoredresourcessearch"
	monitoredresourcessearchassociation "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoredresourcessearchassociation"
	monitoredresourcetask "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoredresourcetask"
	monitoredresourcetype "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoredresourcetype"
	monitoringtemplate "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringtemplate"
	monitoringtemplatealarmcondition "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringtemplatealarmcondition"
	monitoringtemplatemonitoringtemplateongivenresourcesmanagement "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringtemplatemonitoringtemplateongivenresourcesmanagement"
	processset "github.com/oracle/provider-oci/internal/controller/stackmonitoring/processset"
	connectharness "github.com/oracle/provider-oci/internal/controller/streaming/connectharness"
	stream "github.com/oracle/provider-oci/internal/controller/streaming/stream"
	streampool "github.com/oracle/provider-oci/internal/controller/streaming/streampool"
	subscriptionmapping "github.com/oracle/provider-oci/internal/controller/tenantmanagercontrolplane/subscriptionmapping"
	subscriptionredeemableuser "github.com/oracle/provider-oci/internal/controller/usageproxy/subscriptionredeemableuser"
	secret "github.com/oracle/provider-oci/internal/controller/vault/secret"
	vbsinstance "github.com/oracle/provider-oci/internal/controller/vbsinst/vbsinstance"
	vbinstance "github.com/oracle/provider-oci/internal/controller/visualbuilder/vbinstance"
	pathanalysi "github.com/oracle/provider-oci/internal/controller/vnmonitoring/pathanalysi"
	containerscanrecipe "github.com/oracle/provider-oci/internal/controller/vulnerabilityscanning/containerscanrecipe"
	containerscantarget "github.com/oracle/provider-oci/internal/controller/vulnerabilityscanning/containerscantarget"
	hostscanrecipe "github.com/oracle/provider-oci/internal/controller/vulnerabilityscanning/hostscanrecipe"
	hostscantarget "github.com/oracle/provider-oci/internal/controller/vulnerabilityscanning/hostscantarget"
	webappacceleration "github.com/oracle/provider-oci/internal/controller/waa/webappacceleration"
	webappaccelerationpolicy "github.com/oracle/provider-oci/internal/controller/waa/webappaccelerationpolicy"
	addresslist "github.com/oracle/provider-oci/internal/controller/waas/addresslist"
	certificatewaas "github.com/oracle/provider-oci/internal/controller/waas/certificate"
	customprotectionrule "github.com/oracle/provider-oci/internal/controller/waas/customprotectionrule"
	httpredirect "github.com/oracle/provider-oci/internal/controller/waas/httpredirect"
	protectionrule "github.com/oracle/provider-oci/internal/controller/waas/protectionrule"
	purgecache "github.com/oracle/provider-oci/internal/controller/waas/purgecache"
	waaspolicy "github.com/oracle/provider-oci/internal/controller/waas/waaspolicy"
	networkaddresslist "github.com/oracle/provider-oci/internal/controller/waf/networkaddresslist"
	webappfirewall "github.com/oracle/provider-oci/internal/controller/waf/webappfirewall"
	webappfirewallpolicy "github.com/oracle/provider-oci/internal/controller/waf/webappfirewallpolicy"
	configurationzpr "github.com/oracle/provider-oci/internal/controller/zpr/configuration"
	zprpolicy "github.com/oracle/provider-oci/internal/controller/zpr/zprpolicy"
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
		autoscalingconfigurationbds.Setup,
		bdscapacityreport.Setup,
		bdsinstance.Setup,
		bdsinstanceapikey.Setup,
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
		clusterstartcredentialrotationmanagement.Setup,
		clusterworkloadmapping.Setup,
		nodepool.Setup,
		virtualnodepool.Setup,
		containerinstance.Setup,
		byoasn.Setup,
		listingresourceversionagreement.Setup,
		virtualnetwork.Setup,
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
		migrationdatabase.Setup,
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
		catalog.Setup,
		catalogprivateendpoint.Setup,
		connection.Setup,
		dataasset.Setup,
		metastore.Setup,
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
		safeaddsdmcolumns.Setup,
		safealert.Setup,
		safealertpolicy.Setup,
		safealertpolicyrule.Setup,
		safeattributeset.Setup,
		safeauditarchiveretrieval.Setup,
		safeauditpolicy.Setup,
		safeauditpolicymanagement.Setup,
		safeauditprofile.Setup,
		safeauditprofilemanagement.Setup,
		safeaudittrail.Setup,
		safeaudittrailmanagement.Setup,
		safecalculateauditvolumeavailable.Setup,
		safecalculateauditvolumecollected.Setup,
		safecomparesecurityassessment.Setup,
		safecompareuserassessment.Setup,
		safedatabasesecurityconfig.Setup,
		safedatabasesecurityconfigmanagement.Setup,
		safedatasafeconfiguration.Setup,
		safedatasafeprivateendpoint.Setup,
		safediscoveryjob.Setup,
		safediscoveryjobsresult.Setup,
		safegenerateonpremconnectorconfiguration.Setup,
		safelibrarymaskingformat.Setup,
		safemaskdata.Setup,
		safemaskingpoliciesapplydifferencetomaskingcolumns.Setup,
		safemaskingpoliciesmaskingcolumn.Setup,
		safemaskingpolicy.Setup,
		safemaskingpolicyhealthreportmanagement.Setup,
		safemaskingreportmanagement.Setup,
		safeonpremconnector.Setup,
		safereport.Setup,
		safereportdefinition.Setup,
		safesdmmaskingpolicydifference.Setup,
		safesecurityassessment.Setup,
		safesecurityassessmentcheck.Setup,
		safesecurityassessmentfinding.Setup,
		safesecuritypolicy.Setup,
		safesecuritypolicyconfig.Setup,
		safesecuritypolicydeployment.Setup,
		safesecuritypolicydeploymentmanagement.Setup,
		safesecuritypolicymanagement.Setup,
		safesensitivedatamodel.Setup,
		safesensitivedatamodelreferentialrelation.Setup,
		safesensitivedatamodelsapplydiscoveryjobresults.Setup,
		safesensitivedatamodelssensitivecolumn.Setup,
		safesensitivetype.Setup,
		safesensitivetypegroup.Setup,
		safesensitivetypegroupgroupedsensitivetype.Setup,
		safesensitivetypesexport.Setup,
		safesetsecurityassessmentbaseline.Setup,
		safesetsecurityassessmentbaselinemanagement.Setup,
		safesetuserassessmentbaseline.Setup,
		safesetuserassessmentbaselinemanagement.Setup,
		safesqlcollection.Setup,
		safesqlfirewallpolicy.Setup,
		safesqlfirewallpolicymanagement.Setup,
		safetargetalertpolicyassociation.Setup,
		safetargetdatabase.Setup,
		safetargetdatabasegroup.Setup,
		safetargetdatabasepeertargetdatabase.Setup,
		safeunifiedauditpolicy.Setup,
		safeunifiedauditpolicydefinition.Setup,
		safeunsetsecurityassessmentbaseline.Setup,
		safeunsetsecurityassessmentbaselinemanagement.Setup,
		safeunsetuserassessmentbaseline.Setup,
		safeunsetuserassessmentbaselinemanagement.Setup,
		safeuserassessment.Setup,
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
		delegationcontrol.Setup,
		delegationsubscription.Setup,
		occdemandsignal.Setup,
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
		applicationfunctions.Setup,
		function.Setup,
		invokefunction.Setup,
		fusionenvironment.Setup,
		fusionenvironmentadminuser.Setup,
		fusionenvironmentdatamaskingactivity.Setup,
		fusionenvironmentfamily.Setup,
		fusionenvironmentrefreshactivity.Setup,
		fusionenvironmentserviceattachment.Setup,
		agentagent.Setup,
		agentagentendpoint.Setup,
		agentdataingestionjob.Setup,
		agentdatasource.Setup,
		agentknowledgebase.Setup,
		agenttool.Setup,
		dedicatedaicluster.Setup,
		endpointgenerativeai.Setup,
		generativeaiprivateendpoint.Setup,
		modelgenerativeai.Setup,
		artifactbypath.Setup,
		privateendpointgloballydistributeddatabase.Setup,
		shardeddatabase.Setup,
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
		identitypropagationtrust.Setup,
		identityprovideridentitydomains.Setup,
		identitysetting.Setup,
		kmsisetting.Setup,
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
		loganalyticsentitytype.Setup,
		loganalyticsimportcustomcontent.Setup,
		loganalyticsloggroup.Setup,
		loganalyticsobjectcollectionrule.Setup,
		loganalyticspreferencesmanagement.Setup,
		loganalyticsresourcecategoriesmanagement.Setup,
		loganalyticsunprocesseddatabucketmanagement.Setup,
		namespace.Setup,
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
		kafkaclusterconfig.Setup,
		kafkaclustersuperusersmanagement.Setup,
		managementagent.Setup,
		managementagentdatasource.Setup,
		managementagentinstallkey.Setup,
		namedcredential.Setup,
		managementdashboardsimport.Setup,
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
		namespacemetadata.Setup,
		object.Setup,
		objectlifecyclepolicy.Setup,
		preauthrequest.Setup,
		privateendpointobjectstorage.Setup,
		replicationpolicy.Setup,
		oceinstance.Setup,
		clusterocvp.Setup,
		datastore.Setup,
		datastorecluster.Setup,
		esxihost.Setup,
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
		profileosmanagementhub.Setup,
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
		addressactionverification.Setup,
		subscriptionospgateway.Setup,
		providerconfig.Setup,
		privateserviceaccess.Setup,
		psqlbackup.Setup,
		psqlconfiguration.Setup,
		psqldbsystem.Setup,
		queue.Setup,
		protecteddatabase.Setup,
		protectionpolicy.Setup,
		recoveryservicesubnet.Setup,
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
		privateapplication.Setup,
		servicecatalog.Setup,
		servicecatalogassociation.Setup,
		baselineablemetric.Setup,
		configstackmonitoring.Setup,
		discoveryjob.Setup,
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
		monitoringtemplatemonitoringtemplateongivenresourcesmanagement.Setup,
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

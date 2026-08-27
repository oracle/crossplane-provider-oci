/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	knowledgebase "github.com/oracle/provider-oci/internal/controller/namespaced/adm/knowledgebase"
	remediationrecipe "github.com/oracle/provider-oci/internal/controller/namespaced/adm/remediationrecipe"
	remediationrun "github.com/oracle/provider-oci/internal/controller/namespaced/adm/remediationrun"
	vulnerabilityaudit "github.com/oracle/provider-oci/internal/controller/namespaced/adm/vulnerabilityaudit"
	aidataplatform "github.com/oracle/provider-oci/internal/controller/namespaced/aidataplatform/aidataplatform"
	model "github.com/oracle/provider-oci/internal/controller/namespaced/aidocument/model"
	processorjob "github.com/oracle/provider-oci/internal/controller/namespaced/aidocument/processorjob"
	project "github.com/oracle/provider-oci/internal/controller/namespaced/aidocument/project"
	endpoint "github.com/oracle/provider-oci/internal/controller/namespaced/ailanguage/endpoint"
	job "github.com/oracle/provider-oci/internal/controller/namespaced/ailanguage/job"
	modelailanguage "github.com/oracle/provider-oci/internal/controller/namespaced/ailanguage/model"
	projectailanguage "github.com/oracle/provider-oci/internal/controller/namespaced/ailanguage/project"
	modelaivision "github.com/oracle/provider-oci/internal/controller/namespaced/aivision/model"
	projectaivision "github.com/oracle/provider-oci/internal/controller/namespaced/aivision/project"
	streamgroup "github.com/oracle/provider-oci/internal/controller/namespaced/aivision/streamgroup"
	streamjob "github.com/oracle/provider-oci/internal/controller/namespaced/aivision/streamjob"
	streamsource "github.com/oracle/provider-oci/internal/controller/namespaced/aivision/streamsource"
	visionprivateendpoint "github.com/oracle/provider-oci/internal/controller/namespaced/aivision/visionprivateendpoint"
	analyticsinstance "github.com/oracle/provider-oci/internal/controller/namespaced/analytics/analyticsinstance"
	analyticsinstanceprivateaccesschannel "github.com/oracle/provider-oci/internal/controller/namespaced/analytics/analyticsinstanceprivateaccesschannel"
	analyticsinstanceresourcegroup "github.com/oracle/provider-oci/internal/controller/namespaced/analytics/analyticsinstanceresourcegroup"
	analyticsinstancevanityurl "github.com/oracle/provider-oci/internal/controller/namespaced/analytics/analyticsinstancevanityurl"
	announcementsubscription "github.com/oracle/provider-oci/internal/controller/namespaced/announcementsservice/announcementsubscription"
	announcementsubscriptionsactionschangecompartment "github.com/oracle/provider-oci/internal/controller/namespaced/announcementsservice/announcementsubscriptionsactionschangecompartment"
	announcementsubscriptionsfiltergroup "github.com/oracle/provider-oci/internal/controller/namespaced/announcementsservice/announcementsubscriptionsfiltergroup"
	privilegedapicontrol "github.com/oracle/provider-oci/internal/controller/namespaced/apiaccesscontrol/privilegedapicontrol"
	privilegedapirequest "github.com/oracle/provider-oci/internal/controller/namespaced/apiaccesscontrol/privilegedapirequest"
	api "github.com/oracle/provider-oci/internal/controller/namespaced/apigateway/api"
	certificate "github.com/oracle/provider-oci/internal/controller/namespaced/apigateway/certificate"
	deployment "github.com/oracle/provider-oci/internal/controller/namespaced/apigateway/deployment"
	gateway "github.com/oracle/provider-oci/internal/controller/namespaced/apigateway/gateway"
	subscriber "github.com/oracle/provider-oci/internal/controller/namespaced/apigateway/subscriber"
	usageplan "github.com/oracle/provider-oci/internal/controller/namespaced/apigateway/usageplan"
	apiplatforminstance "github.com/oracle/provider-oci/internal/controller/namespaced/apiplatform/apiplatforminstance"
	apmdomain "github.com/oracle/provider-oci/internal/controller/namespaced/apm/apmdomain"
	config "github.com/oracle/provider-oci/internal/controller/namespaced/apmconfig/config"
	datafile "github.com/oracle/provider-oci/internal/controller/namespaced/apmconfig/datafile"
	dedicatedvantagepoint "github.com/oracle/provider-oci/internal/controller/namespaced/apmsynthetics/dedicatedvantagepoint"
	monitor "github.com/oracle/provider-oci/internal/controller/namespaced/apmsynthetics/monitor"
	onpremisevantagepoint "github.com/oracle/provider-oci/internal/controller/namespaced/apmsynthetics/onpremisevantagepoint"
	onpremisevantagepointworker "github.com/oracle/provider-oci/internal/controller/namespaced/apmsynthetics/onpremisevantagepointworker"
	script "github.com/oracle/provider-oci/internal/controller/namespaced/apmsynthetics/script"
	scheduledquery "github.com/oracle/provider-oci/internal/controller/namespaced/apmtraces/scheduledquery"
	controlmonitorpluginmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/appmgmt/controlmonitorpluginmanagement"
	containerconfiguration "github.com/oracle/provider-oci/internal/controller/namespaced/artifacts/containerconfiguration"
	containerimagesignature "github.com/oracle/provider-oci/internal/controller/namespaced/artifacts/containerimagesignature"
	containerrepository "github.com/oracle/provider-oci/internal/controller/namespaced/artifacts/containerrepository"
	genericartifact "github.com/oracle/provider-oci/internal/controller/namespaced/artifacts/genericartifact"
	repository "github.com/oracle/provider-oci/internal/controller/namespaced/artifacts/repository"
	configuration "github.com/oracle/provider-oci/internal/controller/namespaced/audit/configuration"
	autoscalingconfiguration "github.com/oracle/provider-oci/internal/controller/namespaced/autoscaling/autoscalingconfiguration"
	bastion "github.com/oracle/provider-oci/internal/controller/namespaced/bastion/bastion"
	session "github.com/oracle/provider-oci/internal/controller/namespaced/bastion/session"
	batchcontext "github.com/oracle/provider-oci/internal/controller/namespaced/batch/batchcontext"
	batchjobpool "github.com/oracle/provider-oci/internal/controller/namespaced/batch/batchjobpool"
	batchtaskenvironment "github.com/oracle/provider-oci/internal/controller/namespaced/batch/batchtaskenvironment"
	batchtaskprofile "github.com/oracle/provider-oci/internal/controller/namespaced/batch/batchtaskprofile"
	autoscalingconfigurationbds "github.com/oracle/provider-oci/internal/controller/namespaced/bds/autoscalingconfiguration"
	bdscapacityreport "github.com/oracle/provider-oci/internal/controller/namespaced/bds/bdscapacityreport"
	bdscapacityreservation "github.com/oracle/provider-oci/internal/controller/namespaced/bds/bdscapacityreservation"
	bdsclusteradminpasswordresetaction "github.com/oracle/provider-oci/internal/controller/namespaced/bds/bdsclusteradminpasswordresetaction"
	bdsinstance "github.com/oracle/provider-oci/internal/controller/namespaced/bds/bdsinstance"
	bdsinstanceapikey "github.com/oracle/provider-oci/internal/controller/namespaced/bds/bdsinstanceapikey"
	bdsinstancebdscapacityreservationconfiguration "github.com/oracle/provider-oci/internal/controller/namespaced/bds/bdsinstancebdscapacityreservationconfiguration"
	bdsinstancebdscertificateconfiguration "github.com/oracle/provider-oci/internal/controller/namespaced/bds/bdsinstancebdscertificateconfiguration"
	bdsinstanceexecutebootstrapscriptaction "github.com/oracle/provider-oci/internal/controller/namespaced/bds/bdsinstanceexecutebootstrapscriptaction"
	bdsinstanceidentityconfiguration "github.com/oracle/provider-oci/internal/controller/namespaced/bds/bdsinstanceidentityconfiguration"
	bdsinstancemetastoreconfig "github.com/oracle/provider-oci/internal/controller/namespaced/bds/bdsinstancemetastoreconfig"
	bdsinstancenodebackup "github.com/oracle/provider-oci/internal/controller/namespaced/bds/bdsinstancenodebackup"
	bdsinstancenodebackupconfiguration "github.com/oracle/provider-oci/internal/controller/namespaced/bds/bdsinstancenodebackupconfiguration"
	bdsinstancenodereplaceconfiguration "github.com/oracle/provider-oci/internal/controller/namespaced/bds/bdsinstancenodereplaceconfiguration"
	bdsinstanceoperationcertificatemanagementsmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/bds/bdsinstanceoperationcertificatemanagementsmanagement"
	bdsinstanceospatchaction "github.com/oracle/provider-oci/internal/controller/namespaced/bds/bdsinstanceospatchaction"
	bdsinstancepatchaction "github.com/oracle/provider-oci/internal/controller/namespaced/bds/bdsinstancepatchaction"
	bdsinstancereplacenodeaction "github.com/oracle/provider-oci/internal/controller/namespaced/bds/bdsinstancereplacenodeaction"
	bdsinstanceresourceprincipalconfiguration "github.com/oracle/provider-oci/internal/controller/namespaced/bds/bdsinstanceresourceprincipalconfiguration"
	bdsinstancesoftwareupdateaction "github.com/oracle/provider-oci/internal/controller/namespaced/bds/bdsinstancesoftwareupdateaction"
	blockchainplatform "github.com/oracle/provider-oci/internal/controller/namespaced/blockchain/blockchainplatform"
	osn "github.com/oracle/provider-oci/internal/controller/namespaced/blockchain/osn"
	peer "github.com/oracle/provider-oci/internal/controller/namespaced/blockchain/peer"
	bootvolume "github.com/oracle/provider-oci/internal/controller/namespaced/blockstorage/bootvolume"
	bootvolumebackup "github.com/oracle/provider-oci/internal/controller/namespaced/blockstorage/bootvolumebackup"
	volume "github.com/oracle/provider-oci/internal/controller/namespaced/blockstorage/volume"
	volumeattachment "github.com/oracle/provider-oci/internal/controller/namespaced/blockstorage/volumeattachment"
	volumebackup "github.com/oracle/provider-oci/internal/controller/namespaced/blockstorage/volumebackup"
	volumebackuppolicy "github.com/oracle/provider-oci/internal/controller/namespaced/blockstorage/volumebackuppolicy"
	volumebackuppolicyassignment "github.com/oracle/provider-oci/internal/controller/namespaced/blockstorage/volumebackuppolicyassignment"
	volumegroup "github.com/oracle/provider-oci/internal/controller/namespaced/blockstorage/volumegroup"
	volumegroupbackup "github.com/oracle/provider-oci/internal/controller/namespaced/blockstorage/volumegroupbackup"
	alertrule "github.com/oracle/provider-oci/internal/controller/namespaced/budget/alertrule"
	budget "github.com/oracle/provider-oci/internal/controller/namespaced/budget/budget"
	costalertsubscription "github.com/oracle/provider-oci/internal/controller/namespaced/budget/costalertsubscription"
	costanomalyevent "github.com/oracle/provider-oci/internal/controller/namespaced/budget/costanomalyevent"
	costanomalymonitor "github.com/oracle/provider-oci/internal/controller/namespaced/budget/costanomalymonitor"
	costanomalymonitorcostanomalymonitorenabletogglesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/budget/costanomalymonitorcostanomalymonitorenabletogglesmanagement"
	internaloccmdemandsignal "github.com/oracle/provider-oci/internal/controller/namespaced/capacitymanagement/internaloccmdemandsignal"
	internaloccmdemandsignaldelivery "github.com/oracle/provider-oci/internal/controller/namespaced/capacitymanagement/internaloccmdemandsignaldelivery"
	occavailabilitycatalog "github.com/oracle/provider-oci/internal/controller/namespaced/capacitymanagement/occavailabilitycatalog"
	occcapacityrequest "github.com/oracle/provider-oci/internal/controller/namespaced/capacitymanagement/occcapacityrequest"
	occcustomergroup "github.com/oracle/provider-oci/internal/controller/namespaced/capacitymanagement/occcustomergroup"
	occcustomergroupocccustomer "github.com/oracle/provider-oci/internal/controller/namespaced/capacitymanagement/occcustomergroupocccustomer"
	occmdemandsignal "github.com/oracle/provider-oci/internal/controller/namespaced/capacitymanagement/occmdemandsignal"
	occmdemandsignalitem "github.com/oracle/provider-oci/internal/controller/namespaced/capacitymanagement/occmdemandsignalitem"
	cabundle "github.com/oracle/provider-oci/internal/controller/namespaced/certificatesmanagement/cabundle"
	certificatecertificatesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/certificatesmanagement/certificate"
	certificateauthority "github.com/oracle/provider-oci/internal/controller/namespaced/certificatesmanagement/certificateauthority"
	agent "github.com/oracle/provider-oci/internal/controller/namespaced/cloudbridge/agent"
	agentdependency "github.com/oracle/provider-oci/internal/controller/namespaced/cloudbridge/agentdependency"
	agentplugin "github.com/oracle/provider-oci/internal/controller/namespaced/cloudbridge/agentplugin"
	asset "github.com/oracle/provider-oci/internal/controller/namespaced/cloudbridge/asset"
	assetsource "github.com/oracle/provider-oci/internal/controller/namespaced/cloudbridge/assetsource"
	discoveryschedule "github.com/oracle/provider-oci/internal/controller/namespaced/cloudbridge/discoveryschedule"
	environment "github.com/oracle/provider-oci/internal/controller/namespaced/cloudbridge/environment"
	inventory "github.com/oracle/provider-oci/internal/controller/namespaced/cloudbridge/inventory"
	adhocquery "github.com/oracle/provider-oci/internal/controller/namespaced/cloudguard/adhocquery"
	cloudguardconfiguration "github.com/oracle/provider-oci/internal/controller/namespaced/cloudguard/cloudguardconfiguration"
	datamaskrule "github.com/oracle/provider-oci/internal/controller/namespaced/cloudguard/datamaskrule"
	datasource "github.com/oracle/provider-oci/internal/controller/namespaced/cloudguard/datasource"
	detectorrecipe "github.com/oracle/provider-oci/internal/controller/namespaced/cloudguard/detectorrecipe"
	managedlist "github.com/oracle/provider-oci/internal/controller/namespaced/cloudguard/managedlist"
	responderrecipe "github.com/oracle/provider-oci/internal/controller/namespaced/cloudguard/responderrecipe"
	savedquery "github.com/oracle/provider-oci/internal/controller/namespaced/cloudguard/savedquery"
	securityrecipe "github.com/oracle/provider-oci/internal/controller/namespaced/cloudguard/securityrecipe"
	securityzone "github.com/oracle/provider-oci/internal/controller/namespaced/cloudguard/securityzone"
	target "github.com/oracle/provider-oci/internal/controller/namespaced/cloudguard/target"
	wlpagent "github.com/oracle/provider-oci/internal/controller/namespaced/cloudguard/wlpagent"
	migration "github.com/oracle/provider-oci/internal/controller/namespaced/cloudmigrations/migration"
	migrationasset "github.com/oracle/provider-oci/internal/controller/namespaced/cloudmigrations/migrationasset"
	migrationplan "github.com/oracle/provider-oci/internal/controller/namespaced/cloudmigrations/migrationplan"
	replicationschedule "github.com/oracle/provider-oci/internal/controller/namespaced/cloudmigrations/replicationschedule"
	targetasset "github.com/oracle/provider-oci/internal/controller/namespaced/cloudmigrations/targetasset"
	clusterplacementgroup "github.com/oracle/provider-oci/internal/controller/namespaced/clusterplacementgroups/clusterplacementgroup"
	appcataloglistingresourceversionagreement "github.com/oracle/provider-oci/internal/controller/namespaced/compute/appcataloglistingresourceversionagreement"
	appcatalogsubscription "github.com/oracle/provider-oci/internal/controller/namespaced/compute/appcatalogsubscription"
	clusternetwork "github.com/oracle/provider-oci/internal/controller/namespaced/compute/clusternetwork"
	computecapacityreport "github.com/oracle/provider-oci/internal/controller/namespaced/compute/computecapacityreport"
	computecapacityreservation "github.com/oracle/provider-oci/internal/controller/namespaced/compute/computecapacityreservation"
	computecapacitytopology "github.com/oracle/provider-oci/internal/controller/namespaced/compute/computecapacitytopology"
	computecluster "github.com/oracle/provider-oci/internal/controller/namespaced/compute/computecluster"
	computegpumemorycluster "github.com/oracle/provider-oci/internal/controller/namespaced/compute/computegpumemorycluster"
	computegpumemoryfabric "github.com/oracle/provider-oci/internal/controller/namespaced/compute/computegpumemoryfabric"
	computehost "github.com/oracle/provider-oci/internal/controller/namespaced/compute/computehost"
	computehostgroup "github.com/oracle/provider-oci/internal/controller/namespaced/compute/computehostgroup"
	computeimagecapabilityschema "github.com/oracle/provider-oci/internal/controller/namespaced/compute/computeimagecapabilityschema"
	consolehistory "github.com/oracle/provider-oci/internal/controller/namespaced/compute/consolehistory"
	dedicatedvmhost "github.com/oracle/provider-oci/internal/controller/namespaced/compute/dedicatedvmhost"
	image "github.com/oracle/provider-oci/internal/controller/namespaced/compute/image"
	instance "github.com/oracle/provider-oci/internal/controller/namespaced/compute/instance"
	instanceconfiguration "github.com/oracle/provider-oci/internal/controller/namespaced/compute/instanceconfiguration"
	instanceconsoleconnection "github.com/oracle/provider-oci/internal/controller/namespaced/compute/instanceconsoleconnection"
	instancemaintenanceevent "github.com/oracle/provider-oci/internal/controller/namespaced/compute/instancemaintenanceevent"
	instancepool "github.com/oracle/provider-oci/internal/controller/namespaced/compute/instancepool"
	instancepoolinstance "github.com/oracle/provider-oci/internal/controller/namespaced/compute/instancepoolinstance"
	shapemanagement "github.com/oracle/provider-oci/internal/controller/namespaced/compute/shapemanagement"
	cccinfrastructure "github.com/oracle/provider-oci/internal/controller/namespaced/computecloudatcustomer/cccinfrastructure"
	cccupgradeschedule "github.com/oracle/provider-oci/internal/controller/namespaced/computecloudatcustomer/cccupgradeschedule"
	addon "github.com/oracle/provider-oci/internal/controller/namespaced/containerengine/addon"
	cluster "github.com/oracle/provider-oci/internal/controller/namespaced/containerengine/cluster"
	clustercompletecredentialrotationmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/containerengine/clustercompletecredentialrotationmanagement"
	clusterpublicapiendpointdecommissionmanager "github.com/oracle/provider-oci/internal/controller/namespaced/containerengine/clusterpublicapiendpointdecommissionmanager"
	clusterstartcredentialrotationmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/containerengine/clusterstartcredentialrotationmanagement"
	clusterworkloadmapping "github.com/oracle/provider-oci/internal/controller/namespaced/containerengine/clusterworkloadmapping"
	nodepool "github.com/oracle/provider-oci/internal/controller/namespaced/containerengine/nodepool"
	virtualnodepool "github.com/oracle/provider-oci/internal/controller/namespaced/containerengine/virtualnodepool"
	containerinstance "github.com/oracle/provider-oci/internal/controller/namespaced/containerinstances/containerinstance"
	byoasn "github.com/oracle/provider-oci/internal/controller/namespaced/core/byoasn"
	listingresourceversionagreement "github.com/oracle/provider-oci/internal/controller/namespaced/core/listingresourceversionagreement"
	virtualnetwork "github.com/oracle/provider-oci/internal/controller/namespaced/core/virtualnetwork"
	costalertsubscriptioncostad "github.com/oracle/provider-oci/internal/controller/namespaced/costad/costalertsubscription"
	costanomalyeventcostad "github.com/oracle/provider-oci/internal/controller/namespaced/costad/costanomalyevent"
	costanomalymonitorcostad "github.com/oracle/provider-oci/internal/controller/namespaced/costad/costanomalymonitor"
	costanomalymonitorcostanomalymonitorenabletogglesmanagementcostad "github.com/oracle/provider-oci/internal/controller/namespaced/costad/costanomalymonitorcostanomalymonitorenabletogglesmanagement"
	advancedclusterfilesystem "github.com/oracle/provider-oci/internal/controller/namespaced/database/advancedclusterfilesystem"
	advancedclusterfilesystemmount "github.com/oracle/provider-oci/internal/controller/namespaced/database/advancedclusterfilesystemmount"
	advancedclusterfilesystemunmount "github.com/oracle/provider-oci/internal/controller/namespaced/database/advancedclusterfilesystemunmount"
	applicationvip "github.com/oracle/provider-oci/internal/controller/namespaced/database/applicationvip"
	autonomouscontainerdatabase "github.com/oracle/provider-oci/internal/controller/namespaced/database/autonomouscontainerdatabase"
	autonomouscontainerdatabaseaddstandby "github.com/oracle/provider-oci/internal/controller/namespaced/database/autonomouscontainerdatabaseaddstandby"
	autonomouscontainerdatabasedataguardassociation "github.com/oracle/provider-oci/internal/controller/namespaced/database/autonomouscontainerdatabasedataguardassociation"
	autonomouscontainerdatabasedataguardassociationoperation "github.com/oracle/provider-oci/internal/controller/namespaced/database/autonomouscontainerdatabasedataguardassociationoperation"
	autonomouscontainerdatabasedataguardrolechange "github.com/oracle/provider-oci/internal/controller/namespaced/database/autonomouscontainerdatabasedataguardrolechange"
	autonomouscontainerdatabasesnapshotstandby "github.com/oracle/provider-oci/internal/controller/namespaced/database/autonomouscontainerdatabasesnapshotstandby"
	autonomousdatabase "github.com/oracle/provider-oci/internal/controller/namespaced/database/autonomousdatabase"
	autonomousdatabasebackup "github.com/oracle/provider-oci/internal/controller/namespaced/database/autonomousdatabasebackup"
	autonomousdatabaseinstancewalletmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/database/autonomousdatabaseinstancewalletmanagement"
	autonomousdatabaseregionalwalletmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/database/autonomousdatabaseregionalwalletmanagement"
	autonomousdatabasesaasadminuser "github.com/oracle/provider-oci/internal/controller/namespaced/database/autonomousdatabasesaasadminuser"
	autonomousdatabasesoftwareimage "github.com/oracle/provider-oci/internal/controller/namespaced/database/autonomousdatabasesoftwareimage"
	autonomousdatabasewallet "github.com/oracle/provider-oci/internal/controller/namespaced/database/autonomousdatabasewallet"
	autonomousexadatainfrastructure "github.com/oracle/provider-oci/internal/controller/namespaced/database/autonomousexadatainfrastructure"
	autonomousvmcluster "github.com/oracle/provider-oci/internal/controller/namespaced/database/autonomousvmcluster"
	autonomousvmclusterordscertificatemanagement "github.com/oracle/provider-oci/internal/controller/namespaced/database/autonomousvmclusterordscertificatemanagement"
	autonomousvmclustersslcertificatemanagement "github.com/oracle/provider-oci/internal/controller/namespaced/database/autonomousvmclustersslcertificatemanagement"
	backup "github.com/oracle/provider-oci/internal/controller/namespaced/database/backup"
	backupcancelmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/database/backupcancelmanagement"
	backupdestination "github.com/oracle/provider-oci/internal/controller/namespaced/database/backupdestination"
	cloudautonomousvmcluster "github.com/oracle/provider-oci/internal/controller/namespaced/database/cloudautonomousvmcluster"
	clouddatabasemanagement "github.com/oracle/provider-oci/internal/controller/namespaced/database/clouddatabasemanagement"
	cloudexadatainfrastructure "github.com/oracle/provider-oci/internal/controller/namespaced/database/cloudexadatainfrastructure"
	cloudexadatainfrastructureconfigureexascalemanagement "github.com/oracle/provider-oci/internal/controller/namespaced/database/cloudexadatainfrastructureconfigureexascalemanagement"
	cloudvmcluster "github.com/oracle/provider-oci/internal/controller/namespaced/database/cloudvmcluster"
	cloudvmclusteriormconfig "github.com/oracle/provider-oci/internal/controller/namespaced/database/cloudvmclusteriormconfig"
	database "github.com/oracle/provider-oci/internal/controller/namespaced/database/database"
	databasesnapshotstandby "github.com/oracle/provider-oci/internal/controller/namespaced/database/databasesnapshotstandby"
	databasesoftwareimage "github.com/oracle/provider-oci/internal/controller/namespaced/database/databasesoftwareimage"
	databasesoftwareschedulemanagement "github.com/oracle/provider-oci/internal/controller/namespaced/database/databasesoftwareschedulemanagement"
	databaseupgrade "github.com/oracle/provider-oci/internal/controller/namespaced/database/databaseupgrade"
	dataguardassociation "github.com/oracle/provider-oci/internal/controller/namespaced/database/dataguardassociation"
	datapatch "github.com/oracle/provider-oci/internal/controller/namespaced/database/datapatch"
	dbhome "github.com/oracle/provider-oci/internal/controller/namespaced/database/dbhome"
	dbnode "github.com/oracle/provider-oci/internal/controller/namespaced/database/dbnode"
	dbnodeconsoleconnection "github.com/oracle/provider-oci/internal/controller/namespaced/database/dbnodeconsoleconnection"
	dbnodeconsolehistory "github.com/oracle/provider-oci/internal/controller/namespaced/database/dbnodeconsolehistory"
	dbnodesnapshot "github.com/oracle/provider-oci/internal/controller/namespaced/database/dbnodesnapshot"
	dbnodesnapshotmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/database/dbnodesnapshotmanagement"
	dbsystem "github.com/oracle/provider-oci/internal/controller/namespaced/database/dbsystem"
	dbsystemsupgrade "github.com/oracle/provider-oci/internal/controller/namespaced/database/dbsystemsupgrade"
	exadatainfrastructure "github.com/oracle/provider-oci/internal/controller/namespaced/database/exadatainfrastructure"
	exadatainfrastructurecompute "github.com/oracle/provider-oci/internal/controller/namespaced/database/exadatainfrastructurecompute"
	exadatainfrastructureconfigureexascalemanagement "github.com/oracle/provider-oci/internal/controller/namespaced/database/exadatainfrastructureconfigureexascalemanagement"
	exadatainfrastructurestorage "github.com/oracle/provider-oci/internal/controller/namespaced/database/exadatainfrastructurestorage"
	exadataiormconfig "github.com/oracle/provider-oci/internal/controller/namespaced/database/exadataiormconfig"
	exadbvmcluster "github.com/oracle/provider-oci/internal/controller/namespaced/database/exadbvmcluster"
	exascaledbstoragevault "github.com/oracle/provider-oci/internal/controller/namespaced/database/exascaledbstoragevault"
	executionaction "github.com/oracle/provider-oci/internal/controller/namespaced/database/executionaction"
	executionwindow "github.com/oracle/provider-oci/internal/controller/namespaced/database/executionwindow"
	externalcontainerdatabase "github.com/oracle/provider-oci/internal/controller/namespaced/database/externalcontainerdatabase"
	externalcontainerdatabasemanagement "github.com/oracle/provider-oci/internal/controller/namespaced/database/externalcontainerdatabasemanagement"
	externalcontainerdatabasesstackmonitoring "github.com/oracle/provider-oci/internal/controller/namespaced/database/externalcontainerdatabasesstackmonitoring"
	externaldatabaseconnector "github.com/oracle/provider-oci/internal/controller/namespaced/database/externaldatabaseconnector"
	externalnoncontainerdatabase "github.com/oracle/provider-oci/internal/controller/namespaced/database/externalnoncontainerdatabase"
	externalnoncontainerdatabasemanagement "github.com/oracle/provider-oci/internal/controller/namespaced/database/externalnoncontainerdatabasemanagement"
	externalnoncontainerdatabaseoperationsinsightsmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/database/externalnoncontainerdatabaseoperationsinsightsmanagement"
	externalnoncontainerdatabasesstackmonitoring "github.com/oracle/provider-oci/internal/controller/namespaced/database/externalnoncontainerdatabasesstackmonitoring"
	externalpluggabledatabase "github.com/oracle/provider-oci/internal/controller/namespaced/database/externalpluggabledatabase"
	externalpluggabledatabasemanagement "github.com/oracle/provider-oci/internal/controller/namespaced/database/externalpluggabledatabasemanagement"
	externalpluggabledatabaseoperationsinsightsmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/database/externalpluggabledatabaseoperationsinsightsmanagement"
	externalpluggabledatabasesstackmonitoring "github.com/oracle/provider-oci/internal/controller/namespaced/database/externalpluggabledatabasesstackmonitoring"
	keystore "github.com/oracle/provider-oci/internal/controller/namespaced/database/keystore"
	maintenancerun "github.com/oracle/provider-oci/internal/controller/namespaced/database/maintenancerun"
	managementautonomousdatabasedbmfeaturesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementautonomousdatabasedbmfeaturesmanagement"
	managementcloudasm "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementcloudasm"
	managementcloudasminstance "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementcloudasminstance"
	managementcloudcluster "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementcloudcluster"
	managementcloudclusterinstance "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementcloudclusterinstance"
	managementclouddbhome "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementclouddbhome"
	managementclouddbnode "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementclouddbnode"
	managementclouddbsystem "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementclouddbsystem"
	managementclouddbsystemclouddatabasemanagementsmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementclouddbsystemclouddatabasemanagementsmanagement"
	managementclouddbsystemcloudstackmonitoringsmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementclouddbsystemcloudstackmonitoringsmanagement"
	managementclouddbsystemconnector "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementclouddbsystemconnector"
	managementclouddbsystemdiscovery "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementclouddbsystemdiscovery"
	managementcloudexadatainfrastructure "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementcloudexadatainfrastructure"
	managementcloudexadatainfrastructuremanagedexadata "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementcloudexadatainfrastructuremanagedexadata"
	managementcloudexadatastorageconnector "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementcloudexadatastorageconnector"
	managementcloudexadatastoragegrid "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementcloudexadatastoragegrid"
	managementcloudexadatastorageserver "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementcloudexadatastorageserver"
	managementcloudlistener "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementcloudlistener"
	managementdatabasedbmfeaturesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementdatabasedbmfeaturesmanagement"
	managementdbmanagementprivateendpoint "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementdbmanagementprivateendpoint"
	managementexternalasm "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementexternalasm"
	managementexternalasminstance "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementexternalasminstance"
	managementexternalcluster "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementexternalcluster"
	managementexternalclusterinstance "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementexternalclusterinstance"
	managementexternalcontainerdatabasedbmfeaturesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementexternalcontainerdatabasedbmfeaturesmanagement"
	managementexternaldbhome "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementexternaldbhome"
	managementexternaldbnode "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementexternaldbnode"
	managementexternaldbsystem "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementexternaldbsystem"
	managementexternaldbsystemconnector "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementexternaldbsystemconnector"
	managementexternaldbsystemdatabasemanagementsmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementexternaldbsystemdatabasemanagementsmanagement"
	managementexternaldbsystemdiscovery "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementexternaldbsystemdiscovery"
	managementexternaldbsystemstackmonitoringsmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementexternaldbsystemstackmonitoringsmanagement"
	managementexternalexadatainfrastructure "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementexternalexadatainfrastructure"
	managementexternalexadatainfrastructureexadatamanagement "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementexternalexadatainfrastructureexadatamanagement"
	managementexternalexadatastorageconnector "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementexternalexadatastorageconnector"
	managementexternalexadatastoragegrid "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementexternalexadatastoragegrid"
	managementexternalexadatastorageserver "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementexternalexadatastorageserver"
	managementexternallistener "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementexternallistener"
	managementexternalmysqldatabase "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementexternalmysqldatabase"
	managementexternalmysqldatabaseconnector "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementexternalmysqldatabaseconnector"
	managementexternalmysqldatabaseexternalmysqldatabases "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementexternalmysqldatabaseexternalmysqldatabases"
	managementexternalnoncontainerdatabasedbmfeaturesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementexternalnoncontainerdatabasedbmfeaturesmanagement"
	managementexternalpluggabledatabasedbmfeaturesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementexternalpluggabledatabasedbmfeaturesmanagement"
	managementmanageddatabase "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementmanageddatabase"
	managementmanageddatabasegroup "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementmanageddatabasegroup"
	managementmanageddatabaseschangedatabaseparameter "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementmanageddatabaseschangedatabaseparameter"
	managementmanageddatabasesresetdatabaseparameter "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementmanageddatabasesresetdatabaseparameter"
	managementnamedcredential "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementnamedcredential"
	managementpluggabledatabasedbmfeaturesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/database/managementpluggabledatabasedbmfeaturesmanagement"
	migrationdatabase "github.com/oracle/provider-oci/internal/controller/namespaced/database/migration"
	migrationassessment "github.com/oracle/provider-oci/internal/controller/namespaced/database/migrationassessment"
	migrationassessmentassessoraction "github.com/oracle/provider-oci/internal/controller/namespaced/database/migrationassessmentassessoraction"
	migrationconnection "github.com/oracle/provider-oci/internal/controller/namespaced/database/migrationconnection"
	migrationjob "github.com/oracle/provider-oci/internal/controller/namespaced/database/migrationjob"
	migrationjobadvisorreportcheck "github.com/oracle/provider-oci/internal/controller/namespaced/database/migrationjobadvisorreportcheck"
	migrationmigration "github.com/oracle/provider-oci/internal/controller/namespaced/database/migrationmigration"
	oneoffpatch "github.com/oracle/provider-oci/internal/controller/namespaced/database/oneoffpatch"
	pluggabledatabase "github.com/oracle/provider-oci/internal/controller/namespaced/database/pluggabledatabase"
	pluggabledatabasepluggabledatabasemanagementsmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/database/pluggabledatabasepluggabledatabasemanagementsmanagement"
	pluggabledatabaseslocalclone "github.com/oracle/provider-oci/internal/controller/namespaced/database/pluggabledatabaseslocalclone"
	pluggabledatabasesnapshot "github.com/oracle/provider-oci/internal/controller/namespaced/database/pluggabledatabasesnapshot"
	pluggabledatabasesremoteclone "github.com/oracle/provider-oci/internal/controller/namespaced/database/pluggabledatabasesremoteclone"
	scheduledaction "github.com/oracle/provider-oci/internal/controller/namespaced/database/scheduledaction"
	schedulingplan "github.com/oracle/provider-oci/internal/controller/namespaced/database/schedulingplan"
	schedulingpolicy "github.com/oracle/provider-oci/internal/controller/namespaced/database/schedulingpolicy"
	schedulingpolicyschedulingwindow "github.com/oracle/provider-oci/internal/controller/namespaced/database/schedulingpolicyschedulingwindow"
	toolsdatabasetoolsconnection "github.com/oracle/provider-oci/internal/controller/namespaced/database/toolsdatabasetoolsconnection"
	toolsdatabasetoolsdatabaseapigatewayconfig "github.com/oracle/provider-oci/internal/controller/namespaced/database/toolsdatabasetoolsdatabaseapigatewayconfig"
	toolsdatabasetoolsidentity "github.com/oracle/provider-oci/internal/controller/namespaced/database/toolsdatabasetoolsidentity"
	toolsdatabasetoolsmcpserver "github.com/oracle/provider-oci/internal/controller/namespaced/database/toolsdatabasetoolsmcpserver"
	toolsdatabasetoolsmcptoolset "github.com/oracle/provider-oci/internal/controller/namespaced/database/toolsdatabasetoolsmcptoolset"
	toolsdatabasetoolsprivateendpoint "github.com/oracle/provider-oci/internal/controller/namespaced/database/toolsdatabasetoolsprivateendpoint"
	toolsdatabasetoolssqlreport "github.com/oracle/provider-oci/internal/controller/namespaced/database/toolsdatabasetoolssqlreport"
	toolsruntimedatabasetoolsapigatewayconfigpoolapispec "github.com/oracle/provider-oci/internal/controller/namespaced/database/toolsruntimedatabasetoolsapigatewayconfigpoolapispec"
	toolsruntimedatabasetoolsapigatewayconfigpoolautoapispec "github.com/oracle/provider-oci/internal/controller/namespaced/database/toolsruntimedatabasetoolsapigatewayconfigpoolautoapispec"
	toolsruntimedatabasetoolsconnectioncredential "github.com/oracle/provider-oci/internal/controller/namespaced/database/toolsruntimedatabasetoolsconnectioncredential"
	toolsruntimedatabasetoolsconnectioncredentialexecutegrantee "github.com/oracle/provider-oci/internal/controller/namespaced/database/toolsruntimedatabasetoolsconnectioncredentialexecutegrantee"
	toolsruntimedatabasetoolsconnectioncredentialpublicsynonym "github.com/oracle/provider-oci/internal/controller/namespaced/database/toolsruntimedatabasetoolsconnectioncredentialpublicsynonym"
	toolsruntimedatabasetoolsconnectionpropertyset "github.com/oracle/provider-oci/internal/controller/namespaced/database/toolsruntimedatabasetoolsconnectionpropertyset"
	toolsruntimedatabasetoolsdatabaseapigatewayconfigglobal "github.com/oracle/provider-oci/internal/controller/namespaced/database/toolsruntimedatabasetoolsdatabaseapigatewayconfigglobal"
	toolsruntimedatabasetoolsdatabaseapigatewayconfigpool "github.com/oracle/provider-oci/internal/controller/namespaced/database/toolsruntimedatabasetoolsdatabaseapigatewayconfigpool"
	vmcluster "github.com/oracle/provider-oci/internal/controller/namespaced/database/vmcluster"
	vmclusteraddvirtualmachine "github.com/oracle/provider-oci/internal/controller/namespaced/database/vmclusteraddvirtualmachine"
	vmclusternetwork "github.com/oracle/provider-oci/internal/controller/namespaced/database/vmclusternetwork"
	vmclusterremovevirtualmachine "github.com/oracle/provider-oci/internal/controller/namespaced/database/vmclusterremovevirtualmachine"
	catalog "github.com/oracle/provider-oci/internal/controller/namespaced/datacatalog/catalog"
	catalogprivateendpoint "github.com/oracle/provider-oci/internal/controller/namespaced/datacatalog/catalogprivateendpoint"
	connection "github.com/oracle/provider-oci/internal/controller/namespaced/datacatalog/connection"
	dataasset "github.com/oracle/provider-oci/internal/controller/namespaced/datacatalog/dataasset"
	metastore "github.com/oracle/provider-oci/internal/controller/namespaced/datacatalog/metastore"
	infrastructure "github.com/oracle/provider-oci/internal/controller/namespaced/datacc/infrastructure"
	vmclusternetworkdatacc "github.com/oracle/provider-oci/internal/controller/namespaced/datacc/vmclusternetwork"
	vminstance "github.com/oracle/provider-oci/internal/controller/namespaced/datacc/vminstance"
	application "github.com/oracle/provider-oci/internal/controller/namespaced/dataflow/application"
	invokerun "github.com/oracle/provider-oci/internal/controller/namespaced/dataflow/invokerun"
	pool "github.com/oracle/provider-oci/internal/controller/namespaced/dataflow/pool"
	privateendpoint "github.com/oracle/provider-oci/internal/controller/namespaced/dataflow/privateendpoint"
	runstatement "github.com/oracle/provider-oci/internal/controller/namespaced/dataflow/runstatement"
	sqlendpoint "github.com/oracle/provider-oci/internal/controller/namespaced/dataflow/sqlendpoint"
	workspace "github.com/oracle/provider-oci/internal/controller/namespaced/dataintegration/workspace"
	workspaceapplication "github.com/oracle/provider-oci/internal/controller/namespaced/dataintegration/workspaceapplication"
	workspaceapplicationpatch "github.com/oracle/provider-oci/internal/controller/namespaced/dataintegration/workspaceapplicationpatch"
	workspaceapplicationschedule "github.com/oracle/provider-oci/internal/controller/namespaced/dataintegration/workspaceapplicationschedule"
	workspaceapplicationtaskschedule "github.com/oracle/provider-oci/internal/controller/namespaced/dataintegration/workspaceapplicationtaskschedule"
	workspaceexportrequest "github.com/oracle/provider-oci/internal/controller/namespaced/dataintegration/workspaceexportrequest"
	workspacefolder "github.com/oracle/provider-oci/internal/controller/namespaced/dataintegration/workspacefolder"
	workspaceimportrequest "github.com/oracle/provider-oci/internal/controller/namespaced/dataintegration/workspaceimportrequest"
	workspaceproject "github.com/oracle/provider-oci/internal/controller/namespaced/dataintegration/workspaceproject"
	workspacetask "github.com/oracle/provider-oci/internal/controller/namespaced/dataintegration/workspacetask"
	dataset "github.com/oracle/provider-oci/internal/controller/namespaced/datalabelingservice/dataset"
	addsdmcolumns "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/addsdmcolumns"
	alert "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/alert"
	alertpolicy "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/alertpolicy"
	alertpolicyrule "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/alertpolicyrule"
	attributeset "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/attributeset"
	auditarchiveretrieval "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/auditarchiveretrieval"
	auditpolicy "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/auditpolicy"
	auditpolicymanagement "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/auditpolicymanagement"
	auditprofile "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/auditprofile"
	auditprofilemanagement "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/auditprofilemanagement"
	audittrail "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/audittrail"
	audittrailmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/audittrailmanagement"
	calculateauditvolumeavailable "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/calculateauditvolumeavailable"
	calculateauditvolumecollected "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/calculateauditvolumecollected"
	comparesecurityassessment "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/comparesecurityassessment"
	compareuserassessment "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/compareuserassessment"
	databasesecurityconfig "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/databasesecurityconfig"
	databasesecurityconfigmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/databasesecurityconfigmanagement"
	datasafeconfiguration "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/datasafeconfiguration"
	datasafeprivateendpoint "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/datasafeprivateendpoint"
	discoveryjob "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/discoveryjob"
	discoveryjobsresult "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/discoveryjobsresult"
	generateonpremconnectorconfiguration "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/generateonpremconnectorconfiguration"
	librarymaskingformat "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/librarymaskingformat"
	maskdata "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/maskdata"
	maskingpoliciesapplydifferencetomaskingcolumns "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/maskingpoliciesapplydifferencetomaskingcolumns"
	maskingpoliciesmaskingcolumn "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/maskingpoliciesmaskingcolumn"
	maskingpolicy "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/maskingpolicy"
	maskingpolicyhealthreportmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/maskingpolicyhealthreportmanagement"
	maskingreportmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/maskingreportmanagement"
	onpremconnector "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/onpremconnector"
	report "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/report"
	reportdefinition "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/reportdefinition"
	sdmmaskingpolicydifference "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/sdmmaskingpolicydifference"
	securityassessment "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/securityassessment"
	securityassessmentcheck "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/securityassessmentcheck"
	securityassessmentfinding "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/securityassessmentfinding"
	securitypolicy "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/securitypolicy"
	securitypolicyconfig "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/securitypolicyconfig"
	securitypolicydeployment "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/securitypolicydeployment"
	securitypolicydeploymentmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/securitypolicydeploymentmanagement"
	securitypolicymanagement "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/securitypolicymanagement"
	sensitivedatamodel "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/sensitivedatamodel"
	sensitivedatamodelreferentialrelation "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/sensitivedatamodelreferentialrelation"
	sensitivedatamodelsapplydiscoveryjobresults "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/sensitivedatamodelsapplydiscoveryjobresults"
	sensitivedatamodelssensitivecolumn "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/sensitivedatamodelssensitivecolumn"
	sensitivetype "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/sensitivetype"
	sensitivetypegroup "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/sensitivetypegroup"
	sensitivetypegroupgroupedsensitivetype "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/sensitivetypegroupgroupedsensitivetype"
	sensitivetypesexport "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/sensitivetypesexport"
	setsecurityassessmentbaseline "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/setsecurityassessmentbaseline"
	setsecurityassessmentbaselinemanagement "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/setsecurityassessmentbaselinemanagement"
	setuserassessmentbaseline "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/setuserassessmentbaseline"
	setuserassessmentbaselinemanagement "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/setuserassessmentbaselinemanagement"
	sqlcollection "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/sqlcollection"
	sqlfirewallpolicy "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/sqlfirewallpolicy"
	sqlfirewallpolicymanagement "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/sqlfirewallpolicymanagement"
	targetalertpolicyassociation "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/targetalertpolicyassociation"
	targetdatabase "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/targetdatabase"
	targetdatabasegroup "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/targetdatabasegroup"
	targetdatabasepeertargetdatabase "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/targetdatabasepeertargetdatabase"
	unifiedauditpolicy "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/unifiedauditpolicy"
	unifiedauditpolicydefinition "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/unifiedauditpolicydefinition"
	unsetsecurityassessmentbaseline "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/unsetsecurityassessmentbaseline"
	unsetsecurityassessmentbaselinemanagement "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/unsetsecurityassessmentbaselinemanagement"
	unsetuserassessmentbaseline "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/unsetuserassessmentbaseline"
	unsetuserassessmentbaselinemanagement "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/unsetuserassessmentbaselinemanagement"
	userassessment "github.com/oracle/provider-oci/internal/controller/namespaced/datasafe/userassessment"
	computetarget "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/computetarget"
	jobdatascience "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/job"
	jobrun "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/jobrun"
	mlapplication "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/mlapplication"
	mlapplicationimplementation "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/mlapplicationimplementation"
	mlapplicationinstance "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/mlapplicationinstance"
	modeldatascience "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/model"
	modelartifactexport "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/modelartifactexport"
	modelartifactimport "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/modelartifactimport"
	modelcustommetadataartifact "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/modelcustommetadataartifact"
	modeldefinedmetadataartifact "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/modeldefinedmetadataartifact"
	modeldeployment "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/modeldeployment"
	modelgroup "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/modelgroup"
	modelgroupartifact "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/modelgroupartifact"
	modelgroupversionhistory "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/modelgroupversionhistory"
	modelprovenance "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/modelprovenance"
	modelversionset "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/modelversionset"
	notebooksession "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/notebooksession"
	pipeline "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/pipeline"
	pipelinerun "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/pipelinerun"
	privateendpointdatascience "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/privateendpoint"
	projectdatascience "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/project"
	schedule "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/schedule"
	vulnerabilityscan "github.com/oracle/provider-oci/internal/controller/namespaced/dblm/vulnerabilityscan"
	multicloudresourcediscovery "github.com/oracle/provider-oci/internal/controller/namespaced/dbmulticloud/multicloudresourcediscovery"
	oracledbawsidentityconnector "github.com/oracle/provider-oci/internal/controller/namespaced/dbmulticloud/oracledbawsidentityconnector"
	oracledbawskey "github.com/oracle/provider-oci/internal/controller/namespaced/dbmulticloud/oracledbawskey"
	oracledbazureblobcontainer "github.com/oracle/provider-oci/internal/controller/namespaced/dbmulticloud/oracledbazureblobcontainer"
	oracledbazureblobmount "github.com/oracle/provider-oci/internal/controller/namespaced/dbmulticloud/oracledbazureblobmount"
	oracledbazureconnector "github.com/oracle/provider-oci/internal/controller/namespaced/dbmulticloud/oracledbazureconnector"
	oracledbazurevault "github.com/oracle/provider-oci/internal/controller/namespaced/dbmulticloud/oracledbazurevault"
	oracledbazurevaultassociation "github.com/oracle/provider-oci/internal/controller/namespaced/dbmulticloud/oracledbazurevaultassociation"
	oracledbgcpidentityconnector "github.com/oracle/provider-oci/internal/controller/namespaced/dbmulticloud/oracledbgcpidentityconnector"
	oracledbgcpkeyring "github.com/oracle/provider-oci/internal/controller/namespaced/dbmulticloud/oracledbgcpkeyring"
	instanceddfs "github.com/oracle/provider-oci/internal/controller/namespaced/ddfs/instance"
	delegationcontrol "github.com/oracle/provider-oci/internal/controller/namespaced/delegateaccesscontrol/delegationcontrol"
	delegationsubscription "github.com/oracle/provider-oci/internal/controller/namespaced/delegateaccesscontrol/delegationsubscription"
	occdemandsignal "github.com/oracle/provider-oci/internal/controller/namespaced/demandsignal/occdemandsignal"
	occmetricalarm "github.com/oracle/provider-oci/internal/controller/namespaced/demandsignal/occmetricalarm"
	desktoppool "github.com/oracle/provider-oci/internal/controller/namespaced/desktops/desktoppool"
	buildpipeline "github.com/oracle/provider-oci/internal/controller/namespaced/devops/buildpipeline"
	buildpipelinestage "github.com/oracle/provider-oci/internal/controller/namespaced/devops/buildpipelinestage"
	buildrun "github.com/oracle/provider-oci/internal/controller/namespaced/devops/buildrun"
	connectiondevops "github.com/oracle/provider-oci/internal/controller/namespaced/devops/connection"
	deployartifact "github.com/oracle/provider-oci/internal/controller/namespaced/devops/deployartifact"
	deployenvironment "github.com/oracle/provider-oci/internal/controller/namespaced/devops/deployenvironment"
	deploymentdevops "github.com/oracle/provider-oci/internal/controller/namespaced/devops/deployment"
	deploypipeline "github.com/oracle/provider-oci/internal/controller/namespaced/devops/deploypipeline"
	deploystage "github.com/oracle/provider-oci/internal/controller/namespaced/devops/deploystage"
	projectdevops "github.com/oracle/provider-oci/internal/controller/namespaced/devops/project"
	projectrepositorysetting "github.com/oracle/provider-oci/internal/controller/namespaced/devops/projectrepositorysetting"
	repositorydevops "github.com/oracle/provider-oci/internal/controller/namespaced/devops/repository"
	repositorymirror "github.com/oracle/provider-oci/internal/controller/namespaced/devops/repositorymirror"
	repositoryprotectedbranchmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/devops/repositoryprotectedbranchmanagement"
	repositoryref "github.com/oracle/provider-oci/internal/controller/namespaced/devops/repositoryref"
	repositorysetting "github.com/oracle/provider-oci/internal/controller/namespaced/devops/repositorysetting"
	trigger "github.com/oracle/provider-oci/internal/controller/namespaced/devops/trigger"
	stack "github.com/oracle/provider-oci/internal/controller/namespaced/dif/stack"
	automaticdrconfiguration "github.com/oracle/provider-oci/internal/controller/namespaced/disasterrecovery/automaticdrconfiguration"
	drplan "github.com/oracle/provider-oci/internal/controller/namespaced/disasterrecovery/drplan"
	drplanexecution "github.com/oracle/provider-oci/internal/controller/namespaced/disasterrecovery/drplanexecution"
	drprotectiongroup "github.com/oracle/provider-oci/internal/controller/namespaced/disasterrecovery/drprotectiongroup"
	distributedautonomousdatabase "github.com/oracle/provider-oci/internal/controller/namespaced/distributeddatabase/distributedautonomousdatabase"
	distributeddatabase "github.com/oracle/provider-oci/internal/controller/namespaced/distributeddatabase/distributeddatabase"
	distributeddatabaseprivateendpoint "github.com/oracle/provider-oci/internal/controller/namespaced/distributeddatabase/distributeddatabaseprivateendpoint"
	actioncreatezonefromzonefile "github.com/oracle/provider-oci/internal/controller/namespaced/dns/actioncreatezonefromzonefile"
	record "github.com/oracle/provider-oci/internal/controller/namespaced/dns/record"
	resolver "github.com/oracle/provider-oci/internal/controller/namespaced/dns/resolver"
	resolverendpoint "github.com/oracle/provider-oci/internal/controller/namespaced/dns/resolverendpoint"
	rrset "github.com/oracle/provider-oci/internal/controller/namespaced/dns/rrset"
	steeringpolicy "github.com/oracle/provider-oci/internal/controller/namespaced/dns/steeringpolicy"
	steeringpolicyattachment "github.com/oracle/provider-oci/internal/controller/namespaced/dns/steeringpolicyattachment"
	tsigkey "github.com/oracle/provider-oci/internal/controller/namespaced/dns/tsigkey"
	view "github.com/oracle/provider-oci/internal/controller/namespaced/dns/view"
	zone "github.com/oracle/provider-oci/internal/controller/namespaced/dns/zone"
	zonepromotednsseckeyversion "github.com/oracle/provider-oci/internal/controller/namespaced/dns/zonepromotednsseckeyversion"
	zonestagednsseckeyversion "github.com/oracle/provider-oci/internal/controller/namespaced/dns/zonestagednsseckeyversion"
	dkim "github.com/oracle/provider-oci/internal/controller/namespaced/email/dkim"
	emaildomain "github.com/oracle/provider-oci/internal/controller/namespaced/email/emaildomain"
	emailippool "github.com/oracle/provider-oci/internal/controller/namespaced/email/emailippool"
	emailreturnpath "github.com/oracle/provider-oci/internal/controller/namespaced/email/emailreturnpath"
	sender "github.com/oracle/provider-oci/internal/controller/namespaced/email/sender"
	suppression "github.com/oracle/provider-oci/internal/controller/namespaced/email/suppression"
	rule "github.com/oracle/provider-oci/internal/controller/namespaced/events/rule"
	export "github.com/oracle/provider-oci/internal/controller/namespaced/filestorage/export"
	exportset "github.com/oracle/provider-oci/internal/controller/namespaced/filestorage/exportset"
	filesystem "github.com/oracle/provider-oci/internal/controller/namespaced/filestorage/filesystem"
	filesystemquotarule "github.com/oracle/provider-oci/internal/controller/namespaced/filestorage/filesystemquotarule"
	filesystemsnapshotpolicy "github.com/oracle/provider-oci/internal/controller/namespaced/filestorage/filesystemsnapshotpolicy"
	mounttarget "github.com/oracle/provider-oci/internal/controller/namespaced/filestorage/mounttarget"
	outboundconnector "github.com/oracle/provider-oci/internal/controller/namespaced/filestorage/outboundconnector"
	replication "github.com/oracle/provider-oci/internal/controller/namespaced/filestorage/replication"
	snapshot "github.com/oracle/provider-oci/internal/controller/namespaced/filestorage/snapshot"
	catalogitem "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/catalogitem"
	compliancepolicyrule "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/compliancepolicyrule"
	fleet "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/fleet"
	fleetcredential "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/fleetcredential"
	fleetproperty "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/fleetproperty"
	fleetresource "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/fleetresource"
	maintenancewindow "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/maintenancewindow"
	onboarding "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/onboarding"
	patch "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/patch"
	platformconfiguration "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/platformconfiguration"
	property "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/property"
	provision "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/provision"
	runbook "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/runbook"
	runbookversion "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/runbookversion"
	schedulerdefinition "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/schedulerdefinition"
	taskrecord "github.com/oracle/provider-oci/internal/controller/namespaced/fleetappsmanagement/taskrecord"
	fsucollection "github.com/oracle/provider-oci/internal/controller/namespaced/fleetsoftwareupdate/fsucollection"
	fsucycle "github.com/oracle/provider-oci/internal/controller/namespaced/fleetsoftwareupdate/fsucycle"
	fsureadinesscheck "github.com/oracle/provider-oci/internal/controller/namespaced/fleetsoftwareupdate/fsureadinesscheck"
	applicationfunctions "github.com/oracle/provider-oci/internal/controller/namespaced/functions/application"
	function "github.com/oracle/provider-oci/internal/controller/namespaced/functions/function"
	invokefunction "github.com/oracle/provider-oci/internal/controller/namespaced/functions/invokefunction"
	fusionenvironment "github.com/oracle/provider-oci/internal/controller/namespaced/fusionapps/fusionenvironment"
	fusionenvironmentadminuser "github.com/oracle/provider-oci/internal/controller/namespaced/fusionapps/fusionenvironmentadminuser"
	fusionenvironmentdatamaskingactivity "github.com/oracle/provider-oci/internal/controller/namespaced/fusionapps/fusionenvironmentdatamaskingactivity"
	fusionenvironmentfamily "github.com/oracle/provider-oci/internal/controller/namespaced/fusionapps/fusionenvironmentfamily"
	fusionenvironmentrefreshactivity "github.com/oracle/provider-oci/internal/controller/namespaced/fusionapps/fusionenvironmentrefreshactivity"
	fusionenvironmentserviceattachment "github.com/oracle/provider-oci/internal/controller/namespaced/fusionapps/fusionenvironmentserviceattachment"
	gdppipeline "github.com/oracle/provider-oci/internal/controller/namespaced/gdp/gdppipeline"
	agentagent "github.com/oracle/provider-oci/internal/controller/namespaced/generativeai/agentagent"
	agentagentendpoint "github.com/oracle/provider-oci/internal/controller/namespaced/generativeai/agentagentendpoint"
	agentdataingestionjob "github.com/oracle/provider-oci/internal/controller/namespaced/generativeai/agentdataingestionjob"
	agentdatasource "github.com/oracle/provider-oci/internal/controller/namespaced/generativeai/agentdatasource"
	agentknowledgebase "github.com/oracle/provider-oci/internal/controller/namespaced/generativeai/agentknowledgebase"
	agentprovisionedcapacity "github.com/oracle/provider-oci/internal/controller/namespaced/generativeai/agentprovisionedcapacity"
	agenttool "github.com/oracle/provider-oci/internal/controller/namespaced/generativeai/agenttool"
	dedicatedaicluster "github.com/oracle/provider-oci/internal/controller/namespaced/generativeai/dedicatedaicluster"
	endpointgenerativeai "github.com/oracle/provider-oci/internal/controller/namespaced/generativeai/endpoint"
	generativeaiprivateendpoint "github.com/oracle/provider-oci/internal/controller/namespaced/generativeai/generativeaiprivateendpoint"
	hostedapplication "github.com/oracle/provider-oci/internal/controller/namespaced/generativeai/hostedapplication"
	hostedapplicationiam "github.com/oracle/provider-oci/internal/controller/namespaced/generativeai/hostedapplicationiam"
	hostedapplicationstorage "github.com/oracle/provider-oci/internal/controller/namespaced/generativeai/hostedapplicationstorage"
	hosteddeployment "github.com/oracle/provider-oci/internal/controller/namespaced/generativeai/hosteddeployment"
	importedmodel "github.com/oracle/provider-oci/internal/controller/namespaced/generativeai/importedmodel"
	modelgenerativeai "github.com/oracle/provider-oci/internal/controller/namespaced/generativeai/model"
	projectgenerativeai "github.com/oracle/provider-oci/internal/controller/namespaced/generativeai/project"
	semanticstore "github.com/oracle/provider-oci/internal/controller/namespaced/generativeai/semanticstore"
	artifactbypath "github.com/oracle/provider-oci/internal/controller/namespaced/genericartifactscontent/artifactbypath"
	connectiongoldengate "github.com/oracle/provider-oci/internal/controller/namespaced/goldengate/connection"
	connectionassignment "github.com/oracle/provider-oci/internal/controller/namespaced/goldengate/connectionassignment"
	databaseregistration "github.com/oracle/provider-oci/internal/controller/namespaced/goldengate/databaseregistration"
	deploymentgoldengate "github.com/oracle/provider-oci/internal/controller/namespaced/goldengate/deployment"
	deploymentbackup "github.com/oracle/provider-oci/internal/controller/namespaced/goldengate/deploymentbackup"
	deploymentcertificate "github.com/oracle/provider-oci/internal/controller/namespaced/goldengate/deploymentcertificate"
	pipelinegoldengate "github.com/oracle/provider-oci/internal/controller/namespaced/goldengate/pipeline"
	httpmonitor "github.com/oracle/provider-oci/internal/controller/namespaced/healthchecks/httpmonitor"
	httpprobe "github.com/oracle/provider-oci/internal/controller/namespaced/healthchecks/httpprobe"
	pingmonitor "github.com/oracle/provider-oci/internal/controller/namespaced/healthchecks/pingmonitor"
	pingprobe "github.com/oracle/provider-oci/internal/controller/namespaced/healthchecks/pingprobe"
	apikey "github.com/oracle/provider-oci/internal/controller/namespaced/identity/apikey"
	authenticationpolicy "github.com/oracle/provider-oci/internal/controller/namespaced/identity/authenticationpolicy"
	authtoken "github.com/oracle/provider-oci/internal/controller/namespaced/identity/authtoken"
	compartment "github.com/oracle/provider-oci/internal/controller/namespaced/identity/compartment"
	customersecretkey "github.com/oracle/provider-oci/internal/controller/namespaced/identity/customersecretkey"
	dbcredential "github.com/oracle/provider-oci/internal/controller/namespaced/identity/dbcredential"
	domain "github.com/oracle/provider-oci/internal/controller/namespaced/identity/domain"
	domainreplicationtoregion "github.com/oracle/provider-oci/internal/controller/namespaced/identity/domainreplicationtoregion"
	dynamicgroup "github.com/oracle/provider-oci/internal/controller/namespaced/identity/dynamicgroup"
	group "github.com/oracle/provider-oci/internal/controller/namespaced/identity/group"
	identityprovider "github.com/oracle/provider-oci/internal/controller/namespaced/identity/identityprovider"
	idpgroupmapping "github.com/oracle/provider-oci/internal/controller/namespaced/identity/idpgroupmapping"
	importstandardtagsmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/identity/importstandardtagsmanagement"
	networksource "github.com/oracle/provider-oci/internal/controller/namespaced/identity/networksource"
	policy "github.com/oracle/provider-oci/internal/controller/namespaced/identity/policy"
	smtpcredential "github.com/oracle/provider-oci/internal/controller/namespaced/identity/smtpcredential"
	tag "github.com/oracle/provider-oci/internal/controller/namespaced/identity/tag"
	tagdefault "github.com/oracle/provider-oci/internal/controller/namespaced/identity/tagdefault"
	tagnamespace "github.com/oracle/provider-oci/internal/controller/namespaced/identity/tagnamespace"
	uipassword "github.com/oracle/provider-oci/internal/controller/namespaced/identity/uipassword"
	user "github.com/oracle/provider-oci/internal/controller/namespaced/identity/user"
	usercapabilitiesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/identity/usercapabilitiesmanagement"
	usergroupmembership "github.com/oracle/provider-oci/internal/controller/namespaced/identity/usergroupmembership"
	generatescopedaccesstoken "github.com/oracle/provider-oci/internal/controller/namespaced/identitydataplane/generatescopedaccesstoken"
	accountrecoverysetting "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/accountrecoverysetting"
	apikeyidentitydomains "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/apikey"
	app "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/app"
	approle "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/approle"
	approvalworkflow "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/approvalworkflow"
	approvalworkflowassignment "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/approvalworkflowassignment"
	approvalworkflowstep "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/approvalworkflowstep"
	authenticationfactorsetting "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/authenticationfactorsetting"
	authtokenidentitydomains "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/authtoken"
	cloudgate "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/cloudgate"
	cloudgatemapping "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/cloudgatemapping"
	cloudgateserver "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/cloudgateserver"
	condition "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/condition"
	customersecretkeyidentitydomains "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/customersecretkey"
	dynamicresourcegroup "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/dynamicresourcegroup"
	grant "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/grant"
	groupidentitydomains "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/group"
	identityproofingprovider "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/identityproofingprovider"
	identityproofingprovidertemplate "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/identityproofingprovidertemplate"
	identitypropagationtrust "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/identitypropagationtrust"
	identityprovideridentitydomains "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/identityprovider"
	identitysetting "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/identitysetting"
	kmsisetting "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/kmsisetting"
	mappedattribute "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/mappedattribute"
	myapikey "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/myapikey"
	myauthtoken "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/myauthtoken"
	mycustomersecretkey "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/mycustomersecretkey"
	myoauth2clientcredential "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/myoauth2clientcredential"
	myrequest "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/myrequest"
	mysmtpcredential "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/mysmtpcredential"
	mysupportaccount "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/mysupportaccount"
	myuserdbcredential "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/myuserdbcredential"
	networkperimeter "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/networkperimeter"
	notificationsetting "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/notificationsetting"
	oauth2clientcredential "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/oauth2clientcredential"
	oauthclientcertificate "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/oauthclientcertificate"
	oauthpartnercertificate "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/oauthpartnercertificate"
	passwordpolicy "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/passwordpolicy"
	policyidentitydomains "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/policy"
	ruleidentitydomains "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/rule"
	securityquestion "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/securityquestion"
	securityquestionsetting "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/securityquestionsetting"
	selfregistrationprofile "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/selfregistrationprofile"
	setting "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/setting"
	smtpcredentialidentitydomains "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/smtpcredential"
	socialidentityprovider "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/socialidentityprovider"
	useridentitydomains "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/user"
	userdbcredential "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/userdbcredential"
	integrationinstance "github.com/oracle/provider-oci/internal/controller/namespaced/integration/integrationinstance"
	oraclemanagedcustomendpoint "github.com/oracle/provider-oci/internal/controller/namespaced/integration/oraclemanagedcustomendpoint"
	privateendpointoutboundconnection "github.com/oracle/provider-oci/internal/controller/namespaced/integration/privateendpointoutboundconnection"
	digitaltwinadapter "github.com/oracle/provider-oci/internal/controller/namespaced/iot/digitaltwinadapter"
	digitaltwininstance "github.com/oracle/provider-oci/internal/controller/namespaced/iot/digitaltwininstance"
	digitaltwininstanceinvokerawcommand "github.com/oracle/provider-oci/internal/controller/namespaced/iot/digitaltwininstanceinvokerawcommand"
	digitaltwinmodel "github.com/oracle/provider-oci/internal/controller/namespaced/iot/digitaltwinmodel"
	digitaltwinrelationship "github.com/oracle/provider-oci/internal/controller/namespaced/iot/digitaltwinrelationship"
	iotdomain "github.com/oracle/provider-oci/internal/controller/namespaced/iot/iotdomain"
	iotdomainchangedataretentionperiod "github.com/oracle/provider-oci/internal/controller/namespaced/iot/iotdomainchangedataretentionperiod"
	iotdomainconfiguredataaccess "github.com/oracle/provider-oci/internal/controller/namespaced/iot/iotdomainconfiguredataaccess"
	iotdomaingroup "github.com/oracle/provider-oci/internal/controller/namespaced/iot/iotdomaingroup"
	iotdomaingroupconfiguredataaccess "github.com/oracle/provider-oci/internal/controller/namespaced/iot/iotdomaingroupconfiguredataaccess"
	fleetjms "github.com/oracle/provider-oci/internal/controller/namespaced/jms/fleet"
	fleetadvancedfeatureconfiguration "github.com/oracle/provider-oci/internal/controller/namespaced/jms/fleetadvancedfeatureconfiguration"
	fleetagentconfiguration "github.com/oracle/provider-oci/internal/controller/namespaced/jms/fleetagentconfiguration"
	jmsplugin "github.com/oracle/provider-oci/internal/controller/namespaced/jms/jmsplugin"
	taskschedule "github.com/oracle/provider-oci/internal/controller/namespaced/jms/taskschedule"
	javadownloadreport "github.com/oracle/provider-oci/internal/controller/namespaced/jmsjavadownloads/javadownloadreport"
	javadownloadtoken "github.com/oracle/provider-oci/internal/controller/namespaced/jmsjavadownloads/javadownloadtoken"
	javalicenseacceptancerecord "github.com/oracle/provider-oci/internal/controller/namespaced/jmsjavadownloads/javalicenseacceptancerecord"
	analyzeapplicationsconfiguration "github.com/oracle/provider-oci/internal/controller/namespaced/jmsutils/analyzeapplicationsconfiguration"
	subscriptionacknowledgmentconfiguration "github.com/oracle/provider-oci/internal/controller/namespaced/jmsutils/subscriptionacknowledgmentconfiguration"
	ekmsprivateendpoint "github.com/oracle/provider-oci/internal/controller/namespaced/kms/ekmsprivateendpoint"
	encrypteddata "github.com/oracle/provider-oci/internal/controller/namespaced/kms/encrypteddata"
	generatedkey "github.com/oracle/provider-oci/internal/controller/namespaced/kms/generatedkey"
	key "github.com/oracle/provider-oci/internal/controller/namespaced/kms/key"
	keyversion "github.com/oracle/provider-oci/internal/controller/namespaced/kms/keyversion"
	sign "github.com/oracle/provider-oci/internal/controller/namespaced/kms/sign"
	vault "github.com/oracle/provider-oci/internal/controller/namespaced/kms/vault"
	vaultreplication "github.com/oracle/provider-oci/internal/controller/namespaced/kms/vaultreplication"
	verify "github.com/oracle/provider-oci/internal/controller/namespaced/kms/verify"
	configurationlicensemanager "github.com/oracle/provider-oci/internal/controller/namespaced/licensemanager/configuration"
	licenserecord "github.com/oracle/provider-oci/internal/controller/namespaced/licensemanager/licenserecord"
	productlicense "github.com/oracle/provider-oci/internal/controller/namespaced/licensemanager/productlicense"
	quota "github.com/oracle/provider-oci/internal/controller/namespaced/limits/quota"
	backend "github.com/oracle/provider-oci/internal/controller/namespaced/loadbalancer/backend"
	backendset "github.com/oracle/provider-oci/internal/controller/namespaced/loadbalancer/backendset"
	certificateloadbalancer "github.com/oracle/provider-oci/internal/controller/namespaced/loadbalancer/certificate"
	lbhostname "github.com/oracle/provider-oci/internal/controller/namespaced/loadbalancer/lbhostname"
	listener "github.com/oracle/provider-oci/internal/controller/namespaced/loadbalancer/listener"
	loadbalancer "github.com/oracle/provider-oci/internal/controller/namespaced/loadbalancer/loadbalancer"
	pathrouteset "github.com/oracle/provider-oci/internal/controller/namespaced/loadbalancer/pathrouteset"
	routingpolicy "github.com/oracle/provider-oci/internal/controller/namespaced/loadbalancer/routingpolicy"
	ruleset "github.com/oracle/provider-oci/internal/controller/namespaced/loadbalancer/ruleset"
	sslciphersuite "github.com/oracle/provider-oci/internal/controller/namespaced/loadbalancer/sslciphersuite"
	loganalyticsentity "github.com/oracle/provider-oci/internal/controller/namespaced/loganalytics/loganalyticsentity"
	loganalyticsentityassociationsadd "github.com/oracle/provider-oci/internal/controller/namespaced/loganalytics/loganalyticsentityassociationsadd"
	loganalyticsentityassociationsremove "github.com/oracle/provider-oci/internal/controller/namespaced/loganalytics/loganalyticsentityassociationsremove"
	loganalyticsentitytype "github.com/oracle/provider-oci/internal/controller/namespaced/loganalytics/loganalyticsentitytype"
	loganalyticsimportcustomcontent "github.com/oracle/provider-oci/internal/controller/namespaced/loganalytics/loganalyticsimportcustomcontent"
	loganalyticsloggroup "github.com/oracle/provider-oci/internal/controller/namespaced/loganalytics/loganalyticsloggroup"
	loganalyticsobjectcollectionrule "github.com/oracle/provider-oci/internal/controller/namespaced/loganalytics/loganalyticsobjectcollectionrule"
	loganalyticspreferencesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/loganalytics/loganalyticspreferencesmanagement"
	loganalyticsresourcecategoriesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/loganalytics/loganalyticsresourcecategoriesmanagement"
	loganalyticsunprocesseddatabucketmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/loganalytics/loganalyticsunprocesseddatabucketmanagement"
	namespace "github.com/oracle/provider-oci/internal/controller/namespaced/loganalytics/namespace"
	namespaceassociation "github.com/oracle/provider-oci/internal/controller/namespaced/loganalytics/namespaceassociation"
	namespaceingesttimerule "github.com/oracle/provider-oci/internal/controller/namespaced/loganalytics/namespaceingesttimerule"
	namespaceingesttimerulesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/loganalytics/namespaceingesttimerulesmanagement"
	namespacelookup "github.com/oracle/provider-oci/internal/controller/namespaced/loganalytics/namespacelookup"
	namespacelookupsappenddatamanagement "github.com/oracle/provider-oci/internal/controller/namespaced/loganalytics/namespacelookupsappenddatamanagement"
	namespacelookupsupdatedatamanagement "github.com/oracle/provider-oci/internal/controller/namespaced/loganalytics/namespacelookupsupdatedatamanagement"
	namespacescheduledtask "github.com/oracle/provider-oci/internal/controller/namespaced/loganalytics/namespacescheduledtask"
	namespacestoragearchivalconfig "github.com/oracle/provider-oci/internal/controller/namespaced/loganalytics/namespacestoragearchivalconfig"
	namespacestorageenabledisablearchiving "github.com/oracle/provider-oci/internal/controller/namespaced/loganalytics/namespacestorageenabledisablearchiving"
	log "github.com/oracle/provider-oci/internal/controller/namespaced/logging/log"
	loggroup "github.com/oracle/provider-oci/internal/controller/namespaced/logging/loggroup"
	logsavedsearch "github.com/oracle/provider-oci/internal/controller/namespaced/logging/logsavedsearch"
	unifiedagentconfiguration "github.com/oracle/provider-oci/internal/controller/namespaced/logging/unifiedagentconfiguration"
	lustrefilesystem "github.com/oracle/provider-oci/internal/controller/namespaced/lustrefilestorage/lustrefilesystem"
	objectstoragelink "github.com/oracle/provider-oci/internal/controller/namespaced/lustrefilestorage/objectstoragelink"
	kafkacluster "github.com/oracle/provider-oci/internal/controller/namespaced/managedkafka/kafkacluster"
	kafkaclusteraddon "github.com/oracle/provider-oci/internal/controller/namespaced/managedkafka/kafkaclusteraddon"
	kafkaclusterconfig "github.com/oracle/provider-oci/internal/controller/namespaced/managedkafka/kafkaclusterconfig"
	kafkaclustersuperusersmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/managedkafka/kafkaclustersuperusersmanagement"
	managementagent "github.com/oracle/provider-oci/internal/controller/namespaced/managementagent/managementagent"
	managementagentdatasource "github.com/oracle/provider-oci/internal/controller/namespaced/managementagent/managementagentdatasource"
	managementagentinstallkey "github.com/oracle/provider-oci/internal/controller/namespaced/managementagent/managementagentinstallkey"
	namedcredential "github.com/oracle/provider-oci/internal/controller/namespaced/managementagent/namedcredential"
	managementdashboardsimport "github.com/oracle/provider-oci/internal/controller/namespaced/managementdashboard/managementdashboardsimport"
	managementsavedsearch "github.com/oracle/provider-oci/internal/controller/namespaced/managementdashboard/managementsavedsearch"
	acceptedagreement "github.com/oracle/provider-oci/internal/controller/namespaced/marketplace/acceptedagreement"
	listingpackageagreement "github.com/oracle/provider-oci/internal/controller/namespaced/marketplace/listingpackageagreement"
	marketplaceexternalattestedmetadata "github.com/oracle/provider-oci/internal/controller/namespaced/marketplace/marketplaceexternalattestedmetadata"
	publication "github.com/oracle/provider-oci/internal/controller/namespaced/marketplace/publication"
	mediaasset "github.com/oracle/provider-oci/internal/controller/namespaced/mediaservices/mediaasset"
	mediaworkflow "github.com/oracle/provider-oci/internal/controller/namespaced/mediaservices/mediaworkflow"
	mediaworkflowconfiguration "github.com/oracle/provider-oci/internal/controller/namespaced/mediaservices/mediaworkflowconfiguration"
	mediaworkflowjob "github.com/oracle/provider-oci/internal/controller/namespaced/mediaservices/mediaworkflowjob"
	streamcdnconfig "github.com/oracle/provider-oci/internal/controller/namespaced/mediaservices/streamcdnconfig"
	streamdistributionchannel "github.com/oracle/provider-oci/internal/controller/namespaced/mediaservices/streamdistributionchannel"
	streampackagingconfig "github.com/oracle/provider-oci/internal/controller/namespaced/mediaservices/streampackagingconfig"
	customtable "github.com/oracle/provider-oci/internal/controller/namespaced/meteringcomputation/customtable"
	query "github.com/oracle/provider-oci/internal/controller/namespaced/meteringcomputation/query"
	schedulemeteringcomputation "github.com/oracle/provider-oci/internal/controller/namespaced/meteringcomputation/schedule"
	usage "github.com/oracle/provider-oci/internal/controller/namespaced/meteringcomputation/usage"
	usagecarbonemission "github.com/oracle/provider-oci/internal/controller/namespaced/meteringcomputation/usagecarbonemission"
	usagecarbonemissionsquery "github.com/oracle/provider-oci/internal/controller/namespaced/meteringcomputation/usagecarbonemissionsquery"
	usagestatementemailrecipientsgroup "github.com/oracle/provider-oci/internal/controller/namespaced/meteringcomputation/usagestatementemailrecipientsgroup"
	alarm "github.com/oracle/provider-oci/internal/controller/namespaced/monitoring/alarm"
	alarmsuppression "github.com/oracle/provider-oci/internal/controller/namespaced/monitoring/alarmsuppression"
	capturefilter "github.com/oracle/provider-oci/internal/controller/namespaced/monitoring/capturefilter"
	vtap "github.com/oracle/provider-oci/internal/controller/namespaced/monitoring/vtap"
	bluegreendeployment "github.com/oracle/provider-oci/internal/controller/namespaced/mysql/bluegreendeployment"
	mysqlbackup "github.com/oracle/provider-oci/internal/controller/namespaced/mysql/mysqlbackup"
	mysqlchannel "github.com/oracle/provider-oci/internal/controller/namespaced/mysql/mysqlchannel"
	mysqlconfiguration "github.com/oracle/provider-oci/internal/controller/namespaced/mysql/mysqlconfiguration"
	mysqldbsystem "github.com/oracle/provider-oci/internal/controller/namespaced/mysql/mysqldbsystem"
	mysqlheatwavecluster "github.com/oracle/provider-oci/internal/controller/namespaced/mysql/mysqlheatwavecluster"
	mysqlreplica "github.com/oracle/provider-oci/internal/controller/namespaced/mysql/mysqlreplica"
	cpe "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/cpe"
	crossconnect "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/crossconnect"
	crossconnectgroup "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/crossconnectgroup"
	defaultdrgroutetable "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/defaultdrgroutetable"
	drg "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/drg"
	drgattachment "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/drgattachment"
	drgattachmentmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/drgattachmentmanagement"
	drgattachmentslist "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/drgattachmentslist"
	drgroutedistribution "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/drgroutedistribution"
	drgroutedistributionstatement "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/drgroutedistributionstatement"
	drgroutetable "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/drgroutetable"
	drgroutetablerouterule "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/drgroutetablerouterule"
	ipsec "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/ipsec"
	ipsecconnectiontunnelmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/ipsecconnectiontunnelmanagement"
	virtualcircuit "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/virtualcircuit"
	networkfirewall "github.com/oracle/provider-oci/internal/controller/namespaced/networkfirewall/networkfirewall"
	networkfirewallpolicy "github.com/oracle/provider-oci/internal/controller/namespaced/networkfirewall/networkfirewallpolicy"
	networkfirewallpolicyaddresslist "github.com/oracle/provider-oci/internal/controller/namespaced/networkfirewall/networkfirewallpolicyaddresslist"
	networkfirewallpolicyapplication "github.com/oracle/provider-oci/internal/controller/namespaced/networkfirewall/networkfirewallpolicyapplication"
	networkfirewallpolicyapplicationgroup "github.com/oracle/provider-oci/internal/controller/namespaced/networkfirewall/networkfirewallpolicyapplicationgroup"
	networkfirewallpolicydecryptionprofile "github.com/oracle/provider-oci/internal/controller/namespaced/networkfirewall/networkfirewallpolicydecryptionprofile"
	networkfirewallpolicydecryptionrule "github.com/oracle/provider-oci/internal/controller/namespaced/networkfirewall/networkfirewallpolicydecryptionrule"
	networkfirewallpolicymappedsecret "github.com/oracle/provider-oci/internal/controller/namespaced/networkfirewall/networkfirewallpolicymappedsecret"
	networkfirewallpolicynatrule "github.com/oracle/provider-oci/internal/controller/namespaced/networkfirewall/networkfirewallpolicynatrule"
	networkfirewallpolicysecurityrule "github.com/oracle/provider-oci/internal/controller/namespaced/networkfirewall/networkfirewallpolicysecurityrule"
	networkfirewallpolicyservice "github.com/oracle/provider-oci/internal/controller/namespaced/networkfirewall/networkfirewallpolicyservice"
	networkfirewallpolicytunnelinspectionrule "github.com/oracle/provider-oci/internal/controller/namespaced/networkfirewall/networkfirewallpolicytunnelinspectionrule"
	defaultdhcpoptions "github.com/oracle/provider-oci/internal/controller/namespaced/networking/defaultdhcpoptions"
	defaultroutetable "github.com/oracle/provider-oci/internal/controller/namespaced/networking/defaultroutetable"
	defaultsecuritylist "github.com/oracle/provider-oci/internal/controller/namespaced/networking/defaultsecuritylist"
	dhcpoptions "github.com/oracle/provider-oci/internal/controller/namespaced/networking/dhcpoptions"
	internetgateway "github.com/oracle/provider-oci/internal/controller/namespaced/networking/internetgateway"
	ipv6 "github.com/oracle/provider-oci/internal/controller/namespaced/networking/ipv6"
	localpeeringgateway "github.com/oracle/provider-oci/internal/controller/namespaced/networking/localpeeringgateway"
	natgateway "github.com/oracle/provider-oci/internal/controller/namespaced/networking/natgateway"
	networksecuritygroup "github.com/oracle/provider-oci/internal/controller/namespaced/networking/networksecuritygroup"
	networksecuritygroupsecurityrule "github.com/oracle/provider-oci/internal/controller/namespaced/networking/networksecuritygroupsecurityrule"
	privateip "github.com/oracle/provider-oci/internal/controller/namespaced/networking/privateip"
	publicip "github.com/oracle/provider-oci/internal/controller/namespaced/networking/publicip"
	publicippool "github.com/oracle/provider-oci/internal/controller/namespaced/networking/publicippool"
	publicippoolcapacity "github.com/oracle/provider-oci/internal/controller/namespaced/networking/publicippoolcapacity"
	remotepeeringconnection "github.com/oracle/provider-oci/internal/controller/namespaced/networking/remotepeeringconnection"
	routetable "github.com/oracle/provider-oci/internal/controller/namespaced/networking/routetable"
	routetableattachment "github.com/oracle/provider-oci/internal/controller/namespaced/networking/routetableattachment"
	securitylist "github.com/oracle/provider-oci/internal/controller/namespaced/networking/securitylist"
	servicegateway "github.com/oracle/provider-oci/internal/controller/namespaced/networking/servicegateway"
	subnet "github.com/oracle/provider-oci/internal/controller/namespaced/networking/subnet"
	vcn "github.com/oracle/provider-oci/internal/controller/namespaced/networking/vcn"
	vlan "github.com/oracle/provider-oci/internal/controller/namespaced/networking/vlan"
	vnicattachment "github.com/oracle/provider-oci/internal/controller/namespaced/networking/vnicattachment"
	backendnetworkloadbalancer "github.com/oracle/provider-oci/internal/controller/namespaced/networkloadbalancer/backend"
	backendsetnetworkloadbalancer "github.com/oracle/provider-oci/internal/controller/namespaced/networkloadbalancer/backendset"
	listenernetworkloadbalancer "github.com/oracle/provider-oci/internal/controller/namespaced/networkloadbalancer/listener"
	networkloadbalancer "github.com/oracle/provider-oci/internal/controller/namespaced/networkloadbalancer/networkloadbalancer"
	networkloadbalancersbackendsetsunified "github.com/oracle/provider-oci/internal/controller/namespaced/networkloadbalancer/networkloadbalancersbackendsetsunified"
	configurationnosql "github.com/oracle/provider-oci/internal/controller/namespaced/nosql/configuration"
	index "github.com/oracle/provider-oci/internal/controller/namespaced/nosql/index"
	table "github.com/oracle/provider-oci/internal/controller/namespaced/nosql/table"
	tablereplica "github.com/oracle/provider-oci/internal/controller/namespaced/nosql/tablereplica"
	bucket "github.com/oracle/provider-oci/internal/controller/namespaced/objectstorage/bucket"
	object "github.com/oracle/provider-oci/internal/controller/namespaced/objectstorage/object"
	objectlifecyclepolicy "github.com/oracle/provider-oci/internal/controller/namespaced/objectstorage/objectlifecyclepolicy"
	preauthrequest "github.com/oracle/provider-oci/internal/controller/namespaced/objectstorage/preauthrequest"
	privateendpointobjectstorage "github.com/oracle/provider-oci/internal/controller/namespaced/objectstorage/privateendpoint"
	replicationpolicy "github.com/oracle/provider-oci/internal/controller/namespaced/objectstorage/replicationpolicy"
	oceinstance "github.com/oracle/provider-oci/internal/controller/namespaced/oce/oceinstance"
	byol "github.com/oracle/provider-oci/internal/controller/namespaced/ocvp/byol"
	byolallocation "github.com/oracle/provider-oci/internal/controller/namespaced/ocvp/byolallocation"
	clusterocvp "github.com/oracle/provider-oci/internal/controller/namespaced/ocvp/cluster"
	datastore "github.com/oracle/provider-oci/internal/controller/namespaced/ocvp/datastore"
	datastorecluster "github.com/oracle/provider-oci/internal/controller/namespaced/ocvp/datastorecluster"
	esxihost "github.com/oracle/provider-oci/internal/controller/namespaced/ocvp/esxihost"
	managementappliance "github.com/oracle/provider-oci/internal/controller/namespaced/ocvp/managementappliance"
	sddc "github.com/oracle/provider-oci/internal/controller/namespaced/ocvp/sddc"
	odainstance "github.com/oracle/provider-oci/internal/controller/namespaced/oda/odainstance"
	odaprivateendpoint "github.com/oracle/provider-oci/internal/controller/namespaced/oda/odaprivateendpoint"
	odaprivateendpointattachment "github.com/oracle/provider-oci/internal/controller/namespaced/oda/odaprivateendpointattachment"
	odaprivateendpointscanproxy "github.com/oracle/provider-oci/internal/controller/namespaced/oda/odaprivateendpointscanproxy"
	notificationtopic "github.com/oracle/provider-oci/internal/controller/namespaced/ons/notificationtopic"
	subscription "github.com/oracle/provider-oci/internal/controller/namespaced/ons/subscription"
	opainstance "github.com/oracle/provider-oci/internal/controller/namespaced/opa/opainstance"
	opensearchcluster "github.com/oracle/provider-oci/internal/controller/namespaced/opensearch/opensearchcluster"
	opensearchclusterpipeline "github.com/oracle/provider-oci/internal/controller/namespaced/opensearch/opensearchclusterpipeline"
	operatorcontrol "github.com/oracle/provider-oci/internal/controller/namespaced/operatoraccesscontrol/operatorcontrol"
	operatorcontrolassignment "github.com/oracle/provider-oci/internal/controller/namespaced/operatoraccesscontrol/operatorcontrolassignment"
	awrhub "github.com/oracle/provider-oci/internal/controller/namespaced/opsi/awrhub"
	awrhubsource "github.com/oracle/provider-oci/internal/controller/namespaced/opsi/awrhubsource"
	awrhubsourceawrhubsourcesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/opsi/awrhubsourceawrhubsourcesmanagement"
	chargebackplan "github.com/oracle/provider-oci/internal/controller/namespaced/opsi/chargebackplan"
	databaseinsight "github.com/oracle/provider-oci/internal/controller/namespaced/opsi/databaseinsight"
	enterprisemanagerbridge "github.com/oracle/provider-oci/internal/controller/namespaced/opsi/enterprisemanagerbridge"
	exadatainsight "github.com/oracle/provider-oci/internal/controller/namespaced/opsi/exadatainsight"
	hostinsight "github.com/oracle/provider-oci/internal/controller/namespaced/opsi/hostinsight"
	newsreport "github.com/oracle/provider-oci/internal/controller/namespaced/opsi/newsreport"
	operationsinsightsprivateendpoint "github.com/oracle/provider-oci/internal/controller/namespaced/opsi/operationsinsightsprivateendpoint"
	operationsinsightswarehouse "github.com/oracle/provider-oci/internal/controller/namespaced/opsi/operationsinsightswarehouse"
	operationsinsightswarehousedownloadwarehousewallet "github.com/oracle/provider-oci/internal/controller/namespaced/opsi/operationsinsightswarehousedownloadwarehousewallet"
	operationsinsightswarehouserotatewarehousewallet "github.com/oracle/provider-oci/internal/controller/namespaced/opsi/operationsinsightswarehouserotatewarehousewallet"
	operationsinsightswarehouseuser "github.com/oracle/provider-oci/internal/controller/namespaced/opsi/operationsinsightswarehouseuser"
	opsiconfiguration "github.com/oracle/provider-oci/internal/controller/namespaced/opsi/opsiconfiguration"
	enrollmentstatus "github.com/oracle/provider-oci/internal/controller/namespaced/optimizer/enrollmentstatus"
	profile "github.com/oracle/provider-oci/internal/controller/namespaced/optimizer/profile"
	recommendation "github.com/oracle/provider-oci/internal/controller/namespaced/optimizer/recommendation"
	resourceaction "github.com/oracle/provider-oci/internal/controller/namespaced/optimizer/resourceaction"
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
	profileosmanagementhub "github.com/oracle/provider-oci/internal/controller/namespaced/osmanagementhub/profile"
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
	addressactionverification "github.com/oracle/provider-oci/internal/controller/namespaced/ospgateway/addressactionverification"
	subscriptionospgateway "github.com/oracle/provider-oci/internal/controller/namespaced/ospgateway/subscription"
	providerconfig "github.com/oracle/provider-oci/internal/controller/namespaced/providerconfig"
	privateserviceaccess "github.com/oracle/provider-oci/internal/controller/namespaced/psa/privateserviceaccess"
	psqlbackup "github.com/oracle/provider-oci/internal/controller/namespaced/psql/psqlbackup"
	psqlconfiguration "github.com/oracle/provider-oci/internal/controller/namespaced/psql/psqlconfiguration"
	psqldbsystem "github.com/oracle/provider-oci/internal/controller/namespaced/psql/psqldbsystem"
	consumergroup "github.com/oracle/provider-oci/internal/controller/namespaced/queue/consumergroup"
	queue "github.com/oracle/provider-oci/internal/controller/namespaced/queue/queue"
	protecteddatabase "github.com/oracle/provider-oci/internal/controller/namespaced/recovery/protecteddatabase"
	protectionpolicy "github.com/oracle/provider-oci/internal/controller/namespaced/recovery/protectionpolicy"
	recoveryservicesubnet "github.com/oracle/provider-oci/internal/controller/namespaced/recovery/recoveryservicesubnet"
	ocicachebackup "github.com/oracle/provider-oci/internal/controller/namespaced/redis/ocicachebackup"
	ocicachebackupexporttoobjectstorage "github.com/oracle/provider-oci/internal/controller/namespaced/redis/ocicachebackupexporttoobjectstorage"
	ocicacheconfigset "github.com/oracle/provider-oci/internal/controller/namespaced/redis/ocicacheconfigset"
	ocicacheconfigsetlistassociatedocicachecluster "github.com/oracle/provider-oci/internal/controller/namespaced/redis/ocicacheconfigsetlistassociatedocicachecluster"
	ocicacheuser "github.com/oracle/provider-oci/internal/controller/namespaced/redis/ocicacheuser"
	ocicacheusergetrediscluster "github.com/oracle/provider-oci/internal/controller/namespaced/redis/ocicacheusergetrediscluster"
	rediscluster "github.com/oracle/provider-oci/internal/controller/namespaced/redis/rediscluster"
	redisclusterattachocicacheuser "github.com/oracle/provider-oci/internal/controller/namespaced/redis/redisclusterattachocicacheuser"
	redisclustercreateidentitytoken "github.com/oracle/provider-oci/internal/controller/namespaced/redis/redisclustercreateidentitytoken"
	redisclusterdetachocicacheuser "github.com/oracle/provider-oci/internal/controller/namespaced/redis/redisclusterdetachocicacheuser"
	redisclustergetocicacheuser "github.com/oracle/provider-oci/internal/controller/namespaced/redis/redisclustergetocicacheuser"
	monitoredregion "github.com/oracle/provider-oci/internal/controller/namespaced/resourceanalytics/monitoredregion"
	resourceanalyticsinstance "github.com/oracle/provider-oci/internal/controller/namespaced/resourceanalytics/resourceanalyticsinstance"
	resourceanalyticsinstanceoacmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/resourceanalytics/resourceanalyticsinstanceoacmanagement"
	tenancyattachment "github.com/oracle/provider-oci/internal/controller/namespaced/resourceanalytics/tenancyattachment"
	privateendpointresourcemanager "github.com/oracle/provider-oci/internal/controller/namespaced/resourcemanager/privateendpoint"
	scheduleresourcescheduler "github.com/oracle/provider-oci/internal/controller/namespaced/resourcescheduler/schedule"
	serviceconnector "github.com/oracle/provider-oci/internal/controller/namespaced/sch/serviceconnector"
	securityattribute "github.com/oracle/provider-oci/internal/controller/namespaced/securityattribute/securityattribute"
	securityattributenamespace "github.com/oracle/provider-oci/internal/controller/namespaced/securityattribute/securityattributenamespace"
	subscriptionself "github.com/oracle/provider-oci/internal/controller/namespaced/self/subscription"
	privateapplication "github.com/oracle/provider-oci/internal/controller/namespaced/servicecatalog/privateapplication"
	servicecatalog "github.com/oracle/provider-oci/internal/controller/namespaced/servicecatalog/servicecatalog"
	servicecatalogassociation "github.com/oracle/provider-oci/internal/controller/namespaced/servicecatalog/servicecatalogassociation"
	baselineablemetric "github.com/oracle/provider-oci/internal/controller/namespaced/stackmonitoring/baselineablemetric"
	configstackmonitoring "github.com/oracle/provider-oci/internal/controller/namespaced/stackmonitoring/config"
	discoveryjobstackmonitoring "github.com/oracle/provider-oci/internal/controller/namespaced/stackmonitoring/discoveryjob"
	maintenancewindowstackmonitoring "github.com/oracle/provider-oci/internal/controller/namespaced/stackmonitoring/maintenancewindow"
	maintenancewindowsretryfailedoperation "github.com/oracle/provider-oci/internal/controller/namespaced/stackmonitoring/maintenancewindowsretryfailedoperation"
	maintenancewindowsstop "github.com/oracle/provider-oci/internal/controller/namespaced/stackmonitoring/maintenancewindowsstop"
	metricextension "github.com/oracle/provider-oci/internal/controller/namespaced/stackmonitoring/metricextension"
	metricextensionmetricextensionongivenresourcesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/stackmonitoring/metricextensionmetricextensionongivenresourcesmanagement"
	monitoredresource "github.com/oracle/provider-oci/internal/controller/namespaced/stackmonitoring/monitoredresource"
	monitoredresourcesassociatemonitoredresource "github.com/oracle/provider-oci/internal/controller/namespaced/stackmonitoring/monitoredresourcesassociatemonitoredresource"
	monitoredresourceslistmember "github.com/oracle/provider-oci/internal/controller/namespaced/stackmonitoring/monitoredresourceslistmember"
	monitoredresourcessearch "github.com/oracle/provider-oci/internal/controller/namespaced/stackmonitoring/monitoredresourcessearch"
	monitoredresourcessearchassociation "github.com/oracle/provider-oci/internal/controller/namespaced/stackmonitoring/monitoredresourcessearchassociation"
	monitoredresourcetask "github.com/oracle/provider-oci/internal/controller/namespaced/stackmonitoring/monitoredresourcetask"
	monitoredresourcetype "github.com/oracle/provider-oci/internal/controller/namespaced/stackmonitoring/monitoredresourcetype"
	monitoringtemplate "github.com/oracle/provider-oci/internal/controller/namespaced/stackmonitoring/monitoringtemplate"
	monitoringtemplatealarmcondition "github.com/oracle/provider-oci/internal/controller/namespaced/stackmonitoring/monitoringtemplatealarmcondition"
	monitoringtemplateongivenresourcesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/stackmonitoring/monitoringtemplateongivenresourcesmanagement"
	processset "github.com/oracle/provider-oci/internal/controller/namespaced/stackmonitoring/processset"
	connectharness "github.com/oracle/provider-oci/internal/controller/namespaced/streaming/connectharness"
	stream "github.com/oracle/provider-oci/internal/controller/namespaced/streaming/stream"
	streampool "github.com/oracle/provider-oci/internal/controller/namespaced/streaming/streampool"
	subscriptionmapping "github.com/oracle/provider-oci/internal/controller/namespaced/tenantmanagercontrolplane/subscriptionmapping"
	subscriptionredeemableuser "github.com/oracle/provider-oci/internal/controller/namespaced/usageproxy/subscriptionredeemableuser"
	secret "github.com/oracle/provider-oci/internal/controller/namespaced/vault/secret"
	vbsinstance "github.com/oracle/provider-oci/internal/controller/namespaced/vbsinst/vbsinstance"
	vbinstance "github.com/oracle/provider-oci/internal/controller/namespaced/visualbuilder/vbinstance"
	pathanalysi "github.com/oracle/provider-oci/internal/controller/namespaced/vnmonitoring/pathanalysi"
	containerscanrecipe "github.com/oracle/provider-oci/internal/controller/namespaced/vulnerabilityscanning/containerscanrecipe"
	containerscantarget "github.com/oracle/provider-oci/internal/controller/namespaced/vulnerabilityscanning/containerscantarget"
	hostscanrecipe "github.com/oracle/provider-oci/internal/controller/namespaced/vulnerabilityscanning/hostscanrecipe"
	hostscantarget "github.com/oracle/provider-oci/internal/controller/namespaced/vulnerabilityscanning/hostscantarget"
	webappacceleration "github.com/oracle/provider-oci/internal/controller/namespaced/waa/webappacceleration"
	webappaccelerationpolicy "github.com/oracle/provider-oci/internal/controller/namespaced/waa/webappaccelerationpolicy"
	addresslist "github.com/oracle/provider-oci/internal/controller/namespaced/waas/addresslist"
	certificatewaas "github.com/oracle/provider-oci/internal/controller/namespaced/waas/certificate"
	customprotectionrule "github.com/oracle/provider-oci/internal/controller/namespaced/waas/customprotectionrule"
	httpredirect "github.com/oracle/provider-oci/internal/controller/namespaced/waas/httpredirect"
	protectionrule "github.com/oracle/provider-oci/internal/controller/namespaced/waas/protectionrule"
	purgecache "github.com/oracle/provider-oci/internal/controller/namespaced/waas/purgecache"
	waaspolicy "github.com/oracle/provider-oci/internal/controller/namespaced/waas/waaspolicy"
	networkaddresslist "github.com/oracle/provider-oci/internal/controller/namespaced/waf/networkaddresslist"
	webappfirewall "github.com/oracle/provider-oci/internal/controller/namespaced/waf/webappfirewall"
	webappfirewallpolicy "github.com/oracle/provider-oci/internal/controller/namespaced/waf/webappfirewallpolicy"
	configurationzpr "github.com/oracle/provider-oci/internal/controller/namespaced/zpr/configuration"
	zprpolicy "github.com/oracle/provider-oci/internal/controller/namespaced/zpr/zprpolicy"
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
		defaultdrgroutetable.Setup,
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
		defaultdrgroutetable.SetupGated,
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

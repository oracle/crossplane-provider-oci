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
	documentmodel "github.com/oracle/provider-oci/internal/controller/ailanguage/documentmodel"
	documentprocessorjob "github.com/oracle/provider-oci/internal/controller/ailanguage/documentprocessorjob"
	documentproject "github.com/oracle/provider-oci/internal/controller/ailanguage/documentproject"
	languageendpoint "github.com/oracle/provider-oci/internal/controller/ailanguage/languageendpoint"
	languagemodel "github.com/oracle/provider-oci/internal/controller/ailanguage/languagemodel"
	languageproject "github.com/oracle/provider-oci/internal/controller/ailanguage/languageproject"
	visionmodel "github.com/oracle/provider-oci/internal/controller/ailanguage/visionmodel"
	visionproject "github.com/oracle/provider-oci/internal/controller/ailanguage/visionproject"
	analyticsinstance "github.com/oracle/provider-oci/internal/controller/analytics/analyticsinstance"
	instanceprivateaccesschannel "github.com/oracle/provider-oci/internal/controller/analytics/instanceprivateaccesschannel"
	instancevanityurl "github.com/oracle/provider-oci/internal/controller/analytics/instancevanityurl"
	serviceannouncementsubscription "github.com/oracle/provider-oci/internal/controller/announcements/serviceannouncementsubscription"
	serviceannouncementsubscriptionsactionschangecompartment "github.com/oracle/provider-oci/internal/controller/announcements/serviceannouncementsubscriptionsactionschangecompartment"
	serviceannouncementsubscriptionsfiltergroup "github.com/oracle/provider-oci/internal/controller/announcements/serviceannouncementsubscriptionsfiltergroup"
	platformapiplatforminstance "github.com/oracle/provider-oci/internal/controller/api/platformapiplatforminstance"
	privilegedapicontrol "github.com/oracle/provider-oci/internal/controller/apiaccesscontrol/privilegedapicontrol"
	privilegedapirequest "github.com/oracle/provider-oci/internal/controller/apiaccesscontrol/privilegedapirequest"
	api "github.com/oracle/provider-oci/internal/controller/apigateway/api"
	apigatewaycertificate "github.com/oracle/provider-oci/internal/controller/apigateway/apigatewaycertificate"
	apigatewaydeployment "github.com/oracle/provider-oci/internal/controller/apigateway/apigatewaydeployment"
	gateway "github.com/oracle/provider-oci/internal/controller/apigateway/gateway"
	subscriber "github.com/oracle/provider-oci/internal/controller/apigateway/subscriber"
	usageplan "github.com/oracle/provider-oci/internal/controller/apigateway/usageplan"
	apmdomain "github.com/oracle/provider-oci/internal/controller/apm/apmdomain"
	config "github.com/oracle/provider-oci/internal/controller/apm/config"
	syntheticsdedicatedvantagepoint "github.com/oracle/provider-oci/internal/controller/apm/syntheticsdedicatedvantagepoint"
	syntheticsmonitor "github.com/oracle/provider-oci/internal/controller/apm/syntheticsmonitor"
	syntheticsonpremisevantagepoint "github.com/oracle/provider-oci/internal/controller/apm/syntheticsonpremisevantagepoint"
	syntheticsonpremisevantagepointworker "github.com/oracle/provider-oci/internal/controller/apm/syntheticsonpremisevantagepointworker"
	syntheticsscript "github.com/oracle/provider-oci/internal/controller/apm/syntheticsscript"
	tracesscheduledquery "github.com/oracle/provider-oci/internal/controller/apm/tracesscheduledquery"
	controlmonitorpluginmanagement "github.com/oracle/provider-oci/internal/controller/applicationmanagement/controlmonitorpluginmanagement"
	containerconfiguration "github.com/oracle/provider-oci/internal/controller/artifacts/containerconfiguration"
	containerimagesignature "github.com/oracle/provider-oci/internal/controller/artifacts/containerimagesignature"
	containerrepository "github.com/oracle/provider-oci/internal/controller/artifacts/containerrepository"
	genericartifact "github.com/oracle/provider-oci/internal/controller/artifacts/genericartifact"
	repository "github.com/oracle/provider-oci/internal/controller/artifacts/repository"
	auditconfiguration "github.com/oracle/provider-oci/internal/controller/audit/auditconfiguration"
	autoscalingautoscalingconfiguration "github.com/oracle/provider-oci/internal/controller/autoscaling/autoscalingautoscalingconfiguration"
	bastionresource "github.com/oracle/provider-oci/internal/controller/bastion/bastionresource"
	session "github.com/oracle/provider-oci/internal/controller/bastion/session"
	bigdataserviceautoscalingconfiguration "github.com/oracle/provider-oci/internal/controller/bigdataservice/bigdataserviceautoscalingconfiguration"
	bigdataserviceinstance "github.com/oracle/provider-oci/internal/controller/bigdataservice/bigdataserviceinstance"
	capacityreport "github.com/oracle/provider-oci/internal/controller/bigdataservice/capacityreport"
	instanceapikey "github.com/oracle/provider-oci/internal/controller/bigdataservice/instanceapikey"
	instanceidentityconfiguration "github.com/oracle/provider-oci/internal/controller/bigdataservice/instanceidentityconfiguration"
	instancemetastoreconfig "github.com/oracle/provider-oci/internal/controller/bigdataservice/instancemetastoreconfig"
	instancenodebackup "github.com/oracle/provider-oci/internal/controller/bigdataservice/instancenodebackup"
	instancenodebackupconfiguration "github.com/oracle/provider-oci/internal/controller/bigdataservice/instancenodebackupconfiguration"
	instancenodereplaceconfiguration "github.com/oracle/provider-oci/internal/controller/bigdataservice/instancenodereplaceconfiguration"
	instanceoperationcertificatemanagement "github.com/oracle/provider-oci/internal/controller/bigdataservice/instanceoperationcertificatemanagement"
	instanceospatchaction "github.com/oracle/provider-oci/internal/controller/bigdataservice/instanceospatchaction"
	instancepatchaction "github.com/oracle/provider-oci/internal/controller/bigdataservice/instancepatchaction"
	instancereplacenodeaction "github.com/oracle/provider-oci/internal/controller/bigdataservice/instancereplacenodeaction"
	instanceresourceprincipalconfiguration "github.com/oracle/provider-oci/internal/controller/bigdataservice/instanceresourceprincipalconfiguration"
	instancesoftwareupdateaction "github.com/oracle/provider-oci/internal/controller/bigdataservice/instancesoftwareupdateaction"
	osn "github.com/oracle/provider-oci/internal/controller/blockchain/osn"
	peer "github.com/oracle/provider-oci/internal/controller/blockchain/peer"
	platform "github.com/oracle/provider-oci/internal/controller/blockchain/platform"
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
	budgetresource "github.com/oracle/provider-oci/internal/controller/budget/budgetresource"
	managementinternaloccmdemandsignal "github.com/oracle/provider-oci/internal/controller/capacity/managementinternaloccmdemandsignal"
	managementinternaloccmdemandsignaldelivery "github.com/oracle/provider-oci/internal/controller/capacity/managementinternaloccmdemandsignaldelivery"
	managementoccavailabilitycatalog "github.com/oracle/provider-oci/internal/controller/capacity/managementoccavailabilitycatalog"
	managementocccapacityrequest "github.com/oracle/provider-oci/internal/controller/capacity/managementocccapacityrequest"
	managementocccustomergroup "github.com/oracle/provider-oci/internal/controller/capacity/managementocccustomergroup"
	managementocccustomergroupocccustomer "github.com/oracle/provider-oci/internal/controller/capacity/managementocccustomergroupocccustomer"
	managementoccmdemandsignal "github.com/oracle/provider-oci/internal/controller/capacity/managementoccmdemandsignal"
	managementoccmdemandsignalitem "github.com/oracle/provider-oci/internal/controller/capacity/managementoccmdemandsignalitem"
	certificateauthority "github.com/oracle/provider-oci/internal/controller/certificatesmanagement/certificateauthority"
	managementcabundle "github.com/oracle/provider-oci/internal/controller/certificatesmanagement/managementcabundle"
	managementcertificate "github.com/oracle/provider-oci/internal/controller/certificatesmanagement/managementcertificate"
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
	placementgroupsclusterplacementgroup "github.com/oracle/provider-oci/internal/controller/cluster/placementgroupsclusterplacementgroup"
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
	cloudatcustomercccinfrastructure "github.com/oracle/provider-oci/internal/controller/computemanagement/cloudatcustomercccinfrastructure"
	cloudatcustomercccupgradeschedule "github.com/oracle/provider-oci/internal/controller/computemanagement/cloudatcustomercccupgradeschedule"
	instancescontainerinstance "github.com/oracle/provider-oci/internal/controller/container/instancescontainerinstance"
	addon "github.com/oracle/provider-oci/internal/controller/containerengine/addon"
	cluster "github.com/oracle/provider-oci/internal/controller/containerengine/cluster"
	clustercompletecredentialrotationmanagement "github.com/oracle/provider-oci/internal/controller/containerengine/clustercompletecredentialrotationmanagement"
	clusterstartcredentialrotationmanagement "github.com/oracle/provider-oci/internal/controller/containerengine/clusterstartcredentialrotationmanagement"
	clusterworkloadmapping "github.com/oracle/provider-oci/internal/controller/containerengine/clusterworkloadmapping"
	nodepool "github.com/oracle/provider-oci/internal/controller/containerengine/nodepool"
	virtualnodepool "github.com/oracle/provider-oci/internal/controller/containerengine/virtualnodepool"
	contentexperienceinstance "github.com/oracle/provider-oci/internal/controller/contentexperience/contentexperienceinstance"
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
	databasebackup "github.com/oracle/provider-oci/internal/controller/database/databasebackup"
	databasedbmfeatmgmt "github.com/oracle/provider-oci/internal/controller/database/databasedbmfeatmgmt"
	databasedbsystem "github.com/oracle/provider-oci/internal/controller/database/databasedbsystem"
	databasemigration "github.com/oracle/provider-oci/internal/controller/database/databasemigration"
	databaseresource "github.com/oracle/provider-oci/internal/controller/database/databaseresource"
	dbhome "github.com/oracle/provider-oci/internal/controller/database/dbhome"
	dbmgmtprivateendpoint "github.com/oracle/provider-oci/internal/controller/database/dbmgmtprivateendpoint"
	dbnode "github.com/oracle/provider-oci/internal/controller/database/dbnode"
	dbnodeconsoleconnection "github.com/oracle/provider-oci/internal/controller/database/dbnodeconsoleconnection"
	dbnodeconsolehistory "github.com/oracle/provider-oci/internal/controller/database/dbnodeconsolehistory"
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
	softwareimage "github.com/oracle/provider-oci/internal/controller/database/softwareimage"
	toolsdatabasetoolsconnection "github.com/oracle/provider-oci/internal/controller/database/toolsdatabasetoolsconnection"
	toolsdatabasetoolsprivateendpoint "github.com/oracle/provider-oci/internal/controller/database/toolsdatabasetoolsprivateendpoint"
	upgrade "github.com/oracle/provider-oci/internal/controller/database/upgrade"
	vmcluster "github.com/oracle/provider-oci/internal/controller/database/vmcluster"
	vmclusteraddvirtualmachine "github.com/oracle/provider-oci/internal/controller/database/vmclusteraddvirtualmachine"
	vmclusternetwork "github.com/oracle/provider-oci/internal/controller/database/vmclusternetwork"
	vmclusterremovevirtualmachine "github.com/oracle/provider-oci/internal/controller/database/vmclusterremovevirtualmachine"
	catalog "github.com/oracle/provider-oci/internal/controller/datacatalog/catalog"
	catalogprivateendpoint "github.com/oracle/provider-oci/internal/controller/datacatalog/catalogprivateendpoint"
	datacatalogconnection "github.com/oracle/provider-oci/internal/controller/datacatalog/datacatalogconnection"
	metastore "github.com/oracle/provider-oci/internal/controller/datacatalog/metastore"
	dataflowapplication "github.com/oracle/provider-oci/internal/controller/dataflow/dataflowapplication"
	dataflowprivateendpoint "github.com/oracle/provider-oci/internal/controller/dataflow/dataflowprivateendpoint"
	invokerun "github.com/oracle/provider-oci/internal/controller/dataflow/invokerun"
	pool "github.com/oracle/provider-oci/internal/controller/dataflow/pool"
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
	datascienceprivateendpoint "github.com/oracle/provider-oci/internal/controller/datascience/datascienceprivateendpoint"
	datascienceproject "github.com/oracle/provider-oci/internal/controller/datascience/datascienceproject"
	job "github.com/oracle/provider-oci/internal/controller/datascience/job"
	jobrun "github.com/oracle/provider-oci/internal/controller/datascience/jobrun"
	mlapplication "github.com/oracle/provider-oci/internal/controller/datascience/mlapplication"
	mlapplicationimplementation "github.com/oracle/provider-oci/internal/controller/datascience/mlapplicationimplementation"
	mlapplicationinstance "github.com/oracle/provider-oci/internal/controller/datascience/mlapplicationinstance"
	model "github.com/oracle/provider-oci/internal/controller/datascience/model"
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
	schedule "github.com/oracle/provider-oci/internal/controller/datascience/schedule"
	vulnerabilityscan "github.com/oracle/provider-oci/internal/controller/dblm/vulnerabilityscan"
	multicloudresourcediscovery "github.com/oracle/provider-oci/internal/controller/dbmulticloud/multicloudresourcediscovery"
	oracledbazureblobcontainer "github.com/oracle/provider-oci/internal/controller/dbmulticloud/oracledbazureblobcontainer"
	oracledbazureblobmount "github.com/oracle/provider-oci/internal/controller/dbmulticloud/oracledbazureblobmount"
	oracledbazureconnector "github.com/oracle/provider-oci/internal/controller/dbmulticloud/oracledbazureconnector"
	oracledbazurevault "github.com/oracle/provider-oci/internal/controller/dbmulticloud/oracledbazurevault"
	oracledbazurevaultassociation "github.com/oracle/provider-oci/internal/controller/dbmulticloud/oracledbazurevaultassociation"
	accesscontroldelegationcontrol "github.com/oracle/provider-oci/internal/controller/delegate/accesscontroldelegationcontrol"
	accesscontroldelegationsubscription "github.com/oracle/provider-oci/internal/controller/delegate/accesscontroldelegationsubscription"
	signaloccdemandsignal "github.com/oracle/provider-oci/internal/controller/demand/signaloccdemandsignal"
	desktoppool "github.com/oracle/provider-oci/internal/controller/desktops/desktoppool"
	buildpipeline "github.com/oracle/provider-oci/internal/controller/devops/buildpipeline"
	buildpipelinestage "github.com/oracle/provider-oci/internal/controller/devops/buildpipelinestage"
	buildrun "github.com/oracle/provider-oci/internal/controller/devops/buildrun"
	deployartifact "github.com/oracle/provider-oci/internal/controller/devops/deployartifact"
	deployenvironment "github.com/oracle/provider-oci/internal/controller/devops/deployenvironment"
	deploypipeline "github.com/oracle/provider-oci/internal/controller/devops/deploypipeline"
	deploystage "github.com/oracle/provider-oci/internal/controller/devops/deploystage"
	devopsconnection "github.com/oracle/provider-oci/internal/controller/devops/devopsconnection"
	devopsdeployment "github.com/oracle/provider-oci/internal/controller/devops/devopsdeployment"
	devopsproject "github.com/oracle/provider-oci/internal/controller/devops/devopsproject"
	devopsrepository "github.com/oracle/provider-oci/internal/controller/devops/devopsrepository"
	projectrepositorysetting "github.com/oracle/provider-oci/internal/controller/devops/projectrepositorysetting"
	repositorymirror "github.com/oracle/provider-oci/internal/controller/devops/repositorymirror"
	repositoryprotectedbranchmanagement "github.com/oracle/provider-oci/internal/controller/devops/repositoryprotectedbranchmanagement"
	repositoryref "github.com/oracle/provider-oci/internal/controller/devops/repositoryref"
	repositorysetting "github.com/oracle/provider-oci/internal/controller/devops/repositorysetting"
	trigger "github.com/oracle/provider-oci/internal/controller/devops/trigger"
	digitalassistantinstance "github.com/oracle/provider-oci/internal/controller/digitalassistant/digitalassistantinstance"
	digitalassistantprivateendpoint "github.com/oracle/provider-oci/internal/controller/digitalassistant/digitalassistantprivateendpoint"
	privateendpointattachment "github.com/oracle/provider-oci/internal/controller/digitalassistant/privateendpointattachment"
	privateendpointscanproxy "github.com/oracle/provider-oci/internal/controller/digitalassistant/privateendpointscanproxy"
	recoverydrplan "github.com/oracle/provider-oci/internal/controller/disasterrecovery/recoverydrplan"
	recoverydrplanexecution "github.com/oracle/provider-oci/internal/controller/disasterrecovery/recoverydrplanexecution"
	recoverydrprotectiongroup "github.com/oracle/provider-oci/internal/controller/disasterrecovery/recoverydrprotectiongroup"
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
	dkim "github.com/oracle/provider-oci/internal/controller/emaildataplane/dkim"
	emaildataplanedomain "github.com/oracle/provider-oci/internal/controller/emaildataplane/emaildataplanedomain"
	returnpath "github.com/oracle/provider-oci/internal/controller/emaildataplane/returnpath"
	sender "github.com/oracle/provider-oci/internal/controller/emaildataplane/sender"
	suppression "github.com/oracle/provider-oci/internal/controller/emaildataplane/suppression"
	rule "github.com/oracle/provider-oci/internal/controller/events/rule"
	export "github.com/oracle/provider-oci/internal/controller/filestorage/export"
	exportset "github.com/oracle/provider-oci/internal/controller/filestorage/exportset"
	filesystem "github.com/oracle/provider-oci/internal/controller/filestorage/filesystem"
	mounttarget "github.com/oracle/provider-oci/internal/controller/filestorage/mounttarget"
	replication "github.com/oracle/provider-oci/internal/controller/filestorage/replication"
	snapshot "github.com/oracle/provider-oci/internal/controller/filestorage/snapshot"
	storagefilesystemquotarule "github.com/oracle/provider-oci/internal/controller/filestorage/storagefilesystemquotarule"
	storagefilesystemsnapshotpolicy "github.com/oracle/provider-oci/internal/controller/filestorage/storagefilesystemsnapshotpolicy"
	storageoutboundconnector "github.com/oracle/provider-oci/internal/controller/filestorage/storageoutboundconnector"
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
	application "github.com/oracle/provider-oci/internal/controller/functions/application"
	function "github.com/oracle/provider-oci/internal/controller/functions/function"
	invokefunction "github.com/oracle/provider-oci/internal/controller/functions/invokefunction"
	appsfusionenvironment "github.com/oracle/provider-oci/internal/controller/fusionapps/appsfusionenvironment"
	appsfusionenvironmentadminuser "github.com/oracle/provider-oci/internal/controller/fusionapps/appsfusionenvironmentadminuser"
	appsfusionenvironmentfamily "github.com/oracle/provider-oci/internal/controller/fusionapps/appsfusionenvironmentfamily"
	appsfusionenvironmentrefreshactivity "github.com/oracle/provider-oci/internal/controller/fusionapps/appsfusionenvironmentrefreshactivity"
	appsfusionenvironmentserviceattachment "github.com/oracle/provider-oci/internal/controller/fusionapps/appsfusionenvironmentserviceattachment"
	aiagent "github.com/oracle/provider-oci/internal/controller/generativeai/aiagent"
	aiagentendpoint "github.com/oracle/provider-oci/internal/controller/generativeai/aiagentendpoint"
	aiagentknowledgebase "github.com/oracle/provider-oci/internal/controller/generativeai/aiagentknowledgebase"
	aiagenttool "github.com/oracle/provider-oci/internal/controller/generativeai/aiagenttool"
	aidedicatedaicluster "github.com/oracle/provider-oci/internal/controller/generativeai/aidedicatedaicluster"
	aiendpoint "github.com/oracle/provider-oci/internal/controller/generativeai/aiendpoint"
	aimodel "github.com/oracle/provider-oci/internal/controller/generativeai/aimodel"
	artifactscontentartifactbypath "github.com/oracle/provider-oci/internal/controller/generic/artifactscontentartifactbypath"
	distributeddatabaseprivateendpoint "github.com/oracle/provider-oci/internal/controller/globally/distributeddatabaseprivateendpoint"
	distributeddatabaseshardeddatabase "github.com/oracle/provider-oci/internal/controller/globally/distributeddatabaseshardeddatabase"
	gateconnection "github.com/oracle/provider-oci/internal/controller/goldengate/gateconnection"
	gateconnectionassignment "github.com/oracle/provider-oci/internal/controller/goldengate/gateconnectionassignment"
	gatedatabaseregistration "github.com/oracle/provider-oci/internal/controller/goldengate/gatedatabaseregistration"
	gatedeployment "github.com/oracle/provider-oci/internal/controller/goldengate/gatedeployment"
	gatedeploymentbackup "github.com/oracle/provider-oci/internal/controller/goldengate/gatedeploymentbackup"
	gatedeploymentcertificate "github.com/oracle/provider-oci/internal/controller/goldengate/gatedeploymentcertificate"
	gatepipeline "github.com/oracle/provider-oci/internal/controller/goldengate/gatepipeline"
	checkshttpprobe "github.com/oracle/provider-oci/internal/controller/healthchecks/checkshttpprobe"
	checkspingprobe "github.com/oracle/provider-oci/internal/controller/healthchecks/checkspingprobe"
	httpmonitor "github.com/oracle/provider-oci/internal/controller/healthchecks/httpmonitor"
	pingmonitor "github.com/oracle/provider-oci/internal/controller/healthchecks/pingmonitor"
	apikey "github.com/oracle/provider-oci/internal/controller/identity/apikey"
	authenticationpolicy "github.com/oracle/provider-oci/internal/controller/identity/authenticationpolicy"
	authtoken "github.com/oracle/provider-oci/internal/controller/identity/authtoken"
	compartment "github.com/oracle/provider-oci/internal/controller/identity/compartment"
	customersecretkey "github.com/oracle/provider-oci/internal/controller/identity/customersecretkey"
	dbcredential "github.com/oracle/provider-oci/internal/controller/identity/dbcredential"
	domainreplicationtoregion "github.com/oracle/provider-oci/internal/controller/identity/domainreplicationtoregion"
	domainsaccountrecoverysetting "github.com/oracle/provider-oci/internal/controller/identity/domainsaccountrecoverysetting"
	domainsapikey "github.com/oracle/provider-oci/internal/controller/identity/domainsapikey"
	domainsapp "github.com/oracle/provider-oci/internal/controller/identity/domainsapp"
	domainsapprole "github.com/oracle/provider-oci/internal/controller/identity/domainsapprole"
	domainsapprovalworkflow "github.com/oracle/provider-oci/internal/controller/identity/domainsapprovalworkflow"
	domainsapprovalworkflowassignment "github.com/oracle/provider-oci/internal/controller/identity/domainsapprovalworkflowassignment"
	domainsapprovalworkflowstep "github.com/oracle/provider-oci/internal/controller/identity/domainsapprovalworkflowstep"
	domainsauthenticationfactorsetting "github.com/oracle/provider-oci/internal/controller/identity/domainsauthenticationfactorsetting"
	domainsauthtoken "github.com/oracle/provider-oci/internal/controller/identity/domainsauthtoken"
	domainscloudgate "github.com/oracle/provider-oci/internal/controller/identity/domainscloudgate"
	domainscloudgatemapping "github.com/oracle/provider-oci/internal/controller/identity/domainscloudgatemapping"
	domainscloudgateserver "github.com/oracle/provider-oci/internal/controller/identity/domainscloudgateserver"
	domainscondition "github.com/oracle/provider-oci/internal/controller/identity/domainscondition"
	domainscustomersecretkey "github.com/oracle/provider-oci/internal/controller/identity/domainscustomersecretkey"
	domainsdynamicresourcegroup "github.com/oracle/provider-oci/internal/controller/identity/domainsdynamicresourcegroup"
	domainsgrant "github.com/oracle/provider-oci/internal/controller/identity/domainsgrant"
	domainsgroup "github.com/oracle/provider-oci/internal/controller/identity/domainsgroup"
	domainsidentitypropagationtrust "github.com/oracle/provider-oci/internal/controller/identity/domainsidentitypropagationtrust"
	domainsidentityprovider "github.com/oracle/provider-oci/internal/controller/identity/domainsidentityprovider"
	domainsidentitysetting "github.com/oracle/provider-oci/internal/controller/identity/domainsidentitysetting"
	domainskmsisetting "github.com/oracle/provider-oci/internal/controller/identity/domainskmsisetting"
	domainsmyapikey "github.com/oracle/provider-oci/internal/controller/identity/domainsmyapikey"
	domainsmyauthtoken "github.com/oracle/provider-oci/internal/controller/identity/domainsmyauthtoken"
	domainsmycustomersecretkey "github.com/oracle/provider-oci/internal/controller/identity/domainsmycustomersecretkey"
	domainsmyoauth2clientcredential "github.com/oracle/provider-oci/internal/controller/identity/domainsmyoauth2clientcredential"
	domainsmyrequest "github.com/oracle/provider-oci/internal/controller/identity/domainsmyrequest"
	domainsmysmtpcredential "github.com/oracle/provider-oci/internal/controller/identity/domainsmysmtpcredential"
	domainsmysupportaccount "github.com/oracle/provider-oci/internal/controller/identity/domainsmysupportaccount"
	domainsmyuserdbcredential "github.com/oracle/provider-oci/internal/controller/identity/domainsmyuserdbcredential"
	domainsnetworkperimeter "github.com/oracle/provider-oci/internal/controller/identity/domainsnetworkperimeter"
	domainsnotificationsetting "github.com/oracle/provider-oci/internal/controller/identity/domainsnotificationsetting"
	domainsoauth2clientcredential "github.com/oracle/provider-oci/internal/controller/identity/domainsoauth2clientcredential"
	domainsoauthclientcertificate "github.com/oracle/provider-oci/internal/controller/identity/domainsoauthclientcertificate"
	domainsoauthpartnercertificate "github.com/oracle/provider-oci/internal/controller/identity/domainsoauthpartnercertificate"
	domainspasswordpolicy "github.com/oracle/provider-oci/internal/controller/identity/domainspasswordpolicy"
	domainspolicy "github.com/oracle/provider-oci/internal/controller/identity/domainspolicy"
	domainsrule "github.com/oracle/provider-oci/internal/controller/identity/domainsrule"
	domainssecurityquestion "github.com/oracle/provider-oci/internal/controller/identity/domainssecurityquestion"
	domainssecurityquestionsetting "github.com/oracle/provider-oci/internal/controller/identity/domainssecurityquestionsetting"
	domainsselfregistrationprofile "github.com/oracle/provider-oci/internal/controller/identity/domainsselfregistrationprofile"
	domainssetting "github.com/oracle/provider-oci/internal/controller/identity/domainssetting"
	domainssmtpcredential "github.com/oracle/provider-oci/internal/controller/identity/domainssmtpcredential"
	domainssocialidentityprovider "github.com/oracle/provider-oci/internal/controller/identity/domainssocialidentityprovider"
	domainsuser "github.com/oracle/provider-oci/internal/controller/identity/domainsuser"
	domainsuserdbcredential "github.com/oracle/provider-oci/internal/controller/identity/domainsuserdbcredential"
	dynamicgroup "github.com/oracle/provider-oci/internal/controller/identity/dynamicgroup"
	group "github.com/oracle/provider-oci/internal/controller/identity/group"
	identitydomain "github.com/oracle/provider-oci/internal/controller/identity/identitydomain"
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
	integrationinstance "github.com/oracle/provider-oci/internal/controller/integration/integrationinstance"
	oraclemanagedcustomendpoint "github.com/oracle/provider-oci/internal/controller/integration/oraclemanagedcustomendpoint"
	privateendpointoutboundconnection "github.com/oracle/provider-oci/internal/controller/integration/privateendpointoutboundconnection"
	fleet "github.com/oracle/provider-oci/internal/controller/jms/fleet"
	fleetadvancedfeatureconfiguration "github.com/oracle/provider-oci/internal/controller/jms/fleetadvancedfeatureconfiguration"
	javadownloadsjavadownloadreport "github.com/oracle/provider-oci/internal/controller/jms/javadownloadsjavadownloadreport"
	javadownloadsjavadownloadtoken "github.com/oracle/provider-oci/internal/controller/jms/javadownloadsjavadownloadtoken"
	javadownloadsjavalicenseacceptancerecord "github.com/oracle/provider-oci/internal/controller/jms/javadownloadsjavalicenseacceptancerecord"
	plugin "github.com/oracle/provider-oci/internal/controller/jms/plugin"
	ekmsprivateendpoint "github.com/oracle/provider-oci/internal/controller/kms/ekmsprivateendpoint"
	encrypteddata "github.com/oracle/provider-oci/internal/controller/kms/encrypteddata"
	generatedkey "github.com/oracle/provider-oci/internal/controller/kms/generatedkey"
	key "github.com/oracle/provider-oci/internal/controller/kms/key"
	keyversion "github.com/oracle/provider-oci/internal/controller/kms/keyversion"
	sign "github.com/oracle/provider-oci/internal/controller/kms/sign"
	vault "github.com/oracle/provider-oci/internal/controller/kms/vault"
	vaultreplication "github.com/oracle/provider-oci/internal/controller/kms/vaultreplication"
	verify "github.com/oracle/provider-oci/internal/controller/kms/verify"
	managerconfiguration "github.com/oracle/provider-oci/internal/controller/license/managerconfiguration"
	managerlicenserecord "github.com/oracle/provider-oci/internal/controller/license/managerlicenserecord"
	managerproductlicense "github.com/oracle/provider-oci/internal/controller/license/managerproductlicense"
	quota "github.com/oracle/provider-oci/internal/controller/limits/quota"
	backend "github.com/oracle/provider-oci/internal/controller/loadbalancer/backend"
	backendset "github.com/oracle/provider-oci/internal/controller/loadbalancer/backendset"
	balancer "github.com/oracle/provider-oci/internal/controller/loadbalancer/balancer"
	balancerbackendset "github.com/oracle/provider-oci/internal/controller/loadbalancer/balancerbackendset"
	certificate "github.com/oracle/provider-oci/internal/controller/loadbalancer/certificate"
	lbhostname "github.com/oracle/provider-oci/internal/controller/loadbalancer/lbhostname"
	listener "github.com/oracle/provider-oci/internal/controller/loadbalancer/listener"
	loadbalancer "github.com/oracle/provider-oci/internal/controller/loadbalancer/loadbalancer"
	pathrouteset "github.com/oracle/provider-oci/internal/controller/loadbalancer/pathrouteset"
	routingpolicy "github.com/oracle/provider-oci/internal/controller/loadbalancer/routingpolicy"
	ruleset "github.com/oracle/provider-oci/internal/controller/loadbalancer/ruleset"
	sslciphersuite "github.com/oracle/provider-oci/internal/controller/loadbalancer/sslciphersuite"
	analyticsloganalyticsentity "github.com/oracle/provider-oci/internal/controller/log/analyticsloganalyticsentity"
	analyticsloganalyticsentitytype "github.com/oracle/provider-oci/internal/controller/log/analyticsloganalyticsentitytype"
	analyticsloganalyticsimportcustomcontent "github.com/oracle/provider-oci/internal/controller/log/analyticsloganalyticsimportcustomcontent"
	analyticsloganalyticsloggroup "github.com/oracle/provider-oci/internal/controller/log/analyticsloganalyticsloggroup"
	analyticsloganalyticsobjectcollectionrule "github.com/oracle/provider-oci/internal/controller/log/analyticsloganalyticsobjectcollectionrule"
	analyticsloganalyticspreferencesmanagement "github.com/oracle/provider-oci/internal/controller/log/analyticsloganalyticspreferencesmanagement"
	analyticsloganalyticsresourcecategoriesmanagement "github.com/oracle/provider-oci/internal/controller/log/analyticsloganalyticsresourcecategoriesmanagement"
	analyticsnamespace "github.com/oracle/provider-oci/internal/controller/log/analyticsnamespace"
	analyticsnamespaceingesttimerule "github.com/oracle/provider-oci/internal/controller/log/analyticsnamespaceingesttimerule"
	analyticsnamespaceingesttimerulesmanagement "github.com/oracle/provider-oci/internal/controller/log/analyticsnamespaceingesttimerulesmanagement"
	analyticsnamespacelookup "github.com/oracle/provider-oci/internal/controller/log/analyticsnamespacelookup"
	analyticsnamespacescheduledtask "github.com/oracle/provider-oci/internal/controller/log/analyticsnamespacescheduledtask"
	analyticsnamespacestoragearchivalconfig "github.com/oracle/provider-oci/internal/controller/log/analyticsnamespacestoragearchivalconfig"
	analyticsnamespacestorageenabledisablearchiving "github.com/oracle/provider-oci/internal/controller/log/analyticsnamespacestorageenabledisablearchiving"
	log "github.com/oracle/provider-oci/internal/controller/logging/log"
	loggroup "github.com/oracle/provider-oci/internal/controller/logging/loggroup"
	logsavedsearch "github.com/oracle/provider-oci/internal/controller/logging/logsavedsearch"
	unifiedagentconfiguration "github.com/oracle/provider-oci/internal/controller/logging/unifiedagentconfiguration"
	filestoragelustrefilesystem "github.com/oracle/provider-oci/internal/controller/lustre/filestoragelustrefilesystem"
	agentmanagementagent "github.com/oracle/provider-oci/internal/controller/management/agentmanagementagent"
	agentmanagementagentinstallkey "github.com/oracle/provider-oci/internal/controller/management/agentmanagementagentinstallkey"
	agentnamedcredential "github.com/oracle/provider-oci/internal/controller/management/agentnamedcredential"
	dashboardmanagementdashboardsimport "github.com/oracle/provider-oci/internal/controller/management/dashboardmanagementdashboardsimport"
	acceptedagreement "github.com/oracle/provider-oci/internal/controller/marketplace/acceptedagreement"
	listingpackageagreement "github.com/oracle/provider-oci/internal/controller/marketplace/listingpackageagreement"
	publication "github.com/oracle/provider-oci/internal/controller/marketplace/publication"
	servicesmediaasset "github.com/oracle/provider-oci/internal/controller/mediaservices/servicesmediaasset"
	servicesmediaworkflow "github.com/oracle/provider-oci/internal/controller/mediaservices/servicesmediaworkflow"
	servicesmediaworkflowconfiguration "github.com/oracle/provider-oci/internal/controller/mediaservices/servicesmediaworkflowconfiguration"
	servicesmediaworkflowjob "github.com/oracle/provider-oci/internal/controller/mediaservices/servicesmediaworkflowjob"
	servicesstreamcdnconfig "github.com/oracle/provider-oci/internal/controller/mediaservices/servicesstreamcdnconfig"
	servicesstreamdistributionchannel "github.com/oracle/provider-oci/internal/controller/mediaservices/servicesstreamdistributionchannel"
	servicesstreampackagingconfig "github.com/oracle/provider-oci/internal/controller/mediaservices/servicesstreampackagingconfig"
	computationcustomtable "github.com/oracle/provider-oci/internal/controller/meteringcomputation/computationcustomtable"
	computationquery "github.com/oracle/provider-oci/internal/controller/meteringcomputation/computationquery"
	computationschedule "github.com/oracle/provider-oci/internal/controller/meteringcomputation/computationschedule"
	computationusage "github.com/oracle/provider-oci/internal/controller/meteringcomputation/computationusage"
	computationusagecarbonemission "github.com/oracle/provider-oci/internal/controller/meteringcomputation/computationusagecarbonemission"
	computationusagecarbonemissionsquery "github.com/oracle/provider-oci/internal/controller/meteringcomputation/computationusagecarbonemissionsquery"
	computationusagestatementemailrecipientsgroup "github.com/oracle/provider-oci/internal/controller/meteringcomputation/computationusagestatementemailrecipientsgroup"
	alarm "github.com/oracle/provider-oci/internal/controller/monitoring/alarm"
	alarmsuppression "github.com/oracle/provider-oci/internal/controller/monitoring/alarmsuppression"
	capturefilter "github.com/oracle/provider-oci/internal/controller/monitoring/capturefilter"
	vtap "github.com/oracle/provider-oci/internal/controller/monitoring/vtap"
	channel "github.com/oracle/provider-oci/internal/controller/mysql/channel"
	heatwavecluster "github.com/oracle/provider-oci/internal/controller/mysql/heatwavecluster"
	mysqlbackup "github.com/oracle/provider-oci/internal/controller/mysql/mysqlbackup"
	mysqlconfiguration "github.com/oracle/provider-oci/internal/controller/mysql/mysqlconfiguration"
	mysqldbsystem "github.com/oracle/provider-oci/internal/controller/mysql/mysqldbsystem"
	replica "github.com/oracle/provider-oci/internal/controller/mysql/replica"
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
	defaultdhcpoptions "github.com/oracle/provider-oci/internal/controller/networking/defaultdhcpoptions"
	defaultroutetable "github.com/oracle/provider-oci/internal/controller/networking/defaultroutetable"
	defaultsecuritylist "github.com/oracle/provider-oci/internal/controller/networking/defaultsecuritylist"
	dhcpoptions "github.com/oracle/provider-oci/internal/controller/networking/dhcpoptions"
	firewallnetworkfirewallpolicyaddresslist "github.com/oracle/provider-oci/internal/controller/networking/firewallnetworkfirewallpolicyaddresslist"
	firewallnetworkfirewallpolicyapplication "github.com/oracle/provider-oci/internal/controller/networking/firewallnetworkfirewallpolicyapplication"
	firewallnetworkfirewallpolicyapplicationgroup "github.com/oracle/provider-oci/internal/controller/networking/firewallnetworkfirewallpolicyapplicationgroup"
	firewallnetworkfirewallpolicydecryptionprofile "github.com/oracle/provider-oci/internal/controller/networking/firewallnetworkfirewallpolicydecryptionprofile"
	firewallnetworkfirewallpolicydecryptionrule "github.com/oracle/provider-oci/internal/controller/networking/firewallnetworkfirewallpolicydecryptionrule"
	firewallnetworkfirewallpolicymappedsecret "github.com/oracle/provider-oci/internal/controller/networking/firewallnetworkfirewallpolicymappedsecret"
	firewallnetworkfirewallpolicynatrule "github.com/oracle/provider-oci/internal/controller/networking/firewallnetworkfirewallpolicynatrule"
	firewallnetworkfirewallpolicysecurityrule "github.com/oracle/provider-oci/internal/controller/networking/firewallnetworkfirewallpolicysecurityrule"
	firewallnetworkfirewallpolicyservice "github.com/oracle/provider-oci/internal/controller/networking/firewallnetworkfirewallpolicyservice"
	firewallnetworkfirewallpolicytunnelinspectionrule "github.com/oracle/provider-oci/internal/controller/networking/firewallnetworkfirewallpolicytunnelinspectionrule"
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
	index "github.com/oracle/provider-oci/internal/controller/nosql/index"
	nosqlconfiguration "github.com/oracle/provider-oci/internal/controller/nosql/nosqlconfiguration"
	table "github.com/oracle/provider-oci/internal/controller/nosql/table"
	tablereplica "github.com/oracle/provider-oci/internal/controller/nosql/tablereplica"
	bucket "github.com/oracle/provider-oci/internal/controller/objectstorage/bucket"
	namespacemetadata "github.com/oracle/provider-oci/internal/controller/objectstorage/namespacemetadata"
	object "github.com/oracle/provider-oci/internal/controller/objectstorage/object"
	objectlifecyclepolicy "github.com/oracle/provider-oci/internal/controller/objectstorage/objectlifecyclepolicy"
	objectstorageprivateendpoint "github.com/oracle/provider-oci/internal/controller/objectstorage/objectstorageprivateendpoint"
	preauthrequest "github.com/oracle/provider-oci/internal/controller/objectstorage/preauthrequest"
	replicationpolicy "github.com/oracle/provider-oci/internal/controller/objectstorage/replicationpolicy"
	esxihost "github.com/oracle/provider-oci/internal/controller/ocvs/esxihost"
	ocvscluster "github.com/oracle/provider-oci/internal/controller/ocvs/ocvscluster"
	sddc "github.com/oracle/provider-oci/internal/controller/ocvs/sddc"
	notificationtopic "github.com/oracle/provider-oci/internal/controller/ons/notificationtopic"
	subscription "github.com/oracle/provider-oci/internal/controller/ons/subscription"
	opainstance "github.com/oracle/provider-oci/internal/controller/opa/opainstance"
	clusterpipeline "github.com/oracle/provider-oci/internal/controller/opensearch/clusterpipeline"
	opensearchcluster "github.com/oracle/provider-oci/internal/controller/opensearch/opensearchcluster"
	accesscontroloperatorcontrol "github.com/oracle/provider-oci/internal/controller/operator/accesscontroloperatorcontrol"
	accesscontroloperatorcontrolassignment "github.com/oracle/provider-oci/internal/controller/operator/accesscontroloperatorcontrolassignment"
	awrhub "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/awrhub"
	awrhubsource "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/awrhubsource"
	awrhubsourceawrhubsourcesmanagement "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/awrhubsourceawrhubsourcesmanagement"
	databaseinsight "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/databaseinsight"
	enterprisemanagerbridge "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/enterprisemanagerbridge"
	exadatainsight "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/exadatainsight"
	hostinsight "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/hostinsight"
	newsreport "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/newsreport"
	operationsinsightsprivateendpoint "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/operationsinsightsprivateendpoint"
	operationsinsightswarehouse "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/operationsinsightswarehouse"
	operationsinsightswarehousedownloadwarehousewallet "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/operationsinsightswarehousedownloadwarehousewallet"
	operationsinsightswarehouserotatewarehousewallet "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/operationsinsightswarehouserotatewarehousewallet"
	operationsinsightswarehouseuser "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/operationsinsightswarehouseuser"
	operatoraccesscontrolconfiguration "github.com/oracle/provider-oci/internal/controller/operatoraccesscontrol/operatoraccesscontrolconfiguration"
	managementhubevent "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubevent"
	managementhublifecycleenvironment "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhublifecycleenvironment"
	managementhublifecyclestageattachmanagedinstancesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhublifecyclestageattachmanagedinstancesmanagement"
	managementhublifecyclestagedetachmanagedinstancesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhublifecyclestagedetachmanagedinstancesmanagement"
	managementhublifecyclestagepromotesoftwaresourcemanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhublifecyclestagepromotesoftwaresourcemanagement"
	managementhublifecyclestagerebootmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhublifecyclestagerebootmanagement"
	managementhubmanagedinstance "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstance"
	managementhubmanagedinstanceattachprofilemanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstanceattachprofilemanagement"
	managementhubmanagedinstancedetachprofilemanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstancedetachprofilemanagement"
	managementhubmanagedinstancegroup "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstancegroup"
	managementhubmanagedinstancegroupattachmanagedinstancesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstancegroupattachmanagedinstancesmanagement"
	managementhubmanagedinstancegroupattachsoftwaresourcesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstancegroupattachsoftwaresourcesmanagement"
	managementhubmanagedinstancegroupdetachmanagedinstancesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstancegroupdetachmanagedinstancesmanagement"
	managementhubmanagedinstancegroupdetachsoftwaresourcesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstancegroupdetachsoftwaresourcesmanagement"
	managementhubmanagedinstancegroupinstallpackagesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstancegroupinstallpackagesmanagement"
	managementhubmanagedinstancegroupinstallwindowsupdatesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstancegroupinstallwindowsupdatesmanagement"
	managementhubmanagedinstancegroupmanagemodulestreamsmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstancegroupmanagemodulestreamsmanagement"
	managementhubmanagedinstancegrouprebootmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstancegrouprebootmanagement"
	managementhubmanagedinstancegroupremovepackagesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstancegroupremovepackagesmanagement"
	managementhubmanagedinstancegroupupdateallpackagesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstancegroupupdateallpackagesmanagement"
	managementhubmanagedinstanceinstallwindowsupdatesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstanceinstallwindowsupdatesmanagement"
	managementhubmanagedinstancerebootmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstancerebootmanagement"
	managementhubmanagedinstanceupdatepackagesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagedinstanceupdatepackagesmanagement"
	managementhubmanagementstation "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagementstation"
	managementhubmanagementstationassociatemanagedinstancesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagementstationassociatemanagedinstancesmanagement"
	managementhubmanagementstationmirrorsynchronizemanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagementstationmirrorsynchronizemanagement"
	managementhubmanagementstationrefreshmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagementstationrefreshmanagement"
	managementhubmanagementstationsynchronizemirrorsmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubmanagementstationsynchronizemirrorsmanagement"
	managementhubprofile "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubprofile"
	managementhubprofileattachlifecyclestagemanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubprofileattachlifecyclestagemanagement"
	managementhubprofileattachmanagedinstancegroupmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubprofileattachmanagedinstancegroupmanagement"
	managementhubprofileattachmanagementstationmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubprofileattachmanagementstationmanagement"
	managementhubprofileattachsoftwaresourcesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubprofileattachsoftwaresourcesmanagement"
	managementhubprofiledetachsoftwaresourcesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubprofiledetachsoftwaresourcesmanagement"
	managementhubscheduledjob "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubscheduledjob"
	managementhubsoftwaresource "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubsoftwaresource"
	managementhubsoftwaresourceaddpackagesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubsoftwaresourceaddpackagesmanagement"
	managementhubsoftwaresourcechangeavailabilitymanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubsoftwaresourcechangeavailabilitymanagement"
	managementhubsoftwaresourcegeneratemetadatamanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubsoftwaresourcegeneratemetadatamanagement"
	managementhubsoftwaresourcemanifest "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubsoftwaresourcemanifest"
	managementhubsoftwaresourceremovepackagesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubsoftwaresourceremovepackagesmanagement"
	managementhubsoftwaresourcereplacepackagesmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubsoftwaresourcereplacepackagesmanagement"
	managementhubworkrequestrerunmanagement "github.com/oracle/provider-oci/internal/controller/osmanagement/managementhubworkrequestrerunmanagement"
	gatewayaddressactionverification "github.com/oracle/provider-oci/internal/controller/osp/gatewayaddressactionverification"
	gatewaysubscription "github.com/oracle/provider-oci/internal/controller/osp/gatewaysubscription"
	providerconfig "github.com/oracle/provider-oci/internal/controller/providerconfig"
	psqlbackup "github.com/oracle/provider-oci/internal/controller/psql/psqlbackup"
	psqlconfiguration "github.com/oracle/provider-oci/internal/controller/psql/psqlconfiguration"
	psqldbsystem "github.com/oracle/provider-oci/internal/controller/psql/psqldbsystem"
	queueresource "github.com/oracle/provider-oci/internal/controller/queue/queueresource"
	protecteddatabase "github.com/oracle/provider-oci/internal/controller/recovery/protecteddatabase"
	protectionpolicy "github.com/oracle/provider-oci/internal/controller/recovery/protectionpolicy"
	servicesubnet "github.com/oracle/provider-oci/internal/controller/recovery/servicesubnet"
	clusterattachocicacheuser "github.com/oracle/provider-oci/internal/controller/redis/clusterattachocicacheuser"
	clustercreateidentitytoken "github.com/oracle/provider-oci/internal/controller/redis/clustercreateidentitytoken"
	clusterdetachocicacheuser "github.com/oracle/provider-oci/internal/controller/redis/clusterdetachocicacheuser"
	clustergetocicacheuser "github.com/oracle/provider-oci/internal/controller/redis/clustergetocicacheuser"
	ocicacheuser "github.com/oracle/provider-oci/internal/controller/redis/ocicacheuser"
	ocicacheusergetrediscluster "github.com/oracle/provider-oci/internal/controller/redis/ocicacheusergetrediscluster"
	rediscluster "github.com/oracle/provider-oci/internal/controller/redis/rediscluster"
	schedulerschedule "github.com/oracle/provider-oci/internal/controller/resource/schedulerschedule"
	resourcemanagerprivateendpoint "github.com/oracle/provider-oci/internal/controller/resourcemanager/resourcemanagerprivateendpoint"
	serviceconnector "github.com/oracle/provider-oci/internal/controller/sch/serviceconnector"
	attributesecurityattribute "github.com/oracle/provider-oci/internal/controller/securityattribute/attributesecurityattribute"
	attributesecurityattributenamespace "github.com/oracle/provider-oci/internal/controller/securityattribute/attributesecurityattributenamespace"
	catalogprivateapplication "github.com/oracle/provider-oci/internal/controller/service/catalogprivateapplication"
	catalogservicecatalog "github.com/oracle/provider-oci/internal/controller/service/catalogservicecatalog"
	catalogservicecatalogassociation "github.com/oracle/provider-oci/internal/controller/service/catalogservicecatalogassociation"
	monitoringbaselineablemetric "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringbaselineablemetric"
	monitoringconfig "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringconfig"
	monitoringdiscoveryjob "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringdiscoveryjob"
	monitoringmaintenancewindow "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringmaintenancewindow"
	monitoringmaintenancewindowsretryfailedoperation "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringmaintenancewindowsretryfailedoperation"
	monitoringmaintenancewindowsstop "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringmaintenancewindowsstop"
	monitoringmetricextension "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringmetricextension"
	monitoringmetricextensionmetricextensionongivenresourcesmanagement "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringmetricextensionmetricextensionongivenresourcesmanagement"
	monitoringmonitoredresource "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringmonitoredresource"
	monitoringmonitoredresourcesassociatemonitoredresource "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringmonitoredresourcesassociatemonitoredresource"
	monitoringmonitoredresourceslistmember "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringmonitoredresourceslistmember"
	monitoringmonitoredresourcessearch "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringmonitoredresourcessearch"
	monitoringmonitoredresourcessearchassociation "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringmonitoredresourcessearchassociation"
	monitoringmonitoredresourcetask "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringmonitoredresourcetask"
	monitoringmonitoredresourcetype "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringmonitoredresourcetype"
	monitoringprocessset "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringprocessset"
	monitoringtemplate "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringtemplate"
	monitoringtemplatealarmcondition "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringtemplatealarmcondition"
	monitoringtemplatemonitoringtemplateongivenresourcesmanagement "github.com/oracle/provider-oci/internal/controller/stackmonitoring/monitoringtemplatemonitoringtemplateongivenresourcesmanagement"
	connectharness "github.com/oracle/provider-oci/internal/controller/streaming/connectharness"
	stream "github.com/oracle/provider-oci/internal/controller/streaming/stream"
	streampool "github.com/oracle/provider-oci/internal/controller/streaming/streampool"
	subscriptionmapping "github.com/oracle/provider-oci/internal/controller/tenantmanagercontrolplane/subscriptionmapping"
	proxysubscriptionredeemableuser "github.com/oracle/provider-oci/internal/controller/usageapi/proxysubscriptionredeemableuser"
	secret "github.com/oracle/provider-oci/internal/controller/vault/secret"
	instvbsinstance "github.com/oracle/provider-oci/internal/controller/vbs/instvbsinstance"
	buildervbinstance "github.com/oracle/provider-oci/internal/controller/visualbuilder/buildervbinstance"
	monitoringpathanalysi "github.com/oracle/provider-oci/internal/controller/vn/monitoringpathanalysi"
	scanningcontainerscanrecipe "github.com/oracle/provider-oci/internal/controller/vulnerabilityscanning/scanningcontainerscanrecipe"
	scanningcontainerscantarget "github.com/oracle/provider-oci/internal/controller/vulnerabilityscanning/scanningcontainerscantarget"
	scanninghostscanrecipe "github.com/oracle/provider-oci/internal/controller/vulnerabilityscanning/scanninghostscanrecipe"
	scanninghostscantarget "github.com/oracle/provider-oci/internal/controller/vulnerabilityscanning/scanninghostscantarget"
	webappacceleration "github.com/oracle/provider-oci/internal/controller/waa/webappacceleration"
	webappaccelerationpolicy "github.com/oracle/provider-oci/internal/controller/waa/webappaccelerationpolicy"
	addresslist "github.com/oracle/provider-oci/internal/controller/waas/addresslist"
	customprotectionrule "github.com/oracle/provider-oci/internal/controller/waas/customprotectionrule"
	httpredirect "github.com/oracle/provider-oci/internal/controller/waas/httpredirect"
	protectionrule "github.com/oracle/provider-oci/internal/controller/waas/protectionrule"
	purgecache "github.com/oracle/provider-oci/internal/controller/waas/purgecache"
	waascertificate "github.com/oracle/provider-oci/internal/controller/waas/waascertificate"
	waaspolicy "github.com/oracle/provider-oci/internal/controller/waas/waaspolicy"
	networkaddresslist "github.com/oracle/provider-oci/internal/controller/waf/networkaddresslist"
	webappfirewall "github.com/oracle/provider-oci/internal/controller/waf/webappfirewall"
	webappfirewallpolicy "github.com/oracle/provider-oci/internal/controller/waf/webappfirewallpolicy"
	zprconfiguration "github.com/oracle/provider-oci/internal/controller/zpr/zprconfiguration"
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
		documentmodel.Setup,
		documentprocessorjob.Setup,
		documentproject.Setup,
		languageendpoint.Setup,
		languagemodel.Setup,
		languageproject.Setup,
		visionmodel.Setup,
		visionproject.Setup,
		analyticsinstance.Setup,
		instanceprivateaccesschannel.Setup,
		instancevanityurl.Setup,
		serviceannouncementsubscription.Setup,
		serviceannouncementsubscriptionsactionschangecompartment.Setup,
		serviceannouncementsubscriptionsfiltergroup.Setup,
		platformapiplatforminstance.Setup,
		privilegedapicontrol.Setup,
		privilegedapirequest.Setup,
		api.Setup,
		apigatewaycertificate.Setup,
		apigatewaydeployment.Setup,
		gateway.Setup,
		subscriber.Setup,
		usageplan.Setup,
		apmdomain.Setup,
		config.Setup,
		syntheticsdedicatedvantagepoint.Setup,
		syntheticsmonitor.Setup,
		syntheticsonpremisevantagepoint.Setup,
		syntheticsonpremisevantagepointworker.Setup,
		syntheticsscript.Setup,
		tracesscheduledquery.Setup,
		controlmonitorpluginmanagement.Setup,
		containerconfiguration.Setup,
		containerimagesignature.Setup,
		containerrepository.Setup,
		genericartifact.Setup,
		repository.Setup,
		auditconfiguration.Setup,
		autoscalingautoscalingconfiguration.Setup,
		bastionresource.Setup,
		session.Setup,
		bigdataserviceautoscalingconfiguration.Setup,
		bigdataserviceinstance.Setup,
		capacityreport.Setup,
		instanceapikey.Setup,
		instanceidentityconfiguration.Setup,
		instancemetastoreconfig.Setup,
		instancenodebackup.Setup,
		instancenodebackupconfiguration.Setup,
		instancenodereplaceconfiguration.Setup,
		instanceoperationcertificatemanagement.Setup,
		instanceospatchaction.Setup,
		instancepatchaction.Setup,
		instancereplacenodeaction.Setup,
		instanceresourceprincipalconfiguration.Setup,
		instancesoftwareupdateaction.Setup,
		osn.Setup,
		peer.Setup,
		platform.Setup,
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
		budgetresource.Setup,
		managementinternaloccmdemandsignal.Setup,
		managementinternaloccmdemandsignaldelivery.Setup,
		managementoccavailabilitycatalog.Setup,
		managementocccapacityrequest.Setup,
		managementocccustomergroup.Setup,
		managementocccustomergroupocccustomer.Setup,
		managementoccmdemandsignal.Setup,
		managementoccmdemandsignalitem.Setup,
		certificateauthority.Setup,
		managementcabundle.Setup,
		managementcertificate.Setup,
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
		placementgroupsclusterplacementgroup.Setup,
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
		cloudatcustomercccinfrastructure.Setup,
		cloudatcustomercccupgradeschedule.Setup,
		instancescontainerinstance.Setup,
		addon.Setup,
		cluster.Setup,
		clustercompletecredentialrotationmanagement.Setup,
		clusterstartcredentialrotationmanagement.Setup,
		clusterworkloadmapping.Setup,
		nodepool.Setup,
		virtualnodepool.Setup,
		contentexperienceinstance.Setup,
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
		databasebackup.Setup,
		databasedbmfeatmgmt.Setup,
		databasedbsystem.Setup,
		databasemigration.Setup,
		databaseresource.Setup,
		dbhome.Setup,
		dbmgmtprivateendpoint.Setup,
		dbnode.Setup,
		dbnodeconsoleconnection.Setup,
		dbnodeconsolehistory.Setup,
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
		softwareimage.Setup,
		toolsdatabasetoolsconnection.Setup,
		toolsdatabasetoolsprivateendpoint.Setup,
		upgrade.Setup,
		vmcluster.Setup,
		vmclusteraddvirtualmachine.Setup,
		vmclusternetwork.Setup,
		vmclusterremovevirtualmachine.Setup,
		catalog.Setup,
		catalogprivateendpoint.Setup,
		datacatalogconnection.Setup,
		metastore.Setup,
		dataflowapplication.Setup,
		dataflowprivateendpoint.Setup,
		invokerun.Setup,
		pool.Setup,
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
		datascienceprivateendpoint.Setup,
		datascienceproject.Setup,
		job.Setup,
		jobrun.Setup,
		mlapplication.Setup,
		mlapplicationimplementation.Setup,
		mlapplicationinstance.Setup,
		model.Setup,
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
		schedule.Setup,
		vulnerabilityscan.Setup,
		multicloudresourcediscovery.Setup,
		oracledbazureblobcontainer.Setup,
		oracledbazureblobmount.Setup,
		oracledbazureconnector.Setup,
		oracledbazurevault.Setup,
		oracledbazurevaultassociation.Setup,
		accesscontroldelegationcontrol.Setup,
		accesscontroldelegationsubscription.Setup,
		signaloccdemandsignal.Setup,
		desktoppool.Setup,
		buildpipeline.Setup,
		buildpipelinestage.Setup,
		buildrun.Setup,
		deployartifact.Setup,
		deployenvironment.Setup,
		deploypipeline.Setup,
		deploystage.Setup,
		devopsconnection.Setup,
		devopsdeployment.Setup,
		devopsproject.Setup,
		devopsrepository.Setup,
		projectrepositorysetting.Setup,
		repositorymirror.Setup,
		repositoryprotectedbranchmanagement.Setup,
		repositoryref.Setup,
		repositorysetting.Setup,
		trigger.Setup,
		digitalassistantinstance.Setup,
		digitalassistantprivateendpoint.Setup,
		privateendpointattachment.Setup,
		privateendpointscanproxy.Setup,
		recoverydrplan.Setup,
		recoverydrplanexecution.Setup,
		recoverydrprotectiongroup.Setup,
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
		emaildataplanedomain.Setup,
		returnpath.Setup,
		sender.Setup,
		suppression.Setup,
		rule.Setup,
		export.Setup,
		exportset.Setup,
		filesystem.Setup,
		mounttarget.Setup,
		replication.Setup,
		snapshot.Setup,
		storagefilesystemquotarule.Setup,
		storagefilesystemsnapshotpolicy.Setup,
		storageoutboundconnector.Setup,
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
		application.Setup,
		function.Setup,
		invokefunction.Setup,
		appsfusionenvironment.Setup,
		appsfusionenvironmentadminuser.Setup,
		appsfusionenvironmentfamily.Setup,
		appsfusionenvironmentrefreshactivity.Setup,
		appsfusionenvironmentserviceattachment.Setup,
		aiagent.Setup,
		aiagentendpoint.Setup,
		aiagentknowledgebase.Setup,
		aiagenttool.Setup,
		aidedicatedaicluster.Setup,
		aiendpoint.Setup,
		aimodel.Setup,
		artifactscontentartifactbypath.Setup,
		distributeddatabaseprivateendpoint.Setup,
		distributeddatabaseshardeddatabase.Setup,
		gateconnection.Setup,
		gateconnectionassignment.Setup,
		gatedatabaseregistration.Setup,
		gatedeployment.Setup,
		gatedeploymentbackup.Setup,
		gatedeploymentcertificate.Setup,
		gatepipeline.Setup,
		checkshttpprobe.Setup,
		checkspingprobe.Setup,
		httpmonitor.Setup,
		pingmonitor.Setup,
		apikey.Setup,
		authenticationpolicy.Setup,
		authtoken.Setup,
		compartment.Setup,
		customersecretkey.Setup,
		dbcredential.Setup,
		domainreplicationtoregion.Setup,
		domainsaccountrecoverysetting.Setup,
		domainsapikey.Setup,
		domainsapp.Setup,
		domainsapprole.Setup,
		domainsapprovalworkflow.Setup,
		domainsapprovalworkflowassignment.Setup,
		domainsapprovalworkflowstep.Setup,
		domainsauthenticationfactorsetting.Setup,
		domainsauthtoken.Setup,
		domainscloudgate.Setup,
		domainscloudgatemapping.Setup,
		domainscloudgateserver.Setup,
		domainscondition.Setup,
		domainscustomersecretkey.Setup,
		domainsdynamicresourcegroup.Setup,
		domainsgrant.Setup,
		domainsgroup.Setup,
		domainsidentitypropagationtrust.Setup,
		domainsidentityprovider.Setup,
		domainsidentitysetting.Setup,
		domainskmsisetting.Setup,
		domainsmyapikey.Setup,
		domainsmyauthtoken.Setup,
		domainsmycustomersecretkey.Setup,
		domainsmyoauth2clientcredential.Setup,
		domainsmyrequest.Setup,
		domainsmysmtpcredential.Setup,
		domainsmysupportaccount.Setup,
		domainsmyuserdbcredential.Setup,
		domainsnetworkperimeter.Setup,
		domainsnotificationsetting.Setup,
		domainsoauth2clientcredential.Setup,
		domainsoauthclientcertificate.Setup,
		domainsoauthpartnercertificate.Setup,
		domainspasswordpolicy.Setup,
		domainspolicy.Setup,
		domainsrule.Setup,
		domainssecurityquestion.Setup,
		domainssecurityquestionsetting.Setup,
		domainsselfregistrationprofile.Setup,
		domainssetting.Setup,
		domainssmtpcredential.Setup,
		domainssocialidentityprovider.Setup,
		domainsuser.Setup,
		domainsuserdbcredential.Setup,
		dynamicgroup.Setup,
		group.Setup,
		identitydomain.Setup,
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
		integrationinstance.Setup,
		oraclemanagedcustomendpoint.Setup,
		privateendpointoutboundconnection.Setup,
		fleet.Setup,
		fleetadvancedfeatureconfiguration.Setup,
		javadownloadsjavadownloadreport.Setup,
		javadownloadsjavadownloadtoken.Setup,
		javadownloadsjavalicenseacceptancerecord.Setup,
		plugin.Setup,
		ekmsprivateendpoint.Setup,
		encrypteddata.Setup,
		generatedkey.Setup,
		key.Setup,
		keyversion.Setup,
		sign.Setup,
		vault.Setup,
		vaultreplication.Setup,
		verify.Setup,
		managerconfiguration.Setup,
		managerlicenserecord.Setup,
		managerproductlicense.Setup,
		quota.Setup,
		backend.Setup,
		backendset.Setup,
		balancer.Setup,
		balancerbackendset.Setup,
		certificate.Setup,
		lbhostname.Setup,
		listener.Setup,
		loadbalancer.Setup,
		pathrouteset.Setup,
		routingpolicy.Setup,
		ruleset.Setup,
		sslciphersuite.Setup,
		analyticsloganalyticsentity.Setup,
		analyticsloganalyticsentitytype.Setup,
		analyticsloganalyticsimportcustomcontent.Setup,
		analyticsloganalyticsloggroup.Setup,
		analyticsloganalyticsobjectcollectionrule.Setup,
		analyticsloganalyticspreferencesmanagement.Setup,
		analyticsloganalyticsresourcecategoriesmanagement.Setup,
		analyticsnamespace.Setup,
		analyticsnamespaceingesttimerule.Setup,
		analyticsnamespaceingesttimerulesmanagement.Setup,
		analyticsnamespacelookup.Setup,
		analyticsnamespacescheduledtask.Setup,
		analyticsnamespacestoragearchivalconfig.Setup,
		analyticsnamespacestorageenabledisablearchiving.Setup,
		log.Setup,
		loggroup.Setup,
		logsavedsearch.Setup,
		unifiedagentconfiguration.Setup,
		filestoragelustrefilesystem.Setup,
		agentmanagementagent.Setup,
		agentmanagementagentinstallkey.Setup,
		agentnamedcredential.Setup,
		dashboardmanagementdashboardsimport.Setup,
		acceptedagreement.Setup,
		listingpackageagreement.Setup,
		publication.Setup,
		servicesmediaasset.Setup,
		servicesmediaworkflow.Setup,
		servicesmediaworkflowconfiguration.Setup,
		servicesmediaworkflowjob.Setup,
		servicesstreamcdnconfig.Setup,
		servicesstreamdistributionchannel.Setup,
		servicesstreampackagingconfig.Setup,
		computationcustomtable.Setup,
		computationquery.Setup,
		computationschedule.Setup,
		computationusage.Setup,
		computationusagecarbonemission.Setup,
		computationusagecarbonemissionsquery.Setup,
		computationusagestatementemailrecipientsgroup.Setup,
		alarm.Setup,
		alarmsuppression.Setup,
		capturefilter.Setup,
		vtap.Setup,
		channel.Setup,
		heatwavecluster.Setup,
		mysqlbackup.Setup,
		mysqlconfiguration.Setup,
		mysqldbsystem.Setup,
		replica.Setup,
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
		defaultdhcpoptions.Setup,
		defaultroutetable.Setup,
		defaultsecuritylist.Setup,
		dhcpoptions.Setup,
		firewallnetworkfirewallpolicyaddresslist.Setup,
		firewallnetworkfirewallpolicyapplication.Setup,
		firewallnetworkfirewallpolicyapplicationgroup.Setup,
		firewallnetworkfirewallpolicydecryptionprofile.Setup,
		firewallnetworkfirewallpolicydecryptionrule.Setup,
		firewallnetworkfirewallpolicymappedsecret.Setup,
		firewallnetworkfirewallpolicynatrule.Setup,
		firewallnetworkfirewallpolicysecurityrule.Setup,
		firewallnetworkfirewallpolicyservice.Setup,
		firewallnetworkfirewallpolicytunnelinspectionrule.Setup,
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
		index.Setup,
		nosqlconfiguration.Setup,
		table.Setup,
		tablereplica.Setup,
		bucket.Setup,
		namespacemetadata.Setup,
		object.Setup,
		objectlifecyclepolicy.Setup,
		objectstorageprivateendpoint.Setup,
		preauthrequest.Setup,
		replicationpolicy.Setup,
		esxihost.Setup,
		ocvscluster.Setup,
		sddc.Setup,
		notificationtopic.Setup,
		subscription.Setup,
		opainstance.Setup,
		clusterpipeline.Setup,
		opensearchcluster.Setup,
		accesscontroloperatorcontrol.Setup,
		accesscontroloperatorcontrolassignment.Setup,
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
		operatoraccesscontrolconfiguration.Setup,
		managementhubevent.Setup,
		managementhublifecycleenvironment.Setup,
		managementhublifecyclestageattachmanagedinstancesmanagement.Setup,
		managementhublifecyclestagedetachmanagedinstancesmanagement.Setup,
		managementhublifecyclestagepromotesoftwaresourcemanagement.Setup,
		managementhublifecyclestagerebootmanagement.Setup,
		managementhubmanagedinstance.Setup,
		managementhubmanagedinstanceattachprofilemanagement.Setup,
		managementhubmanagedinstancedetachprofilemanagement.Setup,
		managementhubmanagedinstancegroup.Setup,
		managementhubmanagedinstancegroupattachmanagedinstancesmanagement.Setup,
		managementhubmanagedinstancegroupattachsoftwaresourcesmanagement.Setup,
		managementhubmanagedinstancegroupdetachmanagedinstancesmanagement.Setup,
		managementhubmanagedinstancegroupdetachsoftwaresourcesmanagement.Setup,
		managementhubmanagedinstancegroupinstallpackagesmanagement.Setup,
		managementhubmanagedinstancegroupinstallwindowsupdatesmanagement.Setup,
		managementhubmanagedinstancegroupmanagemodulestreamsmanagement.Setup,
		managementhubmanagedinstancegrouprebootmanagement.Setup,
		managementhubmanagedinstancegroupremovepackagesmanagement.Setup,
		managementhubmanagedinstancegroupupdateallpackagesmanagement.Setup,
		managementhubmanagedinstanceinstallwindowsupdatesmanagement.Setup,
		managementhubmanagedinstancerebootmanagement.Setup,
		managementhubmanagedinstanceupdatepackagesmanagement.Setup,
		managementhubmanagementstation.Setup,
		managementhubmanagementstationassociatemanagedinstancesmanagement.Setup,
		managementhubmanagementstationmirrorsynchronizemanagement.Setup,
		managementhubmanagementstationrefreshmanagement.Setup,
		managementhubmanagementstationsynchronizemirrorsmanagement.Setup,
		managementhubprofile.Setup,
		managementhubprofileattachlifecyclestagemanagement.Setup,
		managementhubprofileattachmanagedinstancegroupmanagement.Setup,
		managementhubprofileattachmanagementstationmanagement.Setup,
		managementhubprofileattachsoftwaresourcesmanagement.Setup,
		managementhubprofiledetachsoftwaresourcesmanagement.Setup,
		managementhubscheduledjob.Setup,
		managementhubsoftwaresource.Setup,
		managementhubsoftwaresourceaddpackagesmanagement.Setup,
		managementhubsoftwaresourcechangeavailabilitymanagement.Setup,
		managementhubsoftwaresourcegeneratemetadatamanagement.Setup,
		managementhubsoftwaresourcemanifest.Setup,
		managementhubsoftwaresourceremovepackagesmanagement.Setup,
		managementhubsoftwaresourcereplacepackagesmanagement.Setup,
		managementhubworkrequestrerunmanagement.Setup,
		gatewayaddressactionverification.Setup,
		gatewaysubscription.Setup,
		providerconfig.Setup,
		psqlbackup.Setup,
		psqlconfiguration.Setup,
		psqldbsystem.Setup,
		queueresource.Setup,
		protecteddatabase.Setup,
		protectionpolicy.Setup,
		servicesubnet.Setup,
		clusterattachocicacheuser.Setup,
		clustercreateidentitytoken.Setup,
		clusterdetachocicacheuser.Setup,
		clustergetocicacheuser.Setup,
		ocicacheuser.Setup,
		ocicacheusergetrediscluster.Setup,
		rediscluster.Setup,
		schedulerschedule.Setup,
		resourcemanagerprivateendpoint.Setup,
		serviceconnector.Setup,
		attributesecurityattribute.Setup,
		attributesecurityattributenamespace.Setup,
		catalogprivateapplication.Setup,
		catalogservicecatalog.Setup,
		catalogservicecatalogassociation.Setup,
		monitoringbaselineablemetric.Setup,
		monitoringconfig.Setup,
		monitoringdiscoveryjob.Setup,
		monitoringmaintenancewindow.Setup,
		monitoringmaintenancewindowsretryfailedoperation.Setup,
		monitoringmaintenancewindowsstop.Setup,
		monitoringmetricextension.Setup,
		monitoringmetricextensionmetricextensionongivenresourcesmanagement.Setup,
		monitoringmonitoredresource.Setup,
		monitoringmonitoredresourcesassociatemonitoredresource.Setup,
		monitoringmonitoredresourceslistmember.Setup,
		monitoringmonitoredresourcessearch.Setup,
		monitoringmonitoredresourcessearchassociation.Setup,
		monitoringmonitoredresourcetask.Setup,
		monitoringmonitoredresourcetype.Setup,
		monitoringprocessset.Setup,
		monitoringtemplate.Setup,
		monitoringtemplatealarmcondition.Setup,
		monitoringtemplatemonitoringtemplateongivenresourcesmanagement.Setup,
		connectharness.Setup,
		stream.Setup,
		streampool.Setup,
		subscriptionmapping.Setup,
		proxysubscriptionredeemableuser.Setup,
		secret.Setup,
		instvbsinstance.Setup,
		buildervbinstance.Setup,
		monitoringpathanalysi.Setup,
		scanningcontainerscanrecipe.Setup,
		scanningcontainerscantarget.Setup,
		scanninghostscanrecipe.Setup,
		scanninghostscantarget.Setup,
		webappacceleration.Setup,
		webappaccelerationpolicy.Setup,
		addresslist.Setup,
		customprotectionrule.Setup,
		httpredirect.Setup,
		protectionrule.Setup,
		purgecache.Setup,
		waascertificate.Setup,
		waaspolicy.Setup,
		networkaddresslist.Setup,
		webappfirewall.Setup,
		webappfirewallpolicy.Setup,
		zprconfiguration.Setup,
		zprpolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

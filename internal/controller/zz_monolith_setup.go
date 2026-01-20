/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	containerconfiguration "github.com/oracle/provider-oci/internal/controller/artifacts/containerconfiguration"
	containerrepository "github.com/oracle/provider-oci/internal/controller/artifacts/containerrepository"
	genericartifact "github.com/oracle/provider-oci/internal/controller/artifacts/genericartifact"
	repository "github.com/oracle/provider-oci/internal/controller/artifacts/repository"
	bootvolume "github.com/oracle/provider-oci/internal/controller/blockstorage/bootvolume"
	bootvolumebackup "github.com/oracle/provider-oci/internal/controller/blockstorage/bootvolumebackup"
	volume "github.com/oracle/provider-oci/internal/controller/blockstorage/volume"
	volumeattachment "github.com/oracle/provider-oci/internal/controller/blockstorage/volumeattachment"
	volumebackup "github.com/oracle/provider-oci/internal/controller/blockstorage/volumebackup"
	volumebackuppolicy "github.com/oracle/provider-oci/internal/controller/blockstorage/volumebackuppolicy"
	volumebackuppolicyassignment "github.com/oracle/provider-oci/internal/controller/blockstorage/volumebackuppolicyassignment"
	volumegroup "github.com/oracle/provider-oci/internal/controller/blockstorage/volumegroup"
	volumegroupbackup "github.com/oracle/provider-oci/internal/controller/blockstorage/volumegroupbackup"
	certificateauthority "github.com/oracle/provider-oci/internal/controller/certificatesmanagement/certificateauthority"
	appcataloglistingresourceversionagreement "github.com/oracle/provider-oci/internal/controller/compute/appcataloglistingresourceversionagreement"
	appcatalogsubscription "github.com/oracle/provider-oci/internal/controller/compute/appcatalogsubscription"
	clusternetwork "github.com/oracle/provider-oci/internal/controller/compute/clusternetwork"
	computecapacityreservation "github.com/oracle/provider-oci/internal/controller/compute/computecapacityreservation"
	computecluster "github.com/oracle/provider-oci/internal/controller/compute/computecluster"
	computeimagecapabilityschema "github.com/oracle/provider-oci/internal/controller/compute/computeimagecapabilityschema"
	consolehistory "github.com/oracle/provider-oci/internal/controller/compute/consolehistory"
	dedicatedvmhost "github.com/oracle/provider-oci/internal/controller/compute/dedicatedvmhost"
	image "github.com/oracle/provider-oci/internal/controller/compute/image"
	instance "github.com/oracle/provider-oci/internal/controller/compute/instance"
	instanceconfiguration "github.com/oracle/provider-oci/internal/controller/compute/instanceconfiguration"
	instanceconsoleconnection "github.com/oracle/provider-oci/internal/controller/compute/instanceconsoleconnection"
	instancepool "github.com/oracle/provider-oci/internal/controller/compute/instancepool"
	instancepoolinstance "github.com/oracle/provider-oci/internal/controller/compute/instancepoolinstance"
	shapemanagement "github.com/oracle/provider-oci/internal/controller/compute/shapemanagement"
	cluster "github.com/oracle/provider-oci/internal/controller/containerengine/cluster"
	nodepool "github.com/oracle/provider-oci/internal/controller/containerengine/nodepool"
	record "github.com/oracle/provider-oci/internal/controller/dns/record"
	resolver "github.com/oracle/provider-oci/internal/controller/dns/resolver"
	resolverendpoint "github.com/oracle/provider-oci/internal/controller/dns/resolverendpoint"
	rrset "github.com/oracle/provider-oci/internal/controller/dns/rrset"
	steeringpolicy "github.com/oracle/provider-oci/internal/controller/dns/steeringpolicy"
	steeringpolicyattachment "github.com/oracle/provider-oci/internal/controller/dns/steeringpolicyattachment"
	tsigkey "github.com/oracle/provider-oci/internal/controller/dns/tsigkey"
	view "github.com/oracle/provider-oci/internal/controller/dns/view"
	zone "github.com/oracle/provider-oci/internal/controller/dns/zone"
	rule "github.com/oracle/provider-oci/internal/controller/events/rule"
	export "github.com/oracle/provider-oci/internal/controller/filestorage/export"
	exportset "github.com/oracle/provider-oci/internal/controller/filestorage/exportset"
	filesystem "github.com/oracle/provider-oci/internal/controller/filestorage/filesystem"
	mounttarget "github.com/oracle/provider-oci/internal/controller/filestorage/mounttarget"
	replication "github.com/oracle/provider-oci/internal/controller/filestorage/replication"
	snapshot "github.com/oracle/provider-oci/internal/controller/filestorage/snapshot"
	application "github.com/oracle/provider-oci/internal/controller/functions/application"
	function "github.com/oracle/provider-oci/internal/controller/functions/function"
	invokefunction "github.com/oracle/provider-oci/internal/controller/functions/invokefunction"
	httpmonitor "github.com/oracle/provider-oci/internal/controller/healthchecks/httpmonitor"
	pingmonitor "github.com/oracle/provider-oci/internal/controller/healthchecks/pingmonitor"
	apikey "github.com/oracle/provider-oci/internal/controller/identity/apikey"
	authenticationpolicy "github.com/oracle/provider-oci/internal/controller/identity/authenticationpolicy"
	authtoken "github.com/oracle/provider-oci/internal/controller/identity/authtoken"
	compartment "github.com/oracle/provider-oci/internal/controller/identity/compartment"
	customersecretkey "github.com/oracle/provider-oci/internal/controller/identity/customersecretkey"
	dataplanegeneratescopedaccesstoken "github.com/oracle/provider-oci/internal/controller/identity/dataplanegeneratescopedaccesstoken"
	dbcredential "github.com/oracle/provider-oci/internal/controller/identity/dbcredential"
	domain "github.com/oracle/provider-oci/internal/controller/identity/domain"
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
	key "github.com/oracle/provider-oci/internal/controller/kms/key"
	keyversion "github.com/oracle/provider-oci/internal/controller/kms/keyversion"
	vault "github.com/oracle/provider-oci/internal/controller/kms/vault"
	backend "github.com/oracle/provider-oci/internal/controller/loadbalancer/backend"
	backendset "github.com/oracle/provider-oci/internal/controller/loadbalancer/backendset"
	certificate "github.com/oracle/provider-oci/internal/controller/loadbalancer/certificate"
	lbhostname "github.com/oracle/provider-oci/internal/controller/loadbalancer/lbhostname"
	listener "github.com/oracle/provider-oci/internal/controller/loadbalancer/listener"
	loadbalancer "github.com/oracle/provider-oci/internal/controller/loadbalancer/loadbalancer"
	pathrouteset "github.com/oracle/provider-oci/internal/controller/loadbalancer/pathrouteset"
	routingpolicy "github.com/oracle/provider-oci/internal/controller/loadbalancer/routingpolicy"
	ruleset "github.com/oracle/provider-oci/internal/controller/loadbalancer/ruleset"
	sslciphersuite "github.com/oracle/provider-oci/internal/controller/loadbalancer/sslciphersuite"
	log "github.com/oracle/provider-oci/internal/controller/logging/log"
	loggroup "github.com/oracle/provider-oci/internal/controller/logging/loggroup"
	logsavedsearch "github.com/oracle/provider-oci/internal/controller/logging/logsavedsearch"
	unifiedagentconfiguration "github.com/oracle/provider-oci/internal/controller/logging/unifiedagentconfiguration"
	alarm "github.com/oracle/provider-oci/internal/controller/monitoring/alarm"
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
	bucket "github.com/oracle/provider-oci/internal/controller/objectstorage/bucket"
	namespacemetadata "github.com/oracle/provider-oci/internal/controller/objectstorage/namespacemetadata"
	object "github.com/oracle/provider-oci/internal/controller/objectstorage/object"
	objectlifecyclepolicy "github.com/oracle/provider-oci/internal/controller/objectstorage/objectlifecyclepolicy"
	preauthrequest "github.com/oracle/provider-oci/internal/controller/objectstorage/preauthrequest"
	privateendpoint "github.com/oracle/provider-oci/internal/controller/objectstorage/privateendpoint"
	replicationpolicy "github.com/oracle/provider-oci/internal/controller/objectstorage/replicationpolicy"
	notificationtopic "github.com/oracle/provider-oci/internal/controller/ons/notificationtopic"
	subscription "github.com/oracle/provider-oci/internal/controller/ons/subscription"
	providerconfig "github.com/oracle/provider-oci/internal/controller/providerconfig"
	psqlbackup "github.com/oracle/provider-oci/internal/controller/psql/psqlbackup"
	psqlconfiguration "github.com/oracle/provider-oci/internal/controller/psql/psqlconfiguration"
	psqldbsystem "github.com/oracle/provider-oci/internal/controller/psql/psqldbsystem"
	connectharness "github.com/oracle/provider-oci/internal/controller/streaming/connectharness"
	stream "github.com/oracle/provider-oci/internal/controller/streaming/stream"
	streampool "github.com/oracle/provider-oci/internal/controller/streaming/streampool"
	secret "github.com/oracle/provider-oci/internal/controller/vault/secret"
)

// Setup_monolith creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_monolith(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		containerconfiguration.Setup,
		containerrepository.Setup,
		genericartifact.Setup,
		repository.Setup,
		bootvolume.Setup,
		bootvolumebackup.Setup,
		volume.Setup,
		volumeattachment.Setup,
		volumebackup.Setup,
		volumebackuppolicy.Setup,
		volumebackuppolicyassignment.Setup,
		volumegroup.Setup,
		volumegroupbackup.Setup,
		certificateauthority.Setup,
		appcataloglistingresourceversionagreement.Setup,
		appcatalogsubscription.Setup,
		clusternetwork.Setup,
		computecapacityreservation.Setup,
		computecluster.Setup,
		computeimagecapabilityschema.Setup,
		consolehistory.Setup,
		dedicatedvmhost.Setup,
		image.Setup,
		instance.Setup,
		instanceconfiguration.Setup,
		instanceconsoleconnection.Setup,
		instancepool.Setup,
		instancepoolinstance.Setup,
		shapemanagement.Setup,
		cluster.Setup,
		nodepool.Setup,
		record.Setup,
		resolver.Setup,
		resolverendpoint.Setup,
		rrset.Setup,
		steeringpolicy.Setup,
		steeringpolicyattachment.Setup,
		tsigkey.Setup,
		view.Setup,
		zone.Setup,
		rule.Setup,
		export.Setup,
		exportset.Setup,
		filesystem.Setup,
		mounttarget.Setup,
		replication.Setup,
		snapshot.Setup,
		application.Setup,
		function.Setup,
		invokefunction.Setup,
		httpmonitor.Setup,
		pingmonitor.Setup,
		apikey.Setup,
		authenticationpolicy.Setup,
		authtoken.Setup,
		compartment.Setup,
		customersecretkey.Setup,
		dataplanegeneratescopedaccesstoken.Setup,
		dbcredential.Setup,
		domain.Setup,
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
		key.Setup,
		keyversion.Setup,
		vault.Setup,
		backend.Setup,
		backendset.Setup,
		certificate.Setup,
		lbhostname.Setup,
		listener.Setup,
		loadbalancer.Setup,
		pathrouteset.Setup,
		routingpolicy.Setup,
		ruleset.Setup,
		sslciphersuite.Setup,
		log.Setup,
		loggroup.Setup,
		logsavedsearch.Setup,
		unifiedagentconfiguration.Setup,
		alarm.Setup,
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
		bucket.Setup,
		namespacemetadata.Setup,
		object.Setup,
		objectlifecyclepolicy.Setup,
		preauthrequest.Setup,
		privateendpoint.Setup,
		replicationpolicy.Setup,
		notificationtopic.Setup,
		subscription.Setup,
		providerconfig.Setup,
		psqlbackup.Setup,
		psqlconfiguration.Setup,
		psqldbsystem.Setup,
		connectharness.Setup,
		stream.Setup,
		streampool.Setup,
		secret.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

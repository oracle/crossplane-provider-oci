/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

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
)

// Setup_datasafe creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_datasafe(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
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
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_datasafe creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_datasafe(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
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
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

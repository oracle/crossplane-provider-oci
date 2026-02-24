/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	addsdmcolumns "github.com/oracle/provider-oci/internal/controller/datasafe/addsdmcolumns"
	alert "github.com/oracle/provider-oci/internal/controller/datasafe/alert"
	alertpolicy "github.com/oracle/provider-oci/internal/controller/datasafe/alertpolicy"
	alertpolicyrule "github.com/oracle/provider-oci/internal/controller/datasafe/alertpolicyrule"
	attributeset "github.com/oracle/provider-oci/internal/controller/datasafe/attributeset"
	auditarchiveretrieval "github.com/oracle/provider-oci/internal/controller/datasafe/auditarchiveretrieval"
	auditpolicy "github.com/oracle/provider-oci/internal/controller/datasafe/auditpolicy"
	auditpolicymanagement "github.com/oracle/provider-oci/internal/controller/datasafe/auditpolicymanagement"
	auditprofile "github.com/oracle/provider-oci/internal/controller/datasafe/auditprofile"
	auditprofilemanagement "github.com/oracle/provider-oci/internal/controller/datasafe/auditprofilemanagement"
	audittrail "github.com/oracle/provider-oci/internal/controller/datasafe/audittrail"
	audittrailmanagement "github.com/oracle/provider-oci/internal/controller/datasafe/audittrailmanagement"
	calculateauditvolumeavailable "github.com/oracle/provider-oci/internal/controller/datasafe/calculateauditvolumeavailable"
	calculateauditvolumecollected "github.com/oracle/provider-oci/internal/controller/datasafe/calculateauditvolumecollected"
	comparesecurityassessment "github.com/oracle/provider-oci/internal/controller/datasafe/comparesecurityassessment"
	compareuserassessment "github.com/oracle/provider-oci/internal/controller/datasafe/compareuserassessment"
	databasesecurityconfig "github.com/oracle/provider-oci/internal/controller/datasafe/databasesecurityconfig"
	databasesecurityconfigmanagement "github.com/oracle/provider-oci/internal/controller/datasafe/databasesecurityconfigmanagement"
	datasafeconfiguration "github.com/oracle/provider-oci/internal/controller/datasafe/datasafeconfiguration"
	datasafeprivateendpoint "github.com/oracle/provider-oci/internal/controller/datasafe/datasafeprivateendpoint"
	discoveryjob "github.com/oracle/provider-oci/internal/controller/datasafe/discoveryjob"
	discoveryjobsresult "github.com/oracle/provider-oci/internal/controller/datasafe/discoveryjobsresult"
	generateonpremconnectorconfiguration "github.com/oracle/provider-oci/internal/controller/datasafe/generateonpremconnectorconfiguration"
	librarymaskingformat "github.com/oracle/provider-oci/internal/controller/datasafe/librarymaskingformat"
	maskdata "github.com/oracle/provider-oci/internal/controller/datasafe/maskdata"
	maskingpoliciesapplydifferencetomaskingcolumns "github.com/oracle/provider-oci/internal/controller/datasafe/maskingpoliciesapplydifferencetomaskingcolumns"
	maskingpoliciesmaskingcolumn "github.com/oracle/provider-oci/internal/controller/datasafe/maskingpoliciesmaskingcolumn"
	maskingpolicy "github.com/oracle/provider-oci/internal/controller/datasafe/maskingpolicy"
	maskingpolicyhealthreportmanagement "github.com/oracle/provider-oci/internal/controller/datasafe/maskingpolicyhealthreportmanagement"
	maskingreportmanagement "github.com/oracle/provider-oci/internal/controller/datasafe/maskingreportmanagement"
	onpremconnector "github.com/oracle/provider-oci/internal/controller/datasafe/onpremconnector"
	report "github.com/oracle/provider-oci/internal/controller/datasafe/report"
	reportdefinition "github.com/oracle/provider-oci/internal/controller/datasafe/reportdefinition"
	sdmmaskingpolicydifference "github.com/oracle/provider-oci/internal/controller/datasafe/sdmmaskingpolicydifference"
	securityassessment "github.com/oracle/provider-oci/internal/controller/datasafe/securityassessment"
	securityassessmentcheck "github.com/oracle/provider-oci/internal/controller/datasafe/securityassessmentcheck"
	securityassessmentfinding "github.com/oracle/provider-oci/internal/controller/datasafe/securityassessmentfinding"
	securitypolicy "github.com/oracle/provider-oci/internal/controller/datasafe/securitypolicy"
	securitypolicyconfig "github.com/oracle/provider-oci/internal/controller/datasafe/securitypolicyconfig"
	securitypolicydeployment "github.com/oracle/provider-oci/internal/controller/datasafe/securitypolicydeployment"
	securitypolicydeploymentmanagement "github.com/oracle/provider-oci/internal/controller/datasafe/securitypolicydeploymentmanagement"
	securitypolicymanagement "github.com/oracle/provider-oci/internal/controller/datasafe/securitypolicymanagement"
	sensitivedatamodel "github.com/oracle/provider-oci/internal/controller/datasafe/sensitivedatamodel"
	sensitivedatamodelreferentialrelation "github.com/oracle/provider-oci/internal/controller/datasafe/sensitivedatamodelreferentialrelation"
	sensitivedatamodelsapplydiscoveryjobresults "github.com/oracle/provider-oci/internal/controller/datasafe/sensitivedatamodelsapplydiscoveryjobresults"
	sensitivedatamodelssensitivecolumn "github.com/oracle/provider-oci/internal/controller/datasafe/sensitivedatamodelssensitivecolumn"
	sensitivetype "github.com/oracle/provider-oci/internal/controller/datasafe/sensitivetype"
	sensitivetypegroup "github.com/oracle/provider-oci/internal/controller/datasafe/sensitivetypegroup"
	sensitivetypegroupgroupedsensitivetype "github.com/oracle/provider-oci/internal/controller/datasafe/sensitivetypegroupgroupedsensitivetype"
	sensitivetypesexport "github.com/oracle/provider-oci/internal/controller/datasafe/sensitivetypesexport"
	setsecurityassessmentbaseline "github.com/oracle/provider-oci/internal/controller/datasafe/setsecurityassessmentbaseline"
	setsecurityassessmentbaselinemanagement "github.com/oracle/provider-oci/internal/controller/datasafe/setsecurityassessmentbaselinemanagement"
	setuserassessmentbaseline "github.com/oracle/provider-oci/internal/controller/datasafe/setuserassessmentbaseline"
	setuserassessmentbaselinemanagement "github.com/oracle/provider-oci/internal/controller/datasafe/setuserassessmentbaselinemanagement"
	sqlcollection "github.com/oracle/provider-oci/internal/controller/datasafe/sqlcollection"
	sqlfirewallpolicy "github.com/oracle/provider-oci/internal/controller/datasafe/sqlfirewallpolicy"
	sqlfirewallpolicymanagement "github.com/oracle/provider-oci/internal/controller/datasafe/sqlfirewallpolicymanagement"
	targetalertpolicyassociation "github.com/oracle/provider-oci/internal/controller/datasafe/targetalertpolicyassociation"
	targetdatabase "github.com/oracle/provider-oci/internal/controller/datasafe/targetdatabase"
	targetdatabasegroup "github.com/oracle/provider-oci/internal/controller/datasafe/targetdatabasegroup"
	targetdatabasepeertargetdatabase "github.com/oracle/provider-oci/internal/controller/datasafe/targetdatabasepeertargetdatabase"
	unifiedauditpolicy "github.com/oracle/provider-oci/internal/controller/datasafe/unifiedauditpolicy"
	unifiedauditpolicydefinition "github.com/oracle/provider-oci/internal/controller/datasafe/unifiedauditpolicydefinition"
	unsetsecurityassessmentbaseline "github.com/oracle/provider-oci/internal/controller/datasafe/unsetsecurityassessmentbaseline"
	unsetsecurityassessmentbaselinemanagement "github.com/oracle/provider-oci/internal/controller/datasafe/unsetsecurityassessmentbaselinemanagement"
	unsetuserassessmentbaseline "github.com/oracle/provider-oci/internal/controller/datasafe/unsetuserassessmentbaseline"
	unsetuserassessmentbaselinemanagement "github.com/oracle/provider-oci/internal/controller/datasafe/unsetuserassessmentbaselinemanagement"
	userassessment "github.com/oracle/provider-oci/internal/controller/datasafe/userassessment"
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

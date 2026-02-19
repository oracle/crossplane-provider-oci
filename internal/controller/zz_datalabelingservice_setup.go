/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

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
)

// Setup_datalabelingservice creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_datalabelingservice(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
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
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

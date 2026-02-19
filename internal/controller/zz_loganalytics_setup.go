/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

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
)

// Setup_loganalytics creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_loganalytics(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
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
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

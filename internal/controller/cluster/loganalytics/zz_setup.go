/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

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
)

// Setup_loganalytics creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_loganalytics(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
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
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_loganalytics creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_loganalytics(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
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
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

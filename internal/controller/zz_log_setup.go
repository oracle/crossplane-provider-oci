/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

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
)

// Setup_log creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_log(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
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
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

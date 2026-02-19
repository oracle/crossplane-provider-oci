/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	monitoredregion "github.com/oracle/provider-oci/internal/controller/resourceanalytics/monitoredregion"
	resourceanalyticsinstance "github.com/oracle/provider-oci/internal/controller/resourceanalytics/resourceanalyticsinstance"
	resourceanalyticsinstanceoacmanagement "github.com/oracle/provider-oci/internal/controller/resourceanalytics/resourceanalyticsinstanceoacmanagement"
	tenancyattachment "github.com/oracle/provider-oci/internal/controller/resourceanalytics/tenancyattachment"
)

// Setup_resourceanalytics creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_resourceanalytics(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		monitoredregion.Setup,
		resourceanalyticsinstance.Setup,
		resourceanalyticsinstanceoacmanagement.Setup,
		tenancyattachment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

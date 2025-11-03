/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	serviceannouncementsubscription "github.com/oracle/provider-oci/internal/controller/announcements/serviceannouncementsubscription"
	serviceannouncementsubscriptionsactionschangecompartment "github.com/oracle/provider-oci/internal/controller/announcements/serviceannouncementsubscriptionsactionschangecompartment"
	serviceannouncementsubscriptionsfiltergroup "github.com/oracle/provider-oci/internal/controller/announcements/serviceannouncementsubscriptionsfiltergroup"
)

// Setup_announcements creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_announcements(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		serviceannouncementsubscription.Setup,
		serviceannouncementsubscriptionsactionschangecompartment.Setup,
		serviceannouncementsubscriptionsfiltergroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

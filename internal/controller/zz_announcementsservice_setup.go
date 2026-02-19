/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	announcementsubscription "github.com/oracle/provider-oci/internal/controller/announcementsservice/announcementsubscription"
	announcementsubscriptionsactionschangecompartment "github.com/oracle/provider-oci/internal/controller/announcementsservice/announcementsubscriptionsactionschangecompartment"
	announcementsubscriptionsfiltergroup "github.com/oracle/provider-oci/internal/controller/announcementsservice/announcementsubscriptionsfiltergroup"
)

// Setup_announcementsservice creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_announcementsservice(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		announcementsubscription.Setup,
		announcementsubscriptionsactionschangecompartment.Setup,
		announcementsubscriptionsfiltergroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

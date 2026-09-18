/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	announcementsubscription "github.com/oracle/provider-oci/internal/controller/cluster/announcementsservice/announcementsubscription"
	announcementsubscriptionsactionschangecompartment "github.com/oracle/provider-oci/internal/controller/cluster/announcementsservice/announcementsubscriptionsactionschangecompartment"
	announcementsubscriptionsfiltergroup "github.com/oracle/provider-oci/internal/controller/cluster/announcementsservice/announcementsubscriptionsfiltergroup"
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

// SetupGated_announcementsservice creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_announcementsservice(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		announcementsubscription.SetupGated,
		announcementsubscriptionsactionschangecompartment.SetupGated,
		announcementsubscriptionsfiltergroup.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

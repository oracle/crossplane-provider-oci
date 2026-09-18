/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	notificationtopic "github.com/oracle/provider-oci/internal/controller/namespaced/ons/notificationtopic"
	subscription "github.com/oracle/provider-oci/internal/controller/namespaced/ons/subscription"
)

// Setup_ons creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_ons(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		notificationtopic.Setup,
		subscription.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_ons creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_ons(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		notificationtopic.SetupGated,
		subscription.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

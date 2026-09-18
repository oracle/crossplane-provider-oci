/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	consumergroup "github.com/oracle/provider-oci/internal/controller/namespaced/queue/consumergroup"
	queue "github.com/oracle/provider-oci/internal/controller/namespaced/queue/queue"
)

// Setup_queue creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_queue(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		consumergroup.Setup,
		queue.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_queue creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_queue(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		consumergroup.SetupGated,
		queue.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

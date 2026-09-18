/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	connectharness "github.com/oracle/provider-oci/internal/controller/cluster/streaming/connectharness"
	stream "github.com/oracle/provider-oci/internal/controller/cluster/streaming/stream"
	streampool "github.com/oracle/provider-oci/internal/controller/cluster/streaming/streampool"
)

// Setup_streaming creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_streaming(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		connectharness.Setup,
		stream.Setup,
		streampool.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_streaming creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_streaming(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		connectharness.SetupGated,
		stream.SetupGated,
		streampool.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

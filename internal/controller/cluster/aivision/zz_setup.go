/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	model "github.com/oracle/provider-oci/internal/controller/cluster/aivision/model"
	project "github.com/oracle/provider-oci/internal/controller/cluster/aivision/project"
	streamgroup "github.com/oracle/provider-oci/internal/controller/cluster/aivision/streamgroup"
	streamjob "github.com/oracle/provider-oci/internal/controller/cluster/aivision/streamjob"
	streamsource "github.com/oracle/provider-oci/internal/controller/cluster/aivision/streamsource"
	visionprivateendpoint "github.com/oracle/provider-oci/internal/controller/cluster/aivision/visionprivateendpoint"
)

// Setup_aivision creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_aivision(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		model.Setup,
		project.Setup,
		streamgroup.Setup,
		streamjob.Setup,
		streamsource.Setup,
		visionprivateendpoint.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_aivision creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_aivision(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		model.SetupGated,
		project.SetupGated,
		streamgroup.SetupGated,
		streamjob.SetupGated,
		streamsource.SetupGated,
		visionprivateendpoint.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	model "github.com/oracle/provider-oci/internal/controller/aivision/model"
	project "github.com/oracle/provider-oci/internal/controller/aivision/project"
	streamgroup "github.com/oracle/provider-oci/internal/controller/aivision/streamgroup"
	streamjob "github.com/oracle/provider-oci/internal/controller/aivision/streamjob"
	streamsource "github.com/oracle/provider-oci/internal/controller/aivision/streamsource"
	visionprivateendpoint "github.com/oracle/provider-oci/internal/controller/aivision/visionprivateendpoint"
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

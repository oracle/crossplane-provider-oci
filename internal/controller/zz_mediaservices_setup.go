/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	mediaasset "github.com/oracle/provider-oci/internal/controller/mediaservices/mediaasset"
	mediaworkflow "github.com/oracle/provider-oci/internal/controller/mediaservices/mediaworkflow"
	mediaworkflowconfiguration "github.com/oracle/provider-oci/internal/controller/mediaservices/mediaworkflowconfiguration"
	mediaworkflowjob "github.com/oracle/provider-oci/internal/controller/mediaservices/mediaworkflowjob"
	streamcdnconfig "github.com/oracle/provider-oci/internal/controller/mediaservices/streamcdnconfig"
	streamdistributionchannel "github.com/oracle/provider-oci/internal/controller/mediaservices/streamdistributionchannel"
	streampackagingconfig "github.com/oracle/provider-oci/internal/controller/mediaservices/streampackagingconfig"
)

// Setup_mediaservices creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_mediaservices(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		mediaasset.Setup,
		mediaworkflow.Setup,
		mediaworkflowconfiguration.Setup,
		mediaworkflowjob.Setup,
		streamcdnconfig.Setup,
		streamdistributionchannel.Setup,
		streampackagingconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

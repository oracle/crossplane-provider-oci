/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	mediaasset "github.com/oracle/provider-oci/internal/controller/namespaced/mediaservices/mediaasset"
	mediaworkflow "github.com/oracle/provider-oci/internal/controller/namespaced/mediaservices/mediaworkflow"
	mediaworkflowconfiguration "github.com/oracle/provider-oci/internal/controller/namespaced/mediaservices/mediaworkflowconfiguration"
	mediaworkflowjob "github.com/oracle/provider-oci/internal/controller/namespaced/mediaservices/mediaworkflowjob"
	streamcdnconfig "github.com/oracle/provider-oci/internal/controller/namespaced/mediaservices/streamcdnconfig"
	streamdistributionchannel "github.com/oracle/provider-oci/internal/controller/namespaced/mediaservices/streamdistributionchannel"
	streampackagingconfig "github.com/oracle/provider-oci/internal/controller/namespaced/mediaservices/streampackagingconfig"
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

// SetupGated_mediaservices creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_mediaservices(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		mediaasset.SetupGated,
		mediaworkflow.SetupGated,
		mediaworkflowconfiguration.SetupGated,
		mediaworkflowjob.SetupGated,
		streamcdnconfig.SetupGated,
		streamdistributionchannel.SetupGated,
		streampackagingconfig.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	servicesmediaasset "github.com/oracle/provider-oci/internal/controller/mediaservices/servicesmediaasset"
	servicesmediaworkflow "github.com/oracle/provider-oci/internal/controller/mediaservices/servicesmediaworkflow"
	servicesmediaworkflowconfiguration "github.com/oracle/provider-oci/internal/controller/mediaservices/servicesmediaworkflowconfiguration"
	servicesmediaworkflowjob "github.com/oracle/provider-oci/internal/controller/mediaservices/servicesmediaworkflowjob"
	servicesstreamcdnconfig "github.com/oracle/provider-oci/internal/controller/mediaservices/servicesstreamcdnconfig"
	servicesstreamdistributionchannel "github.com/oracle/provider-oci/internal/controller/mediaservices/servicesstreamdistributionchannel"
	servicesstreampackagingconfig "github.com/oracle/provider-oci/internal/controller/mediaservices/servicesstreampackagingconfig"
)

// Setup_mediaservices creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_mediaservices(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		servicesmediaasset.Setup,
		servicesmediaworkflow.Setup,
		servicesmediaworkflowconfiguration.Setup,
		servicesmediaworkflowjob.Setup,
		servicesstreamcdnconfig.Setup,
		servicesstreamdistributionchannel.Setup,
		servicesstreampackagingconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	containerconfiguration "github.com/oracle/provider-oci/internal/controller/artifacts/containerconfiguration"
	containerimagesignature "github.com/oracle/provider-oci/internal/controller/artifacts/containerimagesignature"
	containerrepository "github.com/oracle/provider-oci/internal/controller/artifacts/containerrepository"
	genericartifact "github.com/oracle/provider-oci/internal/controller/artifacts/genericartifact"
	repository "github.com/oracle/provider-oci/internal/controller/artifacts/repository"
)

// Setup_artifacts creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_artifacts(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		containerconfiguration.Setup,
		containerimagesignature.Setup,
		containerrepository.Setup,
		genericartifact.Setup,
		repository.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

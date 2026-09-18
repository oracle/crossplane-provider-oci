/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	artifactbypath "github.com/oracle/provider-oci/internal/controller/cluster/genericartifactscontent/artifactbypath"
)

// Setup_genericartifactscontent creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_genericartifactscontent(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		artifactbypath.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_genericartifactscontent creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_genericartifactscontent(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		artifactbypath.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

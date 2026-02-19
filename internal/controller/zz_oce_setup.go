/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	oceinstance "github.com/oracle/provider-oci/internal/controller/oce/oceinstance"
)

// Setup_oce creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_oce(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		oceinstance.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

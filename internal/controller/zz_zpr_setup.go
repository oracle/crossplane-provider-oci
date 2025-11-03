/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	zprconfiguration "github.com/oracle/provider-oci/internal/controller/zpr/zprconfiguration"
	zprpolicy "github.com/oracle/provider-oci/internal/controller/zpr/zprpolicy"
)

// Setup_zpr creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_zpr(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		zprconfiguration.Setup,
		zprpolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

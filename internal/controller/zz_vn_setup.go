/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	monitoringpathanalysi "github.com/oracle/provider-oci/internal/controller/vn/monitoringpathanalysi"
)

// Setup_vn creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_vn(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		monitoringpathanalysi.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	gatewayaddressactionverification "github.com/oracle/provider-oci/internal/controller/osp/gatewayaddressactionverification"
	gatewaysubscription "github.com/oracle/provider-oci/internal/controller/osp/gatewaysubscription"
)

// Setup_osp creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_osp(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		gatewayaddressactionverification.Setup,
		gatewaysubscription.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

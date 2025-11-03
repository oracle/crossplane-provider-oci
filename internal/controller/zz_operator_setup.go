/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	accesscontroloperatorcontrol "github.com/oracle/provider-oci/internal/controller/operator/accesscontroloperatorcontrol"
	accesscontroloperatorcontrolassignment "github.com/oracle/provider-oci/internal/controller/operator/accesscontroloperatorcontrolassignment"
)

// Setup_operator creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_operator(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesscontroloperatorcontrol.Setup,
		accesscontroloperatorcontrolassignment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

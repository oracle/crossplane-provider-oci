/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	esxihost "github.com/oracle/provider-oci/internal/controller/ocvs/esxihost"
	ocvscluster "github.com/oracle/provider-oci/internal/controller/ocvs/ocvscluster"
	sddc "github.com/oracle/provider-oci/internal/controller/ocvs/sddc"
)

// Setup_ocvs creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_ocvs(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		esxihost.Setup,
		ocvscluster.Setup,
		sddc.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

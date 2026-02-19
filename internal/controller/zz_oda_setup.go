/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	odainstance "github.com/oracle/provider-oci/internal/controller/oda/odainstance"
	odaprivateendpoint "github.com/oracle/provider-oci/internal/controller/oda/odaprivateendpoint"
	odaprivateendpointattachment "github.com/oracle/provider-oci/internal/controller/oda/odaprivateendpointattachment"
	odaprivateendpointscanproxy "github.com/oracle/provider-oci/internal/controller/oda/odaprivateendpointscanproxy"
)

// Setup_oda creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_oda(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		odainstance.Setup,
		odaprivateendpoint.Setup,
		odaprivateendpointattachment.Setup,
		odaprivateendpointscanproxy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

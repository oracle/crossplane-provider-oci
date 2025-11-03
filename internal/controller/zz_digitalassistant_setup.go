/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	digitalassistantinstance "github.com/oracle/provider-oci/internal/controller/digitalassistant/digitalassistantinstance"
	digitalassistantprivateendpoint "github.com/oracle/provider-oci/internal/controller/digitalassistant/digitalassistantprivateendpoint"
	privateendpointattachment "github.com/oracle/provider-oci/internal/controller/digitalassistant/privateendpointattachment"
	privateendpointscanproxy "github.com/oracle/provider-oci/internal/controller/digitalassistant/privateendpointscanproxy"
)

// Setup_digitalassistant creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_digitalassistant(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		digitalassistantinstance.Setup,
		digitalassistantprivateendpoint.Setup,
		privateendpointattachment.Setup,
		privateendpointscanproxy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

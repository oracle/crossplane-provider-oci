/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	api "github.com/oracle/provider-oci/internal/controller/apigateway/api"
	certificate "github.com/oracle/provider-oci/internal/controller/apigateway/certificate"
	deployment "github.com/oracle/provider-oci/internal/controller/apigateway/deployment"
	gateway "github.com/oracle/provider-oci/internal/controller/apigateway/gateway"
	subscriber "github.com/oracle/provider-oci/internal/controller/apigateway/subscriber"
	usageplan "github.com/oracle/provider-oci/internal/controller/apigateway/usageplan"
)

// Setup_apigateway creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_apigateway(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		api.Setup,
		certificate.Setup,
		deployment.Setup,
		gateway.Setup,
		subscriber.Setup,
		usageplan.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

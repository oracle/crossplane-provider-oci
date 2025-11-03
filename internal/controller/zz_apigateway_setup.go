/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	api "github.com/oracle/provider-oci/internal/controller/apigateway/api"
	apigatewaycertificate "github.com/oracle/provider-oci/internal/controller/apigateway/apigatewaycertificate"
	apigatewaydeployment "github.com/oracle/provider-oci/internal/controller/apigateway/apigatewaydeployment"
	gateway "github.com/oracle/provider-oci/internal/controller/apigateway/gateway"
	subscriber "github.com/oracle/provider-oci/internal/controller/apigateway/subscriber"
	usageplan "github.com/oracle/provider-oci/internal/controller/apigateway/usageplan"
)

// Setup_apigateway creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_apigateway(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		api.Setup,
		apigatewaycertificate.Setup,
		apigatewaydeployment.Setup,
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

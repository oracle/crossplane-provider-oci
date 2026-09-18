/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	api "github.com/oracle/provider-oci/internal/controller/namespaced/apigateway/api"
	certificate "github.com/oracle/provider-oci/internal/controller/namespaced/apigateway/certificate"
	deployment "github.com/oracle/provider-oci/internal/controller/namespaced/apigateway/deployment"
	gateway "github.com/oracle/provider-oci/internal/controller/namespaced/apigateway/gateway"
	subscriber "github.com/oracle/provider-oci/internal/controller/namespaced/apigateway/subscriber"
	usageplan "github.com/oracle/provider-oci/internal/controller/namespaced/apigateway/usageplan"
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

// SetupGated_apigateway creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_apigateway(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		api.SetupGated,
		certificate.SetupGated,
		deployment.SetupGated,
		gateway.SetupGated,
		subscriber.SetupGated,
		usageplan.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

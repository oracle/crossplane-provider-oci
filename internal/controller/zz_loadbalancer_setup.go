/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	backend "github.com/oracle/provider-oci/internal/controller/loadbalancer/backend"
	backendset "github.com/oracle/provider-oci/internal/controller/loadbalancer/backendset"
	balancer "github.com/oracle/provider-oci/internal/controller/loadbalancer/balancer"
	balancerbackendset "github.com/oracle/provider-oci/internal/controller/loadbalancer/balancerbackendset"
	certificate "github.com/oracle/provider-oci/internal/controller/loadbalancer/certificate"
	lbhostname "github.com/oracle/provider-oci/internal/controller/loadbalancer/lbhostname"
	listener "github.com/oracle/provider-oci/internal/controller/loadbalancer/listener"
	loadbalancer "github.com/oracle/provider-oci/internal/controller/loadbalancer/loadbalancer"
	pathrouteset "github.com/oracle/provider-oci/internal/controller/loadbalancer/pathrouteset"
	routingpolicy "github.com/oracle/provider-oci/internal/controller/loadbalancer/routingpolicy"
	ruleset "github.com/oracle/provider-oci/internal/controller/loadbalancer/ruleset"
	sslciphersuite "github.com/oracle/provider-oci/internal/controller/loadbalancer/sslciphersuite"
)

// Setup_loadbalancer creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_loadbalancer(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backend.Setup,
		backendset.Setup,
		balancer.Setup,
		balancerbackendset.Setup,
		certificate.Setup,
		lbhostname.Setup,
		listener.Setup,
		loadbalancer.Setup,
		pathrouteset.Setup,
		routingpolicy.Setup,
		ruleset.Setup,
		sslciphersuite.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

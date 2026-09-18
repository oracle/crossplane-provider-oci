/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	backend "github.com/oracle/provider-oci/internal/controller/cluster/loadbalancer/backend"
	backendset "github.com/oracle/provider-oci/internal/controller/cluster/loadbalancer/backendset"
	certificate "github.com/oracle/provider-oci/internal/controller/cluster/loadbalancer/certificate"
	lbhostname "github.com/oracle/provider-oci/internal/controller/cluster/loadbalancer/lbhostname"
	listener "github.com/oracle/provider-oci/internal/controller/cluster/loadbalancer/listener"
	loadbalancer "github.com/oracle/provider-oci/internal/controller/cluster/loadbalancer/loadbalancer"
	pathrouteset "github.com/oracle/provider-oci/internal/controller/cluster/loadbalancer/pathrouteset"
	routingpolicy "github.com/oracle/provider-oci/internal/controller/cluster/loadbalancer/routingpolicy"
	ruleset "github.com/oracle/provider-oci/internal/controller/cluster/loadbalancer/ruleset"
	sslciphersuite "github.com/oracle/provider-oci/internal/controller/cluster/loadbalancer/sslciphersuite"
)

// Setup_loadbalancer creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_loadbalancer(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backend.Setup,
		backendset.Setup,
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

// SetupGated_loadbalancer creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_loadbalancer(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backend.SetupGated,
		backendset.SetupGated,
		certificate.SetupGated,
		lbhostname.SetupGated,
		listener.SetupGated,
		loadbalancer.SetupGated,
		pathrouteset.SetupGated,
		routingpolicy.SetupGated,
		ruleset.SetupGated,
		sslciphersuite.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

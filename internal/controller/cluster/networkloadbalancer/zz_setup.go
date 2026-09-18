/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	backend "github.com/oracle/provider-oci/internal/controller/cluster/networkloadbalancer/backend"
	backendset "github.com/oracle/provider-oci/internal/controller/cluster/networkloadbalancer/backendset"
	listener "github.com/oracle/provider-oci/internal/controller/cluster/networkloadbalancer/listener"
	networkloadbalancer "github.com/oracle/provider-oci/internal/controller/cluster/networkloadbalancer/networkloadbalancer"
	networkloadbalancersbackendsetsunified "github.com/oracle/provider-oci/internal/controller/cluster/networkloadbalancer/networkloadbalancersbackendsetsunified"
)

// Setup_networkloadbalancer creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_networkloadbalancer(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backend.Setup,
		backendset.Setup,
		listener.Setup,
		networkloadbalancer.Setup,
		networkloadbalancersbackendsetsunified.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_networkloadbalancer creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_networkloadbalancer(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backend.SetupGated,
		backendset.SetupGated,
		listener.SetupGated,
		networkloadbalancer.SetupGated,
		networkloadbalancersbackendsetsunified.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

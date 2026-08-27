/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	infrastructure "github.com/oracle/provider-oci/internal/controller/cluster/datacc/infrastructure"
	vmclusternetwork "github.com/oracle/provider-oci/internal/controller/cluster/datacc/vmclusternetwork"
	vminstance "github.com/oracle/provider-oci/internal/controller/cluster/datacc/vminstance"
)

// Setup_datacc creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_datacc(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		infrastructure.Setup,
		vmclusternetwork.Setup,
		vminstance.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_datacc creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_datacc(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		infrastructure.SetupGated,
		vmclusternetwork.SetupGated,
		vminstance.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	lustrefilesystem "github.com/oracle/provider-oci/internal/controller/cluster/lustrefilestorage/lustrefilesystem"
	objectstoragelink "github.com/oracle/provider-oci/internal/controller/cluster/lustrefilestorage/objectstoragelink"
)

// Setup_lustrefilestorage creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_lustrefilestorage(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		lustrefilesystem.Setup,
		objectstoragelink.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_lustrefilestorage creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_lustrefilestorage(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		lustrefilesystem.SetupGated,
		objectstoragelink.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	lustrefilesystem "github.com/oracle/provider-oci/internal/controller/lustrefilestorage/lustrefilesystem"
	objectstoragelink "github.com/oracle/provider-oci/internal/controller/lustrefilestorage/objectstoragelink"
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

/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	bootvolume "github.com/oracle/provider-oci/internal/controller/cluster/blockstorage/bootvolume"
	bootvolumebackup "github.com/oracle/provider-oci/internal/controller/cluster/blockstorage/bootvolumebackup"
	volume "github.com/oracle/provider-oci/internal/controller/cluster/blockstorage/volume"
	volumeattachment "github.com/oracle/provider-oci/internal/controller/cluster/blockstorage/volumeattachment"
	volumebackup "github.com/oracle/provider-oci/internal/controller/cluster/blockstorage/volumebackup"
	volumebackuppolicy "github.com/oracle/provider-oci/internal/controller/cluster/blockstorage/volumebackuppolicy"
	volumebackuppolicyassignment "github.com/oracle/provider-oci/internal/controller/cluster/blockstorage/volumebackuppolicyassignment"
	volumegroup "github.com/oracle/provider-oci/internal/controller/cluster/blockstorage/volumegroup"
	volumegroupbackup "github.com/oracle/provider-oci/internal/controller/cluster/blockstorage/volumegroupbackup"
)

// Setup_blockstorage creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_blockstorage(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bootvolume.Setup,
		bootvolumebackup.Setup,
		volume.Setup,
		volumeattachment.Setup,
		volumebackup.Setup,
		volumebackuppolicy.Setup,
		volumebackuppolicyassignment.Setup,
		volumegroup.Setup,
		volumegroupbackup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_blockstorage creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_blockstorage(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bootvolume.SetupGated,
		bootvolumebackup.SetupGated,
		volume.SetupGated,
		volumeattachment.SetupGated,
		volumebackup.SetupGated,
		volumebackuppolicy.SetupGated,
		volumebackuppolicyassignment.SetupGated,
		volumegroup.SetupGated,
		volumegroupbackup.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

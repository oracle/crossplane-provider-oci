/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	export "github.com/oracle/provider-oci/internal/controller/namespaced/filestorage/export"
	exportset "github.com/oracle/provider-oci/internal/controller/namespaced/filestorage/exportset"
	filesystem "github.com/oracle/provider-oci/internal/controller/namespaced/filestorage/filesystem"
	filesystemquotarule "github.com/oracle/provider-oci/internal/controller/namespaced/filestorage/filesystemquotarule"
	filesystemsnapshotpolicy "github.com/oracle/provider-oci/internal/controller/namespaced/filestorage/filesystemsnapshotpolicy"
	mounttarget "github.com/oracle/provider-oci/internal/controller/namespaced/filestorage/mounttarget"
	outboundconnector "github.com/oracle/provider-oci/internal/controller/namespaced/filestorage/outboundconnector"
	replication "github.com/oracle/provider-oci/internal/controller/namespaced/filestorage/replication"
	snapshot "github.com/oracle/provider-oci/internal/controller/namespaced/filestorage/snapshot"
)

// Setup_filestorage creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_filestorage(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		export.Setup,
		exportset.Setup,
		filesystem.Setup,
		filesystemquotarule.Setup,
		filesystemsnapshotpolicy.Setup,
		mounttarget.Setup,
		outboundconnector.Setup,
		replication.Setup,
		snapshot.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_filestorage creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_filestorage(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		export.SetupGated,
		exportset.SetupGated,
		filesystem.SetupGated,
		filesystemquotarule.SetupGated,
		filesystemsnapshotpolicy.SetupGated,
		mounttarget.SetupGated,
		outboundconnector.SetupGated,
		replication.SetupGated,
		snapshot.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	export "github.com/oracle/provider-oci/internal/controller/filestorage/export"
	exportset "github.com/oracle/provider-oci/internal/controller/filestorage/exportset"
	filesystem "github.com/oracle/provider-oci/internal/controller/filestorage/filesystem"
	filesystemquotarule "github.com/oracle/provider-oci/internal/controller/filestorage/filesystemquotarule"
	filesystemsnapshotpolicy "github.com/oracle/provider-oci/internal/controller/filestorage/filesystemsnapshotpolicy"
	mounttarget "github.com/oracle/provider-oci/internal/controller/filestorage/mounttarget"
	outboundconnector "github.com/oracle/provider-oci/internal/controller/filestorage/outboundconnector"
	replication "github.com/oracle/provider-oci/internal/controller/filestorage/replication"
	snapshot "github.com/oracle/provider-oci/internal/controller/filestorage/snapshot"
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

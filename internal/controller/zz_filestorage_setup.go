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
	mounttarget "github.com/oracle/provider-oci/internal/controller/filestorage/mounttarget"
	replication "github.com/oracle/provider-oci/internal/controller/filestorage/replication"
	snapshot "github.com/oracle/provider-oci/internal/controller/filestorage/snapshot"
	storagefilesystemquotarule "github.com/oracle/provider-oci/internal/controller/filestorage/storagefilesystemquotarule"
	storagefilesystemsnapshotpolicy "github.com/oracle/provider-oci/internal/controller/filestorage/storagefilesystemsnapshotpolicy"
	storageoutboundconnector "github.com/oracle/provider-oci/internal/controller/filestorage/storageoutboundconnector"
)

// Setup_filestorage creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_filestorage(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		export.Setup,
		exportset.Setup,
		filesystem.Setup,
		mounttarget.Setup,
		replication.Setup,
		snapshot.Setup,
		storagefilesystemquotarule.Setup,
		storagefilesystemsnapshotpolicy.Setup,
		storageoutboundconnector.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	multicloudresourcediscovery "github.com/oracle/provider-oci/internal/controller/dbmulticloud/multicloudresourcediscovery"
	oracledbazureblobcontainer "github.com/oracle/provider-oci/internal/controller/dbmulticloud/oracledbazureblobcontainer"
	oracledbazureblobmount "github.com/oracle/provider-oci/internal/controller/dbmulticloud/oracledbazureblobmount"
	oracledbazureconnector "github.com/oracle/provider-oci/internal/controller/dbmulticloud/oracledbazureconnector"
	oracledbazurevault "github.com/oracle/provider-oci/internal/controller/dbmulticloud/oracledbazurevault"
	oracledbazurevaultassociation "github.com/oracle/provider-oci/internal/controller/dbmulticloud/oracledbazurevaultassociation"
)

// Setup_dbmulticloud creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dbmulticloud(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		multicloudresourcediscovery.Setup,
		oracledbazureblobcontainer.Setup,
		oracledbazureblobmount.Setup,
		oracledbazureconnector.Setup,
		oracledbazurevault.Setup,
		oracledbazurevaultassociation.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

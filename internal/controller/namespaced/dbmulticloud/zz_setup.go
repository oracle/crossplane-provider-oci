/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	multicloudresourcediscovery "github.com/oracle/provider-oci/internal/controller/namespaced/dbmulticloud/multicloudresourcediscovery"
	oracledbawsidentityconnector "github.com/oracle/provider-oci/internal/controller/namespaced/dbmulticloud/oracledbawsidentityconnector"
	oracledbawskey "github.com/oracle/provider-oci/internal/controller/namespaced/dbmulticloud/oracledbawskey"
	oracledbazureblobcontainer "github.com/oracle/provider-oci/internal/controller/namespaced/dbmulticloud/oracledbazureblobcontainer"
	oracledbazureblobmount "github.com/oracle/provider-oci/internal/controller/namespaced/dbmulticloud/oracledbazureblobmount"
	oracledbazureconnector "github.com/oracle/provider-oci/internal/controller/namespaced/dbmulticloud/oracledbazureconnector"
	oracledbazurevault "github.com/oracle/provider-oci/internal/controller/namespaced/dbmulticloud/oracledbazurevault"
	oracledbazurevaultassociation "github.com/oracle/provider-oci/internal/controller/namespaced/dbmulticloud/oracledbazurevaultassociation"
	oracledbgcpidentityconnector "github.com/oracle/provider-oci/internal/controller/namespaced/dbmulticloud/oracledbgcpidentityconnector"
	oracledbgcpkeyring "github.com/oracle/provider-oci/internal/controller/namespaced/dbmulticloud/oracledbgcpkeyring"
)

// Setup_dbmulticloud creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dbmulticloud(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		multicloudresourcediscovery.Setup,
		oracledbawsidentityconnector.Setup,
		oracledbawskey.Setup,
		oracledbazureblobcontainer.Setup,
		oracledbazureblobmount.Setup,
		oracledbazureconnector.Setup,
		oracledbazurevault.Setup,
		oracledbazurevaultassociation.Setup,
		oracledbgcpidentityconnector.Setup,
		oracledbgcpkeyring.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_dbmulticloud creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_dbmulticloud(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		multicloudresourcediscovery.SetupGated,
		oracledbawsidentityconnector.SetupGated,
		oracledbawskey.SetupGated,
		oracledbazureblobcontainer.SetupGated,
		oracledbazureblobmount.SetupGated,
		oracledbazureconnector.SetupGated,
		oracledbazurevault.SetupGated,
		oracledbazurevaultassociation.SetupGated,
		oracledbgcpidentityconnector.SetupGated,
		oracledbgcpkeyring.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

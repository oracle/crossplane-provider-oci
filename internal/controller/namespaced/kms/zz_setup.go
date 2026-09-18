/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	ekmsprivateendpoint "github.com/oracle/provider-oci/internal/controller/namespaced/kms/ekmsprivateendpoint"
	encrypteddata "github.com/oracle/provider-oci/internal/controller/namespaced/kms/encrypteddata"
	generatedkey "github.com/oracle/provider-oci/internal/controller/namespaced/kms/generatedkey"
	key "github.com/oracle/provider-oci/internal/controller/namespaced/kms/key"
	keyversion "github.com/oracle/provider-oci/internal/controller/namespaced/kms/keyversion"
	sign "github.com/oracle/provider-oci/internal/controller/namespaced/kms/sign"
	vault "github.com/oracle/provider-oci/internal/controller/namespaced/kms/vault"
	vaultreplication "github.com/oracle/provider-oci/internal/controller/namespaced/kms/vaultreplication"
	verify "github.com/oracle/provider-oci/internal/controller/namespaced/kms/verify"
)

// Setup_kms creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_kms(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		ekmsprivateendpoint.Setup,
		encrypteddata.Setup,
		generatedkey.Setup,
		key.Setup,
		keyversion.Setup,
		sign.Setup,
		vault.Setup,
		vaultreplication.Setup,
		verify.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_kms creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_kms(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		ekmsprivateendpoint.SetupGated,
		encrypteddata.SetupGated,
		generatedkey.SetupGated,
		key.SetupGated,
		keyversion.SetupGated,
		sign.SetupGated,
		vault.SetupGated,
		vaultreplication.SetupGated,
		verify.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

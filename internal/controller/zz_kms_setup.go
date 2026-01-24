/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	ekmsprivateendpoint "github.com/oracle/provider-oci/internal/controller/kms/ekmsprivateendpoint"
	encrypteddata "github.com/oracle/provider-oci/internal/controller/kms/encrypteddata"
	generatedkey "github.com/oracle/provider-oci/internal/controller/kms/generatedkey"
	key "github.com/oracle/provider-oci/internal/controller/kms/key"
	keyversion "github.com/oracle/provider-oci/internal/controller/kms/keyversion"
	sign "github.com/oracle/provider-oci/internal/controller/kms/sign"
	vault "github.com/oracle/provider-oci/internal/controller/kms/vault"
	vaultreplication "github.com/oracle/provider-oci/internal/controller/kms/vaultreplication"
	verify "github.com/oracle/provider-oci/internal/controller/kms/verify"
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

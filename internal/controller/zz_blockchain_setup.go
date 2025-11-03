/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	osn "github.com/oracle/provider-oci/internal/controller/blockchain/osn"
	peer "github.com/oracle/provider-oci/internal/controller/blockchain/peer"
	platform "github.com/oracle/provider-oci/internal/controller/blockchain/platform"
)

// Setup_blockchain creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_blockchain(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		osn.Setup,
		peer.Setup,
		platform.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

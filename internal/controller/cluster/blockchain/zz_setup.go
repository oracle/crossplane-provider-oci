/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	blockchainplatform "github.com/oracle/provider-oci/internal/controller/cluster/blockchain/blockchainplatform"
	osn "github.com/oracle/provider-oci/internal/controller/cluster/blockchain/osn"
	peer "github.com/oracle/provider-oci/internal/controller/cluster/blockchain/peer"
)

// Setup_blockchain creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_blockchain(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		blockchainplatform.Setup,
		osn.Setup,
		peer.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_blockchain creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_blockchain(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		blockchainplatform.SetupGated,
		osn.SetupGated,
		peer.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	batchcontext "github.com/oracle/provider-oci/internal/controller/namespaced/batch/batchcontext"
	batchjobpool "github.com/oracle/provider-oci/internal/controller/namespaced/batch/batchjobpool"
	batchtaskenvironment "github.com/oracle/provider-oci/internal/controller/namespaced/batch/batchtaskenvironment"
	batchtaskprofile "github.com/oracle/provider-oci/internal/controller/namespaced/batch/batchtaskprofile"
)

// Setup_batch creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_batch(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		batchcontext.Setup,
		batchjobpool.Setup,
		batchtaskenvironment.Setup,
		batchtaskprofile.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_batch creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_batch(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		batchcontext.SetupGated,
		batchjobpool.SetupGated,
		batchtaskenvironment.SetupGated,
		batchtaskprofile.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	byoasn "github.com/oracle/provider-oci/internal/controller/core/byoasn"
	listingresourceversionagreement "github.com/oracle/provider-oci/internal/controller/core/listingresourceversionagreement"
	virtualnetwork "github.com/oracle/provider-oci/internal/controller/core/virtualnetwork"
)

// Setup_core creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_core(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		byoasn.Setup,
		listingresourceversionagreement.Setup,
		virtualnetwork.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

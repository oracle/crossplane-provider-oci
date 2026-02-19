/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	acceptedagreement "github.com/oracle/provider-oci/internal/controller/marketplace/acceptedagreement"
	listingpackageagreement "github.com/oracle/provider-oci/internal/controller/marketplace/listingpackageagreement"
	marketplaceexternalattestedmetadata "github.com/oracle/provider-oci/internal/controller/marketplace/marketplaceexternalattestedmetadata"
	publication "github.com/oracle/provider-oci/internal/controller/marketplace/publication"
)

// Setup_marketplace creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_marketplace(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		acceptedagreement.Setup,
		listingpackageagreement.Setup,
		marketplaceexternalattestedmetadata.Setup,
		publication.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

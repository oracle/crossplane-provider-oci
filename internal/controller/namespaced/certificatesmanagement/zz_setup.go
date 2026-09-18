/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	cabundle "github.com/oracle/provider-oci/internal/controller/namespaced/certificatesmanagement/cabundle"
	certificate "github.com/oracle/provider-oci/internal/controller/namespaced/certificatesmanagement/certificate"
	certificateauthority "github.com/oracle/provider-oci/internal/controller/namespaced/certificatesmanagement/certificateauthority"
)

// Setup_certificatesmanagement creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_certificatesmanagement(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cabundle.Setup,
		certificate.Setup,
		certificateauthority.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_certificatesmanagement creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_certificatesmanagement(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cabundle.SetupGated,
		certificate.SetupGated,
		certificateauthority.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

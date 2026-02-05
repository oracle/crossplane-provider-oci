/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	cabundle "github.com/oracle/provider-oci/internal/controller/certificatesmanagement/cabundle"
	certificate "github.com/oracle/provider-oci/internal/controller/certificatesmanagement/certificate"
	certificateauthority "github.com/oracle/provider-oci/internal/controller/certificatesmanagement/certificateauthority"
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

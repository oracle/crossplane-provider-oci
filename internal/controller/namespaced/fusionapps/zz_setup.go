/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	fusionenvironment "github.com/oracle/provider-oci/internal/controller/namespaced/fusionapps/fusionenvironment"
	fusionenvironmentadminuser "github.com/oracle/provider-oci/internal/controller/namespaced/fusionapps/fusionenvironmentadminuser"
	fusionenvironmentdatamaskingactivity "github.com/oracle/provider-oci/internal/controller/namespaced/fusionapps/fusionenvironmentdatamaskingactivity"
	fusionenvironmentfamily "github.com/oracle/provider-oci/internal/controller/namespaced/fusionapps/fusionenvironmentfamily"
	fusionenvironmentrefreshactivity "github.com/oracle/provider-oci/internal/controller/namespaced/fusionapps/fusionenvironmentrefreshactivity"
	fusionenvironmentserviceattachment "github.com/oracle/provider-oci/internal/controller/namespaced/fusionapps/fusionenvironmentserviceattachment"
)

// Setup_fusionapps creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_fusionapps(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		fusionenvironment.Setup,
		fusionenvironmentadminuser.Setup,
		fusionenvironmentdatamaskingactivity.Setup,
		fusionenvironmentfamily.Setup,
		fusionenvironmentrefreshactivity.Setup,
		fusionenvironmentserviceattachment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_fusionapps creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_fusionapps(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		fusionenvironment.SetupGated,
		fusionenvironmentadminuser.SetupGated,
		fusionenvironmentdatamaskingactivity.SetupGated,
		fusionenvironmentfamily.SetupGated,
		fusionenvironmentrefreshactivity.SetupGated,
		fusionenvironmentserviceattachment.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

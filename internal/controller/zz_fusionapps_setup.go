/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	appsfusionenvironment "github.com/oracle/provider-oci/internal/controller/fusionapps/appsfusionenvironment"
	appsfusionenvironmentadminuser "github.com/oracle/provider-oci/internal/controller/fusionapps/appsfusionenvironmentadminuser"
	appsfusionenvironmentfamily "github.com/oracle/provider-oci/internal/controller/fusionapps/appsfusionenvironmentfamily"
	appsfusionenvironmentrefreshactivity "github.com/oracle/provider-oci/internal/controller/fusionapps/appsfusionenvironmentrefreshactivity"
	appsfusionenvironmentserviceattachment "github.com/oracle/provider-oci/internal/controller/fusionapps/appsfusionenvironmentserviceattachment"
)

// Setup_fusionapps creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_fusionapps(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		appsfusionenvironment.Setup,
		appsfusionenvironmentadminuser.Setup,
		appsfusionenvironmentfamily.Setup,
		appsfusionenvironmentrefreshactivity.Setup,
		appsfusionenvironmentserviceattachment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

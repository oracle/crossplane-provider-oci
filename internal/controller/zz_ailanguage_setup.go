/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	documentmodel "github.com/oracle/provider-oci/internal/controller/ailanguage/documentmodel"
	documentprocessorjob "github.com/oracle/provider-oci/internal/controller/ailanguage/documentprocessorjob"
	documentproject "github.com/oracle/provider-oci/internal/controller/ailanguage/documentproject"
	languageendpoint "github.com/oracle/provider-oci/internal/controller/ailanguage/languageendpoint"
	languagemodel "github.com/oracle/provider-oci/internal/controller/ailanguage/languagemodel"
	languageproject "github.com/oracle/provider-oci/internal/controller/ailanguage/languageproject"
	visionmodel "github.com/oracle/provider-oci/internal/controller/ailanguage/visionmodel"
	visionproject "github.com/oracle/provider-oci/internal/controller/ailanguage/visionproject"
)

// Setup_ailanguage creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_ailanguage(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		documentmodel.Setup,
		documentprocessorjob.Setup,
		documentproject.Setup,
		languageendpoint.Setup,
		languagemodel.Setup,
		languageproject.Setup,
		visionmodel.Setup,
		visionproject.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

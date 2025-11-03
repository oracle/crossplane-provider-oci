/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	dkim "github.com/oracle/provider-oci/internal/controller/emaildataplane/dkim"
	emaildataplanedomain "github.com/oracle/provider-oci/internal/controller/emaildataplane/emaildataplanedomain"
	returnpath "github.com/oracle/provider-oci/internal/controller/emaildataplane/returnpath"
	sender "github.com/oracle/provider-oci/internal/controller/emaildataplane/sender"
	suppression "github.com/oracle/provider-oci/internal/controller/emaildataplane/suppression"
)

// Setup_emaildataplane creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_emaildataplane(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		dkim.Setup,
		emaildataplanedomain.Setup,
		returnpath.Setup,
		sender.Setup,
		suppression.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

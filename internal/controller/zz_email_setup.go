/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	dkim "github.com/oracle/provider-oci/internal/controller/email/dkim"
	emaildomain "github.com/oracle/provider-oci/internal/controller/email/emaildomain"
	emailreturnpath "github.com/oracle/provider-oci/internal/controller/email/emailreturnpath"
	sender "github.com/oracle/provider-oci/internal/controller/email/sender"
	suppression "github.com/oracle/provider-oci/internal/controller/email/suppression"
)

// Setup_email creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_email(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		dkim.Setup,
		emaildomain.Setup,
		emailreturnpath.Setup,
		sender.Setup,
		suppression.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

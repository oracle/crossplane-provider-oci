/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	dkim "github.com/oracle/provider-oci/internal/controller/cluster/email/dkim"
	emaildomain "github.com/oracle/provider-oci/internal/controller/cluster/email/emaildomain"
	emailippool "github.com/oracle/provider-oci/internal/controller/cluster/email/emailippool"
	emailreturnpath "github.com/oracle/provider-oci/internal/controller/cluster/email/emailreturnpath"
	sender "github.com/oracle/provider-oci/internal/controller/cluster/email/sender"
	suppression "github.com/oracle/provider-oci/internal/controller/cluster/email/suppression"
)

// Setup_email creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_email(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		dkim.Setup,
		emaildomain.Setup,
		emailippool.Setup,
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

// SetupGated_email creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_email(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		dkim.SetupGated,
		emaildomain.SetupGated,
		emailippool.SetupGated,
		emailreturnpath.SetupGated,
		sender.SetupGated,
		suppression.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

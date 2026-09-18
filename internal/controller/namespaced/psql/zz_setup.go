/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	psqlbackup "github.com/oracle/provider-oci/internal/controller/namespaced/psql/psqlbackup"
	psqlconfiguration "github.com/oracle/provider-oci/internal/controller/namespaced/psql/psqlconfiguration"
	psqldbsystem "github.com/oracle/provider-oci/internal/controller/namespaced/psql/psqldbsystem"
)

// Setup_psql creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_psql(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		psqlbackup.Setup,
		psqlconfiguration.Setup,
		psqldbsystem.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_psql creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_psql(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		psqlbackup.SetupGated,
		psqlconfiguration.SetupGated,
		psqldbsystem.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

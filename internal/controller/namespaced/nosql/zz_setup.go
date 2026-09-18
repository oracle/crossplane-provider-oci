/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	configuration "github.com/oracle/provider-oci/internal/controller/namespaced/nosql/configuration"
	index "github.com/oracle/provider-oci/internal/controller/namespaced/nosql/index"
	table "github.com/oracle/provider-oci/internal/controller/namespaced/nosql/table"
	tablereplica "github.com/oracle/provider-oci/internal/controller/namespaced/nosql/tablereplica"
)

// Setup_nosql creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_nosql(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		configuration.Setup,
		index.Setup,
		table.Setup,
		tablereplica.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_nosql creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_nosql(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		configuration.SetupGated,
		index.SetupGated,
		table.SetupGated,
		tablereplica.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

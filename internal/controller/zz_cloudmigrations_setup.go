/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	migration "github.com/oracle/provider-oci/internal/controller/cloudmigrations/migration"
	migrationasset "github.com/oracle/provider-oci/internal/controller/cloudmigrations/migrationasset"
	migrationplan "github.com/oracle/provider-oci/internal/controller/cloudmigrations/migrationplan"
	replicationschedule "github.com/oracle/provider-oci/internal/controller/cloudmigrations/replicationschedule"
	targetasset "github.com/oracle/provider-oci/internal/controller/cloudmigrations/targetasset"
)

// Setup_cloudmigrations creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cloudmigrations(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		migration.Setup,
		migrationasset.Setup,
		migrationplan.Setup,
		replicationschedule.Setup,
		targetasset.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

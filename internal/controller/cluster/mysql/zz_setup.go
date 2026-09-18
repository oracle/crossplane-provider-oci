/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	bluegreendeployment "github.com/oracle/provider-oci/internal/controller/cluster/mysql/bluegreendeployment"
	mysqlbackup "github.com/oracle/provider-oci/internal/controller/cluster/mysql/mysqlbackup"
	mysqlchannel "github.com/oracle/provider-oci/internal/controller/cluster/mysql/mysqlchannel"
	mysqlconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/mysql/mysqlconfiguration"
	mysqldbsystem "github.com/oracle/provider-oci/internal/controller/cluster/mysql/mysqldbsystem"
	mysqlheatwavecluster "github.com/oracle/provider-oci/internal/controller/cluster/mysql/mysqlheatwavecluster"
	mysqlreplica "github.com/oracle/provider-oci/internal/controller/cluster/mysql/mysqlreplica"
)

// Setup_mysql creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_mysql(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bluegreendeployment.Setup,
		mysqlbackup.Setup,
		mysqlchannel.Setup,
		mysqlconfiguration.Setup,
		mysqldbsystem.Setup,
		mysqlheatwavecluster.Setup,
		mysqlreplica.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_mysql creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_mysql(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bluegreendeployment.SetupGated,
		mysqlbackup.SetupGated,
		mysqlchannel.SetupGated,
		mysqlconfiguration.SetupGated,
		mysqldbsystem.SetupGated,
		mysqlheatwavecluster.SetupGated,
		mysqlreplica.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

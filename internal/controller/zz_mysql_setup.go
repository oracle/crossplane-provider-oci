/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	channel "github.com/oracle/provider-oci/internal/controller/mysql/channel"
	heatwavecluster "github.com/oracle/provider-oci/internal/controller/mysql/heatwavecluster"
	mysqlbackup "github.com/oracle/provider-oci/internal/controller/mysql/mysqlbackup"
	mysqlconfiguration "github.com/oracle/provider-oci/internal/controller/mysql/mysqlconfiguration"
	mysqldbsystem "github.com/oracle/provider-oci/internal/controller/mysql/mysqldbsystem"
	replica "github.com/oracle/provider-oci/internal/controller/mysql/replica"
)

// Setup_mysql creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_mysql(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		channel.Setup,
		heatwavecluster.Setup,
		mysqlbackup.Setup,
		mysqlconfiguration.Setup,
		mysqldbsystem.Setup,
		replica.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

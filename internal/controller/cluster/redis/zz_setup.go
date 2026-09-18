/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	ocicachebackup "github.com/oracle/provider-oci/internal/controller/cluster/redis/ocicachebackup"
	ocicachebackupexporttoobjectstorage "github.com/oracle/provider-oci/internal/controller/cluster/redis/ocicachebackupexporttoobjectstorage"
	ocicacheconfigset "github.com/oracle/provider-oci/internal/controller/cluster/redis/ocicacheconfigset"
	ocicacheconfigsetlistassociatedocicachecluster "github.com/oracle/provider-oci/internal/controller/cluster/redis/ocicacheconfigsetlistassociatedocicachecluster"
	ocicacheuser "github.com/oracle/provider-oci/internal/controller/cluster/redis/ocicacheuser"
	ocicacheusergetrediscluster "github.com/oracle/provider-oci/internal/controller/cluster/redis/ocicacheusergetrediscluster"
	rediscluster "github.com/oracle/provider-oci/internal/controller/cluster/redis/rediscluster"
	redisclusterattachocicacheuser "github.com/oracle/provider-oci/internal/controller/cluster/redis/redisclusterattachocicacheuser"
	redisclustercreateidentitytoken "github.com/oracle/provider-oci/internal/controller/cluster/redis/redisclustercreateidentitytoken"
	redisclusterdetachocicacheuser "github.com/oracle/provider-oci/internal/controller/cluster/redis/redisclusterdetachocicacheuser"
	redisclustergetocicacheuser "github.com/oracle/provider-oci/internal/controller/cluster/redis/redisclustergetocicacheuser"
)

// Setup_redis creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_redis(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		ocicachebackup.Setup,
		ocicachebackupexporttoobjectstorage.Setup,
		ocicacheconfigset.Setup,
		ocicacheconfigsetlistassociatedocicachecluster.Setup,
		ocicacheuser.Setup,
		ocicacheusergetrediscluster.Setup,
		rediscluster.Setup,
		redisclusterattachocicacheuser.Setup,
		redisclustercreateidentitytoken.Setup,
		redisclusterdetachocicacheuser.Setup,
		redisclustergetocicacheuser.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_redis creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_redis(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		ocicachebackup.SetupGated,
		ocicachebackupexporttoobjectstorage.SetupGated,
		ocicacheconfigset.SetupGated,
		ocicacheconfigsetlistassociatedocicachecluster.SetupGated,
		ocicacheuser.SetupGated,
		ocicacheusergetrediscluster.SetupGated,
		rediscluster.SetupGated,
		redisclusterattachocicacheuser.SetupGated,
		redisclustercreateidentitytoken.SetupGated,
		redisclusterdetachocicacheuser.SetupGated,
		redisclustergetocicacheuser.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

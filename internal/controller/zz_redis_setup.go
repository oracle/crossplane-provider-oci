/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	clusterattachocicacheuser "github.com/oracle/provider-oci/internal/controller/redis/clusterattachocicacheuser"
	clustercreateidentitytoken "github.com/oracle/provider-oci/internal/controller/redis/clustercreateidentitytoken"
	clusterdetachocicacheuser "github.com/oracle/provider-oci/internal/controller/redis/clusterdetachocicacheuser"
	clustergetocicacheuser "github.com/oracle/provider-oci/internal/controller/redis/clustergetocicacheuser"
	ocicacheuser "github.com/oracle/provider-oci/internal/controller/redis/ocicacheuser"
	ocicacheusergetrediscluster "github.com/oracle/provider-oci/internal/controller/redis/ocicacheusergetrediscluster"
	rediscluster "github.com/oracle/provider-oci/internal/controller/redis/rediscluster"
)

// Setup_redis creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_redis(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		clusterattachocicacheuser.Setup,
		clustercreateidentitytoken.Setup,
		clusterdetachocicacheuser.Setup,
		clustergetocicacheuser.Setup,
		ocicacheuser.Setup,
		ocicacheusergetrediscluster.Setup,
		rediscluster.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	kafkacluster "github.com/oracle/provider-oci/internal/controller/managedkafka/kafkacluster"
	kafkaclusterconfig "github.com/oracle/provider-oci/internal/controller/managedkafka/kafkaclusterconfig"
	kafkaclustersuperusersmanagement "github.com/oracle/provider-oci/internal/controller/managedkafka/kafkaclustersuperusersmanagement"
)

// Setup_managedkafka creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_managedkafka(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		kafkacluster.Setup,
		kafkaclusterconfig.Setup,
		kafkaclustersuperusersmanagement.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

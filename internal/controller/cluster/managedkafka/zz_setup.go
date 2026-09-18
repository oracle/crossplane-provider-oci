/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	kafkacluster "github.com/oracle/provider-oci/internal/controller/cluster/managedkafka/kafkacluster"
	kafkaclusteraddon "github.com/oracle/provider-oci/internal/controller/cluster/managedkafka/kafkaclusteraddon"
	kafkaclusterconfig "github.com/oracle/provider-oci/internal/controller/cluster/managedkafka/kafkaclusterconfig"
	kafkaclustersuperusersmanagement "github.com/oracle/provider-oci/internal/controller/cluster/managedkafka/kafkaclustersuperusersmanagement"
)

// Setup_managedkafka creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_managedkafka(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		kafkacluster.Setup,
		kafkaclusteraddon.Setup,
		kafkaclusterconfig.Setup,
		kafkaclustersuperusersmanagement.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_managedkafka creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_managedkafka(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		kafkacluster.SetupGated,
		kafkaclusteraddon.SetupGated,
		kafkaclusterconfig.SetupGated,
		kafkaclustersuperusersmanagement.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

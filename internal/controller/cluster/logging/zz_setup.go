/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	log "github.com/oracle/provider-oci/internal/controller/cluster/logging/log"
	loggroup "github.com/oracle/provider-oci/internal/controller/cluster/logging/loggroup"
	logsavedsearch "github.com/oracle/provider-oci/internal/controller/cluster/logging/logsavedsearch"
	unifiedagentconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/logging/unifiedagentconfiguration"
)

// Setup_logging creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_logging(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		log.Setup,
		loggroup.Setup,
		logsavedsearch.Setup,
		unifiedagentconfiguration.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_logging creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_logging(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		log.SetupGated,
		loggroup.SetupGated,
		logsavedsearch.SetupGated,
		unifiedagentconfiguration.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

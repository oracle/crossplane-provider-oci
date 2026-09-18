/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	analyzeapplicationsconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/jmsutils/analyzeapplicationsconfiguration"
	subscriptionacknowledgmentconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/jmsutils/subscriptionacknowledgmentconfiguration"
)

// Setup_jmsutils creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_jmsutils(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		analyzeapplicationsconfiguration.Setup,
		subscriptionacknowledgmentconfiguration.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_jmsutils creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_jmsutils(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		analyzeapplicationsconfiguration.SetupGated,
		subscriptionacknowledgmentconfiguration.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

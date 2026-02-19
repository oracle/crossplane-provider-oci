/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	knowledgebase "github.com/oracle/provider-oci/internal/controller/adm/knowledgebase"
	remediationrecipe "github.com/oracle/provider-oci/internal/controller/adm/remediationrecipe"
	remediationrun "github.com/oracle/provider-oci/internal/controller/adm/remediationrun"
	vulnerabilityaudit "github.com/oracle/provider-oci/internal/controller/adm/vulnerabilityaudit"
)

// Setup_adm creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_adm(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		knowledgebase.Setup,
		remediationrecipe.Setup,
		remediationrun.Setup,
		vulnerabilityaudit.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

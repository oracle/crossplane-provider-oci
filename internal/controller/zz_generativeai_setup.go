/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	aiagent "github.com/oracle/provider-oci/internal/controller/generativeai/aiagent"
	aiagentendpoint "github.com/oracle/provider-oci/internal/controller/generativeai/aiagentendpoint"
	aiagentknowledgebase "github.com/oracle/provider-oci/internal/controller/generativeai/aiagentknowledgebase"
	aiagenttool "github.com/oracle/provider-oci/internal/controller/generativeai/aiagenttool"
	aidedicatedaicluster "github.com/oracle/provider-oci/internal/controller/generativeai/aidedicatedaicluster"
	aiendpoint "github.com/oracle/provider-oci/internal/controller/generativeai/aiendpoint"
	aimodel "github.com/oracle/provider-oci/internal/controller/generativeai/aimodel"
)

// Setup_generativeai creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_generativeai(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		aiagent.Setup,
		aiagentendpoint.Setup,
		aiagentknowledgebase.Setup,
		aiagenttool.Setup,
		aidedicatedaicluster.Setup,
		aiendpoint.Setup,
		aimodel.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

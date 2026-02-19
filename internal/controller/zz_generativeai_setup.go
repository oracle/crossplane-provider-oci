/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	agentagent "github.com/oracle/provider-oci/internal/controller/generativeai/agentagent"
	agentagentendpoint "github.com/oracle/provider-oci/internal/controller/generativeai/agentagentendpoint"
	agentdataingestionjob "github.com/oracle/provider-oci/internal/controller/generativeai/agentdataingestionjob"
	agentdatasource "github.com/oracle/provider-oci/internal/controller/generativeai/agentdatasource"
	agentknowledgebase "github.com/oracle/provider-oci/internal/controller/generativeai/agentknowledgebase"
	agenttool "github.com/oracle/provider-oci/internal/controller/generativeai/agenttool"
	dedicatedaicluster "github.com/oracle/provider-oci/internal/controller/generativeai/dedicatedaicluster"
	endpoint "github.com/oracle/provider-oci/internal/controller/generativeai/endpoint"
	generativeaiprivateendpoint "github.com/oracle/provider-oci/internal/controller/generativeai/generativeaiprivateendpoint"
	model "github.com/oracle/provider-oci/internal/controller/generativeai/model"
)

// Setup_generativeai creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_generativeai(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		agentagent.Setup,
		agentagentendpoint.Setup,
		agentdataingestionjob.Setup,
		agentdatasource.Setup,
		agentknowledgebase.Setup,
		agenttool.Setup,
		dedicatedaicluster.Setup,
		endpoint.Setup,
		generativeaiprivateendpoint.Setup,
		model.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

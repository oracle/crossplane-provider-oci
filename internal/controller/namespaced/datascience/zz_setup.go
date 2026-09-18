/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	computetarget "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/computetarget"
	job "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/job"
	jobrun "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/jobrun"
	mlapplication "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/mlapplication"
	mlapplicationimplementation "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/mlapplicationimplementation"
	mlapplicationinstance "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/mlapplicationinstance"
	model "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/model"
	modelartifactexport "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/modelartifactexport"
	modelartifactimport "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/modelartifactimport"
	modelcustommetadataartifact "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/modelcustommetadataartifact"
	modeldefinedmetadataartifact "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/modeldefinedmetadataartifact"
	modeldeployment "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/modeldeployment"
	modelgroup "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/modelgroup"
	modelgroupartifact "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/modelgroupartifact"
	modelgroupversionhistory "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/modelgroupversionhistory"
	modelprovenance "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/modelprovenance"
	modelversionset "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/modelversionset"
	notebooksession "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/notebooksession"
	pipeline "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/pipeline"
	pipelinerun "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/pipelinerun"
	privateendpoint "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/privateendpoint"
	project "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/project"
	schedule "github.com/oracle/provider-oci/internal/controller/namespaced/datascience/schedule"
)

// Setup_datascience creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_datascience(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		computetarget.Setup,
		job.Setup,
		jobrun.Setup,
		mlapplication.Setup,
		mlapplicationimplementation.Setup,
		mlapplicationinstance.Setup,
		model.Setup,
		modelartifactexport.Setup,
		modelartifactimport.Setup,
		modelcustommetadataartifact.Setup,
		modeldefinedmetadataartifact.Setup,
		modeldeployment.Setup,
		modelgroup.Setup,
		modelgroupartifact.Setup,
		modelgroupversionhistory.Setup,
		modelprovenance.Setup,
		modelversionset.Setup,
		notebooksession.Setup,
		pipeline.Setup,
		pipelinerun.Setup,
		privateendpoint.Setup,
		project.Setup,
		schedule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_datascience creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_datascience(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		computetarget.SetupGated,
		job.SetupGated,
		jobrun.SetupGated,
		mlapplication.SetupGated,
		mlapplicationimplementation.SetupGated,
		mlapplicationinstance.SetupGated,
		model.SetupGated,
		modelartifactexport.SetupGated,
		modelartifactimport.SetupGated,
		modelcustommetadataartifact.SetupGated,
		modeldefinedmetadataartifact.SetupGated,
		modeldeployment.SetupGated,
		modelgroup.SetupGated,
		modelgroupartifact.SetupGated,
		modelgroupversionhistory.SetupGated,
		modelprovenance.SetupGated,
		modelversionset.SetupGated,
		notebooksession.SetupGated,
		pipeline.SetupGated,
		pipelinerun.SetupGated,
		privateendpoint.SetupGated,
		project.SetupGated,
		schedule.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

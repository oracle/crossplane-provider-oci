/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	job "github.com/oracle/provider-oci/internal/controller/datascience/job"
	jobrun "github.com/oracle/provider-oci/internal/controller/datascience/jobrun"
	mlapplication "github.com/oracle/provider-oci/internal/controller/datascience/mlapplication"
	mlapplicationimplementation "github.com/oracle/provider-oci/internal/controller/datascience/mlapplicationimplementation"
	mlapplicationinstance "github.com/oracle/provider-oci/internal/controller/datascience/mlapplicationinstance"
	model "github.com/oracle/provider-oci/internal/controller/datascience/model"
	modelartifactexport "github.com/oracle/provider-oci/internal/controller/datascience/modelartifactexport"
	modelartifactimport "github.com/oracle/provider-oci/internal/controller/datascience/modelartifactimport"
	modelcustommetadataartifact "github.com/oracle/provider-oci/internal/controller/datascience/modelcustommetadataartifact"
	modeldefinedmetadataartifact "github.com/oracle/provider-oci/internal/controller/datascience/modeldefinedmetadataartifact"
	modeldeployment "github.com/oracle/provider-oci/internal/controller/datascience/modeldeployment"
	modelgroup "github.com/oracle/provider-oci/internal/controller/datascience/modelgroup"
	modelgroupartifact "github.com/oracle/provider-oci/internal/controller/datascience/modelgroupartifact"
	modelgroupversionhistory "github.com/oracle/provider-oci/internal/controller/datascience/modelgroupversionhistory"
	modelprovenance "github.com/oracle/provider-oci/internal/controller/datascience/modelprovenance"
	modelversionset "github.com/oracle/provider-oci/internal/controller/datascience/modelversionset"
	notebooksession "github.com/oracle/provider-oci/internal/controller/datascience/notebooksession"
	pipeline "github.com/oracle/provider-oci/internal/controller/datascience/pipeline"
	pipelinerun "github.com/oracle/provider-oci/internal/controller/datascience/pipelinerun"
	privateendpoint "github.com/oracle/provider-oci/internal/controller/datascience/privateendpoint"
	project "github.com/oracle/provider-oci/internal/controller/datascience/project"
	schedule "github.com/oracle/provider-oci/internal/controller/datascience/schedule"
)

// Setup_datascience creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_datascience(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
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

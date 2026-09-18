/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	computetarget "github.com/oracle/provider-oci/internal/controller/cluster/datascience/computetarget"
	job "github.com/oracle/provider-oci/internal/controller/cluster/datascience/job"
	jobrun "github.com/oracle/provider-oci/internal/controller/cluster/datascience/jobrun"
	mlapplication "github.com/oracle/provider-oci/internal/controller/cluster/datascience/mlapplication"
	mlapplicationimplementation "github.com/oracle/provider-oci/internal/controller/cluster/datascience/mlapplicationimplementation"
	mlapplicationinstance "github.com/oracle/provider-oci/internal/controller/cluster/datascience/mlapplicationinstance"
	model "github.com/oracle/provider-oci/internal/controller/cluster/datascience/model"
	modelartifactexport "github.com/oracle/provider-oci/internal/controller/cluster/datascience/modelartifactexport"
	modelartifactimport "github.com/oracle/provider-oci/internal/controller/cluster/datascience/modelartifactimport"
	modelcustommetadataartifact "github.com/oracle/provider-oci/internal/controller/cluster/datascience/modelcustommetadataartifact"
	modeldefinedmetadataartifact "github.com/oracle/provider-oci/internal/controller/cluster/datascience/modeldefinedmetadataartifact"
	modeldeployment "github.com/oracle/provider-oci/internal/controller/cluster/datascience/modeldeployment"
	modelgroup "github.com/oracle/provider-oci/internal/controller/cluster/datascience/modelgroup"
	modelgroupartifact "github.com/oracle/provider-oci/internal/controller/cluster/datascience/modelgroupartifact"
	modelgroupversionhistory "github.com/oracle/provider-oci/internal/controller/cluster/datascience/modelgroupversionhistory"
	modelprovenance "github.com/oracle/provider-oci/internal/controller/cluster/datascience/modelprovenance"
	modelversionset "github.com/oracle/provider-oci/internal/controller/cluster/datascience/modelversionset"
	notebooksession "github.com/oracle/provider-oci/internal/controller/cluster/datascience/notebooksession"
	pipeline "github.com/oracle/provider-oci/internal/controller/cluster/datascience/pipeline"
	pipelinerun "github.com/oracle/provider-oci/internal/controller/cluster/datascience/pipelinerun"
	privateendpoint "github.com/oracle/provider-oci/internal/controller/cluster/datascience/privateendpoint"
	project "github.com/oracle/provider-oci/internal/controller/cluster/datascience/project"
	schedule "github.com/oracle/provider-oci/internal/controller/cluster/datascience/schedule"
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

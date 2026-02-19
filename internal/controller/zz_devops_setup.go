/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	buildpipeline "github.com/oracle/provider-oci/internal/controller/devops/buildpipeline"
	buildpipelinestage "github.com/oracle/provider-oci/internal/controller/devops/buildpipelinestage"
	buildrun "github.com/oracle/provider-oci/internal/controller/devops/buildrun"
	connection "github.com/oracle/provider-oci/internal/controller/devops/connection"
	deployartifact "github.com/oracle/provider-oci/internal/controller/devops/deployartifact"
	deployenvironment "github.com/oracle/provider-oci/internal/controller/devops/deployenvironment"
	deployment "github.com/oracle/provider-oci/internal/controller/devops/deployment"
	deploypipeline "github.com/oracle/provider-oci/internal/controller/devops/deploypipeline"
	deploystage "github.com/oracle/provider-oci/internal/controller/devops/deploystage"
	project "github.com/oracle/provider-oci/internal/controller/devops/project"
	projectrepositorysetting "github.com/oracle/provider-oci/internal/controller/devops/projectrepositorysetting"
	repository "github.com/oracle/provider-oci/internal/controller/devops/repository"
	repositorymirror "github.com/oracle/provider-oci/internal/controller/devops/repositorymirror"
	repositoryprotectedbranchmanagement "github.com/oracle/provider-oci/internal/controller/devops/repositoryprotectedbranchmanagement"
	repositoryref "github.com/oracle/provider-oci/internal/controller/devops/repositoryref"
	repositorysetting "github.com/oracle/provider-oci/internal/controller/devops/repositorysetting"
	trigger "github.com/oracle/provider-oci/internal/controller/devops/trigger"
)

// Setup_devops creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_devops(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		buildpipeline.Setup,
		buildpipelinestage.Setup,
		buildrun.Setup,
		connection.Setup,
		deployartifact.Setup,
		deployenvironment.Setup,
		deployment.Setup,
		deploypipeline.Setup,
		deploystage.Setup,
		project.Setup,
		projectrepositorysetting.Setup,
		repository.Setup,
		repositorymirror.Setup,
		repositoryprotectedbranchmanagement.Setup,
		repositoryref.Setup,
		repositorysetting.Setup,
		trigger.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

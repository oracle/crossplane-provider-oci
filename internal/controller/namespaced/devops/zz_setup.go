/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	buildpipeline "github.com/oracle/provider-oci/internal/controller/namespaced/devops/buildpipeline"
	buildpipelinestage "github.com/oracle/provider-oci/internal/controller/namespaced/devops/buildpipelinestage"
	buildrun "github.com/oracle/provider-oci/internal/controller/namespaced/devops/buildrun"
	connection "github.com/oracle/provider-oci/internal/controller/namespaced/devops/connection"
	deployartifact "github.com/oracle/provider-oci/internal/controller/namespaced/devops/deployartifact"
	deployenvironment "github.com/oracle/provider-oci/internal/controller/namespaced/devops/deployenvironment"
	deployment "github.com/oracle/provider-oci/internal/controller/namespaced/devops/deployment"
	deploypipeline "github.com/oracle/provider-oci/internal/controller/namespaced/devops/deploypipeline"
	deploystage "github.com/oracle/provider-oci/internal/controller/namespaced/devops/deploystage"
	project "github.com/oracle/provider-oci/internal/controller/namespaced/devops/project"
	projectrepositorysetting "github.com/oracle/provider-oci/internal/controller/namespaced/devops/projectrepositorysetting"
	repository "github.com/oracle/provider-oci/internal/controller/namespaced/devops/repository"
	repositorymirror "github.com/oracle/provider-oci/internal/controller/namespaced/devops/repositorymirror"
	repositoryprotectedbranchmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/devops/repositoryprotectedbranchmanagement"
	repositoryref "github.com/oracle/provider-oci/internal/controller/namespaced/devops/repositoryref"
	repositorysetting "github.com/oracle/provider-oci/internal/controller/namespaced/devops/repositorysetting"
	trigger "github.com/oracle/provider-oci/internal/controller/namespaced/devops/trigger"
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

// SetupGated_devops creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_devops(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		buildpipeline.SetupGated,
		buildpipelinestage.SetupGated,
		buildrun.SetupGated,
		connection.SetupGated,
		deployartifact.SetupGated,
		deployenvironment.SetupGated,
		deployment.SetupGated,
		deploypipeline.SetupGated,
		deploystage.SetupGated,
		project.SetupGated,
		projectrepositorysetting.SetupGated,
		repository.SetupGated,
		repositorymirror.SetupGated,
		repositoryprotectedbranchmanagement.SetupGated,
		repositoryref.SetupGated,
		repositorysetting.SetupGated,
		trigger.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

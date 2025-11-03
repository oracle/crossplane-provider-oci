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
	deployartifact "github.com/oracle/provider-oci/internal/controller/devops/deployartifact"
	deployenvironment "github.com/oracle/provider-oci/internal/controller/devops/deployenvironment"
	deploypipeline "github.com/oracle/provider-oci/internal/controller/devops/deploypipeline"
	deploystage "github.com/oracle/provider-oci/internal/controller/devops/deploystage"
	devopsconnection "github.com/oracle/provider-oci/internal/controller/devops/devopsconnection"
	devopsdeployment "github.com/oracle/provider-oci/internal/controller/devops/devopsdeployment"
	devopsproject "github.com/oracle/provider-oci/internal/controller/devops/devopsproject"
	devopsrepository "github.com/oracle/provider-oci/internal/controller/devops/devopsrepository"
	projectrepositorysetting "github.com/oracle/provider-oci/internal/controller/devops/projectrepositorysetting"
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
		deployartifact.Setup,
		deployenvironment.Setup,
		deploypipeline.Setup,
		deploystage.Setup,
		devopsconnection.Setup,
		devopsdeployment.Setup,
		devopsproject.Setup,
		devopsrepository.Setup,
		projectrepositorysetting.Setup,
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

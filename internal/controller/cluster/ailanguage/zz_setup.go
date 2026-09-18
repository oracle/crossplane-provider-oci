/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	endpoint "github.com/oracle/provider-oci/internal/controller/cluster/ailanguage/endpoint"
	job "github.com/oracle/provider-oci/internal/controller/cluster/ailanguage/job"
	model "github.com/oracle/provider-oci/internal/controller/cluster/ailanguage/model"
	project "github.com/oracle/provider-oci/internal/controller/cluster/ailanguage/project"
)

// Setup_ailanguage creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_ailanguage(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		endpoint.Setup,
		job.Setup,
		model.Setup,
		project.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_ailanguage creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_ailanguage(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		endpoint.SetupGated,
		job.SetupGated,
		model.SetupGated,
		project.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

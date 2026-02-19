/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	endpoint "github.com/oracle/provider-oci/internal/controller/ailanguage/endpoint"
	job "github.com/oracle/provider-oci/internal/controller/ailanguage/job"
	model "github.com/oracle/provider-oci/internal/controller/ailanguage/model"
	project "github.com/oracle/provider-oci/internal/controller/ailanguage/project"
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

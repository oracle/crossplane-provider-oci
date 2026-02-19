/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	model "github.com/oracle/provider-oci/internal/controller/aidocument/model"
	processorjob "github.com/oracle/provider-oci/internal/controller/aidocument/processorjob"
	project "github.com/oracle/provider-oci/internal/controller/aidocument/project"
)

// Setup_aidocument creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_aidocument(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		model.Setup,
		processorjob.Setup,
		project.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	httpmonitor "github.com/oracle/provider-oci/internal/controller/healthchecks/httpmonitor"
	httpprobe "github.com/oracle/provider-oci/internal/controller/healthchecks/httpprobe"
	pingmonitor "github.com/oracle/provider-oci/internal/controller/healthchecks/pingmonitor"
	pingprobe "github.com/oracle/provider-oci/internal/controller/healthchecks/pingprobe"
)

// Setup_healthchecks creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_healthchecks(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		httpmonitor.Setup,
		httpprobe.Setup,
		pingmonitor.Setup,
		pingprobe.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

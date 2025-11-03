/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	checkshttpprobe "github.com/oracle/provider-oci/internal/controller/healthchecks/checkshttpprobe"
	checkspingprobe "github.com/oracle/provider-oci/internal/controller/healthchecks/checkspingprobe"
	httpmonitor "github.com/oracle/provider-oci/internal/controller/healthchecks/httpmonitor"
	pingmonitor "github.com/oracle/provider-oci/internal/controller/healthchecks/pingmonitor"
)

// Setup_healthchecks creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_healthchecks(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		checkshttpprobe.Setup,
		checkspingprobe.Setup,
		httpmonitor.Setup,
		pingmonitor.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

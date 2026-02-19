/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	javadownloadreport "github.com/oracle/provider-oci/internal/controller/jmsjavadownloads/javadownloadreport"
	javadownloadtoken "github.com/oracle/provider-oci/internal/controller/jmsjavadownloads/javadownloadtoken"
	javalicenseacceptancerecord "github.com/oracle/provider-oci/internal/controller/jmsjavadownloads/javalicenseacceptancerecord"
)

// Setup_jmsjavadownloads creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_jmsjavadownloads(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		javadownloadreport.Setup,
		javadownloadtoken.Setup,
		javalicenseacceptancerecord.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

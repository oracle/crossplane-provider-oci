/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	javadownloadreport "github.com/oracle/provider-oci/internal/controller/namespaced/jmsjavadownloads/javadownloadreport"
	javadownloadtoken "github.com/oracle/provider-oci/internal/controller/namespaced/jmsjavadownloads/javadownloadtoken"
	javalicenseacceptancerecord "github.com/oracle/provider-oci/internal/controller/namespaced/jmsjavadownloads/javalicenseacceptancerecord"
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

// SetupGated_jmsjavadownloads creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_jmsjavadownloads(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		javadownloadreport.SetupGated,
		javadownloadtoken.SetupGated,
		javalicenseacceptancerecord.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

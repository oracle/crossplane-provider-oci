/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	digitaltwinadapter "github.com/oracle/provider-oci/internal/controller/namespaced/iot/digitaltwinadapter"
	digitaltwininstance "github.com/oracle/provider-oci/internal/controller/namespaced/iot/digitaltwininstance"
	digitaltwininstanceinvokerawcommand "github.com/oracle/provider-oci/internal/controller/namespaced/iot/digitaltwininstanceinvokerawcommand"
	digitaltwinmodel "github.com/oracle/provider-oci/internal/controller/namespaced/iot/digitaltwinmodel"
	digitaltwinrelationship "github.com/oracle/provider-oci/internal/controller/namespaced/iot/digitaltwinrelationship"
	iotdomain "github.com/oracle/provider-oci/internal/controller/namespaced/iot/iotdomain"
	iotdomainchangedataretentionperiod "github.com/oracle/provider-oci/internal/controller/namespaced/iot/iotdomainchangedataretentionperiod"
	iotdomainconfiguredataaccess "github.com/oracle/provider-oci/internal/controller/namespaced/iot/iotdomainconfiguredataaccess"
	iotdomaingroup "github.com/oracle/provider-oci/internal/controller/namespaced/iot/iotdomaingroup"
	iotdomaingroupconfiguredataaccess "github.com/oracle/provider-oci/internal/controller/namespaced/iot/iotdomaingroupconfiguredataaccess"
	iotflowruntime "github.com/oracle/provider-oci/internal/controller/namespaced/iot/iotflowruntime"
	iotflowruntimeactivate "github.com/oracle/provider-oci/internal/controller/namespaced/iot/iotflowruntimeactivate"
	iotflowruntimedeactivate "github.com/oracle/provider-oci/internal/controller/namespaced/iot/iotflowruntimedeactivate"
	iotflowruntimeflow "github.com/oracle/provider-oci/internal/controller/namespaced/iot/iotflowruntimeflow"
)

// Setup_iot creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_iot(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		digitaltwinadapter.Setup,
		digitaltwininstance.Setup,
		digitaltwininstanceinvokerawcommand.Setup,
		digitaltwinmodel.Setup,
		digitaltwinrelationship.Setup,
		iotdomain.Setup,
		iotdomainchangedataretentionperiod.Setup,
		iotdomainconfiguredataaccess.Setup,
		iotdomaingroup.Setup,
		iotdomaingroupconfiguredataaccess.Setup,
		iotflowruntime.Setup,
		iotflowruntimeactivate.Setup,
		iotflowruntimedeactivate.Setup,
		iotflowruntimeflow.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_iot creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_iot(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		digitaltwinadapter.SetupGated,
		digitaltwininstance.SetupGated,
		digitaltwininstanceinvokerawcommand.SetupGated,
		digitaltwinmodel.SetupGated,
		digitaltwinrelationship.SetupGated,
		iotdomain.SetupGated,
		iotdomainchangedataretentionperiod.SetupGated,
		iotdomainconfiguredataaccess.SetupGated,
		iotdomaingroup.SetupGated,
		iotdomaingroupconfiguredataaccess.SetupGated,
		iotflowruntime.SetupGated,
		iotflowruntimeactivate.SetupGated,
		iotflowruntimedeactivate.SetupGated,
		iotflowruntimeflow.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

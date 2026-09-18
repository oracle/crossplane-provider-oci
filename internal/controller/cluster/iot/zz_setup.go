/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	digitaltwinadapter "github.com/oracle/provider-oci/internal/controller/cluster/iot/digitaltwinadapter"
	digitaltwininstance "github.com/oracle/provider-oci/internal/controller/cluster/iot/digitaltwininstance"
	digitaltwininstanceinvokerawcommand "github.com/oracle/provider-oci/internal/controller/cluster/iot/digitaltwininstanceinvokerawcommand"
	digitaltwinmodel "github.com/oracle/provider-oci/internal/controller/cluster/iot/digitaltwinmodel"
	digitaltwinrelationship "github.com/oracle/provider-oci/internal/controller/cluster/iot/digitaltwinrelationship"
	iotdomain "github.com/oracle/provider-oci/internal/controller/cluster/iot/iotdomain"
	iotdomainchangedataretentionperiod "github.com/oracle/provider-oci/internal/controller/cluster/iot/iotdomainchangedataretentionperiod"
	iotdomainconfiguredataaccess "github.com/oracle/provider-oci/internal/controller/cluster/iot/iotdomainconfiguredataaccess"
	iotdomaingroup "github.com/oracle/provider-oci/internal/controller/cluster/iot/iotdomaingroup"
	iotdomaingroupconfiguredataaccess "github.com/oracle/provider-oci/internal/controller/cluster/iot/iotdomaingroupconfiguredataaccess"
	iotflowruntime "github.com/oracle/provider-oci/internal/controller/cluster/iot/iotflowruntime"
	iotflowruntimeactivate "github.com/oracle/provider-oci/internal/controller/cluster/iot/iotflowruntimeactivate"
	iotflowruntimedeactivate "github.com/oracle/provider-oci/internal/controller/cluster/iot/iotflowruntimedeactivate"
	iotflowruntimeflow "github.com/oracle/provider-oci/internal/controller/cluster/iot/iotflowruntimeflow"
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

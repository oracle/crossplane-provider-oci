/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	digitaltwinadapter "github.com/oracle/provider-oci/internal/controller/iot/digitaltwinadapter"
	digitaltwininstance "github.com/oracle/provider-oci/internal/controller/iot/digitaltwininstance"
	digitaltwininstanceinvokerawcommand "github.com/oracle/provider-oci/internal/controller/iot/digitaltwininstanceinvokerawcommand"
	digitaltwinmodel "github.com/oracle/provider-oci/internal/controller/iot/digitaltwinmodel"
	digitaltwinrelationship "github.com/oracle/provider-oci/internal/controller/iot/digitaltwinrelationship"
	iotdomain "github.com/oracle/provider-oci/internal/controller/iot/iotdomain"
	iotdomainchangedataretentionperiod "github.com/oracle/provider-oci/internal/controller/iot/iotdomainchangedataretentionperiod"
	iotdomainconfiguredataaccess "github.com/oracle/provider-oci/internal/controller/iot/iotdomainconfiguredataaccess"
	iotdomaingroup "github.com/oracle/provider-oci/internal/controller/iot/iotdomaingroup"
	iotdomaingroupconfiguredataaccess "github.com/oracle/provider-oci/internal/controller/iot/iotdomaingroupconfiguredataaccess"
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
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

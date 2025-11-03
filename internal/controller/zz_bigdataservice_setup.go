/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	bigdataserviceautoscalingconfiguration "github.com/oracle/provider-oci/internal/controller/bigdataservice/bigdataserviceautoscalingconfiguration"
	bigdataserviceinstance "github.com/oracle/provider-oci/internal/controller/bigdataservice/bigdataserviceinstance"
	capacityreport "github.com/oracle/provider-oci/internal/controller/bigdataservice/capacityreport"
	instanceapikey "github.com/oracle/provider-oci/internal/controller/bigdataservice/instanceapikey"
	instanceidentityconfiguration "github.com/oracle/provider-oci/internal/controller/bigdataservice/instanceidentityconfiguration"
	instancemetastoreconfig "github.com/oracle/provider-oci/internal/controller/bigdataservice/instancemetastoreconfig"
	instancenodebackup "github.com/oracle/provider-oci/internal/controller/bigdataservice/instancenodebackup"
	instancenodebackupconfiguration "github.com/oracle/provider-oci/internal/controller/bigdataservice/instancenodebackupconfiguration"
	instancenodereplaceconfiguration "github.com/oracle/provider-oci/internal/controller/bigdataservice/instancenodereplaceconfiguration"
	instanceoperationcertificatemanagement "github.com/oracle/provider-oci/internal/controller/bigdataservice/instanceoperationcertificatemanagement"
	instanceospatchaction "github.com/oracle/provider-oci/internal/controller/bigdataservice/instanceospatchaction"
	instancepatchaction "github.com/oracle/provider-oci/internal/controller/bigdataservice/instancepatchaction"
	instancereplacenodeaction "github.com/oracle/provider-oci/internal/controller/bigdataservice/instancereplacenodeaction"
	instanceresourceprincipalconfiguration "github.com/oracle/provider-oci/internal/controller/bigdataservice/instanceresourceprincipalconfiguration"
	instancesoftwareupdateaction "github.com/oracle/provider-oci/internal/controller/bigdataservice/instancesoftwareupdateaction"
)

// Setup_bigdataservice creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_bigdataservice(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bigdataserviceautoscalingconfiguration.Setup,
		bigdataserviceinstance.Setup,
		capacityreport.Setup,
		instanceapikey.Setup,
		instanceidentityconfiguration.Setup,
		instancemetastoreconfig.Setup,
		instancenodebackup.Setup,
		instancenodebackupconfiguration.Setup,
		instancenodereplaceconfiguration.Setup,
		instanceoperationcertificatemanagement.Setup,
		instanceospatchaction.Setup,
		instancepatchaction.Setup,
		instancereplacenodeaction.Setup,
		instanceresourceprincipalconfiguration.Setup,
		instancesoftwareupdateaction.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

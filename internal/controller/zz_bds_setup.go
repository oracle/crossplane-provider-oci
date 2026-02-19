/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	autoscalingconfiguration "github.com/oracle/provider-oci/internal/controller/bds/autoscalingconfiguration"
	bdscapacityreport "github.com/oracle/provider-oci/internal/controller/bds/bdscapacityreport"
	bdsinstance "github.com/oracle/provider-oci/internal/controller/bds/bdsinstance"
	bdsinstanceapikey "github.com/oracle/provider-oci/internal/controller/bds/bdsinstanceapikey"
	bdsinstanceidentityconfiguration "github.com/oracle/provider-oci/internal/controller/bds/bdsinstanceidentityconfiguration"
	bdsinstancemetastoreconfig "github.com/oracle/provider-oci/internal/controller/bds/bdsinstancemetastoreconfig"
	bdsinstancenodebackup "github.com/oracle/provider-oci/internal/controller/bds/bdsinstancenodebackup"
	bdsinstancenodebackupconfiguration "github.com/oracle/provider-oci/internal/controller/bds/bdsinstancenodebackupconfiguration"
	bdsinstancenodereplaceconfiguration "github.com/oracle/provider-oci/internal/controller/bds/bdsinstancenodereplaceconfiguration"
	bdsinstanceoperationcertificatemanagementsmanagement "github.com/oracle/provider-oci/internal/controller/bds/bdsinstanceoperationcertificatemanagementsmanagement"
	bdsinstanceospatchaction "github.com/oracle/provider-oci/internal/controller/bds/bdsinstanceospatchaction"
	bdsinstancepatchaction "github.com/oracle/provider-oci/internal/controller/bds/bdsinstancepatchaction"
	bdsinstancereplacenodeaction "github.com/oracle/provider-oci/internal/controller/bds/bdsinstancereplacenodeaction"
	bdsinstanceresourceprincipalconfiguration "github.com/oracle/provider-oci/internal/controller/bds/bdsinstanceresourceprincipalconfiguration"
	bdsinstancesoftwareupdateaction "github.com/oracle/provider-oci/internal/controller/bds/bdsinstancesoftwareupdateaction"
)

// Setup_bds creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_bds(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		autoscalingconfiguration.Setup,
		bdscapacityreport.Setup,
		bdsinstance.Setup,
		bdsinstanceapikey.Setup,
		bdsinstanceidentityconfiguration.Setup,
		bdsinstancemetastoreconfig.Setup,
		bdsinstancenodebackup.Setup,
		bdsinstancenodebackupconfiguration.Setup,
		bdsinstancenodereplaceconfiguration.Setup,
		bdsinstanceoperationcertificatemanagementsmanagement.Setup,
		bdsinstanceospatchaction.Setup,
		bdsinstancepatchaction.Setup,
		bdsinstancereplacenodeaction.Setup,
		bdsinstanceresourceprincipalconfiguration.Setup,
		bdsinstancesoftwareupdateaction.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	autoscalingconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/bds/autoscalingconfiguration"
	bdscapacityreport "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdscapacityreport"
	bdscapacityreservation "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdscapacityreservation"
	bdsclusteradminpasswordresetaction "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsclusteradminpasswordresetaction"
	bdsinstance "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstance"
	bdsinstanceapikey "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstanceapikey"
	bdsinstancebdscapacityreservationconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstancebdscapacityreservationconfiguration"
	bdsinstancebdscertificateconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstancebdscertificateconfiguration"
	bdsinstanceexecutebootstrapscriptaction "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstanceexecutebootstrapscriptaction"
	bdsinstanceidentityconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstanceidentityconfiguration"
	bdsinstancemetastoreconfig "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstancemetastoreconfig"
	bdsinstancenodebackup "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstancenodebackup"
	bdsinstancenodebackupconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstancenodebackupconfiguration"
	bdsinstancenodereplaceconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstancenodereplaceconfiguration"
	bdsinstanceoperationcertificatemanagementsmanagement "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstanceoperationcertificatemanagementsmanagement"
	bdsinstanceospatchaction "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstanceospatchaction"
	bdsinstancepatchaction "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstancepatchaction"
	bdsinstancereplacenodeaction "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstancereplacenodeaction"
	bdsinstanceresourceprincipalconfiguration "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstanceresourceprincipalconfiguration"
	bdsinstancesoftwareupdateaction "github.com/oracle/provider-oci/internal/controller/cluster/bds/bdsinstancesoftwareupdateaction"
)

// Setup_bds creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_bds(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		autoscalingconfiguration.Setup,
		bdscapacityreport.Setup,
		bdscapacityreservation.Setup,
		bdsclusteradminpasswordresetaction.Setup,
		bdsinstance.Setup,
		bdsinstanceapikey.Setup,
		bdsinstancebdscapacityreservationconfiguration.Setup,
		bdsinstancebdscertificateconfiguration.Setup,
		bdsinstanceexecutebootstrapscriptaction.Setup,
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

// SetupGated_bds creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_bds(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		autoscalingconfiguration.SetupGated,
		bdscapacityreport.SetupGated,
		bdscapacityreservation.SetupGated,
		bdsclusteradminpasswordresetaction.SetupGated,
		bdsinstance.SetupGated,
		bdsinstanceapikey.SetupGated,
		bdsinstancebdscapacityreservationconfiguration.SetupGated,
		bdsinstancebdscertificateconfiguration.SetupGated,
		bdsinstanceexecutebootstrapscriptaction.SetupGated,
		bdsinstanceidentityconfiguration.SetupGated,
		bdsinstancemetastoreconfig.SetupGated,
		bdsinstancenodebackup.SetupGated,
		bdsinstancenodebackupconfiguration.SetupGated,
		bdsinstancenodereplaceconfiguration.SetupGated,
		bdsinstanceoperationcertificatemanagementsmanagement.SetupGated,
		bdsinstanceospatchaction.SetupGated,
		bdsinstancepatchaction.SetupGated,
		bdsinstancereplacenodeaction.SetupGated,
		bdsinstanceresourceprincipalconfiguration.SetupGated,
		bdsinstancesoftwareupdateaction.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

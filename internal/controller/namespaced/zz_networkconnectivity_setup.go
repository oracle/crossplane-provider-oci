/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	cpe "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/cpe"
	crossconnect "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/crossconnect"
	crossconnectgroup "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/crossconnectgroup"
	defaultdrgroutetable "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/defaultdrgroutetable"
	drg "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/drg"
	drgattachment "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/drgattachment"
	drgattachmentmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/drgattachmentmanagement"
	drgattachmentslist "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/drgattachmentslist"
	drgroutedistribution "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/drgroutedistribution"
	drgroutedistributionstatement "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/drgroutedistributionstatement"
	drgroutetable "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/drgroutetable"
	drgroutetablerouterule "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/drgroutetablerouterule"
	ipsec "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/ipsec"
	ipsecconnectiontunnelmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/ipsecconnectiontunnelmanagement"
	virtualcircuit "github.com/oracle/provider-oci/internal/controller/namespaced/networkconnectivity/virtualcircuit"
)

// Setup_networkconnectivity creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_networkconnectivity(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cpe.Setup,
		crossconnect.Setup,
		crossconnectgroup.Setup,
		defaultdrgroutetable.Setup,
		drg.Setup,
		drgattachment.Setup,
		drgattachmentmanagement.Setup,
		drgattachmentslist.Setup,
		drgroutedistribution.Setup,
		drgroutedistributionstatement.Setup,
		drgroutetable.Setup,
		drgroutetablerouterule.Setup,
		ipsec.Setup,
		ipsecconnectiontunnelmanagement.Setup,
		virtualcircuit.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_networkconnectivity creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_networkconnectivity(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cpe.SetupGated,
		crossconnect.SetupGated,
		crossconnectgroup.SetupGated,
		defaultdrgroutetable.SetupGated,
		drg.SetupGated,
		drgattachment.SetupGated,
		drgattachmentmanagement.SetupGated,
		drgattachmentslist.SetupGated,
		drgroutedistribution.SetupGated,
		drgroutedistributionstatement.SetupGated,
		drgroutetable.SetupGated,
		drgroutetablerouterule.SetupGated,
		ipsec.SetupGated,
		ipsecconnectiontunnelmanagement.SetupGated,
		virtualcircuit.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

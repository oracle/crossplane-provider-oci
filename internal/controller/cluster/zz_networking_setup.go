/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	defaultdhcpoptions "github.com/oracle/provider-oci/internal/controller/cluster/networking/defaultdhcpoptions"
	defaultroutetable "github.com/oracle/provider-oci/internal/controller/cluster/networking/defaultroutetable"
	defaultsecuritylist "github.com/oracle/provider-oci/internal/controller/cluster/networking/defaultsecuritylist"
	dhcpoptions "github.com/oracle/provider-oci/internal/controller/cluster/networking/dhcpoptions"
	internetgateway "github.com/oracle/provider-oci/internal/controller/cluster/networking/internetgateway"
	ipv6 "github.com/oracle/provider-oci/internal/controller/cluster/networking/ipv6"
	localpeeringgateway "github.com/oracle/provider-oci/internal/controller/cluster/networking/localpeeringgateway"
	natgateway "github.com/oracle/provider-oci/internal/controller/cluster/networking/natgateway"
	networksecuritygroup "github.com/oracle/provider-oci/internal/controller/cluster/networking/networksecuritygroup"
	networksecuritygroupsecurityrule "github.com/oracle/provider-oci/internal/controller/cluster/networking/networksecuritygroupsecurityrule"
	privateip "github.com/oracle/provider-oci/internal/controller/cluster/networking/privateip"
	publicip "github.com/oracle/provider-oci/internal/controller/cluster/networking/publicip"
	publicippool "github.com/oracle/provider-oci/internal/controller/cluster/networking/publicippool"
	publicippoolcapacity "github.com/oracle/provider-oci/internal/controller/cluster/networking/publicippoolcapacity"
	remotepeeringconnection "github.com/oracle/provider-oci/internal/controller/cluster/networking/remotepeeringconnection"
	routetable "github.com/oracle/provider-oci/internal/controller/cluster/networking/routetable"
	routetableattachment "github.com/oracle/provider-oci/internal/controller/cluster/networking/routetableattachment"
	securitylist "github.com/oracle/provider-oci/internal/controller/cluster/networking/securitylist"
	servicegateway "github.com/oracle/provider-oci/internal/controller/cluster/networking/servicegateway"
	subnet "github.com/oracle/provider-oci/internal/controller/cluster/networking/subnet"
	vcn "github.com/oracle/provider-oci/internal/controller/cluster/networking/vcn"
	vlan "github.com/oracle/provider-oci/internal/controller/cluster/networking/vlan"
	vnicattachment "github.com/oracle/provider-oci/internal/controller/cluster/networking/vnicattachment"
)

// Setup_networking creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_networking(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		defaultdhcpoptions.Setup,
		defaultroutetable.Setup,
		defaultsecuritylist.Setup,
		dhcpoptions.Setup,
		internetgateway.Setup,
		ipv6.Setup,
		localpeeringgateway.Setup,
		natgateway.Setup,
		networksecuritygroup.Setup,
		networksecuritygroupsecurityrule.Setup,
		privateip.Setup,
		publicip.Setup,
		publicippool.Setup,
		publicippoolcapacity.Setup,
		remotepeeringconnection.Setup,
		routetable.Setup,
		routetableattachment.Setup,
		securitylist.Setup,
		servicegateway.Setup,
		subnet.Setup,
		vcn.Setup,
		vlan.Setup,
		vnicattachment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_networking creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_networking(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		defaultdhcpoptions.SetupGated,
		defaultroutetable.SetupGated,
		defaultsecuritylist.SetupGated,
		dhcpoptions.SetupGated,
		internetgateway.SetupGated,
		ipv6.SetupGated,
		localpeeringgateway.SetupGated,
		natgateway.SetupGated,
		networksecuritygroup.SetupGated,
		networksecuritygroupsecurityrule.SetupGated,
		privateip.SetupGated,
		publicip.SetupGated,
		publicippool.SetupGated,
		publicippoolcapacity.SetupGated,
		remotepeeringconnection.SetupGated,
		routetable.SetupGated,
		routetableattachment.SetupGated,
		securitylist.SetupGated,
		servicegateway.SetupGated,
		subnet.SetupGated,
		vcn.SetupGated,
		vlan.SetupGated,
		vnicattachment.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

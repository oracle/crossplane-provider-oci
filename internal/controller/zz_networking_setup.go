/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	defaultdhcpoptions "github.com/oracle/provider-oci/internal/controller/networking/defaultdhcpoptions"
	defaultroutetable "github.com/oracle/provider-oci/internal/controller/networking/defaultroutetable"
	defaultsecuritylist "github.com/oracle/provider-oci/internal/controller/networking/defaultsecuritylist"
	dhcpoptions "github.com/oracle/provider-oci/internal/controller/networking/dhcpoptions"
	internetgateway "github.com/oracle/provider-oci/internal/controller/networking/internetgateway"
	ipv6 "github.com/oracle/provider-oci/internal/controller/networking/ipv6"
	localpeeringgateway "github.com/oracle/provider-oci/internal/controller/networking/localpeeringgateway"
	natgateway "github.com/oracle/provider-oci/internal/controller/networking/natgateway"
	networksecuritygroup "github.com/oracle/provider-oci/internal/controller/networking/networksecuritygroup"
	networksecuritygroupsecurityrule "github.com/oracle/provider-oci/internal/controller/networking/networksecuritygroupsecurityrule"
	privateip "github.com/oracle/provider-oci/internal/controller/networking/privateip"
	publicip "github.com/oracle/provider-oci/internal/controller/networking/publicip"
	publicippool "github.com/oracle/provider-oci/internal/controller/networking/publicippool"
	publicippoolcapacity "github.com/oracle/provider-oci/internal/controller/networking/publicippoolcapacity"
	remotepeeringconnection "github.com/oracle/provider-oci/internal/controller/networking/remotepeeringconnection"
	routetable "github.com/oracle/provider-oci/internal/controller/networking/routetable"
	routetableattachment "github.com/oracle/provider-oci/internal/controller/networking/routetableattachment"
	securitylist "github.com/oracle/provider-oci/internal/controller/networking/securitylist"
	servicegateway "github.com/oracle/provider-oci/internal/controller/networking/servicegateway"
	subnet "github.com/oracle/provider-oci/internal/controller/networking/subnet"
	vcn "github.com/oracle/provider-oci/internal/controller/networking/vcn"
	vlan "github.com/oracle/provider-oci/internal/controller/networking/vlan"
	vnicattachment "github.com/oracle/provider-oci/internal/controller/networking/vnicattachment"
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

/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	apikey "github.com/oracle/provider-oci/internal/controller/namespaced/identity/apikey"
	authenticationpolicy "github.com/oracle/provider-oci/internal/controller/namespaced/identity/authenticationpolicy"
	authtoken "github.com/oracle/provider-oci/internal/controller/namespaced/identity/authtoken"
	compartment "github.com/oracle/provider-oci/internal/controller/namespaced/identity/compartment"
	customersecretkey "github.com/oracle/provider-oci/internal/controller/namespaced/identity/customersecretkey"
	dbcredential "github.com/oracle/provider-oci/internal/controller/namespaced/identity/dbcredential"
	domain "github.com/oracle/provider-oci/internal/controller/namespaced/identity/domain"
	domainreplicationtoregion "github.com/oracle/provider-oci/internal/controller/namespaced/identity/domainreplicationtoregion"
	dynamicgroup "github.com/oracle/provider-oci/internal/controller/namespaced/identity/dynamicgroup"
	group "github.com/oracle/provider-oci/internal/controller/namespaced/identity/group"
	identityprovider "github.com/oracle/provider-oci/internal/controller/namespaced/identity/identityprovider"
	idpgroupmapping "github.com/oracle/provider-oci/internal/controller/namespaced/identity/idpgroupmapping"
	importstandardtagsmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/identity/importstandardtagsmanagement"
	networksource "github.com/oracle/provider-oci/internal/controller/namespaced/identity/networksource"
	policy "github.com/oracle/provider-oci/internal/controller/namespaced/identity/policy"
	smtpcredential "github.com/oracle/provider-oci/internal/controller/namespaced/identity/smtpcredential"
	tag "github.com/oracle/provider-oci/internal/controller/namespaced/identity/tag"
	tagdefault "github.com/oracle/provider-oci/internal/controller/namespaced/identity/tagdefault"
	tagnamespace "github.com/oracle/provider-oci/internal/controller/namespaced/identity/tagnamespace"
	uipassword "github.com/oracle/provider-oci/internal/controller/namespaced/identity/uipassword"
	user "github.com/oracle/provider-oci/internal/controller/namespaced/identity/user"
	usercapabilitiesmanagement "github.com/oracle/provider-oci/internal/controller/namespaced/identity/usercapabilitiesmanagement"
	usergroupmembership "github.com/oracle/provider-oci/internal/controller/namespaced/identity/usergroupmembership"
)

// Setup_identity creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_identity(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		apikey.Setup,
		authenticationpolicy.Setup,
		authtoken.Setup,
		compartment.Setup,
		customersecretkey.Setup,
		dbcredential.Setup,
		domain.Setup,
		domainreplicationtoregion.Setup,
		dynamicgroup.Setup,
		group.Setup,
		identityprovider.Setup,
		idpgroupmapping.Setup,
		importstandardtagsmanagement.Setup,
		networksource.Setup,
		policy.Setup,
		smtpcredential.Setup,
		tag.Setup,
		tagdefault.Setup,
		tagnamespace.Setup,
		uipassword.Setup,
		user.Setup,
		usercapabilitiesmanagement.Setup,
		usergroupmembership.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_identity creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_identity(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		apikey.SetupGated,
		authenticationpolicy.SetupGated,
		authtoken.SetupGated,
		compartment.SetupGated,
		customersecretkey.SetupGated,
		dbcredential.SetupGated,
		domain.SetupGated,
		domainreplicationtoregion.SetupGated,
		dynamicgroup.SetupGated,
		group.SetupGated,
		identityprovider.SetupGated,
		idpgroupmapping.SetupGated,
		importstandardtagsmanagement.SetupGated,
		networksource.SetupGated,
		policy.SetupGated,
		smtpcredential.SetupGated,
		tag.SetupGated,
		tagdefault.SetupGated,
		tagnamespace.SetupGated,
		uipassword.SetupGated,
		user.SetupGated,
		usercapabilitiesmanagement.SetupGated,
		usergroupmembership.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	apikey "github.com/oracle/provider-oci/internal/controller/identity/apikey"
	authenticationpolicy "github.com/oracle/provider-oci/internal/controller/identity/authenticationpolicy"
	authtoken "github.com/oracle/provider-oci/internal/controller/identity/authtoken"
	compartment "github.com/oracle/provider-oci/internal/controller/identity/compartment"
	customersecretkey "github.com/oracle/provider-oci/internal/controller/identity/customersecretkey"
	dbcredential "github.com/oracle/provider-oci/internal/controller/identity/dbcredential"
	domain "github.com/oracle/provider-oci/internal/controller/identity/domain"
	domainreplicationtoregion "github.com/oracle/provider-oci/internal/controller/identity/domainreplicationtoregion"
	dynamicgroup "github.com/oracle/provider-oci/internal/controller/identity/dynamicgroup"
	group "github.com/oracle/provider-oci/internal/controller/identity/group"
	identityprovider "github.com/oracle/provider-oci/internal/controller/identity/identityprovider"
	idpgroupmapping "github.com/oracle/provider-oci/internal/controller/identity/idpgroupmapping"
	importstandardtagsmanagement "github.com/oracle/provider-oci/internal/controller/identity/importstandardtagsmanagement"
	networksource "github.com/oracle/provider-oci/internal/controller/identity/networksource"
	policy "github.com/oracle/provider-oci/internal/controller/identity/policy"
	smtpcredential "github.com/oracle/provider-oci/internal/controller/identity/smtpcredential"
	tag "github.com/oracle/provider-oci/internal/controller/identity/tag"
	tagdefault "github.com/oracle/provider-oci/internal/controller/identity/tagdefault"
	tagnamespace "github.com/oracle/provider-oci/internal/controller/identity/tagnamespace"
	uipassword "github.com/oracle/provider-oci/internal/controller/identity/uipassword"
	user "github.com/oracle/provider-oci/internal/controller/identity/user"
	usercapabilitiesmanagement "github.com/oracle/provider-oci/internal/controller/identity/usercapabilitiesmanagement"
	usergroupmembership "github.com/oracle/provider-oci/internal/controller/identity/usergroupmembership"
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

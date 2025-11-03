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
	domainreplicationtoregion "github.com/oracle/provider-oci/internal/controller/identity/domainreplicationtoregion"
	domainsaccountrecoverysetting "github.com/oracle/provider-oci/internal/controller/identity/domainsaccountrecoverysetting"
	domainsapikey "github.com/oracle/provider-oci/internal/controller/identity/domainsapikey"
	domainsapp "github.com/oracle/provider-oci/internal/controller/identity/domainsapp"
	domainsapprole "github.com/oracle/provider-oci/internal/controller/identity/domainsapprole"
	domainsapprovalworkflow "github.com/oracle/provider-oci/internal/controller/identity/domainsapprovalworkflow"
	domainsapprovalworkflowassignment "github.com/oracle/provider-oci/internal/controller/identity/domainsapprovalworkflowassignment"
	domainsapprovalworkflowstep "github.com/oracle/provider-oci/internal/controller/identity/domainsapprovalworkflowstep"
	domainsauthenticationfactorsetting "github.com/oracle/provider-oci/internal/controller/identity/domainsauthenticationfactorsetting"
	domainsauthtoken "github.com/oracle/provider-oci/internal/controller/identity/domainsauthtoken"
	domainscloudgate "github.com/oracle/provider-oci/internal/controller/identity/domainscloudgate"
	domainscloudgatemapping "github.com/oracle/provider-oci/internal/controller/identity/domainscloudgatemapping"
	domainscloudgateserver "github.com/oracle/provider-oci/internal/controller/identity/domainscloudgateserver"
	domainscondition "github.com/oracle/provider-oci/internal/controller/identity/domainscondition"
	domainscustomersecretkey "github.com/oracle/provider-oci/internal/controller/identity/domainscustomersecretkey"
	domainsdynamicresourcegroup "github.com/oracle/provider-oci/internal/controller/identity/domainsdynamicresourcegroup"
	domainsgrant "github.com/oracle/provider-oci/internal/controller/identity/domainsgrant"
	domainsgroup "github.com/oracle/provider-oci/internal/controller/identity/domainsgroup"
	domainsidentitypropagationtrust "github.com/oracle/provider-oci/internal/controller/identity/domainsidentitypropagationtrust"
	domainsidentityprovider "github.com/oracle/provider-oci/internal/controller/identity/domainsidentityprovider"
	domainsidentitysetting "github.com/oracle/provider-oci/internal/controller/identity/domainsidentitysetting"
	domainskmsisetting "github.com/oracle/provider-oci/internal/controller/identity/domainskmsisetting"
	domainsmyapikey "github.com/oracle/provider-oci/internal/controller/identity/domainsmyapikey"
	domainsmyauthtoken "github.com/oracle/provider-oci/internal/controller/identity/domainsmyauthtoken"
	domainsmycustomersecretkey "github.com/oracle/provider-oci/internal/controller/identity/domainsmycustomersecretkey"
	domainsmyoauth2clientcredential "github.com/oracle/provider-oci/internal/controller/identity/domainsmyoauth2clientcredential"
	domainsmyrequest "github.com/oracle/provider-oci/internal/controller/identity/domainsmyrequest"
	domainsmysmtpcredential "github.com/oracle/provider-oci/internal/controller/identity/domainsmysmtpcredential"
	domainsmysupportaccount "github.com/oracle/provider-oci/internal/controller/identity/domainsmysupportaccount"
	domainsmyuserdbcredential "github.com/oracle/provider-oci/internal/controller/identity/domainsmyuserdbcredential"
	domainsnetworkperimeter "github.com/oracle/provider-oci/internal/controller/identity/domainsnetworkperimeter"
	domainsnotificationsetting "github.com/oracle/provider-oci/internal/controller/identity/domainsnotificationsetting"
	domainsoauth2clientcredential "github.com/oracle/provider-oci/internal/controller/identity/domainsoauth2clientcredential"
	domainsoauthclientcertificate "github.com/oracle/provider-oci/internal/controller/identity/domainsoauthclientcertificate"
	domainsoauthpartnercertificate "github.com/oracle/provider-oci/internal/controller/identity/domainsoauthpartnercertificate"
	domainspasswordpolicy "github.com/oracle/provider-oci/internal/controller/identity/domainspasswordpolicy"
	domainspolicy "github.com/oracle/provider-oci/internal/controller/identity/domainspolicy"
	domainsrule "github.com/oracle/provider-oci/internal/controller/identity/domainsrule"
	domainssecurityquestion "github.com/oracle/provider-oci/internal/controller/identity/domainssecurityquestion"
	domainssecurityquestionsetting "github.com/oracle/provider-oci/internal/controller/identity/domainssecurityquestionsetting"
	domainsselfregistrationprofile "github.com/oracle/provider-oci/internal/controller/identity/domainsselfregistrationprofile"
	domainssetting "github.com/oracle/provider-oci/internal/controller/identity/domainssetting"
	domainssmtpcredential "github.com/oracle/provider-oci/internal/controller/identity/domainssmtpcredential"
	domainssocialidentityprovider "github.com/oracle/provider-oci/internal/controller/identity/domainssocialidentityprovider"
	domainsuser "github.com/oracle/provider-oci/internal/controller/identity/domainsuser"
	domainsuserdbcredential "github.com/oracle/provider-oci/internal/controller/identity/domainsuserdbcredential"
	dynamicgroup "github.com/oracle/provider-oci/internal/controller/identity/dynamicgroup"
	group "github.com/oracle/provider-oci/internal/controller/identity/group"
	identitydomain "github.com/oracle/provider-oci/internal/controller/identity/identitydomain"
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
		domainreplicationtoregion.Setup,
		domainsaccountrecoverysetting.Setup,
		domainsapikey.Setup,
		domainsapp.Setup,
		domainsapprole.Setup,
		domainsapprovalworkflow.Setup,
		domainsapprovalworkflowassignment.Setup,
		domainsapprovalworkflowstep.Setup,
		domainsauthenticationfactorsetting.Setup,
		domainsauthtoken.Setup,
		domainscloudgate.Setup,
		domainscloudgatemapping.Setup,
		domainscloudgateserver.Setup,
		domainscondition.Setup,
		domainscustomersecretkey.Setup,
		domainsdynamicresourcegroup.Setup,
		domainsgrant.Setup,
		domainsgroup.Setup,
		domainsidentitypropagationtrust.Setup,
		domainsidentityprovider.Setup,
		domainsidentitysetting.Setup,
		domainskmsisetting.Setup,
		domainsmyapikey.Setup,
		domainsmyauthtoken.Setup,
		domainsmycustomersecretkey.Setup,
		domainsmyoauth2clientcredential.Setup,
		domainsmyrequest.Setup,
		domainsmysmtpcredential.Setup,
		domainsmysupportaccount.Setup,
		domainsmyuserdbcredential.Setup,
		domainsnetworkperimeter.Setup,
		domainsnotificationsetting.Setup,
		domainsoauth2clientcredential.Setup,
		domainsoauthclientcertificate.Setup,
		domainsoauthpartnercertificate.Setup,
		domainspasswordpolicy.Setup,
		domainspolicy.Setup,
		domainsrule.Setup,
		domainssecurityquestion.Setup,
		domainssecurityquestionsetting.Setup,
		domainsselfregistrationprofile.Setup,
		domainssetting.Setup,
		domainssmtpcredential.Setup,
		domainssocialidentityprovider.Setup,
		domainsuser.Setup,
		domainsuserdbcredential.Setup,
		dynamicgroup.Setup,
		group.Setup,
		identitydomain.Setup,
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

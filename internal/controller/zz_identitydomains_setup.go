/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	accountrecoverysetting "github.com/oracle/provider-oci/internal/controller/identitydomains/accountrecoverysetting"
	apikey "github.com/oracle/provider-oci/internal/controller/identitydomains/apikey"
	app "github.com/oracle/provider-oci/internal/controller/identitydomains/app"
	approle "github.com/oracle/provider-oci/internal/controller/identitydomains/approle"
	approvalworkflow "github.com/oracle/provider-oci/internal/controller/identitydomains/approvalworkflow"
	approvalworkflowassignment "github.com/oracle/provider-oci/internal/controller/identitydomains/approvalworkflowassignment"
	approvalworkflowstep "github.com/oracle/provider-oci/internal/controller/identitydomains/approvalworkflowstep"
	authenticationfactorsetting "github.com/oracle/provider-oci/internal/controller/identitydomains/authenticationfactorsetting"
	authtoken "github.com/oracle/provider-oci/internal/controller/identitydomains/authtoken"
	cloudgate "github.com/oracle/provider-oci/internal/controller/identitydomains/cloudgate"
	cloudgatemapping "github.com/oracle/provider-oci/internal/controller/identitydomains/cloudgatemapping"
	cloudgateserver "github.com/oracle/provider-oci/internal/controller/identitydomains/cloudgateserver"
	condition "github.com/oracle/provider-oci/internal/controller/identitydomains/condition"
	customersecretkey "github.com/oracle/provider-oci/internal/controller/identitydomains/customersecretkey"
	dynamicresourcegroup "github.com/oracle/provider-oci/internal/controller/identitydomains/dynamicresourcegroup"
	grant "github.com/oracle/provider-oci/internal/controller/identitydomains/grant"
	group "github.com/oracle/provider-oci/internal/controller/identitydomains/group"
	identitypropagationtrust "github.com/oracle/provider-oci/internal/controller/identitydomains/identitypropagationtrust"
	identityprovider "github.com/oracle/provider-oci/internal/controller/identitydomains/identityprovider"
	identitysetting "github.com/oracle/provider-oci/internal/controller/identitydomains/identitysetting"
	kmsisetting "github.com/oracle/provider-oci/internal/controller/identitydomains/kmsisetting"
	myapikey "github.com/oracle/provider-oci/internal/controller/identitydomains/myapikey"
	myauthtoken "github.com/oracle/provider-oci/internal/controller/identitydomains/myauthtoken"
	mycustomersecretkey "github.com/oracle/provider-oci/internal/controller/identitydomains/mycustomersecretkey"
	myoauth2clientcredential "github.com/oracle/provider-oci/internal/controller/identitydomains/myoauth2clientcredential"
	myrequest "github.com/oracle/provider-oci/internal/controller/identitydomains/myrequest"
	mysmtpcredential "github.com/oracle/provider-oci/internal/controller/identitydomains/mysmtpcredential"
	mysupportaccount "github.com/oracle/provider-oci/internal/controller/identitydomains/mysupportaccount"
	myuserdbcredential "github.com/oracle/provider-oci/internal/controller/identitydomains/myuserdbcredential"
	networkperimeter "github.com/oracle/provider-oci/internal/controller/identitydomains/networkperimeter"
	notificationsetting "github.com/oracle/provider-oci/internal/controller/identitydomains/notificationsetting"
	oauth2clientcredential "github.com/oracle/provider-oci/internal/controller/identitydomains/oauth2clientcredential"
	oauthclientcertificate "github.com/oracle/provider-oci/internal/controller/identitydomains/oauthclientcertificate"
	oauthpartnercertificate "github.com/oracle/provider-oci/internal/controller/identitydomains/oauthpartnercertificate"
	passwordpolicy "github.com/oracle/provider-oci/internal/controller/identitydomains/passwordpolicy"
	policy "github.com/oracle/provider-oci/internal/controller/identitydomains/policy"
	rule "github.com/oracle/provider-oci/internal/controller/identitydomains/rule"
	securityquestion "github.com/oracle/provider-oci/internal/controller/identitydomains/securityquestion"
	securityquestionsetting "github.com/oracle/provider-oci/internal/controller/identitydomains/securityquestionsetting"
	selfregistrationprofile "github.com/oracle/provider-oci/internal/controller/identitydomains/selfregistrationprofile"
	setting "github.com/oracle/provider-oci/internal/controller/identitydomains/setting"
	smtpcredential "github.com/oracle/provider-oci/internal/controller/identitydomains/smtpcredential"
	socialidentityprovider "github.com/oracle/provider-oci/internal/controller/identitydomains/socialidentityprovider"
	user "github.com/oracle/provider-oci/internal/controller/identitydomains/user"
	userdbcredential "github.com/oracle/provider-oci/internal/controller/identitydomains/userdbcredential"
)

// Setup_identitydomains creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_identitydomains(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accountrecoverysetting.Setup,
		apikey.Setup,
		app.Setup,
		approle.Setup,
		approvalworkflow.Setup,
		approvalworkflowassignment.Setup,
		approvalworkflowstep.Setup,
		authenticationfactorsetting.Setup,
		authtoken.Setup,
		cloudgate.Setup,
		cloudgatemapping.Setup,
		cloudgateserver.Setup,
		condition.Setup,
		customersecretkey.Setup,
		dynamicresourcegroup.Setup,
		grant.Setup,
		group.Setup,
		identitypropagationtrust.Setup,
		identityprovider.Setup,
		identitysetting.Setup,
		kmsisetting.Setup,
		myapikey.Setup,
		myauthtoken.Setup,
		mycustomersecretkey.Setup,
		myoauth2clientcredential.Setup,
		myrequest.Setup,
		mysmtpcredential.Setup,
		mysupportaccount.Setup,
		myuserdbcredential.Setup,
		networkperimeter.Setup,
		notificationsetting.Setup,
		oauth2clientcredential.Setup,
		oauthclientcertificate.Setup,
		oauthpartnercertificate.Setup,
		passwordpolicy.Setup,
		policy.Setup,
		rule.Setup,
		securityquestion.Setup,
		securityquestionsetting.Setup,
		selfregistrationprofile.Setup,
		setting.Setup,
		smtpcredential.Setup,
		socialidentityprovider.Setup,
		user.Setup,
		userdbcredential.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

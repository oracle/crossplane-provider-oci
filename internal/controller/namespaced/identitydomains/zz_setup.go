/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	accountrecoverysetting "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/accountrecoverysetting"
	apikey "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/apikey"
	app "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/app"
	approle "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/approle"
	approvalworkflow "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/approvalworkflow"
	approvalworkflowassignment "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/approvalworkflowassignment"
	approvalworkflowstep "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/approvalworkflowstep"
	authenticationfactorsetting "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/authenticationfactorsetting"
	authtoken "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/authtoken"
	cloudgate "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/cloudgate"
	cloudgatemapping "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/cloudgatemapping"
	cloudgateserver "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/cloudgateserver"
	condition "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/condition"
	customersecretkey "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/customersecretkey"
	dynamicresourcegroup "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/dynamicresourcegroup"
	grant "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/grant"
	group "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/group"
	identityproofingprovider "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/identityproofingprovider"
	identityproofingprovidertemplate "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/identityproofingprovidertemplate"
	identitypropagationtrust "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/identitypropagationtrust"
	identityprovider "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/identityprovider"
	identitysetting "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/identitysetting"
	kmsisetting "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/kmsisetting"
	mappedattribute "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/mappedattribute"
	myapikey "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/myapikey"
	myauthtoken "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/myauthtoken"
	mycustomersecretkey "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/mycustomersecretkey"
	myoauth2clientcredential "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/myoauth2clientcredential"
	myrequest "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/myrequest"
	mysmtpcredential "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/mysmtpcredential"
	mysupportaccount "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/mysupportaccount"
	myuserdbcredential "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/myuserdbcredential"
	networkperimeter "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/networkperimeter"
	notificationsetting "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/notificationsetting"
	oauth2clientcredential "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/oauth2clientcredential"
	oauthclientcertificate "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/oauthclientcertificate"
	oauthpartnercertificate "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/oauthpartnercertificate"
	passwordpolicy "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/passwordpolicy"
	policy "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/policy"
	rule "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/rule"
	securityquestion "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/securityquestion"
	securityquestionsetting "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/securityquestionsetting"
	selfregistrationprofile "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/selfregistrationprofile"
	setting "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/setting"
	smtpcredential "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/smtpcredential"
	socialidentityprovider "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/socialidentityprovider"
	user "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/user"
	userdbcredential "github.com/oracle/provider-oci/internal/controller/namespaced/identitydomains/userdbcredential"
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
		identityproofingprovider.Setup,
		identityproofingprovidertemplate.Setup,
		identitypropagationtrust.Setup,
		identityprovider.Setup,
		identitysetting.Setup,
		kmsisetting.Setup,
		mappedattribute.Setup,
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

// SetupGated_identitydomains creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_identitydomains(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accountrecoverysetting.SetupGated,
		apikey.SetupGated,
		app.SetupGated,
		approle.SetupGated,
		approvalworkflow.SetupGated,
		approvalworkflowassignment.SetupGated,
		approvalworkflowstep.SetupGated,
		authenticationfactorsetting.SetupGated,
		authtoken.SetupGated,
		cloudgate.SetupGated,
		cloudgatemapping.SetupGated,
		cloudgateserver.SetupGated,
		condition.SetupGated,
		customersecretkey.SetupGated,
		dynamicresourcegroup.SetupGated,
		grant.SetupGated,
		group.SetupGated,
		identityproofingprovider.SetupGated,
		identityproofingprovidertemplate.SetupGated,
		identitypropagationtrust.SetupGated,
		identityprovider.SetupGated,
		identitysetting.SetupGated,
		kmsisetting.SetupGated,
		mappedattribute.SetupGated,
		myapikey.SetupGated,
		myauthtoken.SetupGated,
		mycustomersecretkey.SetupGated,
		myoauth2clientcredential.SetupGated,
		myrequest.SetupGated,
		mysmtpcredential.SetupGated,
		mysupportaccount.SetupGated,
		myuserdbcredential.SetupGated,
		networkperimeter.SetupGated,
		notificationsetting.SetupGated,
		oauth2clientcredential.SetupGated,
		oauthclientcertificate.SetupGated,
		oauthpartnercertificate.SetupGated,
		passwordpolicy.SetupGated,
		policy.SetupGated,
		rule.SetupGated,
		securityquestion.SetupGated,
		securityquestionsetting.SetupGated,
		selfregistrationprofile.SetupGated,
		setting.SetupGated,
		smtpcredential.SetupGated,
		socialidentityprovider.SetupGated,
		user.SetupGated,
		userdbcredential.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

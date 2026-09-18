/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	accountrecoverysetting "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/accountrecoverysetting"
	apikey "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/apikey"
	app "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/app"
	approle "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/approle"
	approvalworkflow "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/approvalworkflow"
	approvalworkflowassignment "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/approvalworkflowassignment"
	approvalworkflowstep "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/approvalworkflowstep"
	authenticationfactorsetting "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/authenticationfactorsetting"
	authtoken "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/authtoken"
	cloudgate "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/cloudgate"
	cloudgatemapping "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/cloudgatemapping"
	cloudgateserver "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/cloudgateserver"
	condition "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/condition"
	customersecretkey "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/customersecretkey"
	dynamicresourcegroup "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/dynamicresourcegroup"
	grant "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/grant"
	group "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/group"
	identityproofingprovider "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/identityproofingprovider"
	identityproofingprovidertemplate "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/identityproofingprovidertemplate"
	identitypropagationtrust "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/identitypropagationtrust"
	identityprovider "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/identityprovider"
	identitysetting "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/identitysetting"
	kmsisetting "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/kmsisetting"
	mappedattribute "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/mappedattribute"
	myapikey "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/myapikey"
	myauthtoken "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/myauthtoken"
	mycustomersecretkey "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/mycustomersecretkey"
	myoauth2clientcredential "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/myoauth2clientcredential"
	myrequest "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/myrequest"
	mysmtpcredential "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/mysmtpcredential"
	mysupportaccount "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/mysupportaccount"
	myuserdbcredential "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/myuserdbcredential"
	networkperimeter "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/networkperimeter"
	notificationsetting "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/notificationsetting"
	oauth2clientcredential "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/oauth2clientcredential"
	oauthclientcertificate "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/oauthclientcertificate"
	oauthpartnercertificate "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/oauthpartnercertificate"
	passwordpolicy "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/passwordpolicy"
	policy "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/policy"
	rule "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/rule"
	securityquestion "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/securityquestion"
	securityquestionsetting "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/securityquestionsetting"
	selfregistrationprofile "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/selfregistrationprofile"
	setting "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/setting"
	smtpcredential "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/smtpcredential"
	socialidentityprovider "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/socialidentityprovider"
	user "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/user"
	userdbcredential "github.com/oracle/provider-oci/internal/controller/cluster/identitydomains/userdbcredential"
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

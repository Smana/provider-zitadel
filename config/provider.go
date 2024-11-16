/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/Smana/provider-zitadel/config/action"
	applicationapi "github.com/Smana/provider-zitadel/config/application_api"
	applicationkey "github.com/Smana/provider-zitadel/config/application_key"
	applicationoidc "github.com/Smana/provider-zitadel/config/application_oidc"
	applicationsaml "github.com/Smana/provider-zitadel/config/application_saml"
	defaultdomainpolicy "github.com/Smana/provider-zitadel/config/default_domain_policy"
	defaultlabelpolicy "github.com/Smana/provider-zitadel/config/default_label_policy"
	defaultlockoutpolicy "github.com/Smana/provider-zitadel/config/default_lockout_policy"
	defaultloginpolicy "github.com/Smana/provider-zitadel/config/default_login_policy"
	defaultnotificationpolicy "github.com/Smana/provider-zitadel/config/default_notification_policy"
	defaultoidcsettings "github.com/Smana/provider-zitadel/config/default_oidc_settings"
	defaultpasswordchangemessagetext "github.com/Smana/provider-zitadel/config/default_password_change_message_text"
	defaultpasswordcomplexitypolicy "github.com/Smana/provider-zitadel/config/default_password_complexity_policy"
	defaultpasswordresetmessagetext "github.com/Smana/provider-zitadel/config/default_password_reset_message_text"
	defaultpasswordlessregistrationmessagetext "github.com/Smana/provider-zitadel/config/default_passwordless_registration_message_text"
	defaultprivacypolicy "github.com/Smana/provider-zitadel/config/default_privacy_policy"
	defaultverifyemailmessagetext "github.com/Smana/provider-zitadel/config/default_verify_email_message_text"
	defaultverifyemailotpmessagetext "github.com/Smana/provider-zitadel/config/default_verify_email_otp_message_text"
	defaultverifyphonemessagetext "github.com/Smana/provider-zitadel/config/default_verify_phone_message_text"
	defaultverifysmsotpmessagetext "github.com/Smana/provider-zitadel/config/default_verify_sms_otp_message_text"
	domain "github.com/Smana/provider-zitadel/config/domain"
	domainclaimedmessagetext "github.com/Smana/provider-zitadel/config/domain_claimed_message_text"
	domainpolicy "github.com/Smana/provider-zitadel/config/domain_policy"
	humanuser "github.com/Smana/provider-zitadel/config/human_user"
	idpazuread "github.com/Smana/provider-zitadel/config/idp_azure_ad"
	idpgithub "github.com/Smana/provider-zitadel/config/idp_github"
	idpgithubes "github.com/Smana/provider-zitadel/config/idp_github_es"
	idpgitlab "github.com/Smana/provider-zitadel/config/idp_gitlab"
	idpgitlabselfhosted "github.com/Smana/provider-zitadel/config/idp_gitlab_self_hosted"
	idpgoogle "github.com/Smana/provider-zitadel/config/idp_google"
	idpldap "github.com/Smana/provider-zitadel/config/idp_ldap"
	idpoauth "github.com/Smana/provider-zitadel/config/idp_oauth"
	idpsaml "github.com/Smana/provider-zitadel/config/idp_saml"
	initmessagetext "github.com/Smana/provider-zitadel/config/init_message_text"
	instancemember "github.com/Smana/provider-zitadel/config/instance_member"
	labelpolicy "github.com/Smana/provider-zitadel/config/label_policy"
	lockoutpolicy "github.com/Smana/provider-zitadel/config/lockout_policy"
	loginpolicy "github.com/Smana/provider-zitadel/config/login_policy"
	machinekey "github.com/Smana/provider-zitadel/config/machine_key"
	machineuser "github.com/Smana/provider-zitadel/config/machine_user"
	notificationpolicy "github.com/Smana/provider-zitadel/config/notification_policy"
	org "github.com/Smana/provider-zitadel/config/org"
	orgidpazuread "github.com/Smana/provider-zitadel/config/org_idp_azure_ad"
	orgidpgithub "github.com/Smana/provider-zitadel/config/org_idp_github"
	orgidpgithubes "github.com/Smana/provider-zitadel/config/org_idp_github_es"
	orgidpgitlab "github.com/Smana/provider-zitadel/config/org_idp_gitlab"
	orgidpgitlabselfhosted "github.com/Smana/provider-zitadel/config/org_idp_gitlab_self_hosted"
	orgidpgoogle "github.com/Smana/provider-zitadel/config/org_idp_google"
	orgidpjwt "github.com/Smana/provider-zitadel/config/org_idp_jwt"
	orgidpldap "github.com/Smana/provider-zitadel/config/org_idp_ldap"
	orgidpoauth "github.com/Smana/provider-zitadel/config/org_idp_oauth"
	orgidpoidc "github.com/Smana/provider-zitadel/config/org_idp_oidc"
	orgidpsaml "github.com/Smana/provider-zitadel/config/org_idp_saml"
	orgmember "github.com/Smana/provider-zitadel/config/org_member"
	orgmetadata "github.com/Smana/provider-zitadel/config/org_metadata"
	passwordchangemessagetext "github.com/Smana/provider-zitadel/config/password_change_message_text"
	passwordcomplexitypolicy "github.com/Smana/provider-zitadel/config/password_complexity_policy"
	passwordresetmessagetext "github.com/Smana/provider-zitadel/config/password_reset_message_text"
	passwordlessregistrationmessagetext "github.com/Smana/provider-zitadel/config/passwordless_registration_message_text"
	personalaccesstoken "github.com/Smana/provider-zitadel/config/personal_access_token"
	privacypolicy "github.com/Smana/provider-zitadel/config/privacy_policy"
	project "github.com/Smana/provider-zitadel/config/project"
	projectgrant "github.com/Smana/provider-zitadel/config/project_grant"
	projectgrantmember "github.com/Smana/provider-zitadel/config/project_grant_member"
	projectmember "github.com/Smana/provider-zitadel/config/project_member"
	projectrole "github.com/Smana/provider-zitadel/config/project_role"
	smsprovidertwilio "github.com/Smana/provider-zitadel/config/sms_provider_twilio"
	smtpconfig "github.com/Smana/provider-zitadel/config/smtp_config"
	triggeractions "github.com/Smana/provider-zitadel/config/trigger_actions"
	usergrant "github.com/Smana/provider-zitadel/config/user_grant"
	usermetadata "github.com/Smana/provider-zitadel/config/user_metadata"
	verifyemailmessagetext "github.com/Smana/provider-zitadel/config/verify_email_message_text"
	verifyemailotpmessagetext "github.com/Smana/provider-zitadel/config/verify_email_otp_message_text"
	verifyphonemessagetext "github.com/Smana/provider-zitadel/config/verify_phone_message_text"
	verifysmsotpmessagetext "github.com/Smana/provider-zitadel/config/verify_sms_otp_message_text"
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	resourcePrefix = "zitadel"
	modulePath     = "github.com/Smana/provider-zitadel"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("zitadel.crossplane.io"),
		ujconfig.WithShortName("zitadel"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		action.Configure,
		applicationapi.Configure,
		applicationkey.Configure,
		applicationoidc.Configure,
		applicationsaml.Configure,
		defaultdomainpolicy.Configure,
		defaultlabelpolicy.Configure,
		defaultlockoutpolicy.Configure,
		defaultloginpolicy.Configure,
		defaultnotificationpolicy.Configure,
		defaultoidcsettings.Configure,
		defaultpasswordchangemessagetext.Configure,
		defaultpasswordcomplexitypolicy.Configure,
		defaultpasswordlessregistrationmessagetext.Configure,
		defaultpasswordresetmessagetext.Configure,
		defaultprivacypolicy.Configure,
		defaultverifyemailmessagetext.Configure,
		defaultverifyemailotpmessagetext.Configure,
		defaultverifyphonemessagetext.Configure,
		defaultverifysmsotpmessagetext.Configure,
		domain.Configure,
		domainclaimedmessagetext.Configure,
		domainpolicy.Configure,
		humanuser.Configure,
		idpazuread.Configure,
		idpgithub.Configure,
		idpgithubes.Configure,
		idpgitlab.Configure,
		idpgitlabselfhosted.Configure,
		idpgoogle.Configure,
		idpldap.Configure,
		idpoauth.Configure,
		idpsaml.Configure,
		initmessagetext.Configure,
		instancemember.Configure,
		labelpolicy.Configure,
		lockoutpolicy.Configure,
		loginpolicy.Configure,
		machinekey.Configure,
		machineuser.Configure,
		notificationpolicy.Configure,
		org.Configure,
		orgidpazuread.Configure,
		orgidpgithub.Configure,
		orgidpgithubes.Configure,
		orgidpgitlab.Configure,
		orgidpgitlabselfhosted.Configure,
		orgidpgoogle.Configure,
		orgidpjwt.Configure,
		orgidpldap.Configure,
		orgidpoauth.Configure,
		orgidpoidc.Configure,
		orgidpsaml.Configure,
		orgmember.Configure,
		orgmetadata.Configure,
		passwordchangemessagetext.Configure,
		passwordcomplexitypolicy.Configure,
		passwordlessregistrationmessagetext.Configure,
		passwordresetmessagetext.Configure,
		personalaccesstoken.Configure,
		privacypolicy.Configure,
		project.Configure,
		projectgrant.Configure,
		projectgrantmember.Configure,
		projectmember.Configure,
		projectrole.Configure,
		smsprovidertwilio.Configure,
		smtpconfig.Configure,
		triggeractions.Configure,
		usergrant.Configure,
		usermetadata.Configure,
		verifyemailmessagetext.Configure,
		verifyemailotpmessagetext.Configure,
		verifyphonemessagetext.Configure,
		verifysmsotpmessagetext.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

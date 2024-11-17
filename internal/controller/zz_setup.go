// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	action "github.com/Smana/provider-zitadel/internal/controller/action/action"
	triggeractions "github.com/Smana/provider-zitadel/internal/controller/action/triggeractions"
	applicationapi "github.com/Smana/provider-zitadel/internal/controller/app/applicationapi"
	applicationkey "github.com/Smana/provider-zitadel/internal/controller/app/applicationkey"
	applicationoidc "github.com/Smana/provider-zitadel/internal/controller/app/applicationoidc"
	applicationsaml "github.com/Smana/provider-zitadel/internal/controller/app/applicationsaml"
	domainclaimedmessagetext "github.com/Smana/provider-zitadel/internal/controller/default/domainclaimedmessagetext"
	initmessagetext "github.com/Smana/provider-zitadel/internal/controller/default/initmessagetext"
	verifysmsotpmessagetext "github.com/Smana/provider-zitadel/internal/controller/default/verifysmsotpmessagetext"
	idpazuread "github.com/Smana/provider-zitadel/internal/controller/idp/idpazuread"
	idpgithub "github.com/Smana/provider-zitadel/internal/controller/idp/idpgithub"
	idpgithubes "github.com/Smana/provider-zitadel/internal/controller/idp/idpgithubes"
	idpgitlab "github.com/Smana/provider-zitadel/internal/controller/idp/idpgitlab"
	idpgitlabselfhosted "github.com/Smana/provider-zitadel/internal/controller/idp/idpgitlabselfhosted"
	idpgoogle "github.com/Smana/provider-zitadel/internal/controller/idp/idpgoogle"
	idpldap "github.com/Smana/provider-zitadel/internal/controller/idp/idpldap"
	idpoauth "github.com/Smana/provider-zitadel/internal/controller/idp/idpoauth"
	idpsaml "github.com/Smana/provider-zitadel/internal/controller/idp/idpsaml"
	orgidpazuread "github.com/Smana/provider-zitadel/internal/controller/idp/orgidpazuread"
	orgidpgithub "github.com/Smana/provider-zitadel/internal/controller/idp/orgidpgithub"
	orgidpgithubes "github.com/Smana/provider-zitadel/internal/controller/idp/orgidpgithubes"
	orgidpgitlab "github.com/Smana/provider-zitadel/internal/controller/idp/orgidpgitlab"
	orgidpgitlabselfhosted "github.com/Smana/provider-zitadel/internal/controller/idp/orgidpgitlabselfhosted"
	orgidpgoogle "github.com/Smana/provider-zitadel/internal/controller/idp/orgidpgoogle"
	orgidpjwt "github.com/Smana/provider-zitadel/internal/controller/idp/orgidpjwt"
	orgidpldap "github.com/Smana/provider-zitadel/internal/controller/idp/orgidpldap"
	orgidpoauth "github.com/Smana/provider-zitadel/internal/controller/idp/orgidpoauth"
	orgidpoidc "github.com/Smana/provider-zitadel/internal/controller/idp/orgidpoidc"
	orgidpsaml "github.com/Smana/provider-zitadel/internal/controller/idp/orgidpsaml"
	machinekey "github.com/Smana/provider-zitadel/internal/controller/machine/machinekey"
	defaultpasswordchangemessagetext "github.com/Smana/provider-zitadel/internal/controller/notif/defaultpasswordchangemessagetext"
	defaultpasswordcomplexitypolicy "github.com/Smana/provider-zitadel/internal/controller/notif/defaultpasswordcomplexitypolicy"
	defaultpasswordlessregistrationmessagetext "github.com/Smana/provider-zitadel/internal/controller/notif/defaultpasswordlessregistrationmessagetext"
	defaultpasswordresetmessagetext "github.com/Smana/provider-zitadel/internal/controller/notif/defaultpasswordresetmessagetext"
	defaultverifyemailmessagetext "github.com/Smana/provider-zitadel/internal/controller/notif/defaultverifyemailmessagetext"
	defaultverifyemailotpmessagetext "github.com/Smana/provider-zitadel/internal/controller/notif/defaultverifyemailotpmessagetext"
	defaultverifyphonemessagetext "github.com/Smana/provider-zitadel/internal/controller/notif/defaultverifyphonemessagetext"
	domainclaimedmessagetextnotif "github.com/Smana/provider-zitadel/internal/controller/notif/domainclaimedmessagetext"
	initmessagetextnotif "github.com/Smana/provider-zitadel/internal/controller/notif/initmessagetext"
	passwordchangemessagetext "github.com/Smana/provider-zitadel/internal/controller/notif/passwordchangemessagetext"
	passwordlessregistrationmessagetext "github.com/Smana/provider-zitadel/internal/controller/notif/passwordlessregistrationmessagetext"
	passwordresetmessagetext "github.com/Smana/provider-zitadel/internal/controller/notif/passwordresetmessagetext"
	smsprovidertwilio "github.com/Smana/provider-zitadel/internal/controller/notif/smsprovidertwilio"
	smtpconfig "github.com/Smana/provider-zitadel/internal/controller/notif/smtpconfig"
	verifyemailmessagetext "github.com/Smana/provider-zitadel/internal/controller/notif/verifyemailmessagetext"
	verifyemailotpmessagetext "github.com/Smana/provider-zitadel/internal/controller/notif/verifyemailotpmessagetext"
	verifyphonemessagetext "github.com/Smana/provider-zitadel/internal/controller/notif/verifyphonemessagetext"
	verifysmsotpmessagetextnotif "github.com/Smana/provider-zitadel/internal/controller/notif/verifysmsotpmessagetext"
	defaultdomainpolicy "github.com/Smana/provider-zitadel/internal/controller/org/defaultdomainpolicy"
	defaultlabelpolicy "github.com/Smana/provider-zitadel/internal/controller/org/defaultlabelpolicy"
	defaultlockoutpolicy "github.com/Smana/provider-zitadel/internal/controller/org/defaultlockoutpolicy"
	defaultloginpolicy "github.com/Smana/provider-zitadel/internal/controller/org/defaultloginpolicy"
	defaultnotificationpolicy "github.com/Smana/provider-zitadel/internal/controller/org/defaultnotificationpolicy"
	defaultoidcsettings "github.com/Smana/provider-zitadel/internal/controller/org/defaultoidcsettings"
	defaultprivacypolicy "github.com/Smana/provider-zitadel/internal/controller/org/defaultprivacypolicy"
	domain "github.com/Smana/provider-zitadel/internal/controller/org/domain"
	org "github.com/Smana/provider-zitadel/internal/controller/org/org"
	orgmember "github.com/Smana/provider-zitadel/internal/controller/org/orgmember"
	orgmetadata "github.com/Smana/provider-zitadel/internal/controller/org/orgmetadata"
	resetmessagetext "github.com/Smana/provider-zitadel/internal/controller/password/resetmessagetext"
	domainpolicy "github.com/Smana/provider-zitadel/internal/controller/policy/domainpolicy"
	labelpolicy "github.com/Smana/provider-zitadel/internal/controller/policy/labelpolicy"
	lockoutpolicy "github.com/Smana/provider-zitadel/internal/controller/policy/lockoutpolicy"
	loginpolicy "github.com/Smana/provider-zitadel/internal/controller/policy/loginpolicy"
	notificationpolicy "github.com/Smana/provider-zitadel/internal/controller/policy/notificationpolicy"
	privacypolicy "github.com/Smana/provider-zitadel/internal/controller/policy/privacypolicy"
	project "github.com/Smana/provider-zitadel/internal/controller/project/project"
	projectgrant "github.com/Smana/provider-zitadel/internal/controller/project/projectgrant"
	projectmember "github.com/Smana/provider-zitadel/internal/controller/project/projectmember"
	projectrole "github.com/Smana/provider-zitadel/internal/controller/project/projectrole"
	providerconfig "github.com/Smana/provider-zitadel/internal/controller/providerconfig"
	humanuser "github.com/Smana/provider-zitadel/internal/controller/user/humanuser"
	instancemember "github.com/Smana/provider-zitadel/internal/controller/user/instancemember"
	machineuser "github.com/Smana/provider-zitadel/internal/controller/user/machineuser"
	personalaccesstoken "github.com/Smana/provider-zitadel/internal/controller/user/personalaccesstoken"
	projectgrantmember "github.com/Smana/provider-zitadel/internal/controller/user/projectgrantmember"
	usergrant "github.com/Smana/provider-zitadel/internal/controller/user/usergrant"
	usermetadata "github.com/Smana/provider-zitadel/internal/controller/user/usermetadata"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		action.Setup,
		triggeractions.Setup,
		applicationapi.Setup,
		applicationkey.Setup,
		applicationoidc.Setup,
		applicationsaml.Setup,
		domainclaimedmessagetext.Setup,
		initmessagetext.Setup,
		verifysmsotpmessagetext.Setup,
		idpazuread.Setup,
		idpgithub.Setup,
		idpgithubes.Setup,
		idpgitlab.Setup,
		idpgitlabselfhosted.Setup,
		idpgoogle.Setup,
		idpldap.Setup,
		idpoauth.Setup,
		idpsaml.Setup,
		orgidpazuread.Setup,
		orgidpgithub.Setup,
		orgidpgithubes.Setup,
		orgidpgitlab.Setup,
		orgidpgitlabselfhosted.Setup,
		orgidpgoogle.Setup,
		orgidpjwt.Setup,
		orgidpldap.Setup,
		orgidpoauth.Setup,
		orgidpoidc.Setup,
		orgidpsaml.Setup,
		machinekey.Setup,
		defaultpasswordchangemessagetext.Setup,
		defaultpasswordcomplexitypolicy.Setup,
		defaultpasswordlessregistrationmessagetext.Setup,
		defaultpasswordresetmessagetext.Setup,
		defaultverifyemailmessagetext.Setup,
		defaultverifyemailotpmessagetext.Setup,
		defaultverifyphonemessagetext.Setup,
		domainclaimedmessagetextnotif.Setup,
		initmessagetextnotif.Setup,
		passwordchangemessagetext.Setup,
		passwordlessregistrationmessagetext.Setup,
		passwordresetmessagetext.Setup,
		smsprovidertwilio.Setup,
		smtpconfig.Setup,
		verifyemailmessagetext.Setup,
		verifyemailotpmessagetext.Setup,
		verifyphonemessagetext.Setup,
		verifysmsotpmessagetextnotif.Setup,
		defaultdomainpolicy.Setup,
		defaultlabelpolicy.Setup,
		defaultlockoutpolicy.Setup,
		defaultloginpolicy.Setup,
		defaultnotificationpolicy.Setup,
		defaultoidcsettings.Setup,
		defaultprivacypolicy.Setup,
		domain.Setup,
		org.Setup,
		orgmember.Setup,
		orgmetadata.Setup,
		resetmessagetext.Setup,
		domainpolicy.Setup,
		labelpolicy.Setup,
		lockoutpolicy.Setup,
		loginpolicy.Setup,
		notificationpolicy.Setup,
		privacypolicy.Setup,
		project.Setup,
		projectgrant.Setup,
		projectmember.Setup,
		projectrole.Setup,
		providerconfig.Setup,
		humanuser.Setup,
		instancemember.Setup,
		machineuser.Setup,
		personalaccesstoken.Setup,
		projectgrantmember.Setup,
		usergrant.Setup,
		usermetadata.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

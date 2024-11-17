package defaultpasswordlessregistrationmessagetext

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_default_passwordless_registration_message_text resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_default_passwordless_registration_message_text", func(r *config.Resource) {
		r.Kind = "DefaultPasswordLessRegistrationMessageText"
		r.ShortGroup = "notif"
	})
}

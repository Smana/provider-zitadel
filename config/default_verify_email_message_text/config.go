package defaultverifyemailmessagetext

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_default_verify_email_message_text resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_default_verify_email_message_text", func(r *config.Resource) {
		r.Kind = "DefaultVerifyEmailMessageText"
		r.ShortGroup = "notif"
	})
}

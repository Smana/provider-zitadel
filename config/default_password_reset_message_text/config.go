package defaultpasswordresetmessagetext

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_default_password_reset_message_text resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_default_password_reset_message_text", func(r *config.Resource) {
		r.Kind = "DefaultPasswordResetMessageText"
		r.ShortGroup = "notif"
	})
}

package passwordresetmessagetext

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_password_reset_message_text resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_password_complexity_policy", func(r *config.Resource) {
		r.Kind = "PasswordResetMessageText"
		r.ShortGroup = "notif"
	})
}

package passwordchangemessagetext

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_password_change_message_text resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_password_change_message_text", func(r *config.Resource) {
		r.Kind = "PasswordChangeMessageText"
		r.ShortGroup = "notif"
	})
}

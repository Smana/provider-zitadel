package defaultpasswordchangemessagetext

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_default_password_change_message_text resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_default_password_change_message_text", func(r *config.Resource) {
		r.Kind = "DefaultPasswordChangeMessageText"
		r.ShortGroup = "notif"
	})
}

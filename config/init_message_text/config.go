package initmessagetext

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_init_message_text resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_init_message_text", func(r *config.Resource) {
		r.Kind = "InitMessageText"
		r.ShortGroup = "notif"
	})
}

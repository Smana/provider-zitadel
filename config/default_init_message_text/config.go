package defaultinitmessagetext

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_default_init_message_text resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_default_init_message_text", func(r *config.Resource) {
		r.Kind = "DefaultInitMessageText"
		r.ShortGroup = "org"
	})
}

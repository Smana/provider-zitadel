package defaultdomainclaimedmessagetext

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_default_domain_claimed_message_text resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_default_domain_claimed_message_text", func(r *config.Resource) {
		r.Kind = "DefaultDomainClaimedMessageText"
		r.ShortGroup = "org"
	})
}

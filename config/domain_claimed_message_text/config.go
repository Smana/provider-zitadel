package domainclaimedmessagetext

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_domain_claimed_message_text resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_domain_claimed_message_text", func(r *config.Resource) {
		r.Kind = "DomainClaimedMessageText"
		r.ShortGroup = "notif"
	})
}

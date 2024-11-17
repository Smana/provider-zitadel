package defaultdomainpolicy

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_default_domain_policy resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_default_domain_policy", func(r *config.Resource) {
		r.Kind = "DefaultDomainPolicy"
		r.ShortGroup = "org"
	})
}

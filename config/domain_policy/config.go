package domainpolicy

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_domain_policy resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_domain_policy", func(r *config.Resource) {
		r.Kind = "DomainPolicy"
		r.ShortGroup = "policy"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

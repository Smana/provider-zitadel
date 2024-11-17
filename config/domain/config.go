package domain

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_domain resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_domain", func(r *config.Resource) {
		r.Kind = "Domain"
		r.ShortGroup = "org"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

package loginpolicy

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_login_policy resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_login_policy", func(r *config.Resource) {
		r.Kind = "LoginPolicy"
		r.ShortGroup = "policy"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

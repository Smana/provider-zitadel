package defaultpasswordcomplexitypolicy

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_default_password_complexity_policy resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_default_password_complexity_policy", func(r *config.Resource) {
		r.Kind = "DefaultPasswordComplexityPolicy"
		r.ShortGroup = "notif"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

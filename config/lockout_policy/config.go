package lockoutpolicy

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_lockout_policy resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_lockout_policy", func(r *config.Resource) {
		r.Kind = "LockoutPolicy"
		r.ShortGroup = "policy"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

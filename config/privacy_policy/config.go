package privacypolicy

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_privacy_policy resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_privacy_policy", func(r *config.Resource) {
		r.Kind = "PrivacyPolicy"
		r.ShortGroup = "policy"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

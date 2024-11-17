package defaultlogintexts

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_default_login_texts resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_default_login_texts", func(r *config.Resource) {
		r.Kind = "DefaultLoginTexts"
		r.ShortGroup = "dlogintexts"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

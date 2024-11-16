package logintexts

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_login_texts resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_login_texts", func(r *config.Resource) {
		r.Kind = "LoginTexts"
		r.ShortGroup = "org"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

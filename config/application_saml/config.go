package applicationsaml

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_application_saml resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_application_saml", func(r *config.Resource) {
		r.Kind = "ApplicationSaml"
		r.ShortGroup = "app"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}

		r.References["project_id"] = config.Reference{
			TerraformName: "zitadel_project",
		}
	})
}

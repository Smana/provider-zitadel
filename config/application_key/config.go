package applicationkey

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_application_key resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_application_key", func(r *config.Resource) {
		r.Kind = "ApplicationKey"
		r.ShortGroup = "app"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}

		r.References["project_id"] = config.Reference{
			TerraformName: "zitadel_project",
		}

		r.References["app_id"] = config.Reference{
			TerraformName: "zitadel_application_api",
		}
	})
}

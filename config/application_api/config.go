package applicationapi

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_application_api resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_application_api", func(r *config.Resource) {
		r.Kind = "ApplicationApi"
		r.ShortGroup = "app"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}

		r.References["project_id"] = config.Reference{
			TerraformName: "zitadel_project",
		}
	})
}

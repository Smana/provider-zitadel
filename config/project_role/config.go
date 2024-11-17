package projectrole

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_project_role resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_project_role", func(r *config.Resource) {
		r.Kind = "ProjectRole"
		r.ShortGroup = "project"

		r.References["project_id"] = config.Reference{
			TerraformName: "zitadel_project",
		}

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

package projectmember

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_project_member resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_project_member", func(r *config.Resource) {
		r.Kind = "ProjectMember"
		r.ShortGroup = "project"

		r.References["project_id"] = config.Reference{
			TerraformName: "zitadel_project",
		}

		r.References["user_id"] = config.Reference{
			TerraformName: "zitadel_human_user",
		}

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

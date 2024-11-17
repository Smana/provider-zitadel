package usergrant

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_user_grant resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_user_grant", func(r *config.Resource) {
		r.Kind = "UserGrant"
		r.ShortGroup = "user"

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

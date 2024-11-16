package usermetadata

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_user_metadata resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_user_metadata", func(r *config.Resource) {
		r.Kind = "UserMetadata"
		r.ShortGroup = "user"

		r.References["user_id"] = config.Reference{
			TerraformName: "zitadel_human_user",
		}

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

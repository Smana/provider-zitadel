package projectgrant

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_project_grant resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_project_grant", func(r *config.Resource) {
		r.Kind = "ProjectGrant"
		r.ShortGroup = "project"

		r.References["granted_org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}

		r.References["project_id"] = config.Reference{
			TerraformName: "zitadel_project",
		}

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

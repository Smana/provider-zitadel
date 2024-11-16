package project

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_project resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_project", func(r *config.Resource) {
		r.Kind = "Project"
		r.ShortGroup = "project"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

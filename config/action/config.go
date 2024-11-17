package action

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_action resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_action", func(r *config.Resource) {
		r.Kind = "Action"
		r.ShortGroup = "action"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

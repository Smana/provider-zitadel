package humanuser

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_human_user resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_human_user", func(r *config.Resource) {
		r.Kind = "HumanUser"
		r.ShortGroup = "user"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

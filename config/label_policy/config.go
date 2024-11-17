package labelpolicy

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_label_policy resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_label_policy", func(r *config.Resource) {
		r.Kind = "LabelPolicy"
		r.ShortGroup = "policy"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

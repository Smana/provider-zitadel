package orgmember

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_org_member resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_org_member", func(r *config.Resource) {
		r.Kind = "OrgMember"
		r.ShortGroup = "org"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

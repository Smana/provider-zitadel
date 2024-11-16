package orgidpoauth

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_org_idp_oauth resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_org_idp_oauth", func(r *config.Resource) {
		r.Kind = "OrgIdpOauth"
		r.ShortGroup = "idp"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

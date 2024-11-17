package orgidpjwt

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_org_idp_jwt resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_org_idp_jwt", func(r *config.Resource) {
		r.Kind = "OrgIdpJwt"
		r.ShortGroup = "idp"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

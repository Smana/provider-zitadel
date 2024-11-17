package orgidpoidc

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_org_idp_oidc resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_org_idp_oidc", func(r *config.Resource) {
		r.Kind = "OrgIdpOidc"
		r.ShortGroup = "idp"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

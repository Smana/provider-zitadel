package orgidpsaml

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_org_idp_saml resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_org_idp_saml", func(r *config.Resource) {
		r.Kind = "OrgIdpSaml"
		r.ShortGroup = "idp"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

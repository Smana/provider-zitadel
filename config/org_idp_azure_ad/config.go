package orgidpazuread

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_org_idp_azure_ad resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_org_idp_azure_ad", func(r *config.Resource) {
		r.Kind = "OrgIdpAzureAd"
		r.ShortGroup = "idp"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

package orgidpgitlab

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_org_idp_gitlab resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_org_idp_gitlab", func(r *config.Resource) {
		r.Kind = "OrgIdpGitlab"
		r.ShortGroup = "idp"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

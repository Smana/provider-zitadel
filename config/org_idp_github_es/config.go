package orgidpgithubes

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_org_idp_github_es resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_org_idp_github_es", func(r *config.Resource) {
		r.Kind = "OrgIdpGithubEs"
		r.ShortGroup = "idp"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

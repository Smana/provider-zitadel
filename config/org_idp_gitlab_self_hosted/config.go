package orgidpgitlabselfhosted

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_org_idp_gitlab_self_hosted resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_org_idp_gitlab_self_hosted", func(r *config.Resource) {
		r.Kind = "OrgIdpGitlabSelfHosted"
		r.ShortGroup = "idp"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

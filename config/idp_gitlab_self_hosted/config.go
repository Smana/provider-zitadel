package idpgitlabselfhosted

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_idp_gitlab_self_hosted resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_idp_gitlab_self_hosted", func(r *config.Resource) {
		r.Kind = "IdpGitlabSelfHosted"
		r.ShortGroup = "idp"
	})
}

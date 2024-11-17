package idpgithubes

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_idp_github_es resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_idp_github_es", func(r *config.Resource) {
		r.Kind = "IdpGithubEs"
		r.ShortGroup = "idp"
	})
}

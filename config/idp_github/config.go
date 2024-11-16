package idpgithub

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_idp_github resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_idp_github", func(r *config.Resource) {
		r.Kind = "IdpGithub"
		r.ShortGroup = "idp"
	})
}

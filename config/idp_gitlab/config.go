package idpgitlab

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_idp_gitlab resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_idp_gitlab", func(r *config.Resource) {
		r.Kind = "IdpGitlab"
		r.ShortGroup = "idp"
	})
}

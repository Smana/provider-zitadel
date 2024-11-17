package idpgoogle

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_idp_google resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_idp_google", func(r *config.Resource) {
		r.Kind = "IdpGoogle"
		r.ShortGroup = "idp"
	})
}

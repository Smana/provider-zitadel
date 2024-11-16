package idpoauth

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_idp_oauth resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_idp_oauth", func(r *config.Resource) {
		r.Kind = "IdpOauth"
		r.ShortGroup = "idp"
	})
}

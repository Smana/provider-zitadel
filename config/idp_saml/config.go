package idpsaml

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_idp_saml resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_idp_saml", func(r *config.Resource) {
		r.Kind = "IdpSaml"
		r.ShortGroup = "idp"
	})
}

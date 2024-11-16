package idpazuread

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_idp_azure_ad resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_idp_azure_ad", func(r *config.Resource) {
		r.Kind = "IdpAzureAd"
		r.ShortGroup = "idp"
	})
}

package defaultoidcsettings

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_default_oidc_settings resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_default_oidc_settings", func(r *config.Resource) {
		r.Kind = "DefaultOIDCSettings"
		r.ShortGroup = "org"
	})
}

package defaultloginpolicy

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_default_login_policy resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_default_login_policy", func(r *config.Resource) {
		r.Kind = "DefaultLoginPolicy"
		r.ShortGroup = "org"
	})
}

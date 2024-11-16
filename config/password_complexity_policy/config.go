package passwordcomplexitypolicy

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_password_complexity_policy resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_password_complexity_policy", func(r *config.Resource) {
		r.Kind = "PasswordComplexityPolicy"
		r.ShortGroup = "policy"
	})
}

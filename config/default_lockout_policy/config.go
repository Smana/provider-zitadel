package defaultlockoutpolicy

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_default_lockout_policy resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_default_lockout_policy", func(r *config.Resource) {
		r.Kind = "DefaultLockoutPolicy"
		r.ShortGroup = "org"
	})
}

package defaultprivacypolicy

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_default_privacy_policy resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_default_privacy_policy", func(r *config.Resource) {
		r.Kind = "DefaultPrivacyPolicy"
		r.ShortGroup = "org"
	})
}

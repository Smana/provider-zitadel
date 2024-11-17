package defaultlabelpolicy

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_default_label_policy resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_default_label_policy", func(r *config.Resource) {
		r.Kind = "DefaultLabelPolicy"
		r.ShortGroup = "org"
	})
}

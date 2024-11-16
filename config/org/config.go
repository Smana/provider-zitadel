package org

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_org resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_org", func(r *config.Resource) {
		r.Kind = "Org"
		r.ShortGroup = "org"
	})
}

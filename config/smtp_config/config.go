package smtpconfig

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_smtp_config resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_smtp_config", func(r *config.Resource) {
		r.Kind = "SmtpConfig"
		r.ShortGroup = "notif"
	})
}

package smsprovidertwilio

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_sms_provider_twilio resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_sms_provider_twilio", func(r *config.Resource) {
		r.Kind = "SmsProviderTwilio"
		r.ShortGroup = "notif"
	})
}

package defaultverifysmsotpmessagetext

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_default_sms_otp_message_text resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_default_sms_otp_message_text", func(r *config.Resource) {
		r.Kind = "DefaultVerifySmsOtpMessageText"
		r.ShortGroup = "notif"
	})
}

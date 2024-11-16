package verifyemailotpmessagetext

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_verify_email_otp_message_text resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_verify_email_otp_message_text", func(r *config.Resource) {
		r.Kind = "VerifyEmailOtpMessageText"
		r.ShortGroup = "notif"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

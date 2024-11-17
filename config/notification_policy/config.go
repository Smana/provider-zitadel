package notificationpolicy

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_notification_policy resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_notification_policy", func(r *config.Resource) {
		r.Kind = "NotificationPolicy"
		r.ShortGroup = "policy"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

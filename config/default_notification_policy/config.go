package defaultnotificationpolicy

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_default_notification_policy resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_default_notification_policy", func(r *config.Resource) {
		r.Kind = "DefaultNotificationPolicy"
		r.ShortGroup = "org"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

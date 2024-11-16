package triggeractions

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_trigger_actions resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_trigger_actions", func(r *config.Resource) {
		r.Kind = "TriggerActions"
		r.ShortGroup = "action"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

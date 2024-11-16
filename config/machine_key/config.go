package machinekey

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_machine_key resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_machine_key", func(r *config.Resource) {
		r.Kind = "MachineKey"
		r.ShortGroup = "machine"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}

		r.References["user_id"] = config.Reference{
			TerraformName: "zitadel_machine_user",
		}
	})
}

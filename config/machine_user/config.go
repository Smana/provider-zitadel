package machineuser

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_machine_user resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_machine_user", func(r *config.Resource) {
		r.Kind = "MachineUser"
		r.ShortGroup = "user"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}

	})
}

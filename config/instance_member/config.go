package instancemember

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_instance_member resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_instance_member", func(r *config.Resource) {
		r.Kind = "InstanceMember"
		r.ShortGroup = "user"

		r.References["user_id"] = config.Reference{
			TerraformName: "zitadel_human_user",
		}
	})
}

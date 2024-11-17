package personalaccesstoken

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_personal_access_token resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_personal_access_token", func(r *config.Resource) {
		r.Kind = "PersonalAccessToken"
		r.ShortGroup = "user"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}

		r.References["user_id"] = config.Reference{
			TerraformName: "zitadel_machine_user",
		}
	})
}

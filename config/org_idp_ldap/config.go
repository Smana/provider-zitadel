package orgidpldap

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_org_idp_ldap resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_org_idp_ldap", func(r *config.Resource) {
		r.Kind = "OrgIdpLdap"
		r.ShortGroup = "idp"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

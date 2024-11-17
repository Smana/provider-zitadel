package orgmetadata

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_org_metadata resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_org_metadata", func(r *config.Resource) {
		r.Kind = "OrgMetadata"
		r.ShortGroup = "org"

		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}

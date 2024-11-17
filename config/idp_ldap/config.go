package idpldap

import "github.com/crossplane/upjet/pkg/config"

// Configure zitadel_idp_ldap resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_idp_ldap", func(r *config.Resource) {
		r.Kind = "IdpLdap"
		r.ShortGroup = "idp"
	})
}

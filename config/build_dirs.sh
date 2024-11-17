#!/bin/bash

# Define the list of Zitadel resources
resources=(
  "zitadel_action"
  "zitadel_application_api"
  "zitadel_application_key"
  "zitadel_application_oidc"
  "zitadel_application_saml"
  "zitadel_default_domain_policy"
  "zitadel_default_label_policy"
  "zitadel_default_lockout_policy"
  "zitadel_default_login_policy"
  "zitadel_default_notification_policy"
  "zitadel_default_password_complexity_policy"
  "zitadel_default_privacy_policy"
  "zitadel_domain"
  "zitadel_domain_policy"
  "zitadel_human_user"
  "zitadel_idp_azure_ad"
  "zitadel_idp_github"
  "zitadel_idp_github_es"
  "zitadel_idp_gitlab"
  "zitadel_idp_gitlab_self_hosted"
  "zitadel_idp_google"
  "zitadel_idp_ldap"
  "zitadel_idp_oauth"
  "zitadel_idp_saml"
  "zitadel_instance_member"
  "zitadel_label_policy"
  "zitadel_lockout_policy"
  "zitadel_login_policy"
  "zitadel_machine_key"
  "zitadel_machine_user"
  "zitadel_notification_policy"
  "zitadel_org"
  "zitadel_org_idp_azure_ad"
  "zitadel_org_idp_github"
  "zitadel_org_idp_github_es"
  "zitadel_org_idp_gitlab"
  "zitadel_org_idp_gitlab_self_hosted"
  "zitadel_org_idp_google"
  "zitadel_org_idp_jwt"
  "zitadel_org_idp_ldap"
  "zitadel_org_idp_oauth"
  "zitadel_org_idp_oidc"
  "zitadel_org_idp_saml"
  "zitadel_org_member"
  "zitadel_org_metadata"
  "zitadel_personal_access_token"
  "zitadel_privacy_policy"
  "zitadel_project"
  "zitadel_project_grant"
  "zitadel_project_grant_member"
  "zitadel_project_member"
  "zitadel_project_role"
  "zitadel_sms_provider_twilio"
  "zitadel_smtp_config"
  "zitadel_trigger_actions"
  "zitadel_user_grant"
  "zitadel_user_metadata"
)

# Function to convert a snake_case string to PascalCase
snake_to_pascal() {
  echo "$1" | sed -r 's/(^|_)([a-z])/\U\2/g'
}

# Function to convert a snake_case string to camelCase
snake_to_camel() {
  echo "$1" | sed -r 's/_(.)/\U\1/g'
}

# Create directories and files
for resource in "${resources[@]}"; do
  # Remove prefix 'zitadel_'
  dir_name="${resource#zitadel_}"

  # Create directory
  mkdir -p "$dir_name"

  # Format names for placeholders
  pascal_case_name=$(snake_to_pascal "$dir_name")
  camel_case_name=$(snake_to_camel "$dir_name")

  # Write the config.go file
  cat <<EOF > "$dir_name/config.go"
package ${dir_name//_}

import "github.com/crossplane/upjet/pkg/config"

// Configure $resource resource
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("$resource", func(r *config.Resource) {
        r.Kind = "$pascal_case_name"
        r.ShortGroup = "$camel_case_name"

        r.References["team_id"] = config.Reference{
            TerraformName: "$resource",
        }
    })
}
EOF

  echo "Created $dir_name/config.go"
done

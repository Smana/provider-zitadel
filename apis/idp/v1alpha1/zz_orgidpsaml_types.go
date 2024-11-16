// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type OrgIdpSamlInitParameters struct {

	// (String) The binding, supported values: SAML_BINDING_UNSPECIFIED, SAML_BINDING_POST, SAML_BINDING_REDIRECT, SAML_BINDING_ARTIFACT
	// The binding, supported values: SAML_BINDING_UNSPECIFIED, SAML_BINDING_POST, SAML_BINDING_REDIRECT, SAML_BINDING_ARTIFACT
	Binding *string `json:"binding,omitempty" tf:"binding,omitempty"`

	// (Boolean) enable if a new account in ZITADEL should be created automatically on login with an external account
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation *bool `json:"isAutoCreation,omitempty" tf:"is_auto_creation,omitempty"`

	// (Boolean) enable if a the ZITADEL account fields should be updated automatically on each login
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate *bool `json:"isAutoUpdate,omitempty" tf:"is_auto_update,omitempty"`

	// (Boolean) enable if users should be able to create a new account in ZITADEL when using an external account
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed *bool `json:"isCreationAllowed,omitempty" tf:"is_creation_allowed,omitempty"`

	// (Boolean) enable if users should be able to link an existing ZITADEL user with an external account
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed *bool `json:"isLinkingAllowed,omitempty" tf:"is_linking_allowed,omitempty"`

	// (String) The metadata XML as plain string
	// The metadata XML as plain string
	MetadataXML *string `json:"metadataXml,omitempty" tf:"metadata_xml,omitempty"`

	// (String) Name of the IDP
	// Name of the IDP
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) ID of the organization
	// ID of the organization
	// +crossplane:generate:reference:type=github.com/Smana/provider-zitadel/apis/org/v1alpha1.Org
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Reference to a Org in org to populate orgId.
	// +kubebuilder:validation:Optional
	OrgIDRef *v1.Reference `json:"orgIdRef,omitempty" tf:"-"`

	// Selector for a Org in org to populate orgId.
	// +kubebuilder:validation:Optional
	OrgIDSelector *v1.Selector `json:"orgIdSelector,omitempty" tf:"-"`

	// (Boolean) Whether the SAML IDP requires signed requests
	// Whether the SAML IDP requires signed requests
	WithSignedRequest *bool `json:"withSignedRequest,omitempty" tf:"with_signed_request,omitempty"`
}

type OrgIdpSamlObservation struct {

	// (String) The binding, supported values: SAML_BINDING_UNSPECIFIED, SAML_BINDING_POST, SAML_BINDING_REDIRECT, SAML_BINDING_ARTIFACT
	// The binding, supported values: SAML_BINDING_UNSPECIFIED, SAML_BINDING_POST, SAML_BINDING_REDIRECT, SAML_BINDING_ARTIFACT
	Binding *string `json:"binding,omitempty" tf:"binding,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Boolean) enable if a new account in ZITADEL should be created automatically on login with an external account
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation *bool `json:"isAutoCreation,omitempty" tf:"is_auto_creation,omitempty"`

	// (Boolean) enable if a the ZITADEL account fields should be updated automatically on each login
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate *bool `json:"isAutoUpdate,omitempty" tf:"is_auto_update,omitempty"`

	// (Boolean) enable if users should be able to create a new account in ZITADEL when using an external account
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed *bool `json:"isCreationAllowed,omitempty" tf:"is_creation_allowed,omitempty"`

	// (Boolean) enable if users should be able to link an existing ZITADEL user with an external account
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed *bool `json:"isLinkingAllowed,omitempty" tf:"is_linking_allowed,omitempty"`

	// (String) The metadata XML as plain string
	// The metadata XML as plain string
	MetadataXML *string `json:"metadataXml,omitempty" tf:"metadata_xml,omitempty"`

	// (String) Name of the IDP
	// Name of the IDP
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) ID of the organization
	// ID of the organization
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// (Boolean) Whether the SAML IDP requires signed requests
	// Whether the SAML IDP requires signed requests
	WithSignedRequest *bool `json:"withSignedRequest,omitempty" tf:"with_signed_request,omitempty"`
}

type OrgIdpSamlParameters struct {

	// (String) The binding, supported values: SAML_BINDING_UNSPECIFIED, SAML_BINDING_POST, SAML_BINDING_REDIRECT, SAML_BINDING_ARTIFACT
	// The binding, supported values: SAML_BINDING_UNSPECIFIED, SAML_BINDING_POST, SAML_BINDING_REDIRECT, SAML_BINDING_ARTIFACT
	// +kubebuilder:validation:Optional
	Binding *string `json:"binding,omitempty" tf:"binding,omitempty"`

	// (Boolean) enable if a new account in ZITADEL should be created automatically on login with an external account
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	// +kubebuilder:validation:Optional
	IsAutoCreation *bool `json:"isAutoCreation,omitempty" tf:"is_auto_creation,omitempty"`

	// (Boolean) enable if a the ZITADEL account fields should be updated automatically on each login
	// enable if a the ZITADEL account fields should be updated automatically on each login
	// +kubebuilder:validation:Optional
	IsAutoUpdate *bool `json:"isAutoUpdate,omitempty" tf:"is_auto_update,omitempty"`

	// (Boolean) enable if users should be able to create a new account in ZITADEL when using an external account
	// enable if users should be able to create a new account in ZITADEL when using an external account
	// +kubebuilder:validation:Optional
	IsCreationAllowed *bool `json:"isCreationAllowed,omitempty" tf:"is_creation_allowed,omitempty"`

	// (Boolean) enable if users should be able to link an existing ZITADEL user with an external account
	// enable if users should be able to link an existing ZITADEL user with an external account
	// +kubebuilder:validation:Optional
	IsLinkingAllowed *bool `json:"isLinkingAllowed,omitempty" tf:"is_linking_allowed,omitempty"`

	// (String) The metadata XML as plain string
	// The metadata XML as plain string
	// +kubebuilder:validation:Optional
	MetadataXML *string `json:"metadataXml,omitempty" tf:"metadata_xml,omitempty"`

	// (String) Name of the IDP
	// Name of the IDP
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) ID of the organization
	// ID of the organization
	// +crossplane:generate:reference:type=github.com/Smana/provider-zitadel/apis/org/v1alpha1.Org
	// +kubebuilder:validation:Optional
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Reference to a Org in org to populate orgId.
	// +kubebuilder:validation:Optional
	OrgIDRef *v1.Reference `json:"orgIdRef,omitempty" tf:"-"`

	// Selector for a Org in org to populate orgId.
	// +kubebuilder:validation:Optional
	OrgIDSelector *v1.Selector `json:"orgIdSelector,omitempty" tf:"-"`

	// (Boolean) Whether the SAML IDP requires signed requests
	// Whether the SAML IDP requires signed requests
	// +kubebuilder:validation:Optional
	WithSignedRequest *bool `json:"withSignedRequest,omitempty" tf:"with_signed_request,omitempty"`
}

// OrgIdpSamlSpec defines the desired state of OrgIdpSaml
type OrgIdpSamlSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrgIdpSamlParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider OrgIdpSamlInitParameters `json:"initProvider,omitempty"`
}

// OrgIdpSamlStatus defines the observed state of OrgIdpSaml.
type OrgIdpSamlStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrgIdpSamlObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// OrgIdpSaml is the Schema for the OrgIdpSamls API. Resource representing a SAML IdP on the organization.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,zitadel}
type OrgIdpSaml struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.isAutoCreation) || (has(self.initProvider) && has(self.initProvider.isAutoCreation))",message="spec.forProvider.isAutoCreation is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.isAutoUpdate) || (has(self.initProvider) && has(self.initProvider.isAutoUpdate))",message="spec.forProvider.isAutoUpdate is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.isCreationAllowed) || (has(self.initProvider) && has(self.initProvider.isCreationAllowed))",message="spec.forProvider.isCreationAllowed is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.isLinkingAllowed) || (has(self.initProvider) && has(self.initProvider.isLinkingAllowed))",message="spec.forProvider.isLinkingAllowed is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.metadataXml) || (has(self.initProvider) && has(self.initProvider.metadataXml))",message="spec.forProvider.metadataXml is a required parameter"
	Spec   OrgIdpSamlSpec   `json:"spec"`
	Status OrgIdpSamlStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrgIdpSamlList contains a list of OrgIdpSamls
type OrgIdpSamlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrgIdpSaml `json:"items"`
}

// Repository type metadata.
var (
	OrgIdpSaml_Kind             = "OrgIdpSaml"
	OrgIdpSaml_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OrgIdpSaml_Kind}.String()
	OrgIdpSaml_KindAPIVersion   = OrgIdpSaml_Kind + "." + CRDGroupVersion.String()
	OrgIdpSaml_GroupVersionKind = CRDGroupVersion.WithKind(OrgIdpSaml_Kind)
)

func init() {
	SchemeBuilder.Register(&OrgIdpSaml{}, &OrgIdpSamlList{})
}

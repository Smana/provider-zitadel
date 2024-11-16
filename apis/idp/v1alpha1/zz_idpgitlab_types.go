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

type IdpGitlabInitParameters struct {

	// (String) client id generated by the identity provider
	// client id generated by the identity provider
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// (String, Sensitive) client secret generated by the identity provider
	// client secret generated by the identity provider
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

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

	// (String) Name of the IDP
	// Name of the IDP
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Set of String) the scopes requested by ZITADEL during the request on the identity provider
	// the scopes requested by ZITADEL during the request on the identity provider
	// +listType=set
	Scopes []*string `json:"scopes,omitempty" tf:"scopes,omitempty"`
}

type IdpGitlabObservation struct {

	// (String) client id generated by the identity provider
	// client id generated by the identity provider
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

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

	// (String) Name of the IDP
	// Name of the IDP
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Set of String) the scopes requested by ZITADEL during the request on the identity provider
	// the scopes requested by ZITADEL during the request on the identity provider
	// +listType=set
	Scopes []*string `json:"scopes,omitempty" tf:"scopes,omitempty"`
}

type IdpGitlabParameters struct {

	// (String) client id generated by the identity provider
	// client id generated by the identity provider
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// (String, Sensitive) client secret generated by the identity provider
	// client secret generated by the identity provider
	// +kubebuilder:validation:Optional
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

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

	// (String) Name of the IDP
	// Name of the IDP
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Set of String) the scopes requested by ZITADEL during the request on the identity provider
	// the scopes requested by ZITADEL during the request on the identity provider
	// +kubebuilder:validation:Optional
	// +listType=set
	Scopes []*string `json:"scopes,omitempty" tf:"scopes,omitempty"`
}

// IdpGitlabSpec defines the desired state of IdpGitlab
type IdpGitlabSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IdpGitlabParameters `json:"forProvider"`
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
	InitProvider IdpGitlabInitParameters `json:"initProvider,omitempty"`
}

// IdpGitlabStatus defines the observed state of IdpGitlab.
type IdpGitlabStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IdpGitlabObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// IdpGitlab is the Schema for the IdpGitlabs API. Resource representing a GitLab IDP on the instance.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,zitadel}
type IdpGitlab struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientId) || (has(self.initProvider) && has(self.initProvider.clientId))",message="spec.forProvider.clientId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientSecretSecretRef)",message="spec.forProvider.clientSecretSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.isAutoCreation) || (has(self.initProvider) && has(self.initProvider.isAutoCreation))",message="spec.forProvider.isAutoCreation is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.isAutoUpdate) || (has(self.initProvider) && has(self.initProvider.isAutoUpdate))",message="spec.forProvider.isAutoUpdate is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.isCreationAllowed) || (has(self.initProvider) && has(self.initProvider.isCreationAllowed))",message="spec.forProvider.isCreationAllowed is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.isLinkingAllowed) || (has(self.initProvider) && has(self.initProvider.isLinkingAllowed))",message="spec.forProvider.isLinkingAllowed is a required parameter"
	Spec   IdpGitlabSpec   `json:"spec"`
	Status IdpGitlabStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IdpGitlabList contains a list of IdpGitlabs
type IdpGitlabList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IdpGitlab `json:"items"`
}

// Repository type metadata.
var (
	IdpGitlab_Kind             = "IdpGitlab"
	IdpGitlab_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IdpGitlab_Kind}.String()
	IdpGitlab_KindAPIVersion   = IdpGitlab_Kind + "." + CRDGroupVersion.String()
	IdpGitlab_GroupVersionKind = CRDGroupVersion.WithKind(IdpGitlab_Kind)
)

func init() {
	SchemeBuilder.Register(&IdpGitlab{}, &IdpGitlabList{})
}

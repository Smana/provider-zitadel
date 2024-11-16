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

type DefaultPasswordComplexityPolicyInitParameters struct {

	// (Boolean) defines if the password MUST contain a lower case letter
	// defines if the password MUST contain a lower case letter
	HasLowercase *bool `json:"hasLowercase,omitempty" tf:"has_lowercase,omitempty"`

	// (Boolean) defines if the password MUST contain a number
	// defines if the password MUST contain a number
	HasNumber *bool `json:"hasNumber,omitempty" tf:"has_number,omitempty"`

	// (Boolean) defines if the password MUST contain a symbol. E.g. "$"
	// defines if the password MUST contain a symbol. E.g. "$"
	HasSymbol *bool `json:"hasSymbol,omitempty" tf:"has_symbol,omitempty"`

	// (Boolean) defines if the password MUST contain an upper case letter
	// defines if the password MUST contain an upper case letter
	HasUppercase *bool `json:"hasUppercase,omitempty" tf:"has_uppercase,omitempty"`

	// (Number) Minimal length for the password
	// Minimal length for the password
	MinLength *float64 `json:"minLength,omitempty" tf:"min_length,omitempty"`
}

type DefaultPasswordComplexityPolicyObservation struct {

	// (Boolean) defines if the password MUST contain a lower case letter
	// defines if the password MUST contain a lower case letter
	HasLowercase *bool `json:"hasLowercase,omitempty" tf:"has_lowercase,omitempty"`

	// (Boolean) defines if the password MUST contain a number
	// defines if the password MUST contain a number
	HasNumber *bool `json:"hasNumber,omitempty" tf:"has_number,omitempty"`

	// (Boolean) defines if the password MUST contain a symbol. E.g. "$"
	// defines if the password MUST contain a symbol. E.g. "$"
	HasSymbol *bool `json:"hasSymbol,omitempty" tf:"has_symbol,omitempty"`

	// (Boolean) defines if the password MUST contain an upper case letter
	// defines if the password MUST contain an upper case letter
	HasUppercase *bool `json:"hasUppercase,omitempty" tf:"has_uppercase,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Number) Minimal length for the password
	// Minimal length for the password
	MinLength *float64 `json:"minLength,omitempty" tf:"min_length,omitempty"`
}

type DefaultPasswordComplexityPolicyParameters struct {

	// (Boolean) defines if the password MUST contain a lower case letter
	// defines if the password MUST contain a lower case letter
	// +kubebuilder:validation:Optional
	HasLowercase *bool `json:"hasLowercase,omitempty" tf:"has_lowercase,omitempty"`

	// (Boolean) defines if the password MUST contain a number
	// defines if the password MUST contain a number
	// +kubebuilder:validation:Optional
	HasNumber *bool `json:"hasNumber,omitempty" tf:"has_number,omitempty"`

	// (Boolean) defines if the password MUST contain a symbol. E.g. "$"
	// defines if the password MUST contain a symbol. E.g. "$"
	// +kubebuilder:validation:Optional
	HasSymbol *bool `json:"hasSymbol,omitempty" tf:"has_symbol,omitempty"`

	// (Boolean) defines if the password MUST contain an upper case letter
	// defines if the password MUST contain an upper case letter
	// +kubebuilder:validation:Optional
	HasUppercase *bool `json:"hasUppercase,omitempty" tf:"has_uppercase,omitempty"`

	// (Number) Minimal length for the password
	// Minimal length for the password
	// +kubebuilder:validation:Optional
	MinLength *float64 `json:"minLength,omitempty" tf:"min_length,omitempty"`
}

// DefaultPasswordComplexityPolicySpec defines the desired state of DefaultPasswordComplexityPolicy
type DefaultPasswordComplexityPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DefaultPasswordComplexityPolicyParameters `json:"forProvider"`
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
	InitProvider DefaultPasswordComplexityPolicyInitParameters `json:"initProvider,omitempty"`
}

// DefaultPasswordComplexityPolicyStatus defines the observed state of DefaultPasswordComplexityPolicy.
type DefaultPasswordComplexityPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DefaultPasswordComplexityPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DefaultPasswordComplexityPolicy is the Schema for the DefaultPasswordComplexityPolicys API. Resource representing the default password complexity policy.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,zitadel}
type DefaultPasswordComplexityPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hasLowercase) || (has(self.initProvider) && has(self.initProvider.hasLowercase))",message="spec.forProvider.hasLowercase is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hasNumber) || (has(self.initProvider) && has(self.initProvider.hasNumber))",message="spec.forProvider.hasNumber is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hasSymbol) || (has(self.initProvider) && has(self.initProvider.hasSymbol))",message="spec.forProvider.hasSymbol is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hasUppercase) || (has(self.initProvider) && has(self.initProvider.hasUppercase))",message="spec.forProvider.hasUppercase is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.minLength) || (has(self.initProvider) && has(self.initProvider.minLength))",message="spec.forProvider.minLength is a required parameter"
	Spec   DefaultPasswordComplexityPolicySpec   `json:"spec"`
	Status DefaultPasswordComplexityPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultPasswordComplexityPolicyList contains a list of DefaultPasswordComplexityPolicys
type DefaultPasswordComplexityPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultPasswordComplexityPolicy `json:"items"`
}

// Repository type metadata.
var (
	DefaultPasswordComplexityPolicy_Kind             = "DefaultPasswordComplexityPolicy"
	DefaultPasswordComplexityPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DefaultPasswordComplexityPolicy_Kind}.String()
	DefaultPasswordComplexityPolicy_KindAPIVersion   = DefaultPasswordComplexityPolicy_Kind + "." + CRDGroupVersion.String()
	DefaultPasswordComplexityPolicy_GroupVersionKind = CRDGroupVersion.WithKind(DefaultPasswordComplexityPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&DefaultPasswordComplexityPolicy{}, &DefaultPasswordComplexityPolicyList{})
}

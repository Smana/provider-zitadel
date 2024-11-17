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

type DefaultLockoutPolicyInitParameters struct {

	// (Number) Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correctly or the password is reset.
	// Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correctly or the password is reset.
	MaxPasswordAttempts *float64 `json:"maxPasswordAttempts,omitempty" tf:"max_password_attempts,omitempty"`
}

type DefaultLockoutPolicyObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Number) Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correctly or the password is reset.
	// Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correctly or the password is reset.
	MaxPasswordAttempts *float64 `json:"maxPasswordAttempts,omitempty" tf:"max_password_attempts,omitempty"`
}

type DefaultLockoutPolicyParameters struct {

	// (Number) Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correctly or the password is reset.
	// Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correctly or the password is reset.
	// +kubebuilder:validation:Optional
	MaxPasswordAttempts *float64 `json:"maxPasswordAttempts,omitempty" tf:"max_password_attempts,omitempty"`
}

// DefaultLockoutPolicySpec defines the desired state of DefaultLockoutPolicy
type DefaultLockoutPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DefaultLockoutPolicyParameters `json:"forProvider"`
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
	InitProvider DefaultLockoutPolicyInitParameters `json:"initProvider,omitempty"`
}

// DefaultLockoutPolicyStatus defines the observed state of DefaultLockoutPolicy.
type DefaultLockoutPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DefaultLockoutPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DefaultLockoutPolicy is the Schema for the DefaultLockoutPolicys API. Resource representing the default lockout policy.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,zitadel}
type DefaultLockoutPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.maxPasswordAttempts) || (has(self.initProvider) && has(self.initProvider.maxPasswordAttempts))",message="spec.forProvider.maxPasswordAttempts is a required parameter"
	Spec   DefaultLockoutPolicySpec   `json:"spec"`
	Status DefaultLockoutPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultLockoutPolicyList contains a list of DefaultLockoutPolicys
type DefaultLockoutPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultLockoutPolicy `json:"items"`
}

// Repository type metadata.
var (
	DefaultLockoutPolicy_Kind             = "DefaultLockoutPolicy"
	DefaultLockoutPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DefaultLockoutPolicy_Kind}.String()
	DefaultLockoutPolicy_KindAPIVersion   = DefaultLockoutPolicy_Kind + "." + CRDGroupVersion.String()
	DefaultLockoutPolicy_GroupVersionKind = CRDGroupVersion.WithKind(DefaultLockoutPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&DefaultLockoutPolicy{}, &DefaultLockoutPolicyList{})
}

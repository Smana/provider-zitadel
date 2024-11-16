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

type DefaultDomainPolicyInitParameters struct {

	// (Boolean)
	SMTPSenderAddressMatchesInstanceDomain *bool `json:"smtpSenderAddressMatchesInstanceDomain,omitempty" tf:"smtp_sender_address_matches_instance_domain,omitempty"`

	// (Boolean) User login must be domain
	// User login must be domain
	UserLoginMustBeDomain *bool `json:"userLoginMustBeDomain,omitempty" tf:"user_login_must_be_domain,omitempty"`

	// (Boolean) Validate organization domains
	// Validate organization domains
	ValidateOrgDomains *bool `json:"validateOrgDomains,omitempty" tf:"validate_org_domains,omitempty"`
}

type DefaultDomainPolicyObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Boolean)
	SMTPSenderAddressMatchesInstanceDomain *bool `json:"smtpSenderAddressMatchesInstanceDomain,omitempty" tf:"smtp_sender_address_matches_instance_domain,omitempty"`

	// (Boolean) User login must be domain
	// User login must be domain
	UserLoginMustBeDomain *bool `json:"userLoginMustBeDomain,omitempty" tf:"user_login_must_be_domain,omitempty"`

	// (Boolean) Validate organization domains
	// Validate organization domains
	ValidateOrgDomains *bool `json:"validateOrgDomains,omitempty" tf:"validate_org_domains,omitempty"`
}

type DefaultDomainPolicyParameters struct {

	// (Boolean)
	// +kubebuilder:validation:Optional
	SMTPSenderAddressMatchesInstanceDomain *bool `json:"smtpSenderAddressMatchesInstanceDomain,omitempty" tf:"smtp_sender_address_matches_instance_domain,omitempty"`

	// (Boolean) User login must be domain
	// User login must be domain
	// +kubebuilder:validation:Optional
	UserLoginMustBeDomain *bool `json:"userLoginMustBeDomain,omitempty" tf:"user_login_must_be_domain,omitempty"`

	// (Boolean) Validate organization domains
	// Validate organization domains
	// +kubebuilder:validation:Optional
	ValidateOrgDomains *bool `json:"validateOrgDomains,omitempty" tf:"validate_org_domains,omitempty"`
}

// DefaultDomainPolicySpec defines the desired state of DefaultDomainPolicy
type DefaultDomainPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DefaultDomainPolicyParameters `json:"forProvider"`
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
	InitProvider DefaultDomainPolicyInitParameters `json:"initProvider,omitempty"`
}

// DefaultDomainPolicyStatus defines the observed state of DefaultDomainPolicy.
type DefaultDomainPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DefaultDomainPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DefaultDomainPolicy is the Schema for the DefaultDomainPolicys API. Resource representing the default domain policy.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,zitadel}
type DefaultDomainPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.smtpSenderAddressMatchesInstanceDomain) || (has(self.initProvider) && has(self.initProvider.smtpSenderAddressMatchesInstanceDomain))",message="spec.forProvider.smtpSenderAddressMatchesInstanceDomain is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.userLoginMustBeDomain) || (has(self.initProvider) && has(self.initProvider.userLoginMustBeDomain))",message="spec.forProvider.userLoginMustBeDomain is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.validateOrgDomains) || (has(self.initProvider) && has(self.initProvider.validateOrgDomains))",message="spec.forProvider.validateOrgDomains is a required parameter"
	Spec   DefaultDomainPolicySpec   `json:"spec"`
	Status DefaultDomainPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultDomainPolicyList contains a list of DefaultDomainPolicys
type DefaultDomainPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultDomainPolicy `json:"items"`
}

// Repository type metadata.
var (
	DefaultDomainPolicy_Kind             = "DefaultDomainPolicy"
	DefaultDomainPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DefaultDomainPolicy_Kind}.String()
	DefaultDomainPolicy_KindAPIVersion   = DefaultDomainPolicy_Kind + "." + CRDGroupVersion.String()
	DefaultDomainPolicy_GroupVersionKind = CRDGroupVersion.WithKind(DefaultDomainPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&DefaultDomainPolicy{}, &DefaultDomainPolicyList{})
}

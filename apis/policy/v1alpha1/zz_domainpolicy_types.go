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

type DomainPolicyInitParameters struct {

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

	// (Boolean)
	SMTPSenderAddressMatchesInstanceDomain *bool `json:"smtpSenderAddressMatchesInstanceDomain,omitempty" tf:"smtp_sender_address_matches_instance_domain,omitempty"`

	// (Boolean) User login must be domain
	// User login must be domain
	UserLoginMustBeDomain *bool `json:"userLoginMustBeDomain,omitempty" tf:"user_login_must_be_domain,omitempty"`

	// (Boolean) Validate organization domains
	// Validate organization domains
	ValidateOrgDomains *bool `json:"validateOrgDomains,omitempty" tf:"validate_org_domains,omitempty"`
}

type DomainPolicyObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) ID of the organization
	// ID of the organization
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// (Boolean)
	SMTPSenderAddressMatchesInstanceDomain *bool `json:"smtpSenderAddressMatchesInstanceDomain,omitempty" tf:"smtp_sender_address_matches_instance_domain,omitempty"`

	// (Boolean) User login must be domain
	// User login must be domain
	UserLoginMustBeDomain *bool `json:"userLoginMustBeDomain,omitempty" tf:"user_login_must_be_domain,omitempty"`

	// (Boolean) Validate organization domains
	// Validate organization domains
	ValidateOrgDomains *bool `json:"validateOrgDomains,omitempty" tf:"validate_org_domains,omitempty"`
}

type DomainPolicyParameters struct {

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

// DomainPolicySpec defines the desired state of DomainPolicy
type DomainPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainPolicyParameters `json:"forProvider"`
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
	InitProvider DomainPolicyInitParameters `json:"initProvider,omitempty"`
}

// DomainPolicyStatus defines the observed state of DomainPolicy.
type DomainPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DomainPolicy is the Schema for the DomainPolicys API. Resource representing the custom domain policy of an organization.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,zitadel}
type DomainPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.smtpSenderAddressMatchesInstanceDomain) || (has(self.initProvider) && has(self.initProvider.smtpSenderAddressMatchesInstanceDomain))",message="spec.forProvider.smtpSenderAddressMatchesInstanceDomain is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.userLoginMustBeDomain) || (has(self.initProvider) && has(self.initProvider.userLoginMustBeDomain))",message="spec.forProvider.userLoginMustBeDomain is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.validateOrgDomains) || (has(self.initProvider) && has(self.initProvider.validateOrgDomains))",message="spec.forProvider.validateOrgDomains is a required parameter"
	Spec   DomainPolicySpec   `json:"spec"`
	Status DomainPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainPolicyList contains a list of DomainPolicys
type DomainPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainPolicy `json:"items"`
}

// Repository type metadata.
var (
	DomainPolicy_Kind             = "DomainPolicy"
	DomainPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DomainPolicy_Kind}.String()
	DomainPolicy_KindAPIVersion   = DomainPolicy_Kind + "." + CRDGroupVersion.String()
	DomainPolicy_GroupVersionKind = CRDGroupVersion.WithKind(DomainPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&DomainPolicy{}, &DomainPolicyList{})
}

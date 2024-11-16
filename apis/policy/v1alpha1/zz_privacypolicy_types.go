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

type PrivacyPolicyInitParameters struct {

	// (String)
	HelpLink *string `json:"helpLink,omitempty" tf:"help_link,omitempty"`

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

	// (String)
	PrivacyLink *string `json:"privacyLink,omitempty" tf:"privacy_link,omitempty"`

	// (String)
	SupportEmail *string `json:"supportEmail,omitempty" tf:"support_email,omitempty"`

	// (String)
	TosLink *string `json:"tosLink,omitempty" tf:"tos_link,omitempty"`
}

type PrivacyPolicyObservation struct {

	// (String)
	HelpLink *string `json:"helpLink,omitempty" tf:"help_link,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) ID of the organization
	// ID of the organization
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// (String)
	PrivacyLink *string `json:"privacyLink,omitempty" tf:"privacy_link,omitempty"`

	// (String)
	SupportEmail *string `json:"supportEmail,omitempty" tf:"support_email,omitempty"`

	// (String)
	TosLink *string `json:"tosLink,omitempty" tf:"tos_link,omitempty"`
}

type PrivacyPolicyParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	HelpLink *string `json:"helpLink,omitempty" tf:"help_link,omitempty"`

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

	// (String)
	// +kubebuilder:validation:Optional
	PrivacyLink *string `json:"privacyLink,omitempty" tf:"privacy_link,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	SupportEmail *string `json:"supportEmail,omitempty" tf:"support_email,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	TosLink *string `json:"tosLink,omitempty" tf:"tos_link,omitempty"`
}

// PrivacyPolicySpec defines the desired state of PrivacyPolicy
type PrivacyPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivacyPolicyParameters `json:"forProvider"`
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
	InitProvider PrivacyPolicyInitParameters `json:"initProvider,omitempty"`
}

// PrivacyPolicyStatus defines the observed state of PrivacyPolicy.
type PrivacyPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivacyPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PrivacyPolicy is the Schema for the PrivacyPolicys API. Resource representing the custom privacy policy of an organization.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,zitadel}
type PrivacyPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivacyPolicySpec   `json:"spec"`
	Status            PrivacyPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivacyPolicyList contains a list of PrivacyPolicys
type PrivacyPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivacyPolicy `json:"items"`
}

// Repository type metadata.
var (
	PrivacyPolicy_Kind             = "PrivacyPolicy"
	PrivacyPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivacyPolicy_Kind}.String()
	PrivacyPolicy_KindAPIVersion   = PrivacyPolicy_Kind + "." + CRDGroupVersion.String()
	PrivacyPolicy_GroupVersionKind = CRDGroupVersion.WithKind(PrivacyPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivacyPolicy{}, &PrivacyPolicyList{})
}

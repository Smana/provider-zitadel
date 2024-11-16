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

type VerifySMSOtpMessageTextInitParameters struct {

	// (String)
	ButtonText *string `json:"buttonText,omitempty" tf:"button_text,omitempty"`

	// (String)
	FooterText *string `json:"footerText,omitempty" tf:"footer_text,omitempty"`

	// (String)
	Greeting *string `json:"greeting,omitempty" tf:"greeting,omitempty"`

	// (String)
	Language *string `json:"language,omitempty" tf:"language,omitempty"`

	// (String)
	PreHeader *string `json:"preHeader,omitempty" tf:"pre_header,omitempty"`

	// (String)
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// (String)
	Text *string `json:"text,omitempty" tf:"text,omitempty"`

	// (String)
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type VerifySMSOtpMessageTextObservation struct {

	// (String)
	ButtonText *string `json:"buttonText,omitempty" tf:"button_text,omitempty"`

	// (String)
	FooterText *string `json:"footerText,omitempty" tf:"footer_text,omitempty"`

	// (String)
	Greeting *string `json:"greeting,omitempty" tf:"greeting,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String)
	Language *string `json:"language,omitempty" tf:"language,omitempty"`

	// (String)
	PreHeader *string `json:"preHeader,omitempty" tf:"pre_header,omitempty"`

	// (String)
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// (String)
	Text *string `json:"text,omitempty" tf:"text,omitempty"`

	// (String)
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type VerifySMSOtpMessageTextParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	ButtonText *string `json:"buttonText,omitempty" tf:"button_text,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	FooterText *string `json:"footerText,omitempty" tf:"footer_text,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Greeting *string `json:"greeting,omitempty" tf:"greeting,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Language *string `json:"language,omitempty" tf:"language,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	PreHeader *string `json:"preHeader,omitempty" tf:"pre_header,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Text *string `json:"text,omitempty" tf:"text,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

// VerifySMSOtpMessageTextSpec defines the desired state of VerifySMSOtpMessageText
type VerifySMSOtpMessageTextSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VerifySMSOtpMessageTextParameters `json:"forProvider"`
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
	InitProvider VerifySMSOtpMessageTextInitParameters `json:"initProvider,omitempty"`
}

// VerifySMSOtpMessageTextStatus defines the observed state of VerifySMSOtpMessageText.
type VerifySMSOtpMessageTextStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VerifySMSOtpMessageTextObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VerifySMSOtpMessageText is the Schema for the VerifySMSOtpMessageTexts API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,zitadel}
type VerifySMSOtpMessageText struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.language) || (has(self.initProvider) && has(self.initProvider.language))",message="spec.forProvider.language is a required parameter"
	Spec   VerifySMSOtpMessageTextSpec   `json:"spec"`
	Status VerifySMSOtpMessageTextStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VerifySMSOtpMessageTextList contains a list of VerifySMSOtpMessageTexts
type VerifySMSOtpMessageTextList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VerifySMSOtpMessageText `json:"items"`
}

// Repository type metadata.
var (
	VerifySMSOtpMessageText_Kind             = "VerifySMSOtpMessageText"
	VerifySMSOtpMessageText_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VerifySMSOtpMessageText_Kind}.String()
	VerifySMSOtpMessageText_KindAPIVersion   = VerifySMSOtpMessageText_Kind + "." + CRDGroupVersion.String()
	VerifySMSOtpMessageText_GroupVersionKind = CRDGroupVersion.WithKind(VerifySMSOtpMessageText_Kind)
)

func init() {
	SchemeBuilder.Register(&VerifySMSOtpMessageText{}, &VerifySMSOtpMessageTextList{})
}

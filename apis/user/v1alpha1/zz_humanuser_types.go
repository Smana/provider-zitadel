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

type HumanUserInitParameters struct {

	// (String) Display name of the user
	// Display name of the user
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (String) Email of the user
	// Email of the user
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// (String) First name of the user
	// First name of the user
	FirstName *string `json:"firstName,omitempty" tf:"first_name,omitempty"`

	// (String) Gender of the user, supported values: GENDER_UNSPECIFIED, GENDER_FEMALE, GENDER_MALE, GENDER_DIVERSE
	// Gender of the user, supported values: GENDER_UNSPECIFIED, GENDER_FEMALE, GENDER_MALE, GENDER_DIVERSE
	Gender *string `json:"gender,omitempty" tf:"gender,omitempty"`

	// (String, Sensitive) Initially set password for the user, not changeable after creation
	// Initially set password for the user, not changeable after creation
	InitialPasswordSecretRef *v1.SecretKeySelector `json:"initialPasswordSecretRef,omitempty" tf:"-"`

	// (Boolean) Is the email verified of the user, can only be true if password of the user is set
	// Is the email verified of the user, can only be true if password of the user is set
	IsEmailVerified *bool `json:"isEmailVerified,omitempty" tf:"is_email_verified,omitempty"`

	// (Boolean) Is the phone verified of the user
	// Is the phone verified of the user
	IsPhoneVerified *bool `json:"isPhoneVerified,omitempty" tf:"is_phone_verified,omitempty"`

	// (String) Last name of the user
	// Last name of the user
	LastName *string `json:"lastName,omitempty" tf:"last_name,omitempty"`

	// (String) Nick name of the user
	// Nick name of the user
	NickName *string `json:"nickName,omitempty" tf:"nick_name,omitempty"`

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

	// (String) Phone of the user
	// Phone of the user
	Phone *string `json:"phone,omitempty" tf:"phone,omitempty"`

	// (String) Preferred language of the user
	// Preferred language of the user
	PreferredLanguage *string `json:"preferredLanguage,omitempty" tf:"preferred_language,omitempty"`

	// (String) Username
	// Username
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type HumanUserObservation struct {

	// (String) Display name of the user
	// Display name of the user
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (String) Email of the user
	// Email of the user
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// (String) First name of the user
	// First name of the user
	FirstName *string `json:"firstName,omitempty" tf:"first_name,omitempty"`

	// (String) Gender of the user, supported values: GENDER_UNSPECIFIED, GENDER_FEMALE, GENDER_MALE, GENDER_DIVERSE
	// Gender of the user, supported values: GENDER_UNSPECIFIED, GENDER_FEMALE, GENDER_MALE, GENDER_DIVERSE
	Gender *string `json:"gender,omitempty" tf:"gender,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Boolean) Is the email verified of the user, can only be true if password of the user is set
	// Is the email verified of the user, can only be true if password of the user is set
	IsEmailVerified *bool `json:"isEmailVerified,omitempty" tf:"is_email_verified,omitempty"`

	// (Boolean) Is the phone verified of the user
	// Is the phone verified of the user
	IsPhoneVerified *bool `json:"isPhoneVerified,omitempty" tf:"is_phone_verified,omitempty"`

	// (String) Last name of the user
	// Last name of the user
	LastName *string `json:"lastName,omitempty" tf:"last_name,omitempty"`

	// (List of String) Loginnames
	// Loginnames
	LoginNames []*string `json:"loginNames,omitempty" tf:"login_names,omitempty"`

	// (String) Nick name of the user
	// Nick name of the user
	NickName *string `json:"nickName,omitempty" tf:"nick_name,omitempty"`

	// (String) ID of the organization
	// ID of the organization
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// (String) Phone of the user
	// Phone of the user
	Phone *string `json:"phone,omitempty" tf:"phone,omitempty"`

	// (String) Preferred language of the user
	// Preferred language of the user
	PreferredLanguage *string `json:"preferredLanguage,omitempty" tf:"preferred_language,omitempty"`

	// (String) Preferred login name
	// Preferred login name
	PreferredLoginName *string `json:"preferredLoginName,omitempty" tf:"preferred_login_name,omitempty"`

	// (String) State of the user
	// State of the user
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// (String) Username
	// Username
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type HumanUserParameters struct {

	// (String) Display name of the user
	// Display name of the user
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (String) Email of the user
	// Email of the user
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// (String) First name of the user
	// First name of the user
	// +kubebuilder:validation:Optional
	FirstName *string `json:"firstName,omitempty" tf:"first_name,omitempty"`

	// (String) Gender of the user, supported values: GENDER_UNSPECIFIED, GENDER_FEMALE, GENDER_MALE, GENDER_DIVERSE
	// Gender of the user, supported values: GENDER_UNSPECIFIED, GENDER_FEMALE, GENDER_MALE, GENDER_DIVERSE
	// +kubebuilder:validation:Optional
	Gender *string `json:"gender,omitempty" tf:"gender,omitempty"`

	// (String, Sensitive) Initially set password for the user, not changeable after creation
	// Initially set password for the user, not changeable after creation
	// +kubebuilder:validation:Optional
	InitialPasswordSecretRef *v1.SecretKeySelector `json:"initialPasswordSecretRef,omitempty" tf:"-"`

	// (Boolean) Is the email verified of the user, can only be true if password of the user is set
	// Is the email verified of the user, can only be true if password of the user is set
	// +kubebuilder:validation:Optional
	IsEmailVerified *bool `json:"isEmailVerified,omitempty" tf:"is_email_verified,omitempty"`

	// (Boolean) Is the phone verified of the user
	// Is the phone verified of the user
	// +kubebuilder:validation:Optional
	IsPhoneVerified *bool `json:"isPhoneVerified,omitempty" tf:"is_phone_verified,omitempty"`

	// (String) Last name of the user
	// Last name of the user
	// +kubebuilder:validation:Optional
	LastName *string `json:"lastName,omitempty" tf:"last_name,omitempty"`

	// (String) Nick name of the user
	// Nick name of the user
	// +kubebuilder:validation:Optional
	NickName *string `json:"nickName,omitempty" tf:"nick_name,omitempty"`

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

	// (String) Phone of the user
	// Phone of the user
	// +kubebuilder:validation:Optional
	Phone *string `json:"phone,omitempty" tf:"phone,omitempty"`

	// (String) Preferred language of the user
	// Preferred language of the user
	// +kubebuilder:validation:Optional
	PreferredLanguage *string `json:"preferredLanguage,omitempty" tf:"preferred_language,omitempty"`

	// (String) Username
	// Username
	// +kubebuilder:validation:Optional
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

// HumanUserSpec defines the desired state of HumanUser
type HumanUserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HumanUserParameters `json:"forProvider"`
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
	InitProvider HumanUserInitParameters `json:"initProvider,omitempty"`
}

// HumanUserStatus defines the observed state of HumanUser.
type HumanUserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HumanUserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// HumanUser is the Schema for the HumanUsers API. Resource representing a human user situated under an organization, which then can be authorized through memberships or direct grants on other resources.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,zitadel}
type HumanUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.email) || (has(self.initProvider) && has(self.initProvider.email))",message="spec.forProvider.email is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.firstName) || (has(self.initProvider) && has(self.initProvider.firstName))",message="spec.forProvider.firstName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.lastName) || (has(self.initProvider) && has(self.initProvider.lastName))",message="spec.forProvider.lastName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.userName) || (has(self.initProvider) && has(self.initProvider.userName))",message="spec.forProvider.userName is a required parameter"
	Spec   HumanUserSpec   `json:"spec"`
	Status HumanUserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HumanUserList contains a list of HumanUsers
type HumanUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HumanUser `json:"items"`
}

// Repository type metadata.
var (
	HumanUser_Kind             = "HumanUser"
	HumanUser_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HumanUser_Kind}.String()
	HumanUser_KindAPIVersion   = HumanUser_Kind + "." + CRDGroupVersion.String()
	HumanUser_GroupVersionKind = CRDGroupVersion.WithKind(HumanUser_Kind)
)

func init() {
	SchemeBuilder.Register(&HumanUser{}, &HumanUserList{})
}

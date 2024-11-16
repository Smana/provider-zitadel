//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResetMessageText) DeepCopyInto(out *ResetMessageText) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResetMessageText.
func (in *ResetMessageText) DeepCopy() *ResetMessageText {
	if in == nil {
		return nil
	}
	out := new(ResetMessageText)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResetMessageText) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResetMessageTextInitParameters) DeepCopyInto(out *ResetMessageTextInitParameters) {
	*out = *in
	if in.ButtonText != nil {
		in, out := &in.ButtonText, &out.ButtonText
		*out = new(string)
		**out = **in
	}
	if in.FooterText != nil {
		in, out := &in.FooterText, &out.FooterText
		*out = new(string)
		**out = **in
	}
	if in.Greeting != nil {
		in, out := &in.Greeting, &out.Greeting
		*out = new(string)
		**out = **in
	}
	if in.Language != nil {
		in, out := &in.Language, &out.Language
		*out = new(string)
		**out = **in
	}
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(string)
		**out = **in
	}
	if in.PreHeader != nil {
		in, out := &in.PreHeader, &out.PreHeader
		*out = new(string)
		**out = **in
	}
	if in.Subject != nil {
		in, out := &in.Subject, &out.Subject
		*out = new(string)
		**out = **in
	}
	if in.Text != nil {
		in, out := &in.Text, &out.Text
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResetMessageTextInitParameters.
func (in *ResetMessageTextInitParameters) DeepCopy() *ResetMessageTextInitParameters {
	if in == nil {
		return nil
	}
	out := new(ResetMessageTextInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResetMessageTextList) DeepCopyInto(out *ResetMessageTextList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResetMessageText, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResetMessageTextList.
func (in *ResetMessageTextList) DeepCopy() *ResetMessageTextList {
	if in == nil {
		return nil
	}
	out := new(ResetMessageTextList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResetMessageTextList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResetMessageTextObservation) DeepCopyInto(out *ResetMessageTextObservation) {
	*out = *in
	if in.ButtonText != nil {
		in, out := &in.ButtonText, &out.ButtonText
		*out = new(string)
		**out = **in
	}
	if in.FooterText != nil {
		in, out := &in.FooterText, &out.FooterText
		*out = new(string)
		**out = **in
	}
	if in.Greeting != nil {
		in, out := &in.Greeting, &out.Greeting
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Language != nil {
		in, out := &in.Language, &out.Language
		*out = new(string)
		**out = **in
	}
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(string)
		**out = **in
	}
	if in.PreHeader != nil {
		in, out := &in.PreHeader, &out.PreHeader
		*out = new(string)
		**out = **in
	}
	if in.Subject != nil {
		in, out := &in.Subject, &out.Subject
		*out = new(string)
		**out = **in
	}
	if in.Text != nil {
		in, out := &in.Text, &out.Text
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResetMessageTextObservation.
func (in *ResetMessageTextObservation) DeepCopy() *ResetMessageTextObservation {
	if in == nil {
		return nil
	}
	out := new(ResetMessageTextObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResetMessageTextParameters) DeepCopyInto(out *ResetMessageTextParameters) {
	*out = *in
	if in.ButtonText != nil {
		in, out := &in.ButtonText, &out.ButtonText
		*out = new(string)
		**out = **in
	}
	if in.FooterText != nil {
		in, out := &in.FooterText, &out.FooterText
		*out = new(string)
		**out = **in
	}
	if in.Greeting != nil {
		in, out := &in.Greeting, &out.Greeting
		*out = new(string)
		**out = **in
	}
	if in.Language != nil {
		in, out := &in.Language, &out.Language
		*out = new(string)
		**out = **in
	}
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(string)
		**out = **in
	}
	if in.PreHeader != nil {
		in, out := &in.PreHeader, &out.PreHeader
		*out = new(string)
		**out = **in
	}
	if in.Subject != nil {
		in, out := &in.Subject, &out.Subject
		*out = new(string)
		**out = **in
	}
	if in.Text != nil {
		in, out := &in.Text, &out.Text
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResetMessageTextParameters.
func (in *ResetMessageTextParameters) DeepCopy() *ResetMessageTextParameters {
	if in == nil {
		return nil
	}
	out := new(ResetMessageTextParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResetMessageTextSpec) DeepCopyInto(out *ResetMessageTextSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResetMessageTextSpec.
func (in *ResetMessageTextSpec) DeepCopy() *ResetMessageTextSpec {
	if in == nil {
		return nil
	}
	out := new(ResetMessageTextSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResetMessageTextStatus) DeepCopyInto(out *ResetMessageTextStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResetMessageTextStatus.
func (in *ResetMessageTextStatus) DeepCopy() *ResetMessageTextStatus {
	if in == nil {
		return nil
	}
	out := new(ResetMessageTextStatus)
	in.DeepCopyInto(out)
	return out
}

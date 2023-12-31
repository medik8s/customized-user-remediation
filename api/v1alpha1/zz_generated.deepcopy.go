//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizedUserRemediation) DeepCopyInto(out *CustomizedUserRemediation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizedUserRemediation.
func (in *CustomizedUserRemediation) DeepCopy() *CustomizedUserRemediation {
	if in == nil {
		return nil
	}
	out := new(CustomizedUserRemediation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomizedUserRemediation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizedUserRemediationConfig) DeepCopyInto(out *CustomizedUserRemediationConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizedUserRemediationConfig.
func (in *CustomizedUserRemediationConfig) DeepCopy() *CustomizedUserRemediationConfig {
	if in == nil {
		return nil
	}
	out := new(CustomizedUserRemediationConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomizedUserRemediationConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizedUserRemediationConfigList) DeepCopyInto(out *CustomizedUserRemediationConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CustomizedUserRemediationConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizedUserRemediationConfigList.
func (in *CustomizedUserRemediationConfigList) DeepCopy() *CustomizedUserRemediationConfigList {
	if in == nil {
		return nil
	}
	out := new(CustomizedUserRemediationConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomizedUserRemediationConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizedUserRemediationConfigSpec) DeepCopyInto(out *CustomizedUserRemediationConfigSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizedUserRemediationConfigSpec.
func (in *CustomizedUserRemediationConfigSpec) DeepCopy() *CustomizedUserRemediationConfigSpec {
	if in == nil {
		return nil
	}
	out := new(CustomizedUserRemediationConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizedUserRemediationConfigStatus) DeepCopyInto(out *CustomizedUserRemediationConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizedUserRemediationConfigStatus.
func (in *CustomizedUserRemediationConfigStatus) DeepCopy() *CustomizedUserRemediationConfigStatus {
	if in == nil {
		return nil
	}
	out := new(CustomizedUserRemediationConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizedUserRemediationList) DeepCopyInto(out *CustomizedUserRemediationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CustomizedUserRemediation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizedUserRemediationList.
func (in *CustomizedUserRemediationList) DeepCopy() *CustomizedUserRemediationList {
	if in == nil {
		return nil
	}
	out := new(CustomizedUserRemediationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomizedUserRemediationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizedUserRemediationSpec) DeepCopyInto(out *CustomizedUserRemediationSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizedUserRemediationSpec.
func (in *CustomizedUserRemediationSpec) DeepCopy() *CustomizedUserRemediationSpec {
	if in == nil {
		return nil
	}
	out := new(CustomizedUserRemediationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizedUserRemediationStatus) DeepCopyInto(out *CustomizedUserRemediationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizedUserRemediationStatus.
func (in *CustomizedUserRemediationStatus) DeepCopy() *CustomizedUserRemediationStatus {
	if in == nil {
		return nil
	}
	out := new(CustomizedUserRemediationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizedUserRemediationTemplate) DeepCopyInto(out *CustomizedUserRemediationTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizedUserRemediationTemplate.
func (in *CustomizedUserRemediationTemplate) DeepCopy() *CustomizedUserRemediationTemplate {
	if in == nil {
		return nil
	}
	out := new(CustomizedUserRemediationTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomizedUserRemediationTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizedUserRemediationTemplateList) DeepCopyInto(out *CustomizedUserRemediationTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CustomizedUserRemediationTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizedUserRemediationTemplateList.
func (in *CustomizedUserRemediationTemplateList) DeepCopy() *CustomizedUserRemediationTemplateList {
	if in == nil {
		return nil
	}
	out := new(CustomizedUserRemediationTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomizedUserRemediationTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizedUserRemediationTemplateResource) DeepCopyInto(out *CustomizedUserRemediationTemplateResource) {
	*out = *in
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizedUserRemediationTemplateResource.
func (in *CustomizedUserRemediationTemplateResource) DeepCopy() *CustomizedUserRemediationTemplateResource {
	if in == nil {
		return nil
	}
	out := new(CustomizedUserRemediationTemplateResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizedUserRemediationTemplateSpec) DeepCopyInto(out *CustomizedUserRemediationTemplateSpec) {
	*out = *in
	out.Template = in.Template
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizedUserRemediationTemplateSpec.
func (in *CustomizedUserRemediationTemplateSpec) DeepCopy() *CustomizedUserRemediationTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(CustomizedUserRemediationTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizedUserRemediationTemplateStatus) DeepCopyInto(out *CustomizedUserRemediationTemplateStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizedUserRemediationTemplateStatus.
func (in *CustomizedUserRemediationTemplateStatus) DeepCopy() *CustomizedUserRemediationTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(CustomizedUserRemediationTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

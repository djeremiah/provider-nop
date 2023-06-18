//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NopResource) DeepCopyInto(out *NopResource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NopResource.
func (in *NopResource) DeepCopy() *NopResource {
	if in == nil {
		return nil
	}
	out := new(NopResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NopResource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NopResourceList) DeepCopyInto(out *NopResourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NopResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NopResourceList.
func (in *NopResourceList) DeepCopy() *NopResourceList {
	if in == nil {
		return nil
	}
	out := new(NopResourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NopResourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NopResourceObservation) DeepCopyInto(out *NopResourceObservation) {
	*out = *in
	in.Fields.DeepCopyInto(&out.Fields)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NopResourceObservation.
func (in *NopResourceObservation) DeepCopy() *NopResourceObservation {
	if in == nil {
		return nil
	}
	out := new(NopResourceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NopResourceParameters) DeepCopyInto(out *NopResourceParameters) {
	*out = *in
	if in.ConditionAfter != nil {
		in, out := &in.ConditionAfter, &out.ConditionAfter
		*out = make([]ResourceConditionAfter, len(*in))
		copy(*out, *in)
	}
	if in.ConnectionDetails != nil {
		in, out := &in.ConnectionDetails, &out.ConnectionDetails
		*out = make([]ResourceConnectionDetail, len(*in))
		copy(*out, *in)
	}
	in.Fields.DeepCopyInto(&out.Fields)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NopResourceParameters.
func (in *NopResourceParameters) DeepCopy() *NopResourceParameters {
	if in == nil {
		return nil
	}
	out := new(NopResourceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NopResourceSpec) DeepCopyInto(out *NopResourceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NopResourceSpec.
func (in *NopResourceSpec) DeepCopy() *NopResourceSpec {
	if in == nil {
		return nil
	}
	out := new(NopResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NopResourceStatus) DeepCopyInto(out *NopResourceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NopResourceStatus.
func (in *NopResourceStatus) DeepCopy() *NopResourceStatus {
	if in == nil {
		return nil
	}
	out := new(NopResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceConditionAfter) DeepCopyInto(out *ResourceConditionAfter) {
	*out = *in
	out.Time = in.Time
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceConditionAfter.
func (in *ResourceConditionAfter) DeepCopy() *ResourceConditionAfter {
	if in == nil {
		return nil
	}
	out := new(ResourceConditionAfter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceConnectionDetail) DeepCopyInto(out *ResourceConnectionDetail) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceConnectionDetail.
func (in *ResourceConnectionDetail) DeepCopy() *ResourceConnectionDetail {
	if in == nil {
		return nil
	}
	out := new(ResourceConnectionDetail)
	in.DeepCopyInto(out)
	return out
}

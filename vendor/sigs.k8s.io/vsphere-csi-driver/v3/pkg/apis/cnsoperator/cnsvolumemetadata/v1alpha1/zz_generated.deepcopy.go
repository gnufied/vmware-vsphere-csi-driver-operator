// build : ignore_autogenerated

/*
Copyright 2019 The Kubernetes Authors.

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

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnsVolumeMetadata) DeepCopyInto(out *CnsVolumeMetadata) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnsVolumeMetadata.
func (in *CnsVolumeMetadata) DeepCopy() *CnsVolumeMetadata {
	if in == nil {
		return nil
	}
	out := new(CnsVolumeMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CnsVolumeMetadata) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnsVolumeMetadataList) DeepCopyInto(out *CnsVolumeMetadataList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CnsVolumeMetadata, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnsVolumeMetadataList.
func (in *CnsVolumeMetadataList) DeepCopy() *CnsVolumeMetadataList {
	if in == nil {
		return nil
	}
	out := new(CnsVolumeMetadataList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CnsVolumeMetadataList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnsVolumeMetadataSpec) DeepCopyInto(out *CnsVolumeMetadataSpec) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnsVolumeMetadataSpec.
func (in *CnsVolumeMetadataSpec) DeepCopy() *CnsVolumeMetadataSpec {
	if in == nil {
		return nil
	}
	out := new(CnsVolumeMetadataSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnsVolumeMetadataStatus) DeepCopyInto(out *CnsVolumeMetadataStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnsVolumeMetadataStatus.
func (in *CnsVolumeMetadataStatus) DeepCopy() *CnsVolumeMetadataStatus {
	if in == nil {
		return nil
	}
	out := new(CnsVolumeMetadataStatus)
	in.DeepCopyInto(out)
	return out
}

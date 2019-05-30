// +build !ignore_autogenerated

/*
Copyright 2017 The Portworx Operator Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta2

import (
	v1beta1 "github.com/portworx/talisman/pkg/apis/portworx/v1beta1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonPlacementSpec) DeepCopyInto(out *CommonPlacementSpec) {
	*out = *in
	if in.MatchExpressions != nil {
		in, out := &in.MatchExpressions, &out.MatchExpressions
		*out = make([]*v1beta1.LabelSelectorRequirement, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(v1beta1.LabelSelectorRequirement)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonPlacementSpec.
func (in *CommonPlacementSpec) DeepCopy() *CommonPlacementSpec {
	if in == nil {
		return nil
	}
	out := new(CommonPlacementSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaPlacementSpec) DeepCopyInto(out *ReplicaPlacementSpec) {
	*out = *in
	in.CommonPlacementSpec.DeepCopyInto(&out.CommonPlacementSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaPlacementSpec.
func (in *ReplicaPlacementSpec) DeepCopy() *ReplicaPlacementSpec {
	if in == nil {
		return nil
	}
	out := new(ReplicaPlacementSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumePlacementSpec) DeepCopyInto(out *VolumePlacementSpec) {
	*out = *in
	if in.ReplicaAffinity != nil {
		in, out := &in.ReplicaAffinity, &out.ReplicaAffinity
		*out = make([]*ReplicaPlacementSpec, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ReplicaPlacementSpec)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ReplicaAntiAffinity != nil {
		in, out := &in.ReplicaAntiAffinity, &out.ReplicaAntiAffinity
		*out = make([]*ReplicaPlacementSpec, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ReplicaPlacementSpec)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.VolumeAffinity != nil {
		in, out := &in.VolumeAffinity, &out.VolumeAffinity
		*out = make([]*CommonPlacementSpec, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(CommonPlacementSpec)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.VolumeAntiAffinity != nil {
		in, out := &in.VolumeAntiAffinity, &out.VolumeAntiAffinity
		*out = make([]*CommonPlacementSpec, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(CommonPlacementSpec)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumePlacementSpec.
func (in *VolumePlacementSpec) DeepCopy() *VolumePlacementSpec {
	if in == nil {
		return nil
	}
	out := new(VolumePlacementSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumePlacementStrategy) DeepCopyInto(out *VolumePlacementStrategy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumePlacementStrategy.
func (in *VolumePlacementStrategy) DeepCopy() *VolumePlacementStrategy {
	if in == nil {
		return nil
	}
	out := new(VolumePlacementStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumePlacementStrategy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumePlacementStrategyList) DeepCopyInto(out *VolumePlacementStrategyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VolumePlacementStrategy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumePlacementStrategyList.
func (in *VolumePlacementStrategyList) DeepCopy() *VolumePlacementStrategyList {
	if in == nil {
		return nil
	}
	out := new(VolumePlacementStrategyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumePlacementStrategyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
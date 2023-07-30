//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023 The Nephio Authors.

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
	invv1alpha1 "github.com/nokia/k8s-ipam/apis/inv/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogicalInterconnect) DeepCopyInto(out *LogicalInterconnect) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogicalInterconnect.
func (in *LogicalInterconnect) DeepCopy() *LogicalInterconnect {
	if in == nil {
		return nil
	}
	out := new(LogicalInterconnect)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LogicalInterconnect) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogicalInterconnectEndpoint) DeepCopyInto(out *LogicalInterconnectEndpoint) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Topologies != nil {
		in, out := &in.Topologies, &out.Topologies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.SelectorPolicy != nil {
		in, out := &in.SelectorPolicy, &out.SelectorPolicy
		*out = new(SelectorPolicy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogicalInterconnectEndpoint.
func (in *LogicalInterconnectEndpoint) DeepCopy() *LogicalInterconnectEndpoint {
	if in == nil {
		return nil
	}
	out := new(LogicalInterconnectEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogicalInterconnectList) DeepCopyInto(out *LogicalInterconnectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LogicalInterconnect, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogicalInterconnectList.
func (in *LogicalInterconnectList) DeepCopy() *LogicalInterconnectList {
	if in == nil {
		return nil
	}
	out := new(LogicalInterconnectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LogicalInterconnectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogicalInterconnectSpec) DeepCopyInto(out *LogicalInterconnectSpec) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Lacp != nil {
		in, out := &in.Lacp, &out.Lacp
		*out = new(bool)
		**out = **in
	}
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]LogicalInterconnectEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.UserDefinedLabels.DeepCopyInto(&out.UserDefinedLabels)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogicalInterconnectSpec.
func (in *LogicalInterconnectSpec) DeepCopy() *LogicalInterconnectSpec {
	if in == nil {
		return nil
	}
	out := new(LogicalInterconnectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogicalInterconnectStatus) DeepCopyInto(out *LogicalInterconnectStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogicalInterconnectStatus.
func (in *LogicalInterconnectStatus) DeepCopy() *LogicalInterconnectStatus {
	if in == nil {
		return nil
	}
	out := new(LogicalInterconnectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RawTopology) DeepCopyInto(out *RawTopology) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RawTopology.
func (in *RawTopology) DeepCopy() *RawTopology {
	if in == nil {
		return nil
	}
	out := new(RawTopology)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RawTopology) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RawTopologyList) DeepCopyInto(out *RawTopologyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RawTopology, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RawTopologyList.
func (in *RawTopologyList) DeepCopy() *RawTopologyList {
	if in == nil {
		return nil
	}
	out := new(RawTopologyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RawTopologyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RawTopologySpec) DeepCopyInto(out *RawTopologySpec) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make(map[string]invv1alpha1.NodeSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]invv1alpha1.LinkSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.UserDefinedLabels.DeepCopyInto(&out.UserDefinedLabels)
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(invv1alpha1.Location)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RawTopologySpec.
func (in *RawTopologySpec) DeepCopy() *RawTopologySpec {
	if in == nil {
		return nil
	}
	out := new(RawTopologySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RawTopologyStatus) DeepCopyInto(out *RawTopologyStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RawTopologyStatus.
func (in *RawTopologyStatus) DeepCopy() *RawTopologyStatus {
	if in == nil {
		return nil
	}
	out := new(RawTopologyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelectorPolicy) DeepCopyInto(out *SelectorPolicy) {
	*out = *in
	if in.NodeDiversity != nil {
		in, out := &in.NodeDiversity, &out.NodeDiversity
		*out = new(uint16)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelectorPolicy.
func (in *SelectorPolicy) DeepCopy() *SelectorPolicy {
	if in == nil {
		return nil
	}
	out := new(SelectorPolicy)
	in.DeepCopyInto(out)
	return out
}

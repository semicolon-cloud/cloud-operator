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

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudCluster) DeepCopyInto(out *CloudCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudCluster.
func (in *CloudCluster) DeepCopy() *CloudCluster {
	if in == nil {
		return nil
	}
	out := new(CloudCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudClusterList) DeepCopyInto(out *CloudClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudClusterList.
func (in *CloudClusterList) DeepCopy() *CloudClusterList {
	if in == nil {
		return nil
	}
	out := new(CloudClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudClusterSpec) DeepCopyInto(out *CloudClusterSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudClusterSpec.
func (in *CloudClusterSpec) DeepCopy() *CloudClusterSpec {
	if in == nil {
		return nil
	}
	out := new(CloudClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudClusterStatus) DeepCopyInto(out *CloudClusterStatus) {
	*out = *in
	if in.KeycloakCredential != nil {
		in, out := &in.KeycloakCredential, &out.KeycloakCredential
		*out = new(string)
		**out = **in
	}
	if in.KeycloakClientID != nil {
		in, out := &in.KeycloakClientID, &out.KeycloakClientID
		*out = new(string)
		**out = **in
	}
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.AddressName != nil {
		in, out := &in.AddressName, &out.AddressName
		*out = new(string)
		**out = **in
	}
	if in.BindingName != nil {
		in, out := &in.BindingName, &out.BindingName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudClusterStatus.
func (in *CloudClusterStatus) DeepCopy() *CloudClusterStatus {
	if in == nil {
		return nil
	}
	out := new(CloudClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkAddress) DeepCopyInto(out *NetworkAddress) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAddress.
func (in *NetworkAddress) DeepCopy() *NetworkAddress {
	if in == nil {
		return nil
	}
	out := new(NetworkAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkAddress) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkAddressBinding) DeepCopyInto(out *NetworkAddressBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAddressBinding.
func (in *NetworkAddressBinding) DeepCopy() *NetworkAddressBinding {
	if in == nil {
		return nil
	}
	out := new(NetworkAddressBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkAddressBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkAddressBindingList) DeepCopyInto(out *NetworkAddressBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkAddressBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAddressBindingList.
func (in *NetworkAddressBindingList) DeepCopy() *NetworkAddressBindingList {
	if in == nil {
		return nil
	}
	out := new(NetworkAddressBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkAddressBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkAddressBindingSpec) DeepCopyInto(out *NetworkAddressBindingSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAddressBindingSpec.
func (in *NetworkAddressBindingSpec) DeepCopy() *NetworkAddressBindingSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkAddressBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkAddressBindingStatus) DeepCopyInto(out *NetworkAddressBindingStatus) {
	*out = *in
	if in.TunnelName != nil {
		in, out := &in.TunnelName, &out.TunnelName
		*out = new(string)
		**out = **in
	}
	if in.CurrentRouteMapping != nil {
		in, out := &in.CurrentRouteMapping, &out.CurrentRouteMapping
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DNSRecordID != nil {
		in, out := &in.DNSRecordID, &out.DNSRecordID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAddressBindingStatus.
func (in *NetworkAddressBindingStatus) DeepCopy() *NetworkAddressBindingStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkAddressBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkAddressList) DeepCopyInto(out *NetworkAddressList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkAddress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAddressList.
func (in *NetworkAddressList) DeepCopy() *NetworkAddressList {
	if in == nil {
		return nil
	}
	out := new(NetworkAddressList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkAddressList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkAddressSpec) DeepCopyInto(out *NetworkAddressSpec) {
	*out = *in
	if in.AllowedHTTPPorts != nil {
		in, out := &in.AllowedHTTPPorts, &out.AllowedHTTPPorts
		*out = new(string)
		**out = **in
	}
	if in.AllowedTLSPorts != nil {
		in, out := &in.AllowedTLSPorts, &out.AllowedTLSPorts
		*out = new(string)
		**out = **in
	}
	if in.AllowedForwardPorts != nil {
		in, out := &in.AllowedForwardPorts, &out.AllowedForwardPorts
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAddressSpec.
func (in *NetworkAddressSpec) DeepCopy() *NetworkAddressSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkAddressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkAddressStatus) DeepCopyInto(out *NetworkAddressStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkAddressStatus.
func (in *NetworkAddressStatus) DeepCopy() *NetworkAddressStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkAddressStatus)
	in.DeepCopyInto(out)
	return out
}

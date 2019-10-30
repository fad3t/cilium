// +build !ignore_autogenerated

// Copyright 2017-2019 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by deepcopy-gen. DO NOT EDIT.

package v2

import (
	models "github.com/cilium/cilium/api/v1/models"
	api "github.com/cilium/cilium/pkg/policy/api"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationIP) DeepCopyInto(out *AllocationIP) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationIP.
func (in *AllocationIP) DeepCopy() *AllocationIP {
	if in == nil {
		return nil
	}
	out := new(AllocationIP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsSubnet) DeepCopyInto(out *AwsSubnet) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsSubnet.
func (in *AwsSubnet) DeepCopy() *AwsSubnet {
	if in == nil {
		return nil
	}
	out := new(AwsSubnet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsVPC) DeepCopyInto(out *AwsVPC) {
	*out = *in
	if in.CIDRs != nil {
		in, out := &in.CIDRs, &out.CIDRs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsVPC.
func (in *AwsVPC) DeepCopy() *AwsVPC {
	if in == nil {
		return nil
	}
	out := new(AwsVPC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumClusterwideNetworkPolicy) DeepCopyInto(out *CiliumClusterwideNetworkPolicy) {
	*out = *in
	if in.CiliumNetworkPolicy != nil {
		in, out := &in.CiliumNetworkPolicy, &out.CiliumNetworkPolicy
		*out = new(CiliumNetworkPolicy)
		(*in).DeepCopyInto(*out)
	}
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumClusterwideNetworkPolicy.
func (in *CiliumClusterwideNetworkPolicy) DeepCopy() *CiliumClusterwideNetworkPolicy {
	if in == nil {
		return nil
	}
	out := new(CiliumClusterwideNetworkPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CiliumClusterwideNetworkPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumClusterwideNetworkPolicyList) DeepCopyInto(out *CiliumClusterwideNetworkPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CiliumClusterwideNetworkPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumClusterwideNetworkPolicyList.
func (in *CiliumClusterwideNetworkPolicyList) DeepCopy() *CiliumClusterwideNetworkPolicyList {
	if in == nil {
		return nil
	}
	out := new(CiliumClusterwideNetworkPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CiliumClusterwideNetworkPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumEndpoint) DeepCopyInto(out *CiliumEndpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumEndpoint.
func (in *CiliumEndpoint) DeepCopy() *CiliumEndpoint {
	if in == nil {
		return nil
	}
	out := new(CiliumEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CiliumEndpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumEndpointList) DeepCopyInto(out *CiliumEndpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CiliumEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumEndpointList.
func (in *CiliumEndpointList) DeepCopy() *CiliumEndpointList {
	if in == nil {
		return nil
	}
	out := new(CiliumEndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CiliumEndpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumIdentity) DeepCopyInto(out *CiliumIdentity) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.SecurityLabels != nil {
		in, out := &in.SecurityLabels, &out.SecurityLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumIdentity.
func (in *CiliumIdentity) DeepCopy() *CiliumIdentity {
	if in == nil {
		return nil
	}
	out := new(CiliumIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CiliumIdentity) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumIdentityList) DeepCopyInto(out *CiliumIdentityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CiliumIdentity, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumIdentityList.
func (in *CiliumIdentityList) DeepCopy() *CiliumIdentityList {
	if in == nil {
		return nil
	}
	out := new(CiliumIdentityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CiliumIdentityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumNetworkPolicy) DeepCopyInto(out *CiliumNetworkPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(api.Rule)
		(*in).DeepCopyInto(*out)
	}
	if in.Specs != nil {
		in, out := &in.Specs, &out.Specs
		*out = make(api.Rules, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(api.Rule)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumNetworkPolicy.
func (in *CiliumNetworkPolicy) DeepCopy() *CiliumNetworkPolicy {
	if in == nil {
		return nil
	}
	out := new(CiliumNetworkPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CiliumNetworkPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumNetworkPolicyList) DeepCopyInto(out *CiliumNetworkPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CiliumNetworkPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumNetworkPolicyList.
func (in *CiliumNetworkPolicyList) DeepCopy() *CiliumNetworkPolicyList {
	if in == nil {
		return nil
	}
	out := new(CiliumNetworkPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CiliumNetworkPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumNetworkPolicyNodeStatus) DeepCopyInto(out *CiliumNetworkPolicyNodeStatus) {
	*out = *in
	in.LastUpdated.DeepCopyInto(&out.LastUpdated)
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumNetworkPolicyNodeStatus.
func (in *CiliumNetworkPolicyNodeStatus) DeepCopy() *CiliumNetworkPolicyNodeStatus {
	if in == nil {
		return nil
	}
	out := new(CiliumNetworkPolicyNodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumNetworkPolicyStatus) DeepCopyInto(out *CiliumNetworkPolicyStatus) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make(map[string]CiliumNetworkPolicyNodeStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.DerivativePolicies != nil {
		in, out := &in.DerivativePolicies, &out.DerivativePolicies
		*out = make(map[string]CiliumNetworkPolicyNodeStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumNetworkPolicyStatus.
func (in *CiliumNetworkPolicyStatus) DeepCopy() *CiliumNetworkPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(CiliumNetworkPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumNode) DeepCopyInto(out *CiliumNode) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumNode.
func (in *CiliumNode) DeepCopy() *CiliumNode {
	if in == nil {
		return nil
	}
	out := new(CiliumNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CiliumNode) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CiliumNodeList) DeepCopyInto(out *CiliumNodeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CiliumNode, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CiliumNodeList.
func (in *CiliumNodeList) DeepCopy() *CiliumNodeList {
	if in == nil {
		return nil
	}
	out := new(CiliumNodeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CiliumNodeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ENI) DeepCopyInto(out *ENI) {
	*out = *in
	out.Subnet = in.Subnet
	in.VPC.DeepCopyInto(&out.VPC)
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ENI.
func (in *ENI) DeepCopy() *ENI {
	if in == nil {
		return nil
	}
	out := new(ENI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ENISpec) DeepCopyInto(out *ENISpec) {
	*out = *in
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SubnetTags != nil {
		in, out := &in.SubnetTags, &out.SubnetTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DeleteOnTermination != nil {
		in, out := &in.DeleteOnTermination, &out.DeleteOnTermination
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ENISpec.
func (in *ENISpec) DeepCopy() *ENISpec {
	if in == nil {
		return nil
	}
	out := new(ENISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ENIStatus) DeepCopyInto(out *ENIStatus) {
	*out = *in
	if in.ENIs != nil {
		in, out := &in.ENIs, &out.ENIs
		*out = make(map[string]ENI, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ENIStatus.
func (in *ENIStatus) DeepCopy() *ENIStatus {
	if in == nil {
		return nil
	}
	out := new(ENIStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionSpec) DeepCopyInto(out *EncryptionSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionSpec.
func (in *EncryptionSpec) DeepCopy() *EncryptionSpec {
	if in == nil {
		return nil
	}
	out := new(EncryptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointIdentity) DeepCopyInto(out *EndpointIdentity) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointIdentity.
func (in *EndpointIdentity) DeepCopy() *EndpointIdentity {
	if in == nil {
		return nil
	}
	out := new(EndpointIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointNetworking) DeepCopyInto(out *EndpointNetworking) {
	*out = *in
	if in.Addressing != nil {
		in, out := &in.Addressing, &out.Addressing
		*out = make(AddressPairList, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(AddressPair)
				**out = **in
			}
		}
	}
	if in.HostAddressing != nil {
		in, out := &in.HostAddressing, &out.HostAddressing
		*out = new(models.NodeAddressing)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointNetworking.
func (in *EndpointNetworking) DeepCopy() *EndpointNetworking {
	if in == nil {
		return nil
	}
	out := new(EndpointNetworking)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthAddressingSpec) DeepCopyInto(out *HealthAddressingSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthAddressingSpec.
func (in *HealthAddressingSpec) DeepCopy() *HealthAddressingSpec {
	if in == nil {
		return nil
	}
	out := new(HealthAddressingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAMSpec) DeepCopyInto(out *IPAMSpec) {
	*out = *in
	if in.Pool != nil {
		in, out := &in.Pool, &out.Pool
		*out = make(map[string]AllocationIP, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodCIDRs != nil {
		in, out := &in.PodCIDRs, &out.PodCIDRs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAMSpec.
func (in *IPAMSpec) DeepCopy() *IPAMSpec {
	if in == nil {
		return nil
	}
	out := new(IPAMSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAMStatus) DeepCopyInto(out *IPAMStatus) {
	*out = *in
	if in.Used != nil {
		in, out := &in.Used, &out.Used
		*out = make(map[string]AllocationIP, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAMStatus.
func (in *IPAMStatus) DeepCopy() *IPAMStatus {
	if in == nil {
		return nil
	}
	out := new(IPAMStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityStatus) DeepCopyInto(out *IdentityStatus) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make(map[string]v1.Time, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityStatus.
func (in *IdentityStatus) DeepCopy() *IdentityStatus {
	if in == nil {
		return nil
	}
	out := new(IdentityStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeAddress) DeepCopyInto(out *NodeAddress) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeAddress.
func (in *NodeAddress) DeepCopy() *NodeAddress {
	if in == nil {
		return nil
	}
	out := new(NodeAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeSpec) DeepCopyInto(out *NodeSpec) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]NodeAddress, len(*in))
		copy(*out, *in)
	}
	out.HealthAddressing = in.HealthAddressing
	out.Encryption = in.Encryption
	in.ENI.DeepCopyInto(&out.ENI)
	in.IPAM.DeepCopyInto(&out.IPAM)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeSpec.
func (in *NodeSpec) DeepCopy() *NodeSpec {
	if in == nil {
		return nil
	}
	out := new(NodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeStatus) DeepCopyInto(out *NodeStatus) {
	*out = *in
	in.ENI.DeepCopyInto(&out.ENI)
	in.IPAM.DeepCopyInto(&out.IPAM)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeStatus.
func (in *NodeStatus) DeepCopy() *NodeStatus {
	if in == nil {
		return nil
	}
	out := new(NodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Timestamp.
func (in *Timestamp) DeepCopy() *Timestamp {
	if in == nil {
		return nil
	}
	out := new(Timestamp)
	in.DeepCopyInto(out)
	return out
}

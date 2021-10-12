//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-FileCopyrightText: 2021 iteratec GmbH
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoDiscoveryConfig) DeepCopyInto(out *AutoDiscoveryConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ControllerManagerConfigurationSpec.DeepCopyInto(&out.ControllerManagerConfigurationSpec)
	out.Cluster = in.Cluster
	out.ResourceInclusion = in.ResourceInclusion
	in.ServiceAutoDiscoveryConfig.DeepCopyInto(&out.ServiceAutoDiscoveryConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoDiscoveryConfig.
func (in *AutoDiscoveryConfig) DeepCopy() *AutoDiscoveryConfig {
	if in == nil {
		return nil
	}
	out := new(AutoDiscoveryConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutoDiscoveryConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfig) DeepCopyInto(out *ClusterConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfig.
func (in *ClusterConfig) DeepCopy() *ClusterConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceInclusionConfig) DeepCopyInto(out *ResourceInclusionConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceInclusionConfig.
func (in *ResourceInclusionConfig) DeepCopy() *ResourceInclusionConfig {
	if in == nil {
		return nil
	}
	out := new(ResourceInclusionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScanConfig) DeepCopyInto(out *ScanConfig) {
	*out = *in
	out.RepeatInterval = in.RepeatInterval
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScanConfig.
func (in *ScanConfig) DeepCopy() *ScanConfig {
	if in == nil {
		return nil
	}
	out := new(ScanConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAutoDiscoveryConfig) DeepCopyInto(out *ServiceAutoDiscoveryConfig) {
	*out = *in
	out.PassiveReconcileInterval = in.PassiveReconcileInterval
	in.ScanConfig.DeepCopyInto(&out.ScanConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAutoDiscoveryConfig.
func (in *ServiceAutoDiscoveryConfig) DeepCopy() *ServiceAutoDiscoveryConfig {
	if in == nil {
		return nil
	}
	out := new(ServiceAutoDiscoveryConfig)
	in.DeepCopyInto(out)
	return out
}

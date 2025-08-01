//go:build !ignore_autogenerated

/*
 Copyright (c) 2022 Avesha, Inc. All rights reserved.   SPDX-License-Identifier: Apache-2.0
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
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterHealth) DeepCopyInto(out *ClusterHealth) {
	*out = *in
	if in.ComponentStatuses != nil {
		in, out := &in.ComponentStatuses, &out.ComponentStatuses
		*out = make([]ComponentStatus, len(*in))
		copy(*out, *in)
	}
	in.LastUpdated.DeepCopyInto(&out.LastUpdated)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterHealth.
func (in *ClusterHealth) DeepCopy() *ClusterHealth {
	if in == nil {
		return nil
	}
	out := new(ClusterHealth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterProperty) DeepCopyInto(out *ClusterProperty) {
	*out = *in
	out.Telemetry = in.Telemetry
	out.GeoLocation = in.GeoLocation
	out.Monitoring = in.Monitoring
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterProperty.
func (in *ClusterProperty) DeepCopy() *ClusterProperty {
	if in == nil {
		return nil
	}
	out := new(ClusterProperty)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.NodeIPs != nil {
		in, out := &in.NodeIPs, &out.NodeIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.ClusterProperty = in.ClusterProperty
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	if in.CniSubnet != nil {
		in, out := &in.CniSubnet, &out.CniSubnet
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]NamespacesConfig, len(*in))
		copy(*out, *in)
	}
	if in.ClusterHealth != nil {
		in, out := &in.ClusterHealth, &out.ClusterHealth
		*out = new(ClusterHealth)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeIPs != nil {
		in, out := &in.NodeIPs, &out.NodeIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VCPURestriction != nil {
		in, out := &in.VCPURestriction, &out.VCPURestriction
		*out = new(VCPURestriction)
		(*in).DeepCopyInto(*out)
	}
	if in.GPURestriction != nil {
		in, out := &in.GPURestriction, &out.GPURestriction
		*out = new(GPURestriction)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentStatus) DeepCopyInto(out *ComponentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentStatus.
func (in *ComponentStatus) DeepCopy() *ComponentStatus {
	if in == nil {
		return nil
	}
	out := new(ComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalGatewayConfig) DeepCopyInto(out *ExternalGatewayConfig) {
	*out = *in
	out.Ingress = in.Ingress
	out.Egress = in.Egress
	out.NsIngress = in.NsIngress
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.VPCServiceAccess = in.VPCServiceAccess
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalGatewayConfig.
func (in *ExternalGatewayConfig) DeepCopy() *ExternalGatewayConfig {
	if in == nil {
		return nil
	}
	out := new(ExternalGatewayConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalGatewayConfigOptions) DeepCopyInto(out *ExternalGatewayConfigOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalGatewayConfigOptions.
func (in *ExternalGatewayConfigOptions) DeepCopy() *ExternalGatewayConfigOptions {
	if in == nil {
		return nil
	}
	out := new(ExternalGatewayConfigOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GPURestriction) DeepCopyInto(out *GPURestriction) {
	*out = *in
	in.LastUpdatedTimestamp.DeepCopyInto(&out.LastUpdatedTimestamp)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GPURestriction.
func (in *GPURestriction) DeepCopy() *GPURestriction {
	if in == nil {
		return nil
	}
	out := new(GPURestriction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GeoLocation) DeepCopyInto(out *GeoLocation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GeoLocation.
func (in *GeoLocation) DeepCopy() *GeoLocation {
	if in == nil {
		return nil
	}
	out := new(GeoLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesDashboard) DeepCopyInto(out *KubernetesDashboard) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesDashboard.
func (in *KubernetesDashboard) DeepCopy() *KubernetesDashboard {
	if in == nil {
		return nil
	}
	out := new(KubernetesDashboard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubesliceEvent) DeepCopyInto(out *KubesliceEvent) {
	*out = *in
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Timestamp.DeepCopyInto(&out.Timestamp)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubesliceEvent.
func (in *KubesliceEvent) DeepCopy() *KubesliceEvent {
	if in == nil {
		return nil
	}
	out := new(KubesliceEvent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Monitoring) DeepCopyInto(out *Monitoring) {
	*out = *in
	out.KubernetesDashboard = in.KubernetesDashboard
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Monitoring.
func (in *Monitoring) DeepCopy() *Monitoring {
	if in == nil {
		return nil
	}
	out := new(Monitoring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceIsolationProfile) DeepCopyInto(out *NamespaceIsolationProfile) {
	*out = *in
	if in.ApplicationNamespaces != nil {
		in, out := &in.ApplicationNamespaces, &out.ApplicationNamespaces
		*out = make([]SliceNamespaceSelection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AllowedNamespaces != nil {
		in, out := &in.AllowedNamespaces, &out.AllowedNamespaces
		*out = make([]SliceNamespaceSelection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceIsolationProfile.
func (in *NamespaceIsolationProfile) DeepCopy() *NamespaceIsolationProfile {
	if in == nil {
		return nil
	}
	out := new(NamespaceIsolationProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacesConfig) DeepCopyInto(out *NamespacesConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacesConfig.
func (in *NamespacesConfig) DeepCopy() *NamespacesConfig {
	if in == nil {
		return nil
	}
	out := new(NamespacesConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Project) DeepCopyInto(out *Project) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Project.
func (in *Project) DeepCopy() *Project {
	if in == nil {
		return nil
	}
	out := new(Project)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Project) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectList) DeepCopyInto(out *ProjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Project, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectList.
func (in *ProjectList) DeepCopy() *ProjectList {
	if in == nil {
		return nil
	}
	out := new(ProjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectSpec) DeepCopyInto(out *ProjectSpec) {
	*out = *in
	in.ServiceAccount.DeepCopyInto(&out.ServiceAccount)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectSpec.
func (in *ProjectSpec) DeepCopy() *ProjectSpec {
	if in == nil {
		return nil
	}
	out := new(ProjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectStatus) DeepCopyInto(out *ProjectStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectStatus.
func (in *ProjectStatus) DeepCopy() *ProjectStatus {
	if in == nil {
		return nil
	}
	out := new(ProjectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QOSProfile) DeepCopyInto(out *QOSProfile) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QOSProfile.
func (in *QOSProfile) DeepCopy() *QOSProfile {
	if in == nil {
		return nil
	}
	out := new(QOSProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccess) DeepCopyInto(out *ServiceAccess) {
	*out = *in
	out.Ingress = in.Ingress
	out.Egress = in.Egress
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccess.
func (in *ServiceAccess) DeepCopy() *ServiceAccess {
	if in == nil {
		return nil
	}
	out := new(ServiceAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccount) DeepCopyInto(out *ServiceAccount) {
	*out = *in
	if in.ReadOnly != nil {
		in, out := &in.ReadOnly, &out.ReadOnly
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ReadWrite != nil {
		in, out := &in.ReadWrite, &out.ReadWrite
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccount.
func (in *ServiceAccount) DeepCopy() *ServiceAccount {
	if in == nil {
		return nil
	}
	out := new(ServiceAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceDiscoveryEndpoint) DeepCopyInto(out *ServiceDiscoveryEndpoint) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceDiscoveryEndpoint.
func (in *ServiceDiscoveryEndpoint) DeepCopy() *ServiceDiscoveryEndpoint {
	if in == nil {
		return nil
	}
	out := new(ServiceDiscoveryEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceDiscoveryPort) DeepCopyInto(out *ServiceDiscoveryPort) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceDiscoveryPort.
func (in *ServiceDiscoveryPort) DeepCopy() *ServiceDiscoveryPort {
	if in == nil {
		return nil
	}
	out := new(ServiceDiscoveryPort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceExportConfig) DeepCopyInto(out *ServiceExportConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceExportConfig.
func (in *ServiceExportConfig) DeepCopy() *ServiceExportConfig {
	if in == nil {
		return nil
	}
	out := new(ServiceExportConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceExportConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceExportConfigList) DeepCopyInto(out *ServiceExportConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceExportConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceExportConfigList.
func (in *ServiceExportConfigList) DeepCopy() *ServiceExportConfigList {
	if in == nil {
		return nil
	}
	out := new(ServiceExportConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceExportConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceExportConfigSpec) DeepCopyInto(out *ServiceExportConfigSpec) {
	*out = *in
	if in.ServiceDiscoveryEndpoints != nil {
		in, out := &in.ServiceDiscoveryEndpoints, &out.ServiceDiscoveryEndpoints
		*out = make([]ServiceDiscoveryEndpoint, len(*in))
		copy(*out, *in)
	}
	if in.ServiceDiscoveryPorts != nil {
		in, out := &in.ServiceDiscoveryPorts, &out.ServiceDiscoveryPorts
		*out = make([]ServiceDiscoveryPort, len(*in))
		copy(*out, *in)
	}
	if in.Aliases != nil {
		in, out := &in.Aliases, &out.Aliases
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceExportConfigSpec.
func (in *ServiceExportConfigSpec) DeepCopy() *ServiceExportConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceExportConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceExportConfigStatus) DeepCopyInto(out *ServiceExportConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceExportConfigStatus.
func (in *ServiceExportConfigStatus) DeepCopy() *ServiceExportConfigStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceExportConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SliceConfig) DeepCopyInto(out *SliceConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SliceConfig.
func (in *SliceConfig) DeepCopy() *SliceConfig {
	if in == nil {
		return nil
	}
	out := new(SliceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SliceConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SliceConfigList) DeepCopyInto(out *SliceConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SliceConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SliceConfigList.
func (in *SliceConfigList) DeepCopy() *SliceConfigList {
	if in == nil {
		return nil
	}
	out := new(SliceConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SliceConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SliceConfigSpec) DeepCopyInto(out *SliceConfigSpec) {
	*out = *in
	if in.SliceGatewayProvider != nil {
		in, out := &in.SliceGatewayProvider, &out.SliceGatewayProvider
		*out = new(WorkerSliceGatewayProvider)
		(*in).DeepCopyInto(*out)
	}
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.QosProfileDetails != nil {
		in, out := &in.QosProfileDetails, &out.QosProfileDetails
		*out = new(QOSProfile)
		**out = **in
	}
	in.NamespaceIsolationProfile.DeepCopyInto(&out.NamespaceIsolationProfile)
	if in.ExternalGatewayConfig != nil {
		in, out := &in.ExternalGatewayConfig, &out.ExternalGatewayConfig
		*out = make([]ExternalGatewayConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RenewBefore != nil {
		in, out := &in.RenewBefore, &out.RenewBefore
		*out = (*in).DeepCopy()
	}
	if in.VPNConfig != nil {
		in, out := &in.VPNConfig, &out.VPNConfig
		*out = new(VPNConfiguration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SliceConfigSpec.
func (in *SliceConfigSpec) DeepCopy() *SliceConfigSpec {
	if in == nil {
		return nil
	}
	out := new(SliceConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SliceConfigStatus) DeepCopyInto(out *SliceConfigStatus) {
	*out = *in
	if in.KubesliceEvents != nil {
		in, out := &in.KubesliceEvents, &out.KubesliceEvents
		*out = make([]KubesliceEvent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SliceConfigStatus.
func (in *SliceConfigStatus) DeepCopy() *SliceConfigStatus {
	if in == nil {
		return nil
	}
	out := new(SliceConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SliceGatewayServiceType) DeepCopyInto(out *SliceGatewayServiceType) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SliceGatewayServiceType.
func (in *SliceGatewayServiceType) DeepCopy() *SliceGatewayServiceType {
	if in == nil {
		return nil
	}
	out := new(SliceGatewayServiceType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SliceNamespaceSelection) DeepCopyInto(out *SliceNamespaceSelection) {
	*out = *in
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SliceNamespaceSelection.
func (in *SliceNamespaceSelection) DeepCopy() *SliceNamespaceSelection {
	if in == nil {
		return nil
	}
	out := new(SliceNamespaceSelection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SliceQoSConfig) DeepCopyInto(out *SliceQoSConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SliceQoSConfig.
func (in *SliceQoSConfig) DeepCopy() *SliceQoSConfig {
	if in == nil {
		return nil
	}
	out := new(SliceQoSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SliceQoSConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SliceQoSConfigList) DeepCopyInto(out *SliceQoSConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SliceQoSConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SliceQoSConfigList.
func (in *SliceQoSConfigList) DeepCopy() *SliceQoSConfigList {
	if in == nil {
		return nil
	}
	out := new(SliceQoSConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SliceQoSConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SliceQoSConfigSpec) DeepCopyInto(out *SliceQoSConfigSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SliceQoSConfigSpec.
func (in *SliceQoSConfigSpec) DeepCopy() *SliceQoSConfigSpec {
	if in == nil {
		return nil
	}
	out := new(SliceQoSConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SliceQoSConfigStatus) DeepCopyInto(out *SliceQoSConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SliceQoSConfigStatus.
func (in *SliceQoSConfigStatus) DeepCopy() *SliceQoSConfigStatus {
	if in == nil {
		return nil
	}
	out := new(SliceQoSConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusOfKeyRotation) DeepCopyInto(out *StatusOfKeyRotation) {
	*out = *in
	in.LastUpdatedTimestamp.DeepCopyInto(&out.LastUpdatedTimestamp)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusOfKeyRotation.
func (in *StatusOfKeyRotation) DeepCopy() *StatusOfKeyRotation {
	if in == nil {
		return nil
	}
	out := new(StatusOfKeyRotation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Telemetry) DeepCopyInto(out *Telemetry) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Telemetry.
func (in *Telemetry) DeepCopy() *Telemetry {
	if in == nil {
		return nil
	}
	out := new(Telemetry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VCPURestriction) DeepCopyInto(out *VCPURestriction) {
	*out = *in
	in.LastUpdatedTimestamp.DeepCopyInto(&out.LastUpdatedTimestamp)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VCPURestriction.
func (in *VCPURestriction) DeepCopy() *VCPURestriction {
	if in == nil {
		return nil
	}
	out := new(VCPURestriction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPNConfiguration) DeepCopyInto(out *VPNConfiguration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPNConfiguration.
func (in *VPNConfiguration) DeepCopy() *VPNConfiguration {
	if in == nil {
		return nil
	}
	out := new(VPNConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnKeyRotation) DeepCopyInto(out *VpnKeyRotation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnKeyRotation.
func (in *VpnKeyRotation) DeepCopy() *VpnKeyRotation {
	if in == nil {
		return nil
	}
	out := new(VpnKeyRotation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpnKeyRotation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnKeyRotationList) DeepCopyInto(out *VpnKeyRotationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VpnKeyRotation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnKeyRotationList.
func (in *VpnKeyRotationList) DeepCopy() *VpnKeyRotationList {
	if in == nil {
		return nil
	}
	out := new(VpnKeyRotationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpnKeyRotationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnKeyRotationSpec) DeepCopyInto(out *VpnKeyRotationSpec) {
	*out = *in
	if in.ClusterGatewayMapping != nil {
		in, out := &in.ClusterGatewayMapping, &out.ClusterGatewayMapping
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.CertificateCreationTime != nil {
		in, out := &in.CertificateCreationTime, &out.CertificateCreationTime
		*out = (*in).DeepCopy()
	}
	if in.CertificateExpiryTime != nil {
		in, out := &in.CertificateExpiryTime, &out.CertificateExpiryTime
		*out = (*in).DeepCopy()
	}
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnKeyRotationSpec.
func (in *VpnKeyRotationSpec) DeepCopy() *VpnKeyRotationSpec {
	if in == nil {
		return nil
	}
	out := new(VpnKeyRotationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnKeyRotationStatus) DeepCopyInto(out *VpnKeyRotationStatus) {
	*out = *in
	if in.CurrentRotationState != nil {
		in, out := &in.CurrentRotationState, &out.CurrentRotationState
		*out = make(map[string]StatusOfKeyRotation, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.StatusHistory != nil {
		in, out := &in.StatusHistory, &out.StatusHistory
		*out = make(map[string][]StatusOfKeyRotation, len(*in))
		for key, val := range *in {
			var outVal []StatusOfKeyRotation
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = make([]StatusOfKeyRotation, len(*in))
				for i := range *in {
					(*in)[i].DeepCopyInto(&(*out)[i])
				}
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnKeyRotationStatus.
func (in *VpnKeyRotationStatus) DeepCopy() *VpnKeyRotationStatus {
	if in == nil {
		return nil
	}
	out := new(VpnKeyRotationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerSliceGatewayProvider) DeepCopyInto(out *WorkerSliceGatewayProvider) {
	*out = *in
	if in.SliceGatewayServiceType != nil {
		in, out := &in.SliceGatewayServiceType, &out.SliceGatewayServiceType
		*out = make([]SliceGatewayServiceType, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerSliceGatewayProvider.
func (in *WorkerSliceGatewayProvider) DeepCopy() *WorkerSliceGatewayProvider {
	if in == nil {
		return nil
	}
	out := new(WorkerSliceGatewayProvider)
	in.DeepCopyInto(out)
	return out
}

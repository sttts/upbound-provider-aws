// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DefaultVPCInitParameters struct {

	// and instance_tenancy arguments become computed attributes
	AssignGeneratedIPv6CidrBlock *bool `json:"assignGeneratedIpv6CidrBlock,omitempty" tf:"assign_generated_ipv6_cidr_block,omitempty"`

	// is true
	EnableDNSHostnames *bool `json:"enableDnsHostnames,omitempty" tf:"enable_dns_hostnames,omitempty"`

	EnableDNSSupport *bool `json:"enableDnsSupport,omitempty" tf:"enable_dns_support,omitempty"`

	EnableNetworkAddressUsageMetrics *bool `json:"enableNetworkAddressUsageMetrics,omitempty" tf:"enable_network_address_usage_metrics,omitempty"`

	// Whether destroying the resource deletes the default VPC. Default: false
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// and instance_tenancy arguments become computed attributes
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`

	IPv6CidrBlockNetworkBorderGroup *string `json:"ipv6CidrBlockNetworkBorderGroup,omitempty" tf:"ipv6_cidr_block_network_border_group,omitempty"`

	IPv6IpamPoolID *string `json:"ipv6IpamPoolId,omitempty" tf:"ipv6_ipam_pool_id,omitempty"`

	IPv6NetmaskLength *int64 `json:"ipv6NetmaskLength,omitempty" tf:"ipv6_netmask_length,omitempty"`

	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DefaultVPCObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// and instance_tenancy arguments become computed attributes
	AssignGeneratedIPv6CidrBlock *bool `json:"assignGeneratedIpv6CidrBlock,omitempty" tf:"assign_generated_ipv6_cidr_block,omitempty"`

	// and instance_tenancy arguments become computed attributes
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`

	DHCPOptionsID *string `json:"dhcpOptionsId,omitempty" tf:"dhcp_options_id,omitempty"`

	DefaultNetworkACLID *string `json:"defaultNetworkAclId,omitempty" tf:"default_network_acl_id,omitempty"`

	DefaultRouteTableID *string `json:"defaultRouteTableId,omitempty" tf:"default_route_table_id,omitempty"`

	DefaultSecurityGroupID *string `json:"defaultSecurityGroupId,omitempty" tf:"default_security_group_id,omitempty"`

	// is true
	EnableDNSHostnames *bool `json:"enableDnsHostnames,omitempty" tf:"enable_dns_hostnames,omitempty"`

	EnableDNSSupport *bool `json:"enableDnsSupport,omitempty" tf:"enable_dns_support,omitempty"`

	EnableNetworkAddressUsageMetrics *bool `json:"enableNetworkAddressUsageMetrics,omitempty" tf:"enable_network_address_usage_metrics,omitempty"`

	ExistingDefaultVPC *bool `json:"existingDefaultVpc,omitempty" tf:"existing_default_vpc,omitempty"`

	// Whether destroying the resource deletes the default VPC. Default: false
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IPv6AssociationID *string `json:"ipv6AssociationId,omitempty" tf:"ipv6_association_id,omitempty"`

	// and instance_tenancy arguments become computed attributes
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`

	IPv6CidrBlockNetworkBorderGroup *string `json:"ipv6CidrBlockNetworkBorderGroup,omitempty" tf:"ipv6_cidr_block_network_border_group,omitempty"`

	IPv6IpamPoolID *string `json:"ipv6IpamPoolId,omitempty" tf:"ipv6_ipam_pool_id,omitempty"`

	IPv6NetmaskLength *int64 `json:"ipv6NetmaskLength,omitempty" tf:"ipv6_netmask_length,omitempty"`

	// The allowed tenancy of instances launched into the VPC
	InstanceTenancy *string `json:"instanceTenancy,omitempty" tf:"instance_tenancy,omitempty"`

	MainRouteTableID *string `json:"mainRouteTableId,omitempty" tf:"main_route_table_id,omitempty"`

	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type DefaultVPCParameters struct {

	// and instance_tenancy arguments become computed attributes
	// +kubebuilder:validation:Optional
	AssignGeneratedIPv6CidrBlock *bool `json:"assignGeneratedIpv6CidrBlock,omitempty" tf:"assign_generated_ipv6_cidr_block,omitempty"`

	// is true
	// +kubebuilder:validation:Optional
	EnableDNSHostnames *bool `json:"enableDnsHostnames,omitempty" tf:"enable_dns_hostnames,omitempty"`

	// +kubebuilder:validation:Optional
	EnableDNSSupport *bool `json:"enableDnsSupport,omitempty" tf:"enable_dns_support,omitempty"`

	// +kubebuilder:validation:Optional
	EnableNetworkAddressUsageMetrics *bool `json:"enableNetworkAddressUsageMetrics,omitempty" tf:"enable_network_address_usage_metrics,omitempty"`

	// Whether destroying the resource deletes the default VPC. Default: false
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// and instance_tenancy arguments become computed attributes
	// +kubebuilder:validation:Optional
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6CidrBlockNetworkBorderGroup *string `json:"ipv6CidrBlockNetworkBorderGroup,omitempty" tf:"ipv6_cidr_block_network_border_group,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6IpamPoolID *string `json:"ipv6IpamPoolId,omitempty" tf:"ipv6_ipam_pool_id,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6NetmaskLength *int64 `json:"ipv6NetmaskLength,omitempty" tf:"ipv6_netmask_length,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// DefaultVPCSpec defines the desired state of DefaultVPC
type DefaultVPCSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DefaultVPCParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider DefaultVPCInitParameters `json:"initProvider,omitempty"`
}

// DefaultVPCStatus defines the observed state of DefaultVPC.
type DefaultVPCStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DefaultVPCObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultVPC is the Schema for the DefaultVPCs API. Manage a default VPC resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DefaultVPC struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultVPCSpec   `json:"spec"`
	Status            DefaultVPCStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultVPCList contains a list of DefaultVPCs
type DefaultVPCList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultVPC `json:"items"`
}

// Repository type metadata.
var (
	DefaultVPC_Kind             = "DefaultVPC"
	DefaultVPC_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DefaultVPC_Kind}.String()
	DefaultVPC_KindAPIVersion   = DefaultVPC_Kind + "." + CRDGroupVersion.String()
	DefaultVPC_GroupVersionKind = CRDGroupVersion.WithKind(DefaultVPC_Kind)
)

func init() {
	SchemeBuilder.Register(&DefaultVPC{}, &DefaultVPCList{})
}

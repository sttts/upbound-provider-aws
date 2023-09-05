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

type TransitGatewayConnectInitParameters struct {

	// The tunnel protocol. Valid values: gre. Default is gre.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Boolean whether the Connect should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: true.
	TransitGatewayDefaultRouteTableAssociation *bool `json:"transitGatewayDefaultRouteTableAssociation,omitempty" tf:"transit_gateway_default_route_table_association,omitempty"`

	// Boolean whether the Connect should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: true.
	TransitGatewayDefaultRouteTablePropagation *bool `json:"transitGatewayDefaultRouteTablePropagation,omitempty" tf:"transit_gateway_default_route_table_propagation,omitempty"`
}

type TransitGatewayConnectObservation struct {

	// EC2 Transit Gateway Attachment identifier
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The tunnel protocol. Valid values: gre. Default is gre.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Boolean whether the Connect should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: true.
	TransitGatewayDefaultRouteTableAssociation *bool `json:"transitGatewayDefaultRouteTableAssociation,omitempty" tf:"transit_gateway_default_route_table_association,omitempty"`

	// Boolean whether the Connect should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: true.
	TransitGatewayDefaultRouteTablePropagation *bool `json:"transitGatewayDefaultRouteTablePropagation,omitempty" tf:"transit_gateway_default_route_table_propagation,omitempty"`

	// Identifier of EC2 Transit Gateway.
	TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id,omitempty"`

	// The underlaying VPC attachment
	TransportAttachmentID *string `json:"transportAttachmentId,omitempty" tf:"transport_attachment_id,omitempty"`
}

type TransitGatewayConnectParameters struct {

	// The tunnel protocol. Valid values: gre. Default is gre.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Boolean whether the Connect should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: true.
	// +kubebuilder:validation:Optional
	TransitGatewayDefaultRouteTableAssociation *bool `json:"transitGatewayDefaultRouteTableAssociation,omitempty" tf:"transit_gateway_default_route_table_association,omitempty"`

	// Boolean whether the Connect should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: true.
	// +kubebuilder:validation:Optional
	TransitGatewayDefaultRouteTablePropagation *bool `json:"transitGatewayDefaultRouteTablePropagation,omitempty" tf:"transit_gateway_default_route_table_propagation,omitempty"`

	// Identifier of EC2 Transit Gateway.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.TransitGateway
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id,omitempty"`

	// Reference to a TransitGateway in ec2 to populate transitGatewayId.
	// +kubebuilder:validation:Optional
	TransitGatewayIDRef *v1.Reference `json:"transitGatewayIdRef,omitempty" tf:"-"`

	// Selector for a TransitGateway in ec2 to populate transitGatewayId.
	// +kubebuilder:validation:Optional
	TransitGatewayIDSelector *v1.Selector `json:"transitGatewayIdSelector,omitempty" tf:"-"`

	// The underlaying VPC attachment
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.TransitGatewayVPCAttachment
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TransportAttachmentID *string `json:"transportAttachmentId,omitempty" tf:"transport_attachment_id,omitempty"`

	// Reference to a TransitGatewayVPCAttachment in ec2 to populate transportAttachmentId.
	// +kubebuilder:validation:Optional
	TransportAttachmentIDRef *v1.Reference `json:"transportAttachmentIdRef,omitempty" tf:"-"`

	// Selector for a TransitGatewayVPCAttachment in ec2 to populate transportAttachmentId.
	// +kubebuilder:validation:Optional
	TransportAttachmentIDSelector *v1.Selector `json:"transportAttachmentIdSelector,omitempty" tf:"-"`
}

// TransitGatewayConnectSpec defines the desired state of TransitGatewayConnect
type TransitGatewayConnectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitGatewayConnectParameters `json:"forProvider"`
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
	InitProvider TransitGatewayConnectInitParameters `json:"initProvider,omitempty"`
}

// TransitGatewayConnectStatus defines the observed state of TransitGatewayConnect.
type TransitGatewayConnectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitGatewayConnectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayConnect is the Schema for the TransitGatewayConnects API. Manages an EC2 Transit Gateway Connect
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TransitGatewayConnect struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitGatewayConnectSpec   `json:"spec"`
	Status            TransitGatewayConnectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayConnectList contains a list of TransitGatewayConnects
type TransitGatewayConnectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGatewayConnect `json:"items"`
}

// Repository type metadata.
var (
	TransitGatewayConnect_Kind             = "TransitGatewayConnect"
	TransitGatewayConnect_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitGatewayConnect_Kind}.String()
	TransitGatewayConnect_KindAPIVersion   = TransitGatewayConnect_Kind + "." + CRDGroupVersion.String()
	TransitGatewayConnect_GroupVersionKind = CRDGroupVersion.WithKind(TransitGatewayConnect_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitGatewayConnect{}, &TransitGatewayConnectList{})
}

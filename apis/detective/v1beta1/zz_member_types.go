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

type MemberInitParameters struct {

	// AWS account ID for the account.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// If set to true, then the root user of the invited account will not receive an email notification. This notification is in addition to an alert that the root user receives in AWS Personal Health Dashboard. By default, this is set to false.
	DisableEmailNotification *bool `json:"disableEmailNotification,omitempty" tf:"disable_email_notification,omitempty"`

	// Email address for the account.
	EmailAddress *string `json:"emailAddress,omitempty" tf:"email_address,omitempty"`

	// A custom message to include in the invitation. Amazon Detective adds this message to the standard content that it sends for an invitation.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`
}

type MemberObservation struct {

	// AWS account ID for the account.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// AWS account ID for the administrator account.
	AdministratorID *string `json:"administratorId,omitempty" tf:"administrator_id,omitempty"`

	// If set to true, then the root user of the invited account will not receive an email notification. This notification is in addition to an alert that the root user receives in AWS Personal Health Dashboard. By default, this is set to false.
	DisableEmailNotification *bool `json:"disableEmailNotification,omitempty" tf:"disable_email_notification,omitempty"`

	DisabledReason *string `json:"disabledReason,omitempty" tf:"disabled_reason,omitempty"`

	// Email address for the account.
	EmailAddress *string `json:"emailAddress,omitempty" tf:"email_address,omitempty"`

	// ARN of the behavior graph to invite the member accounts to contribute their data to.
	GraphArn *string `json:"graphArn,omitempty" tf:"graph_arn,omitempty"`

	// Unique identifier (ID) of the Detective.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Date and time, in UTC and extended RFC 3339 format, when an Amazon Detective membership invitation was last sent to the account.
	InvitedTime *string `json:"invitedTime,omitempty" tf:"invited_time,omitempty"`

	// A custom message to include in the invitation. Amazon Detective adds this message to the standard content that it sends for an invitation.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// Current membership status of the member account.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Date and time, in UTC and extended RFC 3339 format, of the most recent change to the member account's status.
	UpdatedTime *string `json:"updatedTime,omitempty" tf:"updated_time,omitempty"`

	// Data volume in bytes per day for the member account.
	VolumeUsageInBytes *string `json:"volumeUsageInBytes,omitempty" tf:"volume_usage_in_bytes,omitempty"`
}

type MemberParameters struct {

	// AWS account ID for the account.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// If set to true, then the root user of the invited account will not receive an email notification. This notification is in addition to an alert that the root user receives in AWS Personal Health Dashboard. By default, this is set to false.
	// +kubebuilder:validation:Optional
	DisableEmailNotification *bool `json:"disableEmailNotification,omitempty" tf:"disable_email_notification,omitempty"`

	// Email address for the account.
	// +kubebuilder:validation:Optional
	EmailAddress *string `json:"emailAddress,omitempty" tf:"email_address,omitempty"`

	// ARN of the behavior graph to invite the member accounts to contribute their data to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/detective/v1beta1.Graph
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	GraphArn *string `json:"graphArn,omitempty" tf:"graph_arn,omitempty"`

	// Reference to a Graph in detective to populate graphArn.
	// +kubebuilder:validation:Optional
	GraphArnRef *v1.Reference `json:"graphArnRef,omitempty" tf:"-"`

	// Selector for a Graph in detective to populate graphArn.
	// +kubebuilder:validation:Optional
	GraphArnSelector *v1.Selector `json:"graphArnSelector,omitempty" tf:"-"`

	// A custom message to include in the invitation. Amazon Detective adds this message to the standard content that it sends for an invitation.
	// +kubebuilder:validation:Optional
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// MemberSpec defines the desired state of Member
type MemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MemberParameters `json:"forProvider"`
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
	InitProvider MemberInitParameters `json:"initProvider,omitempty"`
}

// MemberStatus defines the observed state of Member.
type MemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Member is the Schema for the Members API. Provides a resource to manage an Amazon Detective member.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Member struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.accountId) || (has(self.initProvider) && has(self.initProvider.accountId))",message="spec.forProvider.accountId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.emailAddress) || (has(self.initProvider) && has(self.initProvider.emailAddress))",message="spec.forProvider.emailAddress is a required parameter"
	Spec   MemberSpec   `json:"spec"`
	Status MemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MemberList contains a list of Members
type MemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Member `json:"items"`
}

// Repository type metadata.
var (
	Member_Kind             = "Member"
	Member_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Member_Kind}.String()
	Member_KindAPIVersion   = Member_Kind + "." + CRDGroupVersion.String()
	Member_GroupVersionKind = CRDGroupVersion.WithKind(Member_Kind)
)

func init() {
	SchemeBuilder.Register(&Member{}, &MemberList{})
}

// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type UserGroupInitParameters struct {

	// The current supported value is REDIS.
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type UserGroupObservation struct {

	// The ARN that identifies the user group.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The current supported value is REDIS.
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// The user group identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The list of user IDs that belong to the user group.
	UserIds []*string `json:"userIds,omitempty" tf:"user_ids,omitempty"`
}

type UserGroupParameters struct {

	// The current supported value is REDIS.
	// +kubebuilder:validation:Optional
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// References to User to populate userIds.
	// +kubebuilder:validation:Optional
	UserIDRefs []v1.Reference `json:"userIdRefs,omitempty" tf:"-"`

	// Selector for a list of User to populate userIds.
	// +kubebuilder:validation:Optional
	UserIDSelector *v1.Selector `json:"userIdSelector,omitempty" tf:"-"`

	// The list of user IDs that belong to the user group.
	// +crossplane:generate:reference:type=User
	// +crossplane:generate:reference:refFieldName=UserIDRefs
	// +crossplane:generate:reference:selectorFieldName=UserIDSelector
	// +kubebuilder:validation:Optional
	UserIds []*string `json:"userIds,omitempty" tf:"user_ids,omitempty"`
}

// UserGroupSpec defines the desired state of UserGroup
type UserGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserGroupParameters `json:"forProvider"`
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
	InitProvider UserGroupInitParameters `json:"initProvider,omitempty"`
}

// UserGroupStatus defines the observed state of UserGroup.
type UserGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserGroup is the Schema for the UserGroups API. Provides an ElastiCache user group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type UserGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.engine) || (has(self.initProvider) && has(self.initProvider.engine))",message="spec.forProvider.engine is a required parameter"
	Spec   UserGroupSpec   `json:"spec"`
	Status UserGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserGroupList contains a list of UserGroups
type UserGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserGroup `json:"items"`
}

// Repository type metadata.
var (
	UserGroup_Kind             = "UserGroup"
	UserGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserGroup_Kind}.String()
	UserGroup_KindAPIVersion   = UserGroup_Kind + "." + CRDGroupVersion.String()
	UserGroup_GroupVersionKind = CRDGroupVersion.WithKind(UserGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&UserGroup{}, &UserGroupList{})
}

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

type ConfigurationProfileInitParameters struct {

	// Description of the configuration profile. Can be at most 1024 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// URI to locate the configuration. You can specify the AWS AppConfig hosted configuration store, Systems Manager (SSM) document, an SSM Parameter Store parameter, or an Amazon S3 object. For the hosted configuration store, specify hosted. For an SSM document, specify either the document name in the format ssm-document://<Document_name> or the ARN. For a parameter, specify either the parameter name in the format ssm-parameter://<Parameter_name> or the ARN. For an Amazon S3 object, specify the URI in the following format: s3://<bucket>/<objectKey>.
	LocationURI *string `json:"locationUri,omitempty" tf:"location_uri,omitempty"`

	// Name for the configuration profile. Must be between 1 and 64 characters in length.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Type of configurations contained in the profile. Valid values: AWS.AppConfig.FeatureFlags and AWS.Freeform.  Default: AWS.Freeform.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Set of methods for validating the configuration. Maximum of 2. See Validator below for more details.
	Validator []ValidatorInitParameters `json:"validator,omitempty" tf:"validator,omitempty"`
}

type ConfigurationProfileObservation struct {

	// Application ID. Must be between 4 and 7 characters in length.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// ARN of the AppConfig Configuration Profile.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The configuration profile ID.
	ConfigurationProfileID *string `json:"configurationProfileId,omitempty" tf:"configuration_profile_id,omitempty"`

	// Description of the configuration profile. Can be at most 1024 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// AppConfig configuration profile ID and application ID separated by a colon (:).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// URI to locate the configuration. You can specify the AWS AppConfig hosted configuration store, Systems Manager (SSM) document, an SSM Parameter Store parameter, or an Amazon S3 object. For the hosted configuration store, specify hosted. For an SSM document, specify either the document name in the format ssm-document://<Document_name> or the ARN. For a parameter, specify either the parameter name in the format ssm-parameter://<Parameter_name> or the ARN. For an Amazon S3 object, specify the URI in the following format: s3://<bucket>/<objectKey>.
	LocationURI *string `json:"locationUri,omitempty" tf:"location_uri,omitempty"`

	// Name for the configuration profile. Must be between 1 and 64 characters in length.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ARN of an IAM role with permission to access the configuration at the specified location_uri. A retrieval role ARN is not required for configurations stored in the AWS AppConfig hosted configuration store. It is required for all other sources that store your configuration.
	RetrievalRoleArn *string `json:"retrievalRoleArn,omitempty" tf:"retrieval_role_arn,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Type of configurations contained in the profile. Valid values: AWS.AppConfig.FeatureFlags and AWS.Freeform.  Default: AWS.Freeform.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Set of methods for validating the configuration. Maximum of 2. See Validator below for more details.
	Validator []ValidatorObservation `json:"validator,omitempty" tf:"validator,omitempty"`
}

type ConfigurationProfileParameters struct {

	// Application ID. Must be between 4 and 7 characters in length.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appconfig/v1beta1.Application
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// Reference to a Application in appconfig to populate applicationId.
	// +kubebuilder:validation:Optional
	ApplicationIDRef *v1.Reference `json:"applicationIdRef,omitempty" tf:"-"`

	// Selector for a Application in appconfig to populate applicationId.
	// +kubebuilder:validation:Optional
	ApplicationIDSelector *v1.Selector `json:"applicationIdSelector,omitempty" tf:"-"`

	// Description of the configuration profile. Can be at most 1024 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// URI to locate the configuration. You can specify the AWS AppConfig hosted configuration store, Systems Manager (SSM) document, an SSM Parameter Store parameter, or an Amazon S3 object. For the hosted configuration store, specify hosted. For an SSM document, specify either the document name in the format ssm-document://<Document_name> or the ARN. For a parameter, specify either the parameter name in the format ssm-parameter://<Parameter_name> or the ARN. For an Amazon S3 object, specify the URI in the following format: s3://<bucket>/<objectKey>.
	// +kubebuilder:validation:Optional
	LocationURI *string `json:"locationUri,omitempty" tf:"location_uri,omitempty"`

	// Name for the configuration profile. Must be between 1 and 64 characters in length.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// ARN of an IAM role with permission to access the configuration at the specified location_uri. A retrieval role ARN is not required for configurations stored in the AWS AppConfig hosted configuration store. It is required for all other sources that store your configuration.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RetrievalRoleArn *string `json:"retrievalRoleArn,omitempty" tf:"retrieval_role_arn,omitempty"`

	// Reference to a Role in iam to populate retrievalRoleArn.
	// +kubebuilder:validation:Optional
	RetrievalRoleArnRef *v1.Reference `json:"retrievalRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate retrievalRoleArn.
	// +kubebuilder:validation:Optional
	RetrievalRoleArnSelector *v1.Selector `json:"retrievalRoleArnSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Type of configurations contained in the profile. Valid values: AWS.AppConfig.FeatureFlags and AWS.Freeform.  Default: AWS.Freeform.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Set of methods for validating the configuration. Maximum of 2. See Validator below for more details.
	// +kubebuilder:validation:Optional
	Validator []ValidatorParameters `json:"validator,omitempty" tf:"validator,omitempty"`
}

type ValidatorInitParameters struct {

	// Type of validator. Valid values: JSON_SCHEMA and LAMBDA.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ValidatorObservation struct {

	// Type of validator. Valid values: JSON_SCHEMA and LAMBDA.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ValidatorParameters struct {

	// Either the JSON Schema content or the ARN of an AWS Lambda function.
	// +kubebuilder:validation:Optional
	ContentSecretRef *v1.SecretKeySelector `json:"contentSecretRef,omitempty" tf:"-"`

	// Type of validator. Valid values: JSON_SCHEMA and LAMBDA.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// ConfigurationProfileSpec defines the desired state of ConfigurationProfile
type ConfigurationProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConfigurationProfileParameters `json:"forProvider"`
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
	InitProvider ConfigurationProfileInitParameters `json:"initProvider,omitempty"`
}

// ConfigurationProfileStatus defines the observed state of ConfigurationProfile.
type ConfigurationProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConfigurationProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigurationProfile is the Schema for the ConfigurationProfiles API. Provides an AppConfig Configuration Profile resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ConfigurationProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.locationUri) || (has(self.initProvider) && has(self.initProvider.locationUri))",message="spec.forProvider.locationUri is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ConfigurationProfileSpec   `json:"spec"`
	Status ConfigurationProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigurationProfileList contains a list of ConfigurationProfiles
type ConfigurationProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigurationProfile `json:"items"`
}

// Repository type metadata.
var (
	ConfigurationProfile_Kind             = "ConfigurationProfile"
	ConfigurationProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConfigurationProfile_Kind}.String()
	ConfigurationProfile_KindAPIVersion   = ConfigurationProfile_Kind + "." + CRDGroupVersion.String()
	ConfigurationProfile_GroupVersionKind = CRDGroupVersion.WithKind(ConfigurationProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&ConfigurationProfile{}, &ConfigurationProfileList{})
}

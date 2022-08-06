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

type IntegrationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IntegrationResponseSelectionExpression *string `json:"integrationResponseSelectionExpression,omitempty" tf:"integration_response_selection_expression,omitempty"`
}

type IntegrationParameters struct {

	// +crossplane:generate:reference:type=API
	// +kubebuilder:validation:Optional
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// +kubebuilder:validation:Optional
	APIIDRef *v1.Reference `json:"apiIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	APIIDSelector *v1.Selector `json:"apiIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/apigatewayv2/v1beta1.VPCLink
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ConnectionID *string `json:"connectionId,omitempty" tf:"connection_id,omitempty"`

	// +kubebuilder:validation:Optional
	ConnectionIDRef *v1.Reference `json:"connectionIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ConnectionIDSelector *v1.Selector `json:"connectionIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ConnectionType *string `json:"connectionType,omitempty" tf:"connection_type,omitempty"`

	// +kubebuilder:validation:Optional
	ContentHandlingStrategy *string `json:"contentHandlingStrategy,omitempty" tf:"content_handling_strategy,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	CredentialsArn *string `json:"credentialsArn,omitempty" tf:"credentials_arn,omitempty"`

	// +kubebuilder:validation:Optional
	CredentialsArnRef *v1.Reference `json:"credentialsArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	CredentialsArnSelector *v1.Selector `json:"credentialsArnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	IntegrationMethod *string `json:"integrationMethod,omitempty" tf:"integration_method,omitempty"`

	// +kubebuilder:validation:Optional
	IntegrationSubtype *string `json:"integrationSubtype,omitempty" tf:"integration_subtype,omitempty"`

	// +kubebuilder:validation:Required
	IntegrationType *string `json:"integrationType" tf:"integration_type,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/lambda/v1beta1.Function
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("invoke_arn",true)
	// +kubebuilder:validation:Optional
	IntegrationURI *string `json:"integrationUri,omitempty" tf:"integration_uri,omitempty"`

	// +kubebuilder:validation:Optional
	IntegrationURIRef *v1.Reference `json:"integrationUriRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	IntegrationURISelector *v1.Selector `json:"integrationUriSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PassthroughBehavior *string `json:"passthroughBehavior,omitempty" tf:"passthrough_behavior,omitempty"`

	// +kubebuilder:validation:Optional
	PayloadFormatVersion *string `json:"payloadFormatVersion,omitempty" tf:"payload_format_version,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RequestParameters map[string]*string `json:"requestParameters,omitempty" tf:"request_parameters,omitempty"`

	// +kubebuilder:validation:Optional
	RequestTemplates map[string]*string `json:"requestTemplates,omitempty" tf:"request_templates,omitempty"`

	// +kubebuilder:validation:Optional
	ResponseParameters []ResponseParametersParameters `json:"responseParameters,omitempty" tf:"response_parameters,omitempty"`

	// +kubebuilder:validation:Optional
	TLSConfig []TLSConfigParameters `json:"tlsConfig,omitempty" tf:"tls_config,omitempty"`

	// +kubebuilder:validation:Optional
	TemplateSelectionExpression *string `json:"templateSelectionExpression,omitempty" tf:"template_selection_expression,omitempty"`

	// +kubebuilder:validation:Optional
	TimeoutMilliseconds *float64 `json:"timeoutMilliseconds,omitempty" tf:"timeout_milliseconds,omitempty"`
}

type ResponseParametersObservation struct {
}

type ResponseParametersParameters struct {

	// +kubebuilder:validation:Required
	Mappings map[string]*string `json:"mappings" tf:"mappings,omitempty"`

	// +kubebuilder:validation:Required
	StatusCode *string `json:"statusCode" tf:"status_code,omitempty"`
}

type TLSConfigObservation struct {
}

type TLSConfigParameters struct {

	// +kubebuilder:validation:Optional
	ServerNameToVerify *string `json:"serverNameToVerify,omitempty" tf:"server_name_to_verify,omitempty"`
}

// IntegrationSpec defines the desired state of Integration
type IntegrationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IntegrationParameters `json:"forProvider"`
}

// IntegrationStatus defines the observed state of Integration.
type IntegrationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IntegrationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Integration is the Schema for the Integrations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Integration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IntegrationSpec   `json:"spec"`
	Status            IntegrationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IntegrationList contains a list of Integrations
type IntegrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Integration `json:"items"`
}

// Repository type metadata.
var (
	Integration_Kind             = "Integration"
	Integration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Integration_Kind}.String()
	Integration_KindAPIVersion   = Integration_Kind + "." + CRDGroupVersion.String()
	Integration_GroupVersionKind = CRDGroupVersion.WithKind(Integration_Kind)
)

func init() {
	SchemeBuilder.Register(&Integration{}, &IntegrationList{})
}

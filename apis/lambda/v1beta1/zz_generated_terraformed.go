// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"github.com/pkg/errors"

	"github.com/crossplane/upjet/pkg/resource"
	"github.com/crossplane/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this Alias
func (mg *Alias) GetTerraformResourceType() string {
	return "aws_lambda_alias"
}

// GetConnectionDetailsMapping for this Alias
func (tr *Alias) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Alias
func (tr *Alias) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Alias
func (tr *Alias) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Alias
func (tr *Alias) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Alias
func (tr *Alias) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Alias
func (tr *Alias) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this Alias
func (tr *Alias) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this Alias using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Alias) LateInitialize(attrs []byte) (bool, error) {
	params := &AliasParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Alias) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this CodeSigningConfig
func (mg *CodeSigningConfig) GetTerraformResourceType() string {
	return "aws_lambda_code_signing_config"
}

// GetConnectionDetailsMapping for this CodeSigningConfig
func (tr *CodeSigningConfig) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this CodeSigningConfig
func (tr *CodeSigningConfig) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this CodeSigningConfig
func (tr *CodeSigningConfig) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this CodeSigningConfig
func (tr *CodeSigningConfig) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this CodeSigningConfig
func (tr *CodeSigningConfig) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this CodeSigningConfig
func (tr *CodeSigningConfig) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this CodeSigningConfig
func (tr *CodeSigningConfig) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this CodeSigningConfig using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *CodeSigningConfig) LateInitialize(attrs []byte) (bool, error) {
	params := &CodeSigningConfigParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *CodeSigningConfig) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this EventSourceMapping
func (mg *EventSourceMapping) GetTerraformResourceType() string {
	return "aws_lambda_event_source_mapping"
}

// GetConnectionDetailsMapping for this EventSourceMapping
func (tr *EventSourceMapping) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this EventSourceMapping
func (tr *EventSourceMapping) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this EventSourceMapping
func (tr *EventSourceMapping) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this EventSourceMapping
func (tr *EventSourceMapping) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this EventSourceMapping
func (tr *EventSourceMapping) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this EventSourceMapping
func (tr *EventSourceMapping) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this EventSourceMapping
func (tr *EventSourceMapping) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this EventSourceMapping using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *EventSourceMapping) LateInitialize(attrs []byte) (bool, error) {
	params := &EventSourceMappingParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *EventSourceMapping) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this Function
func (mg *Function) GetTerraformResourceType() string {
	return "aws_lambda_function"
}

// GetConnectionDetailsMapping for this Function
func (tr *Function) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Function
func (tr *Function) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Function
func (tr *Function) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Function
func (tr *Function) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Function
func (tr *Function) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Function
func (tr *Function) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this Function
func (tr *Function) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this Function using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Function) LateInitialize(attrs []byte) (bool, error) {
	params := &FunctionParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Function) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this FunctionEventInvokeConfig
func (mg *FunctionEventInvokeConfig) GetTerraformResourceType() string {
	return "aws_lambda_function_event_invoke_config"
}

// GetConnectionDetailsMapping for this FunctionEventInvokeConfig
func (tr *FunctionEventInvokeConfig) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this FunctionEventInvokeConfig
func (tr *FunctionEventInvokeConfig) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this FunctionEventInvokeConfig
func (tr *FunctionEventInvokeConfig) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this FunctionEventInvokeConfig
func (tr *FunctionEventInvokeConfig) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this FunctionEventInvokeConfig
func (tr *FunctionEventInvokeConfig) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this FunctionEventInvokeConfig
func (tr *FunctionEventInvokeConfig) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this FunctionEventInvokeConfig
func (tr *FunctionEventInvokeConfig) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this FunctionEventInvokeConfig using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *FunctionEventInvokeConfig) LateInitialize(attrs []byte) (bool, error) {
	params := &FunctionEventInvokeConfigParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *FunctionEventInvokeConfig) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this FunctionURL
func (mg *FunctionURL) GetTerraformResourceType() string {
	return "aws_lambda_function_url"
}

// GetConnectionDetailsMapping for this FunctionURL
func (tr *FunctionURL) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this FunctionURL
func (tr *FunctionURL) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this FunctionURL
func (tr *FunctionURL) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this FunctionURL
func (tr *FunctionURL) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this FunctionURL
func (tr *FunctionURL) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this FunctionURL
func (tr *FunctionURL) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this FunctionURL
func (tr *FunctionURL) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this FunctionURL using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *FunctionURL) LateInitialize(attrs []byte) (bool, error) {
	params := &FunctionURLParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *FunctionURL) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this Invocation
func (mg *Invocation) GetTerraformResourceType() string {
	return "aws_lambda_invocation"
}

// GetConnectionDetailsMapping for this Invocation
func (tr *Invocation) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Invocation
func (tr *Invocation) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Invocation
func (tr *Invocation) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Invocation
func (tr *Invocation) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Invocation
func (tr *Invocation) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Invocation
func (tr *Invocation) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this Invocation
func (tr *Invocation) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this Invocation using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Invocation) LateInitialize(attrs []byte) (bool, error) {
	params := &InvocationParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Invocation) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this LayerVersion
func (mg *LayerVersion) GetTerraformResourceType() string {
	return "aws_lambda_layer_version"
}

// GetConnectionDetailsMapping for this LayerVersion
func (tr *LayerVersion) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this LayerVersion
func (tr *LayerVersion) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this LayerVersion
func (tr *LayerVersion) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this LayerVersion
func (tr *LayerVersion) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this LayerVersion
func (tr *LayerVersion) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this LayerVersion
func (tr *LayerVersion) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this LayerVersion
func (tr *LayerVersion) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this LayerVersion using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *LayerVersion) LateInitialize(attrs []byte) (bool, error) {
	params := &LayerVersionParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *LayerVersion) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this LayerVersionPermission
func (mg *LayerVersionPermission) GetTerraformResourceType() string {
	return "aws_lambda_layer_version_permission"
}

// GetConnectionDetailsMapping for this LayerVersionPermission
func (tr *LayerVersionPermission) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this LayerVersionPermission
func (tr *LayerVersionPermission) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this LayerVersionPermission
func (tr *LayerVersionPermission) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this LayerVersionPermission
func (tr *LayerVersionPermission) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this LayerVersionPermission
func (tr *LayerVersionPermission) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this LayerVersionPermission
func (tr *LayerVersionPermission) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this LayerVersionPermission
func (tr *LayerVersionPermission) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this LayerVersionPermission using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *LayerVersionPermission) LateInitialize(attrs []byte) (bool, error) {
	params := &LayerVersionPermissionParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *LayerVersionPermission) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this Permission
func (mg *Permission) GetTerraformResourceType() string {
	return "aws_lambda_permission"
}

// GetConnectionDetailsMapping for this Permission
func (tr *Permission) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Permission
func (tr *Permission) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Permission
func (tr *Permission) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Permission
func (tr *Permission) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Permission
func (tr *Permission) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Permission
func (tr *Permission) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this Permission
func (tr *Permission) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this Permission using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Permission) LateInitialize(attrs []byte) (bool, error) {
	params := &PermissionParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Permission) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this ProvisionedConcurrencyConfig
func (mg *ProvisionedConcurrencyConfig) GetTerraformResourceType() string {
	return "aws_lambda_provisioned_concurrency_config"
}

// GetConnectionDetailsMapping for this ProvisionedConcurrencyConfig
func (tr *ProvisionedConcurrencyConfig) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this ProvisionedConcurrencyConfig
func (tr *ProvisionedConcurrencyConfig) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ProvisionedConcurrencyConfig
func (tr *ProvisionedConcurrencyConfig) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ProvisionedConcurrencyConfig
func (tr *ProvisionedConcurrencyConfig) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ProvisionedConcurrencyConfig
func (tr *ProvisionedConcurrencyConfig) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ProvisionedConcurrencyConfig
func (tr *ProvisionedConcurrencyConfig) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this ProvisionedConcurrencyConfig
func (tr *ProvisionedConcurrencyConfig) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this ProvisionedConcurrencyConfig using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ProvisionedConcurrencyConfig) LateInitialize(attrs []byte) (bool, error) {
	params := &ProvisionedConcurrencyConfigParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}
	opts = append(opts, resource.WithNameFilter("ProvisionedConcurrentExecutions"))

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ProvisionedConcurrencyConfig) GetTerraformSchemaVersion() int {
	return 1
}

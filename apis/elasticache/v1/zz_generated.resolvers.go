/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-aws/apis/ec2/v1beta1"
	v1beta11 "github.com/upbound/provider-aws/apis/kms/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Cluster.
func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ParameterGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ParameterGroupNameRef,
		Selector:     mg.Spec.ForProvider.ParameterGroupNameSelector,
		To: reference.To{
			List:    &ParameterGroupList{},
			Managed: &ParameterGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ParameterGroupName")
	}
	mg.Spec.ForProvider.ParameterGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ParameterGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ReplicationGroupID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ReplicationGroupIDRef,
		Selector:     mg.Spec.ForProvider.ReplicationGroupIDSelector,
		To: reference.To{
			List:    &ReplicationGroupList{},
			Managed: &ReplicationGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ReplicationGroupID")
	}
	mg.Spec.ForProvider.ReplicationGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ReplicationGroupIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupIDRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupIDSelector,
		To: reference.To{
			List:    &v1beta1.SecurityGroupList{},
			Managed: &v1beta1.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupIds")
	}
	mg.Spec.ForProvider.SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupIDRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SubnetGroupNameRef,
		Selector:     mg.Spec.ForProvider.SubnetGroupNameSelector,
		To: reference.To{
			List:    &SubnetGroupList{},
			Managed: &SubnetGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetGroupName")
	}
	mg.Spec.ForProvider.SubnetGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ReplicationGroup.
func (mg *ReplicationGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KMSKeyIDRef,
		Selector:     mg.Spec.ForProvider.KMSKeyIDSelector,
		To: reference.To{
			List:    &v1beta11.KeyList{},
			Managed: &v1beta11.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyID")
	}
	mg.Spec.ForProvider.KMSKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupIDRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupIDSelector,
		To: reference.To{
			List:    &v1beta1.SecurityGroupList{},
			Managed: &v1beta1.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupIds")
	}
	mg.Spec.ForProvider.SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupIDRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SubnetGroupNameRef,
		Selector:     mg.Spec.ForProvider.SubnetGroupNameSelector,
		To: reference.To{
			List:    &SubnetGroupList{},
			Managed: &SubnetGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetGroupName")
	}
	mg.Spec.ForProvider.SubnetGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SubnetGroup.
func (mg *SubnetGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SubnetIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SubnetIDRefs,
		Selector:      mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &v1beta1.SubnetList{},
			Managed: &v1beta1.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetIds")
	}
	mg.Spec.ForProvider.SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SubnetIDRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this UserGroup.
func (mg *UserGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.UserIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.UserIDRefs,
		Selector:      mg.Spec.ForProvider.UserIDSelector,
		To: reference.To{
			List:    &UserList{},
			Managed: &User{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.UserIds")
	}
	mg.Spec.ForProvider.UserIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.UserIDRefs = mrsp.ResolvedReferences

	return nil
}

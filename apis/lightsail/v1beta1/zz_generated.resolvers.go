/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this DiskAttachment.
func (mg *DiskAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DiskName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DiskNameRef,
		Selector:     mg.Spec.ForProvider.DiskNameSelector,
		To: reference.To{
			List:    &DiskList{},
			Managed: &Disk{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DiskName")
	}
	mg.Spec.ForProvider.DiskName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DiskNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstanceNameRef,
		Selector:     mg.Spec.ForProvider.InstanceNameSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceName")
	}
	mg.Spec.ForProvider.InstanceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DomainEntry.
func (mg *DomainEntry) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DomainName),
		Extract:      resource.ExtractParamPath("domain_name", false),
		Reference:    mg.Spec.ForProvider.DomainNameRef,
		Selector:     mg.Spec.ForProvider.DomainNameSelector,
		To: reference.To{
			List:    &DomainList{},
			Managed: &Domain{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DomainName")
	}
	mg.Spec.ForProvider.DomainName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DomainNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this InstancePublicPorts.
func (mg *InstancePublicPorts) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstanceNameRef,
		Selector:     mg.Spec.ForProvider.InstanceNameSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceName")
	}
	mg.Spec.ForProvider.InstanceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LBAttachment.
func (mg *LBAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstanceNameRef,
		Selector:     mg.Spec.ForProvider.InstanceNameSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceName")
	}
	mg.Spec.ForProvider.InstanceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LBName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.LBNameRef,
		Selector:     mg.Spec.ForProvider.LBNameSelector,
		To: reference.To{
			List:    &LBList{},
			Managed: &LB{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LBName")
	}
	mg.Spec.ForProvider.LBName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LBNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LBCertificate.
func (mg *LBCertificate) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LBName),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.LBNameRef,
		Selector:     mg.Spec.ForProvider.LBNameSelector,
		To: reference.To{
			List:    &LBList{},
			Managed: &LB{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LBName")
	}
	mg.Spec.ForProvider.LBName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LBNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this StaticIPAttachment.
func (mg *StaticIPAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceName),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.InstanceNameRef,
		Selector:     mg.Spec.ForProvider.InstanceNameSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceName")
	}
	mg.Spec.ForProvider.InstanceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StaticIPName),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.StaticIPNameRef,
		Selector:     mg.Spec.ForProvider.StaticIPNameSelector,
		To: reference.To{
			List:    &StaticIPList{},
			Managed: &StaticIP{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StaticIPName")
	}
	mg.Spec.ForProvider.StaticIPName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StaticIPNameRef = rsp.ResolvedReference

	return nil
}

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
	v1beta1 "github.com/upbound/provider-aws/apis/ec2/v1beta1"
	common "github.com/upbound/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this AttachmentAccepter.
func (mg *AttachmentAccepter) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AttachmentID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.AttachmentIDRef,
		Selector:     mg.Spec.ForProvider.AttachmentIDSelector,
		To: reference.To{
			List:    &VPCAttachmentList{},
			Managed: &VPCAttachment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AttachmentID")
	}
	mg.Spec.ForProvider.AttachmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AttachmentIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AttachmentType),
		Extract:      resource.ExtractParamPath("attachment_type", true),
		Reference:    mg.Spec.ForProvider.AttachmentTypeRef,
		Selector:     mg.Spec.ForProvider.AttachmentTypeSelector,
		To: reference.To{
			List:    &VPCAttachmentList{},
			Managed: &VPCAttachment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AttachmentType")
	}
	mg.Spec.ForProvider.AttachmentType = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AttachmentTypeRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ConnectAttachment.
func (mg *ConnectAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CoreNetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CoreNetworkIDRef,
		Selector:     mg.Spec.ForProvider.CoreNetworkIDSelector,
		To: reference.To{
			List:    &CoreNetworkList{},
			Managed: &CoreNetwork{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CoreNetworkID")
	}
	mg.Spec.ForProvider.CoreNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CoreNetworkIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EdgeLocation),
		Extract:      resource.ExtractParamPath("edge_location", true),
		Reference:    mg.Spec.ForProvider.EdgeLocationRef,
		Selector:     mg.Spec.ForProvider.EdgeLocationSelector,
		To: reference.To{
			List:    &VPCAttachmentList{},
			Managed: &VPCAttachment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EdgeLocation")
	}
	mg.Spec.ForProvider.EdgeLocation = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EdgeLocationRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TransportAttachmentID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.TransportAttachmentIDRef,
		Selector:     mg.Spec.ForProvider.TransportAttachmentIDSelector,
		To: reference.To{
			List:    &VPCAttachmentList{},
			Managed: &VPCAttachment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TransportAttachmentID")
	}
	mg.Spec.ForProvider.TransportAttachmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TransportAttachmentIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Connection.
func (mg *Connection) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConnectedDeviceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ConnectedDeviceIDRef,
		Selector:     mg.Spec.ForProvider.ConnectedDeviceIDSelector,
		To: reference.To{
			List:    &DeviceList{},
			Managed: &Device{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ConnectedDeviceID")
	}
	mg.Spec.ForProvider.ConnectedDeviceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ConnectedDeviceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DeviceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.DeviceIDRef,
		Selector:     mg.Spec.ForProvider.DeviceIDSelector,
		To: reference.To{
			List:    &DeviceList{},
			Managed: &Device{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DeviceID")
	}
	mg.Spec.ForProvider.DeviceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DeviceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.GlobalNetworkID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.GlobalNetworkIDRef,
		Selector:     mg.Spec.ForProvider.GlobalNetworkIDSelector,
		To: reference.To{
			List:    &GlobalNetworkList{},
			Managed: &GlobalNetwork{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GlobalNetworkID")
	}
	mg.Spec.ForProvider.GlobalNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GlobalNetworkIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this CoreNetwork.
func (mg *CoreNetwork) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.GlobalNetworkID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.GlobalNetworkIDRef,
		Selector:     mg.Spec.ForProvider.GlobalNetworkIDSelector,
		To: reference.To{
			List:    &GlobalNetworkList{},
			Managed: &GlobalNetwork{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GlobalNetworkID")
	}
	mg.Spec.ForProvider.GlobalNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GlobalNetworkIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this CustomerGatewayAssociation.
func (mg *CustomerGatewayAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CustomerGatewayArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.CustomerGatewayArnRef,
		Selector:     mg.Spec.ForProvider.CustomerGatewayArnSelector,
		To: reference.To{
			List:    &v1beta1.CustomerGatewayList{},
			Managed: &v1beta1.CustomerGateway{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CustomerGatewayArn")
	}
	mg.Spec.ForProvider.CustomerGatewayArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CustomerGatewayArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DeviceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.DeviceIDRef,
		Selector:     mg.Spec.ForProvider.DeviceIDSelector,
		To: reference.To{
			List:    &DeviceList{},
			Managed: &Device{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DeviceID")
	}
	mg.Spec.ForProvider.DeviceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DeviceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.GlobalNetworkID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.GlobalNetworkIDRef,
		Selector:     mg.Spec.ForProvider.GlobalNetworkIDSelector,
		To: reference.To{
			List:    &GlobalNetworkList{},
			Managed: &GlobalNetwork{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GlobalNetworkID")
	}
	mg.Spec.ForProvider.GlobalNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GlobalNetworkIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Device.
func (mg *Device) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.GlobalNetworkID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.GlobalNetworkIDRef,
		Selector:     mg.Spec.ForProvider.GlobalNetworkIDSelector,
		To: reference.To{
			List:    &GlobalNetworkList{},
			Managed: &GlobalNetwork{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GlobalNetworkID")
	}
	mg.Spec.ForProvider.GlobalNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GlobalNetworkIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SiteID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SiteIDRef,
		Selector:     mg.Spec.ForProvider.SiteIDSelector,
		To: reference.To{
			List:    &SiteList{},
			Managed: &Site{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SiteID")
	}
	mg.Spec.ForProvider.SiteID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SiteIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Link.
func (mg *Link) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.GlobalNetworkID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.GlobalNetworkIDRef,
		Selector:     mg.Spec.ForProvider.GlobalNetworkIDSelector,
		To: reference.To{
			List:    &GlobalNetworkList{},
			Managed: &GlobalNetwork{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GlobalNetworkID")
	}
	mg.Spec.ForProvider.GlobalNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GlobalNetworkIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SiteID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SiteIDRef,
		Selector:     mg.Spec.ForProvider.SiteIDSelector,
		To: reference.To{
			List:    &SiteList{},
			Managed: &Site{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SiteID")
	}
	mg.Spec.ForProvider.SiteID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SiteIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LinkAssociation.
func (mg *LinkAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DeviceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DeviceIDRef,
		Selector:     mg.Spec.ForProvider.DeviceIDSelector,
		To: reference.To{
			List:    &DeviceList{},
			Managed: &Device{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DeviceID")
	}
	mg.Spec.ForProvider.DeviceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DeviceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.GlobalNetworkID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.GlobalNetworkIDRef,
		Selector:     mg.Spec.ForProvider.GlobalNetworkIDSelector,
		To: reference.To{
			List:    &GlobalNetworkList{},
			Managed: &GlobalNetwork{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GlobalNetworkID")
	}
	mg.Spec.ForProvider.GlobalNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GlobalNetworkIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LinkID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.LinkIDRef,
		Selector:     mg.Spec.ForProvider.LinkIDSelector,
		To: reference.To{
			List:    &LinkList{},
			Managed: &Link{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LinkID")
	}
	mg.Spec.ForProvider.LinkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LinkIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Site.
func (mg *Site) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.GlobalNetworkID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.GlobalNetworkIDRef,
		Selector:     mg.Spec.ForProvider.GlobalNetworkIDSelector,
		To: reference.To{
			List:    &GlobalNetworkList{},
			Managed: &GlobalNetwork{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GlobalNetworkID")
	}
	mg.Spec.ForProvider.GlobalNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GlobalNetworkIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TransitGatewayConnectPeerAssociation.
func (mg *TransitGatewayConnectPeerAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DeviceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.DeviceIDRef,
		Selector:     mg.Spec.ForProvider.DeviceIDSelector,
		To: reference.To{
			List:    &DeviceList{},
			Managed: &Device{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DeviceID")
	}
	mg.Spec.ForProvider.DeviceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DeviceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.GlobalNetworkID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.GlobalNetworkIDRef,
		Selector:     mg.Spec.ForProvider.GlobalNetworkIDSelector,
		To: reference.To{
			List:    &GlobalNetworkList{},
			Managed: &GlobalNetwork{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GlobalNetworkID")
	}
	mg.Spec.ForProvider.GlobalNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GlobalNetworkIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TransitGatewayConnectPeerArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.TransitGatewayConnectPeerArnRef,
		Selector:     mg.Spec.ForProvider.TransitGatewayConnectPeerArnSelector,
		To: reference.To{
			List:    &v1beta1.TransitGatewayConnectPeerList{},
			Managed: &v1beta1.TransitGatewayConnectPeer{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TransitGatewayConnectPeerArn")
	}
	mg.Spec.ForProvider.TransitGatewayConnectPeerArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TransitGatewayConnectPeerArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TransitGatewayRegistration.
func (mg *TransitGatewayRegistration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.GlobalNetworkID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.GlobalNetworkIDRef,
		Selector:     mg.Spec.ForProvider.GlobalNetworkIDSelector,
		To: reference.To{
			List:    &GlobalNetworkList{},
			Managed: &GlobalNetwork{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GlobalNetworkID")
	}
	mg.Spec.ForProvider.GlobalNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GlobalNetworkIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TransitGatewayArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.TransitGatewayArnRef,
		Selector:     mg.Spec.ForProvider.TransitGatewayArnSelector,
		To: reference.To{
			List:    &v1beta1.TransitGatewayList{},
			Managed: &v1beta1.TransitGateway{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TransitGatewayArn")
	}
	mg.Spec.ForProvider.TransitGatewayArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TransitGatewayArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VPCAttachment.
func (mg *VPCAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CoreNetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CoreNetworkIDRef,
		Selector:     mg.Spec.ForProvider.CoreNetworkIDSelector,
		To: reference.To{
			List:    &CoreNetworkList{},
			Managed: &CoreNetwork{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CoreNetworkID")
	}
	mg.Spec.ForProvider.CoreNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CoreNetworkIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SubnetArns),
		Extract:       common.ARNExtractor(),
		References:    mg.Spec.ForProvider.SubnetArnsRefs,
		Selector:      mg.Spec.ForProvider.SubnetArnsSelector,
		To: reference.To{
			List:    &v1beta1.SubnetList{},
			Managed: &v1beta1.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetArns")
	}
	mg.Spec.ForProvider.SubnetArns = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SubnetArnsRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.VPCArnRef,
		Selector:     mg.Spec.ForProvider.VPCArnSelector,
		To: reference.To{
			List:    &v1beta1.VPCList{},
			Managed: &v1beta1.VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCArn")
	}
	mg.Spec.ForProvider.VPCArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCArnRef = rsp.ResolvedReference

	return nil
}

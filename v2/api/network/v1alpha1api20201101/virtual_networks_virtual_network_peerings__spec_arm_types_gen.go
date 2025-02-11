// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

//Deprecated version of VirtualNetworksVirtualNetworkPeerings_Spec. Use v1beta20201101.VirtualNetworksVirtualNetworkPeerings_Spec instead
type VirtualNetworksVirtualNetworkPeerings_SpecARM struct {
	Location   *string                                   `json:"location,omitempty"`
	Name       string                                    `json:"name,omitempty"`
	Properties *VirtualNetworkPeeringPropertiesFormatARM `json:"properties,omitempty"`
	Tags       map[string]string                         `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &VirtualNetworksVirtualNetworkPeerings_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (peerings VirtualNetworksVirtualNetworkPeerings_SpecARM) GetAPIVersion() string {
	return "2020-11-01"
}

// GetName returns the Name of the resource
func (peerings VirtualNetworksVirtualNetworkPeerings_SpecARM) GetName() string {
	return peerings.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/virtualNetworks/virtualNetworkPeerings"
func (peerings VirtualNetworksVirtualNetworkPeerings_SpecARM) GetType() string {
	return "Microsoft.Network/virtualNetworks/virtualNetworkPeerings"
}

//Deprecated version of VirtualNetworkPeeringPropertiesFormat. Use v1beta20201101.VirtualNetworkPeeringPropertiesFormat instead
type VirtualNetworkPeeringPropertiesFormatARM struct {
	AllowForwardedTraffic     *bool                                              `json:"allowForwardedTraffic,omitempty"`
	AllowGatewayTransit       *bool                                              `json:"allowGatewayTransit,omitempty"`
	AllowVirtualNetworkAccess *bool                                              `json:"allowVirtualNetworkAccess,omitempty"`
	PeeringState              *VirtualNetworkPeeringPropertiesFormatPeeringState `json:"peeringState,omitempty"`
	RemoteAddressSpace        *AddressSpaceARM                                   `json:"remoteAddressSpace,omitempty"`
	RemoteBgpCommunities      *VirtualNetworkBgpCommunitiesARM                   `json:"remoteBgpCommunities,omitempty"`
	RemoteVirtualNetwork      *SubResourceARM                                    `json:"remoteVirtualNetwork,omitempty"`
	UseRemoteGateways         *bool                                              `json:"useRemoteGateways,omitempty"`
}

//Deprecated version of VirtualNetworkBgpCommunities. Use v1beta20201101.VirtualNetworkBgpCommunities instead
type VirtualNetworkBgpCommunitiesARM struct {
	VirtualNetworkCommunity *string `json:"virtualNetworkCommunity,omitempty"`
}

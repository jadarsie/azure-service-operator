// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type VirtualNetworkGateways_SPECARM struct {
	AzureName string `json:"azureName"`

	//ExtendedLocation: The extended location of type local virtual network gateway.
	ExtendedLocation *ExtendedLocation_SpecARM `json:"extendedLocation,omitempty"`
	Id               *string                   `json:"id,omitempty"`

	//Location: Resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name"`

	//Properties: Properties of the virtual network gateway.
	Properties VirtualNetworkGatewayPropertiesFormat_SpecARM `json:"properties"`

	//Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &VirtualNetworkGateways_SPECARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (specarm VirtualNetworkGateways_SPECARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (specarm VirtualNetworkGateways_SPECARM) GetName() string {
	return specarm.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (specarm VirtualNetworkGateways_SPECARM) GetType() string {
	return ""
}

type VirtualNetworkGatewayPropertiesFormat_SpecARM struct {
	//ActiveActive: ActiveActive flag.
	ActiveActive *bool `json:"activeActive,omitempty"`

	//BgpSettings: Virtual network gateway's BGP speaker settings.
	BgpSettings *BgpSettings_SpecARM `json:"bgpSettings,omitempty"`

	//CustomRoutes: The reference to the address space resource which represents the
	//custom routes address space specified by the customer for virtual network
	//gateway and VpnClient.
	CustomRoutes *AddressSpace_SpecARM `json:"customRoutes,omitempty"`

	//EnableBgp: Whether BGP is enabled for this virtual network gateway or not.
	EnableBgp *bool `json:"enableBgp,omitempty"`

	//EnableDnsForwarding: Whether dns forwarding is enabled or not.
	EnableDnsForwarding *bool `json:"enableDnsForwarding,omitempty"`

	//EnablePrivateIpAddress: Whether private IP needs to be enabled on this gateway
	//for connections or not.
	EnablePrivateIpAddress *bool `json:"enablePrivateIpAddress,omitempty"`

	//GatewayDefaultSite: The reference to the LocalNetworkGateway resource which
	//represents local network site having default routes. Assign Null value in case
	//of removing existing default site setting.
	GatewayDefaultSite *SubResource_SpecARM `json:"gatewayDefaultSite,omitempty"`

	//GatewayType: The type of this virtual network gateway.
	GatewayType *VirtualNetworkGatewayPropertiesFormat_GatewayType_Spec `json:"gatewayType,omitempty"`

	//IpConfigurations: IP configurations for virtual network gateway.
	IpConfigurations []VirtualNetworkGatewayIPConfiguration_SpecARM `json:"ipConfigurations,omitempty"`

	//Sku: The reference to the VirtualNetworkGatewaySku resource which represents the
	//SKU selected for Virtual network gateway.
	Sku                            *VirtualNetworkGatewaySku_SpecARM `json:"sku,omitempty"`
	VNetExtendedLocationResourceId *string                           `json:"vNetExtendedLocationResourceId,omitempty"`

	//VpnClientConfiguration: The reference to the VpnClientConfiguration resource
	//which represents the P2S VpnClient configurations.
	VpnClientConfiguration *VpnClientConfiguration_SpecARM `json:"vpnClientConfiguration,omitempty"`

	//VpnGatewayGeneration: The generation for this VirtualNetworkGateway. Must be
	//None if gatewayType is not VPN.
	VpnGatewayGeneration *VirtualNetworkGatewayPropertiesFormat_VpnGatewayGeneration_Spec `json:"vpnGatewayGeneration,omitempty"`

	//VpnType: The type of this virtual network gateway.
	VpnType *VirtualNetworkGatewayPropertiesFormat_VpnType_Spec `json:"vpnType,omitempty"`
}

type AddressSpace_SpecARM struct {
	//AddressPrefixes: A list of address blocks reserved for this virtual network in
	//CIDR notation.
	AddressPrefixes []string `json:"addressPrefixes,omitempty"`
}

type BgpSettings_SpecARM struct {
	//Asn: The BGP speaker's ASN.
	Asn *uint32 `json:"asn,omitempty"`

	//BgpPeeringAddress: The BGP peering address and BGP identifier of this BGP
	//speaker.
	BgpPeeringAddress *string `json:"bgpPeeringAddress,omitempty"`

	//BgpPeeringAddresses: BGP peering address with IP configuration ID for virtual
	//network gateway.
	BgpPeeringAddresses []IPConfigurationBgpPeeringAddress_SpecARM `json:"bgpPeeringAddresses,omitempty"`

	//PeerWeight: The weight added to routes learned from this BGP speaker.
	PeerWeight *int `json:"peerWeight,omitempty"`
}

type VirtualNetworkGatewayIPConfiguration_SpecARM struct {
	Id *string `json:"id,omitempty"`

	//Name: The name of the resource that is unique within a resource group. This name
	//can be used to access the resource.
	Name *string `json:"name,omitempty"`

	//Properties: Properties of the virtual network gateway ip configuration.
	Properties *VirtualNetworkGatewayIPConfigurationPropertiesFormat_SpecARM `json:"properties,omitempty"`
}

type VirtualNetworkGatewaySku_SpecARM struct {
	//Name: Gateway SKU name.
	Name *VirtualNetworkGatewaySku_Name_Spec `json:"name,omitempty"`

	//Tier: Gateway SKU tier.
	Tier *VirtualNetworkGatewaySku_Tier_Spec `json:"tier,omitempty"`
}

type VpnClientConfiguration_SpecARM struct {
	//AadAudience: The AADAudience property of the VirtualNetworkGateway resource for
	//vpn client connection used for AAD authentication.
	AadAudience *string `json:"aadAudience,omitempty"`

	//AadIssuer: The AADIssuer property of the VirtualNetworkGateway resource for vpn
	//client connection used for AAD authentication.
	AadIssuer *string `json:"aadIssuer,omitempty"`

	//AadTenant: The AADTenant property of the VirtualNetworkGateway resource for vpn
	//client connection used for AAD authentication.
	AadTenant *string `json:"aadTenant,omitempty"`

	//RadiusServerAddress: The radius server address property of the
	//VirtualNetworkGateway resource for vpn client connection.
	RadiusServerAddress *string `json:"radiusServerAddress,omitempty"`

	//RadiusServerSecret: The radius secret property of the VirtualNetworkGateway
	//resource for vpn client connection.
	RadiusServerSecret *string `json:"radiusServerSecret,omitempty"`

	//RadiusServers: The radiusServers property for multiple radius server
	//configuration.
	RadiusServers []RadiusServer_SpecARM `json:"radiusServers,omitempty"`

	//VpnAuthenticationTypes: VPN authentication types for the virtual network
	//gateway..
	VpnAuthenticationTypes []VpnClientConfiguration_VpnAuthenticationTypes_Spec `json:"vpnAuthenticationTypes,omitempty"`

	//VpnClientAddressPool: The reference to the address space resource which
	//represents Address space for P2S VpnClient.
	VpnClientAddressPool *AddressSpace_SpecARM `json:"vpnClientAddressPool,omitempty"`

	//VpnClientIpsecPolicies: VpnClientIpsecPolicies for virtual network gateway P2S
	//client.
	VpnClientIpsecPolicies []IpsecPolicy_SpecARM `json:"vpnClientIpsecPolicies,omitempty"`

	//VpnClientProtocols: VpnClientProtocols for Virtual network gateway.
	VpnClientProtocols []VpnClientConfiguration_VpnClientProtocols_Spec `json:"vpnClientProtocols,omitempty"`

	//VpnClientRevokedCertificates: VpnClientRevokedCertificate for Virtual network
	//gateway.
	VpnClientRevokedCertificates []VpnClientRevokedCertificate_SpecARM `json:"vpnClientRevokedCertificates,omitempty"`

	//VpnClientRootCertificates: VpnClientRootCertificate for virtual network gateway.
	VpnClientRootCertificates []VpnClientRootCertificate_SpecARM `json:"vpnClientRootCertificates,omitempty"`
}

type IPConfigurationBgpPeeringAddress_SpecARM struct {
	//CustomBgpIpAddresses: The list of custom BGP peering addresses which belong to
	//IP configuration.
	CustomBgpIpAddresses []string `json:"customBgpIpAddresses,omitempty"`

	//IpconfigurationId: The ID of IP configuration which belongs to gateway.
	IpconfigurationId *string `json:"ipconfigurationId,omitempty"`
}

type IpsecPolicy_SpecARM struct {
	//DhGroup: The DH Group used in IKE Phase 1 for initial SA.
	DhGroup DhGroup_Spec `json:"dhGroup"`

	//IkeEncryption: The IKE encryption algorithm (IKE phase 2).
	IkeEncryption IkeEncryption_Spec `json:"ikeEncryption"`

	//IkeIntegrity: The IKE integrity algorithm (IKE phase 2).
	IkeIntegrity IkeIntegrity_Spec `json:"ikeIntegrity"`

	//IpsecEncryption: The IPSec encryption algorithm (IKE phase 1).
	IpsecEncryption IpsecEncryption_Spec `json:"ipsecEncryption"`

	//IpsecIntegrity: The IPSec integrity algorithm (IKE phase 1).
	IpsecIntegrity IpsecIntegrity_Spec `json:"ipsecIntegrity"`

	//PfsGroup: The Pfs Group used in IKE Phase 2 for new child SA.
	PfsGroup PfsGroup_Spec `json:"pfsGroup"`

	//SaDataSizeKilobytes: The IPSec Security Association (also called Quick Mode or
	//Phase 2 SA) payload size in KB for a site to site VPN tunnel.
	SaDataSizeKilobytes int `json:"saDataSizeKilobytes"`

	//SaLifeTimeSeconds: The IPSec Security Association (also called Quick Mode or
	//Phase 2 SA) lifetime in seconds for a site to site VPN tunnel.
	SaLifeTimeSeconds int `json:"saLifeTimeSeconds"`
}

type RadiusServer_SpecARM struct {
	//RadiusServerAddress: The address of this radius server.
	RadiusServerAddress string `json:"radiusServerAddress"`

	//RadiusServerScore: The initial score assigned to this radius server.
	RadiusServerScore *int `json:"radiusServerScore,omitempty"`

	//RadiusServerSecret: The secret used for this radius server.
	RadiusServerSecret *string `json:"radiusServerSecret,omitempty"`
}

type VirtualNetworkGatewayIPConfigurationPropertiesFormat_SpecARM struct {
	//PrivateIPAllocationMethod: The private IP address allocation method.
	PrivateIPAllocationMethod *IPAllocationMethod_Spec `json:"privateIPAllocationMethod,omitempty"`

	//PublicIPAddress: The reference to the public IP resource.
	PublicIPAddress *SubResource_SpecARM `json:"publicIPAddress,omitempty"`

	//Subnet: The reference to the subnet resource.
	Subnet *SubResource_SpecARM `json:"subnet,omitempty"`
}

type VpnClientRevokedCertificate_SpecARM struct {
	Id *string `json:"id,omitempty"`

	//Name: The name of the resource that is unique within a resource group. This name
	//can be used to access the resource.
	Name *string `json:"name,omitempty"`

	//Properties: Properties of the vpn client revoked certificate.
	Properties *VpnClientRevokedCertificatePropertiesFormat_SpecARM `json:"properties,omitempty"`
}

type VpnClientRootCertificate_SpecARM struct {
	Id *string `json:"id,omitempty"`

	//Name: The name of the resource that is unique within a resource group. This name
	//can be used to access the resource.
	Name *string `json:"name,omitempty"`

	//Properties: Properties of the vpn client root certificate.
	Properties VpnClientRootCertificatePropertiesFormat_SpecARM `json:"properties"`
}

type VpnClientRevokedCertificatePropertiesFormat_SpecARM struct {
	//Thumbprint: The revoked VPN client certificate thumbprint.
	Thumbprint *string `json:"thumbprint,omitempty"`
}

type VpnClientRootCertificatePropertiesFormat_SpecARM struct {
	//PublicCertData: The certificate public data.
	PublicCertData string `json:"publicCertData"`
}

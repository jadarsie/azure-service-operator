// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

//Deprecated version of PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded. Use v1beta20201101.PublicIPAddress_Status_PublicIPAddress_SubResourceEmbedded instead
type PublicIPAddress_Status_PublicIPAddress_SubResourceEmbeddedARM struct {
	Etag             *string                                    `json:"etag,omitempty"`
	ExtendedLocation *ExtendedLocation_StatusARM                `json:"extendedLocation,omitempty"`
	Id               *string                                    `json:"id,omitempty"`
	Location         *string                                    `json:"location,omitempty"`
	Name             *string                                    `json:"name,omitempty"`
	Properties       *PublicIPAddressPropertiesFormat_StatusARM `json:"properties,omitempty"`
	Sku              *PublicIPAddressSku_StatusARM              `json:"sku,omitempty"`
	Tags             map[string]string                          `json:"tags,omitempty"`
	Type             *string                                    `json:"type,omitempty"`
	Zones            []string                                   `json:"zones,omitempty"`
}

//Deprecated version of PublicIPAddressPropertiesFormat_Status. Use v1beta20201101.PublicIPAddressPropertiesFormat_Status instead
type PublicIPAddressPropertiesFormat_StatusARM struct {
	DdosSettings             *DdosSettings_StatusARM                                        `json:"ddosSettings,omitempty"`
	DnsSettings              *PublicIPAddressDnsSettings_StatusARM                          `json:"dnsSettings,omitempty"`
	IdleTimeoutInMinutes     *int                                                           `json:"idleTimeoutInMinutes,omitempty"`
	IpAddress                *string                                                        `json:"ipAddress,omitempty"`
	IpConfiguration          *IPConfiguration_Status_PublicIPAddress_SubResourceEmbeddedARM `json:"ipConfiguration,omitempty"`
	IpTags                   []IpTag_StatusARM                                              `json:"ipTags,omitempty"`
	MigrationPhase           *PublicIPAddressPropertiesFormatStatusMigrationPhase           `json:"migrationPhase,omitempty"`
	NatGateway               *NatGateway_Status_PublicIPAddress_SubResourceEmbeddedARM      `json:"natGateway,omitempty"`
	ProvisioningState        *ProvisioningState_Status                                      `json:"provisioningState,omitempty"`
	PublicIPAddressVersion   *IPVersion_Status                                              `json:"publicIPAddressVersion,omitempty"`
	PublicIPAllocationMethod *IPAllocationMethod_Status                                     `json:"publicIPAllocationMethod,omitempty"`
	PublicIPPrefix           *SubResource_StatusARM                                         `json:"publicIPPrefix,omitempty"`
	ResourceGuid             *string                                                        `json:"resourceGuid,omitempty"`
}

//Deprecated version of PublicIPAddressSku_Status. Use v1beta20201101.PublicIPAddressSku_Status instead
type PublicIPAddressSku_StatusARM struct {
	Name *PublicIPAddressSkuStatusName `json:"name,omitempty"`
	Tier *PublicIPAddressSkuStatusTier `json:"tier,omitempty"`
}

//Deprecated version of DdosSettings_Status. Use v1beta20201101.DdosSettings_Status instead
type DdosSettings_StatusARM struct {
	DdosCustomPolicy   *SubResource_StatusARM                `json:"ddosCustomPolicy,omitempty"`
	ProtectedIP        *bool                                 `json:"protectedIP,omitempty"`
	ProtectionCoverage *DdosSettingsStatusProtectionCoverage `json:"protectionCoverage,omitempty"`
}

//Deprecated version of IPConfiguration_Status_PublicIPAddress_SubResourceEmbedded. Use v1beta20201101.IPConfiguration_Status_PublicIPAddress_SubResourceEmbedded instead
type IPConfiguration_Status_PublicIPAddress_SubResourceEmbeddedARM struct {
	Etag       *string                                                                        `json:"etag,omitempty"`
	Id         *string                                                                        `json:"id,omitempty"`
	Name       *string                                                                        `json:"name,omitempty"`
	Properties *IPConfigurationPropertiesFormat_Status_PublicIPAddress_SubResourceEmbeddedARM `json:"properties,omitempty"`
}

//Deprecated version of IpTag_Status. Use v1beta20201101.IpTag_Status instead
type IpTag_StatusARM struct {
	IpTagType *string `json:"ipTagType,omitempty"`
	Tag       *string `json:"tag,omitempty"`
}

//Deprecated version of NatGateway_Status_PublicIPAddress_SubResourceEmbedded. Use v1beta20201101.NatGateway_Status_PublicIPAddress_SubResourceEmbedded instead
type NatGateway_Status_PublicIPAddress_SubResourceEmbeddedARM struct {
	Id    *string                  `json:"id,omitempty"`
	Sku   *NatGatewaySku_StatusARM `json:"sku,omitempty"`
	Zones []string                 `json:"zones,omitempty"`
}

//Deprecated version of PublicIPAddressDnsSettings_Status. Use v1beta20201101.PublicIPAddressDnsSettings_Status instead
type PublicIPAddressDnsSettings_StatusARM struct {
	DomainNameLabel *string `json:"domainNameLabel,omitempty"`
	Fqdn            *string `json:"fqdn,omitempty"`
	ReverseFqdn     *string `json:"reverseFqdn,omitempty"`
}

//Deprecated version of PublicIPAddressSkuStatusName. Use v1beta20201101.PublicIPAddressSkuStatusName instead
type PublicIPAddressSkuStatusName string

const (
	PublicIPAddressSkuStatusNameBasic    = PublicIPAddressSkuStatusName("Basic")
	PublicIPAddressSkuStatusNameStandard = PublicIPAddressSkuStatusName("Standard")
)

//Deprecated version of PublicIPAddressSkuStatusTier. Use v1beta20201101.PublicIPAddressSkuStatusTier instead
type PublicIPAddressSkuStatusTier string

const (
	PublicIPAddressSkuStatusTierGlobal   = PublicIPAddressSkuStatusTier("Global")
	PublicIPAddressSkuStatusTierRegional = PublicIPAddressSkuStatusTier("Regional")
)

//Deprecated version of IPConfigurationPropertiesFormat_Status_PublicIPAddress_SubResourceEmbedded. Use v1beta20201101.IPConfigurationPropertiesFormat_Status_PublicIPAddress_SubResourceEmbedded instead
type IPConfigurationPropertiesFormat_Status_PublicIPAddress_SubResourceEmbeddedARM struct {
	PrivateIPAddress          *string                                               `json:"privateIPAddress,omitempty"`
	PrivateIPAllocationMethod *IPAllocationMethod_Status                            `json:"privateIPAllocationMethod,omitempty"`
	ProvisioningState         *ProvisioningState_Status                             `json:"provisioningState,omitempty"`
	Subnet                    *Subnet_Status_PublicIPAddress_SubResourceEmbeddedARM `json:"subnet,omitempty"`
}

//Deprecated version of NatGatewaySku_Status. Use v1beta20201101.NatGatewaySku_Status instead
type NatGatewaySku_StatusARM struct {
	Name *NatGatewaySkuStatusName `json:"name,omitempty"`
}

//Deprecated version of Subnet_Status_PublicIPAddress_SubResourceEmbedded. Use v1beta20201101.Subnet_Status_PublicIPAddress_SubResourceEmbedded instead
type Subnet_Status_PublicIPAddress_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`
}

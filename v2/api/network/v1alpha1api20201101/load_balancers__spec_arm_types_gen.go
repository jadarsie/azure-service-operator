// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

//Deprecated version of LoadBalancers_Spec. Use v1beta20201101.LoadBalancers_Spec instead
type LoadBalancers_SpecARM struct {
	ExtendedLocation *ExtendedLocationARM              `json:"extendedLocation,omitempty"`
	Location         *string                           `json:"location,omitempty"`
	Name             string                            `json:"name,omitempty"`
	Properties       *LoadBalancers_Spec_PropertiesARM `json:"properties,omitempty"`
	Sku              *LoadBalancerSkuARM               `json:"sku,omitempty"`
	Tags             map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &LoadBalancers_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (balancers LoadBalancers_SpecARM) GetAPIVersion() string {
	return "2020-11-01"
}

// GetName returns the Name of the resource
func (balancers LoadBalancers_SpecARM) GetName() string {
	return balancers.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/loadBalancers"
func (balancers LoadBalancers_SpecARM) GetType() string {
	return "Microsoft.Network/loadBalancers"
}

//Deprecated version of ExtendedLocation. Use v1beta20201101.ExtendedLocation instead
type ExtendedLocationARM struct {
	Name *string               `json:"name,omitempty"`
	Type *ExtendedLocationType `json:"type,omitempty"`
}

//Deprecated version of LoadBalancerSku. Use v1beta20201101.LoadBalancerSku instead
type LoadBalancerSkuARM struct {
	Name *LoadBalancerSkuName `json:"name,omitempty"`
	Tier *LoadBalancerSkuTier `json:"tier,omitempty"`
}

//Deprecated version of LoadBalancers_Spec_Properties. Use v1beta20201101.LoadBalancers_Spec_Properties instead
type LoadBalancers_Spec_PropertiesARM struct {
	BackendAddressPools      []LoadBalancers_Spec_Properties_BackendAddressPoolsARM      `json:"backendAddressPools,omitempty"`
	FrontendIPConfigurations []LoadBalancers_Spec_Properties_FrontendIPConfigurationsARM `json:"frontendIPConfigurations,omitempty"`
	InboundNatPools          []LoadBalancers_Spec_Properties_InboundNatPoolsARM          `json:"inboundNatPools,omitempty"`
	LoadBalancingRules       []LoadBalancers_Spec_Properties_LoadBalancingRulesARM       `json:"loadBalancingRules,omitempty"`
	OutboundRules            []LoadBalancers_Spec_Properties_OutboundRulesARM            `json:"outboundRules,omitempty"`
	Probes                   []LoadBalancers_Spec_Properties_ProbesARM                   `json:"probes,omitempty"`
}

//Deprecated version of ExtendedLocationType. Use v1beta20201101.ExtendedLocationType instead
// +kubebuilder:validation:Enum={"EdgeZone"}
type ExtendedLocationType string

const ExtendedLocationTypeEdgeZone = ExtendedLocationType("EdgeZone")

//Deprecated version of LoadBalancerSkuName. Use v1beta20201101.LoadBalancerSkuName instead
// +kubebuilder:validation:Enum={"Basic","Standard"}
type LoadBalancerSkuName string

const (
	LoadBalancerSkuNameBasic    = LoadBalancerSkuName("Basic")
	LoadBalancerSkuNameStandard = LoadBalancerSkuName("Standard")
)

//Deprecated version of LoadBalancerSkuTier. Use v1beta20201101.LoadBalancerSkuTier instead
// +kubebuilder:validation:Enum={"Global","Regional"}
type LoadBalancerSkuTier string

const (
	LoadBalancerSkuTierGlobal   = LoadBalancerSkuTier("Global")
	LoadBalancerSkuTierRegional = LoadBalancerSkuTier("Regional")
)

//Deprecated version of LoadBalancers_Spec_Properties_BackendAddressPools. Use v1beta20201101.LoadBalancers_Spec_Properties_BackendAddressPools instead
type LoadBalancers_Spec_Properties_BackendAddressPoolsARM struct {
	Name       *string                                                          `json:"name,omitempty"`
	Properties *LoadBalancers_Spec_Properties_BackendAddressPools_PropertiesARM `json:"properties,omitempty"`
}

//Deprecated version of LoadBalancers_Spec_Properties_FrontendIPConfigurations. Use v1beta20201101.LoadBalancers_Spec_Properties_FrontendIPConfigurations instead
type LoadBalancers_Spec_Properties_FrontendIPConfigurationsARM struct {
	Name       *string                                     `json:"name,omitempty"`
	Properties *FrontendIPConfigurationPropertiesFormatARM `json:"properties,omitempty"`
	Zones      []string                                    `json:"zones,omitempty"`
}

//Deprecated version of LoadBalancers_Spec_Properties_InboundNatPools. Use v1beta20201101.LoadBalancers_Spec_Properties_InboundNatPools instead
type LoadBalancers_Spec_Properties_InboundNatPoolsARM struct {
	Name       *string                            `json:"name,omitempty"`
	Properties *InboundNatPoolPropertiesFormatARM `json:"properties,omitempty"`
}

//Deprecated version of LoadBalancers_Spec_Properties_LoadBalancingRules. Use v1beta20201101.LoadBalancers_Spec_Properties_LoadBalancingRules instead
type LoadBalancers_Spec_Properties_LoadBalancingRulesARM struct {
	Name       *string                               `json:"name,omitempty"`
	Properties *LoadBalancingRulePropertiesFormatARM `json:"properties,omitempty"`
}

//Deprecated version of LoadBalancers_Spec_Properties_OutboundRules. Use v1beta20201101.LoadBalancers_Spec_Properties_OutboundRules instead
type LoadBalancers_Spec_Properties_OutboundRulesARM struct {
	Name       *string                          `json:"name,omitempty"`
	Properties *OutboundRulePropertiesFormatARM `json:"properties,omitempty"`
}

//Deprecated version of LoadBalancers_Spec_Properties_Probes. Use v1beta20201101.LoadBalancers_Spec_Properties_Probes instead
type LoadBalancers_Spec_Properties_ProbesARM struct {
	Name       *string                   `json:"name,omitempty"`
	Properties *ProbePropertiesFormatARM `json:"properties,omitempty"`
}

//Deprecated version of FrontendIPConfigurationPropertiesFormat. Use v1beta20201101.FrontendIPConfigurationPropertiesFormat instead
type FrontendIPConfigurationPropertiesFormatARM struct {
	PrivateIPAddress          *string                                                           `json:"privateIPAddress,omitempty"`
	PrivateIPAddressVersion   *FrontendIPConfigurationPropertiesFormatPrivateIPAddressVersion   `json:"privateIPAddressVersion,omitempty"`
	PrivateIPAllocationMethod *FrontendIPConfigurationPropertiesFormatPrivateIPAllocationMethod `json:"privateIPAllocationMethod,omitempty"`
	PublicIPAddress           *SubResourceARM                                                   `json:"publicIPAddress,omitempty"`
	PublicIPPrefix            *SubResourceARM                                                   `json:"publicIPPrefix,omitempty"`
	Subnet                    *SubResourceARM                                                   `json:"subnet,omitempty"`
}

//Deprecated version of InboundNatPoolPropertiesFormat. Use v1beta20201101.InboundNatPoolPropertiesFormat instead
type InboundNatPoolPropertiesFormatARM struct {
	BackendPort             *int                                    `json:"backendPort,omitempty"`
	EnableFloatingIP        *bool                                   `json:"enableFloatingIP,omitempty"`
	EnableTcpReset          *bool                                   `json:"enableTcpReset,omitempty"`
	FrontendIPConfiguration *SubResourceARM                         `json:"frontendIPConfiguration,omitempty"`
	FrontendPortRangeEnd    *int                                    `json:"frontendPortRangeEnd,omitempty"`
	FrontendPortRangeStart  *int                                    `json:"frontendPortRangeStart,omitempty"`
	IdleTimeoutInMinutes    *int                                    `json:"idleTimeoutInMinutes,omitempty"`
	Protocol                *InboundNatPoolPropertiesFormatProtocol `json:"protocol,omitempty"`
}

//Deprecated version of LoadBalancers_Spec_Properties_BackendAddressPools_Properties. Use v1beta20201101.LoadBalancers_Spec_Properties_BackendAddressPools_Properties instead
type LoadBalancers_Spec_Properties_BackendAddressPools_PropertiesARM struct {
	LoadBalancerBackendAddresses []LoadBalancers_Spec_Properties_BackendAddressPools_Properties_LoadBalancerBackendAddressesARM `json:"loadBalancerBackendAddresses,omitempty"`
	Location                     *string                                                                                        `json:"location,omitempty"`
}

//Deprecated version of LoadBalancingRulePropertiesFormat. Use v1beta20201101.LoadBalancingRulePropertiesFormat instead
type LoadBalancingRulePropertiesFormatARM struct {
	BackendAddressPool      *SubResourceARM                                    `json:"backendAddressPool,omitempty"`
	BackendPort             *int                                               `json:"backendPort,omitempty"`
	DisableOutboundSnat     *bool                                              `json:"disableOutboundSnat,omitempty"`
	EnableFloatingIP        *bool                                              `json:"enableFloatingIP,omitempty"`
	EnableTcpReset          *bool                                              `json:"enableTcpReset,omitempty"`
	FrontendIPConfiguration *SubResourceARM                                    `json:"frontendIPConfiguration,omitempty"`
	FrontendPort            *int                                               `json:"frontendPort,omitempty"`
	IdleTimeoutInMinutes    *int                                               `json:"idleTimeoutInMinutes,omitempty"`
	LoadDistribution        *LoadBalancingRulePropertiesFormatLoadDistribution `json:"loadDistribution,omitempty"`
	Probe                   *SubResourceARM                                    `json:"probe,omitempty"`
	Protocol                *LoadBalancingRulePropertiesFormatProtocol         `json:"protocol,omitempty"`
}

//Deprecated version of OutboundRulePropertiesFormat. Use v1beta20201101.OutboundRulePropertiesFormat instead
type OutboundRulePropertiesFormatARM struct {
	AllocatedOutboundPorts   *int                                  `json:"allocatedOutboundPorts,omitempty"`
	BackendAddressPool       *SubResourceARM                       `json:"backendAddressPool,omitempty"`
	EnableTcpReset           *bool                                 `json:"enableTcpReset,omitempty"`
	FrontendIPConfigurations []SubResourceARM                      `json:"frontendIPConfigurations,omitempty"`
	IdleTimeoutInMinutes     *int                                  `json:"idleTimeoutInMinutes,omitempty"`
	Protocol                 *OutboundRulePropertiesFormatProtocol `json:"protocol,omitempty"`
}

//Deprecated version of ProbePropertiesFormat. Use v1beta20201101.ProbePropertiesFormat instead
type ProbePropertiesFormatARM struct {
	IntervalInSeconds *int                           `json:"intervalInSeconds,omitempty"`
	NumberOfProbes    *int                           `json:"numberOfProbes,omitempty"`
	Port              *int                           `json:"port,omitempty"`
	Protocol          *ProbePropertiesFormatProtocol `json:"protocol,omitempty"`
	RequestPath       *string                        `json:"requestPath,omitempty"`
}

//Deprecated version of LoadBalancers_Spec_Properties_BackendAddressPools_Properties_LoadBalancerBackendAddresses. Use v1beta20201101.LoadBalancers_Spec_Properties_BackendAddressPools_Properties_LoadBalancerBackendAddresses instead
type LoadBalancers_Spec_Properties_BackendAddressPools_Properties_LoadBalancerBackendAddressesARM struct {
	Name       *string                                        `json:"name,omitempty"`
	Properties *LoadBalancerBackendAddressPropertiesFormatARM `json:"properties,omitempty"`
}

//Deprecated version of LoadBalancerBackendAddressPropertiesFormat. Use v1beta20201101.LoadBalancerBackendAddressPropertiesFormat instead
type LoadBalancerBackendAddressPropertiesFormatARM struct {
	IpAddress                           *string         `json:"ipAddress,omitempty"`
	LoadBalancerFrontendIPConfiguration *SubResourceARM `json:"loadBalancerFrontendIPConfiguration,omitempty"`
	Subnet                              *SubResourceARM `json:"subnet,omitempty"`
	VirtualNetwork                      *SubResourceARM `json:"virtualNetwork,omitempty"`
}

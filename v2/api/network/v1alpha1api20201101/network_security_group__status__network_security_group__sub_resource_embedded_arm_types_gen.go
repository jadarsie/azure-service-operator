// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

type NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbeddedARM struct {
	//Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	//Id: Resource ID.
	Id *string `json:"id,omitempty"`

	//Location: Resource location.
	Location *string `json:"location,omitempty"`

	//Name: Resource name.
	Name *string `json:"name,omitempty"`

	//Properties: Properties of the network security group.
	Properties *NetworkSecurityGroupPropertiesFormat_StatusARM `json:"properties,omitempty"`

	//Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	//Type: Resource type.
	Type *string `json:"type,omitempty"`
}

type NetworkSecurityGroupPropertiesFormat_StatusARM struct {
	//DefaultSecurityRules: The default security rules of network security group.
	DefaultSecurityRules []SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbeddedARM `json:"defaultSecurityRules,omitempty"`

	//FlowLogs: A collection of references to flow log resources.
	FlowLogs []FlowLog_Status_SubResourceEmbeddedARM `json:"flowLogs,omitempty"`

	//NetworkInterfaces: A collection of references to network interfaces.
	NetworkInterfaces []NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbeddedARM `json:"networkInterfaces,omitempty"`

	//ProvisioningState: The provisioning state of the network security group resource.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	//ResourceGuid: The resource GUID property of the network security group resource.
	ResourceGuid *string `json:"resourceGuid,omitempty"`

	//SecurityRules: A collection of security rules of the network security group.
	SecurityRules []SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbeddedARM `json:"securityRules,omitempty"`

	//Subnets: A collection of references to subnets.
	Subnets []Subnet_Status_NetworkSecurityGroup_SubResourceEmbeddedARM `json:"subnets,omitempty"`
}

type FlowLog_Status_SubResourceEmbeddedARM struct {
	//Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

type NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbeddedARM struct {
	//ExtendedLocation: The extended location of the network interface.
	ExtendedLocation *ExtendedLocation_StatusARM `json:"extendedLocation,omitempty"`

	//Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

type SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbeddedARM struct {
	//Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

type Subnet_Status_NetworkSecurityGroup_SubResourceEmbeddedARM struct {
	//Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

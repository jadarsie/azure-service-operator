// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210401

type QueueServiceProperties_StatusARM struct {
	//Id: Fully qualified resource ID for the resource. Ex -
	///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	//Name: The name of the resource
	Name *string `json:"name,omitempty"`

	//Properties: The properties of a storage account’s Queue service.
	Properties *QueueServiceProperties_Status_PropertiesARM `json:"properties,omitempty"`

	//Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type QueueServiceProperties_Status_PropertiesARM struct {
	//Cors: Specifies CORS rules for the Queue service. You can include up to five CorsRule elements in the request. If no
	//CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the
	//Queue service.
	Cors *CorsRules_StatusARM `json:"cors,omitempty"`
}

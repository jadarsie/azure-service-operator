// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

type AuthorizationRule_StatusARM struct {
	//Id: Fully qualified resource ID for the resource. Ex -
	///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	//Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	//Name: The name of the resource
	Name *string `json:"name,omitempty"`

	//Properties: Properties supplied to create or update AuthorizationRule
	Properties *AuthorizationRule_Properties_StatusARM `json:"properties,omitempty"`

	//SystemData: The system meta data relating to this resource.
	SystemData *SystemData_StatusARM `json:"systemData,omitempty"`

	//Type: The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or
	//"Microsoft.EventHub/Namespaces/EventHubs"
	Type *string `json:"type,omitempty"`
}

type AuthorizationRule_Properties_StatusARM struct {
	//Rights: The rights associated with the rule.
	Rights []AuthorizationRule_Properties_Rights_Status `json:"rights,omitempty"`
}

type SystemData_StatusARM struct {
	//CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	//CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	//CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemData_CreatedByType_Status `json:"createdByType,omitempty"`

	//LastModifiedAt: The type of identity that last modified the resource.
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	//LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	//LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemData_LastModifiedByType_Status `json:"lastModifiedByType,omitempty"`
}

type AuthorizationRule_Properties_Rights_Status string

const (
	AuthorizationRule_Properties_Rights_StatusListen = AuthorizationRule_Properties_Rights_Status("Listen")
	AuthorizationRule_Properties_Rights_StatusManage = AuthorizationRule_Properties_Rights_Status("Manage")
	AuthorizationRule_Properties_Rights_StatusSend   = AuthorizationRule_Properties_Rights_Status("Send")
)

type SystemData_CreatedByType_Status string

const (
	SystemData_CreatedByType_StatusApplication     = SystemData_CreatedByType_Status("Application")
	SystemData_CreatedByType_StatusKey             = SystemData_CreatedByType_Status("Key")
	SystemData_CreatedByType_StatusManagedIdentity = SystemData_CreatedByType_Status("ManagedIdentity")
	SystemData_CreatedByType_StatusUser            = SystemData_CreatedByType_Status("User")
)

type SystemData_LastModifiedByType_Status string

const (
	SystemData_LastModifiedByType_StatusApplication     = SystemData_LastModifiedByType_Status("Application")
	SystemData_LastModifiedByType_StatusKey             = SystemData_LastModifiedByType_Status("Key")
	SystemData_LastModifiedByType_StatusManagedIdentity = SystemData_LastModifiedByType_Status("ManagedIdentity")
	SystemData_LastModifiedByType_StatusUser            = SystemData_LastModifiedByType_Status("User")
)

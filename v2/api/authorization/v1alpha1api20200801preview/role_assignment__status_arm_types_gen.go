// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200801preview

type RoleAssignment_StatusARM struct {
	//Properties: Role assignment properties.
	Properties *RoleAssignmentProperties_StatusARM `json:"properties,omitempty"`
}

type RoleAssignmentProperties_StatusARM struct {
	//Condition: The conditions on the role assignment. This limits the resources it
	//can be assigned to. e.g.:
	//@Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName]
	//StringEqualsIgnoreCase 'foo_storage_container'
	Condition *string `json:"condition,omitempty"`

	//ConditionVersion: Version of the condition. Currently accepted value is '2.0'
	ConditionVersion *string `json:"conditionVersion,omitempty"`

	//CreatedBy: Id of the user who created the assignment
	CreatedBy *string `json:"createdBy,omitempty"`

	//CreatedOn: Time it was created
	CreatedOn *string `json:"createdOn,omitempty"`

	//DelegatedManagedIdentityResourceId: Id of the delegated managed identity resource
	DelegatedManagedIdentityResourceId *string `json:"delegatedManagedIdentityResourceId,omitempty"`

	//Description: Description of role assignment
	Description *string `json:"description,omitempty"`

	//PrincipalId: The principal ID.
	PrincipalId string `json:"principalId"`

	//PrincipalType: The principal type of the assigned principal ID.
	PrincipalType *RoleAssignmentProperties_PrincipalType_Status `json:"principalType,omitempty"`

	//RoleDefinitionId: The role definition ID.
	RoleDefinitionId string `json:"roleDefinitionId"`

	//Scope: The role assignment scope.
	Scope *string `json:"scope,omitempty"`

	//UpdatedBy: Id of the user who updated the assignment
	UpdatedBy *string `json:"updatedBy,omitempty"`

	//UpdatedOn: Time it was updated
	UpdatedOn *string `json:"updatedOn,omitempty"`
}

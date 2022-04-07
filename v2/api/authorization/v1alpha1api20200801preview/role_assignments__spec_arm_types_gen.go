// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200801preview

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

//Deprecated version of RoleAssignments_Spec. Use v1beta20200801preview.RoleAssignments_Spec instead
type RoleAssignments_SpecARM struct {
	Location   *string                      `json:"location,omitempty"`
	Name       string                       `json:"name,omitempty"`
	Properties *RoleAssignmentPropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string            `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &RoleAssignments_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-08-01-preview"
func (assignments RoleAssignments_SpecARM) GetAPIVersion() string {
	return "2020-08-01-preview"
}

// GetName returns the Name of the resource
func (assignments RoleAssignments_SpecARM) GetName() string {
	return assignments.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Authorization/roleAssignments"
func (assignments RoleAssignments_SpecARM) GetType() string {
	return "Microsoft.Authorization/roleAssignments"
}

//Deprecated version of RoleAssignmentProperties. Use v1beta20200801preview.RoleAssignmentProperties instead
type RoleAssignmentPropertiesARM struct {
	Condition                          *string                                `json:"condition,omitempty"`
	ConditionVersion                   *string                                `json:"conditionVersion,omitempty"`
	DelegatedManagedIdentityResourceId *string                                `json:"delegatedManagedIdentityResourceId,omitempty"`
	Description                        *string                                `json:"description,omitempty"`
	PrincipalId                        *string                                `json:"principalId,omitempty"`
	PrincipalType                      *RoleAssignmentPropertiesPrincipalType `json:"principalType,omitempty"`
	RoleDefinitionId                   *string                                `json:"roleDefinitionId,omitempty"`
}

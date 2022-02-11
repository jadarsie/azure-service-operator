// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type NamespacesEventhubsAuthorizationRules_SPECARM struct {
	AzureName string `json:"azureName"`
	Name      string `json:"name"`

	//Properties: Properties supplied to create or update AuthorizationRule
	Properties *NamespacesEventhubsAuthorizationRules_Properties_SPECARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NamespacesEventhubsAuthorizationRules_SPECARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (specarm NamespacesEventhubsAuthorizationRules_SPECARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (specarm NamespacesEventhubsAuthorizationRules_SPECARM) GetName() string {
	return specarm.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (specarm NamespacesEventhubsAuthorizationRules_SPECARM) GetType() string {
	return ""
}

type NamespacesEventhubsAuthorizationRules_Properties_SPECARM struct {
	//Rights: The rights associated with the rule.
	Rights []NamespacesEventhubsAuthorizationRules_Properties_Rights_SPEC `json:"rights"`
}

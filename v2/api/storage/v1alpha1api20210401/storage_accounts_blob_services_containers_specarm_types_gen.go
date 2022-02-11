// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type StorageAccountsBlobServicesContainers_SPECARM struct {
	AzureName string `json:"azureName"`
	Name      string `json:"name"`

	//Properties: Properties of the blob container.
	Properties *ContainerProperties_SpecARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &StorageAccountsBlobServicesContainers_SPECARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (specarm StorageAccountsBlobServicesContainers_SPECARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (specarm StorageAccountsBlobServicesContainers_SPECARM) GetName() string {
	return specarm.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (specarm StorageAccountsBlobServicesContainers_SPECARM) GetType() string {
	return ""
}

type ContainerProperties_SpecARM struct {
	//DefaultEncryptionScope: Default the container to use specified encryption scope
	//for all writes.
	DefaultEncryptionScope *string `json:"defaultEncryptionScope,omitempty"`

	//DenyEncryptionScopeOverride: Block override of encryption scope from the
	//container default.
	DenyEncryptionScopeOverride *bool `json:"denyEncryptionScopeOverride,omitempty"`

	//ImmutableStorageWithVersioning: The object level immutability property of the
	//container. The property is immutable and can only be set to true at the
	//container creation time. Existing containers must undergo a migration process.
	ImmutableStorageWithVersioning *ImmutableStorageWithVersioning_SpecARM `json:"immutableStorageWithVersioning,omitempty"`

	//Metadata: A name-value pair to associate with the container as metadata.
	Metadata map[string]string `json:"metadata,omitempty"`

	//PublicAccess: Specifies whether data in the container may be accessed publicly
	//and the level of access.
	PublicAccess *ContainerProperties_PublicAccess_Spec `json:"publicAccess,omitempty"`
}

type ImmutableStorageWithVersioning_SpecARM struct {
	//Enabled: This is an immutable property, when set to true it enables object level
	//immutability at the container level.
	Enabled *bool `json:"enabled,omitempty"`
}

// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

//Deprecated version of StorageAccountsBlobServicesContainers_Spec. Use v1beta20210401.StorageAccountsBlobServicesContainers_Spec instead
type StorageAccountsBlobServicesContainers_SpecARM struct {
	Location   *string                 `json:"location,omitempty"`
	Name       string                  `json:"name,omitempty"`
	Properties *ContainerPropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string       `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &StorageAccountsBlobServicesContainers_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (containers StorageAccountsBlobServicesContainers_SpecARM) GetAPIVersion() string {
	return "2021-04-01"
}

// GetName returns the Name of the resource
func (containers StorageAccountsBlobServicesContainers_SpecARM) GetName() string {
	return containers.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/blobServices/containers"
func (containers StorageAccountsBlobServicesContainers_SpecARM) GetType() string {
	return "Microsoft.Storage/storageAccounts/blobServices/containers"
}

//Deprecated version of ContainerProperties. Use v1beta20210401.ContainerProperties instead
type ContainerPropertiesARM struct {
	DefaultEncryptionScope         *string                            `json:"defaultEncryptionScope,omitempty"`
	DenyEncryptionScopeOverride    *bool                              `json:"denyEncryptionScopeOverride,omitempty"`
	ImmutableStorageWithVersioning *ImmutableStorageWithVersioningARM `json:"immutableStorageWithVersioning,omitempty"`
	Metadata                       map[string]string                  `json:"metadata,omitempty"`
	PublicAccess                   *ContainerPropertiesPublicAccess   `json:"publicAccess,omitempty"`
}

//Deprecated version of ImmutableStorageWithVersioning. Use v1beta20210401.ImmutableStorageWithVersioning instead
type ImmutableStorageWithVersioningARM struct {
	Enabled *bool `json:"enabled,omitempty"`
}

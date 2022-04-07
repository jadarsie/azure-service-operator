// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	networkv1alpha1api20201101 "github.com/Azure/azure-service-operator/v2/api/network/v1alpha1api20201101"
	"github.com/Azure/azure-service-operator/v2/api/network/v1alpha1api20201101storage"
	networkv1beta20201101 "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101"
	"github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type PublicIPAddressExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *PublicIPAddressExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&networkv1alpha1api20201101.PublicIPAddress{},
		&v1alpha1api20201101storage.PublicIPAddress{},
		&networkv1beta20201101.PublicIPAddress{},
		&v1beta20201101storage.PublicIPAddress{}}
}

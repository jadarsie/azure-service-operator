// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	eventgridv1alpha1api20200601 "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1alpha1api20200601"
	"github.com/Azure/azure-service-operator/v2/api/eventgrid/v1alpha1api20200601storage"
	eventgridv1beta20200601 "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1beta20200601"
	"github.com/Azure/azure-service-operator/v2/api/eventgrid/v1beta20200601storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type TopicExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *TopicExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&eventgridv1alpha1api20200601.Topic{},
		&v1alpha1api20200601storage.Topic{},
		&eventgridv1beta20200601.Topic{},
		&v1beta20200601storage.Topic{}}
}

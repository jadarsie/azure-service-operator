// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	documentdbv1alpha1api20210515 "github.com/Azure/azure-service-operator/v2/api/documentdb/v1alpha1api20210515"
	"github.com/Azure/azure-service-operator/v2/api/documentdb/v1alpha1api20210515storage"
	documentdbv1beta20210515 "github.com/Azure/azure-service-operator/v2/api/documentdb/v1beta20210515"
	"github.com/Azure/azure-service-operator/v2/api/documentdb/v1beta20210515storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type MongodbDatabaseCollectionThroughputSettingExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *MongodbDatabaseCollectionThroughputSettingExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&documentdbv1alpha1api20210515.MongodbDatabaseCollectionThroughputSetting{},
		&v1alpha1api20210515storage.MongodbDatabaseCollectionThroughputSetting{},
		&documentdbv1beta20210515.MongodbDatabaseCollectionThroughputSetting{},
		&v1beta20210515storage.MongodbDatabaseCollectionThroughputSetting{}}
}

// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

//Deprecated version of Topics_Spec. Use v1beta20200601.Topics_Spec instead
type Topics_SpecARM struct {
	Location *string           `json:"location,omitempty"`
	Name     string            `json:"name,omitempty"`
	Tags     map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Topics_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (topics Topics_SpecARM) GetAPIVersion() string {
	return "2020-06-01"
}

// GetName returns the Name of the resource
func (topics Topics_SpecARM) GetName() string {
	return topics.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventGrid/topics"
func (topics Topics_SpecARM) GetType() string {
	return "Microsoft.EventGrid/topics"
}

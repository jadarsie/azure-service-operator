// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201storage

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/api/cache/v1beta20201201storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20201201.RedisLinkedServer
//Deprecated version of RedisLinkedServer. Use v1beta20201201.RedisLinkedServer instead
type RedisLinkedServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisLinkedServers_Spec                `json:"spec,omitempty"`
	Status            RedisLinkedServerWithProperties_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RedisLinkedServer{}

// GetConditions returns the conditions of the resource
func (server *RedisLinkedServer) GetConditions() conditions.Conditions {
	return server.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (server *RedisLinkedServer) SetConditions(conditions conditions.Conditions) {
	server.Status.Conditions = conditions
}

var _ conversion.Convertible = &RedisLinkedServer{}

// ConvertFrom populates our RedisLinkedServer from the provided hub RedisLinkedServer
func (server *RedisLinkedServer) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1beta20201201storage.RedisLinkedServer)
	if !ok {
		return fmt.Errorf("expected cache/v1beta20201201storage/RedisLinkedServer but received %T instead", hub)
	}

	return server.AssignPropertiesFromRedisLinkedServer(source)
}

// ConvertTo populates the provided hub RedisLinkedServer from our RedisLinkedServer
func (server *RedisLinkedServer) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1beta20201201storage.RedisLinkedServer)
	if !ok {
		return fmt.Errorf("expected cache/v1beta20201201storage/RedisLinkedServer but received %T instead", hub)
	}

	return server.AssignPropertiesToRedisLinkedServer(destination)
}

var _ genruntime.KubernetesResource = &RedisLinkedServer{}

// AzureName returns the Azure name of the resource
func (server *RedisLinkedServer) AzureName() string {
	return server.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (server RedisLinkedServer) GetAPIVersion() string {
	return "2020-12-01"
}

// GetResourceKind returns the kind of the resource
func (server *RedisLinkedServer) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (server *RedisLinkedServer) GetSpec() genruntime.ConvertibleSpec {
	return &server.Spec
}

// GetStatus returns the status of this resource
func (server *RedisLinkedServer) GetStatus() genruntime.ConvertibleStatus {
	return &server.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/linkedServers"
func (server *RedisLinkedServer) GetType() string {
	return "Microsoft.Cache/redis/linkedServers"
}

// NewEmptyStatus returns a new empty (blank) status
func (server *RedisLinkedServer) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RedisLinkedServerWithProperties_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (server *RedisLinkedServer) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(server.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  server.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (server *RedisLinkedServer) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RedisLinkedServerWithProperties_Status); ok {
		server.Status = *st
		return nil
	}

	// Convert status to required version
	var st RedisLinkedServerWithProperties_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	server.Status = st
	return nil
}

// AssignPropertiesFromRedisLinkedServer populates our RedisLinkedServer from the provided source RedisLinkedServer
func (server *RedisLinkedServer) AssignPropertiesFromRedisLinkedServer(source *v1beta20201201storage.RedisLinkedServer) error {

	// ObjectMeta
	server.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec RedisLinkedServers_Spec
	err := spec.AssignPropertiesFromRedisLinkedServersSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRedisLinkedServersSpec() to populate field Spec")
	}
	server.Spec = spec

	// Status
	var status RedisLinkedServerWithProperties_Status
	err = status.AssignPropertiesFromRedisLinkedServerWithPropertiesStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRedisLinkedServerWithPropertiesStatus() to populate field Status")
	}
	server.Status = status

	// No error
	return nil
}

// AssignPropertiesToRedisLinkedServer populates the provided destination RedisLinkedServer from our RedisLinkedServer
func (server *RedisLinkedServer) AssignPropertiesToRedisLinkedServer(destination *v1beta20201201storage.RedisLinkedServer) error {

	// ObjectMeta
	destination.ObjectMeta = *server.ObjectMeta.DeepCopy()

	// Spec
	var spec v1beta20201201storage.RedisLinkedServers_Spec
	err := server.Spec.AssignPropertiesToRedisLinkedServersSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRedisLinkedServersSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1beta20201201storage.RedisLinkedServerWithProperties_Status
	err = server.Status.AssignPropertiesToRedisLinkedServerWithPropertiesStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRedisLinkedServerWithPropertiesStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (server *RedisLinkedServer) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: server.Spec.OriginalVersion,
		Kind:    "RedisLinkedServer",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20201201.RedisLinkedServer
//Deprecated version of RedisLinkedServer. Use v1beta20201201.RedisLinkedServer instead
type RedisLinkedServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisLinkedServer `json:"items"`
}

//Storage version of v1alpha1api20201201.RedisLinkedServerWithProperties_Status
//Deprecated version of RedisLinkedServerWithProperties_Status. Use v1beta20201201.RedisLinkedServerWithProperties_Status instead
type RedisLinkedServerWithProperties_Status struct {
	Conditions               []conditions.Condition `json:"conditions,omitempty"`
	Id                       *string                `json:"id,omitempty"`
	LinkedRedisCacheId       *string                `json:"linkedRedisCacheId,omitempty"`
	LinkedRedisCacheLocation *string                `json:"linkedRedisCacheLocation,omitempty"`
	Name                     *string                `json:"name,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState        *string                `json:"provisioningState,omitempty"`
	ServerRole               *string                `json:"serverRole,omitempty"`
	Type                     *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RedisLinkedServerWithProperties_Status{}

// ConvertStatusFrom populates our RedisLinkedServerWithProperties_Status from the provided source
func (properties *RedisLinkedServerWithProperties_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1beta20201201storage.RedisLinkedServerWithProperties_Status)
	if ok {
		// Populate our instance from source
		return properties.AssignPropertiesFromRedisLinkedServerWithPropertiesStatus(src)
	}

	// Convert to an intermediate form
	src = &v1beta20201201storage.RedisLinkedServerWithProperties_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = properties.AssignPropertiesFromRedisLinkedServerWithPropertiesStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our RedisLinkedServerWithProperties_Status
func (properties *RedisLinkedServerWithProperties_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1beta20201201storage.RedisLinkedServerWithProperties_Status)
	if ok {
		// Populate destination from our instance
		return properties.AssignPropertiesToRedisLinkedServerWithPropertiesStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v1beta20201201storage.RedisLinkedServerWithProperties_Status{}
	err := properties.AssignPropertiesToRedisLinkedServerWithPropertiesStatus(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

// AssignPropertiesFromRedisLinkedServerWithPropertiesStatus populates our RedisLinkedServerWithProperties_Status from the provided source RedisLinkedServerWithProperties_Status
func (properties *RedisLinkedServerWithProperties_Status) AssignPropertiesFromRedisLinkedServerWithPropertiesStatus(source *v1beta20201201storage.RedisLinkedServerWithProperties_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	properties.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	properties.Id = genruntime.ClonePointerToString(source.Id)

	// LinkedRedisCacheId
	properties.LinkedRedisCacheId = genruntime.ClonePointerToString(source.LinkedRedisCacheId)

	// LinkedRedisCacheLocation
	properties.LinkedRedisCacheLocation = genruntime.ClonePointerToString(source.LinkedRedisCacheLocation)

	// Name
	properties.Name = genruntime.ClonePointerToString(source.Name)

	// ProvisioningState
	properties.ProvisioningState = genruntime.ClonePointerToString(source.ProvisioningState)

	// ServerRole
	properties.ServerRole = genruntime.ClonePointerToString(source.ServerRole)

	// Type
	properties.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		properties.PropertyBag = propertyBag
	} else {
		properties.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToRedisLinkedServerWithPropertiesStatus populates the provided destination RedisLinkedServerWithProperties_Status from our RedisLinkedServerWithProperties_Status
func (properties *RedisLinkedServerWithProperties_Status) AssignPropertiesToRedisLinkedServerWithPropertiesStatus(destination *v1beta20201201storage.RedisLinkedServerWithProperties_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(properties.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(properties.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(properties.Id)

	// LinkedRedisCacheId
	destination.LinkedRedisCacheId = genruntime.ClonePointerToString(properties.LinkedRedisCacheId)

	// LinkedRedisCacheLocation
	destination.LinkedRedisCacheLocation = genruntime.ClonePointerToString(properties.LinkedRedisCacheLocation)

	// Name
	destination.Name = genruntime.ClonePointerToString(properties.Name)

	// ProvisioningState
	destination.ProvisioningState = genruntime.ClonePointerToString(properties.ProvisioningState)

	// ServerRole
	destination.ServerRole = genruntime.ClonePointerToString(properties.ServerRole)

	// Type
	destination.Type = genruntime.ClonePointerToString(properties.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

//Storage version of v1alpha1api20201201.RedisLinkedServers_Spec
type RedisLinkedServers_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	//doesn't have to be.
	AzureName                string  `json:"azureName,omitempty"`
	LinkedRedisCacheLocation *string `json:"linkedRedisCacheLocation,omitempty"`

	// +kubebuilder:validation:Required
	LinkedRedisCacheReference *genruntime.ResourceReference `armReference:"LinkedRedisCacheId" json:"linkedRedisCacheReference,omitempty"`
	Location                  *string                       `json:"location,omitempty"`
	OriginalVersion           string                        `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	//Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	//controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	//reference to a cache.azure.com/Redis resource
	Owner       *genruntime.KnownResourceReference `group:"cache.azure.com" json:"owner,omitempty" kind:"Redis"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	ServerRole  *string                            `json:"serverRole,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &RedisLinkedServers_Spec{}

// ConvertSpecFrom populates our RedisLinkedServers_Spec from the provided source
func (servers *RedisLinkedServers_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1beta20201201storage.RedisLinkedServers_Spec)
	if ok {
		// Populate our instance from source
		return servers.AssignPropertiesFromRedisLinkedServersSpec(src)
	}

	// Convert to an intermediate form
	src = &v1beta20201201storage.RedisLinkedServers_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = servers.AssignPropertiesFromRedisLinkedServersSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our RedisLinkedServers_Spec
func (servers *RedisLinkedServers_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1beta20201201storage.RedisLinkedServers_Spec)
	if ok {
		// Populate destination from our instance
		return servers.AssignPropertiesToRedisLinkedServersSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v1beta20201201storage.RedisLinkedServers_Spec{}
	err := servers.AssignPropertiesToRedisLinkedServersSpec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignPropertiesFromRedisLinkedServersSpec populates our RedisLinkedServers_Spec from the provided source RedisLinkedServers_Spec
func (servers *RedisLinkedServers_Spec) AssignPropertiesFromRedisLinkedServersSpec(source *v1beta20201201storage.RedisLinkedServers_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	servers.AzureName = source.AzureName

	// LinkedRedisCacheLocation
	servers.LinkedRedisCacheLocation = genruntime.ClonePointerToString(source.LinkedRedisCacheLocation)

	// LinkedRedisCacheReference
	if source.LinkedRedisCacheReference != nil {
		linkedRedisCacheReference := source.LinkedRedisCacheReference.Copy()
		servers.LinkedRedisCacheReference = &linkedRedisCacheReference
	} else {
		servers.LinkedRedisCacheReference = nil
	}

	// Location
	servers.Location = genruntime.ClonePointerToString(source.Location)

	// OriginalVersion
	servers.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		servers.Owner = &owner
	} else {
		servers.Owner = nil
	}

	// ServerRole
	servers.ServerRole = genruntime.ClonePointerToString(source.ServerRole)

	// Tags
	servers.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		servers.PropertyBag = propertyBag
	} else {
		servers.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToRedisLinkedServersSpec populates the provided destination RedisLinkedServers_Spec from our RedisLinkedServers_Spec
func (servers *RedisLinkedServers_Spec) AssignPropertiesToRedisLinkedServersSpec(destination *v1beta20201201storage.RedisLinkedServers_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(servers.PropertyBag)

	// AzureName
	destination.AzureName = servers.AzureName

	// LinkedRedisCacheLocation
	destination.LinkedRedisCacheLocation = genruntime.ClonePointerToString(servers.LinkedRedisCacheLocation)

	// LinkedRedisCacheReference
	if servers.LinkedRedisCacheReference != nil {
		linkedRedisCacheReference := servers.LinkedRedisCacheReference.Copy()
		destination.LinkedRedisCacheReference = &linkedRedisCacheReference
	} else {
		destination.LinkedRedisCacheReference = nil
	}

	// Location
	destination.Location = genruntime.ClonePointerToString(servers.Location)

	// OriginalVersion
	destination.OriginalVersion = servers.OriginalVersion

	// Owner
	if servers.Owner != nil {
		owner := servers.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// ServerRole
	destination.ServerRole = genruntime.ClonePointerToString(servers.ServerRole)

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(servers.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

func init() {
	SchemeBuilder.Register(&RedisLinkedServer{}, &RedisLinkedServerList{})
}

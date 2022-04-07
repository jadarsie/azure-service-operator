// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210301storage

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/api/cache/v1beta20210301storage"
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
//Storage version of v1alpha1api20210301.RedisEnterpriseDatabase
//Deprecated version of RedisEnterpriseDatabase. Use v1beta20210301.RedisEnterpriseDatabase instead
type RedisEnterpriseDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisEnterpriseDatabases_Spec `json:"spec,omitempty"`
	Status            Database_Status               `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RedisEnterpriseDatabase{}

// GetConditions returns the conditions of the resource
func (database *RedisEnterpriseDatabase) GetConditions() conditions.Conditions {
	return database.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (database *RedisEnterpriseDatabase) SetConditions(conditions conditions.Conditions) {
	database.Status.Conditions = conditions
}

var _ conversion.Convertible = &RedisEnterpriseDatabase{}

// ConvertFrom populates our RedisEnterpriseDatabase from the provided hub RedisEnterpriseDatabase
func (database *RedisEnterpriseDatabase) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1beta20210301storage.RedisEnterpriseDatabase)
	if !ok {
		return fmt.Errorf("expected cache/v1beta20210301storage/RedisEnterpriseDatabase but received %T instead", hub)
	}

	return database.AssignPropertiesFromRedisEnterpriseDatabase(source)
}

// ConvertTo populates the provided hub RedisEnterpriseDatabase from our RedisEnterpriseDatabase
func (database *RedisEnterpriseDatabase) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1beta20210301storage.RedisEnterpriseDatabase)
	if !ok {
		return fmt.Errorf("expected cache/v1beta20210301storage/RedisEnterpriseDatabase but received %T instead", hub)
	}

	return database.AssignPropertiesToRedisEnterpriseDatabase(destination)
}

var _ genruntime.KubernetesResource = &RedisEnterpriseDatabase{}

// AzureName returns the Azure name of the resource
func (database *RedisEnterpriseDatabase) AzureName() string {
	return database.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-03-01"
func (database RedisEnterpriseDatabase) GetAPIVersion() string {
	return "2021-03-01"
}

// GetResourceKind returns the kind of the resource
func (database *RedisEnterpriseDatabase) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (database *RedisEnterpriseDatabase) GetSpec() genruntime.ConvertibleSpec {
	return &database.Spec
}

// GetStatus returns the status of this resource
func (database *RedisEnterpriseDatabase) GetStatus() genruntime.ConvertibleStatus {
	return &database.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redisEnterprise/databases"
func (database *RedisEnterpriseDatabase) GetType() string {
	return "Microsoft.Cache/redisEnterprise/databases"
}

// NewEmptyStatus returns a new empty (blank) status
func (database *RedisEnterpriseDatabase) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Database_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (database *RedisEnterpriseDatabase) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(database.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  database.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (database *RedisEnterpriseDatabase) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Database_Status); ok {
		database.Status = *st
		return nil
	}

	// Convert status to required version
	var st Database_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	database.Status = st
	return nil
}

// AssignPropertiesFromRedisEnterpriseDatabase populates our RedisEnterpriseDatabase from the provided source RedisEnterpriseDatabase
func (database *RedisEnterpriseDatabase) AssignPropertiesFromRedisEnterpriseDatabase(source *v1beta20210301storage.RedisEnterpriseDatabase) error {

	// ObjectMeta
	database.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec RedisEnterpriseDatabases_Spec
	err := spec.AssignPropertiesFromRedisEnterpriseDatabasesSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRedisEnterpriseDatabasesSpec() to populate field Spec")
	}
	database.Spec = spec

	// Status
	var status Database_Status
	err = status.AssignPropertiesFromDatabaseStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDatabaseStatus() to populate field Status")
	}
	database.Status = status

	// No error
	return nil
}

// AssignPropertiesToRedisEnterpriseDatabase populates the provided destination RedisEnterpriseDatabase from our RedisEnterpriseDatabase
func (database *RedisEnterpriseDatabase) AssignPropertiesToRedisEnterpriseDatabase(destination *v1beta20210301storage.RedisEnterpriseDatabase) error {

	// ObjectMeta
	destination.ObjectMeta = *database.ObjectMeta.DeepCopy()

	// Spec
	var spec v1beta20210301storage.RedisEnterpriseDatabases_Spec
	err := database.Spec.AssignPropertiesToRedisEnterpriseDatabasesSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRedisEnterpriseDatabasesSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1beta20210301storage.Database_Status
	err = database.Status.AssignPropertiesToDatabaseStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDatabaseStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (database *RedisEnterpriseDatabase) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: database.Spec.OriginalVersion,
		Kind:    "RedisEnterpriseDatabase",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210301.RedisEnterpriseDatabase
//Deprecated version of RedisEnterpriseDatabase. Use v1beta20210301.RedisEnterpriseDatabase instead
type RedisEnterpriseDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisEnterpriseDatabase `json:"items"`
}

//Storage version of v1alpha1api20210301.Database_Status
//Deprecated version of Database_Status. Use v1beta20210301.Database_Status instead
type Database_Status struct {
	ClientProtocol    *string                `json:"clientProtocol,omitempty"`
	ClusteringPolicy  *string                `json:"clusteringPolicy,omitempty"`
	Conditions        []conditions.Condition `json:"conditions,omitempty"`
	EvictionPolicy    *string                `json:"evictionPolicy,omitempty"`
	Id                *string                `json:"id,omitempty"`
	Modules           []Module_Status        `json:"modules,omitempty"`
	Name              *string                `json:"name,omitempty"`
	Persistence       *Persistence_Status    `json:"persistence,omitempty"`
	Port              *int                   `json:"port,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
	ResourceState     *string                `json:"resourceState,omitempty"`
	Type              *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Database_Status{}

// ConvertStatusFrom populates our Database_Status from the provided source
func (database *Database_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1beta20210301storage.Database_Status)
	if ok {
		// Populate our instance from source
		return database.AssignPropertiesFromDatabaseStatus(src)
	}

	// Convert to an intermediate form
	src = &v1beta20210301storage.Database_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = database.AssignPropertiesFromDatabaseStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Database_Status
func (database *Database_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1beta20210301storage.Database_Status)
	if ok {
		// Populate destination from our instance
		return database.AssignPropertiesToDatabaseStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v1beta20210301storage.Database_Status{}
	err := database.AssignPropertiesToDatabaseStatus(dst)
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

// AssignPropertiesFromDatabaseStatus populates our Database_Status from the provided source Database_Status
func (database *Database_Status) AssignPropertiesFromDatabaseStatus(source *v1beta20210301storage.Database_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// ClientProtocol
	database.ClientProtocol = genruntime.ClonePointerToString(source.ClientProtocol)

	// ClusteringPolicy
	database.ClusteringPolicy = genruntime.ClonePointerToString(source.ClusteringPolicy)

	// Conditions
	database.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// EvictionPolicy
	database.EvictionPolicy = genruntime.ClonePointerToString(source.EvictionPolicy)

	// Id
	database.Id = genruntime.ClonePointerToString(source.Id)

	// Modules
	if source.Modules != nil {
		moduleList := make([]Module_Status, len(source.Modules))
		for moduleIndex, moduleItem := range source.Modules {
			// Shadow the loop variable to avoid aliasing
			moduleItem := moduleItem
			var module Module_Status
			err := module.AssignPropertiesFromModuleStatus(&moduleItem)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesFromModuleStatus() to populate field Modules")
			}
			moduleList[moduleIndex] = module
		}
		database.Modules = moduleList
	} else {
		database.Modules = nil
	}

	// Name
	database.Name = genruntime.ClonePointerToString(source.Name)

	// Persistence
	if source.Persistence != nil {
		var persistence Persistence_Status
		err := persistence.AssignPropertiesFromPersistenceStatus(source.Persistence)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromPersistenceStatus() to populate field Persistence")
		}
		database.Persistence = &persistence
	} else {
		database.Persistence = nil
	}

	// Port
	database.Port = genruntime.ClonePointerToInt(source.Port)

	// ProvisioningState
	database.ProvisioningState = genruntime.ClonePointerToString(source.ProvisioningState)

	// ResourceState
	database.ResourceState = genruntime.ClonePointerToString(source.ResourceState)

	// Type
	database.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		database.PropertyBag = propertyBag
	} else {
		database.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToDatabaseStatus populates the provided destination Database_Status from our Database_Status
func (database *Database_Status) AssignPropertiesToDatabaseStatus(destination *v1beta20210301storage.Database_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(database.PropertyBag)

	// ClientProtocol
	destination.ClientProtocol = genruntime.ClonePointerToString(database.ClientProtocol)

	// ClusteringPolicy
	destination.ClusteringPolicy = genruntime.ClonePointerToString(database.ClusteringPolicy)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(database.Conditions)

	// EvictionPolicy
	destination.EvictionPolicy = genruntime.ClonePointerToString(database.EvictionPolicy)

	// Id
	destination.Id = genruntime.ClonePointerToString(database.Id)

	// Modules
	if database.Modules != nil {
		moduleList := make([]v1beta20210301storage.Module_Status, len(database.Modules))
		for moduleIndex, moduleItem := range database.Modules {
			// Shadow the loop variable to avoid aliasing
			moduleItem := moduleItem
			var module v1beta20210301storage.Module_Status
			err := moduleItem.AssignPropertiesToModuleStatus(&module)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesToModuleStatus() to populate field Modules")
			}
			moduleList[moduleIndex] = module
		}
		destination.Modules = moduleList
	} else {
		destination.Modules = nil
	}

	// Name
	destination.Name = genruntime.ClonePointerToString(database.Name)

	// Persistence
	if database.Persistence != nil {
		var persistence v1beta20210301storage.Persistence_Status
		err := database.Persistence.AssignPropertiesToPersistenceStatus(&persistence)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToPersistenceStatus() to populate field Persistence")
		}
		destination.Persistence = &persistence
	} else {
		destination.Persistence = nil
	}

	// Port
	destination.Port = genruntime.ClonePointerToInt(database.Port)

	// ProvisioningState
	destination.ProvisioningState = genruntime.ClonePointerToString(database.ProvisioningState)

	// ResourceState
	destination.ResourceState = genruntime.ClonePointerToString(database.ResourceState)

	// Type
	destination.Type = genruntime.ClonePointerToString(database.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

//Storage version of v1alpha1api20210301.RedisEnterpriseDatabases_Spec
type RedisEnterpriseDatabases_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	//doesn't have to be.
	AzureName        string   `json:"azureName,omitempty"`
	ClientProtocol   *string  `json:"clientProtocol,omitempty"`
	ClusteringPolicy *string  `json:"clusteringPolicy,omitempty"`
	EvictionPolicy   *string  `json:"evictionPolicy,omitempty"`
	Location         *string  `json:"location,omitempty"`
	Modules          []Module `json:"modules,omitempty"`
	OriginalVersion  string   `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	//Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	//controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	//reference to a cache.azure.com/RedisEnterprise resource
	Owner       *genruntime.KnownResourceReference `group:"cache.azure.com" json:"owner,omitempty" kind:"RedisEnterprise"`
	Persistence *Persistence                       `json:"persistence,omitempty"`
	Port        *int                               `json:"port,omitempty"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &RedisEnterpriseDatabases_Spec{}

// ConvertSpecFrom populates our RedisEnterpriseDatabases_Spec from the provided source
func (databases *RedisEnterpriseDatabases_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1beta20210301storage.RedisEnterpriseDatabases_Spec)
	if ok {
		// Populate our instance from source
		return databases.AssignPropertiesFromRedisEnterpriseDatabasesSpec(src)
	}

	// Convert to an intermediate form
	src = &v1beta20210301storage.RedisEnterpriseDatabases_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = databases.AssignPropertiesFromRedisEnterpriseDatabasesSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our RedisEnterpriseDatabases_Spec
func (databases *RedisEnterpriseDatabases_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1beta20210301storage.RedisEnterpriseDatabases_Spec)
	if ok {
		// Populate destination from our instance
		return databases.AssignPropertiesToRedisEnterpriseDatabasesSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v1beta20210301storage.RedisEnterpriseDatabases_Spec{}
	err := databases.AssignPropertiesToRedisEnterpriseDatabasesSpec(dst)
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

// AssignPropertiesFromRedisEnterpriseDatabasesSpec populates our RedisEnterpriseDatabases_Spec from the provided source RedisEnterpriseDatabases_Spec
func (databases *RedisEnterpriseDatabases_Spec) AssignPropertiesFromRedisEnterpriseDatabasesSpec(source *v1beta20210301storage.RedisEnterpriseDatabases_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	databases.AzureName = source.AzureName

	// ClientProtocol
	databases.ClientProtocol = genruntime.ClonePointerToString(source.ClientProtocol)

	// ClusteringPolicy
	databases.ClusteringPolicy = genruntime.ClonePointerToString(source.ClusteringPolicy)

	// EvictionPolicy
	databases.EvictionPolicy = genruntime.ClonePointerToString(source.EvictionPolicy)

	// Location
	databases.Location = genruntime.ClonePointerToString(source.Location)

	// Modules
	if source.Modules != nil {
		moduleList := make([]Module, len(source.Modules))
		for moduleIndex, moduleItem := range source.Modules {
			// Shadow the loop variable to avoid aliasing
			moduleItem := moduleItem
			var module Module
			err := module.AssignPropertiesFromModule(&moduleItem)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesFromModule() to populate field Modules")
			}
			moduleList[moduleIndex] = module
		}
		databases.Modules = moduleList
	} else {
		databases.Modules = nil
	}

	// OriginalVersion
	databases.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		databases.Owner = &owner
	} else {
		databases.Owner = nil
	}

	// Persistence
	if source.Persistence != nil {
		var persistence Persistence
		err := persistence.AssignPropertiesFromPersistence(source.Persistence)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromPersistence() to populate field Persistence")
		}
		databases.Persistence = &persistence
	} else {
		databases.Persistence = nil
	}

	// Port
	databases.Port = genruntime.ClonePointerToInt(source.Port)

	// Tags
	databases.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		databases.PropertyBag = propertyBag
	} else {
		databases.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToRedisEnterpriseDatabasesSpec populates the provided destination RedisEnterpriseDatabases_Spec from our RedisEnterpriseDatabases_Spec
func (databases *RedisEnterpriseDatabases_Spec) AssignPropertiesToRedisEnterpriseDatabasesSpec(destination *v1beta20210301storage.RedisEnterpriseDatabases_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(databases.PropertyBag)

	// AzureName
	destination.AzureName = databases.AzureName

	// ClientProtocol
	destination.ClientProtocol = genruntime.ClonePointerToString(databases.ClientProtocol)

	// ClusteringPolicy
	destination.ClusteringPolicy = genruntime.ClonePointerToString(databases.ClusteringPolicy)

	// EvictionPolicy
	destination.EvictionPolicy = genruntime.ClonePointerToString(databases.EvictionPolicy)

	// Location
	destination.Location = genruntime.ClonePointerToString(databases.Location)

	// Modules
	if databases.Modules != nil {
		moduleList := make([]v1beta20210301storage.Module, len(databases.Modules))
		for moduleIndex, moduleItem := range databases.Modules {
			// Shadow the loop variable to avoid aliasing
			moduleItem := moduleItem
			var module v1beta20210301storage.Module
			err := moduleItem.AssignPropertiesToModule(&module)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesToModule() to populate field Modules")
			}
			moduleList[moduleIndex] = module
		}
		destination.Modules = moduleList
	} else {
		destination.Modules = nil
	}

	// OriginalVersion
	destination.OriginalVersion = databases.OriginalVersion

	// Owner
	if databases.Owner != nil {
		owner := databases.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Persistence
	if databases.Persistence != nil {
		var persistence v1beta20210301storage.Persistence
		err := databases.Persistence.AssignPropertiesToPersistence(&persistence)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToPersistence() to populate field Persistence")
		}
		destination.Persistence = &persistence
	} else {
		destination.Persistence = nil
	}

	// Port
	destination.Port = genruntime.ClonePointerToInt(databases.Port)

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(databases.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

//Storage version of v1alpha1api20210301.Module
//Deprecated version of Module. Use v1beta20210301.Module instead
type Module struct {
	Args        *string                `json:"args,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignPropertiesFromModule populates our Module from the provided source Module
func (module *Module) AssignPropertiesFromModule(source *v1beta20210301storage.Module) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Args
	module.Args = genruntime.ClonePointerToString(source.Args)

	// Name
	module.Name = genruntime.ClonePointerToString(source.Name)

	// Update the property bag
	if len(propertyBag) > 0 {
		module.PropertyBag = propertyBag
	} else {
		module.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToModule populates the provided destination Module from our Module
func (module *Module) AssignPropertiesToModule(destination *v1beta20210301storage.Module) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(module.PropertyBag)

	// Args
	destination.Args = genruntime.ClonePointerToString(module.Args)

	// Name
	destination.Name = genruntime.ClonePointerToString(module.Name)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

//Storage version of v1alpha1api20210301.Module_Status
//Deprecated version of Module_Status. Use v1beta20210301.Module_Status instead
type Module_Status struct {
	Args        *string                `json:"args,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Version     *string                `json:"version,omitempty"`
}

// AssignPropertiesFromModuleStatus populates our Module_Status from the provided source Module_Status
func (module *Module_Status) AssignPropertiesFromModuleStatus(source *v1beta20210301storage.Module_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Args
	module.Args = genruntime.ClonePointerToString(source.Args)

	// Name
	module.Name = genruntime.ClonePointerToString(source.Name)

	// Version
	module.Version = genruntime.ClonePointerToString(source.Version)

	// Update the property bag
	if len(propertyBag) > 0 {
		module.PropertyBag = propertyBag
	} else {
		module.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToModuleStatus populates the provided destination Module_Status from our Module_Status
func (module *Module_Status) AssignPropertiesToModuleStatus(destination *v1beta20210301storage.Module_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(module.PropertyBag)

	// Args
	destination.Args = genruntime.ClonePointerToString(module.Args)

	// Name
	destination.Name = genruntime.ClonePointerToString(module.Name)

	// Version
	destination.Version = genruntime.ClonePointerToString(module.Version)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

//Storage version of v1alpha1api20210301.Persistence
//Deprecated version of Persistence. Use v1beta20210301.Persistence instead
type Persistence struct {
	AofEnabled   *bool                  `json:"aofEnabled,omitempty"`
	AofFrequency *string                `json:"aofFrequency,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RdbEnabled   *bool                  `json:"rdbEnabled,omitempty"`
	RdbFrequency *string                `json:"rdbFrequency,omitempty"`
}

// AssignPropertiesFromPersistence populates our Persistence from the provided source Persistence
func (persistence *Persistence) AssignPropertiesFromPersistence(source *v1beta20210301storage.Persistence) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AofEnabled
	if source.AofEnabled != nil {
		aofEnabled := *source.AofEnabled
		persistence.AofEnabled = &aofEnabled
	} else {
		persistence.AofEnabled = nil
	}

	// AofFrequency
	persistence.AofFrequency = genruntime.ClonePointerToString(source.AofFrequency)

	// RdbEnabled
	if source.RdbEnabled != nil {
		rdbEnabled := *source.RdbEnabled
		persistence.RdbEnabled = &rdbEnabled
	} else {
		persistence.RdbEnabled = nil
	}

	// RdbFrequency
	persistence.RdbFrequency = genruntime.ClonePointerToString(source.RdbFrequency)

	// Update the property bag
	if len(propertyBag) > 0 {
		persistence.PropertyBag = propertyBag
	} else {
		persistence.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToPersistence populates the provided destination Persistence from our Persistence
func (persistence *Persistence) AssignPropertiesToPersistence(destination *v1beta20210301storage.Persistence) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(persistence.PropertyBag)

	// AofEnabled
	if persistence.AofEnabled != nil {
		aofEnabled := *persistence.AofEnabled
		destination.AofEnabled = &aofEnabled
	} else {
		destination.AofEnabled = nil
	}

	// AofFrequency
	destination.AofFrequency = genruntime.ClonePointerToString(persistence.AofFrequency)

	// RdbEnabled
	if persistence.RdbEnabled != nil {
		rdbEnabled := *persistence.RdbEnabled
		destination.RdbEnabled = &rdbEnabled
	} else {
		destination.RdbEnabled = nil
	}

	// RdbFrequency
	destination.RdbFrequency = genruntime.ClonePointerToString(persistence.RdbFrequency)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

//Storage version of v1alpha1api20210301.Persistence_Status
//Deprecated version of Persistence_Status. Use v1beta20210301.Persistence_Status instead
type Persistence_Status struct {
	AofEnabled   *bool                  `json:"aofEnabled,omitempty"`
	AofFrequency *string                `json:"aofFrequency,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RdbEnabled   *bool                  `json:"rdbEnabled,omitempty"`
	RdbFrequency *string                `json:"rdbFrequency,omitempty"`
}

// AssignPropertiesFromPersistenceStatus populates our Persistence_Status from the provided source Persistence_Status
func (persistence *Persistence_Status) AssignPropertiesFromPersistenceStatus(source *v1beta20210301storage.Persistence_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AofEnabled
	if source.AofEnabled != nil {
		aofEnabled := *source.AofEnabled
		persistence.AofEnabled = &aofEnabled
	} else {
		persistence.AofEnabled = nil
	}

	// AofFrequency
	persistence.AofFrequency = genruntime.ClonePointerToString(source.AofFrequency)

	// RdbEnabled
	if source.RdbEnabled != nil {
		rdbEnabled := *source.RdbEnabled
		persistence.RdbEnabled = &rdbEnabled
	} else {
		persistence.RdbEnabled = nil
	}

	// RdbFrequency
	persistence.RdbFrequency = genruntime.ClonePointerToString(source.RdbFrequency)

	// Update the property bag
	if len(propertyBag) > 0 {
		persistence.PropertyBag = propertyBag
	} else {
		persistence.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToPersistenceStatus populates the provided destination Persistence_Status from our Persistence_Status
func (persistence *Persistence_Status) AssignPropertiesToPersistenceStatus(destination *v1beta20210301storage.Persistence_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(persistence.PropertyBag)

	// AofEnabled
	if persistence.AofEnabled != nil {
		aofEnabled := *persistence.AofEnabled
		destination.AofEnabled = &aofEnabled
	} else {
		destination.AofEnabled = nil
	}

	// AofFrequency
	destination.AofFrequency = genruntime.ClonePointerToString(persistence.AofFrequency)

	// RdbEnabled
	if persistence.RdbEnabled != nil {
		rdbEnabled := *persistence.RdbEnabled
		destination.RdbEnabled = &rdbEnabled
	} else {
		destination.RdbEnabled = nil
	}

	// RdbFrequency
	destination.RdbFrequency = genruntime.ClonePointerToString(persistence.RdbFrequency)

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
	SchemeBuilder.Register(&RedisEnterpriseDatabase{}, &RedisEnterpriseDatabaseList{})
}

// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/api/cache/v1alpha1api20201201storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Generator information:
//- Generated from: /redis/resource-manager/Microsoft.Cache/stable/2020-12-01/redis.json
//- ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/linkedServers/{linkedServerName}
type RedisLinkedServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisLinkedServer_Spec   `json:"spec,omitempty"`
	Status            RedisLinkedServer_Status `json:"status,omitempty"`
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
	source, ok := hub.(*v1alpha1api20201201storage.RedisLinkedServer)
	if !ok {
		return fmt.Errorf("expected storage:cache/v1alpha1api20201201storage/RedisLinkedServer but received %T instead", hub)
	}

	return server.AssignPropertiesFromRedisLinkedServer(source)
}

// ConvertTo populates the provided hub RedisLinkedServer from our RedisLinkedServer
func (server *RedisLinkedServer) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1alpha1api20201201storage.RedisLinkedServer)
	if !ok {
		return fmt.Errorf("expected storage:cache/v1alpha1api20201201storage/RedisLinkedServer but received %T instead", hub)
	}

	return server.AssignPropertiesToRedisLinkedServer(destination)
}

// +kubebuilder:webhook:path=/mutate-cache-azure-com-v1alpha1api20201201-redislinkedserver,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cache.azure.com,resources=redislinkedservers,verbs=create;update,versions=v1alpha1api20201201,name=default.v1alpha1api20201201.redislinkedservers.cache.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &RedisLinkedServer{}

// Default applies defaults to the RedisLinkedServer resource
func (server *RedisLinkedServer) Default() {
	server.defaultImpl()
	var temp interface{} = server
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (server *RedisLinkedServer) defaultAzureName() {
	if server.Spec.AzureName == "" {
		server.Spec.AzureName = server.Name
	}
}

// defaultImpl applies the code generated defaults to the RedisLinkedServer resource
func (server *RedisLinkedServer) defaultImpl() { server.defaultAzureName() }

var _ genruntime.KubernetesResource = &RedisLinkedServer{}

// AzureName returns the Azure name of the resource
func (server *RedisLinkedServer) AzureName() string {
	return server.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (server RedisLinkedServer) GetAPIVersion() string {
	return string(APIVersionValue)
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
	return &RedisLinkedServer_Status{}
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
	if st, ok := status.(*RedisLinkedServer_Status); ok {
		server.Status = *st
		return nil
	}

	// Convert status to required version
	var st RedisLinkedServer_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	server.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-cache-azure-com-v1alpha1api20201201-redislinkedserver,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cache.azure.com,resources=redislinkedservers,verbs=create;update,versions=v1alpha1api20201201,name=validate.v1alpha1api20201201.redislinkedservers.cache.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &RedisLinkedServer{}

// ValidateCreate validates the creation of the resource
func (server *RedisLinkedServer) ValidateCreate() error {
	validations := server.createValidations()
	var temp interface{} = server
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateDelete validates the deletion of the resource
func (server *RedisLinkedServer) ValidateDelete() error {
	validations := server.deleteValidations()
	var temp interface{} = server
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateUpdate validates an update of the resource
func (server *RedisLinkedServer) ValidateUpdate(old runtime.Object) error {
	validations := server.updateValidations()
	var temp interface{} = server
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation(old)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// createValidations validates the creation of the resource
func (server *RedisLinkedServer) createValidations() []func() error {
	return []func() error{server.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (server *RedisLinkedServer) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (server *RedisLinkedServer) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return server.validateResourceReferences()
		},
	}
}

// validateResourceReferences validates all resource references
func (server *RedisLinkedServer) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&server.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromRedisLinkedServer populates our RedisLinkedServer from the provided source RedisLinkedServer
func (server *RedisLinkedServer) AssignPropertiesFromRedisLinkedServer(source *v1alpha1api20201201storage.RedisLinkedServer) error {

	// ObjectMeta
	server.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec RedisLinkedServer_Spec
	err := spec.AssignPropertiesFromRedisLinkedServer_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRedisLinkedServer_Spec() to populate field Spec")
	}
	server.Spec = spec

	// Status
	var status RedisLinkedServer_Status
	err = status.AssignPropertiesFromRedisLinkedServer_Status(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRedisLinkedServer_Status() to populate field Status")
	}
	server.Status = status

	// No error
	return nil
}

// AssignPropertiesToRedisLinkedServer populates the provided destination RedisLinkedServer from our RedisLinkedServer
func (server *RedisLinkedServer) AssignPropertiesToRedisLinkedServer(destination *v1alpha1api20201201storage.RedisLinkedServer) error {

	// ObjectMeta
	destination.ObjectMeta = *server.ObjectMeta.DeepCopy()

	// Spec
	var spec v1alpha1api20201201storage.RedisLinkedServer_Spec
	err := server.Spec.AssignPropertiesToRedisLinkedServer_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRedisLinkedServer_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1alpha1api20201201storage.RedisLinkedServer_Status
	err = server.Status.AssignPropertiesToRedisLinkedServer_Status(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRedisLinkedServer_Status() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (server *RedisLinkedServer) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: server.Spec.OriginalVersion(),
		Kind:    "RedisLinkedServer",
	}
}

// +kubebuilder:object:root=true
//Generator information:
//- Generated from: /redis/resource-manager/Microsoft.Cache/stable/2020-12-01/redis.json
//- ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/linkedServers/{linkedServerName}
type RedisLinkedServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisLinkedServer `json:"items"`
}

type RedisLinkedServer_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName string `json:"azureName"`

	// +kubebuilder:validation:Required
	//LinkedRedisCacheLocation: Location of the linked redis cache.
	LinkedRedisCacheLocation string `json:"linkedRedisCacheLocation"`

	// +kubebuilder:validation:Required
	//LinkedRedisCacheReference: Fully qualified resourceId of the linked redis cache.
	LinkedRedisCacheReference genruntime.ResourceReference `armReference:"LinkedRedisCacheId" json:"linkedRedisCacheReference"`

	// +kubebuilder:validation:Required
	Owner genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`

	// +kubebuilder:validation:Required
	//ServerRole: Role of the linked server.
	ServerRole RedisLinkedServerPropertiesServerRole `json:"serverRole"`
}

var _ genruntime.ARMTransformer = &RedisLinkedServer_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (server *RedisLinkedServer_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if server == nil {
		return nil, nil
	}
	var result RedisLinkedServer_SpecARM

	// Set property ‘AzureName’:
	result.AzureName = server.AzureName

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	result.Properties.LinkedRedisCacheLocation = server.LinkedRedisCacheLocation
	linkedRedisCacheIdARMID, err := resolved.ResolvedReferences.ARMIDOrErr(server.LinkedRedisCacheReference)
	if err != nil {
		return nil, err
	}
	result.Properties.LinkedRedisCacheId = linkedRedisCacheIdARMID
	result.Properties.ServerRole = server.ServerRole
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (server *RedisLinkedServer_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RedisLinkedServer_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (server *RedisLinkedServer_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RedisLinkedServer_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RedisLinkedServer_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	server.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘LinkedRedisCacheLocation’:
	// copying flattened property:
	server.LinkedRedisCacheLocation = typedInput.Properties.LinkedRedisCacheLocation

	// no assignment for property ‘LinkedRedisCacheReference’

	// Set property ‘Owner’:
	server.Owner = genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘ServerRole’:
	// copying flattened property:
	server.ServerRole = typedInput.Properties.ServerRole

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &RedisLinkedServer_Spec{}

// ConvertSpecFrom populates our RedisLinkedServer_Spec from the provided source
func (server *RedisLinkedServer_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1alpha1api20201201storage.RedisLinkedServer_Spec)
	if ok {
		// Populate our instance from source
		return server.AssignPropertiesFromRedisLinkedServer_Spec(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20201201storage.RedisLinkedServer_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = server.AssignPropertiesFromRedisLinkedServer_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our RedisLinkedServer_Spec
func (server *RedisLinkedServer_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1alpha1api20201201storage.RedisLinkedServer_Spec)
	if ok {
		// Populate destination from our instance
		return server.AssignPropertiesToRedisLinkedServer_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20201201storage.RedisLinkedServer_Spec{}
	err := server.AssignPropertiesToRedisLinkedServer_Spec(dst)
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

// AssignPropertiesFromRedisLinkedServer_Spec populates our RedisLinkedServer_Spec from the provided source RedisLinkedServer_Spec
func (server *RedisLinkedServer_Spec) AssignPropertiesFromRedisLinkedServer_Spec(source *v1alpha1api20201201storage.RedisLinkedServer_Spec) error {

	// AzureName
	server.AzureName = source.AzureName

	// LinkedRedisCacheLocation
	server.LinkedRedisCacheLocation = genruntime.GetOptionalStringValue(source.LinkedRedisCacheLocation)

	// LinkedRedisCacheReference
	server.LinkedRedisCacheReference = source.LinkedRedisCacheReference.Copy()

	// Owner
	server.Owner = source.Owner.Copy()

	// ServerRole
	if source.ServerRole != nil {
		server.ServerRole = RedisLinkedServerPropertiesServerRole(*source.ServerRole)
	} else {
		server.ServerRole = ""
	}

	// No error
	return nil
}

// AssignPropertiesToRedisLinkedServer_Spec populates the provided destination RedisLinkedServer_Spec from our RedisLinkedServer_Spec
func (server *RedisLinkedServer_Spec) AssignPropertiesToRedisLinkedServer_Spec(destination *v1alpha1api20201201storage.RedisLinkedServer_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = server.AzureName

	// LinkedRedisCacheLocation
	linkedRedisCacheLocation := server.LinkedRedisCacheLocation
	destination.LinkedRedisCacheLocation = &linkedRedisCacheLocation

	// LinkedRedisCacheReference
	destination.LinkedRedisCacheReference = server.LinkedRedisCacheReference.Copy()

	// OriginalVersion
	destination.OriginalVersion = server.OriginalVersion()

	// Owner
	destination.Owner = server.Owner.Copy()

	// ServerRole
	serverRole := string(server.ServerRole)
	destination.ServerRole = &serverRole

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (server *RedisLinkedServer_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (server *RedisLinkedServer_Spec) SetAzureName(azureName string) { server.AzureName = azureName }

type RedisLinkedServer_Status struct {
	//Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	//LinkedRedisCacheId: Fully qualified resourceId of the linked redis cache.
	LinkedRedisCacheId *string `json:"linkedRedisCacheId,omitempty"`

	//LinkedRedisCacheLocation: Location of the linked redis cache.
	LinkedRedisCacheLocation *string `json:"linkedRedisCacheLocation,omitempty"`

	//ServerRole: Role of the linked server.
	ServerRole *string `json:"serverRole,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RedisLinkedServer_Status{}

// ConvertStatusFrom populates our RedisLinkedServer_Status from the provided source
func (server *RedisLinkedServer_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1alpha1api20201201storage.RedisLinkedServer_Status)
	if ok {
		// Populate our instance from source
		return server.AssignPropertiesFromRedisLinkedServer_Status(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20201201storage.RedisLinkedServer_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = server.AssignPropertiesFromRedisLinkedServer_Status(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our RedisLinkedServer_Status
func (server *RedisLinkedServer_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1alpha1api20201201storage.RedisLinkedServer_Status)
	if ok {
		// Populate destination from our instance
		return server.AssignPropertiesToRedisLinkedServer_Status(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20201201storage.RedisLinkedServer_Status{}
	err := server.AssignPropertiesToRedisLinkedServer_Status(dst)
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

var _ genruntime.FromARMConverter = &RedisLinkedServer_Status{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (server *RedisLinkedServer_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RedisLinkedServer_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (server *RedisLinkedServer_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RedisLinkedServer_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RedisLinkedServer_StatusARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘LinkedRedisCacheId’:
	// copying flattened property:
	if typedInput.Properties != nil {
		server.LinkedRedisCacheId = &typedInput.Properties.LinkedRedisCacheId
	}

	// Set property ‘LinkedRedisCacheLocation’:
	// copying flattened property:
	if typedInput.Properties != nil {
		server.LinkedRedisCacheLocation = &typedInput.Properties.LinkedRedisCacheLocation
	}

	// Set property ‘ServerRole’:
	// copying flattened property:
	if typedInput.Properties != nil {
		server.ServerRole = &typedInput.Properties.ServerRole
	}

	// No error
	return nil
}

// AssignPropertiesFromRedisLinkedServer_Status populates our RedisLinkedServer_Status from the provided source RedisLinkedServer_Status
func (server *RedisLinkedServer_Status) AssignPropertiesFromRedisLinkedServer_Status(source *v1alpha1api20201201storage.RedisLinkedServer_Status) error {

	// Conditions
	server.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// LinkedRedisCacheId
	server.LinkedRedisCacheId = genruntime.ClonePointerToString(source.LinkedRedisCacheId)

	// LinkedRedisCacheLocation
	server.LinkedRedisCacheLocation = genruntime.ClonePointerToString(source.LinkedRedisCacheLocation)

	// ServerRole
	server.ServerRole = genruntime.ClonePointerToString(source.ServerRole)

	// No error
	return nil
}

// AssignPropertiesToRedisLinkedServer_Status populates the provided destination RedisLinkedServer_Status from our RedisLinkedServer_Status
func (server *RedisLinkedServer_Status) AssignPropertiesToRedisLinkedServer_Status(destination *v1alpha1api20201201storage.RedisLinkedServer_Status) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(server.Conditions)

	// LinkedRedisCacheId
	destination.LinkedRedisCacheId = genruntime.ClonePointerToString(server.LinkedRedisCacheId)

	// LinkedRedisCacheLocation
	destination.LinkedRedisCacheLocation = genruntime.ClonePointerToString(server.LinkedRedisCacheLocation)

	// ServerRole
	destination.ServerRole = genruntime.ClonePointerToString(server.ServerRole)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// +kubebuilder:validation:Enum={"Primary","Secondary"}
type RedisLinkedServerPropertiesServerRole string

const (
	RedisLinkedServerPropertiesServerRolePrimary   = RedisLinkedServerPropertiesServerRole("Primary")
	RedisLinkedServerPropertiesServerRoleSecondary = RedisLinkedServerPropertiesServerRole("Secondary")
)

func init() {
	SchemeBuilder.Register(&RedisLinkedServer{}, &RedisLinkedServerList{})
}

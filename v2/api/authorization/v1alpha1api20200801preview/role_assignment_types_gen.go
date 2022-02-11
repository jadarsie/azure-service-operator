// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200801preview

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/api/authorization/v1alpha1api20200801previewstorage"
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
//- Generated from: /authorization/resource-manager/Microsoft.Authorization/preview/2020-08-01-preview/authorization-RoleAssignmentsCalls.json
//- ARM URI: /{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}
type RoleAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoleAssignments_SPEC  `json:"spec,omitempty"`
	Status            RoleAssignment_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RoleAssignment{}

// GetConditions returns the conditions of the resource
func (assignment *RoleAssignment) GetConditions() conditions.Conditions {
	return assignment.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (assignment *RoleAssignment) SetConditions(conditions conditions.Conditions) {
	assignment.Status.Conditions = conditions
}

var _ conversion.Convertible = &RoleAssignment{}

// ConvertFrom populates our RoleAssignment from the provided hub RoleAssignment
func (assignment *RoleAssignment) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1alpha1api20200801previewstorage.RoleAssignment)
	if !ok {
		return fmt.Errorf("expected storage:authorization/v1alpha1api20200801previewstorage/RoleAssignment but received %T instead", hub)
	}

	return assignment.AssignPropertiesFromRoleAssignment(source)
}

// ConvertTo populates the provided hub RoleAssignment from our RoleAssignment
func (assignment *RoleAssignment) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1alpha1api20200801previewstorage.RoleAssignment)
	if !ok {
		return fmt.Errorf("expected storage:authorization/v1alpha1api20200801previewstorage/RoleAssignment but received %T instead", hub)
	}

	return assignment.AssignPropertiesToRoleAssignment(destination)
}

// +kubebuilder:webhook:path=/mutate-authorization-azure-com-v1alpha1api20200801preview-roleassignment,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=authorization.azure.com,resources=roleassignments,verbs=create;update,versions=v1alpha1api20200801preview,name=default.v1alpha1api20200801preview.roleassignments.authorization.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &RoleAssignment{}

// Default applies defaults to the RoleAssignment resource
func (assignment *RoleAssignment) Default() {
	assignment.defaultImpl()
	var temp interface{} = assignment
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (assignment *RoleAssignment) defaultAzureName() {
	if assignment.Spec.AzureName == "" {
		assignment.Spec.AzureName = assignment.Name
	}
}

// defaultImpl applies the code generated defaults to the RoleAssignment resource
func (assignment *RoleAssignment) defaultImpl() { assignment.defaultAzureName() }

var _ genruntime.KubernetesResource = &RoleAssignment{}

// AzureName returns the Azure name of the resource
func (assignment *RoleAssignment) AzureName() string {
	return assignment.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-08-01"
func (assignment RoleAssignment) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (assignment *RoleAssignment) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (assignment *RoleAssignment) GetSpec() genruntime.ConvertibleSpec {
	return &assignment.Spec
}

// GetStatus returns the status of this resource
func (assignment *RoleAssignment) GetStatus() genruntime.ConvertibleStatus {
	return &assignment.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (assignment *RoleAssignment) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (assignment *RoleAssignment) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RoleAssignment_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (assignment *RoleAssignment) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(assignment.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  assignment.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (assignment *RoleAssignment) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RoleAssignment_Status); ok {
		assignment.Status = *st
		return nil
	}

	// Convert status to required version
	var st RoleAssignment_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	assignment.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-authorization-azure-com-v1alpha1api20200801preview-roleassignment,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=authorization.azure.com,resources=roleassignments,verbs=create;update,versions=v1alpha1api20200801preview,name=validate.v1alpha1api20200801preview.roleassignments.authorization.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &RoleAssignment{}

// ValidateCreate validates the creation of the resource
func (assignment *RoleAssignment) ValidateCreate() error {
	validations := assignment.createValidations()
	var temp interface{} = assignment
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
func (assignment *RoleAssignment) ValidateDelete() error {
	validations := assignment.deleteValidations()
	var temp interface{} = assignment
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
func (assignment *RoleAssignment) ValidateUpdate(old runtime.Object) error {
	validations := assignment.updateValidations()
	var temp interface{} = assignment
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
func (assignment *RoleAssignment) createValidations() []func() error {
	return []func() error{assignment.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (assignment *RoleAssignment) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (assignment *RoleAssignment) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return assignment.validateResourceReferences()
		},
	}
}

// validateResourceReferences validates all resource references
func (assignment *RoleAssignment) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&assignment.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromRoleAssignment populates our RoleAssignment from the provided source RoleAssignment
func (assignment *RoleAssignment) AssignPropertiesFromRoleAssignment(source *v1alpha1api20200801previewstorage.RoleAssignment) error {

	// ObjectMeta
	assignment.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec RoleAssignments_SPEC
	err := spec.AssignPropertiesFromRoleAssignments_SPEC(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRoleAssignments_SPEC() to populate field Spec")
	}
	assignment.Spec = spec

	// Status
	var status RoleAssignment_Status
	err = status.AssignPropertiesFromRoleAssignment_Status(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRoleAssignment_Status() to populate field Status")
	}
	assignment.Status = status

	// No error
	return nil
}

// AssignPropertiesToRoleAssignment populates the provided destination RoleAssignment from our RoleAssignment
func (assignment *RoleAssignment) AssignPropertiesToRoleAssignment(destination *v1alpha1api20200801previewstorage.RoleAssignment) error {

	// ObjectMeta
	destination.ObjectMeta = *assignment.ObjectMeta.DeepCopy()

	// Spec
	var spec v1alpha1api20200801previewstorage.RoleAssignments_SPEC
	err := assignment.Spec.AssignPropertiesToRoleAssignments_SPEC(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRoleAssignments_SPEC() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1alpha1api20200801previewstorage.RoleAssignment_Status
	err = assignment.Status.AssignPropertiesToRoleAssignment_Status(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRoleAssignment_Status() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (assignment *RoleAssignment) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: assignment.Spec.OriginalVersion(),
		Kind:    "RoleAssignment",
	}
}

// +kubebuilder:object:root=true
//Generator information:
//- Generated from: /authorization/resource-manager/Microsoft.Authorization/preview/2020-08-01-preview/authorization-RoleAssignmentsCalls.json
//- ARM URI: /{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}
type RoleAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleAssignment `json:"items"`
}

// +kubebuilder:validation:Enum={"2020-08-01"}
type APIVersion string

const APIVersionValue = APIVersion("2020-08-01")

type RoleAssignment_Status struct {
	//Condition: The conditions on the role assignment. This limits the resources it
	//can be assigned to. e.g.:
	//@Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName]
	//StringEqualsIgnoreCase 'foo_storage_container'
	Condition *string `json:"condition,omitempty"`

	//ConditionVersion: Version of the condition. Currently accepted value is '2.0'
	ConditionVersion *string `json:"conditionVersion,omitempty"`

	//Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	//CreatedBy: Id of the user who created the assignment
	CreatedBy *string `json:"createdBy,omitempty"`

	//CreatedOn: Time it was created
	CreatedOn *string `json:"createdOn,omitempty"`

	//DelegatedManagedIdentityResourceId: Id of the delegated managed identity resource
	DelegatedManagedIdentityResourceId *string `json:"delegatedManagedIdentityResourceId,omitempty"`

	//Description: Description of role assignment
	Description *string `json:"description,omitempty"`

	//PrincipalId: The principal ID.
	PrincipalId *string `json:"principalId,omitempty"`

	//PrincipalType: The principal type of the assigned principal ID.
	PrincipalType *RoleAssignmentProperties_PrincipalType_Status `json:"principalType,omitempty"`

	//RoleDefinitionId: The role definition ID.
	RoleDefinitionId *string `json:"roleDefinitionId,omitempty"`

	//Scope: The role assignment scope.
	Scope *string `json:"scope,omitempty"`

	//UpdatedBy: Id of the user who updated the assignment
	UpdatedBy *string `json:"updatedBy,omitempty"`

	//UpdatedOn: Time it was updated
	UpdatedOn *string `json:"updatedOn,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RoleAssignment_Status{}

// ConvertStatusFrom populates our RoleAssignment_Status from the provided source
func (assignment *RoleAssignment_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1alpha1api20200801previewstorage.RoleAssignment_Status)
	if ok {
		// Populate our instance from source
		return assignment.AssignPropertiesFromRoleAssignment_Status(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20200801previewstorage.RoleAssignment_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = assignment.AssignPropertiesFromRoleAssignment_Status(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our RoleAssignment_Status
func (assignment *RoleAssignment_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1alpha1api20200801previewstorage.RoleAssignment_Status)
	if ok {
		// Populate destination from our instance
		return assignment.AssignPropertiesToRoleAssignment_Status(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20200801previewstorage.RoleAssignment_Status{}
	err := assignment.AssignPropertiesToRoleAssignment_Status(dst)
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

var _ genruntime.FromARMConverter = &RoleAssignment_Status{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (assignment *RoleAssignment_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RoleAssignment_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (assignment *RoleAssignment_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RoleAssignment_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RoleAssignment_StatusARM, got %T", armInput)
	}

	// Set property ‘Condition’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Condition != nil {
			condition := *typedInput.Properties.Condition
			assignment.Condition = &condition
		}
	}

	// Set property ‘ConditionVersion’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ConditionVersion != nil {
			conditionVersion := *typedInput.Properties.ConditionVersion
			assignment.ConditionVersion = &conditionVersion
		}
	}

	// no assignment for property ‘Conditions’

	// Set property ‘CreatedBy’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.CreatedBy != nil {
			createdBy := *typedInput.Properties.CreatedBy
			assignment.CreatedBy = &createdBy
		}
	}

	// Set property ‘CreatedOn’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.CreatedOn != nil {
			createdOn := *typedInput.Properties.CreatedOn
			assignment.CreatedOn = &createdOn
		}
	}

	// Set property ‘DelegatedManagedIdentityResourceId’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.DelegatedManagedIdentityResourceId != nil {
			delegatedManagedIdentityResourceId := *typedInput.Properties.DelegatedManagedIdentityResourceId
			assignment.DelegatedManagedIdentityResourceId = &delegatedManagedIdentityResourceId
		}
	}

	// Set property ‘Description’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Description != nil {
			description := *typedInput.Properties.Description
			assignment.Description = &description
		}
	}

	// Set property ‘PrincipalId’:
	// copying flattened property:
	if typedInput.Properties != nil {
		assignment.PrincipalId = &typedInput.Properties.PrincipalId
	}

	// Set property ‘PrincipalType’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.PrincipalType != nil {
			principalType := *typedInput.Properties.PrincipalType
			assignment.PrincipalType = &principalType
		}
	}

	// Set property ‘RoleDefinitionId’:
	// copying flattened property:
	if typedInput.Properties != nil {
		assignment.RoleDefinitionId = &typedInput.Properties.RoleDefinitionId
	}

	// Set property ‘Scope’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Scope != nil {
			scope := *typedInput.Properties.Scope
			assignment.Scope = &scope
		}
	}

	// Set property ‘UpdatedBy’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.UpdatedBy != nil {
			updatedBy := *typedInput.Properties.UpdatedBy
			assignment.UpdatedBy = &updatedBy
		}
	}

	// Set property ‘UpdatedOn’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.UpdatedOn != nil {
			updatedOn := *typedInput.Properties.UpdatedOn
			assignment.UpdatedOn = &updatedOn
		}
	}

	// No error
	return nil
}

// AssignPropertiesFromRoleAssignment_Status populates our RoleAssignment_Status from the provided source RoleAssignment_Status
func (assignment *RoleAssignment_Status) AssignPropertiesFromRoleAssignment_Status(source *v1alpha1api20200801previewstorage.RoleAssignment_Status) error {

	// Condition
	assignment.Condition = genruntime.ClonePointerToString(source.Condition)

	// ConditionVersion
	assignment.ConditionVersion = genruntime.ClonePointerToString(source.ConditionVersion)

	// Conditions
	assignment.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// CreatedBy
	assignment.CreatedBy = genruntime.ClonePointerToString(source.CreatedBy)

	// CreatedOn
	assignment.CreatedOn = genruntime.ClonePointerToString(source.CreatedOn)

	// DelegatedManagedIdentityResourceId
	assignment.DelegatedManagedIdentityResourceId = genruntime.ClonePointerToString(source.DelegatedManagedIdentityResourceId)

	// Description
	assignment.Description = genruntime.ClonePointerToString(source.Description)

	// PrincipalId
	assignment.PrincipalId = genruntime.ClonePointerToString(source.PrincipalId)

	// PrincipalType
	if source.PrincipalType != nil {
		principalType := RoleAssignmentProperties_PrincipalType_Status(*source.PrincipalType)
		assignment.PrincipalType = &principalType
	} else {
		assignment.PrincipalType = nil
	}

	// RoleDefinitionId
	assignment.RoleDefinitionId = genruntime.ClonePointerToString(source.RoleDefinitionId)

	// Scope
	assignment.Scope = genruntime.ClonePointerToString(source.Scope)

	// UpdatedBy
	assignment.UpdatedBy = genruntime.ClonePointerToString(source.UpdatedBy)

	// UpdatedOn
	assignment.UpdatedOn = genruntime.ClonePointerToString(source.UpdatedOn)

	// No error
	return nil
}

// AssignPropertiesToRoleAssignment_Status populates the provided destination RoleAssignment_Status from our RoleAssignment_Status
func (assignment *RoleAssignment_Status) AssignPropertiesToRoleAssignment_Status(destination *v1alpha1api20200801previewstorage.RoleAssignment_Status) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Condition
	destination.Condition = genruntime.ClonePointerToString(assignment.Condition)

	// ConditionVersion
	destination.ConditionVersion = genruntime.ClonePointerToString(assignment.ConditionVersion)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(assignment.Conditions)

	// CreatedBy
	destination.CreatedBy = genruntime.ClonePointerToString(assignment.CreatedBy)

	// CreatedOn
	destination.CreatedOn = genruntime.ClonePointerToString(assignment.CreatedOn)

	// DelegatedManagedIdentityResourceId
	destination.DelegatedManagedIdentityResourceId = genruntime.ClonePointerToString(assignment.DelegatedManagedIdentityResourceId)

	// Description
	destination.Description = genruntime.ClonePointerToString(assignment.Description)

	// PrincipalId
	destination.PrincipalId = genruntime.ClonePointerToString(assignment.PrincipalId)

	// PrincipalType
	if assignment.PrincipalType != nil {
		principalType := string(*assignment.PrincipalType)
		destination.PrincipalType = &principalType
	} else {
		destination.PrincipalType = nil
	}

	// RoleDefinitionId
	destination.RoleDefinitionId = genruntime.ClonePointerToString(assignment.RoleDefinitionId)

	// Scope
	destination.Scope = genruntime.ClonePointerToString(assignment.Scope)

	// UpdatedBy
	destination.UpdatedBy = genruntime.ClonePointerToString(assignment.UpdatedBy)

	// UpdatedOn
	destination.UpdatedOn = genruntime.ClonePointerToString(assignment.UpdatedOn)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type RoleAssignments_SPEC struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName string `json:"azureName"`

	//Condition: The conditions on the role assignment. This limits the resources it
	//can be assigned to. e.g.:
	//@Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName]
	//StringEqualsIgnoreCase 'foo_storage_container'
	Condition *string `json:"condition,omitempty"`

	//ConditionVersion: Version of the condition. Currently accepted value is '2.0'
	ConditionVersion *string `json:"conditionVersion,omitempty"`

	//DelegatedManagedIdentityResourceId: Id of the delegated managed identity resource
	DelegatedManagedIdentityResourceId *string `json:"delegatedManagedIdentityResourceId,omitempty"`

	//Description: Description of role assignment
	Description *string `json:"description,omitempty"`

	// +kubebuilder:validation:Required
	Owner genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`

	// +kubebuilder:validation:Required
	//PrincipalId: The principal ID.
	PrincipalId string `json:"principalId"`

	//PrincipalType: The principal type of the assigned principal ID.
	PrincipalType *RoleAssignmentProperties_PrincipalType_Spec `json:"principalType,omitempty"`

	// +kubebuilder:validation:Required
	//RoleDefinitionId: The role definition ID.
	RoleDefinitionId string `json:"roleDefinitionId"`
}

var _ genruntime.ARMTransformer = &RoleAssignments_SPEC{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (spec *RoleAssignments_SPEC) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if spec == nil {
		return nil, nil
	}
	var result RoleAssignments_SPECARM

	// Set property ‘AzureName’:
	result.AzureName = spec.AzureName

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if spec.Condition != nil {
		condition := *spec.Condition
		result.Properties.Condition = &condition
	}
	if spec.ConditionVersion != nil {
		conditionVersion := *spec.ConditionVersion
		result.Properties.ConditionVersion = &conditionVersion
	}
	if spec.DelegatedManagedIdentityResourceId != nil {
		delegatedManagedIdentityResourceId := *spec.DelegatedManagedIdentityResourceId
		result.Properties.DelegatedManagedIdentityResourceId = &delegatedManagedIdentityResourceId
	}
	if spec.Description != nil {
		description := *spec.Description
		result.Properties.Description = &description
	}
	result.Properties.PrincipalId = spec.PrincipalId
	if spec.PrincipalType != nil {
		principalType := *spec.PrincipalType
		result.Properties.PrincipalType = &principalType
	}
	result.Properties.RoleDefinitionId = spec.RoleDefinitionId
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (spec *RoleAssignments_SPEC) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RoleAssignments_SPECARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (spec *RoleAssignments_SPEC) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RoleAssignments_SPECARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RoleAssignments_SPECARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	spec.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Condition’:
	// copying flattened property:
	if typedInput.Properties.Condition != nil {
		condition := *typedInput.Properties.Condition
		spec.Condition = &condition
	}

	// Set property ‘ConditionVersion’:
	// copying flattened property:
	if typedInput.Properties.ConditionVersion != nil {
		conditionVersion := *typedInput.Properties.ConditionVersion
		spec.ConditionVersion = &conditionVersion
	}

	// Set property ‘DelegatedManagedIdentityResourceId’:
	// copying flattened property:
	if typedInput.Properties.DelegatedManagedIdentityResourceId != nil {
		delegatedManagedIdentityResourceId := *typedInput.Properties.DelegatedManagedIdentityResourceId
		spec.DelegatedManagedIdentityResourceId = &delegatedManagedIdentityResourceId
	}

	// Set property ‘Description’:
	// copying flattened property:
	if typedInput.Properties.Description != nil {
		description := *typedInput.Properties.Description
		spec.Description = &description
	}

	// Set property ‘Owner’:
	spec.Owner = genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘PrincipalId’:
	// copying flattened property:
	spec.PrincipalId = typedInput.Properties.PrincipalId

	// Set property ‘PrincipalType’:
	// copying flattened property:
	if typedInput.Properties.PrincipalType != nil {
		principalType := *typedInput.Properties.PrincipalType
		spec.PrincipalType = &principalType
	}

	// Set property ‘RoleDefinitionId’:
	// copying flattened property:
	spec.RoleDefinitionId = typedInput.Properties.RoleDefinitionId

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &RoleAssignments_SPEC{}

// ConvertSpecFrom populates our RoleAssignments_SPEC from the provided source
func (spec *RoleAssignments_SPEC) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1alpha1api20200801previewstorage.RoleAssignments_SPEC)
	if ok {
		// Populate our instance from source
		return spec.AssignPropertiesFromRoleAssignments_SPEC(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20200801previewstorage.RoleAssignments_SPEC{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = spec.AssignPropertiesFromRoleAssignments_SPEC(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our RoleAssignments_SPEC
func (spec *RoleAssignments_SPEC) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1alpha1api20200801previewstorage.RoleAssignments_SPEC)
	if ok {
		// Populate destination from our instance
		return spec.AssignPropertiesToRoleAssignments_SPEC(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20200801previewstorage.RoleAssignments_SPEC{}
	err := spec.AssignPropertiesToRoleAssignments_SPEC(dst)
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

// AssignPropertiesFromRoleAssignments_SPEC populates our RoleAssignments_SPEC from the provided source RoleAssignments_SPEC
func (spec *RoleAssignments_SPEC) AssignPropertiesFromRoleAssignments_SPEC(source *v1alpha1api20200801previewstorage.RoleAssignments_SPEC) error {

	// AzureName
	spec.AzureName = source.AzureName

	// Condition
	spec.Condition = genruntime.ClonePointerToString(source.Condition)

	// ConditionVersion
	spec.ConditionVersion = genruntime.ClonePointerToString(source.ConditionVersion)

	// DelegatedManagedIdentityResourceId
	spec.DelegatedManagedIdentityResourceId = genruntime.ClonePointerToString(source.DelegatedManagedIdentityResourceId)

	// Description
	spec.Description = genruntime.ClonePointerToString(source.Description)

	// Owner
	spec.Owner = source.Owner.Copy()

	// PrincipalId
	spec.PrincipalId = genruntime.GetOptionalStringValue(source.PrincipalId)

	// PrincipalType
	if source.PrincipalType != nil {
		principalType := RoleAssignmentProperties_PrincipalType_Spec(*source.PrincipalType)
		spec.PrincipalType = &principalType
	} else {
		spec.PrincipalType = nil
	}

	// RoleDefinitionId
	spec.RoleDefinitionId = genruntime.GetOptionalStringValue(source.RoleDefinitionId)

	// No error
	return nil
}

// AssignPropertiesToRoleAssignments_SPEC populates the provided destination RoleAssignments_SPEC from our RoleAssignments_SPEC
func (spec *RoleAssignments_SPEC) AssignPropertiesToRoleAssignments_SPEC(destination *v1alpha1api20200801previewstorage.RoleAssignments_SPEC) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = spec.AzureName

	// Condition
	destination.Condition = genruntime.ClonePointerToString(spec.Condition)

	// ConditionVersion
	destination.ConditionVersion = genruntime.ClonePointerToString(spec.ConditionVersion)

	// DelegatedManagedIdentityResourceId
	destination.DelegatedManagedIdentityResourceId = genruntime.ClonePointerToString(spec.DelegatedManagedIdentityResourceId)

	// Description
	destination.Description = genruntime.ClonePointerToString(spec.Description)

	// OriginalVersion
	destination.OriginalVersion = spec.OriginalVersion()

	// Owner
	destination.Owner = spec.Owner.Copy()

	// PrincipalId
	principalId := spec.PrincipalId
	destination.PrincipalId = &principalId

	// PrincipalType
	if spec.PrincipalType != nil {
		principalType := string(*spec.PrincipalType)
		destination.PrincipalType = &principalType
	} else {
		destination.PrincipalType = nil
	}

	// RoleDefinitionId
	roleDefinitionId := spec.RoleDefinitionId
	destination.RoleDefinitionId = &roleDefinitionId

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
func (spec *RoleAssignments_SPEC) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (spec *RoleAssignments_SPEC) SetAzureName(azureName string) { spec.AzureName = azureName }

// +kubebuilder:validation:Enum={"ForeignGroup","Group","ServicePrincipal","User"}
type RoleAssignmentProperties_PrincipalType_Spec string

const (
	RoleAssignmentProperties_PrincipalType_SpecForeignGroup     = RoleAssignmentProperties_PrincipalType_Spec("ForeignGroup")
	RoleAssignmentProperties_PrincipalType_SpecGroup            = RoleAssignmentProperties_PrincipalType_Spec("Group")
	RoleAssignmentProperties_PrincipalType_SpecServicePrincipal = RoleAssignmentProperties_PrincipalType_Spec("ServicePrincipal")
	RoleAssignmentProperties_PrincipalType_SpecUser             = RoleAssignmentProperties_PrincipalType_Spec("User")
)

type RoleAssignmentProperties_PrincipalType_Status string

const (
	RoleAssignmentProperties_PrincipalType_StatusForeignGroup     = RoleAssignmentProperties_PrincipalType_Status("ForeignGroup")
	RoleAssignmentProperties_PrincipalType_StatusGroup            = RoleAssignmentProperties_PrincipalType_Status("Group")
	RoleAssignmentProperties_PrincipalType_StatusServicePrincipal = RoleAssignmentProperties_PrincipalType_Status("ServicePrincipal")
	RoleAssignmentProperties_PrincipalType_StatusUser             = RoleAssignmentProperties_PrincipalType_Status("User")
)

func init() {
	SchemeBuilder.Register(&RoleAssignment{}, &RoleAssignmentList{})
}

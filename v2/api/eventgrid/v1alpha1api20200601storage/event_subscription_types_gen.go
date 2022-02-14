// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=eventgrid.azure.com,resources=eventsubscriptions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=eventgrid.azure.com,resources={eventsubscriptions/status,eventsubscriptions/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20200601.EventSubscription
//Generator information:
//- Generated from: /eventgrid/resource-manager/Microsoft.EventGrid/stable/2020-06-01/EventGrid.json
//- ARM URI: /{scope}/providers/Microsoft.EventGrid/eventSubscriptions/{eventSubscriptionName}
type EventSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventSubscription_Spec   `json:"spec,omitempty"`
	Status            EventSubscription_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &EventSubscription{}

// GetConditions returns the conditions of the resource
func (subscription *EventSubscription) GetConditions() conditions.Conditions {
	return subscription.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (subscription *EventSubscription) SetConditions(conditions conditions.Conditions) {
	subscription.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &EventSubscription{}

// AzureName returns the Azure name of the resource
func (subscription *EventSubscription) AzureName() string {
	return subscription.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (subscription EventSubscription) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (subscription *EventSubscription) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (subscription *EventSubscription) GetSpec() genruntime.ConvertibleSpec {
	return &subscription.Spec
}

// GetStatus returns the status of this resource
func (subscription *EventSubscription) GetStatus() genruntime.ConvertibleStatus {
	return &subscription.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventGrid/eventSubscriptions"
func (subscription *EventSubscription) GetType() string {
	return "Microsoft.EventGrid/eventSubscriptions"
}

// NewEmptyStatus returns a new empty (blank) status
func (subscription *EventSubscription) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &EventSubscription_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (subscription *EventSubscription) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(subscription.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  subscription.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (subscription *EventSubscription) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*EventSubscription_Status); ok {
		subscription.Status = *st
		return nil
	}

	// Convert status to required version
	var st EventSubscription_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	subscription.Status = st
	return nil
}

// Hub marks that this EventSubscription is the hub type for conversion
func (subscription *EventSubscription) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (subscription *EventSubscription) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: subscription.Spec.OriginalVersion,
		Kind:    "EventSubscription",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20200601.EventSubscription
//Generator information:
//- Generated from: /eventgrid/resource-manager/Microsoft.EventGrid/stable/2020-06-01/EventGrid.json
//- ARM URI: /{scope}/providers/Microsoft.EventGrid/eventSubscriptions/{eventSubscriptionName}
type EventSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventSubscription `json:"items"`
}

//Storage version of v1alpha1api20200601.EventSubscription_Spec
type EventSubscription_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName             string                        `json:"azureName"`
	DeadLetterDestination *DeadLetterDestination        `json:"deadLetterDestination,omitempty"`
	Destination           *EventSubscriptionDestination `json:"destination,omitempty"`
	EventDeliverySchema   *string                       `json:"eventDeliverySchema,omitempty"`
	ExpirationTimeUtc     *string                       `json:"expirationTimeUtc,omitempty"`
	Filter                *EventSubscriptionFilter      `json:"filter,omitempty"`
	Labels                []string                      `json:"labels,omitempty"`
	OriginalVersion       string                        `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	RetryPolicy *RetryPolicy                      `json:"retryPolicy,omitempty"`
}

var _ genruntime.ConvertibleSpec = &EventSubscription_Spec{}

// ConvertSpecFrom populates our EventSubscription_Spec from the provided source
func (subscription *EventSubscription_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == subscription {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(subscription)
}

// ConvertSpecTo populates the provided destination from our EventSubscription_Spec
func (subscription *EventSubscription_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == subscription {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(subscription)
}

//Storage version of v1alpha1api20200601.EventSubscription_Status
type EventSubscription_Status struct {
	Conditions            []conditions.Condition               `json:"conditions,omitempty"`
	DeadLetterDestination *DeadLetterDestination_Status        `json:"deadLetterDestination,omitempty"`
	Destination           *EventSubscriptionDestination_Status `json:"destination,omitempty"`
	EventDeliverySchema   *string                              `json:"eventDeliverySchema,omitempty"`
	ExpirationTimeUtc     *string                              `json:"expirationTimeUtc,omitempty"`
	Filter                *EventSubscriptionFilter_Status      `json:"filter,omitempty"`
	Id                    *string                              `json:"id,omitempty"`
	Labels                []string                             `json:"labels,omitempty"`
	Name                  *string                              `json:"name,omitempty"`
	PropertyBag           genruntime.PropertyBag               `json:"$propertyBag,omitempty"`
	ProvisioningState     *string                              `json:"provisioningState,omitempty"`
	RetryPolicy           *RetryPolicy_Status                  `json:"retryPolicy,omitempty"`
	SystemData            *SystemData_Status                   `json:"systemData,omitempty"`
	Topic                 *string                              `json:"topic,omitempty"`
	Type                  *string                              `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &EventSubscription_Status{}

// ConvertStatusFrom populates our EventSubscription_Status from the provided source
func (subscription *EventSubscription_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == subscription {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(subscription)
}

// ConvertStatusTo populates the provided destination from our EventSubscription_Status
func (subscription *EventSubscription_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == subscription {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(subscription)
}

//Storage version of v1alpha1api20200601.DeadLetterDestination
type DeadLetterDestination struct {
	EndpointType *string                `json:"endpointType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20200601.DeadLetterDestination_Status
type DeadLetterDestination_Status struct {
	EndpointType *string                `json:"endpointType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20200601.EventSubscriptionDestination
type EventSubscriptionDestination struct {
	EndpointType *string                `json:"endpointType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20200601.EventSubscriptionDestination_Status
type EventSubscriptionDestination_Status struct {
	EndpointType *string                `json:"endpointType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20200601.EventSubscriptionFilter
type EventSubscriptionFilter struct {
	AdvancedFilters        []AdvancedFilter       `json:"advancedFilters,omitempty"`
	IncludedEventTypes     []string               `json:"includedEventTypes,omitempty"`
	IsSubjectCaseSensitive *bool                  `json:"isSubjectCaseSensitive,omitempty"`
	PropertyBag            genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SubjectBeginsWith      *string                `json:"subjectBeginsWith,omitempty"`
	SubjectEndsWith        *string                `json:"subjectEndsWith,omitempty"`
}

//Storage version of v1alpha1api20200601.EventSubscriptionFilter_Status
type EventSubscriptionFilter_Status struct {
	AdvancedFilters        []AdvancedFilter_Status `json:"advancedFilters,omitempty"`
	IncludedEventTypes     []string                `json:"includedEventTypes,omitempty"`
	IsSubjectCaseSensitive *bool                   `json:"isSubjectCaseSensitive,omitempty"`
	PropertyBag            genruntime.PropertyBag  `json:"$propertyBag,omitempty"`
	SubjectBeginsWith      *string                 `json:"subjectBeginsWith,omitempty"`
	SubjectEndsWith        *string                 `json:"subjectEndsWith,omitempty"`
}

//Storage version of v1alpha1api20200601.RetryPolicy
type RetryPolicy struct {
	EventTimeToLiveInMinutes *int                   `json:"eventTimeToLiveInMinutes,omitempty"`
	MaxDeliveryAttempts      *int                   `json:"maxDeliveryAttempts,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20200601.RetryPolicy_Status
type RetryPolicy_Status struct {
	EventTimeToLiveInMinutes *int                   `json:"eventTimeToLiveInMinutes,omitempty"`
	MaxDeliveryAttempts      *int                   `json:"maxDeliveryAttempts,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20200601.AdvancedFilter
type AdvancedFilter struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20200601.AdvancedFilter_Status
type AdvancedFilter_Status struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&EventSubscription{}, &EventSubscriptionList{})
}

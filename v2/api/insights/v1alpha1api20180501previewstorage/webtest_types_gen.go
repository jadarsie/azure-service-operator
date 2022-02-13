// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20180501previewstorage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=insights.azure.com,resources=webtests,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=insights.azure.com,resources={webtests/status,webtests/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20180501preview.Webtest
//Generator information:
//- Generated from: /applicationinsights/resource-manager/Microsoft.Insights/preview/2018-05-01-preview/webTests_API.json
//- ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests/{webTestName}
type Webtest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Webtests_SPEC  `json:"spec,omitempty"`
	Status            WebTest_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Webtest{}

// GetConditions returns the conditions of the resource
func (webtest *Webtest) GetConditions() conditions.Conditions {
	return webtest.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (webtest *Webtest) SetConditions(conditions conditions.Conditions) {
	webtest.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &Webtest{}

// AzureName returns the Azure name of the resource
func (webtest *Webtest) AzureName() string {
	return webtest.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-05-01-preview"
func (webtest Webtest) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (webtest *Webtest) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (webtest *Webtest) GetSpec() genruntime.ConvertibleSpec {
	return &webtest.Spec
}

// GetStatus returns the status of this resource
func (webtest *Webtest) GetStatus() genruntime.ConvertibleStatus {
	return &webtest.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (webtest *Webtest) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (webtest *Webtest) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &WebTest_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (webtest *Webtest) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(webtest.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  webtest.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (webtest *Webtest) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*WebTest_Status); ok {
		webtest.Status = *st
		return nil
	}

	// Convert status to required version
	var st WebTest_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	webtest.Status = st
	return nil
}

// Hub marks that this Webtest is the hub type for conversion
func (webtest *Webtest) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (webtest *Webtest) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: webtest.Spec.OriginalVersion,
		Kind:    "Webtest",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20180501preview.Webtest
//Generator information:
//- Generated from: /applicationinsights/resource-manager/Microsoft.Insights/preview/2018-05-01-preview/webTests_API.json
//- ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests/{webTestName}
type WebtestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Webtest `json:"items"`
}

// +kubebuilder:validation:Enum={"2018-05-01-preview"}
type APIVersion string

const APIVersionValue = APIVersion("2018-05-01-preview")

//Storage version of v1alpha1api20180501preview.WebTest_Status
type WebTest_Status struct {
	Conditions         []conditions.Condition                    `json:"conditions,omitempty"`
	Configuration      *WebTestProperties_Configuration_Status   `json:"Configuration,omitempty"`
	Description        *string                                   `json:"Description,omitempty"`
	Enabled            *bool                                     `json:"Enabled,omitempty"`
	Frequency          *int                                      `json:"Frequency,omitempty"`
	Id                 *string                                   `json:"id,omitempty"`
	Kind               *string                                   `json:"Kind,omitempty"`
	Location           *string                                   `json:"location,omitempty"`
	Locations          []WebTestGeolocation_Status               `json:"Locations,omitempty"`
	Name               *string                                   `json:"name,omitempty"`
	PropertiesName     *string                                   `json:"properties_name,omitempty"`
	PropertyBag        genruntime.PropertyBag                    `json:"$propertyBag,omitempty"`
	ProvisioningState  *string                                   `json:"provisioningState,omitempty"`
	Request            *WebTestProperties_Request_Status         `json:"Request,omitempty"`
	RetryEnabled       *bool                                     `json:"RetryEnabled,omitempty"`
	SyntheticMonitorId *string                                   `json:"SyntheticMonitorId,omitempty"`
	Tags               *v1.JSON                                  `json:"tags,omitempty"`
	Timeout            *int                                      `json:"Timeout,omitempty"`
	Type               *string                                   `json:"type,omitempty"`
	ValidationRules    *WebTestProperties_ValidationRules_Status `json:"ValidationRules,omitempty"`
}

var _ genruntime.ConvertibleStatus = &WebTest_Status{}

// ConvertStatusFrom populates our WebTest_Status from the provided source
func (test *WebTest_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == test {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(test)
}

// ConvertStatusTo populates the provided destination from our WebTest_Status
func (test *WebTest_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == test {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(test)
}

//Storage version of v1alpha1api20180501preview.Webtests_SPEC
type Webtests_SPEC struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName       string                                `json:"azureName"`
	Configuration   *WebTestProperties_Configuration_Spec `json:"Configuration,omitempty"`
	Description     *string                               `json:"Description,omitempty"`
	Enabled         *bool                                 `json:"Enabled,omitempty"`
	Frequency       *int                                  `json:"Frequency,omitempty"`
	Kind            *string                               `json:"Kind,omitempty"`
	Location        *string                               `json:"location,omitempty"`
	Locations       []WebTestGeolocation_Spec             `json:"Locations,omitempty"`
	Name            *string                               `json:"Name,omitempty"`
	OriginalVersion string                                `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner              genruntime.KnownResourceReference       `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PropertyBag        genruntime.PropertyBag                  `json:"$propertyBag,omitempty"`
	Request            *WebTestProperties_Request_Spec         `json:"Request,omitempty"`
	RetryEnabled       *bool                                   `json:"RetryEnabled,omitempty"`
	SyntheticMonitorId *string                                 `json:"SyntheticMonitorId,omitempty"`
	Tags               *v1.JSON                                `json:"tags,omitempty"`
	Timeout            *int                                    `json:"Timeout,omitempty"`
	ValidationRules    *WebTestProperties_ValidationRules_Spec `json:"ValidationRules,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Webtests_SPEC{}

// ConvertSpecFrom populates our Webtests_SPEC from the provided source
func (spec *Webtests_SPEC) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(spec)
}

// ConvertSpecTo populates the provided destination from our Webtests_SPEC
func (spec *Webtests_SPEC) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(spec)
}

//Storage version of v1alpha1api20180501preview.WebTestGeolocation_Spec
type WebTestGeolocation_Spec struct {
	Id          *string                `json:"Id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20180501preview.WebTestGeolocation_Status
type WebTestGeolocation_Status struct {
	Id          *string                `json:"Id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20180501preview.WebTestProperties_Configuration_Spec
type WebTestProperties_Configuration_Spec struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	WebTest     *string                `json:"WebTest,omitempty"`
}

//Storage version of v1alpha1api20180501preview.WebTestProperties_Configuration_Status
type WebTestProperties_Configuration_Status struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	WebTest     *string                `json:"WebTest,omitempty"`
}

//Storage version of v1alpha1api20180501preview.WebTestProperties_Request_Spec
type WebTestProperties_Request_Spec struct {
	FollowRedirects        *bool                  `json:"FollowRedirects,omitempty"`
	Headers                []HeaderField_Spec     `json:"Headers,omitempty"`
	HttpVerb               *string                `json:"HttpVerb,omitempty"`
	ParseDependentRequests *bool                  `json:"ParseDependentRequests,omitempty"`
	PropertyBag            genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RequestBody            *string                `json:"RequestBody,omitempty"`
	RequestUrl             *string                `json:"RequestUrl,omitempty"`
}

//Storage version of v1alpha1api20180501preview.WebTestProperties_Request_Status
type WebTestProperties_Request_Status struct {
	FollowRedirects        *bool                  `json:"FollowRedirects,omitempty"`
	Headers                []HeaderField_Status   `json:"Headers,omitempty"`
	HttpVerb               *string                `json:"HttpVerb,omitempty"`
	ParseDependentRequests *bool                  `json:"ParseDependentRequests,omitempty"`
	PropertyBag            genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RequestBody            *string                `json:"RequestBody,omitempty"`
	RequestUrl             *string                `json:"RequestUrl,omitempty"`
}

//Storage version of v1alpha1api20180501preview.WebTestProperties_ValidationRules_Spec
type WebTestProperties_ValidationRules_Spec struct {
	ContentValidation             *WebTestProperties_ValidationRules_ContentValidation_Spec `json:"ContentValidation,omitempty"`
	ExpectedHttpStatusCode        *int                                                      `json:"ExpectedHttpStatusCode,omitempty"`
	IgnoreHttpsStatusCode         *bool                                                     `json:"IgnoreHttpsStatusCode,omitempty"`
	PropertyBag                   genruntime.PropertyBag                                    `json:"$propertyBag,omitempty"`
	SSLCertRemainingLifetimeCheck *int                                                      `json:"SSLCertRemainingLifetimeCheck,omitempty"`
	SSLCheck                      *bool                                                     `json:"SSLCheck,omitempty"`
}

//Storage version of v1alpha1api20180501preview.WebTestProperties_ValidationRules_Status
type WebTestProperties_ValidationRules_Status struct {
	ContentValidation             *WebTestProperties_ValidationRules_ContentValidation_Status `json:"ContentValidation,omitempty"`
	ExpectedHttpStatusCode        *int                                                        `json:"ExpectedHttpStatusCode,omitempty"`
	IgnoreHttpsStatusCode         *bool                                                       `json:"IgnoreHttpsStatusCode,omitempty"`
	PropertyBag                   genruntime.PropertyBag                                      `json:"$propertyBag,omitempty"`
	SSLCertRemainingLifetimeCheck *int                                                        `json:"SSLCertRemainingLifetimeCheck,omitempty"`
	SSLCheck                      *bool                                                       `json:"SSLCheck,omitempty"`
}

//Storage version of v1alpha1api20180501preview.HeaderField_Spec
type HeaderField_Spec struct {
	Key         *string                `json:"key,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

//Storage version of v1alpha1api20180501preview.HeaderField_Status
type HeaderField_Status struct {
	Key         *string                `json:"key,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

//Storage version of v1alpha1api20180501preview.WebTestProperties_ValidationRules_ContentValidation_Spec
type WebTestProperties_ValidationRules_ContentValidation_Spec struct {
	ContentMatch    *string                `json:"ContentMatch,omitempty"`
	IgnoreCase      *bool                  `json:"IgnoreCase,omitempty"`
	PassIfTextFound *bool                  `json:"PassIfTextFound,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20180501preview.WebTestProperties_ValidationRules_ContentValidation_Status
type WebTestProperties_ValidationRules_ContentValidation_Status struct {
	ContentMatch    *string                `json:"ContentMatch,omitempty"`
	IgnoreCase      *bool                  `json:"IgnoreCase,omitempty"`
	PassIfTextFound *bool                  `json:"PassIfTextFound,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Webtest{}, &WebtestList{})
}

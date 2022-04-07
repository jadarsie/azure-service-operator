// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211001

//Deprecated version of SignalRResource_Status. Use v1beta20211001.SignalRResource_Status instead
type SignalRResource_StatusARM struct {
	Id         *string                      `json:"id,omitempty"`
	Identity   *ManagedIdentity_StatusARM   `json:"identity,omitempty"`
	Kind       *ServiceKind_Status          `json:"kind,omitempty"`
	Location   *string                      `json:"location,omitempty"`
	Name       *string                      `json:"name,omitempty"`
	Properties *SignalRProperties_StatusARM `json:"properties,omitempty"`
	Sku        *ResourceSku_StatusARM       `json:"sku,omitempty"`
	SystemData *SystemData_StatusARM        `json:"systemData,omitempty"`
	Tags       map[string]string            `json:"tags,omitempty"`
	Type       *string                      `json:"type,omitempty"`
}

//Deprecated version of ManagedIdentity_Status. Use v1beta20211001.ManagedIdentity_Status instead
type ManagedIdentity_StatusARM struct {
	PrincipalId            *string                                           `json:"principalId,omitempty"`
	TenantId               *string                                           `json:"tenantId,omitempty"`
	Type                   *ManagedIdentityType_Status                       `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentityProperty_StatusARM `json:"userAssignedIdentities,omitempty"`
}

//Deprecated version of ResourceSku_Status. Use v1beta20211001.ResourceSku_Status instead
type ResourceSku_StatusARM struct {
	Capacity *int                   `json:"capacity,omitempty"`
	Family   *string                `json:"family,omitempty"`
	Name     *string                `json:"name,omitempty"`
	Size     *string                `json:"size,omitempty"`
	Tier     *SignalRSkuTier_Status `json:"tier,omitempty"`
}

//Deprecated version of ServiceKind_Status. Use v1beta20211001.ServiceKind_Status instead
type ServiceKind_Status string

const (
	ServiceKind_StatusRawWebSockets = ServiceKind_Status("RawWebSockets")
	ServiceKind_StatusSignalR       = ServiceKind_Status("SignalR")
)

//Deprecated version of SignalRProperties_Status. Use v1beta20211001.SignalRProperties_Status instead
type SignalRProperties_StatusARM struct {
	Cors                       *SignalRCorsSettings_StatusARM                                    `json:"cors,omitempty"`
	DisableAadAuth             *bool                                                             `json:"disableAadAuth,omitempty"`
	DisableLocalAuth           *bool                                                             `json:"disableLocalAuth,omitempty"`
	ExternalIP                 *string                                                           `json:"externalIP,omitempty"`
	Features                   []SignalRFeature_StatusARM                                        `json:"features,omitempty"`
	HostName                   *string                                                           `json:"hostName,omitempty"`
	HostNamePrefix             *string                                                           `json:"hostNamePrefix,omitempty"`
	NetworkACLs                *SignalRNetworkACLs_StatusARM                                     `json:"networkACLs,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_Status_SignalR_SubResourceEmbeddedARM `json:"privateEndpointConnections,omitempty"`
	ProvisioningState          *ProvisioningState_Status                                         `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *string                                                           `json:"publicNetworkAccess,omitempty"`
	PublicPort                 *int                                                              `json:"publicPort,omitempty"`
	ResourceLogConfiguration   *ResourceLogConfiguration_StatusARM                               `json:"resourceLogConfiguration,omitempty"`
	ServerPort                 *int                                                              `json:"serverPort,omitempty"`
	SharedPrivateLinkResources []SharedPrivateLinkResource_Status_SignalR_SubResourceEmbeddedARM `json:"sharedPrivateLinkResources,omitempty"`
	Tls                        *SignalRTlsSettings_StatusARM                                     `json:"tls,omitempty"`
	Upstream                   *ServerlessUpstreamSettings_StatusARM                             `json:"upstream,omitempty"`
	Version                    *string                                                           `json:"version,omitempty"`
}

//Deprecated version of SystemData_Status. Use v1beta20211001.SystemData_Status instead
type SystemData_StatusARM struct {
	CreatedAt          *string                             `json:"createdAt,omitempty"`
	CreatedBy          *string                             `json:"createdBy,omitempty"`
	CreatedByType      *SystemDataStatusCreatedByType      `json:"createdByType,omitempty"`
	LastModifiedAt     *string                             `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                             `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *SystemDataStatusLastModifiedByType `json:"lastModifiedByType,omitempty"`
}

//Deprecated version of ManagedIdentityType_Status. Use v1beta20211001.ManagedIdentityType_Status instead
type ManagedIdentityType_Status string

const (
	ManagedIdentityType_StatusNone           = ManagedIdentityType_Status("None")
	ManagedIdentityType_StatusSystemAssigned = ManagedIdentityType_Status("SystemAssigned")
	ManagedIdentityType_StatusUserAssigned   = ManagedIdentityType_Status("UserAssigned")
)

//Deprecated version of PrivateEndpointConnection_Status_SignalR_SubResourceEmbedded. Use v1beta20211001.PrivateEndpointConnection_Status_SignalR_SubResourceEmbedded instead
type PrivateEndpointConnection_Status_SignalR_SubResourceEmbeddedARM struct {
	Id         *string               `json:"id,omitempty"`
	SystemData *SystemData_StatusARM `json:"systemData,omitempty"`
}

//Deprecated version of ResourceLogConfiguration_Status. Use v1beta20211001.ResourceLogConfiguration_Status instead
type ResourceLogConfiguration_StatusARM struct {
	Categories []ResourceLogCategory_StatusARM `json:"categories,omitempty"`
}

//Deprecated version of ServerlessUpstreamSettings_Status. Use v1beta20211001.ServerlessUpstreamSettings_Status instead
type ServerlessUpstreamSettings_StatusARM struct {
	Templates []UpstreamTemplate_StatusARM `json:"templates,omitempty"`
}

//Deprecated version of SharedPrivateLinkResource_Status_SignalR_SubResourceEmbedded. Use v1beta20211001.SharedPrivateLinkResource_Status_SignalR_SubResourceEmbedded instead
type SharedPrivateLinkResource_Status_SignalR_SubResourceEmbeddedARM struct {
	Id         *string               `json:"id,omitempty"`
	SystemData *SystemData_StatusARM `json:"systemData,omitempty"`
}

//Deprecated version of SignalRCorsSettings_Status. Use v1beta20211001.SignalRCorsSettings_Status instead
type SignalRCorsSettings_StatusARM struct {
	AllowedOrigins []string `json:"allowedOrigins,omitempty"`
}

//Deprecated version of SignalRFeature_Status. Use v1beta20211001.SignalRFeature_Status instead
type SignalRFeature_StatusARM struct {
	Flag       *FeatureFlags_Status `json:"flag,omitempty"`
	Properties map[string]string    `json:"properties,omitempty"`
	Value      *string              `json:"value,omitempty"`
}

//Deprecated version of SignalRNetworkACLs_Status. Use v1beta20211001.SignalRNetworkACLs_Status instead
type SignalRNetworkACLs_StatusARM struct {
	DefaultAction    *ACLAction_Status              `json:"defaultAction,omitempty"`
	PrivateEndpoints []PrivateEndpointACL_StatusARM `json:"privateEndpoints,omitempty"`
	PublicNetwork    *NetworkACL_StatusARM          `json:"publicNetwork,omitempty"`
}

//Deprecated version of SignalRSkuTier_Status. Use v1beta20211001.SignalRSkuTier_Status instead
type SignalRSkuTier_Status string

const (
	SignalRSkuTier_StatusBasic    = SignalRSkuTier_Status("Basic")
	SignalRSkuTier_StatusFree     = SignalRSkuTier_Status("Free")
	SignalRSkuTier_StatusPremium  = SignalRSkuTier_Status("Premium")
	SignalRSkuTier_StatusStandard = SignalRSkuTier_Status("Standard")
)

//Deprecated version of SignalRTlsSettings_Status. Use v1beta20211001.SignalRTlsSettings_Status instead
type SignalRTlsSettings_StatusARM struct {
	ClientCertEnabled *bool `json:"clientCertEnabled,omitempty"`
}

//Deprecated version of SystemDataStatusCreatedByType. Use v1beta20211001.SystemDataStatusCreatedByType instead
type SystemDataStatusCreatedByType string

const (
	SystemDataStatusCreatedByTypeApplication     = SystemDataStatusCreatedByType("Application")
	SystemDataStatusCreatedByTypeKey             = SystemDataStatusCreatedByType("Key")
	SystemDataStatusCreatedByTypeManagedIdentity = SystemDataStatusCreatedByType("ManagedIdentity")
	SystemDataStatusCreatedByTypeUser            = SystemDataStatusCreatedByType("User")
)

//Deprecated version of SystemDataStatusLastModifiedByType. Use v1beta20211001.SystemDataStatusLastModifiedByType instead
type SystemDataStatusLastModifiedByType string

const (
	SystemDataStatusLastModifiedByTypeApplication     = SystemDataStatusLastModifiedByType("Application")
	SystemDataStatusLastModifiedByTypeKey             = SystemDataStatusLastModifiedByType("Key")
	SystemDataStatusLastModifiedByTypeManagedIdentity = SystemDataStatusLastModifiedByType("ManagedIdentity")
	SystemDataStatusLastModifiedByTypeUser            = SystemDataStatusLastModifiedByType("User")
)

//Deprecated version of UserAssignedIdentityProperty_Status. Use v1beta20211001.UserAssignedIdentityProperty_Status instead
type UserAssignedIdentityProperty_StatusARM struct {
	ClientId    *string `json:"clientId,omitempty"`
	PrincipalId *string `json:"principalId,omitempty"`
}

//Deprecated version of NetworkACL_Status. Use v1beta20211001.NetworkACL_Status instead
type NetworkACL_StatusARM struct {
	Allow []SignalRRequestType_Status `json:"allow,omitempty"`
	Deny  []SignalRRequestType_Status `json:"deny,omitempty"`
}

//Deprecated version of PrivateEndpointACL_Status. Use v1beta20211001.PrivateEndpointACL_Status instead
type PrivateEndpointACL_StatusARM struct {
	Allow []SignalRRequestType_Status `json:"allow,omitempty"`
	Deny  []SignalRRequestType_Status `json:"deny,omitempty"`
	Name  *string                     `json:"name,omitempty"`
}

//Deprecated version of ResourceLogCategory_Status. Use v1beta20211001.ResourceLogCategory_Status instead
type ResourceLogCategory_StatusARM struct {
	Enabled *string `json:"enabled,omitempty"`
	Name    *string `json:"name,omitempty"`
}

//Deprecated version of UpstreamTemplate_Status. Use v1beta20211001.UpstreamTemplate_Status instead
type UpstreamTemplate_StatusARM struct {
	Auth            *UpstreamAuthSettings_StatusARM `json:"auth,omitempty"`
	CategoryPattern *string                         `json:"categoryPattern,omitempty"`
	EventPattern    *string                         `json:"eventPattern,omitempty"`
	HubPattern      *string                         `json:"hubPattern,omitempty"`
	UrlTemplate     *string                         `json:"urlTemplate,omitempty"`
}

//Deprecated version of UpstreamAuthSettings_Status. Use v1beta20211001.UpstreamAuthSettings_Status instead
type UpstreamAuthSettings_StatusARM struct {
	ManagedIdentity *ManagedIdentitySettings_StatusARM `json:"managedIdentity,omitempty"`
	Type            *UpstreamAuthType_Status           `json:"type,omitempty"`
}

//Deprecated version of ManagedIdentitySettings_Status. Use v1beta20211001.ManagedIdentitySettings_Status instead
type ManagedIdentitySettings_StatusARM struct {
	Resource *string `json:"resource,omitempty"`
}

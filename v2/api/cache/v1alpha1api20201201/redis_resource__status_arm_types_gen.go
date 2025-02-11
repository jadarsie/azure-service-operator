// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

//Deprecated version of RedisResource_Status. Use v1beta20201201.RedisResource_Status instead
type RedisResource_StatusARM struct {
	Id         *string                    `json:"id,omitempty"`
	Location   *string                    `json:"location,omitempty"`
	Name       *string                    `json:"name,omitempty"`
	Properties *RedisProperties_StatusARM `json:"properties,omitempty"`
	Tags       map[string]string          `json:"tags,omitempty"`
	Type       *string                    `json:"type,omitempty"`
	Zones      []string                   `json:"zones,omitempty"`
}

//Deprecated version of RedisProperties_Status. Use v1beta20201201.RedisProperties_Status instead
type RedisProperties_StatusARM struct {
	EnableNonSslPort           *bool                                                     `json:"enableNonSslPort,omitempty"`
	HostName                   *string                                                   `json:"hostName,omitempty"`
	Instances                  []RedisInstanceDetails_StatusARM                          `json:"instances,omitempty"`
	LinkedServers              []RedisLinkedServer_StatusARM                             `json:"linkedServers,omitempty"`
	MinimumTlsVersion          *RedisPropertiesStatusMinimumTlsVersion                   `json:"minimumTlsVersion,omitempty"`
	Port                       *int                                                      `json:"port,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_Status_SubResourceEmbeddedARM `json:"privateEndpointConnections,omitempty"`
	ProvisioningState          *RedisPropertiesStatusProvisioningState                   `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *RedisPropertiesStatusPublicNetworkAccess                 `json:"publicNetworkAccess,omitempty"`
	RedisConfiguration         map[string]string                                         `json:"redisConfiguration,omitempty"`
	RedisVersion               *string                                                   `json:"redisVersion,omitempty"`
	ReplicasPerMaster          *int                                                      `json:"replicasPerMaster,omitempty"`
	ReplicasPerPrimary         *int                                                      `json:"replicasPerPrimary,omitempty"`
	ShardCount                 *int                                                      `json:"shardCount,omitempty"`
	Sku                        *Sku_StatusARM                                            `json:"sku,omitempty"`
	SslPort                    *int                                                      `json:"sslPort,omitempty"`
	StaticIP                   *string                                                   `json:"staticIP,omitempty"`
	SubnetId                   *string                                                   `json:"subnetId,omitempty"`
	TenantSettings             map[string]string                                         `json:"tenantSettings,omitempty"`
}

//Deprecated version of PrivateEndpointConnection_Status_SubResourceEmbedded. Use v1beta20201201.PrivateEndpointConnection_Status_SubResourceEmbedded instead
type PrivateEndpointConnection_Status_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`
}

//Deprecated version of RedisInstanceDetails_Status. Use v1beta20201201.RedisInstanceDetails_Status instead
type RedisInstanceDetails_StatusARM struct {
	IsMaster   *bool   `json:"isMaster,omitempty"`
	IsPrimary  *bool   `json:"isPrimary,omitempty"`
	NonSslPort *int    `json:"nonSslPort,omitempty"`
	ShardId    *int    `json:"shardId,omitempty"`
	SslPort    *int    `json:"sslPort,omitempty"`
	Zone       *string `json:"zone,omitempty"`
}

//Deprecated version of RedisLinkedServer_Status. Use v1beta20201201.RedisLinkedServer_Status instead
type RedisLinkedServer_StatusARM struct {
	Id *string `json:"id,omitempty"`
}

//Deprecated version of Sku_Status. Use v1beta20201201.Sku_Status instead
type Sku_StatusARM struct {
	Capacity *int             `json:"capacity,omitempty"`
	Family   *SkuStatusFamily `json:"family,omitempty"`
	Name     *SkuStatusName   `json:"name,omitempty"`
}

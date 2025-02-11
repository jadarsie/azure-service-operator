// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package controllers

import (
	authorizationcustomizations "github.com/Azure/azure-service-operator/v2/api/authorization/customizations"
	authorizationv1alpha1api20200801preview "github.com/Azure/azure-service-operator/v2/api/authorization/v1alpha1api20200801preview"
	authorizationv1alpha1api20200801previewstorage "github.com/Azure/azure-service-operator/v2/api/authorization/v1alpha1api20200801previewstorage"
	authorizationv1beta20200801preview "github.com/Azure/azure-service-operator/v2/api/authorization/v1beta20200801preview"
	authorizationv1beta20200801previewstorage "github.com/Azure/azure-service-operator/v2/api/authorization/v1beta20200801previewstorage"
	batchcustomizations "github.com/Azure/azure-service-operator/v2/api/batch/customizations"
	batchv1alpha1api20210101 "github.com/Azure/azure-service-operator/v2/api/batch/v1alpha1api20210101"
	batchv1alpha1api20210101storage "github.com/Azure/azure-service-operator/v2/api/batch/v1alpha1api20210101storage"
	batchv1beta20210101 "github.com/Azure/azure-service-operator/v2/api/batch/v1beta20210101"
	batchv1beta20210101storage "github.com/Azure/azure-service-operator/v2/api/batch/v1beta20210101storage"
	cachecustomizations "github.com/Azure/azure-service-operator/v2/api/cache/customizations"
	cachev1alpha1api20201201 "github.com/Azure/azure-service-operator/v2/api/cache/v1alpha1api20201201"
	cachev1alpha1api20201201storage "github.com/Azure/azure-service-operator/v2/api/cache/v1alpha1api20201201storage"
	cachev1alpha1api20210301 "github.com/Azure/azure-service-operator/v2/api/cache/v1alpha1api20210301"
	cachev1alpha1api20210301storage "github.com/Azure/azure-service-operator/v2/api/cache/v1alpha1api20210301storage"
	cachev1beta20201201 "github.com/Azure/azure-service-operator/v2/api/cache/v1beta20201201"
	cachev1beta20201201storage "github.com/Azure/azure-service-operator/v2/api/cache/v1beta20201201storage"
	cachev1beta20210301 "github.com/Azure/azure-service-operator/v2/api/cache/v1beta20210301"
	cachev1beta20210301storage "github.com/Azure/azure-service-operator/v2/api/cache/v1beta20210301storage"
	computecustomizations "github.com/Azure/azure-service-operator/v2/api/compute/customizations"
	computev1alpha1api20200930 "github.com/Azure/azure-service-operator/v2/api/compute/v1alpha1api20200930"
	computev1alpha1api20200930storage "github.com/Azure/azure-service-operator/v2/api/compute/v1alpha1api20200930storage"
	computev1alpha1api20201201 "github.com/Azure/azure-service-operator/v2/api/compute/v1alpha1api20201201"
	computev1alpha1api20201201storage "github.com/Azure/azure-service-operator/v2/api/compute/v1alpha1api20201201storage"
	computev1alpha1api20210701 "github.com/Azure/azure-service-operator/v2/api/compute/v1alpha1api20210701"
	computev1alpha1api20210701storage "github.com/Azure/azure-service-operator/v2/api/compute/v1alpha1api20210701storage"
	computev1beta20200930 "github.com/Azure/azure-service-operator/v2/api/compute/v1beta20200930"
	computev1beta20200930storage "github.com/Azure/azure-service-operator/v2/api/compute/v1beta20200930storage"
	computev1beta20201201 "github.com/Azure/azure-service-operator/v2/api/compute/v1beta20201201"
	computev1beta20201201storage "github.com/Azure/azure-service-operator/v2/api/compute/v1beta20201201storage"
	computev1beta20210701 "github.com/Azure/azure-service-operator/v2/api/compute/v1beta20210701"
	computev1beta20210701storage "github.com/Azure/azure-service-operator/v2/api/compute/v1beta20210701storage"
	containerregistrycustomizations "github.com/Azure/azure-service-operator/v2/api/containerregistry/customizations"
	containerregistryv1alpha1api20210901 "github.com/Azure/azure-service-operator/v2/api/containerregistry/v1alpha1api20210901"
	containerregistryv1alpha1api20210901storage "github.com/Azure/azure-service-operator/v2/api/containerregistry/v1alpha1api20210901storage"
	containerregistryv1beta20210901 "github.com/Azure/azure-service-operator/v2/api/containerregistry/v1beta20210901"
	containerregistryv1beta20210901storage "github.com/Azure/azure-service-operator/v2/api/containerregistry/v1beta20210901storage"
	containerservicecustomizations "github.com/Azure/azure-service-operator/v2/api/containerservice/customizations"
	containerservicev1alpha1api20210501 "github.com/Azure/azure-service-operator/v2/api/containerservice/v1alpha1api20210501"
	containerservicev1alpha1api20210501storage "github.com/Azure/azure-service-operator/v2/api/containerservice/v1alpha1api20210501storage"
	containerservicev1beta20210501 "github.com/Azure/azure-service-operator/v2/api/containerservice/v1beta20210501"
	containerservicev1beta20210501storage "github.com/Azure/azure-service-operator/v2/api/containerservice/v1beta20210501storage"
	eventgridcustomizations "github.com/Azure/azure-service-operator/v2/api/eventgrid/customizations"
	eventgridv1alpha1api20200601 "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1alpha1api20200601"
	eventgridv1alpha1api20200601storage "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1alpha1api20200601storage"
	eventgridv1beta20200601 "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1beta20200601"
	eventgridv1beta20200601storage "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1beta20200601storage"
	insightscustomizations "github.com/Azure/azure-service-operator/v2/api/insights/customizations"
	insightsv1alpha1api20180501preview "github.com/Azure/azure-service-operator/v2/api/insights/v1alpha1api20180501preview"
	insightsv1alpha1api20180501previewstorage "github.com/Azure/azure-service-operator/v2/api/insights/v1alpha1api20180501previewstorage"
	insightsv1alpha1api20200202 "github.com/Azure/azure-service-operator/v2/api/insights/v1alpha1api20200202"
	insightsv1alpha1api20200202storage "github.com/Azure/azure-service-operator/v2/api/insights/v1alpha1api20200202storage"
	insightsv1beta20180501preview "github.com/Azure/azure-service-operator/v2/api/insights/v1beta20180501preview"
	insightsv1beta20180501previewstorage "github.com/Azure/azure-service-operator/v2/api/insights/v1beta20180501previewstorage"
	insightsv1beta20200202 "github.com/Azure/azure-service-operator/v2/api/insights/v1beta20200202"
	insightsv1beta20200202storage "github.com/Azure/azure-service-operator/v2/api/insights/v1beta20200202storage"
	managedidentitycustomizations "github.com/Azure/azure-service-operator/v2/api/managedidentity/customizations"
	managedidentityv1alpha1api20181130 "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1alpha1api20181130"
	managedidentityv1alpha1api20181130storage "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1alpha1api20181130storage"
	managedidentityv1beta20181130 "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1beta20181130"
	managedidentityv1beta20181130storage "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1beta20181130storage"
	networkcustomizations "github.com/Azure/azure-service-operator/v2/api/network/customizations"
	networkv1alpha1api20201101 "github.com/Azure/azure-service-operator/v2/api/network/v1alpha1api20201101"
	networkv1alpha1api20201101storage "github.com/Azure/azure-service-operator/v2/api/network/v1alpha1api20201101storage"
	networkv1beta20201101 "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101"
	networkv1beta20201101storage "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101storage"
	operationalinsightscustomizations "github.com/Azure/azure-service-operator/v2/api/operationalinsights/customizations"
	operationalinsightsv1alpha1api20210601 "github.com/Azure/azure-service-operator/v2/api/operationalinsights/v1alpha1api20210601"
	operationalinsightsv1alpha1api20210601storage "github.com/Azure/azure-service-operator/v2/api/operationalinsights/v1alpha1api20210601storage"
	operationalinsightsv1beta20210601 "github.com/Azure/azure-service-operator/v2/api/operationalinsights/v1beta20210601"
	operationalinsightsv1beta20210601storage "github.com/Azure/azure-service-operator/v2/api/operationalinsights/v1beta20210601storage"
	signalrservicecustomizations "github.com/Azure/azure-service-operator/v2/api/signalrservice/customizations"
	signalrservicev1alpha1api20211001 "github.com/Azure/azure-service-operator/v2/api/signalrservice/v1alpha1api20211001"
	signalrservicev1alpha1api20211001storage "github.com/Azure/azure-service-operator/v2/api/signalrservice/v1alpha1api20211001storage"
	signalrservicev1beta20211001 "github.com/Azure/azure-service-operator/v2/api/signalrservice/v1beta20211001"
	signalrservicev1beta20211001storage "github.com/Azure/azure-service-operator/v2/api/signalrservice/v1beta20211001storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/registration"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

// getKnownStorageTypes returns the list of storage types which can be reconciled.
func getKnownStorageTypes() []*registration.StorageType {
	var result []*registration.StorageType
	result = append(result, &registration.StorageType{
		Obj:     new(authorizationv1beta20200801previewstorage.RoleAssignment),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(batchv1beta20210101storage.BatchAccount),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(cachev1beta20201201storage.Redis),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(cachev1beta20201201storage.RedisFirewallRule),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(cachev1beta20201201storage.RedisLinkedServer),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(cachev1beta20201201storage.RedisPatchSchedule),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(cachev1beta20210301storage.RedisEnterprise),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(cachev1beta20210301storage.RedisEnterpriseDatabase),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(computev1beta20200930storage.Disk),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(computev1beta20200930storage.Snapshot),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj: new(computev1beta20201201storage.VirtualMachine),
		Indexes: []registration.Index{
			{
				Key:  ".spec.osProfile.adminPassword",
				Func: indexComputeVirtualMachineAdminPassword,
			},
		},
		Watches: []registration.Watch{
			{
				Src:              &source.Kind{Type: &v1.Secret{}},
				MakeEventHandler: watchSecretsFactory([]string{".spec.osProfile.adminPassword"}, &computev1beta20201201storage.VirtualMachineList{}),
			},
		},
	})
	result = append(result, &registration.StorageType{
		Obj: new(computev1beta20201201storage.VirtualMachineScaleSet),
		Indexes: []registration.Index{
			{
				Key:  ".spec.virtualMachineProfile.osProfile.adminPassword",
				Func: indexComputeVirtualMachineScaleSetAdminPassword,
			},
		},
		Watches: []registration.Watch{
			{
				Src:              &source.Kind{Type: &v1.Secret{}},
				MakeEventHandler: watchSecretsFactory([]string{".spec.virtualMachineProfile.osProfile.adminPassword"}, &computev1beta20201201storage.VirtualMachineScaleSetList{}),
			},
		},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(computev1beta20210701storage.Image),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(containerregistryv1beta20210901storage.Registry),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(containerservicev1beta20210501storage.ManagedCluster),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(containerservicev1beta20210501storage.ManagedClustersAgentPool),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(eventgridv1beta20200601storage.Domain),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(eventgridv1beta20200601storage.DomainsTopic),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(eventgridv1beta20200601storage.EventSubscription),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(eventgridv1beta20200601storage.Topic),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(insightsv1beta20180501previewstorage.Webtest),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(insightsv1beta20200202storage.Component),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(managedidentityv1beta20181130storage.UserAssignedIdentity),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(networkv1beta20201101storage.LoadBalancer),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(networkv1beta20201101storage.NetworkInterface),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(networkv1beta20201101storage.NetworkSecurityGroup),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(networkv1beta20201101storage.NetworkSecurityGroupsSecurityRule),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(networkv1beta20201101storage.PublicIPAddress),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(networkv1beta20201101storage.VirtualNetwork),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(networkv1beta20201101storage.VirtualNetworkGateway),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(networkv1beta20201101storage.VirtualNetworksSubnet),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(networkv1beta20201101storage.VirtualNetworksVirtualNetworkPeering),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(operationalinsightsv1beta20210601storage.Workspace),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	result = append(result, &registration.StorageType{
		Obj:     new(signalrservicev1beta20211001storage.SignalR),
		Indexes: []registration.Index{},
		Watches: []registration.Watch{},
	})
	return result
}

// getKnownTypes returns the list of all types.
func getKnownTypes() []client.Object {
	var result []client.Object
	result = append(result, new(authorizationv1alpha1api20200801preview.RoleAssignment))
	result = append(result, new(authorizationv1alpha1api20200801previewstorage.RoleAssignment))
	result = append(result, new(authorizationv1beta20200801preview.RoleAssignment))
	result = append(result, new(authorizationv1beta20200801previewstorage.RoleAssignment))
	result = append(result, new(batchv1alpha1api20210101.BatchAccount))
	result = append(result, new(batchv1alpha1api20210101storage.BatchAccount))
	result = append(result, new(batchv1beta20210101.BatchAccount))
	result = append(result, new(batchv1beta20210101storage.BatchAccount))
	result = append(result, new(cachev1alpha1api20201201.Redis))
	result = append(result, new(cachev1alpha1api20201201.RedisFirewallRule))
	result = append(result, new(cachev1alpha1api20201201.RedisLinkedServer))
	result = append(result, new(cachev1alpha1api20201201.RedisPatchSchedule))
	result = append(result, new(cachev1alpha1api20201201storage.Redis))
	result = append(result, new(cachev1alpha1api20201201storage.RedisFirewallRule))
	result = append(result, new(cachev1alpha1api20201201storage.RedisLinkedServer))
	result = append(result, new(cachev1alpha1api20201201storage.RedisPatchSchedule))
	result = append(result, new(cachev1alpha1api20210301.RedisEnterprise))
	result = append(result, new(cachev1alpha1api20210301.RedisEnterpriseDatabase))
	result = append(result, new(cachev1alpha1api20210301storage.RedisEnterprise))
	result = append(result, new(cachev1alpha1api20210301storage.RedisEnterpriseDatabase))
	result = append(result, new(cachev1beta20201201.Redis))
	result = append(result, new(cachev1beta20201201.RedisFirewallRule))
	result = append(result, new(cachev1beta20201201.RedisLinkedServer))
	result = append(result, new(cachev1beta20201201.RedisPatchSchedule))
	result = append(result, new(cachev1beta20201201storage.Redis))
	result = append(result, new(cachev1beta20201201storage.RedisFirewallRule))
	result = append(result, new(cachev1beta20201201storage.RedisLinkedServer))
	result = append(result, new(cachev1beta20201201storage.RedisPatchSchedule))
	result = append(result, new(cachev1beta20210301.RedisEnterprise))
	result = append(result, new(cachev1beta20210301.RedisEnterpriseDatabase))
	result = append(result, new(cachev1beta20210301storage.RedisEnterprise))
	result = append(result, new(cachev1beta20210301storage.RedisEnterpriseDatabase))
	result = append(result, new(computev1alpha1api20200930.Disk))
	result = append(result, new(computev1alpha1api20200930.Snapshot))
	result = append(result, new(computev1alpha1api20200930storage.Disk))
	result = append(result, new(computev1alpha1api20200930storage.Snapshot))
	result = append(result, new(computev1alpha1api20201201.VirtualMachine))
	result = append(result, new(computev1alpha1api20201201.VirtualMachineScaleSet))
	result = append(result, new(computev1alpha1api20201201storage.VirtualMachine))
	result = append(result, new(computev1alpha1api20201201storage.VirtualMachineScaleSet))
	result = append(result, new(computev1alpha1api20210701.Image))
	result = append(result, new(computev1alpha1api20210701storage.Image))
	result = append(result, new(computev1beta20200930.Disk))
	result = append(result, new(computev1beta20200930.Snapshot))
	result = append(result, new(computev1beta20200930storage.Disk))
	result = append(result, new(computev1beta20200930storage.Snapshot))
	result = append(result, new(computev1beta20201201.VirtualMachine))
	result = append(result, new(computev1beta20201201.VirtualMachineScaleSet))
	result = append(result, new(computev1beta20201201storage.VirtualMachine))
	result = append(result, new(computev1beta20201201storage.VirtualMachineScaleSet))
	result = append(result, new(computev1beta20210701.Image))
	result = append(result, new(computev1beta20210701storage.Image))
	result = append(result, new(containerregistryv1alpha1api20210901.Registry))
	result = append(result, new(containerregistryv1alpha1api20210901storage.Registry))
	result = append(result, new(containerregistryv1beta20210901.Registry))
	result = append(result, new(containerregistryv1beta20210901storage.Registry))
	result = append(result, new(containerservicev1alpha1api20210501.ManagedCluster))
	result = append(result, new(containerservicev1alpha1api20210501.ManagedClustersAgentPool))
	result = append(result, new(containerservicev1alpha1api20210501storage.ManagedCluster))
	result = append(result, new(containerservicev1alpha1api20210501storage.ManagedClustersAgentPool))
	result = append(result, new(containerservicev1beta20210501.ManagedCluster))
	result = append(result, new(containerservicev1beta20210501.ManagedClustersAgentPool))
	result = append(result, new(containerservicev1beta20210501storage.ManagedCluster))
	result = append(result, new(containerservicev1beta20210501storage.ManagedClustersAgentPool))
	result = append(result, new(eventgridv1alpha1api20200601.Domain))
	result = append(result, new(eventgridv1alpha1api20200601.DomainsTopic))
	result = append(result, new(eventgridv1alpha1api20200601.EventSubscription))
	result = append(result, new(eventgridv1alpha1api20200601.Topic))
	result = append(result, new(eventgridv1alpha1api20200601storage.Domain))
	result = append(result, new(eventgridv1alpha1api20200601storage.DomainsTopic))
	result = append(result, new(eventgridv1alpha1api20200601storage.EventSubscription))
	result = append(result, new(eventgridv1alpha1api20200601storage.Topic))
	result = append(result, new(eventgridv1beta20200601.Domain))
	result = append(result, new(eventgridv1beta20200601.DomainsTopic))
	result = append(result, new(eventgridv1beta20200601.EventSubscription))
	result = append(result, new(eventgridv1beta20200601.Topic))
	result = append(result, new(eventgridv1beta20200601storage.Domain))
	result = append(result, new(eventgridv1beta20200601storage.DomainsTopic))
	result = append(result, new(eventgridv1beta20200601storage.EventSubscription))
	result = append(result, new(eventgridv1beta20200601storage.Topic))
	result = append(result, new(insightsv1alpha1api20180501preview.Webtest))
	result = append(result, new(insightsv1alpha1api20180501previewstorage.Webtest))
	result = append(result, new(insightsv1alpha1api20200202.Component))
	result = append(result, new(insightsv1alpha1api20200202storage.Component))
	result = append(result, new(insightsv1beta20180501preview.Webtest))
	result = append(result, new(insightsv1beta20180501previewstorage.Webtest))
	result = append(result, new(insightsv1beta20200202.Component))
	result = append(result, new(insightsv1beta20200202storage.Component))
	result = append(result, new(managedidentityv1alpha1api20181130.UserAssignedIdentity))
	result = append(result, new(managedidentityv1alpha1api20181130storage.UserAssignedIdentity))
	result = append(result, new(managedidentityv1beta20181130.UserAssignedIdentity))
	result = append(result, new(managedidentityv1beta20181130storage.UserAssignedIdentity))
	result = append(result, new(networkv1alpha1api20201101.LoadBalancer))
	result = append(result, new(networkv1alpha1api20201101.NetworkInterface))
	result = append(result, new(networkv1alpha1api20201101.NetworkSecurityGroup))
	result = append(result, new(networkv1alpha1api20201101.NetworkSecurityGroupsSecurityRule))
	result = append(result, new(networkv1alpha1api20201101.PublicIPAddress))
	result = append(result, new(networkv1alpha1api20201101.VirtualNetwork))
	result = append(result, new(networkv1alpha1api20201101.VirtualNetworkGateway))
	result = append(result, new(networkv1alpha1api20201101.VirtualNetworksSubnet))
	result = append(result, new(networkv1alpha1api20201101.VirtualNetworksVirtualNetworkPeering))
	result = append(result, new(networkv1alpha1api20201101storage.LoadBalancer))
	result = append(result, new(networkv1alpha1api20201101storage.NetworkInterface))
	result = append(result, new(networkv1alpha1api20201101storage.NetworkSecurityGroup))
	result = append(result, new(networkv1alpha1api20201101storage.NetworkSecurityGroupsSecurityRule))
	result = append(result, new(networkv1alpha1api20201101storage.PublicIPAddress))
	result = append(result, new(networkv1alpha1api20201101storage.VirtualNetwork))
	result = append(result, new(networkv1alpha1api20201101storage.VirtualNetworkGateway))
	result = append(result, new(networkv1alpha1api20201101storage.VirtualNetworksSubnet))
	result = append(result, new(networkv1alpha1api20201101storage.VirtualNetworksVirtualNetworkPeering))
	result = append(result, new(networkv1beta20201101.LoadBalancer))
	result = append(result, new(networkv1beta20201101.NetworkInterface))
	result = append(result, new(networkv1beta20201101.NetworkSecurityGroup))
	result = append(result, new(networkv1beta20201101.NetworkSecurityGroupsSecurityRule))
	result = append(result, new(networkv1beta20201101.PublicIPAddress))
	result = append(result, new(networkv1beta20201101.VirtualNetwork))
	result = append(result, new(networkv1beta20201101.VirtualNetworkGateway))
	result = append(result, new(networkv1beta20201101.VirtualNetworksSubnet))
	result = append(result, new(networkv1beta20201101.VirtualNetworksVirtualNetworkPeering))
	result = append(result, new(networkv1beta20201101storage.LoadBalancer))
	result = append(result, new(networkv1beta20201101storage.NetworkInterface))
	result = append(result, new(networkv1beta20201101storage.NetworkSecurityGroup))
	result = append(result, new(networkv1beta20201101storage.NetworkSecurityGroupsSecurityRule))
	result = append(result, new(networkv1beta20201101storage.PublicIPAddress))
	result = append(result, new(networkv1beta20201101storage.VirtualNetwork))
	result = append(result, new(networkv1beta20201101storage.VirtualNetworkGateway))
	result = append(result, new(networkv1beta20201101storage.VirtualNetworksSubnet))
	result = append(result, new(networkv1beta20201101storage.VirtualNetworksVirtualNetworkPeering))
	result = append(result, new(operationalinsightsv1alpha1api20210601.Workspace))
	result = append(result, new(operationalinsightsv1alpha1api20210601storage.Workspace))
	result = append(result, new(operationalinsightsv1beta20210601.Workspace))
	result = append(result, new(operationalinsightsv1beta20210601storage.Workspace))
	result = append(result, new(signalrservicev1alpha1api20211001.SignalR))
	result = append(result, new(signalrservicev1alpha1api20211001storage.SignalR))
	result = append(result, new(signalrservicev1beta20211001.SignalR))
	result = append(result, new(signalrservicev1beta20211001storage.SignalR))
	return result
}

// createScheme creates a Scheme containing the clientgo types and all of the custom types returned by getKnownTypes
func createScheme() *runtime.Scheme {
	scheme := runtime.NewScheme()
	_ = clientgoscheme.AddToScheme(scheme)
	_ = authorizationv1alpha1api20200801preview.AddToScheme(scheme)
	_ = authorizationv1alpha1api20200801previewstorage.AddToScheme(scheme)
	_ = authorizationv1beta20200801preview.AddToScheme(scheme)
	_ = authorizationv1beta20200801previewstorage.AddToScheme(scheme)
	_ = batchv1alpha1api20210101.AddToScheme(scheme)
	_ = batchv1alpha1api20210101storage.AddToScheme(scheme)
	_ = batchv1beta20210101.AddToScheme(scheme)
	_ = batchv1beta20210101storage.AddToScheme(scheme)
	_ = cachev1alpha1api20201201.AddToScheme(scheme)
	_ = cachev1alpha1api20201201storage.AddToScheme(scheme)
	_ = cachev1alpha1api20210301.AddToScheme(scheme)
	_ = cachev1alpha1api20210301storage.AddToScheme(scheme)
	_ = cachev1beta20201201.AddToScheme(scheme)
	_ = cachev1beta20201201storage.AddToScheme(scheme)
	_ = cachev1beta20210301.AddToScheme(scheme)
	_ = cachev1beta20210301storage.AddToScheme(scheme)
	_ = computev1alpha1api20200930.AddToScheme(scheme)
	_ = computev1alpha1api20200930storage.AddToScheme(scheme)
	_ = computev1alpha1api20201201.AddToScheme(scheme)
	_ = computev1alpha1api20201201storage.AddToScheme(scheme)
	_ = computev1alpha1api20210701.AddToScheme(scheme)
	_ = computev1alpha1api20210701storage.AddToScheme(scheme)
	_ = computev1beta20200930.AddToScheme(scheme)
	_ = computev1beta20200930storage.AddToScheme(scheme)
	_ = computev1beta20201201.AddToScheme(scheme)
	_ = computev1beta20201201storage.AddToScheme(scheme)
	_ = computev1beta20210701.AddToScheme(scheme)
	_ = computev1beta20210701storage.AddToScheme(scheme)
	_ = containerregistryv1alpha1api20210901.AddToScheme(scheme)
	_ = containerregistryv1alpha1api20210901storage.AddToScheme(scheme)
	_ = containerregistryv1beta20210901.AddToScheme(scheme)
	_ = containerregistryv1beta20210901storage.AddToScheme(scheme)
	_ = containerservicev1alpha1api20210501.AddToScheme(scheme)
	_ = containerservicev1alpha1api20210501storage.AddToScheme(scheme)
	_ = containerservicev1beta20210501.AddToScheme(scheme)
	_ = containerservicev1beta20210501storage.AddToScheme(scheme)
	_ = eventgridv1alpha1api20200601.AddToScheme(scheme)
	_ = eventgridv1alpha1api20200601storage.AddToScheme(scheme)
	_ = eventgridv1beta20200601.AddToScheme(scheme)
	_ = eventgridv1beta20200601storage.AddToScheme(scheme)
	_ = insightsv1alpha1api20180501preview.AddToScheme(scheme)
	_ = insightsv1alpha1api20180501previewstorage.AddToScheme(scheme)
	_ = insightsv1alpha1api20200202.AddToScheme(scheme)
	_ = insightsv1alpha1api20200202storage.AddToScheme(scheme)
	_ = insightsv1beta20180501preview.AddToScheme(scheme)
	_ = insightsv1beta20180501previewstorage.AddToScheme(scheme)
	_ = insightsv1beta20200202.AddToScheme(scheme)
	_ = insightsv1beta20200202storage.AddToScheme(scheme)
	_ = managedidentityv1alpha1api20181130.AddToScheme(scheme)
	_ = managedidentityv1alpha1api20181130storage.AddToScheme(scheme)
	_ = managedidentityv1beta20181130.AddToScheme(scheme)
	_ = managedidentityv1beta20181130storage.AddToScheme(scheme)
	_ = networkv1alpha1api20201101.AddToScheme(scheme)
	_ = networkv1alpha1api20201101storage.AddToScheme(scheme)
	_ = networkv1beta20201101.AddToScheme(scheme)
	_ = networkv1beta20201101storage.AddToScheme(scheme)
	_ = operationalinsightsv1alpha1api20210601.AddToScheme(scheme)
	_ = operationalinsightsv1alpha1api20210601storage.AddToScheme(scheme)
	_ = operationalinsightsv1beta20210601.AddToScheme(scheme)
	_ = operationalinsightsv1beta20210601storage.AddToScheme(scheme)
	_ = signalrservicev1alpha1api20211001.AddToScheme(scheme)
	_ = signalrservicev1alpha1api20211001storage.AddToScheme(scheme)
	_ = signalrservicev1beta20211001.AddToScheme(scheme)
	_ = signalrservicev1beta20211001storage.AddToScheme(scheme)
	return scheme
}

// getResourceExtensions returns a list of resource extensions
func getResourceExtensions() []genruntime.ResourceExtension {
	var result []genruntime.ResourceExtension
	result = append(result, &authorizationcustomizations.RoleAssignmentExtension{})
	result = append(result, &batchcustomizations.BatchAccountExtension{})
	result = append(result, &cachecustomizations.RedisEnterpriseDatabaseExtension{})
	result = append(result, &cachecustomizations.RedisEnterpriseExtension{})
	result = append(result, &cachecustomizations.RedisExtension{})
	result = append(result, &cachecustomizations.RedisFirewallRuleExtension{})
	result = append(result, &cachecustomizations.RedisLinkedServerExtension{})
	result = append(result, &cachecustomizations.RedisPatchScheduleExtension{})
	result = append(result, &computecustomizations.DiskExtension{})
	result = append(result, &computecustomizations.ImageExtension{})
	result = append(result, &computecustomizations.SnapshotExtension{})
	result = append(result, &computecustomizations.VirtualMachineExtension{})
	result = append(result, &computecustomizations.VirtualMachineScaleSetExtension{})
	result = append(result, &containerregistrycustomizations.RegistryExtension{})
	result = append(result, &containerservicecustomizations.ManagedClusterExtension{})
	result = append(result, &containerservicecustomizations.ManagedClustersAgentPoolExtension{})
	result = append(result, &eventgridcustomizations.DomainExtension{})
	result = append(result, &eventgridcustomizations.DomainsTopicExtension{})
	result = append(result, &eventgridcustomizations.EventSubscriptionExtension{})
	result = append(result, &eventgridcustomizations.TopicExtension{})
	result = append(result, &insightscustomizations.ComponentExtension{})
	result = append(result, &insightscustomizations.WebtestExtension{})
	result = append(result, &managedidentitycustomizations.UserAssignedIdentityExtension{})
	result = append(result, &networkcustomizations.LoadBalancerExtension{})
	result = append(result, &networkcustomizations.NetworkInterfaceExtension{})
	result = append(result, &networkcustomizations.NetworkSecurityGroupExtension{})
	result = append(result, &networkcustomizations.NetworkSecurityGroupsSecurityRuleExtension{})
	result = append(result, &networkcustomizations.PublicIPAddressExtension{})
	result = append(result, &networkcustomizations.VirtualNetworkExtension{})
	result = append(result, &networkcustomizations.VirtualNetworkGatewayExtension{})
	result = append(result, &networkcustomizations.VirtualNetworksSubnetExtension{})
	result = append(result, &networkcustomizations.VirtualNetworksVirtualNetworkPeeringExtension{})
	result = append(result, &operationalinsightscustomizations.WorkspaceExtension{})
	result = append(result, &signalrservicecustomizations.SignalRExtension{})
	return result
}

// indexComputeVirtualMachineAdminPassword an index function for computev1beta20201201storage.VirtualMachine .spec.osProfile.adminPassword
func indexComputeVirtualMachineAdminPassword(rawObj client.Object) []string {
	obj, ok := rawObj.(*computev1beta20201201storage.VirtualMachine)
	if !ok {
		return nil
	}
	if obj.Spec.OsProfile == nil {
		return nil
	}
	if obj.Spec.OsProfile.AdminPassword == nil {
		return nil
	}
	return []string{obj.Spec.OsProfile.AdminPassword.Name}
}

// indexComputeVirtualMachineScaleSetAdminPassword an index function for computev1beta20201201storage.VirtualMachineScaleSet .spec.virtualMachineProfile.osProfile.adminPassword
func indexComputeVirtualMachineScaleSetAdminPassword(rawObj client.Object) []string {
	obj, ok := rawObj.(*computev1beta20201201storage.VirtualMachineScaleSet)
	if !ok {
		return nil
	}
	if obj.Spec.VirtualMachineProfile == nil {
		return nil
	}
	if obj.Spec.VirtualMachineProfile.OsProfile == nil {
		return nil
	}
	if obj.Spec.VirtualMachineProfile.OsProfile.AdminPassword == nil {
		return nil
	}
	return []string{obj.Spec.VirtualMachineProfile.OsProfile.AdminPassword.Name}
}

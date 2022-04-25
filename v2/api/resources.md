# Supported Resources

These are the resources with Azure Service Operator support committed to our **main** branch, grouped by the originating ARM service. (Newly supported resources will appear in this list prior to inclusion in any ASO release.)

## Authorization

| Resource       | ARM Version        | CRD Version                | Sample                                                                                                                                           |
|----------------|--------------------|----------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------|
| RoleAssignment | 2020-08-01-preview | v1beta20200801preview      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/authorization/v1beta20200801preview_roleassignment.yaml)      |
| RoleAssignment | 2020-08-01-preview | v1alpha1api20200801preview | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/authorization/v1alpha1api20200801preview_roleassignment.yaml) |

## Batch

| Resource     | ARM Version | CRD Version         | Sample                                                                                                                          |
|--------------|-------------|---------------------|---------------------------------------------------------------------------------------------------------------------------------|
| BatchAccount | 2021-01-01  | v1beta20210101      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/batch/v1beta20210101_batchaccount.yaml)      |
| BatchAccount | 2021-01-01  | v1alpha1api20210101 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/batch/v1alpha1api20210101_batchaccount.yaml) |

## Cache

| Resource                | ARM Version | CRD Version         | Sample                                                                                                                                     |
|-------------------------|-------------|---------------------|--------------------------------------------------------------------------------------------------------------------------------------------|
| Redis                   | 2020-12-01  | v1beta20201201      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/cache/v1beta20201201_redis.yaml)                        |
| Redis                   | 2020-12-01  | v1alpha1api20201201 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/cache/v1alpha1api20201201_redis.yaml)                   |
| RedisEnterprise         | 2021-03-01  | v1beta20210301      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/cache/v1beta20210301_redisenterprise.yaml)              |
| RedisEnterprise         | 2021-03-01  | v1alpha1api20210301 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/cache/v1alpha1api20210301_redisenterprise.yaml)         |
| RedisEnterpriseDatabase | 2021-03-01  | v1beta20210301      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/cache/v1beta20210301_redisenterprisedatabase.yaml)      |
| RedisEnterpriseDatabase | 2021-03-01  | v1alpha1api20210301 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/cache/v1alpha1api20210301_redisenterprisedatabase.yaml) |
| RedisFirewallRule       | 2020-12-01  | v1beta20201201      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/cache/v1beta20201201_redisfirewallrule.yaml)            |
| RedisFirewallRule       | 2020-12-01  | v1alpha1api20201201 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/cache/v1alpha1api20201201_redisfirewallrule.yaml)       |
| RedisLinkedServer       | 2020-12-01  | v1beta20201201      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/cache/v1beta20201201_redislinkedserver.yaml)            |
| RedisLinkedServer       | 2020-12-01  | v1alpha1api20201201 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/cache/v1alpha1api20201201_redislinkedserver.yaml)       |
| RedisPatchSchedule      | 2020-12-01  | v1beta20201201      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/cache/v1beta20201201_redispatchschedule.yaml)           |
| RedisPatchSchedule      | 2020-12-01  | v1alpha1api20201201 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/cache/v1alpha1api20201201_redispatchschedule.yaml)      |

## Compute

| Resource               | ARM Version | CRD Version         | Sample                                                                                                                                      |
|------------------------|-------------|---------------------|---------------------------------------------------------------------------------------------------------------------------------------------|
| Disk                   | 2020-09-30  | v1beta20200930      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/compute/v1beta20200930_disk.yaml)                        |
| Disk                   | 2020-09-30  | v1alpha1api20200930 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/compute/v1alpha1api20200930_disk.yaml)                   |
| Image                  | 2021-07-01  | v1beta20210701      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/compute/v1beta20210701_image.yaml)                       |
| Image                  | 2021-07-01  | v1alpha1api20210701 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/compute/v1alpha1api20210701_image.yaml)                  |
| Snapshot               | 2020-09-30  | v1beta20200930      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/compute/v1beta20200930_snapshot.yaml)                    |
| Snapshot               | 2020-09-30  | v1alpha1api20200930 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/compute/v1alpha1api20200930_snapshot.yaml)               |
| VirtualMachine         | 2020-12-01  | v1beta20201201      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/compute/v1beta20201201_virtualmachine.yaml)              |
| VirtualMachine         | 2020-12-01  | v1alpha1api20201201 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/compute/v1alpha1api20201201_virtualmachine.yaml)         |
| VirtualMachineScaleSet | 2020-12-01  | v1beta20201201      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/compute/v1beta20201201_virtualmachinescaleset.yaml)      |
| VirtualMachineScaleSet | 2020-12-01  | v1alpha1api20201201 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/compute/v1alpha1api20201201_virtualmachinescaleset.yaml) |

## Containerregistry

| Resource | ARM Version | CRD Version         | Sample                                                                                                                                  |
|----------|-------------|---------------------|-----------------------------------------------------------------------------------------------------------------------------------------|
| Registry | 2021-09-01  | v1beta20210901      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/containerregistry/v1beta20210901_registry.yaml)      |
| Registry | 2021-09-01  | v1alpha1api20210901 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/containerregistry/v1alpha1api20210901_registry.yaml) |

## Containerservice

| Resource                 | ARM Version | CRD Version         | Sample                                                                                                                                                 |
|--------------------------|-------------|---------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------|
| ManagedCluster           | 2021-05-01  | v1beta20210501      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/containerservice/v1beta20210501_managedcluster.yaml)                |
| ManagedCluster           | 2021-05-01  | v1alpha1api20210501 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/containerservice/v1alpha1api20210501_managedcluster.yaml)           |
| ManagedClustersAgentPool | 2021-05-01  | v1beta20210501      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/containerservice/v1beta20210501_managedclustersagentpool.yaml)      |
| ManagedClustersAgentPool | 2021-05-01  | v1alpha1api20210501 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/containerservice/v1alpha1api20210501_managedclustersagentpool.yaml) |

## Eventgrid

| Resource          | ARM Version | CRD Version         | Sample                                                                                                                                   |
|-------------------|-------------|---------------------|------------------------------------------------------------------------------------------------------------------------------------------|
| Domain            | 2020-06-01  | v1beta20200601      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/eventgrid/v1beta20200601_domain.yaml)                 |
| Domain            | 2020-06-01  | v1alpha1api20200601 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/eventgrid/v1alpha1api20200601_domain.yaml)            |
| DomainsTopic      | 2020-06-01  | v1beta20200601      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/eventgrid/v1beta20200601_domainstopic.yaml)           |
| DomainsTopic      | 2020-06-01  | v1alpha1api20200601 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/eventgrid/v1alpha1api20200601_domainstopic.yaml)      |
| EventSubscription | 2020-06-01  | v1beta20200601      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/eventgrid/v1beta20200601_eventsubscription.yaml)      |
| EventSubscription | 2020-06-01  | v1alpha1api20200601 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/eventgrid/v1alpha1api20200601_eventsubscription.yaml) |
| Topic             | 2020-06-01  | v1beta20200601      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/eventgrid/v1beta20200601_topic.yaml)                  |
| Topic             | 2020-06-01  | v1alpha1api20200601 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/eventgrid/v1alpha1api20200601_topic.yaml)             |

## Insights

| Resource  | ARM Version        | CRD Version                | Sample                                                                                                                               |
|-----------|--------------------|----------------------------|--------------------------------------------------------------------------------------------------------------------------------------|
| Component | 2020-02-02         | v1beta20200202             | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/insights/v1beta20200202_component.yaml)           |
| Component | 2020-02-02         | v1alpha1api20200202        | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/insights/v1alpha1api20200202_component.yaml)      |
| Webtest   | 2018-05-01-preview | v1beta20180501preview      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/insights/v1beta20180501preview_webtest.yaml)      |
| Webtest   | 2018-05-01-preview | v1alpha1api20180501preview | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/insights/v1alpha1api20180501preview_webtest.yaml) |

## Managedidentity

| Resource             | ARM Version | CRD Version         | Sample                                                                                                                                            |
|----------------------|-------------|---------------------|---------------------------------------------------------------------------------------------------------------------------------------------------|
| UserAssignedIdentity | 2018-11-30  | v1beta20181130      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/managedidentity/v1beta20181130_userassignedidentity.yaml)      |
| UserAssignedIdentity | 2018-11-30  | v1alpha1api20181130 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/managedidentity/v1alpha1api20181130_userassignedidentity.yaml) |

## Network

| Resource                             | ARM Version | CRD Version         | Sample                                                                                                                                                    |
|--------------------------------------|-------------|---------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------|
| LoadBalancer                         | 2020-11-01  | v1beta20201101      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1beta20201101_loadbalancer.yaml)                              |
| LoadBalancer                         | 2020-11-01  | v1alpha1api20201101 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1alpha1api20201101_loadbalancer.yaml)                         |
| NetworkInterface                     | 2020-11-01  | v1beta20201101      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1beta20201101_networkinterface.yaml)                          |
| NetworkInterface                     | 2020-11-01  | v1alpha1api20201101 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1alpha1api20201101_networkinterface.yaml)                     |
| NetworkSecurityGroup                 | 2020-11-01  | v1beta20201101      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1beta20201101_networksecuritygroup.yaml)                      |
| NetworkSecurityGroup                 | 2020-11-01  | v1alpha1api20201101 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1alpha1api20201101_networksecuritygroup.yaml)                 |
| NetworkSecurityGroupsSecurityRule    | 2020-11-01  | v1beta20201101      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1beta20201101_networksecuritygroupssecurityrule.yaml)         |
| NetworkSecurityGroupsSecurityRule    | 2020-11-01  | v1alpha1api20201101 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1alpha1api20201101_networksecuritygroupssecurityrule.yaml)    |
| PublicIPAddress                      | 2020-11-01  | v1beta20201101      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1beta20201101_publicipaddress.yaml)                           |
| PublicIPAddress                      | 2020-11-01  | v1alpha1api20201101 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1alpha1api20201101_publicipaddress.yaml)                      |
| VirtualNetwork                       | 2020-11-01  | v1beta20201101      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1beta20201101_virtualnetwork.yaml)                            |
| VirtualNetwork                       | 2020-11-01  | v1alpha1api20201101 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1alpha1api20201101_virtualnetwork.yaml)                       |
| VirtualNetworkGateway                | 2020-11-01  | v1beta20201101      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1beta20201101_virtualnetworkgateway.yaml)                     |
| VirtualNetworkGateway                | 2020-11-01  | v1alpha1api20201101 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1alpha1api20201101_virtualnetworkgateway.yaml)                |
| VirtualNetworksSubnet                | 2020-11-01  | v1beta20201101      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1beta20201101_virtualnetworkssubnet.yaml)                     |
| VirtualNetworksSubnet                | 2020-11-01  | v1alpha1api20201101 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1alpha1api20201101_virtualnetworkssubnet.yaml)                |
| VirtualNetworksVirtualNetworkPeering | 2020-11-01  | v1beta20201101      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1beta20201101_virtualnetworksvirtualnetworkpeering.yaml)      |
| VirtualNetworksVirtualNetworkPeering | 2020-11-01  | v1alpha1api20201101 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/network/v1alpha1api20201101_virtualnetworksvirtualnetworkpeering.yaml) |

## Operationalinsights

| Resource  | ARM Version | CRD Version         | Sample                                                                                                                                     |
|-----------|-------------|---------------------|--------------------------------------------------------------------------------------------------------------------------------------------|
| Workspace | 2021-06-01  | v1beta20210601      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/operationalinsights/v1beta20210601_workspace.yaml)      |
| Workspace | 2021-06-01  | v1alpha1api20210601 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/operationalinsights/v1alpha1api20210601_workspace.yaml) |

## Signalrservice

| Resource | ARM Version | CRD Version         | Sample                                                                                                                              |
|----------|-------------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------|
| SignalR  | 2021-10-01  | v1beta20211001      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/signalrservice/v1beta20211001_signalr.yaml)      |
| SignalR  | 2021-10-01  | v1alpha1api20211001 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/signalrservice/v1alpha1api20211001_signalr.yaml) |

## Storage

| Resource                             | ARM Version | CRD Version         | Sample                                                                                                                                                    |
|--------------------------------------|-------------|---------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------|
| StorageAccount                       | 2021-04-01  | v1beta20210401      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/storage/v1beta20210401_storageaccount.yaml)                            |
| StorageAccount                       | 2021-04-01  | v1alpha1api20210401 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/storage/v1alpha1api20210401_storageaccount.yaml)                       |
| StorageAccountsBlobService           | 2021-04-01  | v1beta20210401      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/storage/v1beta20210401_storageaccountsblobservice.yaml)                |
| StorageAccountsBlobService           | 2021-04-01  | v1alpha1api20210401 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/storage/v1alpha1api20210401_storageaccountsblobservice.yaml)           |
| StorageAccountsBlobServicesContainer | 2021-04-01  | v1beta20210401      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/storage/v1beta20210401_storageaccountsblobservicescontainer.yaml)      |
| StorageAccountsBlobServicesContainer | 2021-04-01  | v1alpha1api20210401 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/storage/v1alpha1api20210401_storageaccountsblobservicescontainer.yaml) |
| StorageAccountsManagementPolicy      | 2021-04-01  | v1beta20210401      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/storage/v1beta20210401_storageaccountsmanagementpolicy.yaml)           |
| StorageAccountsManagementPolicy      | 2021-04-01  | v1alpha1api20210401 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/storage/v1alpha1api20210401_storageaccountsmanagementpolicy.yaml)      |
| StorageAccountsQueueService          | 2021-04-01  | v1beta20210401      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/storage/v1beta20210401_storageaccountsqueueservice.yaml)               |
| StorageAccountsQueueService          | 2021-04-01  | v1alpha1api20210401 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/storage/v1alpha1api20210401_storageaccountsqueueservice.yaml)          |
| StorageAccountsQueueServicesQueue    | 2021-04-01  | v1beta20210401      | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/storage/v1beta20210401_storageaccountsqueueservicesqueue.yaml)         |
| StorageAccountsQueueServicesQueue    | 2021-04-01  | v1alpha1api20210401 | [View](https://github.com/Azure/azure-service-operator/tree/main/v2/config/samples/storage/v1alpha1api20210401_storageaccountsqueueservicesqueue.yaml)    |


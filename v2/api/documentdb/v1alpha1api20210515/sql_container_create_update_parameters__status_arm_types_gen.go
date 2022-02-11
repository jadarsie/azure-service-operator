// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

type SqlContainerCreateUpdateParameters_StatusARM struct {
	//Id: The unique resource identifier of the ARM resource.
	Id *string `json:"id,omitempty"`

	//Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	//Name: The name of the ARM resource.
	Name *string `json:"name,omitempty"`

	//Properties: Properties to create and update Azure Cosmos DB container.
	Properties *SqlContainerCreateUpdateProperties_StatusARM `json:"properties,omitempty"`
	Tags       map[string]string                             `json:"tags,omitempty"`

	//Type: The type of Azure resource.
	Type *string `json:"type,omitempty"`
}

type SqlContainerCreateUpdateProperties_StatusARM struct {
	//Options: A key-value pair of options to be applied for the request. This
	//corresponds to the headers sent with the request.
	Options *CreateUpdateOptions_StatusARM `json:"options,omitempty"`

	//Resource: The standard JSON format of a container
	Resource SqlContainerResource_StatusARM `json:"resource"`
}

type SqlContainerResource_StatusARM struct {
	//AnalyticalStorageTtl: Analytical TTL.
	AnalyticalStorageTtl *int `json:"analyticalStorageTtl,omitempty"`

	//ConflictResolutionPolicy: The conflict resolution policy for the container.
	ConflictResolutionPolicy *ConflictResolutionPolicy_StatusARM `json:"conflictResolutionPolicy,omitempty"`

	//DefaultTtl: Default time to live
	DefaultTtl *int `json:"defaultTtl,omitempty"`

	//Id: Name of the Cosmos DB SQL container
	Id string `json:"id"`

	//IndexingPolicy: The configuration of the indexing policy. By default, the
	//indexing is automatic for all document paths within the container
	IndexingPolicy *IndexingPolicy_StatusARM `json:"indexingPolicy,omitempty"`

	//PartitionKey: The configuration of the partition key to be used for partitioning
	//data into multiple partitions
	PartitionKey *ContainerPartitionKey_StatusARM `json:"partitionKey,omitempty"`

	//UniqueKeyPolicy: The unique key policy configuration for specifying uniqueness
	//constraints on documents in the collection in the Azure Cosmos DB service.
	UniqueKeyPolicy *UniqueKeyPolicy_StatusARM `json:"uniqueKeyPolicy,omitempty"`
}

type ConflictResolutionPolicy_StatusARM struct {
	//ConflictResolutionPath: The conflict resolution path in the case of
	//LastWriterWins mode.
	ConflictResolutionPath *string `json:"conflictResolutionPath,omitempty"`

	//ConflictResolutionProcedure: The procedure to resolve conflicts in the case of
	//custom mode.
	ConflictResolutionProcedure *string `json:"conflictResolutionProcedure,omitempty"`

	//Mode: Indicates the conflict resolution mode.
	Mode *ConflictResolutionPolicy_Mode_Status `json:"mode,omitempty"`
}

type ContainerPartitionKey_StatusARM struct {
	//Kind: Indicates the kind of algorithm used for partitioning. For MultiHash,
	//multiple partition keys (upto three maximum) are supported for container create
	Kind *ContainerPartitionKey_Kind_Status `json:"kind,omitempty"`

	//Paths: List of paths using which data within the container can be partitioned
	Paths []string `json:"paths,omitempty"`

	//SystemKey: Indicates if the container is using a system generated partition key
	SystemKey *bool `json:"systemKey,omitempty"`

	//Version: Indicates the version of the partition key definition
	Version *int `json:"version,omitempty"`
}

type IndexingPolicy_StatusARM struct {
	//Automatic: Indicates if the indexing policy is automatic
	Automatic *bool `json:"automatic,omitempty"`

	//CompositeIndexes: List of composite path list
	CompositeIndexes [][]CompositePath_StatusARM `json:"compositeIndexes,omitempty"`

	//ExcludedPaths: List of paths to exclude from indexing
	ExcludedPaths []ExcludedPath_StatusARM `json:"excludedPaths,omitempty"`

	//IncludedPaths: List of paths to include in the indexing
	IncludedPaths []IncludedPath_StatusARM `json:"includedPaths,omitempty"`

	//IndexingMode: Indicates the indexing mode.
	IndexingMode *IndexingPolicy_IndexingMode_Status `json:"indexingMode,omitempty"`

	//SpatialIndexes: List of spatial specifics
	SpatialIndexes []SpatialSpec_StatusARM `json:"spatialIndexes,omitempty"`
}

type UniqueKeyPolicy_StatusARM struct {
	//UniqueKeys: List of unique keys on that enforces uniqueness constraint on
	//documents in the collection in the Azure Cosmos DB service.
	UniqueKeys []UniqueKey_StatusARM `json:"uniqueKeys,omitempty"`
}

type CompositePath_StatusARM struct {
	//Order: Sort order for composite paths.
	Order *CompositePath_Order_Status `json:"order,omitempty"`

	//Path: The path for which the indexing behavior applies to. Index paths typically
	//start with root and end with wildcard (/path/*)
	Path *string `json:"path,omitempty"`
}

type ConflictResolutionPolicy_Mode_Status string

const (
	ConflictResolutionPolicy_Mode_StatusCustom         = ConflictResolutionPolicy_Mode_Status("Custom")
	ConflictResolutionPolicy_Mode_StatusLastWriterWins = ConflictResolutionPolicy_Mode_Status("LastWriterWins")
)

type ContainerPartitionKey_Kind_Status string

const (
	ContainerPartitionKey_Kind_StatusHash      = ContainerPartitionKey_Kind_Status("Hash")
	ContainerPartitionKey_Kind_StatusMultiHash = ContainerPartitionKey_Kind_Status("MultiHash")
	ContainerPartitionKey_Kind_StatusRange     = ContainerPartitionKey_Kind_Status("Range")
)

type ExcludedPath_StatusARM struct {
	//Path: The path for which the indexing behavior applies to. Index paths typically
	//start with root and end with wildcard (/path/*)
	Path *string `json:"path,omitempty"`
}

type IncludedPath_StatusARM struct {
	//Indexes: List of indexes for this path
	Indexes []Indexes_StatusARM `json:"indexes,omitempty"`

	//Path: The path for which the indexing behavior applies to. Index paths typically
	//start with root and end with wildcard (/path/*)
	Path *string `json:"path,omitempty"`
}

type IndexingPolicy_IndexingMode_Status string

const (
	IndexingPolicy_IndexingMode_StatusConsistent = IndexingPolicy_IndexingMode_Status("consistent")
	IndexingPolicy_IndexingMode_StatusLazy       = IndexingPolicy_IndexingMode_Status("lazy")
	IndexingPolicy_IndexingMode_StatusNone       = IndexingPolicy_IndexingMode_Status("none")
)

type SpatialSpec_StatusARM struct {
	//Path: The path for which the indexing behavior applies to. Index paths typically
	//start with root and end with wildcard (/path/*)
	Path *string `json:"path,omitempty"`

	//Types: List of path's spatial type
	Types []SpatialType_Status `json:"types,omitempty"`
}

type UniqueKey_StatusARM struct {
	//Paths: List of paths must be unique for each document in the Azure Cosmos DB
	//service
	Paths []string `json:"paths,omitempty"`
}

type CompositePath_Order_Status string

const (
	CompositePath_Order_StatusAscending  = CompositePath_Order_Status("ascending")
	CompositePath_Order_StatusDescending = CompositePath_Order_Status("descending")
)

type Indexes_StatusARM struct {
	//DataType: The datatype for which the indexing behavior is applied to.
	DataType *Indexes_DataType_Status `json:"dataType,omitempty"`

	//Kind: Indicates the type of index.
	Kind *Indexes_Kind_Status `json:"kind,omitempty"`

	//Precision: The precision of the index. -1 is maximum precision.
	Precision *int `json:"precision,omitempty"`
}

type SpatialType_Status string

const (
	SpatialType_StatusLineString   = SpatialType_Status("LineString")
	SpatialType_StatusMultiPolygon = SpatialType_Status("MultiPolygon")
	SpatialType_StatusPoint        = SpatialType_Status("Point")
	SpatialType_StatusPolygon      = SpatialType_Status("Polygon")
)

type Indexes_DataType_Status string

const (
	Indexes_DataType_StatusLineString   = Indexes_DataType_Status("LineString")
	Indexes_DataType_StatusMultiPolygon = Indexes_DataType_Status("MultiPolygon")
	Indexes_DataType_StatusNumber       = Indexes_DataType_Status("Number")
	Indexes_DataType_StatusPoint        = Indexes_DataType_Status("Point")
	Indexes_DataType_StatusPolygon      = Indexes_DataType_Status("Polygon")
	Indexes_DataType_StatusString       = Indexes_DataType_Status("String")
)

type Indexes_Kind_Status string

const (
	Indexes_Kind_StatusHash    = Indexes_Kind_Status("Hash")
	Indexes_Kind_StatusRange   = Indexes_Kind_Status("Range")
	Indexes_Kind_StatusSpatial = Indexes_Kind_Status("Spatial")
)

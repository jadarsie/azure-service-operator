// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_SqlContainer_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlContainer_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlContainer_StatusARM, SqlContainer_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlContainer_StatusARM runs a test to see if a specific instance of SqlContainer_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlContainer_StatusARM(subject SqlContainer_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlContainer_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SqlContainer_StatusARM instances for property testing - lazily instantiated by
//SqlContainer_StatusARMGenerator()
var sqlContainer_statusARMGenerator gopter.Gen

// SqlContainer_StatusARMGenerator returns a generator of SqlContainer_StatusARM instances for property testing.
// We first initialize sqlContainer_statusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlContainer_StatusARMGenerator() gopter.Gen {
	if sqlContainer_statusARMGenerator != nil {
		return sqlContainer_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlContainer_StatusARM(generators)
	sqlContainer_statusARMGenerator = gen.Struct(reflect.TypeOf(SqlContainer_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlContainer_StatusARM(generators)
	AddRelatedPropertyGeneratorsForSqlContainer_StatusARM(generators)
	sqlContainer_statusARMGenerator = gen.Struct(reflect.TypeOf(SqlContainer_StatusARM{}), generators)

	return sqlContainer_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlContainer_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlContainer_StatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlContainer_StatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlContainer_StatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SqlContainerProperties_StatusARMGenerator())
}

func Test_SqlContainerProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlContainerProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlContainerProperties_StatusARM, SqlContainerProperties_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlContainerProperties_StatusARM runs a test to see if a specific instance of SqlContainerProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlContainerProperties_StatusARM(subject SqlContainerProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlContainerProperties_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SqlContainerProperties_StatusARM instances for property testing - lazily instantiated by
//SqlContainerProperties_StatusARMGenerator()
var sqlContainerProperties_statusARMGenerator gopter.Gen

// SqlContainerProperties_StatusARMGenerator returns a generator of SqlContainerProperties_StatusARM instances for property testing.
func SqlContainerProperties_StatusARMGenerator() gopter.Gen {
	if sqlContainerProperties_statusARMGenerator != nil {
		return sqlContainerProperties_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlContainerProperties_StatusARM(generators)
	sqlContainerProperties_statusARMGenerator = gen.Struct(reflect.TypeOf(SqlContainerProperties_StatusARM{}), generators)

	return sqlContainerProperties_statusARMGenerator
}

// AddRelatedPropertyGeneratorsForSqlContainerProperties_StatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlContainerProperties_StatusARM(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptions_StatusARMGenerator())
	gens["Resource"] = SqlContainerResource_StatusARMGenerator()
}

func Test_SqlContainerResource_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlContainerResource_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlContainerResource_StatusARM, SqlContainerResource_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlContainerResource_StatusARM runs a test to see if a specific instance of SqlContainerResource_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlContainerResource_StatusARM(subject SqlContainerResource_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlContainerResource_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SqlContainerResource_StatusARM instances for property testing - lazily instantiated by
//SqlContainerResource_StatusARMGenerator()
var sqlContainerResource_statusARMGenerator gopter.Gen

// SqlContainerResource_StatusARMGenerator returns a generator of SqlContainerResource_StatusARM instances for property testing.
// We first initialize sqlContainerResource_statusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlContainerResource_StatusARMGenerator() gopter.Gen {
	if sqlContainerResource_statusARMGenerator != nil {
		return sqlContainerResource_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlContainerResource_StatusARM(generators)
	sqlContainerResource_statusARMGenerator = gen.Struct(reflect.TypeOf(SqlContainerResource_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlContainerResource_StatusARM(generators)
	AddRelatedPropertyGeneratorsForSqlContainerResource_StatusARM(generators)
	sqlContainerResource_statusARMGenerator = gen.Struct(reflect.TypeOf(SqlContainerResource_StatusARM{}), generators)

	return sqlContainerResource_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlContainerResource_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlContainerResource_StatusARM(gens map[string]gopter.Gen) {
	gens["AnalyticalStorageTtl"] = gen.PtrOf(gen.Int())
	gens["DefaultTtl"] = gen.PtrOf(gen.Int())
	gens["Id"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForSqlContainerResource_StatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlContainerResource_StatusARM(gens map[string]gopter.Gen) {
	gens["ConflictResolutionPolicy"] = gen.PtrOf(ConflictResolutionPolicy_StatusARMGenerator())
	gens["IndexingPolicy"] = gen.PtrOf(IndexingPolicy_StatusARMGenerator())
	gens["PartitionKey"] = gen.PtrOf(ContainerPartitionKey_StatusARMGenerator())
	gens["UniqueKeyPolicy"] = gen.PtrOf(UniqueKeyPolicy_StatusARMGenerator())
}

func Test_ConflictResolutionPolicy_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ConflictResolutionPolicy_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConflictResolutionPolicy_StatusARM, ConflictResolutionPolicy_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConflictResolutionPolicy_StatusARM runs a test to see if a specific instance of ConflictResolutionPolicy_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForConflictResolutionPolicy_StatusARM(subject ConflictResolutionPolicy_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ConflictResolutionPolicy_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ConflictResolutionPolicy_StatusARM instances for property testing - lazily instantiated by
//ConflictResolutionPolicy_StatusARMGenerator()
var conflictResolutionPolicy_statusARMGenerator gopter.Gen

// ConflictResolutionPolicy_StatusARMGenerator returns a generator of ConflictResolutionPolicy_StatusARM instances for property testing.
func ConflictResolutionPolicy_StatusARMGenerator() gopter.Gen {
	if conflictResolutionPolicy_statusARMGenerator != nil {
		return conflictResolutionPolicy_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConflictResolutionPolicy_StatusARM(generators)
	conflictResolutionPolicy_statusARMGenerator = gen.Struct(reflect.TypeOf(ConflictResolutionPolicy_StatusARM{}), generators)

	return conflictResolutionPolicy_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForConflictResolutionPolicy_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForConflictResolutionPolicy_StatusARM(gens map[string]gopter.Gen) {
	gens["ConflictResolutionPath"] = gen.PtrOf(gen.AlphaString())
	gens["ConflictResolutionProcedure"] = gen.PtrOf(gen.AlphaString())
	gens["Mode"] = gen.PtrOf(gen.AlphaString())
}

func Test_ContainerPartitionKey_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ContainerPartitionKey_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForContainerPartitionKey_StatusARM, ContainerPartitionKey_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForContainerPartitionKey_StatusARM runs a test to see if a specific instance of ContainerPartitionKey_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForContainerPartitionKey_StatusARM(subject ContainerPartitionKey_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ContainerPartitionKey_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ContainerPartitionKey_StatusARM instances for property testing - lazily instantiated by
//ContainerPartitionKey_StatusARMGenerator()
var containerPartitionKey_statusARMGenerator gopter.Gen

// ContainerPartitionKey_StatusARMGenerator returns a generator of ContainerPartitionKey_StatusARM instances for property testing.
func ContainerPartitionKey_StatusARMGenerator() gopter.Gen {
	if containerPartitionKey_statusARMGenerator != nil {
		return containerPartitionKey_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForContainerPartitionKey_StatusARM(generators)
	containerPartitionKey_statusARMGenerator = gen.Struct(reflect.TypeOf(ContainerPartitionKey_StatusARM{}), generators)

	return containerPartitionKey_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForContainerPartitionKey_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForContainerPartitionKey_StatusARM(gens map[string]gopter.Gen) {
	gens["Kind"] = gen.PtrOf(gen.AlphaString())
	gens["Paths"] = gen.SliceOf(gen.AlphaString())
	gens["SystemKey"] = gen.PtrOf(gen.Bool())
	gens["Version"] = gen.PtrOf(gen.Int())
}

func Test_IndexingPolicy_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IndexingPolicy_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIndexingPolicy_StatusARM, IndexingPolicy_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIndexingPolicy_StatusARM runs a test to see if a specific instance of IndexingPolicy_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIndexingPolicy_StatusARM(subject IndexingPolicy_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IndexingPolicy_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of IndexingPolicy_StatusARM instances for property testing - lazily instantiated by
//IndexingPolicy_StatusARMGenerator()
var indexingPolicy_statusARMGenerator gopter.Gen

// IndexingPolicy_StatusARMGenerator returns a generator of IndexingPolicy_StatusARM instances for property testing.
// We first initialize indexingPolicy_statusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func IndexingPolicy_StatusARMGenerator() gopter.Gen {
	if indexingPolicy_statusARMGenerator != nil {
		return indexingPolicy_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIndexingPolicy_StatusARM(generators)
	indexingPolicy_statusARMGenerator = gen.Struct(reflect.TypeOf(IndexingPolicy_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIndexingPolicy_StatusARM(generators)
	AddRelatedPropertyGeneratorsForIndexingPolicy_StatusARM(generators)
	indexingPolicy_statusARMGenerator = gen.Struct(reflect.TypeOf(IndexingPolicy_StatusARM{}), generators)

	return indexingPolicy_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForIndexingPolicy_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIndexingPolicy_StatusARM(gens map[string]gopter.Gen) {
	gens["Automatic"] = gen.PtrOf(gen.Bool())
	gens["IndexingMode"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForIndexingPolicy_StatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIndexingPolicy_StatusARM(gens map[string]gopter.Gen) {
	gens["CompositeIndexes"] = gen.SliceOf(gen.SliceOf(CompositePath_StatusARMGenerator()))
	gens["ExcludedPaths"] = gen.SliceOf(ExcludedPath_StatusARMGenerator())
	gens["IncludedPaths"] = gen.SliceOf(IncludedPath_StatusARMGenerator())
	gens["SpatialIndexes"] = gen.SliceOf(SpatialSpec_StatusARMGenerator())
}

func Test_UniqueKeyPolicy_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UniqueKeyPolicy_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUniqueKeyPolicy_StatusARM, UniqueKeyPolicy_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUniqueKeyPolicy_StatusARM runs a test to see if a specific instance of UniqueKeyPolicy_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUniqueKeyPolicy_StatusARM(subject UniqueKeyPolicy_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UniqueKeyPolicy_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of UniqueKeyPolicy_StatusARM instances for property testing - lazily instantiated by
//UniqueKeyPolicy_StatusARMGenerator()
var uniqueKeyPolicy_statusARMGenerator gopter.Gen

// UniqueKeyPolicy_StatusARMGenerator returns a generator of UniqueKeyPolicy_StatusARM instances for property testing.
func UniqueKeyPolicy_StatusARMGenerator() gopter.Gen {
	if uniqueKeyPolicy_statusARMGenerator != nil {
		return uniqueKeyPolicy_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForUniqueKeyPolicy_StatusARM(generators)
	uniqueKeyPolicy_statusARMGenerator = gen.Struct(reflect.TypeOf(UniqueKeyPolicy_StatusARM{}), generators)

	return uniqueKeyPolicy_statusARMGenerator
}

// AddRelatedPropertyGeneratorsForUniqueKeyPolicy_StatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForUniqueKeyPolicy_StatusARM(gens map[string]gopter.Gen) {
	gens["UniqueKeys"] = gen.SliceOf(UniqueKey_StatusARMGenerator())
}

func Test_CompositePath_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CompositePath_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCompositePath_StatusARM, CompositePath_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCompositePath_StatusARM runs a test to see if a specific instance of CompositePath_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCompositePath_StatusARM(subject CompositePath_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CompositePath_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of CompositePath_StatusARM instances for property testing - lazily instantiated by
//CompositePath_StatusARMGenerator()
var compositePath_statusARMGenerator gopter.Gen

// CompositePath_StatusARMGenerator returns a generator of CompositePath_StatusARM instances for property testing.
func CompositePath_StatusARMGenerator() gopter.Gen {
	if compositePath_statusARMGenerator != nil {
		return compositePath_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCompositePath_StatusARM(generators)
	compositePath_statusARMGenerator = gen.Struct(reflect.TypeOf(CompositePath_StatusARM{}), generators)

	return compositePath_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForCompositePath_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCompositePath_StatusARM(gens map[string]gopter.Gen) {
	gens["Order"] = gen.PtrOf(gen.AlphaString())
	gens["Path"] = gen.PtrOf(gen.AlphaString())
}

func Test_ExcludedPath_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ExcludedPath_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForExcludedPath_StatusARM, ExcludedPath_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForExcludedPath_StatusARM runs a test to see if a specific instance of ExcludedPath_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForExcludedPath_StatusARM(subject ExcludedPath_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ExcludedPath_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ExcludedPath_StatusARM instances for property testing - lazily instantiated by
//ExcludedPath_StatusARMGenerator()
var excludedPath_statusARMGenerator gopter.Gen

// ExcludedPath_StatusARMGenerator returns a generator of ExcludedPath_StatusARM instances for property testing.
func ExcludedPath_StatusARMGenerator() gopter.Gen {
	if excludedPath_statusARMGenerator != nil {
		return excludedPath_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForExcludedPath_StatusARM(generators)
	excludedPath_statusARMGenerator = gen.Struct(reflect.TypeOf(ExcludedPath_StatusARM{}), generators)

	return excludedPath_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForExcludedPath_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForExcludedPath_StatusARM(gens map[string]gopter.Gen) {
	gens["Path"] = gen.PtrOf(gen.AlphaString())
}

func Test_IncludedPath_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IncludedPath_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIncludedPath_StatusARM, IncludedPath_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIncludedPath_StatusARM runs a test to see if a specific instance of IncludedPath_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIncludedPath_StatusARM(subject IncludedPath_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IncludedPath_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of IncludedPath_StatusARM instances for property testing - lazily instantiated by
//IncludedPath_StatusARMGenerator()
var includedPath_statusARMGenerator gopter.Gen

// IncludedPath_StatusARMGenerator returns a generator of IncludedPath_StatusARM instances for property testing.
// We first initialize includedPath_statusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func IncludedPath_StatusARMGenerator() gopter.Gen {
	if includedPath_statusARMGenerator != nil {
		return includedPath_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIncludedPath_StatusARM(generators)
	includedPath_statusARMGenerator = gen.Struct(reflect.TypeOf(IncludedPath_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIncludedPath_StatusARM(generators)
	AddRelatedPropertyGeneratorsForIncludedPath_StatusARM(generators)
	includedPath_statusARMGenerator = gen.Struct(reflect.TypeOf(IncludedPath_StatusARM{}), generators)

	return includedPath_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForIncludedPath_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIncludedPath_StatusARM(gens map[string]gopter.Gen) {
	gens["Path"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForIncludedPath_StatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIncludedPath_StatusARM(gens map[string]gopter.Gen) {
	gens["Indexes"] = gen.SliceOf(Indexes_StatusARMGenerator())
}

func Test_SpatialSpec_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SpatialSpec_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSpatialSpec_StatusARM, SpatialSpec_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSpatialSpec_StatusARM runs a test to see if a specific instance of SpatialSpec_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSpatialSpec_StatusARM(subject SpatialSpec_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SpatialSpec_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SpatialSpec_StatusARM instances for property testing - lazily instantiated by
//SpatialSpec_StatusARMGenerator()
var spatialSpec_statusARMGenerator gopter.Gen

// SpatialSpec_StatusARMGenerator returns a generator of SpatialSpec_StatusARM instances for property testing.
func SpatialSpec_StatusARMGenerator() gopter.Gen {
	if spatialSpec_statusARMGenerator != nil {
		return spatialSpec_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSpatialSpec_StatusARM(generators)
	spatialSpec_statusARMGenerator = gen.Struct(reflect.TypeOf(SpatialSpec_StatusARM{}), generators)

	return spatialSpec_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForSpatialSpec_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSpatialSpec_StatusARM(gens map[string]gopter.Gen) {
	gens["Path"] = gen.PtrOf(gen.AlphaString())
	gens["Types"] = gen.SliceOf(gen.AlphaString())
}

func Test_UniqueKey_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UniqueKey_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUniqueKey_StatusARM, UniqueKey_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUniqueKey_StatusARM runs a test to see if a specific instance of UniqueKey_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUniqueKey_StatusARM(subject UniqueKey_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UniqueKey_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of UniqueKey_StatusARM instances for property testing - lazily instantiated by
//UniqueKey_StatusARMGenerator()
var uniqueKey_statusARMGenerator gopter.Gen

// UniqueKey_StatusARMGenerator returns a generator of UniqueKey_StatusARM instances for property testing.
func UniqueKey_StatusARMGenerator() gopter.Gen {
	if uniqueKey_statusARMGenerator != nil {
		return uniqueKey_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUniqueKey_StatusARM(generators)
	uniqueKey_statusARMGenerator = gen.Struct(reflect.TypeOf(UniqueKey_StatusARM{}), generators)

	return uniqueKey_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForUniqueKey_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUniqueKey_StatusARM(gens map[string]gopter.Gen) {
	gens["Paths"] = gen.SliceOf(gen.AlphaString())
}

func Test_Indexes_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Indexes_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIndexes_StatusARM, Indexes_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIndexes_StatusARM runs a test to see if a specific instance of Indexes_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIndexes_StatusARM(subject Indexes_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Indexes_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Indexes_StatusARM instances for property testing - lazily instantiated by Indexes_StatusARMGenerator()
var indexes_statusARMGenerator gopter.Gen

// Indexes_StatusARMGenerator returns a generator of Indexes_StatusARM instances for property testing.
func Indexes_StatusARMGenerator() gopter.Gen {
	if indexes_statusARMGenerator != nil {
		return indexes_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIndexes_StatusARM(generators)
	indexes_statusARMGenerator = gen.Struct(reflect.TypeOf(Indexes_StatusARM{}), generators)

	return indexes_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForIndexes_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIndexes_StatusARM(gens map[string]gopter.Gen) {
	gens["DataType"] = gen.PtrOf(gen.AlphaString())
	gens["Kind"] = gen.PtrOf(gen.AlphaString())
	gens["Precision"] = gen.PtrOf(gen.Int())
}

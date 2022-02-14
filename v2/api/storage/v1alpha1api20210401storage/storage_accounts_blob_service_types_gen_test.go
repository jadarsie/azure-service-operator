// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401storage

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

func Test_StorageAccountsBlobService_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsBlobService via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsBlobService, StorageAccountsBlobServiceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsBlobService runs a test to see if a specific instance of StorageAccountsBlobService round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsBlobService(subject StorageAccountsBlobService) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsBlobService
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

// Generator of StorageAccountsBlobService instances for property testing - lazily instantiated by
//StorageAccountsBlobServiceGenerator()
var storageAccountsBlobServiceGenerator gopter.Gen

// StorageAccountsBlobServiceGenerator returns a generator of StorageAccountsBlobService instances for property testing.
func StorageAccountsBlobServiceGenerator() gopter.Gen {
	if storageAccountsBlobServiceGenerator != nil {
		return storageAccountsBlobServiceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForStorageAccountsBlobService(generators)
	storageAccountsBlobServiceGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobService{}), generators)

	return storageAccountsBlobServiceGenerator
}

// AddRelatedPropertyGeneratorsForStorageAccountsBlobService is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsBlobService(gens map[string]gopter.Gen) {
	gens["Spec"] = StorageAccountsBlobService_SpecGenerator()
	gens["Status"] = BlobServiceProperties_StatusGenerator()
}

func Test_BlobServiceProperties_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BlobServiceProperties_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBlobServiceProperties_Status, BlobServiceProperties_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBlobServiceProperties_Status runs a test to see if a specific instance of BlobServiceProperties_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForBlobServiceProperties_Status(subject BlobServiceProperties_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BlobServiceProperties_Status
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

// Generator of BlobServiceProperties_Status instances for property testing - lazily instantiated by
//BlobServiceProperties_StatusGenerator()
var blobServiceProperties_statusGenerator gopter.Gen

// BlobServiceProperties_StatusGenerator returns a generator of BlobServiceProperties_Status instances for property testing.
// We first initialize blobServiceProperties_statusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BlobServiceProperties_StatusGenerator() gopter.Gen {
	if blobServiceProperties_statusGenerator != nil {
		return blobServiceProperties_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobServiceProperties_Status(generators)
	blobServiceProperties_statusGenerator = gen.Struct(reflect.TypeOf(BlobServiceProperties_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobServiceProperties_Status(generators)
	AddRelatedPropertyGeneratorsForBlobServiceProperties_Status(generators)
	blobServiceProperties_statusGenerator = gen.Struct(reflect.TypeOf(BlobServiceProperties_Status{}), generators)

	return blobServiceProperties_statusGenerator
}

// AddIndependentPropertyGeneratorsForBlobServiceProperties_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBlobServiceProperties_Status(gens map[string]gopter.Gen) {
	gens["AutomaticSnapshotPolicyEnabled"] = gen.PtrOf(gen.Bool())
	gens["DefaultServiceVersion"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IsVersioningEnabled"] = gen.PtrOf(gen.Bool())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForBlobServiceProperties_Status is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBlobServiceProperties_Status(gens map[string]gopter.Gen) {
	gens["ChangeFeed"] = gen.PtrOf(ChangeFeed_StatusGenerator())
	gens["ContainerDeleteRetentionPolicy"] = gen.PtrOf(DeleteRetentionPolicy_StatusGenerator())
	gens["Cors"] = gen.PtrOf(CorsRules_StatusGenerator())
	gens["DeleteRetentionPolicy"] = gen.PtrOf(DeleteRetentionPolicy_StatusGenerator())
	gens["LastAccessTimeTrackingPolicy"] = gen.PtrOf(LastAccessTimeTrackingPolicy_StatusGenerator())
	gens["RestorePolicy"] = gen.PtrOf(RestorePolicyProperties_StatusGenerator())
	gens["Sku"] = gen.PtrOf(Sku_StatusGenerator())
}

func Test_StorageAccountsBlobService_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsBlobService_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsBlobService_Spec, StorageAccountsBlobService_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsBlobService_Spec runs a test to see if a specific instance of StorageAccountsBlobService_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsBlobService_Spec(subject StorageAccountsBlobService_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsBlobService_Spec
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

// Generator of StorageAccountsBlobService_Spec instances for property testing - lazily instantiated by
//StorageAccountsBlobService_SpecGenerator()
var storageAccountsBlobService_specGenerator gopter.Gen

// StorageAccountsBlobService_SpecGenerator returns a generator of StorageAccountsBlobService_Spec instances for property testing.
// We first initialize storageAccountsBlobService_specGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccountsBlobService_SpecGenerator() gopter.Gen {
	if storageAccountsBlobService_specGenerator != nil {
		return storageAccountsBlobService_specGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsBlobService_Spec(generators)
	storageAccountsBlobService_specGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobService_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsBlobService_Spec(generators)
	AddRelatedPropertyGeneratorsForStorageAccountsBlobService_Spec(generators)
	storageAccountsBlobService_specGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobService_Spec{}), generators)

	return storageAccountsBlobService_specGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsBlobService_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsBlobService_Spec(gens map[string]gopter.Gen) {
	gens["AutomaticSnapshotPolicyEnabled"] = gen.PtrOf(gen.Bool())
	gens["AzureName"] = gen.AlphaString()
	gens["DefaultServiceVersion"] = gen.PtrOf(gen.AlphaString())
	gens["IsVersioningEnabled"] = gen.PtrOf(gen.Bool())
	gens["OriginalVersion"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForStorageAccountsBlobService_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsBlobService_Spec(gens map[string]gopter.Gen) {
	gens["ChangeFeed"] = gen.PtrOf(ChangeFeedGenerator())
	gens["ContainerDeleteRetentionPolicy"] = gen.PtrOf(DeleteRetentionPolicyGenerator())
	gens["Cors"] = gen.PtrOf(CorsRulesGenerator())
	gens["DeleteRetentionPolicy"] = gen.PtrOf(DeleteRetentionPolicyGenerator())
	gens["LastAccessTimeTrackingPolicy"] = gen.PtrOf(LastAccessTimeTrackingPolicyGenerator())
	gens["RestorePolicy"] = gen.PtrOf(RestorePolicyPropertiesGenerator())
}

func Test_ChangeFeed_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ChangeFeed via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForChangeFeed, ChangeFeedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForChangeFeed runs a test to see if a specific instance of ChangeFeed round trips to JSON and back losslessly
func RunJSONSerializationTestForChangeFeed(subject ChangeFeed) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ChangeFeed
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

// Generator of ChangeFeed instances for property testing - lazily instantiated by ChangeFeedGenerator()
var changeFeedGenerator gopter.Gen

// ChangeFeedGenerator returns a generator of ChangeFeed instances for property testing.
func ChangeFeedGenerator() gopter.Gen {
	if changeFeedGenerator != nil {
		return changeFeedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForChangeFeed(generators)
	changeFeedGenerator = gen.Struct(reflect.TypeOf(ChangeFeed{}), generators)

	return changeFeedGenerator
}

// AddIndependentPropertyGeneratorsForChangeFeed is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForChangeFeed(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["RetentionInDays"] = gen.PtrOf(gen.Int())
}

func Test_ChangeFeed_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ChangeFeed_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForChangeFeed_Status, ChangeFeed_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForChangeFeed_Status runs a test to see if a specific instance of ChangeFeed_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForChangeFeed_Status(subject ChangeFeed_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ChangeFeed_Status
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

// Generator of ChangeFeed_Status instances for property testing - lazily instantiated by ChangeFeed_StatusGenerator()
var changeFeed_statusGenerator gopter.Gen

// ChangeFeed_StatusGenerator returns a generator of ChangeFeed_Status instances for property testing.
func ChangeFeed_StatusGenerator() gopter.Gen {
	if changeFeed_statusGenerator != nil {
		return changeFeed_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForChangeFeed_Status(generators)
	changeFeed_statusGenerator = gen.Struct(reflect.TypeOf(ChangeFeed_Status{}), generators)

	return changeFeed_statusGenerator
}

// AddIndependentPropertyGeneratorsForChangeFeed_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForChangeFeed_Status(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["RetentionInDays"] = gen.PtrOf(gen.Int())
}

func Test_CorsRules_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CorsRules via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCorsRules, CorsRulesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCorsRules runs a test to see if a specific instance of CorsRules round trips to JSON and back losslessly
func RunJSONSerializationTestForCorsRules(subject CorsRules) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CorsRules
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

// Generator of CorsRules instances for property testing - lazily instantiated by CorsRulesGenerator()
var corsRulesGenerator gopter.Gen

// CorsRulesGenerator returns a generator of CorsRules instances for property testing.
func CorsRulesGenerator() gopter.Gen {
	if corsRulesGenerator != nil {
		return corsRulesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForCorsRules(generators)
	corsRulesGenerator = gen.Struct(reflect.TypeOf(CorsRules{}), generators)

	return corsRulesGenerator
}

// AddRelatedPropertyGeneratorsForCorsRules is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCorsRules(gens map[string]gopter.Gen) {
	gens["CorsRules"] = gen.SliceOf(CorsRuleGenerator())
}

func Test_CorsRules_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CorsRules_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCorsRules_Status, CorsRules_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCorsRules_Status runs a test to see if a specific instance of CorsRules_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForCorsRules_Status(subject CorsRules_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CorsRules_Status
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

// Generator of CorsRules_Status instances for property testing - lazily instantiated by CorsRules_StatusGenerator()
var corsRules_statusGenerator gopter.Gen

// CorsRules_StatusGenerator returns a generator of CorsRules_Status instances for property testing.
func CorsRules_StatusGenerator() gopter.Gen {
	if corsRules_statusGenerator != nil {
		return corsRules_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForCorsRules_Status(generators)
	corsRules_statusGenerator = gen.Struct(reflect.TypeOf(CorsRules_Status{}), generators)

	return corsRules_statusGenerator
}

// AddRelatedPropertyGeneratorsForCorsRules_Status is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCorsRules_Status(gens map[string]gopter.Gen) {
	gens["CorsRules"] = gen.SliceOf(CorsRule_StatusGenerator())
}

func Test_DeleteRetentionPolicy_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DeleteRetentionPolicy via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDeleteRetentionPolicy, DeleteRetentionPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDeleteRetentionPolicy runs a test to see if a specific instance of DeleteRetentionPolicy round trips to JSON and back losslessly
func RunJSONSerializationTestForDeleteRetentionPolicy(subject DeleteRetentionPolicy) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DeleteRetentionPolicy
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

// Generator of DeleteRetentionPolicy instances for property testing - lazily instantiated by
//DeleteRetentionPolicyGenerator()
var deleteRetentionPolicyGenerator gopter.Gen

// DeleteRetentionPolicyGenerator returns a generator of DeleteRetentionPolicy instances for property testing.
func DeleteRetentionPolicyGenerator() gopter.Gen {
	if deleteRetentionPolicyGenerator != nil {
		return deleteRetentionPolicyGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDeleteRetentionPolicy(generators)
	deleteRetentionPolicyGenerator = gen.Struct(reflect.TypeOf(DeleteRetentionPolicy{}), generators)

	return deleteRetentionPolicyGenerator
}

// AddIndependentPropertyGeneratorsForDeleteRetentionPolicy is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDeleteRetentionPolicy(gens map[string]gopter.Gen) {
	gens["Days"] = gen.PtrOf(gen.Int())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

func Test_DeleteRetentionPolicy_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DeleteRetentionPolicy_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDeleteRetentionPolicy_Status, DeleteRetentionPolicy_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDeleteRetentionPolicy_Status runs a test to see if a specific instance of DeleteRetentionPolicy_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForDeleteRetentionPolicy_Status(subject DeleteRetentionPolicy_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DeleteRetentionPolicy_Status
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

// Generator of DeleteRetentionPolicy_Status instances for property testing - lazily instantiated by
//DeleteRetentionPolicy_StatusGenerator()
var deleteRetentionPolicy_statusGenerator gopter.Gen

// DeleteRetentionPolicy_StatusGenerator returns a generator of DeleteRetentionPolicy_Status instances for property testing.
func DeleteRetentionPolicy_StatusGenerator() gopter.Gen {
	if deleteRetentionPolicy_statusGenerator != nil {
		return deleteRetentionPolicy_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDeleteRetentionPolicy_Status(generators)
	deleteRetentionPolicy_statusGenerator = gen.Struct(reflect.TypeOf(DeleteRetentionPolicy_Status{}), generators)

	return deleteRetentionPolicy_statusGenerator
}

// AddIndependentPropertyGeneratorsForDeleteRetentionPolicy_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDeleteRetentionPolicy_Status(gens map[string]gopter.Gen) {
	gens["Days"] = gen.PtrOf(gen.Int())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

func Test_LastAccessTimeTrackingPolicy_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LastAccessTimeTrackingPolicy via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLastAccessTimeTrackingPolicy, LastAccessTimeTrackingPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLastAccessTimeTrackingPolicy runs a test to see if a specific instance of LastAccessTimeTrackingPolicy round trips to JSON and back losslessly
func RunJSONSerializationTestForLastAccessTimeTrackingPolicy(subject LastAccessTimeTrackingPolicy) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LastAccessTimeTrackingPolicy
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

// Generator of LastAccessTimeTrackingPolicy instances for property testing - lazily instantiated by
//LastAccessTimeTrackingPolicyGenerator()
var lastAccessTimeTrackingPolicyGenerator gopter.Gen

// LastAccessTimeTrackingPolicyGenerator returns a generator of LastAccessTimeTrackingPolicy instances for property testing.
func LastAccessTimeTrackingPolicyGenerator() gopter.Gen {
	if lastAccessTimeTrackingPolicyGenerator != nil {
		return lastAccessTimeTrackingPolicyGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicy(generators)
	lastAccessTimeTrackingPolicyGenerator = gen.Struct(reflect.TypeOf(LastAccessTimeTrackingPolicy{}), generators)

	return lastAccessTimeTrackingPolicyGenerator
}

// AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicy is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicy(gens map[string]gopter.Gen) {
	gens["BlobType"] = gen.SliceOf(gen.AlphaString())
	gens["Enable"] = gen.PtrOf(gen.Bool())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["TrackingGranularityInDays"] = gen.PtrOf(gen.Int())
}

func Test_LastAccessTimeTrackingPolicy_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LastAccessTimeTrackingPolicy_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLastAccessTimeTrackingPolicy_Status, LastAccessTimeTrackingPolicy_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLastAccessTimeTrackingPolicy_Status runs a test to see if a specific instance of LastAccessTimeTrackingPolicy_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForLastAccessTimeTrackingPolicy_Status(subject LastAccessTimeTrackingPolicy_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LastAccessTimeTrackingPolicy_Status
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

// Generator of LastAccessTimeTrackingPolicy_Status instances for property testing - lazily instantiated by
//LastAccessTimeTrackingPolicy_StatusGenerator()
var lastAccessTimeTrackingPolicy_statusGenerator gopter.Gen

// LastAccessTimeTrackingPolicy_StatusGenerator returns a generator of LastAccessTimeTrackingPolicy_Status instances for property testing.
func LastAccessTimeTrackingPolicy_StatusGenerator() gopter.Gen {
	if lastAccessTimeTrackingPolicy_statusGenerator != nil {
		return lastAccessTimeTrackingPolicy_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicy_Status(generators)
	lastAccessTimeTrackingPolicy_statusGenerator = gen.Struct(reflect.TypeOf(LastAccessTimeTrackingPolicy_Status{}), generators)

	return lastAccessTimeTrackingPolicy_statusGenerator
}

// AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicy_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicy_Status(gens map[string]gopter.Gen) {
	gens["BlobType"] = gen.SliceOf(gen.AlphaString())
	gens["Enable"] = gen.PtrOf(gen.Bool())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["TrackingGranularityInDays"] = gen.PtrOf(gen.Int())
}

func Test_RestorePolicyProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RestorePolicyProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRestorePolicyProperties, RestorePolicyPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRestorePolicyProperties runs a test to see if a specific instance of RestorePolicyProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForRestorePolicyProperties(subject RestorePolicyProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RestorePolicyProperties
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

// Generator of RestorePolicyProperties instances for property testing - lazily instantiated by
//RestorePolicyPropertiesGenerator()
var restorePolicyPropertiesGenerator gopter.Gen

// RestorePolicyPropertiesGenerator returns a generator of RestorePolicyProperties instances for property testing.
func RestorePolicyPropertiesGenerator() gopter.Gen {
	if restorePolicyPropertiesGenerator != nil {
		return restorePolicyPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRestorePolicyProperties(generators)
	restorePolicyPropertiesGenerator = gen.Struct(reflect.TypeOf(RestorePolicyProperties{}), generators)

	return restorePolicyPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForRestorePolicyProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRestorePolicyProperties(gens map[string]gopter.Gen) {
	gens["Days"] = gen.PtrOf(gen.Int())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

func Test_RestorePolicyProperties_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RestorePolicyProperties_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRestorePolicyProperties_Status, RestorePolicyProperties_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRestorePolicyProperties_Status runs a test to see if a specific instance of RestorePolicyProperties_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForRestorePolicyProperties_Status(subject RestorePolicyProperties_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RestorePolicyProperties_Status
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

// Generator of RestorePolicyProperties_Status instances for property testing - lazily instantiated by
//RestorePolicyProperties_StatusGenerator()
var restorePolicyProperties_statusGenerator gopter.Gen

// RestorePolicyProperties_StatusGenerator returns a generator of RestorePolicyProperties_Status instances for property testing.
func RestorePolicyProperties_StatusGenerator() gopter.Gen {
	if restorePolicyProperties_statusGenerator != nil {
		return restorePolicyProperties_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRestorePolicyProperties_Status(generators)
	restorePolicyProperties_statusGenerator = gen.Struct(reflect.TypeOf(RestorePolicyProperties_Status{}), generators)

	return restorePolicyProperties_statusGenerator
}

// AddIndependentPropertyGeneratorsForRestorePolicyProperties_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRestorePolicyProperties_Status(gens map[string]gopter.Gen) {
	gens["Days"] = gen.PtrOf(gen.Int())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["LastEnabledTime"] = gen.PtrOf(gen.AlphaString())
	gens["MinRestoreTime"] = gen.PtrOf(gen.AlphaString())
}

func Test_CorsRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CorsRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCorsRule, CorsRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCorsRule runs a test to see if a specific instance of CorsRule round trips to JSON and back losslessly
func RunJSONSerializationTestForCorsRule(subject CorsRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CorsRule
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

// Generator of CorsRule instances for property testing - lazily instantiated by CorsRuleGenerator()
var corsRuleGenerator gopter.Gen

// CorsRuleGenerator returns a generator of CorsRule instances for property testing.
func CorsRuleGenerator() gopter.Gen {
	if corsRuleGenerator != nil {
		return corsRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCorsRule(generators)
	corsRuleGenerator = gen.Struct(reflect.TypeOf(CorsRule{}), generators)

	return corsRuleGenerator
}

// AddIndependentPropertyGeneratorsForCorsRule is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCorsRule(gens map[string]gopter.Gen) {
	gens["AllowedHeaders"] = gen.SliceOf(gen.AlphaString())
	gens["AllowedMethods"] = gen.SliceOf(gen.AlphaString())
	gens["AllowedOrigins"] = gen.SliceOf(gen.AlphaString())
	gens["ExposedHeaders"] = gen.SliceOf(gen.AlphaString())
	gens["MaxAgeInSeconds"] = gen.PtrOf(gen.Int())
}

func Test_CorsRule_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CorsRule_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCorsRule_Status, CorsRule_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCorsRule_Status runs a test to see if a specific instance of CorsRule_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForCorsRule_Status(subject CorsRule_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CorsRule_Status
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

// Generator of CorsRule_Status instances for property testing - lazily instantiated by CorsRule_StatusGenerator()
var corsRule_statusGenerator gopter.Gen

// CorsRule_StatusGenerator returns a generator of CorsRule_Status instances for property testing.
func CorsRule_StatusGenerator() gopter.Gen {
	if corsRule_statusGenerator != nil {
		return corsRule_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCorsRule_Status(generators)
	corsRule_statusGenerator = gen.Struct(reflect.TypeOf(CorsRule_Status{}), generators)

	return corsRule_statusGenerator
}

// AddIndependentPropertyGeneratorsForCorsRule_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCorsRule_Status(gens map[string]gopter.Gen) {
	gens["AllowedHeaders"] = gen.SliceOf(gen.AlphaString())
	gens["AllowedMethods"] = gen.SliceOf(gen.AlphaString())
	gens["AllowedOrigins"] = gen.SliceOf(gen.AlphaString())
	gens["ExposedHeaders"] = gen.SliceOf(gen.AlphaString())
	gens["MaxAgeInSeconds"] = gen.PtrOf(gen.Int())
}

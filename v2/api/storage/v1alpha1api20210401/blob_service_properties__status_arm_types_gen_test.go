// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

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

func Test_BlobServiceProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BlobServiceProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBlobServiceProperties_StatusARM, BlobServiceProperties_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBlobServiceProperties_StatusARM runs a test to see if a specific instance of BlobServiceProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBlobServiceProperties_StatusARM(subject BlobServiceProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BlobServiceProperties_StatusARM
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

// Generator of BlobServiceProperties_StatusARM instances for property testing - lazily instantiated by
//BlobServiceProperties_StatusARMGenerator()
var blobServiceProperties_statusARMGenerator gopter.Gen

// BlobServiceProperties_StatusARMGenerator returns a generator of BlobServiceProperties_StatusARM instances for property testing.
// We first initialize blobServiceProperties_statusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BlobServiceProperties_StatusARMGenerator() gopter.Gen {
	if blobServiceProperties_statusARMGenerator != nil {
		return blobServiceProperties_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobServiceProperties_StatusARM(generators)
	blobServiceProperties_statusARMGenerator = gen.Struct(reflect.TypeOf(BlobServiceProperties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobServiceProperties_StatusARM(generators)
	AddRelatedPropertyGeneratorsForBlobServiceProperties_StatusARM(generators)
	blobServiceProperties_statusARMGenerator = gen.Struct(reflect.TypeOf(BlobServiceProperties_StatusARM{}), generators)

	return blobServiceProperties_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForBlobServiceProperties_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBlobServiceProperties_StatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForBlobServiceProperties_StatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBlobServiceProperties_StatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(BlobServiceProperties_Properties_StatusARMGenerator())
	gens["Sku"] = gen.PtrOf(Sku_StatusARMGenerator())
}

func Test_BlobServiceProperties_Properties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BlobServiceProperties_Properties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBlobServiceProperties_Properties_StatusARM, BlobServiceProperties_Properties_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBlobServiceProperties_Properties_StatusARM runs a test to see if a specific instance of BlobServiceProperties_Properties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBlobServiceProperties_Properties_StatusARM(subject BlobServiceProperties_Properties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BlobServiceProperties_Properties_StatusARM
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

// Generator of BlobServiceProperties_Properties_StatusARM instances for property testing - lazily instantiated by
//BlobServiceProperties_Properties_StatusARMGenerator()
var blobServiceProperties_properties_statusARMGenerator gopter.Gen

// BlobServiceProperties_Properties_StatusARMGenerator returns a generator of BlobServiceProperties_Properties_StatusARM instances for property testing.
// We first initialize blobServiceProperties_properties_statusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BlobServiceProperties_Properties_StatusARMGenerator() gopter.Gen {
	if blobServiceProperties_properties_statusARMGenerator != nil {
		return blobServiceProperties_properties_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobServiceProperties_Properties_StatusARM(generators)
	blobServiceProperties_properties_statusARMGenerator = gen.Struct(reflect.TypeOf(BlobServiceProperties_Properties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobServiceProperties_Properties_StatusARM(generators)
	AddRelatedPropertyGeneratorsForBlobServiceProperties_Properties_StatusARM(generators)
	blobServiceProperties_properties_statusARMGenerator = gen.Struct(reflect.TypeOf(BlobServiceProperties_Properties_StatusARM{}), generators)

	return blobServiceProperties_properties_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForBlobServiceProperties_Properties_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBlobServiceProperties_Properties_StatusARM(gens map[string]gopter.Gen) {
	gens["AutomaticSnapshotPolicyEnabled"] = gen.PtrOf(gen.Bool())
	gens["DefaultServiceVersion"] = gen.PtrOf(gen.AlphaString())
	gens["IsVersioningEnabled"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForBlobServiceProperties_Properties_StatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBlobServiceProperties_Properties_StatusARM(gens map[string]gopter.Gen) {
	gens["ChangeFeed"] = gen.PtrOf(ChangeFeed_StatusARMGenerator())
	gens["ContainerDeleteRetentionPolicy"] = gen.PtrOf(DeleteRetentionPolicy_StatusARMGenerator())
	gens["Cors"] = gen.PtrOf(CorsRules_StatusARMGenerator())
	gens["DeleteRetentionPolicy"] = gen.PtrOf(DeleteRetentionPolicy_StatusARMGenerator())
	gens["LastAccessTimeTrackingPolicy"] = gen.PtrOf(LastAccessTimeTrackingPolicy_StatusARMGenerator())
	gens["RestorePolicy"] = gen.PtrOf(RestorePolicyProperties_StatusARMGenerator())
}

func Test_Sku_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku_StatusARM, Sku_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku_StatusARM runs a test to see if a specific instance of Sku_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSku_StatusARM(subject Sku_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_StatusARM
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

// Generator of Sku_StatusARM instances for property testing - lazily instantiated by Sku_StatusARMGenerator()
var sku_statusARMGenerator gopter.Gen

// Sku_StatusARMGenerator returns a generator of Sku_StatusARM instances for property testing.
func Sku_StatusARMGenerator() gopter.Gen {
	if sku_statusARMGenerator != nil {
		return sku_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku_StatusARM(generators)
	sku_statusARMGenerator = gen.Struct(reflect.TypeOf(Sku_StatusARM{}), generators)

	return sku_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForSku_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku_StatusARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.OneConstOf(
		SkuName_StatusPremium_LRS,
		SkuName_StatusPremium_ZRS,
		SkuName_StatusStandard_GRS,
		SkuName_StatusStandard_GZRS,
		SkuName_StatusStandard_LRS,
		SkuName_StatusStandard_RAGRS,
		SkuName_StatusStandard_RAGZRS,
		SkuName_StatusStandard_ZRS)
	gens["Tier"] = gen.PtrOf(gen.OneConstOf(Tier_StatusPremium, Tier_StatusStandard))
}

func Test_ChangeFeed_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ChangeFeed_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForChangeFeed_StatusARM, ChangeFeed_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForChangeFeed_StatusARM runs a test to see if a specific instance of ChangeFeed_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForChangeFeed_StatusARM(subject ChangeFeed_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ChangeFeed_StatusARM
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

// Generator of ChangeFeed_StatusARM instances for property testing - lazily instantiated by
//ChangeFeed_StatusARMGenerator()
var changeFeed_statusARMGenerator gopter.Gen

// ChangeFeed_StatusARMGenerator returns a generator of ChangeFeed_StatusARM instances for property testing.
func ChangeFeed_StatusARMGenerator() gopter.Gen {
	if changeFeed_statusARMGenerator != nil {
		return changeFeed_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForChangeFeed_StatusARM(generators)
	changeFeed_statusARMGenerator = gen.Struct(reflect.TypeOf(ChangeFeed_StatusARM{}), generators)

	return changeFeed_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForChangeFeed_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForChangeFeed_StatusARM(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["RetentionInDays"] = gen.PtrOf(gen.Int())
}

func Test_CorsRules_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CorsRules_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCorsRules_StatusARM, CorsRules_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCorsRules_StatusARM runs a test to see if a specific instance of CorsRules_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCorsRules_StatusARM(subject CorsRules_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CorsRules_StatusARM
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

// Generator of CorsRules_StatusARM instances for property testing - lazily instantiated by
//CorsRules_StatusARMGenerator()
var corsRules_statusARMGenerator gopter.Gen

// CorsRules_StatusARMGenerator returns a generator of CorsRules_StatusARM instances for property testing.
func CorsRules_StatusARMGenerator() gopter.Gen {
	if corsRules_statusARMGenerator != nil {
		return corsRules_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForCorsRules_StatusARM(generators)
	corsRules_statusARMGenerator = gen.Struct(reflect.TypeOf(CorsRules_StatusARM{}), generators)

	return corsRules_statusARMGenerator
}

// AddRelatedPropertyGeneratorsForCorsRules_StatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCorsRules_StatusARM(gens map[string]gopter.Gen) {
	gens["CorsRules"] = gen.SliceOf(CorsRule_StatusARMGenerator())
}

func Test_DeleteRetentionPolicy_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DeleteRetentionPolicy_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDeleteRetentionPolicy_StatusARM, DeleteRetentionPolicy_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDeleteRetentionPolicy_StatusARM runs a test to see if a specific instance of DeleteRetentionPolicy_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDeleteRetentionPolicy_StatusARM(subject DeleteRetentionPolicy_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DeleteRetentionPolicy_StatusARM
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

// Generator of DeleteRetentionPolicy_StatusARM instances for property testing - lazily instantiated by
//DeleteRetentionPolicy_StatusARMGenerator()
var deleteRetentionPolicy_statusARMGenerator gopter.Gen

// DeleteRetentionPolicy_StatusARMGenerator returns a generator of DeleteRetentionPolicy_StatusARM instances for property testing.
func DeleteRetentionPolicy_StatusARMGenerator() gopter.Gen {
	if deleteRetentionPolicy_statusARMGenerator != nil {
		return deleteRetentionPolicy_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDeleteRetentionPolicy_StatusARM(generators)
	deleteRetentionPolicy_statusARMGenerator = gen.Struct(reflect.TypeOf(DeleteRetentionPolicy_StatusARM{}), generators)

	return deleteRetentionPolicy_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForDeleteRetentionPolicy_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDeleteRetentionPolicy_StatusARM(gens map[string]gopter.Gen) {
	gens["Days"] = gen.PtrOf(gen.Int())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

func Test_LastAccessTimeTrackingPolicy_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LastAccessTimeTrackingPolicy_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLastAccessTimeTrackingPolicy_StatusARM, LastAccessTimeTrackingPolicy_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLastAccessTimeTrackingPolicy_StatusARM runs a test to see if a specific instance of LastAccessTimeTrackingPolicy_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForLastAccessTimeTrackingPolicy_StatusARM(subject LastAccessTimeTrackingPolicy_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LastAccessTimeTrackingPolicy_StatusARM
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

// Generator of LastAccessTimeTrackingPolicy_StatusARM instances for property testing - lazily instantiated by
//LastAccessTimeTrackingPolicy_StatusARMGenerator()
var lastAccessTimeTrackingPolicy_statusARMGenerator gopter.Gen

// LastAccessTimeTrackingPolicy_StatusARMGenerator returns a generator of LastAccessTimeTrackingPolicy_StatusARM instances for property testing.
func LastAccessTimeTrackingPolicy_StatusARMGenerator() gopter.Gen {
	if lastAccessTimeTrackingPolicy_statusARMGenerator != nil {
		return lastAccessTimeTrackingPolicy_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicy_StatusARM(generators)
	lastAccessTimeTrackingPolicy_statusARMGenerator = gen.Struct(reflect.TypeOf(LastAccessTimeTrackingPolicy_StatusARM{}), generators)

	return lastAccessTimeTrackingPolicy_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicy_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLastAccessTimeTrackingPolicy_StatusARM(gens map[string]gopter.Gen) {
	gens["BlobType"] = gen.SliceOf(gen.AlphaString())
	gens["Enable"] = gen.Bool()
	gens["Name"] = gen.PtrOf(gen.OneConstOf(LastAccessTimeTrackingPolicy_Name_StatusAccessTimeTracking))
	gens["TrackingGranularityInDays"] = gen.PtrOf(gen.Int())
}

func Test_RestorePolicyProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RestorePolicyProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRestorePolicyProperties_StatusARM, RestorePolicyProperties_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRestorePolicyProperties_StatusARM runs a test to see if a specific instance of RestorePolicyProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRestorePolicyProperties_StatusARM(subject RestorePolicyProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RestorePolicyProperties_StatusARM
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

// Generator of RestorePolicyProperties_StatusARM instances for property testing - lazily instantiated by
//RestorePolicyProperties_StatusARMGenerator()
var restorePolicyProperties_statusARMGenerator gopter.Gen

// RestorePolicyProperties_StatusARMGenerator returns a generator of RestorePolicyProperties_StatusARM instances for property testing.
func RestorePolicyProperties_StatusARMGenerator() gopter.Gen {
	if restorePolicyProperties_statusARMGenerator != nil {
		return restorePolicyProperties_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRestorePolicyProperties_StatusARM(generators)
	restorePolicyProperties_statusARMGenerator = gen.Struct(reflect.TypeOf(RestorePolicyProperties_StatusARM{}), generators)

	return restorePolicyProperties_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForRestorePolicyProperties_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRestorePolicyProperties_StatusARM(gens map[string]gopter.Gen) {
	gens["Days"] = gen.PtrOf(gen.Int())
	gens["Enabled"] = gen.Bool()
	gens["LastEnabledTime"] = gen.PtrOf(gen.AlphaString())
	gens["MinRestoreTime"] = gen.PtrOf(gen.AlphaString())
}

func Test_CorsRule_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CorsRule_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCorsRule_StatusARM, CorsRule_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCorsRule_StatusARM runs a test to see if a specific instance of CorsRule_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCorsRule_StatusARM(subject CorsRule_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CorsRule_StatusARM
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

// Generator of CorsRule_StatusARM instances for property testing - lazily instantiated by CorsRule_StatusARMGenerator()
var corsRule_statusARMGenerator gopter.Gen

// CorsRule_StatusARMGenerator returns a generator of CorsRule_StatusARM instances for property testing.
func CorsRule_StatusARMGenerator() gopter.Gen {
	if corsRule_statusARMGenerator != nil {
		return corsRule_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCorsRule_StatusARM(generators)
	corsRule_statusARMGenerator = gen.Struct(reflect.TypeOf(CorsRule_StatusARM{}), generators)

	return corsRule_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForCorsRule_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCorsRule_StatusARM(gens map[string]gopter.Gen) {
	gens["AllowedHeaders"] = gen.SliceOf(gen.AlphaString())
	gens["AllowedMethods"] = gen.SliceOf(gen.OneConstOf(
		CorsRule_AllowedMethods_StatusDELETE,
		CorsRule_AllowedMethods_StatusGET,
		CorsRule_AllowedMethods_StatusHEAD,
		CorsRule_AllowedMethods_StatusMERGE,
		CorsRule_AllowedMethods_StatusOPTIONS,
		CorsRule_AllowedMethods_StatusPOST,
		CorsRule_AllowedMethods_StatusPUT))
	gens["AllowedOrigins"] = gen.SliceOf(gen.AlphaString())
	gens["ExposedHeaders"] = gen.SliceOf(gen.AlphaString())
	gens["MaxAgeInSeconds"] = gen.Int()
}

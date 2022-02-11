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

func Test_StorageAccountsBlobServicesContainer_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsBlobServicesContainer via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsBlobServicesContainer, StorageAccountsBlobServicesContainerGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsBlobServicesContainer runs a test to see if a specific instance of StorageAccountsBlobServicesContainer round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsBlobServicesContainer(subject StorageAccountsBlobServicesContainer) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsBlobServicesContainer
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

// Generator of StorageAccountsBlobServicesContainer instances for property testing - lazily instantiated by
//StorageAccountsBlobServicesContainerGenerator()
var storageAccountsBlobServicesContainerGenerator gopter.Gen

// StorageAccountsBlobServicesContainerGenerator returns a generator of StorageAccountsBlobServicesContainer instances for property testing.
func StorageAccountsBlobServicesContainerGenerator() gopter.Gen {
	if storageAccountsBlobServicesContainerGenerator != nil {
		return storageAccountsBlobServicesContainerGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesContainer(generators)
	storageAccountsBlobServicesContainerGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobServicesContainer{}), generators)

	return storageAccountsBlobServicesContainerGenerator
}

// AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesContainer is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesContainer(gens map[string]gopter.Gen) {
	gens["Spec"] = StorageAccountsBlobServicesContainers_SPECGenerator()
	gens["Status"] = BlobContainer_StatusGenerator()
}

func Test_BlobContainer_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BlobContainer_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBlobContainer_Status, BlobContainer_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBlobContainer_Status runs a test to see if a specific instance of BlobContainer_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForBlobContainer_Status(subject BlobContainer_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BlobContainer_Status
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

// Generator of BlobContainer_Status instances for property testing - lazily instantiated by
//BlobContainer_StatusGenerator()
var blobContainer_statusGenerator gopter.Gen

// BlobContainer_StatusGenerator returns a generator of BlobContainer_Status instances for property testing.
// We first initialize blobContainer_statusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BlobContainer_StatusGenerator() gopter.Gen {
	if blobContainer_statusGenerator != nil {
		return blobContainer_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobContainer_Status(generators)
	blobContainer_statusGenerator = gen.Struct(reflect.TypeOf(BlobContainer_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBlobContainer_Status(generators)
	AddRelatedPropertyGeneratorsForBlobContainer_Status(generators)
	blobContainer_statusGenerator = gen.Struct(reflect.TypeOf(BlobContainer_Status{}), generators)

	return blobContainer_statusGenerator
}

// AddIndependentPropertyGeneratorsForBlobContainer_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBlobContainer_Status(gens map[string]gopter.Gen) {
	gens["DefaultEncryptionScope"] = gen.PtrOf(gen.AlphaString())
	gens["Deleted"] = gen.PtrOf(gen.Bool())
	gens["DeletedTime"] = gen.PtrOf(gen.AlphaString())
	gens["DenyEncryptionScopeOverride"] = gen.PtrOf(gen.Bool())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["HasImmutabilityPolicy"] = gen.PtrOf(gen.Bool())
	gens["HasLegalHold"] = gen.PtrOf(gen.Bool())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedTime"] = gen.PtrOf(gen.AlphaString())
	gens["LeaseDuration"] = gen.PtrOf(gen.AlphaString())
	gens["LeaseState"] = gen.PtrOf(gen.AlphaString())
	gens["LeaseStatus"] = gen.PtrOf(gen.AlphaString())
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PublicAccess"] = gen.PtrOf(gen.AlphaString())
	gens["RemainingRetentionDays"] = gen.PtrOf(gen.Int())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Version"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForBlobContainer_Status is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBlobContainer_Status(gens map[string]gopter.Gen) {
	gens["ImmutabilityPolicy"] = gen.PtrOf(ImmutabilityPolicyProperties_StatusGenerator())
	gens["ImmutableStorageWithVersioning"] = gen.PtrOf(ImmutableStorageWithVersioning_StatusGenerator())
	gens["LegalHold"] = gen.PtrOf(LegalHoldProperties_StatusGenerator())
}

func Test_StorageAccountsBlobServicesContainers_SPEC_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsBlobServicesContainers_SPEC via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsBlobServicesContainers_SPEC, StorageAccountsBlobServicesContainers_SPECGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsBlobServicesContainers_SPEC runs a test to see if a specific instance of StorageAccountsBlobServicesContainers_SPEC round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsBlobServicesContainers_SPEC(subject StorageAccountsBlobServicesContainers_SPEC) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsBlobServicesContainers_SPEC
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

// Generator of StorageAccountsBlobServicesContainers_SPEC instances for property testing - lazily instantiated by
//StorageAccountsBlobServicesContainers_SPECGenerator()
var storageAccountsBlobServicesContainers_specGenerator gopter.Gen

// StorageAccountsBlobServicesContainers_SPECGenerator returns a generator of StorageAccountsBlobServicesContainers_SPEC instances for property testing.
// We first initialize storageAccountsBlobServicesContainers_specGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccountsBlobServicesContainers_SPECGenerator() gopter.Gen {
	if storageAccountsBlobServicesContainers_specGenerator != nil {
		return storageAccountsBlobServicesContainers_specGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsBlobServicesContainers_SPEC(generators)
	storageAccountsBlobServicesContainers_specGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobServicesContainers_SPEC{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsBlobServicesContainers_SPEC(generators)
	AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesContainers_SPEC(generators)
	storageAccountsBlobServicesContainers_specGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobServicesContainers_SPEC{}), generators)

	return storageAccountsBlobServicesContainers_specGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsBlobServicesContainers_SPEC is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsBlobServicesContainers_SPEC(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["DefaultEncryptionScope"] = gen.PtrOf(gen.AlphaString())
	gens["DenyEncryptionScopeOverride"] = gen.PtrOf(gen.Bool())
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["PublicAccess"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesContainers_SPEC is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesContainers_SPEC(gens map[string]gopter.Gen) {
	gens["ImmutableStorageWithVersioning"] = gen.PtrOf(ImmutableStorageWithVersioning_SpecGenerator())
}

func Test_ImmutabilityPolicyProperties_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImmutabilityPolicyProperties_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImmutabilityPolicyProperties_Status, ImmutabilityPolicyProperties_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImmutabilityPolicyProperties_Status runs a test to see if a specific instance of ImmutabilityPolicyProperties_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForImmutabilityPolicyProperties_Status(subject ImmutabilityPolicyProperties_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImmutabilityPolicyProperties_Status
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

// Generator of ImmutabilityPolicyProperties_Status instances for property testing - lazily instantiated by
//ImmutabilityPolicyProperties_StatusGenerator()
var immutabilityPolicyProperties_statusGenerator gopter.Gen

// ImmutabilityPolicyProperties_StatusGenerator returns a generator of ImmutabilityPolicyProperties_Status instances for property testing.
// We first initialize immutabilityPolicyProperties_statusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImmutabilityPolicyProperties_StatusGenerator() gopter.Gen {
	if immutabilityPolicyProperties_statusGenerator != nil {
		return immutabilityPolicyProperties_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutabilityPolicyProperties_Status(generators)
	immutabilityPolicyProperties_statusGenerator = gen.Struct(reflect.TypeOf(ImmutabilityPolicyProperties_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutabilityPolicyProperties_Status(generators)
	AddRelatedPropertyGeneratorsForImmutabilityPolicyProperties_Status(generators)
	immutabilityPolicyProperties_statusGenerator = gen.Struct(reflect.TypeOf(ImmutabilityPolicyProperties_Status{}), generators)

	return immutabilityPolicyProperties_statusGenerator
}

// AddIndependentPropertyGeneratorsForImmutabilityPolicyProperties_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImmutabilityPolicyProperties_Status(gens map[string]gopter.Gen) {
	gens["AllowProtectedAppendWrites"] = gen.PtrOf(gen.Bool())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["ImmutabilityPeriodSinceCreationInDays"] = gen.PtrOf(gen.Int())
	gens["State"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForImmutabilityPolicyProperties_Status is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImmutabilityPolicyProperties_Status(gens map[string]gopter.Gen) {
	gens["UpdateHistory"] = gen.SliceOf(UpdateHistoryProperty_StatusGenerator())
}

func Test_ImmutableStorageWithVersioning_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImmutableStorageWithVersioning_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImmutableStorageWithVersioning_Spec, ImmutableStorageWithVersioning_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImmutableStorageWithVersioning_Spec runs a test to see if a specific instance of ImmutableStorageWithVersioning_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForImmutableStorageWithVersioning_Spec(subject ImmutableStorageWithVersioning_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImmutableStorageWithVersioning_Spec
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

// Generator of ImmutableStorageWithVersioning_Spec instances for property testing - lazily instantiated by
//ImmutableStorageWithVersioning_SpecGenerator()
var immutableStorageWithVersioning_specGenerator gopter.Gen

// ImmutableStorageWithVersioning_SpecGenerator returns a generator of ImmutableStorageWithVersioning_Spec instances for property testing.
func ImmutableStorageWithVersioning_SpecGenerator() gopter.Gen {
	if immutableStorageWithVersioning_specGenerator != nil {
		return immutableStorageWithVersioning_specGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutableStorageWithVersioning_Spec(generators)
	immutableStorageWithVersioning_specGenerator = gen.Struct(reflect.TypeOf(ImmutableStorageWithVersioning_Spec{}), generators)

	return immutableStorageWithVersioning_specGenerator
}

// AddIndependentPropertyGeneratorsForImmutableStorageWithVersioning_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImmutableStorageWithVersioning_Spec(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

func Test_ImmutableStorageWithVersioning_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImmutableStorageWithVersioning_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImmutableStorageWithVersioning_Status, ImmutableStorageWithVersioning_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImmutableStorageWithVersioning_Status runs a test to see if a specific instance of ImmutableStorageWithVersioning_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForImmutableStorageWithVersioning_Status(subject ImmutableStorageWithVersioning_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImmutableStorageWithVersioning_Status
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

// Generator of ImmutableStorageWithVersioning_Status instances for property testing - lazily instantiated by
//ImmutableStorageWithVersioning_StatusGenerator()
var immutableStorageWithVersioning_statusGenerator gopter.Gen

// ImmutableStorageWithVersioning_StatusGenerator returns a generator of ImmutableStorageWithVersioning_Status instances for property testing.
func ImmutableStorageWithVersioning_StatusGenerator() gopter.Gen {
	if immutableStorageWithVersioning_statusGenerator != nil {
		return immutableStorageWithVersioning_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutableStorageWithVersioning_Status(generators)
	immutableStorageWithVersioning_statusGenerator = gen.Struct(reflect.TypeOf(ImmutableStorageWithVersioning_Status{}), generators)

	return immutableStorageWithVersioning_statusGenerator
}

// AddIndependentPropertyGeneratorsForImmutableStorageWithVersioning_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImmutableStorageWithVersioning_Status(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["MigrationState"] = gen.PtrOf(gen.AlphaString())
	gens["TimeStamp"] = gen.PtrOf(gen.AlphaString())
}

func Test_LegalHoldProperties_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LegalHoldProperties_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLegalHoldProperties_Status, LegalHoldProperties_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLegalHoldProperties_Status runs a test to see if a specific instance of LegalHoldProperties_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForLegalHoldProperties_Status(subject LegalHoldProperties_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LegalHoldProperties_Status
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

// Generator of LegalHoldProperties_Status instances for property testing - lazily instantiated by
//LegalHoldProperties_StatusGenerator()
var legalHoldProperties_statusGenerator gopter.Gen

// LegalHoldProperties_StatusGenerator returns a generator of LegalHoldProperties_Status instances for property testing.
// We first initialize legalHoldProperties_statusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func LegalHoldProperties_StatusGenerator() gopter.Gen {
	if legalHoldProperties_statusGenerator != nil {
		return legalHoldProperties_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLegalHoldProperties_Status(generators)
	legalHoldProperties_statusGenerator = gen.Struct(reflect.TypeOf(LegalHoldProperties_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLegalHoldProperties_Status(generators)
	AddRelatedPropertyGeneratorsForLegalHoldProperties_Status(generators)
	legalHoldProperties_statusGenerator = gen.Struct(reflect.TypeOf(LegalHoldProperties_Status{}), generators)

	return legalHoldProperties_statusGenerator
}

// AddIndependentPropertyGeneratorsForLegalHoldProperties_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLegalHoldProperties_Status(gens map[string]gopter.Gen) {
	gens["HasLegalHold"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForLegalHoldProperties_Status is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForLegalHoldProperties_Status(gens map[string]gopter.Gen) {
	gens["Tags"] = gen.SliceOf(TagProperty_StatusGenerator())
}

func Test_TagProperty_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TagProperty_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTagProperty_Status, TagProperty_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTagProperty_Status runs a test to see if a specific instance of TagProperty_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForTagProperty_Status(subject TagProperty_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TagProperty_Status
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

// Generator of TagProperty_Status instances for property testing - lazily instantiated by TagProperty_StatusGenerator()
var tagProperty_statusGenerator gopter.Gen

// TagProperty_StatusGenerator returns a generator of TagProperty_Status instances for property testing.
func TagProperty_StatusGenerator() gopter.Gen {
	if tagProperty_statusGenerator != nil {
		return tagProperty_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTagProperty_Status(generators)
	tagProperty_statusGenerator = gen.Struct(reflect.TypeOf(TagProperty_Status{}), generators)

	return tagProperty_statusGenerator
}

// AddIndependentPropertyGeneratorsForTagProperty_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTagProperty_Status(gens map[string]gopter.Gen) {
	gens["ObjectIdentifier"] = gen.PtrOf(gen.AlphaString())
	gens["Tag"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Timestamp"] = gen.PtrOf(gen.AlphaString())
	gens["Upn"] = gen.PtrOf(gen.AlphaString())
}

func Test_UpdateHistoryProperty_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UpdateHistoryProperty_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUpdateHistoryProperty_Status, UpdateHistoryProperty_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUpdateHistoryProperty_Status runs a test to see if a specific instance of UpdateHistoryProperty_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForUpdateHistoryProperty_Status(subject UpdateHistoryProperty_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UpdateHistoryProperty_Status
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

// Generator of UpdateHistoryProperty_Status instances for property testing - lazily instantiated by
//UpdateHistoryProperty_StatusGenerator()
var updateHistoryProperty_statusGenerator gopter.Gen

// UpdateHistoryProperty_StatusGenerator returns a generator of UpdateHistoryProperty_Status instances for property testing.
func UpdateHistoryProperty_StatusGenerator() gopter.Gen {
	if updateHistoryProperty_statusGenerator != nil {
		return updateHistoryProperty_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUpdateHistoryProperty_Status(generators)
	updateHistoryProperty_statusGenerator = gen.Struct(reflect.TypeOf(UpdateHistoryProperty_Status{}), generators)

	return updateHistoryProperty_statusGenerator
}

// AddIndependentPropertyGeneratorsForUpdateHistoryProperty_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUpdateHistoryProperty_Status(gens map[string]gopter.Gen) {
	gens["ImmutabilityPeriodSinceCreationInDays"] = gen.PtrOf(gen.Int())
	gens["ObjectIdentifier"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Timestamp"] = gen.PtrOf(gen.AlphaString())
	gens["Update"] = gen.PtrOf(gen.AlphaString())
	gens["Upn"] = gen.PtrOf(gen.AlphaString())
}

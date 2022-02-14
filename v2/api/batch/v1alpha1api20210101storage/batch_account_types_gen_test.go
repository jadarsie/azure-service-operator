// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210101storage

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

func Test_BatchAccount_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BatchAccount via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBatchAccount, BatchAccountGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBatchAccount runs a test to see if a specific instance of BatchAccount round trips to JSON and back losslessly
func RunJSONSerializationTestForBatchAccount(subject BatchAccount) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BatchAccount
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

// Generator of BatchAccount instances for property testing - lazily instantiated by BatchAccountGenerator()
var batchAccountGenerator gopter.Gen

// BatchAccountGenerator returns a generator of BatchAccount instances for property testing.
func BatchAccountGenerator() gopter.Gen {
	if batchAccountGenerator != nil {
		return batchAccountGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForBatchAccount(generators)
	batchAccountGenerator = gen.Struct(reflect.TypeOf(BatchAccount{}), generators)

	return batchAccountGenerator
}

// AddRelatedPropertyGeneratorsForBatchAccount is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBatchAccount(gens map[string]gopter.Gen) {
	gens["Spec"] = BatchAccount_SpecGenerator()
	gens["Status"] = BatchAccount_StatusGenerator()
}

func Test_BatchAccount_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BatchAccount_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBatchAccount_Spec, BatchAccount_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBatchAccount_Spec runs a test to see if a specific instance of BatchAccount_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForBatchAccount_Spec(subject BatchAccount_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BatchAccount_Spec
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

// Generator of BatchAccount_Spec instances for property testing - lazily instantiated by BatchAccount_SpecGenerator()
var batchAccount_specGenerator gopter.Gen

// BatchAccount_SpecGenerator returns a generator of BatchAccount_Spec instances for property testing.
// We first initialize batchAccount_specGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BatchAccount_SpecGenerator() gopter.Gen {
	if batchAccount_specGenerator != nil {
		return batchAccount_specGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBatchAccount_Spec(generators)
	batchAccount_specGenerator = gen.Struct(reflect.TypeOf(BatchAccount_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBatchAccount_Spec(generators)
	AddRelatedPropertyGeneratorsForBatchAccount_Spec(generators)
	batchAccount_specGenerator = gen.Struct(reflect.TypeOf(BatchAccount_Spec{}), generators)

	return batchAccount_specGenerator
}

// AddIndependentPropertyGeneratorsForBatchAccount_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBatchAccount_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["PoolAllocationMode"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForBatchAccount_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBatchAccount_Spec(gens map[string]gopter.Gen) {
	gens["AutoStorage"] = gen.PtrOf(AutoStorageBasePropertiesGenerator())
	gens["Encryption"] = gen.PtrOf(EncryptionPropertiesGenerator())
	gens["Identity"] = gen.PtrOf(BatchAccountIdentityGenerator())
	gens["KeyVaultReference"] = gen.PtrOf(KeyVaultReferenceGenerator())
}

func Test_BatchAccount_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BatchAccount_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBatchAccount_Status, BatchAccount_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBatchAccount_Status runs a test to see if a specific instance of BatchAccount_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForBatchAccount_Status(subject BatchAccount_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BatchAccount_Status
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

// Generator of BatchAccount_Status instances for property testing - lazily instantiated by
//BatchAccount_StatusGenerator()
var batchAccount_statusGenerator gopter.Gen

// BatchAccount_StatusGenerator returns a generator of BatchAccount_Status instances for property testing.
// We first initialize batchAccount_statusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BatchAccount_StatusGenerator() gopter.Gen {
	if batchAccount_statusGenerator != nil {
		return batchAccount_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBatchAccount_Status(generators)
	batchAccount_statusGenerator = gen.Struct(reflect.TypeOf(BatchAccount_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBatchAccount_Status(generators)
	AddRelatedPropertyGeneratorsForBatchAccount_Status(generators)
	batchAccount_statusGenerator = gen.Struct(reflect.TypeOf(BatchAccount_Status{}), generators)

	return batchAccount_statusGenerator
}

// AddIndependentPropertyGeneratorsForBatchAccount_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBatchAccount_Status(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["PoolAllocationMode"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForBatchAccount_Status is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBatchAccount_Status(gens map[string]gopter.Gen) {
	gens["AutoStorage"] = gen.PtrOf(AutoStorageBaseProperties_StatusGenerator())
	gens["Encryption"] = gen.PtrOf(EncryptionProperties_StatusGenerator())
	gens["Identity"] = gen.PtrOf(BatchAccountIdentity_StatusGenerator())
	gens["KeyVaultReference"] = gen.PtrOf(KeyVaultReference_StatusGenerator())
}

func Test_AutoStorageBaseProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoStorageBaseProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoStorageBaseProperties, AutoStorageBasePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoStorageBaseProperties runs a test to see if a specific instance of AutoStorageBaseProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoStorageBaseProperties(subject AutoStorageBaseProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoStorageBaseProperties
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

// Generator of AutoStorageBaseProperties instances for property testing - lazily instantiated by
//AutoStorageBasePropertiesGenerator()
var autoStorageBasePropertiesGenerator gopter.Gen

// AutoStorageBasePropertiesGenerator returns a generator of AutoStorageBaseProperties instances for property testing.
func AutoStorageBasePropertiesGenerator() gopter.Gen {
	if autoStorageBasePropertiesGenerator != nil {
		return autoStorageBasePropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	autoStorageBasePropertiesGenerator = gen.Struct(reflect.TypeOf(AutoStorageBaseProperties{}), generators)

	return autoStorageBasePropertiesGenerator
}

func Test_AutoStorageBaseProperties_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoStorageBaseProperties_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoStorageBaseProperties_Status, AutoStorageBaseProperties_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoStorageBaseProperties_Status runs a test to see if a specific instance of AutoStorageBaseProperties_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoStorageBaseProperties_Status(subject AutoStorageBaseProperties_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoStorageBaseProperties_Status
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

// Generator of AutoStorageBaseProperties_Status instances for property testing - lazily instantiated by
//AutoStorageBaseProperties_StatusGenerator()
var autoStorageBaseProperties_statusGenerator gopter.Gen

// AutoStorageBaseProperties_StatusGenerator returns a generator of AutoStorageBaseProperties_Status instances for property testing.
func AutoStorageBaseProperties_StatusGenerator() gopter.Gen {
	if autoStorageBaseProperties_statusGenerator != nil {
		return autoStorageBaseProperties_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoStorageBaseProperties_Status(generators)
	autoStorageBaseProperties_statusGenerator = gen.Struct(reflect.TypeOf(AutoStorageBaseProperties_Status{}), generators)

	return autoStorageBaseProperties_statusGenerator
}

// AddIndependentPropertyGeneratorsForAutoStorageBaseProperties_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAutoStorageBaseProperties_Status(gens map[string]gopter.Gen) {
	gens["StorageAccountId"] = gen.PtrOf(gen.AlphaString())
}

func Test_BatchAccountIdentity_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BatchAccountIdentity via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBatchAccountIdentity, BatchAccountIdentityGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBatchAccountIdentity runs a test to see if a specific instance of BatchAccountIdentity round trips to JSON and back losslessly
func RunJSONSerializationTestForBatchAccountIdentity(subject BatchAccountIdentity) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BatchAccountIdentity
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

// Generator of BatchAccountIdentity instances for property testing - lazily instantiated by
//BatchAccountIdentityGenerator()
var batchAccountIdentityGenerator gopter.Gen

// BatchAccountIdentityGenerator returns a generator of BatchAccountIdentity instances for property testing.
func BatchAccountIdentityGenerator() gopter.Gen {
	if batchAccountIdentityGenerator != nil {
		return batchAccountIdentityGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBatchAccountIdentity(generators)
	batchAccountIdentityGenerator = gen.Struct(reflect.TypeOf(BatchAccountIdentity{}), generators)

	return batchAccountIdentityGenerator
}

// AddIndependentPropertyGeneratorsForBatchAccountIdentity is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBatchAccountIdentity(gens map[string]gopter.Gen) {
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_BatchAccountIdentity_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BatchAccountIdentity_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBatchAccountIdentity_Status, BatchAccountIdentity_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBatchAccountIdentity_Status runs a test to see if a specific instance of BatchAccountIdentity_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForBatchAccountIdentity_Status(subject BatchAccountIdentity_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BatchAccountIdentity_Status
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

// Generator of BatchAccountIdentity_Status instances for property testing - lazily instantiated by
//BatchAccountIdentity_StatusGenerator()
var batchAccountIdentity_statusGenerator gopter.Gen

// BatchAccountIdentity_StatusGenerator returns a generator of BatchAccountIdentity_Status instances for property testing.
// We first initialize batchAccountIdentity_statusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BatchAccountIdentity_StatusGenerator() gopter.Gen {
	if batchAccountIdentity_statusGenerator != nil {
		return batchAccountIdentity_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBatchAccountIdentity_Status(generators)
	batchAccountIdentity_statusGenerator = gen.Struct(reflect.TypeOf(BatchAccountIdentity_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBatchAccountIdentity_Status(generators)
	AddRelatedPropertyGeneratorsForBatchAccountIdentity_Status(generators)
	batchAccountIdentity_statusGenerator = gen.Struct(reflect.TypeOf(BatchAccountIdentity_Status{}), generators)

	return batchAccountIdentity_statusGenerator
}

// AddIndependentPropertyGeneratorsForBatchAccountIdentity_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBatchAccountIdentity_Status(gens map[string]gopter.Gen) {
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForBatchAccountIdentity_Status is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBatchAccountIdentity_Status(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentities"] = gen.MapOf(gen.AlphaString(), BatchAccountIdentity_StatusUserAssignedIdentitiesGenerator())
}

func Test_EncryptionProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EncryptionProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionProperties, EncryptionPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionProperties runs a test to see if a specific instance of EncryptionProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionProperties(subject EncryptionProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EncryptionProperties
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

// Generator of EncryptionProperties instances for property testing - lazily instantiated by
//EncryptionPropertiesGenerator()
var encryptionPropertiesGenerator gopter.Gen

// EncryptionPropertiesGenerator returns a generator of EncryptionProperties instances for property testing.
// We first initialize encryptionPropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EncryptionPropertiesGenerator() gopter.Gen {
	if encryptionPropertiesGenerator != nil {
		return encryptionPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionProperties(generators)
	encryptionPropertiesGenerator = gen.Struct(reflect.TypeOf(EncryptionProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionProperties(generators)
	AddRelatedPropertyGeneratorsForEncryptionProperties(generators)
	encryptionPropertiesGenerator = gen.Struct(reflect.TypeOf(EncryptionProperties{}), generators)

	return encryptionPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForEncryptionProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryptionProperties(gens map[string]gopter.Gen) {
	gens["KeySource"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForEncryptionProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEncryptionProperties(gens map[string]gopter.Gen) {
	gens["KeyVaultProperties"] = gen.PtrOf(KeyVaultPropertiesGenerator())
}

func Test_EncryptionProperties_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EncryptionProperties_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionProperties_Status, EncryptionProperties_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionProperties_Status runs a test to see if a specific instance of EncryptionProperties_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionProperties_Status(subject EncryptionProperties_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EncryptionProperties_Status
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

// Generator of EncryptionProperties_Status instances for property testing - lazily instantiated by
//EncryptionProperties_StatusGenerator()
var encryptionProperties_statusGenerator gopter.Gen

// EncryptionProperties_StatusGenerator returns a generator of EncryptionProperties_Status instances for property testing.
// We first initialize encryptionProperties_statusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EncryptionProperties_StatusGenerator() gopter.Gen {
	if encryptionProperties_statusGenerator != nil {
		return encryptionProperties_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionProperties_Status(generators)
	encryptionProperties_statusGenerator = gen.Struct(reflect.TypeOf(EncryptionProperties_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionProperties_Status(generators)
	AddRelatedPropertyGeneratorsForEncryptionProperties_Status(generators)
	encryptionProperties_statusGenerator = gen.Struct(reflect.TypeOf(EncryptionProperties_Status{}), generators)

	return encryptionProperties_statusGenerator
}

// AddIndependentPropertyGeneratorsForEncryptionProperties_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryptionProperties_Status(gens map[string]gopter.Gen) {
	gens["KeySource"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForEncryptionProperties_Status is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEncryptionProperties_Status(gens map[string]gopter.Gen) {
	gens["KeyVaultProperties"] = gen.PtrOf(KeyVaultProperties_StatusGenerator())
}

func Test_KeyVaultReference_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultReference via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultReference, KeyVaultReferenceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultReference runs a test to see if a specific instance of KeyVaultReference round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultReference(subject KeyVaultReference) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultReference
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

// Generator of KeyVaultReference instances for property testing - lazily instantiated by KeyVaultReferenceGenerator()
var keyVaultReferenceGenerator gopter.Gen

// KeyVaultReferenceGenerator returns a generator of KeyVaultReference instances for property testing.
func KeyVaultReferenceGenerator() gopter.Gen {
	if keyVaultReferenceGenerator != nil {
		return keyVaultReferenceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultReference(generators)
	keyVaultReferenceGenerator = gen.Struct(reflect.TypeOf(KeyVaultReference{}), generators)

	return keyVaultReferenceGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultReference is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultReference(gens map[string]gopter.Gen) {
	gens["Url"] = gen.PtrOf(gen.AlphaString())
}

func Test_KeyVaultReference_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultReference_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultReference_Status, KeyVaultReference_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultReference_Status runs a test to see if a specific instance of KeyVaultReference_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultReference_Status(subject KeyVaultReference_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultReference_Status
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

// Generator of KeyVaultReference_Status instances for property testing - lazily instantiated by
//KeyVaultReference_StatusGenerator()
var keyVaultReference_statusGenerator gopter.Gen

// KeyVaultReference_StatusGenerator returns a generator of KeyVaultReference_Status instances for property testing.
func KeyVaultReference_StatusGenerator() gopter.Gen {
	if keyVaultReference_statusGenerator != nil {
		return keyVaultReference_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultReference_Status(generators)
	keyVaultReference_statusGenerator = gen.Struct(reflect.TypeOf(KeyVaultReference_Status{}), generators)

	return keyVaultReference_statusGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultReference_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultReference_Status(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Url"] = gen.PtrOf(gen.AlphaString())
}

func Test_BatchAccountIdentity_StatusUserAssignedIdentities_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BatchAccountIdentity_StatusUserAssignedIdentities via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBatchAccountIdentity_StatusUserAssignedIdentities, BatchAccountIdentity_StatusUserAssignedIdentitiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBatchAccountIdentity_StatusUserAssignedIdentities runs a test to see if a specific instance of BatchAccountIdentity_StatusUserAssignedIdentities round trips to JSON and back losslessly
func RunJSONSerializationTestForBatchAccountIdentity_StatusUserAssignedIdentities(subject BatchAccountIdentity_StatusUserAssignedIdentities) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BatchAccountIdentity_StatusUserAssignedIdentities
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

// Generator of BatchAccountIdentity_StatusUserAssignedIdentities instances for property testing - lazily instantiated
//by BatchAccountIdentity_StatusUserAssignedIdentitiesGenerator()
var batchAccountIdentity_statusUserAssignedIdentitiesGenerator gopter.Gen

// BatchAccountIdentity_StatusUserAssignedIdentitiesGenerator returns a generator of BatchAccountIdentity_StatusUserAssignedIdentities instances for property testing.
func BatchAccountIdentity_StatusUserAssignedIdentitiesGenerator() gopter.Gen {
	if batchAccountIdentity_statusUserAssignedIdentitiesGenerator != nil {
		return batchAccountIdentity_statusUserAssignedIdentitiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBatchAccountIdentity_StatusUserAssignedIdentities(generators)
	batchAccountIdentity_statusUserAssignedIdentitiesGenerator = gen.Struct(reflect.TypeOf(BatchAccountIdentity_StatusUserAssignedIdentities{}), generators)

	return batchAccountIdentity_statusUserAssignedIdentitiesGenerator
}

// AddIndependentPropertyGeneratorsForBatchAccountIdentity_StatusUserAssignedIdentities is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBatchAccountIdentity_StatusUserAssignedIdentities(gens map[string]gopter.Gen) {
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
}

func Test_KeyVaultProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultProperties, KeyVaultPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultProperties runs a test to see if a specific instance of KeyVaultProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultProperties(subject KeyVaultProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultProperties
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

// Generator of KeyVaultProperties instances for property testing - lazily instantiated by KeyVaultPropertiesGenerator()
var keyVaultPropertiesGenerator gopter.Gen

// KeyVaultPropertiesGenerator returns a generator of KeyVaultProperties instances for property testing.
func KeyVaultPropertiesGenerator() gopter.Gen {
	if keyVaultPropertiesGenerator != nil {
		return keyVaultPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultProperties(generators)
	keyVaultPropertiesGenerator = gen.Struct(reflect.TypeOf(KeyVaultProperties{}), generators)

	return keyVaultPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultProperties(gens map[string]gopter.Gen) {
	gens["KeyIdentifier"] = gen.PtrOf(gen.AlphaString())
}

func Test_KeyVaultProperties_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultProperties_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultProperties_Status, KeyVaultProperties_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultProperties_Status runs a test to see if a specific instance of KeyVaultProperties_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultProperties_Status(subject KeyVaultProperties_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultProperties_Status
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

// Generator of KeyVaultProperties_Status instances for property testing - lazily instantiated by
//KeyVaultProperties_StatusGenerator()
var keyVaultProperties_statusGenerator gopter.Gen

// KeyVaultProperties_StatusGenerator returns a generator of KeyVaultProperties_Status instances for property testing.
func KeyVaultProperties_StatusGenerator() gopter.Gen {
	if keyVaultProperties_statusGenerator != nil {
		return keyVaultProperties_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultProperties_Status(generators)
	keyVaultProperties_statusGenerator = gen.Struct(reflect.TypeOf(KeyVaultProperties_Status{}), generators)

	return keyVaultProperties_statusGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultProperties_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultProperties_Status(gens map[string]gopter.Gen) {
	gens["KeyIdentifier"] = gen.PtrOf(gen.AlphaString())
}

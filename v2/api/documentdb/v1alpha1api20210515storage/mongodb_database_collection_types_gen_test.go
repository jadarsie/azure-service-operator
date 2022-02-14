// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515storage

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

func Test_MongodbDatabaseCollection_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongodbDatabaseCollection via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongodbDatabaseCollection, MongodbDatabaseCollectionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongodbDatabaseCollection runs a test to see if a specific instance of MongodbDatabaseCollection round trips to JSON and back losslessly
func RunJSONSerializationTestForMongodbDatabaseCollection(subject MongodbDatabaseCollection) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongodbDatabaseCollection
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

// Generator of MongodbDatabaseCollection instances for property testing - lazily instantiated by
//MongodbDatabaseCollectionGenerator()
var mongodbDatabaseCollectionGenerator gopter.Gen

// MongodbDatabaseCollectionGenerator returns a generator of MongodbDatabaseCollection instances for property testing.
func MongodbDatabaseCollectionGenerator() gopter.Gen {
	if mongodbDatabaseCollectionGenerator != nil {
		return mongodbDatabaseCollectionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongodbDatabaseCollection(generators)
	mongodbDatabaseCollectionGenerator = gen.Struct(reflect.TypeOf(MongodbDatabaseCollection{}), generators)

	return mongodbDatabaseCollectionGenerator
}

// AddRelatedPropertyGeneratorsForMongodbDatabaseCollection is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongodbDatabaseCollection(gens map[string]gopter.Gen) {
	gens["Spec"] = DatabaseAccountsMongodbDatabasesCollection_SpecGenerator()
	gens["Status"] = MongoDBCollection_StatusGenerator()
}

func Test_DatabaseAccountsMongodbDatabasesCollection_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsMongodbDatabasesCollection_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollection_Spec, DatabaseAccountsMongodbDatabasesCollection_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollection_Spec runs a test to see if a specific instance of DatabaseAccountsMongodbDatabasesCollection_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollection_Spec(subject DatabaseAccountsMongodbDatabasesCollection_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsMongodbDatabasesCollection_Spec
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

// Generator of DatabaseAccountsMongodbDatabasesCollection_Spec instances for property testing - lazily instantiated by
//DatabaseAccountsMongodbDatabasesCollection_SpecGenerator()
var databaseAccountsMongodbDatabasesCollection_specGenerator gopter.Gen

// DatabaseAccountsMongodbDatabasesCollection_SpecGenerator returns a generator of DatabaseAccountsMongodbDatabasesCollection_Spec instances for property testing.
// We first initialize databaseAccountsMongodbDatabasesCollection_specGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsMongodbDatabasesCollection_SpecGenerator() gopter.Gen {
	if databaseAccountsMongodbDatabasesCollection_specGenerator != nil {
		return databaseAccountsMongodbDatabasesCollection_specGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollection_Spec(generators)
	databaseAccountsMongodbDatabasesCollection_specGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabasesCollection_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollection_Spec(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollection_Spec(generators)
	databaseAccountsMongodbDatabasesCollection_specGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabasesCollection_Spec{}), generators)

	return databaseAccountsMongodbDatabasesCollection_specGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollection_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollection_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollection_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollection_Spec(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsGenerator())
	gens["Resource"] = gen.PtrOf(MongoDBCollectionResourceGenerator())
}

func Test_MongoDBCollection_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBCollection_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBCollection_Status, MongoDBCollection_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBCollection_Status runs a test to see if a specific instance of MongoDBCollection_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBCollection_Status(subject MongoDBCollection_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBCollection_Status
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

// Generator of MongoDBCollection_Status instances for property testing - lazily instantiated by
//MongoDBCollection_StatusGenerator()
var mongoDBCollection_statusGenerator gopter.Gen

// MongoDBCollection_StatusGenerator returns a generator of MongoDBCollection_Status instances for property testing.
// We first initialize mongoDBCollection_statusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MongoDBCollection_StatusGenerator() gopter.Gen {
	if mongoDBCollection_statusGenerator != nil {
		return mongoDBCollection_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollection_Status(generators)
	mongoDBCollection_statusGenerator = gen.Struct(reflect.TypeOf(MongoDBCollection_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollection_Status(generators)
	AddRelatedPropertyGeneratorsForMongoDBCollection_Status(generators)
	mongoDBCollection_statusGenerator = gen.Struct(reflect.TypeOf(MongoDBCollection_Status{}), generators)

	return mongoDBCollection_statusGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBCollection_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBCollection_Status(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMongoDBCollection_Status is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBCollection_Status(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptions_StatusGenerator())
	gens["Resource"] = gen.PtrOf(MongoDBCollectionResource_StatusGenerator())
}

func Test_MongoDBCollectionResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBCollectionResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBCollectionResource, MongoDBCollectionResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBCollectionResource runs a test to see if a specific instance of MongoDBCollectionResource round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBCollectionResource(subject MongoDBCollectionResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBCollectionResource
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

// Generator of MongoDBCollectionResource instances for property testing - lazily instantiated by
//MongoDBCollectionResourceGenerator()
var mongoDBCollectionResourceGenerator gopter.Gen

// MongoDBCollectionResourceGenerator returns a generator of MongoDBCollectionResource instances for property testing.
// We first initialize mongoDBCollectionResourceGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MongoDBCollectionResourceGenerator() gopter.Gen {
	if mongoDBCollectionResourceGenerator != nil {
		return mongoDBCollectionResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionResource(generators)
	mongoDBCollectionResourceGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionResource{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionResource(generators)
	AddRelatedPropertyGeneratorsForMongoDBCollectionResource(generators)
	mongoDBCollectionResourceGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionResource{}), generators)

	return mongoDBCollectionResourceGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBCollectionResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBCollectionResource(gens map[string]gopter.Gen) {
	gens["AnalyticalStorageTtl"] = gen.PtrOf(gen.Int())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["ShardKey"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMongoDBCollectionResource is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBCollectionResource(gens map[string]gopter.Gen) {
	gens["Indexes"] = gen.SliceOf(MongoIndexGenerator())
}

func Test_MongoDBCollectionResource_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBCollectionResource_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBCollectionResource_Status, MongoDBCollectionResource_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBCollectionResource_Status runs a test to see if a specific instance of MongoDBCollectionResource_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBCollectionResource_Status(subject MongoDBCollectionResource_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBCollectionResource_Status
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

// Generator of MongoDBCollectionResource_Status instances for property testing - lazily instantiated by
//MongoDBCollectionResource_StatusGenerator()
var mongoDBCollectionResource_statusGenerator gopter.Gen

// MongoDBCollectionResource_StatusGenerator returns a generator of MongoDBCollectionResource_Status instances for property testing.
// We first initialize mongoDBCollectionResource_statusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MongoDBCollectionResource_StatusGenerator() gopter.Gen {
	if mongoDBCollectionResource_statusGenerator != nil {
		return mongoDBCollectionResource_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionResource_Status(generators)
	mongoDBCollectionResource_statusGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionResource_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionResource_Status(generators)
	AddRelatedPropertyGeneratorsForMongoDBCollectionResource_Status(generators)
	mongoDBCollectionResource_statusGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionResource_Status{}), generators)

	return mongoDBCollectionResource_statusGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBCollectionResource_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBCollectionResource_Status(gens map[string]gopter.Gen) {
	gens["AnalyticalStorageTtl"] = gen.PtrOf(gen.Int())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["ShardKey"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMongoDBCollectionResource_Status is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBCollectionResource_Status(gens map[string]gopter.Gen) {
	gens["Indexes"] = gen.SliceOf(MongoIndex_StatusGenerator())
}

func Test_MongoIndex_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndex via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndex, MongoIndexGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndex runs a test to see if a specific instance of MongoIndex round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndex(subject MongoIndex) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndex
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

// Generator of MongoIndex instances for property testing - lazily instantiated by MongoIndexGenerator()
var mongoIndexGenerator gopter.Gen

// MongoIndexGenerator returns a generator of MongoIndex instances for property testing.
func MongoIndexGenerator() gopter.Gen {
	if mongoIndexGenerator != nil {
		return mongoIndexGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongoIndex(generators)
	mongoIndexGenerator = gen.Struct(reflect.TypeOf(MongoIndex{}), generators)

	return mongoIndexGenerator
}

// AddRelatedPropertyGeneratorsForMongoIndex is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoIndex(gens map[string]gopter.Gen) {
	gens["Key"] = gen.PtrOf(MongoIndexKeysGenerator())
	gens["Options"] = gen.PtrOf(MongoIndexOptionsGenerator())
}

func Test_MongoIndex_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndex_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndex_Status, MongoIndex_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndex_Status runs a test to see if a specific instance of MongoIndex_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndex_Status(subject MongoIndex_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndex_Status
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

// Generator of MongoIndex_Status instances for property testing - lazily instantiated by MongoIndex_StatusGenerator()
var mongoIndex_statusGenerator gopter.Gen

// MongoIndex_StatusGenerator returns a generator of MongoIndex_Status instances for property testing.
func MongoIndex_StatusGenerator() gopter.Gen {
	if mongoIndex_statusGenerator != nil {
		return mongoIndex_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongoIndex_Status(generators)
	mongoIndex_statusGenerator = gen.Struct(reflect.TypeOf(MongoIndex_Status{}), generators)

	return mongoIndex_statusGenerator
}

// AddRelatedPropertyGeneratorsForMongoIndex_Status is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoIndex_Status(gens map[string]gopter.Gen) {
	gens["Key"] = gen.PtrOf(MongoIndexKeys_StatusGenerator())
	gens["Options"] = gen.PtrOf(MongoIndexOptions_StatusGenerator())
}

func Test_MongoIndexKeys_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndexKeys via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexKeys, MongoIndexKeysGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexKeys runs a test to see if a specific instance of MongoIndexKeys round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexKeys(subject MongoIndexKeys) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndexKeys
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

// Generator of MongoIndexKeys instances for property testing - lazily instantiated by MongoIndexKeysGenerator()
var mongoIndexKeysGenerator gopter.Gen

// MongoIndexKeysGenerator returns a generator of MongoIndexKeys instances for property testing.
func MongoIndexKeysGenerator() gopter.Gen {
	if mongoIndexKeysGenerator != nil {
		return mongoIndexKeysGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoIndexKeys(generators)
	mongoIndexKeysGenerator = gen.Struct(reflect.TypeOf(MongoIndexKeys{}), generators)

	return mongoIndexKeysGenerator
}

// AddIndependentPropertyGeneratorsForMongoIndexKeys is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoIndexKeys(gens map[string]gopter.Gen) {
	gens["Keys"] = gen.SliceOf(gen.AlphaString())
}

func Test_MongoIndexKeys_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndexKeys_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexKeys_Status, MongoIndexKeys_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexKeys_Status runs a test to see if a specific instance of MongoIndexKeys_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexKeys_Status(subject MongoIndexKeys_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndexKeys_Status
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

// Generator of MongoIndexKeys_Status instances for property testing - lazily instantiated by
//MongoIndexKeys_StatusGenerator()
var mongoIndexKeys_statusGenerator gopter.Gen

// MongoIndexKeys_StatusGenerator returns a generator of MongoIndexKeys_Status instances for property testing.
func MongoIndexKeys_StatusGenerator() gopter.Gen {
	if mongoIndexKeys_statusGenerator != nil {
		return mongoIndexKeys_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoIndexKeys_Status(generators)
	mongoIndexKeys_statusGenerator = gen.Struct(reflect.TypeOf(MongoIndexKeys_Status{}), generators)

	return mongoIndexKeys_statusGenerator
}

// AddIndependentPropertyGeneratorsForMongoIndexKeys_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoIndexKeys_Status(gens map[string]gopter.Gen) {
	gens["Keys"] = gen.SliceOf(gen.AlphaString())
}

func Test_MongoIndexOptions_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndexOptions via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexOptions, MongoIndexOptionsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexOptions runs a test to see if a specific instance of MongoIndexOptions round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexOptions(subject MongoIndexOptions) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndexOptions
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

// Generator of MongoIndexOptions instances for property testing - lazily instantiated by MongoIndexOptionsGenerator()
var mongoIndexOptionsGenerator gopter.Gen

// MongoIndexOptionsGenerator returns a generator of MongoIndexOptions instances for property testing.
func MongoIndexOptionsGenerator() gopter.Gen {
	if mongoIndexOptionsGenerator != nil {
		return mongoIndexOptionsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoIndexOptions(generators)
	mongoIndexOptionsGenerator = gen.Struct(reflect.TypeOf(MongoIndexOptions{}), generators)

	return mongoIndexOptionsGenerator
}

// AddIndependentPropertyGeneratorsForMongoIndexOptions is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoIndexOptions(gens map[string]gopter.Gen) {
	gens["ExpireAfterSeconds"] = gen.PtrOf(gen.Int())
	gens["Unique"] = gen.PtrOf(gen.Bool())
}

func Test_MongoIndexOptions_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndexOptions_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexOptions_Status, MongoIndexOptions_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexOptions_Status runs a test to see if a specific instance of MongoIndexOptions_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexOptions_Status(subject MongoIndexOptions_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndexOptions_Status
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

// Generator of MongoIndexOptions_Status instances for property testing - lazily instantiated by
//MongoIndexOptions_StatusGenerator()
var mongoIndexOptions_statusGenerator gopter.Gen

// MongoIndexOptions_StatusGenerator returns a generator of MongoIndexOptions_Status instances for property testing.
func MongoIndexOptions_StatusGenerator() gopter.Gen {
	if mongoIndexOptions_statusGenerator != nil {
		return mongoIndexOptions_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoIndexOptions_Status(generators)
	mongoIndexOptions_statusGenerator = gen.Struct(reflect.TypeOf(MongoIndexOptions_Status{}), generators)

	return mongoIndexOptions_statusGenerator
}

// AddIndependentPropertyGeneratorsForMongoIndexOptions_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoIndexOptions_Status(gens map[string]gopter.Gen) {
	gens["ExpireAfterSeconds"] = gen.PtrOf(gen.Int())
	gens["Unique"] = gen.PtrOf(gen.Bool())
}

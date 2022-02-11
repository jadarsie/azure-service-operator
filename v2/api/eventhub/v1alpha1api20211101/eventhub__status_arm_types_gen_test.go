// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

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

func Test_Eventhub_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Eventhub_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEventhub_StatusARM, Eventhub_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEventhub_StatusARM runs a test to see if a specific instance of Eventhub_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEventhub_StatusARM(subject Eventhub_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Eventhub_StatusARM
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

// Generator of Eventhub_StatusARM instances for property testing - lazily instantiated by Eventhub_StatusARMGenerator()
var eventhub_statusARMGenerator gopter.Gen

// Eventhub_StatusARMGenerator returns a generator of Eventhub_StatusARM instances for property testing.
// We first initialize eventhub_statusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Eventhub_StatusARMGenerator() gopter.Gen {
	if eventhub_statusARMGenerator != nil {
		return eventhub_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventhub_StatusARM(generators)
	eventhub_statusARMGenerator = gen.Struct(reflect.TypeOf(Eventhub_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventhub_StatusARM(generators)
	AddRelatedPropertyGeneratorsForEventhub_StatusARM(generators)
	eventhub_statusARMGenerator = gen.Struct(reflect.TypeOf(Eventhub_StatusARM{}), generators)

	return eventhub_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForEventhub_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEventhub_StatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForEventhub_StatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEventhub_StatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(Eventhub_Properties_StatusARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_StatusARMGenerator())
}

func Test_Eventhub_Properties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Eventhub_Properties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEventhub_Properties_StatusARM, Eventhub_Properties_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEventhub_Properties_StatusARM runs a test to see if a specific instance of Eventhub_Properties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEventhub_Properties_StatusARM(subject Eventhub_Properties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Eventhub_Properties_StatusARM
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

// Generator of Eventhub_Properties_StatusARM instances for property testing - lazily instantiated by
//Eventhub_Properties_StatusARMGenerator()
var eventhub_properties_statusARMGenerator gopter.Gen

// Eventhub_Properties_StatusARMGenerator returns a generator of Eventhub_Properties_StatusARM instances for property testing.
// We first initialize eventhub_properties_statusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Eventhub_Properties_StatusARMGenerator() gopter.Gen {
	if eventhub_properties_statusARMGenerator != nil {
		return eventhub_properties_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventhub_Properties_StatusARM(generators)
	eventhub_properties_statusARMGenerator = gen.Struct(reflect.TypeOf(Eventhub_Properties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventhub_Properties_StatusARM(generators)
	AddRelatedPropertyGeneratorsForEventhub_Properties_StatusARM(generators)
	eventhub_properties_statusARMGenerator = gen.Struct(reflect.TypeOf(Eventhub_Properties_StatusARM{}), generators)

	return eventhub_properties_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForEventhub_Properties_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEventhub_Properties_StatusARM(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["MessageRetentionInDays"] = gen.PtrOf(gen.Int())
	gens["PartitionCount"] = gen.PtrOf(gen.Int())
	gens["PartitionIds"] = gen.SliceOf(gen.AlphaString())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(
		Eventhub_Properties_Status_StatusActive,
		Eventhub_Properties_Status_StatusCreating,
		Eventhub_Properties_Status_StatusDeleting,
		Eventhub_Properties_Status_StatusDisabled,
		Eventhub_Properties_Status_StatusReceiveDisabled,
		Eventhub_Properties_Status_StatusRenaming,
		Eventhub_Properties_Status_StatusRestoring,
		Eventhub_Properties_Status_StatusSendDisabled,
		Eventhub_Properties_Status_StatusUnknown))
	gens["UpdatedAt"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForEventhub_Properties_StatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEventhub_Properties_StatusARM(gens map[string]gopter.Gen) {
	gens["CaptureDescription"] = gen.PtrOf(CaptureDescription_StatusARMGenerator())
}

func Test_CaptureDescription_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CaptureDescription_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCaptureDescription_StatusARM, CaptureDescription_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCaptureDescription_StatusARM runs a test to see if a specific instance of CaptureDescription_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCaptureDescription_StatusARM(subject CaptureDescription_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CaptureDescription_StatusARM
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

// Generator of CaptureDescription_StatusARM instances for property testing - lazily instantiated by
//CaptureDescription_StatusARMGenerator()
var captureDescription_statusARMGenerator gopter.Gen

// CaptureDescription_StatusARMGenerator returns a generator of CaptureDescription_StatusARM instances for property testing.
// We first initialize captureDescription_statusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func CaptureDescription_StatusARMGenerator() gopter.Gen {
	if captureDescription_statusARMGenerator != nil {
		return captureDescription_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCaptureDescription_StatusARM(generators)
	captureDescription_statusARMGenerator = gen.Struct(reflect.TypeOf(CaptureDescription_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCaptureDescription_StatusARM(generators)
	AddRelatedPropertyGeneratorsForCaptureDescription_StatusARM(generators)
	captureDescription_statusARMGenerator = gen.Struct(reflect.TypeOf(CaptureDescription_StatusARM{}), generators)

	return captureDescription_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForCaptureDescription_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCaptureDescription_StatusARM(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["Encoding"] = gen.PtrOf(gen.OneConstOf(CaptureDescription_Encoding_StatusAvro, CaptureDescription_Encoding_StatusAvroDeflate))
	gens["IntervalInSeconds"] = gen.PtrOf(gen.Int())
	gens["SizeLimitInBytes"] = gen.PtrOf(gen.Int())
	gens["SkipEmptyArchives"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForCaptureDescription_StatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCaptureDescription_StatusARM(gens map[string]gopter.Gen) {
	gens["Destination"] = gen.PtrOf(Destination_StatusARMGenerator())
}

func Test_Destination_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Destination_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDestination_StatusARM, Destination_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDestination_StatusARM runs a test to see if a specific instance of Destination_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDestination_StatusARM(subject Destination_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Destination_StatusARM
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

// Generator of Destination_StatusARM instances for property testing - lazily instantiated by
//Destination_StatusARMGenerator()
var destination_statusARMGenerator gopter.Gen

// Destination_StatusARMGenerator returns a generator of Destination_StatusARM instances for property testing.
// We first initialize destination_statusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Destination_StatusARMGenerator() gopter.Gen {
	if destination_statusARMGenerator != nil {
		return destination_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDestination_StatusARM(generators)
	destination_statusARMGenerator = gen.Struct(reflect.TypeOf(Destination_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDestination_StatusARM(generators)
	AddRelatedPropertyGeneratorsForDestination_StatusARM(generators)
	destination_statusARMGenerator = gen.Struct(reflect.TypeOf(Destination_StatusARM{}), generators)

	return destination_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForDestination_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDestination_StatusARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDestination_StatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDestination_StatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(Destination_Properties_StatusARMGenerator())
}

func Test_Destination_Properties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Destination_Properties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDestination_Properties_StatusARM, Destination_Properties_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDestination_Properties_StatusARM runs a test to see if a specific instance of Destination_Properties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDestination_Properties_StatusARM(subject Destination_Properties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Destination_Properties_StatusARM
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

// Generator of Destination_Properties_StatusARM instances for property testing - lazily instantiated by
//Destination_Properties_StatusARMGenerator()
var destination_properties_statusARMGenerator gopter.Gen

// Destination_Properties_StatusARMGenerator returns a generator of Destination_Properties_StatusARM instances for property testing.
func Destination_Properties_StatusARMGenerator() gopter.Gen {
	if destination_properties_statusARMGenerator != nil {
		return destination_properties_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDestination_Properties_StatusARM(generators)
	destination_properties_statusARMGenerator = gen.Struct(reflect.TypeOf(Destination_Properties_StatusARM{}), generators)

	return destination_properties_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForDestination_Properties_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDestination_Properties_StatusARM(gens map[string]gopter.Gen) {
	gens["ArchiveNameFormat"] = gen.PtrOf(gen.AlphaString())
	gens["BlobContainer"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeAccountName"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeFolderPath"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeSubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["StorageAccountResourceId"] = gen.PtrOf(gen.AlphaString())
}

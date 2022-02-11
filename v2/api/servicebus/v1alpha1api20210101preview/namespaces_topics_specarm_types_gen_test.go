// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210101preview

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

func Test_NamespacesTopics_SPECARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesTopics_SPECARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesTopics_SPECARM, NamespacesTopics_SPECARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesTopics_SPECARM runs a test to see if a specific instance of NamespacesTopics_SPECARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesTopics_SPECARM(subject NamespacesTopics_SPECARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesTopics_SPECARM
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

// Generator of NamespacesTopics_SPECARM instances for property testing - lazily instantiated by
//NamespacesTopics_SPECARMGenerator()
var namespacesTopics_specarmGenerator gopter.Gen

// NamespacesTopics_SPECARMGenerator returns a generator of NamespacesTopics_SPECARM instances for property testing.
// We first initialize namespacesTopics_specarmGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesTopics_SPECARMGenerator() gopter.Gen {
	if namespacesTopics_specarmGenerator != nil {
		return namespacesTopics_specarmGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesTopics_SPECARM(generators)
	namespacesTopics_specarmGenerator = gen.Struct(reflect.TypeOf(NamespacesTopics_SPECARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesTopics_SPECARM(generators)
	AddRelatedPropertyGeneratorsForNamespacesTopics_SPECARM(generators)
	namespacesTopics_specarmGenerator = gen.Struct(reflect.TypeOf(NamespacesTopics_SPECARM{}), generators)

	return namespacesTopics_specarmGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesTopics_SPECARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesTopics_SPECARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForNamespacesTopics_SPECARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesTopics_SPECARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SBTopicProperties_SpecARMGenerator())
}

func Test_SBTopicProperties_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SBTopicProperties_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSBTopicProperties_SpecARM, SBTopicProperties_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSBTopicProperties_SpecARM runs a test to see if a specific instance of SBTopicProperties_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSBTopicProperties_SpecARM(subject SBTopicProperties_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SBTopicProperties_SpecARM
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

// Generator of SBTopicProperties_SpecARM instances for property testing - lazily instantiated by
//SBTopicProperties_SpecARMGenerator()
var sbTopicProperties_specARMGenerator gopter.Gen

// SBTopicProperties_SpecARMGenerator returns a generator of SBTopicProperties_SpecARM instances for property testing.
func SBTopicProperties_SpecARMGenerator() gopter.Gen {
	if sbTopicProperties_specARMGenerator != nil {
		return sbTopicProperties_specARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBTopicProperties_SpecARM(generators)
	sbTopicProperties_specARMGenerator = gen.Struct(reflect.TypeOf(SBTopicProperties_SpecARM{}), generators)

	return sbTopicProperties_specARMGenerator
}

// AddIndependentPropertyGeneratorsForSBTopicProperties_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSBTopicProperties_SpecARM(gens map[string]gopter.Gen) {
	gens["AutoDeleteOnIdle"] = gen.PtrOf(gen.AlphaString())
	gens["DefaultMessageTimeToLive"] = gen.PtrOf(gen.AlphaString())
	gens["DuplicateDetectionHistoryTimeWindow"] = gen.PtrOf(gen.AlphaString())
	gens["EnableBatchedOperations"] = gen.PtrOf(gen.Bool())
	gens["EnableExpress"] = gen.PtrOf(gen.Bool())
	gens["EnablePartitioning"] = gen.PtrOf(gen.Bool())
	gens["MaxSizeInMegabytes"] = gen.PtrOf(gen.Int())
	gens["RequiresDuplicateDetection"] = gen.PtrOf(gen.Bool())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(
		EntityStatus_SpecActive,
		EntityStatus_SpecCreating,
		EntityStatus_SpecDeleting,
		EntityStatus_SpecDisabled,
		EntityStatus_SpecReceiveDisabled,
		EntityStatus_SpecRenaming,
		EntityStatus_SpecRestoring,
		EntityStatus_SpecSendDisabled,
		EntityStatus_SpecUnknown))
	gens["SupportOrdering"] = gen.PtrOf(gen.Bool())
}

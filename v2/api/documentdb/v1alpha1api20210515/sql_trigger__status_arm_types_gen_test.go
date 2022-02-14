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

func Test_SqlTrigger_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlTrigger_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlTrigger_StatusARM, SqlTrigger_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlTrigger_StatusARM runs a test to see if a specific instance of SqlTrigger_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlTrigger_StatusARM(subject SqlTrigger_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlTrigger_StatusARM
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

// Generator of SqlTrigger_StatusARM instances for property testing - lazily instantiated by
//SqlTrigger_StatusARMGenerator()
var sqlTrigger_statusARMGenerator gopter.Gen

// SqlTrigger_StatusARMGenerator returns a generator of SqlTrigger_StatusARM instances for property testing.
// We first initialize sqlTrigger_statusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlTrigger_StatusARMGenerator() gopter.Gen {
	if sqlTrigger_statusARMGenerator != nil {
		return sqlTrigger_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlTrigger_StatusARM(generators)
	sqlTrigger_statusARMGenerator = gen.Struct(reflect.TypeOf(SqlTrigger_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlTrigger_StatusARM(generators)
	AddRelatedPropertyGeneratorsForSqlTrigger_StatusARM(generators)
	sqlTrigger_statusARMGenerator = gen.Struct(reflect.TypeOf(SqlTrigger_StatusARM{}), generators)

	return sqlTrigger_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlTrigger_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlTrigger_StatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlTrigger_StatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlTrigger_StatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SqlTriggerProperties_StatusARMGenerator())
}

func Test_SqlTriggerProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlTriggerProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlTriggerProperties_StatusARM, SqlTriggerProperties_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlTriggerProperties_StatusARM runs a test to see if a specific instance of SqlTriggerProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlTriggerProperties_StatusARM(subject SqlTriggerProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlTriggerProperties_StatusARM
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

// Generator of SqlTriggerProperties_StatusARM instances for property testing - lazily instantiated by
//SqlTriggerProperties_StatusARMGenerator()
var sqlTriggerProperties_statusARMGenerator gopter.Gen

// SqlTriggerProperties_StatusARMGenerator returns a generator of SqlTriggerProperties_StatusARM instances for property testing.
func SqlTriggerProperties_StatusARMGenerator() gopter.Gen {
	if sqlTriggerProperties_statusARMGenerator != nil {
		return sqlTriggerProperties_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlTriggerProperties_StatusARM(generators)
	sqlTriggerProperties_statusARMGenerator = gen.Struct(reflect.TypeOf(SqlTriggerProperties_StatusARM{}), generators)

	return sqlTriggerProperties_statusARMGenerator
}

// AddRelatedPropertyGeneratorsForSqlTriggerProperties_StatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlTriggerProperties_StatusARM(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptions_StatusARMGenerator())
	gens["Resource"] = SqlTriggerResource_StatusARMGenerator()
}

func Test_SqlTriggerResource_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlTriggerResource_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlTriggerResource_StatusARM, SqlTriggerResource_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlTriggerResource_StatusARM runs a test to see if a specific instance of SqlTriggerResource_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlTriggerResource_StatusARM(subject SqlTriggerResource_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlTriggerResource_StatusARM
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

// Generator of SqlTriggerResource_StatusARM instances for property testing - lazily instantiated by
//SqlTriggerResource_StatusARMGenerator()
var sqlTriggerResource_statusARMGenerator gopter.Gen

// SqlTriggerResource_StatusARMGenerator returns a generator of SqlTriggerResource_StatusARM instances for property testing.
func SqlTriggerResource_StatusARMGenerator() gopter.Gen {
	if sqlTriggerResource_statusARMGenerator != nil {
		return sqlTriggerResource_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlTriggerResource_StatusARM(generators)
	sqlTriggerResource_statusARMGenerator = gen.Struct(reflect.TypeOf(SqlTriggerResource_StatusARM{}), generators)

	return sqlTriggerResource_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlTriggerResource_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlTriggerResource_StatusARM(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.AlphaString()
	gens["TriggerOperation"] = gen.PtrOf(gen.AlphaString())
	gens["TriggerType"] = gen.PtrOf(gen.AlphaString())
}

// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

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

func Test_AuthorizationRule_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationRule_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationRuleStatusARM, AuthorizationRuleStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationRuleStatusARM runs a test to see if a specific instance of AuthorizationRule_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationRuleStatusARM(subject AuthorizationRule_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationRule_StatusARM
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

// Generator of AuthorizationRule_StatusARM instances for property testing - lazily instantiated by
//AuthorizationRuleStatusARMGenerator()
var authorizationRuleStatusARMGenerator gopter.Gen

// AuthorizationRuleStatusARMGenerator returns a generator of AuthorizationRule_StatusARM instances for property testing.
// We first initialize authorizationRuleStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AuthorizationRuleStatusARMGenerator() gopter.Gen {
	if authorizationRuleStatusARMGenerator != nil {
		return authorizationRuleStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationRuleStatusARM(generators)
	authorizationRuleStatusARMGenerator = gen.Struct(reflect.TypeOf(AuthorizationRule_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationRuleStatusARM(generators)
	AddRelatedPropertyGeneratorsForAuthorizationRuleStatusARM(generators)
	authorizationRuleStatusARMGenerator = gen.Struct(reflect.TypeOf(AuthorizationRule_StatusARM{}), generators)

	return authorizationRuleStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForAuthorizationRuleStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAuthorizationRuleStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForAuthorizationRuleStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAuthorizationRuleStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(AuthorizationRuleStatusPropertiesARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataStatusARMGenerator())
}

func Test_AuthorizationRule_Status_PropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationRule_Status_PropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationRuleStatusPropertiesARM, AuthorizationRuleStatusPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationRuleStatusPropertiesARM runs a test to see if a specific instance of AuthorizationRule_Status_PropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationRuleStatusPropertiesARM(subject AuthorizationRule_Status_PropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationRule_Status_PropertiesARM
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

// Generator of AuthorizationRule_Status_PropertiesARM instances for property testing - lazily instantiated by
//AuthorizationRuleStatusPropertiesARMGenerator()
var authorizationRuleStatusPropertiesARMGenerator gopter.Gen

// AuthorizationRuleStatusPropertiesARMGenerator returns a generator of AuthorizationRule_Status_PropertiesARM instances for property testing.
func AuthorizationRuleStatusPropertiesARMGenerator() gopter.Gen {
	if authorizationRuleStatusPropertiesARMGenerator != nil {
		return authorizationRuleStatusPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationRuleStatusPropertiesARM(generators)
	authorizationRuleStatusPropertiesARMGenerator = gen.Struct(reflect.TypeOf(AuthorizationRule_Status_PropertiesARM{}), generators)

	return authorizationRuleStatusPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForAuthorizationRuleStatusPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAuthorizationRuleStatusPropertiesARM(gens map[string]gopter.Gen) {
	gens["Rights"] = gen.SliceOf(gen.OneConstOf(AuthorizationRuleStatusPropertiesRightsListen, AuthorizationRuleStatusPropertiesRightsManage, AuthorizationRuleStatusPropertiesRightsSend))
}

func Test_SystemData_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemDataStatusARM, SystemDataStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemDataStatusARM runs a test to see if a specific instance of SystemData_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemDataStatusARM(subject SystemData_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_StatusARM
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

// Generator of SystemData_StatusARM instances for property testing - lazily instantiated by
//SystemDataStatusARMGenerator()
var systemDataStatusARMGenerator gopter.Gen

// SystemDataStatusARMGenerator returns a generator of SystemData_StatusARM instances for property testing.
func SystemDataStatusARMGenerator() gopter.Gen {
	if systemDataStatusARMGenerator != nil {
		return systemDataStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemDataStatusARM(generators)
	systemDataStatusARMGenerator = gen.Struct(reflect.TypeOf(SystemData_StatusARM{}), generators)

	return systemDataStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForSystemDataStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemDataStatusARM(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemDataStatusCreatedByTypeApplication,
		SystemDataStatusCreatedByTypeKey,
		SystemDataStatusCreatedByTypeManagedIdentity,
		SystemDataStatusCreatedByTypeUser))
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemDataStatusLastModifiedByTypeApplication,
		SystemDataStatusLastModifiedByTypeKey,
		SystemDataStatusLastModifiedByTypeManagedIdentity,
		SystemDataStatusLastModifiedByTypeUser))
}

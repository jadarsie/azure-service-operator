// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201storage

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

func Test_RedisLinkedServer_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisLinkedServer via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisLinkedServer, RedisLinkedServerGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisLinkedServer runs a test to see if a specific instance of RedisLinkedServer round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisLinkedServer(subject RedisLinkedServer) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisLinkedServer
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

// Generator of RedisLinkedServer instances for property testing - lazily instantiated by RedisLinkedServerGenerator()
var redisLinkedServerGenerator gopter.Gen

// RedisLinkedServerGenerator returns a generator of RedisLinkedServer instances for property testing.
func RedisLinkedServerGenerator() gopter.Gen {
	if redisLinkedServerGenerator != nil {
		return redisLinkedServerGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRedisLinkedServer(generators)
	redisLinkedServerGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServer{}), generators)

	return redisLinkedServerGenerator
}

// AddRelatedPropertyGeneratorsForRedisLinkedServer is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisLinkedServer(gens map[string]gopter.Gen) {
	gens["Spec"] = RedisLinkedServer_SpecGenerator()
	gens["Status"] = RedisLinkedServer_StatusGenerator()
}

func Test_RedisLinkedServer_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisLinkedServer_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisLinkedServer_Spec, RedisLinkedServer_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisLinkedServer_Spec runs a test to see if a specific instance of RedisLinkedServer_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisLinkedServer_Spec(subject RedisLinkedServer_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisLinkedServer_Spec
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

// Generator of RedisLinkedServer_Spec instances for property testing - lazily instantiated by
//RedisLinkedServer_SpecGenerator()
var redisLinkedServer_specGenerator gopter.Gen

// RedisLinkedServer_SpecGenerator returns a generator of RedisLinkedServer_Spec instances for property testing.
func RedisLinkedServer_SpecGenerator() gopter.Gen {
	if redisLinkedServer_specGenerator != nil {
		return redisLinkedServer_specGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisLinkedServer_Spec(generators)
	redisLinkedServer_specGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServer_Spec{}), generators)

	return redisLinkedServer_specGenerator
}

// AddIndependentPropertyGeneratorsForRedisLinkedServer_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisLinkedServer_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["LinkedRedisCacheLocation"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["ServerRole"] = gen.PtrOf(gen.AlphaString())
}

func Test_RedisLinkedServer_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisLinkedServer_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisLinkedServer_Status, RedisLinkedServer_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisLinkedServer_Status runs a test to see if a specific instance of RedisLinkedServer_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisLinkedServer_Status(subject RedisLinkedServer_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisLinkedServer_Status
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

// Generator of RedisLinkedServer_Status instances for property testing - lazily instantiated by
//RedisLinkedServer_StatusGenerator()
var redisLinkedServer_statusGenerator gopter.Gen

// RedisLinkedServer_StatusGenerator returns a generator of RedisLinkedServer_Status instances for property testing.
func RedisLinkedServer_StatusGenerator() gopter.Gen {
	if redisLinkedServer_statusGenerator != nil {
		return redisLinkedServer_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisLinkedServer_Status(generators)
	redisLinkedServer_statusGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServer_Status{}), generators)

	return redisLinkedServer_statusGenerator
}

// AddIndependentPropertyGeneratorsForRedisLinkedServer_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisLinkedServer_Status(gens map[string]gopter.Gen) {
	gens["LinkedRedisCacheId"] = gen.PtrOf(gen.AlphaString())
	gens["LinkedRedisCacheLocation"] = gen.PtrOf(gen.AlphaString())
	gens["ServerRole"] = gen.PtrOf(gen.AlphaString())
}

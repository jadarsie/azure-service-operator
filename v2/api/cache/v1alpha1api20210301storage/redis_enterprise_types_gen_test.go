// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210301storage

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

func Test_RedisEnterprise_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisEnterprise via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisEnterprise, RedisEnterpriseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisEnterprise runs a test to see if a specific instance of RedisEnterprise round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisEnterprise(subject RedisEnterprise) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisEnterprise
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

// Generator of RedisEnterprise instances for property testing - lazily instantiated by RedisEnterpriseGenerator()
var redisEnterpriseGenerator gopter.Gen

// RedisEnterpriseGenerator returns a generator of RedisEnterprise instances for property testing.
func RedisEnterpriseGenerator() gopter.Gen {
	if redisEnterpriseGenerator != nil {
		return redisEnterpriseGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRedisEnterprise(generators)
	redisEnterpriseGenerator = gen.Struct(reflect.TypeOf(RedisEnterprise{}), generators)

	return redisEnterpriseGenerator
}

// AddRelatedPropertyGeneratorsForRedisEnterprise is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisEnterprise(gens map[string]gopter.Gen) {
	gens["Spec"] = RedisEnterprise_SPECGenerator()
	gens["Status"] = Cluster_StatusGenerator()
}

func Test_Cluster_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Cluster_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCluster_Status, Cluster_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCluster_Status runs a test to see if a specific instance of Cluster_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForCluster_Status(subject Cluster_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Cluster_Status
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

// Generator of Cluster_Status instances for property testing - lazily instantiated by Cluster_StatusGenerator()
var cluster_statusGenerator gopter.Gen

// Cluster_StatusGenerator returns a generator of Cluster_Status instances for property testing.
// We first initialize cluster_statusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Cluster_StatusGenerator() gopter.Gen {
	if cluster_statusGenerator != nil {
		return cluster_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCluster_Status(generators)
	cluster_statusGenerator = gen.Struct(reflect.TypeOf(Cluster_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCluster_Status(generators)
	AddRelatedPropertyGeneratorsForCluster_Status(generators)
	cluster_statusGenerator = gen.Struct(reflect.TypeOf(Cluster_Status{}), generators)

	return cluster_statusGenerator
}

// AddIndependentPropertyGeneratorsForCluster_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCluster_Status(gens map[string]gopter.Gen) {
	gens["HostName"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MinimumTlsVersion"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["RedisVersion"] = gen.PtrOf(gen.AlphaString())
	gens["ResourceState"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForCluster_Status is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCluster_Status(gens map[string]gopter.Gen) {
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnection_Status_SubResourceEmbeddedGenerator())
	gens["Sku"] = gen.PtrOf(Sku_StatusGenerator())
}

func Test_RedisEnterprise_SPEC_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisEnterprise_SPEC via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisEnterprise_SPEC, RedisEnterprise_SPECGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisEnterprise_SPEC runs a test to see if a specific instance of RedisEnterprise_SPEC round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisEnterprise_SPEC(subject RedisEnterprise_SPEC) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisEnterprise_SPEC
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

// Generator of RedisEnterprise_SPEC instances for property testing - lazily instantiated by
//RedisEnterprise_SPECGenerator()
var redisEnterprise_specGenerator gopter.Gen

// RedisEnterprise_SPECGenerator returns a generator of RedisEnterprise_SPEC instances for property testing.
// We first initialize redisEnterprise_specGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisEnterprise_SPECGenerator() gopter.Gen {
	if redisEnterprise_specGenerator != nil {
		return redisEnterprise_specGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterprise_SPEC(generators)
	redisEnterprise_specGenerator = gen.Struct(reflect.TypeOf(RedisEnterprise_SPEC{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterprise_SPEC(generators)
	AddRelatedPropertyGeneratorsForRedisEnterprise_SPEC(generators)
	redisEnterprise_specGenerator = gen.Struct(reflect.TypeOf(RedisEnterprise_SPEC{}), generators)

	return redisEnterprise_specGenerator
}

// AddIndependentPropertyGeneratorsForRedisEnterprise_SPEC is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisEnterprise_SPEC(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MinimumTlsVersion"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisEnterprise_SPEC is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisEnterprise_SPEC(gens map[string]gopter.Gen) {
	gens["Sku"] = gen.PtrOf(Sku_SpecGenerator())
}

func Test_PrivateEndpointConnection_Status_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_Status_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnection_Status_SubResourceEmbedded, PrivateEndpointConnection_Status_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnection_Status_SubResourceEmbedded runs a test to see if a specific instance of PrivateEndpointConnection_Status_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnection_Status_SubResourceEmbedded(subject PrivateEndpointConnection_Status_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnection_Status_SubResourceEmbedded
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

// Generator of PrivateEndpointConnection_Status_SubResourceEmbedded instances for property testing - lazily
//instantiated by PrivateEndpointConnection_Status_SubResourceEmbeddedGenerator()
var privateEndpointConnection_status_subResourceEmbeddedGenerator gopter.Gen

// PrivateEndpointConnection_Status_SubResourceEmbeddedGenerator returns a generator of PrivateEndpointConnection_Status_SubResourceEmbedded instances for property testing.
func PrivateEndpointConnection_Status_SubResourceEmbeddedGenerator() gopter.Gen {
	if privateEndpointConnection_status_subResourceEmbeddedGenerator != nil {
		return privateEndpointConnection_status_subResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnection_Status_SubResourceEmbedded(generators)
	privateEndpointConnection_status_subResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_Status_SubResourceEmbedded{}), generators)

	return privateEndpointConnection_status_subResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnection_Status_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnection_Status_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_Sku_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku_Spec, Sku_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku_Spec runs a test to see if a specific instance of Sku_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForSku_Spec(subject Sku_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_Spec
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

// Generator of Sku_Spec instances for property testing - lazily instantiated by Sku_SpecGenerator()
var sku_specGenerator gopter.Gen

// Sku_SpecGenerator returns a generator of Sku_Spec instances for property testing.
func Sku_SpecGenerator() gopter.Gen {
	if sku_specGenerator != nil {
		return sku_specGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku_Spec(generators)
	sku_specGenerator = gen.Struct(reflect.TypeOf(Sku_Spec{}), generators)

	return sku_specGenerator
}

// AddIndependentPropertyGeneratorsForSku_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku_Spec(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

func Test_Sku_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku_Status, Sku_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku_Status runs a test to see if a specific instance of Sku_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForSku_Status(subject Sku_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_Status
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

// Generator of Sku_Status instances for property testing - lazily instantiated by Sku_StatusGenerator()
var sku_statusGenerator gopter.Gen

// Sku_StatusGenerator returns a generator of Sku_Status instances for property testing.
func Sku_StatusGenerator() gopter.Gen {
	if sku_statusGenerator != nil {
		return sku_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku_Status(generators)
	sku_statusGenerator = gen.Struct(reflect.TypeOf(Sku_Status{}), generators)

	return sku_statusGenerator
}

// AddIndependentPropertyGeneratorsForSku_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku_Status(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

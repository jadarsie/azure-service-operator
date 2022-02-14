// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

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

func Test_Topic_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Topic_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTopic_StatusARM, Topic_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTopic_StatusARM runs a test to see if a specific instance of Topic_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForTopic_StatusARM(subject Topic_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Topic_StatusARM
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

// Generator of Topic_StatusARM instances for property testing - lazily instantiated by Topic_StatusARMGenerator()
var topic_statusARMGenerator gopter.Gen

// Topic_StatusARMGenerator returns a generator of Topic_StatusARM instances for property testing.
// We first initialize topic_statusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Topic_StatusARMGenerator() gopter.Gen {
	if topic_statusARMGenerator != nil {
		return topic_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTopic_StatusARM(generators)
	topic_statusARMGenerator = gen.Struct(reflect.TypeOf(Topic_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTopic_StatusARM(generators)
	AddRelatedPropertyGeneratorsForTopic_StatusARM(generators)
	topic_statusARMGenerator = gen.Struct(reflect.TypeOf(Topic_StatusARM{}), generators)

	return topic_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForTopic_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTopic_StatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForTopic_StatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTopic_StatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(TopicProperties_StatusARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_StatusARMGenerator())
}

func Test_TopicProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TopicProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTopicProperties_StatusARM, TopicProperties_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTopicProperties_StatusARM runs a test to see if a specific instance of TopicProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForTopicProperties_StatusARM(subject TopicProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TopicProperties_StatusARM
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

// Generator of TopicProperties_StatusARM instances for property testing - lazily instantiated by
//TopicProperties_StatusARMGenerator()
var topicProperties_statusARMGenerator gopter.Gen

// TopicProperties_StatusARMGenerator returns a generator of TopicProperties_StatusARM instances for property testing.
// We first initialize topicProperties_statusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func TopicProperties_StatusARMGenerator() gopter.Gen {
	if topicProperties_statusARMGenerator != nil {
		return topicProperties_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTopicProperties_StatusARM(generators)
	topicProperties_statusARMGenerator = gen.Struct(reflect.TypeOf(TopicProperties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTopicProperties_StatusARM(generators)
	AddRelatedPropertyGeneratorsForTopicProperties_StatusARM(generators)
	topicProperties_statusARMGenerator = gen.Struct(reflect.TypeOf(TopicProperties_StatusARM{}), generators)

	return topicProperties_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForTopicProperties_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTopicProperties_StatusARM(gens map[string]gopter.Gen) {
	gens["Endpoint"] = gen.PtrOf(gen.AlphaString())
	gens["InputSchema"] = gen.PtrOf(gen.AlphaString())
	gens["MetricResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForTopicProperties_StatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTopicProperties_StatusARM(gens map[string]gopter.Gen) {
	gens["InboundIpRules"] = gen.SliceOf(InboundIpRule_StatusARMGenerator())
	gens["InputSchemaMapping"] = gen.PtrOf(InputSchemaMapping_StatusARMGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARMGenerator())
}

func Test_PrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARM, PrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARM runs a test to see if a specific instance of PrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARM(subject PrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARM
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

// Generator of PrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARM instances for property testing - lazily
//instantiated by PrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARMGenerator()
var privateEndpointConnection_status_topic_subResourceEmbeddedARMGenerator gopter.Gen

// PrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARMGenerator returns a generator of PrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARM instances for property testing.
func PrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARMGenerator() gopter.Gen {
	if privateEndpointConnection_status_topic_subResourceEmbeddedARMGenerator != nil {
		return privateEndpointConnection_status_topic_subResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARM(generators)
	privateEndpointConnection_status_topic_subResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARM{}), generators)

	return privateEndpointConnection_status_topic_subResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

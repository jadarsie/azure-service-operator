// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/api/eventgrid/v1alpha1api20200601storage"
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

func Test_Domain_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Domain to hub returns original",
		prop.ForAll(RunResourceConversionTestForDomain, DomainGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForDomain tests if a specific instance of Domain round trips to the hub storage version and back losslessly
func RunResourceConversionTestForDomain(subject Domain) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v1alpha1api20200601storage.Domain
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual Domain
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Domain_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Domain to Domain via AssignPropertiesToDomain & AssignPropertiesFromDomain returns original",
		prop.ForAll(RunPropertyAssignmentTestForDomain, DomainGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDomain tests if a specific instance of Domain can be assigned to v1alpha1api20200601storage and back losslessly
func RunPropertyAssignmentTestForDomain(subject Domain) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20200601storage.Domain
	err := copied.AssignPropertiesToDomain(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Domain
	err = actual.AssignPropertiesFromDomain(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Domain_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Domain via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomain, DomainGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomain runs a test to see if a specific instance of Domain round trips to JSON and back losslessly
func RunJSONSerializationTestForDomain(subject Domain) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Domain
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

// Generator of Domain instances for property testing - lazily instantiated by DomainGenerator()
var domainGenerator gopter.Gen

// DomainGenerator returns a generator of Domain instances for property testing.
func DomainGenerator() gopter.Gen {
	if domainGenerator != nil {
		return domainGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDomain(generators)
	domainGenerator = gen.Struct(reflect.TypeOf(Domain{}), generators)

	return domainGenerator
}

// AddRelatedPropertyGeneratorsForDomain is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDomain(gens map[string]gopter.Gen) {
	gens["Spec"] = Domain_SpecGenerator()
	gens["Status"] = Domain_StatusGenerator()
}

func Test_Domain_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Domain_Spec to Domain_Spec via AssignPropertiesToDomain_Spec & AssignPropertiesFromDomain_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForDomain_Spec, Domain_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDomain_Spec tests if a specific instance of Domain_Spec can be assigned to v1alpha1api20200601storage and back losslessly
func RunPropertyAssignmentTestForDomain_Spec(subject Domain_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20200601storage.Domain_Spec
	err := copied.AssignPropertiesToDomain_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Domain_Spec
	err = actual.AssignPropertiesFromDomain_Spec(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Domain_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Domain_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomain_Spec, Domain_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomain_Spec runs a test to see if a specific instance of Domain_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDomain_Spec(subject Domain_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Domain_Spec
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

// Generator of Domain_Spec instances for property testing - lazily instantiated by Domain_SpecGenerator()
var domain_specGenerator gopter.Gen

// Domain_SpecGenerator returns a generator of Domain_Spec instances for property testing.
// We first initialize domain_specGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Domain_SpecGenerator() gopter.Gen {
	if domain_specGenerator != nil {
		return domain_specGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomain_Spec(generators)
	domain_specGenerator = gen.Struct(reflect.TypeOf(Domain_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomain_Spec(generators)
	AddRelatedPropertyGeneratorsForDomain_Spec(generators)
	domain_specGenerator = gen.Struct(reflect.TypeOf(Domain_Spec{}), generators)

	return domain_specGenerator
}

// AddIndependentPropertyGeneratorsForDomain_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDomain_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["InputSchema"] = gen.PtrOf(gen.OneConstOf(DomainPropertiesInputSchemaCloudEventSchemaV1_0, DomainPropertiesInputSchemaCustomEventSchema, DomainPropertiesInputSchemaEventGridSchema))
	gens["Location"] = gen.AlphaString()
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(DomainPropertiesPublicNetworkAccessDisabled, DomainPropertiesPublicNetworkAccessEnabled))
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDomain_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDomain_Spec(gens map[string]gopter.Gen) {
	gens["InboundIpRules"] = gen.SliceOf(InboundIpRuleGenerator())
	gens["InputSchemaMapping"] = gen.PtrOf(InputSchemaMappingGenerator())
}

func Test_Domain_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Domain_Status to Domain_Status via AssignPropertiesToDomain_Status & AssignPropertiesFromDomain_Status returns original",
		prop.ForAll(RunPropertyAssignmentTestForDomain_Status, Domain_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDomain_Status tests if a specific instance of Domain_Status can be assigned to v1alpha1api20200601storage and back losslessly
func RunPropertyAssignmentTestForDomain_Status(subject Domain_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20200601storage.Domain_Status
	err := copied.AssignPropertiesToDomain_Status(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Domain_Status
	err = actual.AssignPropertiesFromDomain_Status(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Domain_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Domain_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomain_Status, Domain_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomain_Status runs a test to see if a specific instance of Domain_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForDomain_Status(subject Domain_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Domain_Status
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

// Generator of Domain_Status instances for property testing - lazily instantiated by Domain_StatusGenerator()
var domain_statusGenerator gopter.Gen

// Domain_StatusGenerator returns a generator of Domain_Status instances for property testing.
// We first initialize domain_statusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Domain_StatusGenerator() gopter.Gen {
	if domain_statusGenerator != nil {
		return domain_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomain_Status(generators)
	domain_statusGenerator = gen.Struct(reflect.TypeOf(Domain_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomain_Status(generators)
	AddRelatedPropertyGeneratorsForDomain_Status(generators)
	domain_statusGenerator = gen.Struct(reflect.TypeOf(Domain_Status{}), generators)

	return domain_statusGenerator
}

// AddIndependentPropertyGeneratorsForDomain_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDomain_Status(gens map[string]gopter.Gen) {
	gens["Endpoint"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["InputSchema"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MetricResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDomain_Status is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDomain_Status(gens map[string]gopter.Gen) {
	gens["InboundIpRules"] = gen.SliceOf(InboundIpRule_StatusGenerator())
	gens["InputSchemaMapping"] = gen.PtrOf(InputSchemaMapping_StatusGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnection_Status_Domain_SubResourceEmbeddedGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_StatusGenerator())
}

func Test_InboundIpRule_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from InboundIpRule to InboundIpRule via AssignPropertiesToInboundIpRule & AssignPropertiesFromInboundIpRule returns original",
		prop.ForAll(RunPropertyAssignmentTestForInboundIpRule, InboundIpRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForInboundIpRule tests if a specific instance of InboundIpRule can be assigned to v1alpha1api20200601storage and back losslessly
func RunPropertyAssignmentTestForInboundIpRule(subject InboundIpRule) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20200601storage.InboundIpRule
	err := copied.AssignPropertiesToInboundIpRule(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual InboundIpRule
	err = actual.AssignPropertiesFromInboundIpRule(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_InboundIpRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of InboundIpRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForInboundIpRule, InboundIpRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForInboundIpRule runs a test to see if a specific instance of InboundIpRule round trips to JSON and back losslessly
func RunJSONSerializationTestForInboundIpRule(subject InboundIpRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual InboundIpRule
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

// Generator of InboundIpRule instances for property testing - lazily instantiated by InboundIpRuleGenerator()
var inboundIpRuleGenerator gopter.Gen

// InboundIpRuleGenerator returns a generator of InboundIpRule instances for property testing.
func InboundIpRuleGenerator() gopter.Gen {
	if inboundIpRuleGenerator != nil {
		return inboundIpRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInboundIpRule(generators)
	inboundIpRuleGenerator = gen.Struct(reflect.TypeOf(InboundIpRule{}), generators)

	return inboundIpRuleGenerator
}

// AddIndependentPropertyGeneratorsForInboundIpRule is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForInboundIpRule(gens map[string]gopter.Gen) {
	gens["Action"] = gen.PtrOf(gen.OneConstOf(InboundIpRuleActionAllow))
	gens["IpMask"] = gen.PtrOf(gen.AlphaString())
}

func Test_InboundIpRule_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from InboundIpRule_Status to InboundIpRule_Status via AssignPropertiesToInboundIpRule_Status & AssignPropertiesFromInboundIpRule_Status returns original",
		prop.ForAll(RunPropertyAssignmentTestForInboundIpRule_Status, InboundIpRule_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForInboundIpRule_Status tests if a specific instance of InboundIpRule_Status can be assigned to v1alpha1api20200601storage and back losslessly
func RunPropertyAssignmentTestForInboundIpRule_Status(subject InboundIpRule_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20200601storage.InboundIpRule_Status
	err := copied.AssignPropertiesToInboundIpRule_Status(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual InboundIpRule_Status
	err = actual.AssignPropertiesFromInboundIpRule_Status(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_InboundIpRule_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of InboundIpRule_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForInboundIpRule_Status, InboundIpRule_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForInboundIpRule_Status runs a test to see if a specific instance of InboundIpRule_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForInboundIpRule_Status(subject InboundIpRule_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual InboundIpRule_Status
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

// Generator of InboundIpRule_Status instances for property testing - lazily instantiated by
//InboundIpRule_StatusGenerator()
var inboundIpRule_statusGenerator gopter.Gen

// InboundIpRule_StatusGenerator returns a generator of InboundIpRule_Status instances for property testing.
func InboundIpRule_StatusGenerator() gopter.Gen {
	if inboundIpRule_statusGenerator != nil {
		return inboundIpRule_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInboundIpRule_Status(generators)
	inboundIpRule_statusGenerator = gen.Struct(reflect.TypeOf(InboundIpRule_Status{}), generators)

	return inboundIpRule_statusGenerator
}

// AddIndependentPropertyGeneratorsForInboundIpRule_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForInboundIpRule_Status(gens map[string]gopter.Gen) {
	gens["Action"] = gen.PtrOf(gen.AlphaString())
	gens["IpMask"] = gen.PtrOf(gen.AlphaString())
}

func Test_InputSchemaMapping_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from InputSchemaMapping to InputSchemaMapping via AssignPropertiesToInputSchemaMapping & AssignPropertiesFromInputSchemaMapping returns original",
		prop.ForAll(RunPropertyAssignmentTestForInputSchemaMapping, InputSchemaMappingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForInputSchemaMapping tests if a specific instance of InputSchemaMapping can be assigned to v1alpha1api20200601storage and back losslessly
func RunPropertyAssignmentTestForInputSchemaMapping(subject InputSchemaMapping) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20200601storage.InputSchemaMapping
	err := copied.AssignPropertiesToInputSchemaMapping(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual InputSchemaMapping
	err = actual.AssignPropertiesFromInputSchemaMapping(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_InputSchemaMapping_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of InputSchemaMapping via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForInputSchemaMapping, InputSchemaMappingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForInputSchemaMapping runs a test to see if a specific instance of InputSchemaMapping round trips to JSON and back losslessly
func RunJSONSerializationTestForInputSchemaMapping(subject InputSchemaMapping) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual InputSchemaMapping
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

// Generator of InputSchemaMapping instances for property testing - lazily instantiated by InputSchemaMappingGenerator()
var inputSchemaMappingGenerator gopter.Gen

// InputSchemaMappingGenerator returns a generator of InputSchemaMapping instances for property testing.
func InputSchemaMappingGenerator() gopter.Gen {
	if inputSchemaMappingGenerator != nil {
		return inputSchemaMappingGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInputSchemaMapping(generators)
	inputSchemaMappingGenerator = gen.Struct(reflect.TypeOf(InputSchemaMapping{}), generators)

	return inputSchemaMappingGenerator
}

// AddIndependentPropertyGeneratorsForInputSchemaMapping is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForInputSchemaMapping(gens map[string]gopter.Gen) {
	gens["InputSchemaMappingType"] = gen.OneConstOf(InputSchemaMappingInputSchemaMappingTypeJson)
}

func Test_InputSchemaMapping_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from InputSchemaMapping_Status to InputSchemaMapping_Status via AssignPropertiesToInputSchemaMapping_Status & AssignPropertiesFromInputSchemaMapping_Status returns original",
		prop.ForAll(RunPropertyAssignmentTestForInputSchemaMapping_Status, InputSchemaMapping_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForInputSchemaMapping_Status tests if a specific instance of InputSchemaMapping_Status can be assigned to v1alpha1api20200601storage and back losslessly
func RunPropertyAssignmentTestForInputSchemaMapping_Status(subject InputSchemaMapping_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20200601storage.InputSchemaMapping_Status
	err := copied.AssignPropertiesToInputSchemaMapping_Status(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual InputSchemaMapping_Status
	err = actual.AssignPropertiesFromInputSchemaMapping_Status(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_InputSchemaMapping_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of InputSchemaMapping_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForInputSchemaMapping_Status, InputSchemaMapping_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForInputSchemaMapping_Status runs a test to see if a specific instance of InputSchemaMapping_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForInputSchemaMapping_Status(subject InputSchemaMapping_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual InputSchemaMapping_Status
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

// Generator of InputSchemaMapping_Status instances for property testing - lazily instantiated by
//InputSchemaMapping_StatusGenerator()
var inputSchemaMapping_statusGenerator gopter.Gen

// InputSchemaMapping_StatusGenerator returns a generator of InputSchemaMapping_Status instances for property testing.
func InputSchemaMapping_StatusGenerator() gopter.Gen {
	if inputSchemaMapping_statusGenerator != nil {
		return inputSchemaMapping_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInputSchemaMapping_Status(generators)
	inputSchemaMapping_statusGenerator = gen.Struct(reflect.TypeOf(InputSchemaMapping_Status{}), generators)

	return inputSchemaMapping_statusGenerator
}

// AddIndependentPropertyGeneratorsForInputSchemaMapping_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForInputSchemaMapping_Status(gens map[string]gopter.Gen) {
	gens["InputSchemaMappingType"] = gen.AlphaString()
}

func Test_PrivateEndpointConnection_Status_Domain_SubResourceEmbedded_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateEndpointConnection_Status_Domain_SubResourceEmbedded to PrivateEndpointConnection_Status_Domain_SubResourceEmbedded via AssignPropertiesToPrivateEndpointConnection_Status_Domain_SubResourceEmbedded & AssignPropertiesFromPrivateEndpointConnection_Status_Domain_SubResourceEmbedded returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateEndpointConnection_Status_Domain_SubResourceEmbedded, PrivateEndpointConnection_Status_Domain_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateEndpointConnection_Status_Domain_SubResourceEmbedded tests if a specific instance of PrivateEndpointConnection_Status_Domain_SubResourceEmbedded can be assigned to v1alpha1api20200601storage and back losslessly
func RunPropertyAssignmentTestForPrivateEndpointConnection_Status_Domain_SubResourceEmbedded(subject PrivateEndpointConnection_Status_Domain_SubResourceEmbedded) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20200601storage.PrivateEndpointConnection_Status_Domain_SubResourceEmbedded
	err := copied.AssignPropertiesToPrivateEndpointConnection_Status_Domain_SubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateEndpointConnection_Status_Domain_SubResourceEmbedded
	err = actual.AssignPropertiesFromPrivateEndpointConnection_Status_Domain_SubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_PrivateEndpointConnection_Status_Domain_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_Status_Domain_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnection_Status_Domain_SubResourceEmbedded, PrivateEndpointConnection_Status_Domain_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnection_Status_Domain_SubResourceEmbedded runs a test to see if a specific instance of PrivateEndpointConnection_Status_Domain_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnection_Status_Domain_SubResourceEmbedded(subject PrivateEndpointConnection_Status_Domain_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnection_Status_Domain_SubResourceEmbedded
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

// Generator of PrivateEndpointConnection_Status_Domain_SubResourceEmbedded instances for property testing - lazily
//instantiated by PrivateEndpointConnection_Status_Domain_SubResourceEmbeddedGenerator()
var privateEndpointConnection_status_domain_subResourceEmbeddedGenerator gopter.Gen

// PrivateEndpointConnection_Status_Domain_SubResourceEmbeddedGenerator returns a generator of PrivateEndpointConnection_Status_Domain_SubResourceEmbedded instances for property testing.
func PrivateEndpointConnection_Status_Domain_SubResourceEmbeddedGenerator() gopter.Gen {
	if privateEndpointConnection_status_domain_subResourceEmbeddedGenerator != nil {
		return privateEndpointConnection_status_domain_subResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnection_Status_Domain_SubResourceEmbedded(generators)
	privateEndpointConnection_status_domain_subResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_Status_Domain_SubResourceEmbedded{}), generators)

	return privateEndpointConnection_status_domain_subResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnection_Status_Domain_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnection_Status_Domain_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_SystemData_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SystemData_Status to SystemData_Status via AssignPropertiesToSystemData_Status & AssignPropertiesFromSystemData_Status returns original",
		prop.ForAll(RunPropertyAssignmentTestForSystemData_Status, SystemData_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSystemData_Status tests if a specific instance of SystemData_Status can be assigned to v1alpha1api20200601storage and back losslessly
func RunPropertyAssignmentTestForSystemData_Status(subject SystemData_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20200601storage.SystemData_Status
	err := copied.AssignPropertiesToSystemData_Status(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SystemData_Status
	err = actual.AssignPropertiesFromSystemData_Status(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SystemData_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemData_Status, SystemData_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemData_Status runs a test to see if a specific instance of SystemData_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemData_Status(subject SystemData_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_Status
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

// Generator of SystemData_Status instances for property testing - lazily instantiated by SystemData_StatusGenerator()
var systemData_statusGenerator gopter.Gen

// SystemData_StatusGenerator returns a generator of SystemData_Status instances for property testing.
func SystemData_StatusGenerator() gopter.Gen {
	if systemData_statusGenerator != nil {
		return systemData_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemData_Status(generators)
	systemData_statusGenerator = gen.Struct(reflect.TypeOf(SystemData_Status{}), generators)

	return systemData_statusGenerator
}

// AddIndependentPropertyGeneratorsForSystemData_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemData_Status(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.AlphaString())
}

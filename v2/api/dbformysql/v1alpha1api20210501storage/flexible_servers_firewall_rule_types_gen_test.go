// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210501storage

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

func Test_FlexibleServersFirewallRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersFirewallRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersFirewallRule, FlexibleServersFirewallRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersFirewallRule runs a test to see if a specific instance of FlexibleServersFirewallRule round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersFirewallRule(subject FlexibleServersFirewallRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersFirewallRule
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

// Generator of FlexibleServersFirewallRule instances for property testing - lazily instantiated by
//FlexibleServersFirewallRuleGenerator()
var flexibleServersFirewallRuleGenerator gopter.Gen

// FlexibleServersFirewallRuleGenerator returns a generator of FlexibleServersFirewallRule instances for property testing.
func FlexibleServersFirewallRuleGenerator() gopter.Gen {
	if flexibleServersFirewallRuleGenerator != nil {
		return flexibleServersFirewallRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForFlexibleServersFirewallRule(generators)
	flexibleServersFirewallRuleGenerator = gen.Struct(reflect.TypeOf(FlexibleServersFirewallRule{}), generators)

	return flexibleServersFirewallRuleGenerator
}

// AddRelatedPropertyGeneratorsForFlexibleServersFirewallRule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersFirewallRule(gens map[string]gopter.Gen) {
	gens["Spec"] = FlexibleServersFirewallRules_SPECGenerator()
	gens["Status"] = FirewallRule_StatusGenerator()
}

func Test_FirewallRule_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FirewallRule_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFirewallRule_Status, FirewallRule_StatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFirewallRule_Status runs a test to see if a specific instance of FirewallRule_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForFirewallRule_Status(subject FirewallRule_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FirewallRule_Status
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

// Generator of FirewallRule_Status instances for property testing - lazily instantiated by
//FirewallRule_StatusGenerator()
var firewallRule_statusGenerator gopter.Gen

// FirewallRule_StatusGenerator returns a generator of FirewallRule_Status instances for property testing.
// We first initialize firewallRule_statusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FirewallRule_StatusGenerator() gopter.Gen {
	if firewallRule_statusGenerator != nil {
		return firewallRule_statusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFirewallRule_Status(generators)
	firewallRule_statusGenerator = gen.Struct(reflect.TypeOf(FirewallRule_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFirewallRule_Status(generators)
	AddRelatedPropertyGeneratorsForFirewallRule_Status(generators)
	firewallRule_statusGenerator = gen.Struct(reflect.TypeOf(FirewallRule_Status{}), generators)

	return firewallRule_statusGenerator
}

// AddIndependentPropertyGeneratorsForFirewallRule_Status is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFirewallRule_Status(gens map[string]gopter.Gen) {
	gens["EndIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["StartIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFirewallRule_Status is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFirewallRule_Status(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_StatusGenerator())
}

func Test_FlexibleServersFirewallRules_SPEC_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersFirewallRules_SPEC via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersFirewallRules_SPEC, FlexibleServersFirewallRules_SPECGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersFirewallRules_SPEC runs a test to see if a specific instance of FlexibleServersFirewallRules_SPEC round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersFirewallRules_SPEC(subject FlexibleServersFirewallRules_SPEC) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersFirewallRules_SPEC
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

// Generator of FlexibleServersFirewallRules_SPEC instances for property testing - lazily instantiated by
//FlexibleServersFirewallRules_SPECGenerator()
var flexibleServersFirewallRules_specGenerator gopter.Gen

// FlexibleServersFirewallRules_SPECGenerator returns a generator of FlexibleServersFirewallRules_SPEC instances for property testing.
func FlexibleServersFirewallRules_SPECGenerator() gopter.Gen {
	if flexibleServersFirewallRules_specGenerator != nil {
		return flexibleServersFirewallRules_specGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersFirewallRules_SPEC(generators)
	flexibleServersFirewallRules_specGenerator = gen.Struct(reflect.TypeOf(FlexibleServersFirewallRules_SPEC{}), generators)

	return flexibleServersFirewallRules_specGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersFirewallRules_SPEC is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersFirewallRules_SPEC(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["EndIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["StartIpAddress"] = gen.PtrOf(gen.AlphaString())
}

// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210501

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

func Test_FlexibleServersFirewallRule_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersFirewallRule_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersFirewallRule_SpecARM, FlexibleServersFirewallRule_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersFirewallRule_SpecARM runs a test to see if a specific instance of FlexibleServersFirewallRule_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersFirewallRule_SpecARM(subject FlexibleServersFirewallRule_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersFirewallRule_SpecARM
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

// Generator of FlexibleServersFirewallRule_SpecARM instances for property testing - lazily instantiated by
//FlexibleServersFirewallRule_SpecARMGenerator()
var flexibleServersFirewallRule_specARMGenerator gopter.Gen

// FlexibleServersFirewallRule_SpecARMGenerator returns a generator of FlexibleServersFirewallRule_SpecARM instances for property testing.
// We first initialize flexibleServersFirewallRule_specARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServersFirewallRule_SpecARMGenerator() gopter.Gen {
	if flexibleServersFirewallRule_specARMGenerator != nil {
		return flexibleServersFirewallRule_specARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersFirewallRule_SpecARM(generators)
	flexibleServersFirewallRule_specARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServersFirewallRule_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersFirewallRule_SpecARM(generators)
	AddRelatedPropertyGeneratorsForFlexibleServersFirewallRule_SpecARM(generators)
	flexibleServersFirewallRule_specARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServersFirewallRule_SpecARM{}), generators)

	return flexibleServersFirewallRule_specARMGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersFirewallRule_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersFirewallRule_SpecARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForFlexibleServersFirewallRule_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersFirewallRule_SpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = FirewallRulePropertiesARMGenerator()
}

func Test_FirewallRulePropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FirewallRulePropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFirewallRulePropertiesARM, FirewallRulePropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFirewallRulePropertiesARM runs a test to see if a specific instance of FirewallRulePropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFirewallRulePropertiesARM(subject FirewallRulePropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FirewallRulePropertiesARM
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

// Generator of FirewallRulePropertiesARM instances for property testing - lazily instantiated by
//FirewallRulePropertiesARMGenerator()
var firewallRulePropertiesARMGenerator gopter.Gen

// FirewallRulePropertiesARMGenerator returns a generator of FirewallRulePropertiesARM instances for property testing.
func FirewallRulePropertiesARMGenerator() gopter.Gen {
	if firewallRulePropertiesARMGenerator != nil {
		return firewallRulePropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFirewallRulePropertiesARM(generators)
	firewallRulePropertiesARMGenerator = gen.Struct(reflect.TypeOf(FirewallRulePropertiesARM{}), generators)

	return firewallRulePropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForFirewallRulePropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFirewallRulePropertiesARM(gens map[string]gopter.Gen) {
	gens["EndIpAddress"] = gen.AlphaString()
	gens["StartIpAddress"] = gen.AlphaString()
}

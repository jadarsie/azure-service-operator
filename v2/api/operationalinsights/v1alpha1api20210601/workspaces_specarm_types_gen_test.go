// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601

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

func Test_Workspaces_SPECARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Workspaces_SPECARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaces_SPECARM, Workspaces_SPECARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaces_SPECARM runs a test to see if a specific instance of Workspaces_SPECARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaces_SPECARM(subject Workspaces_SPECARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Workspaces_SPECARM
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

// Generator of Workspaces_SPECARM instances for property testing - lazily instantiated by Workspaces_SPECARMGenerator()
var workspaces_specarmGenerator gopter.Gen

// Workspaces_SPECARMGenerator returns a generator of Workspaces_SPECARM instances for property testing.
// We first initialize workspaces_specarmGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Workspaces_SPECARMGenerator() gopter.Gen {
	if workspaces_specarmGenerator != nil {
		return workspaces_specarmGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaces_SPECARM(generators)
	workspaces_specarmGenerator = gen.Struct(reflect.TypeOf(Workspaces_SPECARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaces_SPECARM(generators)
	AddRelatedPropertyGeneratorsForWorkspaces_SPECARM(generators)
	workspaces_specarmGenerator = gen.Struct(reflect.TypeOf(Workspaces_SPECARM{}), generators)

	return workspaces_specarmGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaces_SPECARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaces_SPECARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["ETag"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.AlphaString()
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForWorkspaces_SPECARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspaces_SPECARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(WorkspaceProperties_SpecARMGenerator())
}

func Test_WorkspaceProperties_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceProperties_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceProperties_SpecARM, WorkspaceProperties_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceProperties_SpecARM runs a test to see if a specific instance of WorkspaceProperties_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceProperties_SpecARM(subject WorkspaceProperties_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceProperties_SpecARM
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

// Generator of WorkspaceProperties_SpecARM instances for property testing - lazily instantiated by
//WorkspaceProperties_SpecARMGenerator()
var workspaceProperties_specARMGenerator gopter.Gen

// WorkspaceProperties_SpecARMGenerator returns a generator of WorkspaceProperties_SpecARM instances for property testing.
// We first initialize workspaceProperties_specARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WorkspaceProperties_SpecARMGenerator() gopter.Gen {
	if workspaceProperties_specARMGenerator != nil {
		return workspaceProperties_specARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceProperties_SpecARM(generators)
	workspaceProperties_specARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceProperties_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceProperties_SpecARM(generators)
	AddRelatedPropertyGeneratorsForWorkspaceProperties_SpecARM(generators)
	workspaceProperties_specARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceProperties_SpecARM{}), generators)

	return workspaceProperties_specARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceProperties_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceProperties_SpecARM(gens map[string]gopter.Gen) {
	gens["ForceCmkForQuery"] = gen.PtrOf(gen.Bool())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		WorkspaceProperties_ProvisioningState_SpecCanceled,
		WorkspaceProperties_ProvisioningState_SpecCreating,
		WorkspaceProperties_ProvisioningState_SpecDeleting,
		WorkspaceProperties_ProvisioningState_SpecFailed,
		WorkspaceProperties_ProvisioningState_SpecProvisioningAccount,
		WorkspaceProperties_ProvisioningState_SpecSucceeded,
		WorkspaceProperties_ProvisioningState_SpecUpdating))
	gens["PublicNetworkAccessForIngestion"] = gen.PtrOf(gen.OneConstOf(PublicNetworkAccessType_SpecDisabled, PublicNetworkAccessType_SpecEnabled))
	gens["PublicNetworkAccessForQuery"] = gen.PtrOf(gen.OneConstOf(PublicNetworkAccessType_SpecDisabled, PublicNetworkAccessType_SpecEnabled))
	gens["RetentionInDays"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForWorkspaceProperties_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspaceProperties_SpecARM(gens map[string]gopter.Gen) {
	gens["Features"] = gen.PtrOf(WorkspaceFeatures_SpecARMGenerator())
	gens["Sku"] = gen.PtrOf(WorkspaceSku_SpecARMGenerator())
	gens["WorkspaceCapping"] = gen.PtrOf(WorkspaceCapping_SpecARMGenerator())
}

func Test_WorkspaceCapping_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceCapping_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceCapping_SpecARM, WorkspaceCapping_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceCapping_SpecARM runs a test to see if a specific instance of WorkspaceCapping_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceCapping_SpecARM(subject WorkspaceCapping_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceCapping_SpecARM
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

// Generator of WorkspaceCapping_SpecARM instances for property testing - lazily instantiated by
//WorkspaceCapping_SpecARMGenerator()
var workspaceCapping_specARMGenerator gopter.Gen

// WorkspaceCapping_SpecARMGenerator returns a generator of WorkspaceCapping_SpecARM instances for property testing.
func WorkspaceCapping_SpecARMGenerator() gopter.Gen {
	if workspaceCapping_specARMGenerator != nil {
		return workspaceCapping_specARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceCapping_SpecARM(generators)
	workspaceCapping_specARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceCapping_SpecARM{}), generators)

	return workspaceCapping_specARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceCapping_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceCapping_SpecARM(gens map[string]gopter.Gen) {
	gens["DailyQuotaGb"] = gen.PtrOf(gen.Float64())
}

func Test_WorkspaceFeatures_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceFeatures_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceFeatures_SpecARM, WorkspaceFeatures_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceFeatures_SpecARM runs a test to see if a specific instance of WorkspaceFeatures_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceFeatures_SpecARM(subject WorkspaceFeatures_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceFeatures_SpecARM
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

// Generator of WorkspaceFeatures_SpecARM instances for property testing - lazily instantiated by
//WorkspaceFeatures_SpecARMGenerator()
var workspaceFeatures_specARMGenerator gopter.Gen

// WorkspaceFeatures_SpecARMGenerator returns a generator of WorkspaceFeatures_SpecARM instances for property testing.
func WorkspaceFeatures_SpecARMGenerator() gopter.Gen {
	if workspaceFeatures_specARMGenerator != nil {
		return workspaceFeatures_specARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceFeatures_SpecARM(generators)
	workspaceFeatures_specARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceFeatures_SpecARM{}), generators)

	return workspaceFeatures_specARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceFeatures_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceFeatures_SpecARM(gens map[string]gopter.Gen) {
	gens["ClusterResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["DisableLocalAuth"] = gen.PtrOf(gen.Bool())
	gens["EnableDataExport"] = gen.PtrOf(gen.Bool())
	gens["EnableLogAccessUsingOnlyResourcePermissions"] = gen.PtrOf(gen.Bool())
	gens["ImmediatePurgeDataOn30Days"] = gen.PtrOf(gen.Bool())
}

func Test_WorkspaceSku_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceSku_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceSku_SpecARM, WorkspaceSku_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceSku_SpecARM runs a test to see if a specific instance of WorkspaceSku_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceSku_SpecARM(subject WorkspaceSku_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceSku_SpecARM
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

// Generator of WorkspaceSku_SpecARM instances for property testing - lazily instantiated by
//WorkspaceSku_SpecARMGenerator()
var workspaceSku_specARMGenerator gopter.Gen

// WorkspaceSku_SpecARMGenerator returns a generator of WorkspaceSku_SpecARM instances for property testing.
func WorkspaceSku_SpecARMGenerator() gopter.Gen {
	if workspaceSku_specARMGenerator != nil {
		return workspaceSku_specARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceSku_SpecARM(generators)
	workspaceSku_specARMGenerator = gen.Struct(reflect.TypeOf(WorkspaceSku_SpecARM{}), generators)

	return workspaceSku_specARMGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceSku_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceSku_SpecARM(gens map[string]gopter.Gen) {
	gens["CapacityReservationLevel"] = gen.PtrOf(gen.OneConstOf(
		WorkspaceSku_CapacityReservationLevel_Spec100,
		WorkspaceSku_CapacityReservationLevel_Spec1000,
		WorkspaceSku_CapacityReservationLevel_Spec200,
		WorkspaceSku_CapacityReservationLevel_Spec2000,
		WorkspaceSku_CapacityReservationLevel_Spec300,
		WorkspaceSku_CapacityReservationLevel_Spec400,
		WorkspaceSku_CapacityReservationLevel_Spec500,
		WorkspaceSku_CapacityReservationLevel_Spec5000))
	gens["Name"] = gen.OneConstOf(
		WorkspaceSku_Name_SpecCapacityReservation,
		WorkspaceSku_Name_SpecFree,
		WorkspaceSku_Name_SpecLACluster,
		WorkspaceSku_Name_SpecPerGB2018,
		WorkspaceSku_Name_SpecPerNode,
		WorkspaceSku_Name_SpecPremium,
		WorkspaceSku_Name_SpecStandalone,
		WorkspaceSku_Name_SpecStandard)
}

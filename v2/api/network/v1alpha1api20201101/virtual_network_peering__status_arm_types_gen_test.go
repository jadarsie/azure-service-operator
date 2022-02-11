// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

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

func Test_VirtualNetworkPeering_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkPeering_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkPeering_StatusARM, VirtualNetworkPeering_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkPeering_StatusARM runs a test to see if a specific instance of VirtualNetworkPeering_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkPeering_StatusARM(subject VirtualNetworkPeering_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkPeering_StatusARM
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

// Generator of VirtualNetworkPeering_StatusARM instances for property testing - lazily instantiated by
//VirtualNetworkPeering_StatusARMGenerator()
var virtualNetworkPeering_statusARMGenerator gopter.Gen

// VirtualNetworkPeering_StatusARMGenerator returns a generator of VirtualNetworkPeering_StatusARM instances for property testing.
// We first initialize virtualNetworkPeering_statusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworkPeering_StatusARMGenerator() gopter.Gen {
	if virtualNetworkPeering_statusARMGenerator != nil {
		return virtualNetworkPeering_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPeering_StatusARM(generators)
	virtualNetworkPeering_statusARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPeering_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPeering_StatusARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworkPeering_StatusARM(generators)
	virtualNetworkPeering_statusARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPeering_StatusARM{}), generators)

	return virtualNetworkPeering_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkPeering_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkPeering_StatusARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetworkPeering_StatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworkPeering_StatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(VirtualNetworkPeeringPropertiesFormat_StatusARMGenerator())
}

func Test_VirtualNetworkPeeringPropertiesFormat_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkPeeringPropertiesFormat_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkPeeringPropertiesFormat_StatusARM, VirtualNetworkPeeringPropertiesFormat_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkPeeringPropertiesFormat_StatusARM runs a test to see if a specific instance of VirtualNetworkPeeringPropertiesFormat_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkPeeringPropertiesFormat_StatusARM(subject VirtualNetworkPeeringPropertiesFormat_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkPeeringPropertiesFormat_StatusARM
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

// Generator of VirtualNetworkPeeringPropertiesFormat_StatusARM instances for property testing - lazily instantiated by
//VirtualNetworkPeeringPropertiesFormat_StatusARMGenerator()
var virtualNetworkPeeringPropertiesFormat_statusARMGenerator gopter.Gen

// VirtualNetworkPeeringPropertiesFormat_StatusARMGenerator returns a generator of VirtualNetworkPeeringPropertiesFormat_StatusARM instances for property testing.
// We first initialize virtualNetworkPeeringPropertiesFormat_statusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworkPeeringPropertiesFormat_StatusARMGenerator() gopter.Gen {
	if virtualNetworkPeeringPropertiesFormat_statusARMGenerator != nil {
		return virtualNetworkPeeringPropertiesFormat_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormat_StatusARM(generators)
	virtualNetworkPeeringPropertiesFormat_statusARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPeeringPropertiesFormat_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormat_StatusARM(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormat_StatusARM(generators)
	virtualNetworkPeeringPropertiesFormat_statusARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkPeeringPropertiesFormat_StatusARM{}), generators)

	return virtualNetworkPeeringPropertiesFormat_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormat_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormat_StatusARM(gens map[string]gopter.Gen) {
	gens["AllowForwardedTraffic"] = gen.PtrOf(gen.Bool())
	gens["AllowGatewayTransit"] = gen.PtrOf(gen.Bool())
	gens["AllowVirtualNetworkAccess"] = gen.PtrOf(gen.Bool())
	gens["DoNotVerifyRemoteGateways"] = gen.PtrOf(gen.Bool())
	gens["PeeringState"] = gen.PtrOf(gen.OneConstOf(VirtualNetworkPeeringPropertiesFormat_PeeringState_StatusConnected, VirtualNetworkPeeringPropertiesFormat_PeeringState_StatusDisconnected, VirtualNetworkPeeringPropertiesFormat_PeeringState_StatusInitiated))
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_StatusDeleting,
		ProvisioningState_StatusFailed,
		ProvisioningState_StatusSucceeded,
		ProvisioningState_StatusUpdating))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
	gens["UseRemoteGateways"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormat_StatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworkPeeringPropertiesFormat_StatusARM(gens map[string]gopter.Gen) {
	gens["RemoteAddressSpace"] = gen.PtrOf(AddressSpace_StatusARMGenerator())
	gens["RemoteBgpCommunities"] = gen.PtrOf(VirtualNetworkBgpCommunities_StatusARMGenerator())
	gens["RemoteVirtualNetwork"] = gen.PtrOf(SubResource_StatusARMGenerator())
}

func Test_VirtualNetworkBgpCommunities_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkBgpCommunities_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkBgpCommunities_StatusARM, VirtualNetworkBgpCommunities_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkBgpCommunities_StatusARM runs a test to see if a specific instance of VirtualNetworkBgpCommunities_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkBgpCommunities_StatusARM(subject VirtualNetworkBgpCommunities_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkBgpCommunities_StatusARM
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

// Generator of VirtualNetworkBgpCommunities_StatusARM instances for property testing - lazily instantiated by
//VirtualNetworkBgpCommunities_StatusARMGenerator()
var virtualNetworkBgpCommunities_statusARMGenerator gopter.Gen

// VirtualNetworkBgpCommunities_StatusARMGenerator returns a generator of VirtualNetworkBgpCommunities_StatusARM instances for property testing.
func VirtualNetworkBgpCommunities_StatusARMGenerator() gopter.Gen {
	if virtualNetworkBgpCommunities_statusARMGenerator != nil {
		return virtualNetworkBgpCommunities_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunities_StatusARM(generators)
	virtualNetworkBgpCommunities_statusARMGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkBgpCommunities_StatusARM{}), generators)

	return virtualNetworkBgpCommunities_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunities_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunities_StatusARM(gens map[string]gopter.Gen) {
	gens["RegionalCommunity"] = gen.PtrOf(gen.AlphaString())
	gens["VirtualNetworkCommunity"] = gen.AlphaString()
}

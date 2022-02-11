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

func Test_AgentPool_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AgentPool_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAgentPool_StatusARM, AgentPool_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAgentPool_StatusARM runs a test to see if a specific instance of AgentPool_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAgentPool_StatusARM(subject AgentPool_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AgentPool_StatusARM
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

// Generator of AgentPool_StatusARM instances for property testing - lazily instantiated by
//AgentPool_StatusARMGenerator()
var agentPool_statusARMGenerator gopter.Gen

// AgentPool_StatusARMGenerator returns a generator of AgentPool_StatusARM instances for property testing.
// We first initialize agentPool_statusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AgentPool_StatusARMGenerator() gopter.Gen {
	if agentPool_statusARMGenerator != nil {
		return agentPool_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAgentPool_StatusARM(generators)
	agentPool_statusARMGenerator = gen.Struct(reflect.TypeOf(AgentPool_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAgentPool_StatusARM(generators)
	AddRelatedPropertyGeneratorsForAgentPool_StatusARM(generators)
	agentPool_statusARMGenerator = gen.Struct(reflect.TypeOf(AgentPool_StatusARM{}), generators)

	return agentPool_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForAgentPool_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAgentPool_StatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForAgentPool_StatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAgentPool_StatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ManagedClusterAgentPoolProfileProperties_StatusARMGenerator())
}

func Test_ManagedClusterAgentPoolProfileProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedClusterAgentPoolProfileProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedClusterAgentPoolProfileProperties_StatusARM, ManagedClusterAgentPoolProfileProperties_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedClusterAgentPoolProfileProperties_StatusARM runs a test to see if a specific instance of ManagedClusterAgentPoolProfileProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedClusterAgentPoolProfileProperties_StatusARM(subject ManagedClusterAgentPoolProfileProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedClusterAgentPoolProfileProperties_StatusARM
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

// Generator of ManagedClusterAgentPoolProfileProperties_StatusARM instances for property testing - lazily instantiated
//by ManagedClusterAgentPoolProfileProperties_StatusARMGenerator()
var managedClusterAgentPoolProfileProperties_statusARMGenerator gopter.Gen

// ManagedClusterAgentPoolProfileProperties_StatusARMGenerator returns a generator of ManagedClusterAgentPoolProfileProperties_StatusARM instances for property testing.
// We first initialize managedClusterAgentPoolProfileProperties_statusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagedClusterAgentPoolProfileProperties_StatusARMGenerator() gopter.Gen {
	if managedClusterAgentPoolProfileProperties_statusARMGenerator != nil {
		return managedClusterAgentPoolProfileProperties_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClusterAgentPoolProfileProperties_StatusARM(generators)
	managedClusterAgentPoolProfileProperties_statusARMGenerator = gen.Struct(reflect.TypeOf(ManagedClusterAgentPoolProfileProperties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClusterAgentPoolProfileProperties_StatusARM(generators)
	AddRelatedPropertyGeneratorsForManagedClusterAgentPoolProfileProperties_StatusARM(generators)
	managedClusterAgentPoolProfileProperties_statusARMGenerator = gen.Struct(reflect.TypeOf(ManagedClusterAgentPoolProfileProperties_StatusARM{}), generators)

	return managedClusterAgentPoolProfileProperties_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForManagedClusterAgentPoolProfileProperties_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedClusterAgentPoolProfileProperties_StatusARM(gens map[string]gopter.Gen) {
	gens["AvailabilityZones"] = gen.SliceOf(gen.AlphaString())
	gens["Count"] = gen.PtrOf(gen.Int())
	gens["EnableAutoScaling"] = gen.PtrOf(gen.Bool())
	gens["EnableEncryptionAtHost"] = gen.PtrOf(gen.Bool())
	gens["EnableFIPS"] = gen.PtrOf(gen.Bool())
	gens["EnableNodePublicIP"] = gen.PtrOf(gen.Bool())
	gens["EnableUltraSSD"] = gen.PtrOf(gen.Bool())
	gens["GpuInstanceProfile"] = gen.PtrOf(gen.OneConstOf(
		GPUInstanceProfile_StatusMIG1G,
		GPUInstanceProfile_StatusMIG2G,
		GPUInstanceProfile_StatusMIG3G,
		GPUInstanceProfile_StatusMIG4G,
		GPUInstanceProfile_StatusMIG7G))
	gens["KubeletDiskType"] = gen.PtrOf(gen.OneConstOf(KubeletDiskType_StatusOS, KubeletDiskType_StatusTemporary))
	gens["MaxCount"] = gen.PtrOf(gen.Int())
	gens["MaxPods"] = gen.PtrOf(gen.Int())
	gens["MinCount"] = gen.PtrOf(gen.Int())
	gens["Mode"] = gen.PtrOf(gen.OneConstOf(AgentPoolMode_StatusSystem, AgentPoolMode_StatusUser))
	gens["NodeImageVersion"] = gen.PtrOf(gen.AlphaString())
	gens["NodeLabels"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["NodePublicIPPrefixID"] = gen.PtrOf(gen.AlphaString())
	gens["NodeTaints"] = gen.SliceOf(gen.AlphaString())
	gens["OrchestratorVersion"] = gen.PtrOf(gen.AlphaString())
	gens["OsDiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["OsDiskType"] = gen.PtrOf(gen.OneConstOf(OSDiskType_StatusEphemeral, OSDiskType_StatusManaged))
	gens["OsSKU"] = gen.PtrOf(gen.OneConstOf(OSSKU_StatusCBLMariner, OSSKU_StatusUbuntu))
	gens["OsType"] = gen.PtrOf(gen.OneConstOf(OSType_StatusLinux, OSType_StatusWindows))
	gens["PodSubnetID"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["ProximityPlacementGroupID"] = gen.PtrOf(gen.AlphaString())
	gens["ScaleSetEvictionPolicy"] = gen.PtrOf(gen.OneConstOf(ScaleSetEvictionPolicy_StatusDeallocate, ScaleSetEvictionPolicy_StatusDelete))
	gens["ScaleSetPriority"] = gen.PtrOf(gen.OneConstOf(ScaleSetPriority_StatusRegular, ScaleSetPriority_StatusSpot))
	gens["SpotMaxPrice"] = gen.PtrOf(gen.Float64())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(AgentPoolType_StatusAvailabilitySet, AgentPoolType_StatusVirtualMachineScaleSets))
	gens["VmSize"] = gen.PtrOf(gen.AlphaString())
	gens["VnetSubnetID"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForManagedClusterAgentPoolProfileProperties_StatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagedClusterAgentPoolProfileProperties_StatusARM(gens map[string]gopter.Gen) {
	gens["KubeletConfig"] = gen.PtrOf(KubeletConfig_StatusARMGenerator())
	gens["LinuxOSConfig"] = gen.PtrOf(LinuxOSConfig_StatusARMGenerator())
	gens["PowerState"] = gen.PtrOf(PowerState_StatusARMGenerator())
	gens["UpgradeSettings"] = gen.PtrOf(AgentPoolUpgradeSettings_StatusARMGenerator())
}

func Test_AgentPoolUpgradeSettings_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AgentPoolUpgradeSettings_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAgentPoolUpgradeSettings_StatusARM, AgentPoolUpgradeSettings_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAgentPoolUpgradeSettings_StatusARM runs a test to see if a specific instance of AgentPoolUpgradeSettings_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAgentPoolUpgradeSettings_StatusARM(subject AgentPoolUpgradeSettings_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AgentPoolUpgradeSettings_StatusARM
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

// Generator of AgentPoolUpgradeSettings_StatusARM instances for property testing - lazily instantiated by
//AgentPoolUpgradeSettings_StatusARMGenerator()
var agentPoolUpgradeSettings_statusARMGenerator gopter.Gen

// AgentPoolUpgradeSettings_StatusARMGenerator returns a generator of AgentPoolUpgradeSettings_StatusARM instances for property testing.
func AgentPoolUpgradeSettings_StatusARMGenerator() gopter.Gen {
	if agentPoolUpgradeSettings_statusARMGenerator != nil {
		return agentPoolUpgradeSettings_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettings_StatusARM(generators)
	agentPoolUpgradeSettings_statusARMGenerator = gen.Struct(reflect.TypeOf(AgentPoolUpgradeSettings_StatusARM{}), generators)

	return agentPoolUpgradeSettings_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettings_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettings_StatusARM(gens map[string]gopter.Gen) {
	gens["MaxSurge"] = gen.PtrOf(gen.AlphaString())
}

func Test_KubeletConfig_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KubeletConfig_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKubeletConfig_StatusARM, KubeletConfig_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKubeletConfig_StatusARM runs a test to see if a specific instance of KubeletConfig_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForKubeletConfig_StatusARM(subject KubeletConfig_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KubeletConfig_StatusARM
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

// Generator of KubeletConfig_StatusARM instances for property testing - lazily instantiated by
//KubeletConfig_StatusARMGenerator()
var kubeletConfig_statusARMGenerator gopter.Gen

// KubeletConfig_StatusARMGenerator returns a generator of KubeletConfig_StatusARM instances for property testing.
func KubeletConfig_StatusARMGenerator() gopter.Gen {
	if kubeletConfig_statusARMGenerator != nil {
		return kubeletConfig_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKubeletConfig_StatusARM(generators)
	kubeletConfig_statusARMGenerator = gen.Struct(reflect.TypeOf(KubeletConfig_StatusARM{}), generators)

	return kubeletConfig_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForKubeletConfig_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKubeletConfig_StatusARM(gens map[string]gopter.Gen) {
	gens["AllowedUnsafeSysctls"] = gen.SliceOf(gen.AlphaString())
	gens["ContainerLogMaxFiles"] = gen.PtrOf(gen.Int())
	gens["ContainerLogMaxSizeMB"] = gen.PtrOf(gen.Int())
	gens["CpuCfsQuota"] = gen.PtrOf(gen.Bool())
	gens["CpuCfsQuotaPeriod"] = gen.PtrOf(gen.AlphaString())
	gens["CpuManagerPolicy"] = gen.PtrOf(gen.AlphaString())
	gens["FailSwapOn"] = gen.PtrOf(gen.Bool())
	gens["ImageGcHighThreshold"] = gen.PtrOf(gen.Int())
	gens["ImageGcLowThreshold"] = gen.PtrOf(gen.Int())
	gens["PodMaxPids"] = gen.PtrOf(gen.Int())
	gens["TopologyManagerPolicy"] = gen.PtrOf(gen.AlphaString())
}

func Test_LinuxOSConfig_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LinuxOSConfig_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLinuxOSConfig_StatusARM, LinuxOSConfig_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLinuxOSConfig_StatusARM runs a test to see if a specific instance of LinuxOSConfig_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForLinuxOSConfig_StatusARM(subject LinuxOSConfig_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LinuxOSConfig_StatusARM
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

// Generator of LinuxOSConfig_StatusARM instances for property testing - lazily instantiated by
//LinuxOSConfig_StatusARMGenerator()
var linuxOSConfig_statusARMGenerator gopter.Gen

// LinuxOSConfig_StatusARMGenerator returns a generator of LinuxOSConfig_StatusARM instances for property testing.
// We first initialize linuxOSConfig_statusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func LinuxOSConfig_StatusARMGenerator() gopter.Gen {
	if linuxOSConfig_statusARMGenerator != nil {
		return linuxOSConfig_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLinuxOSConfig_StatusARM(generators)
	linuxOSConfig_statusARMGenerator = gen.Struct(reflect.TypeOf(LinuxOSConfig_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLinuxOSConfig_StatusARM(generators)
	AddRelatedPropertyGeneratorsForLinuxOSConfig_StatusARM(generators)
	linuxOSConfig_statusARMGenerator = gen.Struct(reflect.TypeOf(LinuxOSConfig_StatusARM{}), generators)

	return linuxOSConfig_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForLinuxOSConfig_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLinuxOSConfig_StatusARM(gens map[string]gopter.Gen) {
	gens["SwapFileSizeMB"] = gen.PtrOf(gen.Int())
	gens["TransparentHugePageDefrag"] = gen.PtrOf(gen.AlphaString())
	gens["TransparentHugePageEnabled"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForLinuxOSConfig_StatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForLinuxOSConfig_StatusARM(gens map[string]gopter.Gen) {
	gens["Sysctls"] = gen.PtrOf(SysctlConfig_StatusARMGenerator())
}

func Test_PowerState_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PowerState_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPowerState_StatusARM, PowerState_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPowerState_StatusARM runs a test to see if a specific instance of PowerState_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPowerState_StatusARM(subject PowerState_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PowerState_StatusARM
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

// Generator of PowerState_StatusARM instances for property testing - lazily instantiated by
//PowerState_StatusARMGenerator()
var powerState_statusARMGenerator gopter.Gen

// PowerState_StatusARMGenerator returns a generator of PowerState_StatusARM instances for property testing.
func PowerState_StatusARMGenerator() gopter.Gen {
	if powerState_statusARMGenerator != nil {
		return powerState_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPowerState_StatusARM(generators)
	powerState_statusARMGenerator = gen.Struct(reflect.TypeOf(PowerState_StatusARM{}), generators)

	return powerState_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForPowerState_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPowerState_StatusARM(gens map[string]gopter.Gen) {
	gens["Code"] = gen.PtrOf(gen.OneConstOf(PowerState_Code_StatusRunning, PowerState_Code_StatusStopped))
}

func Test_SysctlConfig_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SysctlConfig_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSysctlConfig_StatusARM, SysctlConfig_StatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSysctlConfig_StatusARM runs a test to see if a specific instance of SysctlConfig_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSysctlConfig_StatusARM(subject SysctlConfig_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SysctlConfig_StatusARM
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

// Generator of SysctlConfig_StatusARM instances for property testing - lazily instantiated by
//SysctlConfig_StatusARMGenerator()
var sysctlConfig_statusARMGenerator gopter.Gen

// SysctlConfig_StatusARMGenerator returns a generator of SysctlConfig_StatusARM instances for property testing.
func SysctlConfig_StatusARMGenerator() gopter.Gen {
	if sysctlConfig_statusARMGenerator != nil {
		return sysctlConfig_statusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSysctlConfig_StatusARM(generators)
	sysctlConfig_statusARMGenerator = gen.Struct(reflect.TypeOf(SysctlConfig_StatusARM{}), generators)

	return sysctlConfig_statusARMGenerator
}

// AddIndependentPropertyGeneratorsForSysctlConfig_StatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSysctlConfig_StatusARM(gens map[string]gopter.Gen) {
	gens["FsAioMaxNr"] = gen.PtrOf(gen.Int())
	gens["FsFileMax"] = gen.PtrOf(gen.Int())
	gens["FsInotifyMaxUserWatches"] = gen.PtrOf(gen.Int())
	gens["FsNrOpen"] = gen.PtrOf(gen.Int())
	gens["KernelThreadsMax"] = gen.PtrOf(gen.Int())
	gens["NetCoreNetdevMaxBacklog"] = gen.PtrOf(gen.Int())
	gens["NetCoreOptmemMax"] = gen.PtrOf(gen.Int())
	gens["NetCoreRmemDefault"] = gen.PtrOf(gen.Int())
	gens["NetCoreRmemMax"] = gen.PtrOf(gen.Int())
	gens["NetCoreSomaxconn"] = gen.PtrOf(gen.Int())
	gens["NetCoreWmemDefault"] = gen.PtrOf(gen.Int())
	gens["NetCoreWmemMax"] = gen.PtrOf(gen.Int())
	gens["NetIpv4IpLocalPortRange"] = gen.PtrOf(gen.AlphaString())
	gens["NetIpv4NeighDefaultGcThresh1"] = gen.PtrOf(gen.Int())
	gens["NetIpv4NeighDefaultGcThresh2"] = gen.PtrOf(gen.Int())
	gens["NetIpv4NeighDefaultGcThresh3"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpFinTimeout"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpKeepaliveProbes"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpKeepaliveTime"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpMaxSynBacklog"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpMaxTwBuckets"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpTwReuse"] = gen.PtrOf(gen.Bool())
	gens["NetIpv4TcpkeepaliveIntvl"] = gen.PtrOf(gen.Int())
	gens["NetNetfilterNfConntrackBuckets"] = gen.PtrOf(gen.Int())
	gens["NetNetfilterNfConntrackMax"] = gen.PtrOf(gen.Int())
	gens["VmMaxMapCount"] = gen.PtrOf(gen.Int())
	gens["VmSwappiness"] = gen.PtrOf(gen.Int())
	gens["VmVfsCachePressure"] = gen.PtrOf(gen.Int())
}

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

/*

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	authorization "github.com/Azure/azure-service-operator/v2/api/authorization/v1alpha1api20200801preview"
	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1alpha1api20181130"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1alpha1api20210401"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// NOTE: If you're rerecording these tests make sure you use an SP
// with the Owner role - if it only has Contributor it won't be able
// to create the RoleAssignment.

func Test_Authorization_RoleAssignment_OnResourceGroup_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Create a dummy managed identity which we will assign to a role
	mi := &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
		},
	}

	tc.CreateResourceAndWait(mi)
	tc.Expect(mi.Status.TenantId).ToNot(BeNil())
	tc.Expect(mi.Status.PrincipalId).ToNot(BeNil())

	// Now assign that managed identity to a new role
	roleAssignmentGUID, err := tc.Namer.GenerateUUID()
	tc.Expect(err).ToNot(HaveOccurred())

	roleAssignment := &authorization.RoleAssignment{
		ObjectMeta: tc.MakeObjectMetaWithName(roleAssignmentGUID.String()),
		Spec: authorization.RoleAssignment_Spec{
			Location:    &tc.AzureRegion,
			Owner:       tc.AsExtensionOwner(rg),
			PrincipalId: *mi.Status.PrincipalId,
			RoleDefinitionReference: genruntime.ResourceReference{
				ARMID: fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Authorization/roleDefinitions/b24988ac-6180-42a0-ab88-20f7382dd24c", tc.AzureSubscription), // This is contributor
			},
		},
	}

	tc.CreateResourceAndWait(roleAssignment)

	tc.Expect(roleAssignment.Status).ToNot(BeNil())
	armId := *roleAssignment.Status

	tc.DeleteResourceAndWait(roleAssignment)

	// Ensure that the resource group was really deleted in Azure
	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(authorization.APIVersionValue))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func Test_Authorization_RoleAssignment_OnStorageAccount_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Create a dummy managed identity which we will assign to a role
	mi := &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
		},
	}

	tc.CreateResourceAndWait(mi)
	tc.Expect(mi.Status.TenantId).ToNot(BeNil())
	tc.Expect(mi.Status.PrincipalId).ToNot(BeNil())

	// Create a storage account which we will put the role assignment on
	// Custom namer because storage accounts have strict names
	storageNamer := tc.Namer.WithSeparator("")

	// Create a storage account
	accessTier := storage.StorageAccountPropertiesAccessTierHot
	acct := &storage.StorageAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(storageNamer.GenerateName("stor")),
		Spec: storage.StorageAccount_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Kind:     storage.StorageAccount_SpecKindBlobStorage,
			Sku: storage.Sku{
				Name: storage.SkuNameStandard_LRS,
			},
			AccessTier: &accessTier,
		},
	}

	tc.CreateResourceAndWait(acct)

	// Now assign that managed identity to a new role on the storage account
	roleAssignmentGUID, err := tc.Namer.GenerateUUID()
	tc.Expect(err).ToNot(HaveOccurred())

	roleAssignment := &authorization.RoleAssignment{
		ObjectMeta: tc.MakeObjectMetaWithName(roleAssignmentGUID.String()),
		Spec: authorization.RoleAssignment_Spec{
			Location:    &tc.AzureRegion,
			Owner:       tc.AsExtensionOwner(acct),
			PrincipalId: *mi.Status.PrincipalId,
			RoleDefinitionReference: genruntime.ResourceReference{
				ARMID: fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Authorization/roleDefinitions/b24988ac-6180-42a0-ab88-20f7382dd24c", tc.AzureSubscription), // This is contributor
			},
		},
	}

	tc.CreateResourceAndWait(roleAssignment)

	tc.Expect(roleAssignment.Status.Id).ToNot(BeNil())
	armId := *roleAssignment.Status.Id

	// Ensure that we're actually on the storage account - this means the storage account id should be in the roleAssignment ID
	tc.Expect(armId).To(ContainSubstring(acct.AzureName()))

	tc.DeleteResourceAndWait(roleAssignment)

	// Ensure that the resource group was really deleted in Azure
	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(authorization.APIVersionValue))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

*/

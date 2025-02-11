/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genericarmclient

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// MakeSubscriptionScopeARMID makes an ARM ID at the subscription scope. This has the format:
// /subscriptions/00000000-0000-0000-0000-000000000000/providers/<provider>/<resourceType>/<resourceName>/...
func MakeSubscriptionScopeARMID(subscription string, provider string, params ...string) (string, error) {
	if len(params) == 0 {
		return "", errors.New("At least 2 params must be specified")
	}
	if len(params)%2 != 0 {
		return "", errors.New("ARM Id params must come in resourceKind/name pairs")
	}

	suffix := strings.Join(params, "/")

	return fmt.Sprintf("/subscriptions/%s/providers/%s/%s", subscription, provider, suffix), nil
}

// MakeResourceGroupScopeARMID makes an ARM ID at the resource group scope. This has the format:
// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/<rgName>/providers/<provider>/<resourceType>/<resourceName>/...
func MakeResourceGroupScopeARMID(subscription string, resourceGroup string, provider string, params ...string) (string, error) {
	if len(params) == 0 {
		return "", errors.New("At least 2 params must be specified")
	}
	if len(params)%2 != 0 {
		return "", errors.New("ARM Id params must come in resourceKind/name pairs")
	}

	suffix := strings.Join(params, "/")
	return fmt.Sprintf("%s/providers/%s/%s", MakeResourceGroupID(subscription, resourceGroup), provider, suffix), nil
}

// MakeResourceGroupID makes an ARM ID representing a resource group. This has the format:
// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/<rgName>
// This is "special" because there is no provider at all - but that is what creating/getting a resourceGroup expects.
func MakeResourceGroupID(subscription string, resourceGroup string) string {
	// This is a "special" format as it doesn't actually list a provider in it anywhere. I was expecting that it would
	// have a Microsoft.Resources provider but it doesn't
	return fmt.Sprintf("/subscriptions/%s/resourceGroups/%s", subscription, resourceGroup)
}

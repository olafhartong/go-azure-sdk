package replicationprotecteditems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EnableProtectionProviderSpecificInput = InMageRcmEnableProtectionInput{}

type InMageRcmEnableProtectionInput struct {
	DisksDefault                          *InMageRcmDisksDefaultInput `json:"disksDefault,omitempty"`
	DisksToInclude                        *[]InMageRcmDiskInput       `json:"disksToInclude,omitempty"`
	FabricDiscoveryMachineId              string                      `json:"fabricDiscoveryMachineId"`
	LicenseType                           *LicenseType                `json:"licenseType,omitempty"`
	MultiVmGroupName                      *string                     `json:"multiVmGroupName,omitempty"`
	ProcessServerId                       string                      `json:"processServerId"`
	RunAsAccountId                        *string                     `json:"runAsAccountId,omitempty"`
	TargetAvailabilitySetId               *string                     `json:"targetAvailabilitySetId,omitempty"`
	TargetAvailabilityZone                *string                     `json:"targetAvailabilityZone,omitempty"`
	TargetBootDiagnosticsStorageAccountId *string                     `json:"targetBootDiagnosticsStorageAccountId,omitempty"`
	TargetNetworkId                       *string                     `json:"targetNetworkId,omitempty"`
	TargetProximityPlacementGroupId       *string                     `json:"targetProximityPlacementGroupId,omitempty"`
	TargetResourceGroupId                 string                      `json:"targetResourceGroupId"`
	TargetSubnetName                      *string                     `json:"targetSubnetName,omitempty"`
	TargetVmName                          *string                     `json:"targetVmName,omitempty"`
	TargetVmSize                          *string                     `json:"targetVmSize,omitempty"`
	TestNetworkId                         *string                     `json:"testNetworkId,omitempty"`
	TestSubnetName                        *string                     `json:"testSubnetName,omitempty"`

	// Fields inherited from EnableProtectionProviderSpecificInput
}

var _ json.Marshaler = InMageRcmEnableProtectionInput{}

func (s InMageRcmEnableProtectionInput) MarshalJSON() ([]byte, error) {
	type wrapper InMageRcmEnableProtectionInput
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InMageRcmEnableProtectionInput: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InMageRcmEnableProtectionInput: %+v", err)
	}
	decoded["instanceType"] = "InMageRcm"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InMageRcmEnableProtectionInput: %+v", err)
	}

	return encoded, nil
}

package recoverypoints

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProviderSpecificRecoveryPointDetails = InMageRcmRecoveryPointDetails{}

type InMageRcmRecoveryPointDetails struct {
	IsMultiVmSyncPoint *string `json:"isMultiVmSyncPoint,omitempty"`

	// Fields inherited from ProviderSpecificRecoveryPointDetails
}

var _ json.Marshaler = InMageRcmRecoveryPointDetails{}

func (s InMageRcmRecoveryPointDetails) MarshalJSON() ([]byte, error) {
	type wrapper InMageRcmRecoveryPointDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InMageRcmRecoveryPointDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InMageRcmRecoveryPointDetails: %+v", err)
	}
	decoded["instanceType"] = "InMageRcm"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InMageRcmRecoveryPointDetails: %+v", err)
	}

	return encoded, nil
}

package replicationprotecteditems

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InMageRcmFailbackDiscoveredProtectedVmDetails struct {
	CreatedTimestamp       *string   `json:"createdTimestamp,omitempty"`
	DataStores             *[]string `json:"datastores,omitempty"`
	IPAddresses            *[]string `json:"ipAddresses,omitempty"`
	IsDeleted              *bool     `json:"isDeleted,omitempty"`
	LastDiscoveryTimeInUtc *string   `json:"lastDiscoveryTimeInUtc,omitempty"`
	OsName                 *string   `json:"osName,omitempty"`
	PowerStatus            *string   `json:"powerStatus,omitempty"`
	UpdatedTimestamp       *string   `json:"updatedTimestamp,omitempty"`
	VCenterFqdn            *string   `json:"vCenterFqdn,omitempty"`
	VCenterId              *string   `json:"vCenterId,omitempty"`
	VmFqdn                 *string   `json:"vmFqdn,omitempty"`
	VmwareToolsStatus      *string   `json:"vmwareToolsStatus,omitempty"`
}

func (o *InMageRcmFailbackDiscoveredProtectedVmDetails) GetCreatedTimestampAsTime() (*time.Time, error) {
	if o.CreatedTimestamp == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedTimestamp, "2006-01-02T15:04:05Z07:00")
}

func (o *InMageRcmFailbackDiscoveredProtectedVmDetails) SetCreatedTimestampAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreatedTimestamp = &formatted
}

func (o *InMageRcmFailbackDiscoveredProtectedVmDetails) GetLastDiscoveryTimeInUtcAsTime() (*time.Time, error) {
	if o.LastDiscoveryTimeInUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastDiscoveryTimeInUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *InMageRcmFailbackDiscoveredProtectedVmDetails) SetLastDiscoveryTimeInUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastDiscoveryTimeInUtc = &formatted
}

func (o *InMageRcmFailbackDiscoveredProtectedVmDetails) GetUpdatedTimestampAsTime() (*time.Time, error) {
	if o.UpdatedTimestamp == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.UpdatedTimestamp, "2006-01-02T15:04:05Z07:00")
}

func (o *InMageRcmFailbackDiscoveredProtectedVmDetails) SetUpdatedTimestampAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.UpdatedTimestamp = &formatted
}

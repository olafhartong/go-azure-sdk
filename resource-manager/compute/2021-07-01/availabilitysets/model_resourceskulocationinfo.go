package availabilitysets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceSkuLocationInfo struct {
	ExtendedLocations *[]string                 `json:"extendedLocations,omitempty"`
	Location          *string                   `json:"location,omitempty"`
	Type              *ExtendedLocationType     `json:"type,omitempty"`
	ZoneDetails       *[]ResourceSkuZoneDetails `json:"zoneDetails,omitempty"`
	Zones             *[]string                 `json:"zones,omitempty"`
}

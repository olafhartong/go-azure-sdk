package azuretrafficcollectors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureTrafficCollector struct {
	Etag       *string                                `json:"etag,omitempty"`
	Id         *string                                `json:"id,omitempty"`
	Location   *string                                `json:"location,omitempty"`
	Name       *string                                `json:"name,omitempty"`
	Properties *AzureTrafficCollectorPropertiesFormat `json:"properties,omitempty"`
	SystemData *SystemData                            `json:"systemData,omitempty"`
	Tags       *map[string]string                     `json:"tags,omitempty"`
	Type       *string                                `json:"type,omitempty"`
}

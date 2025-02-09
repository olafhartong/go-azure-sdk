package apioperation

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiOperationClient struct {
	Client  autorest.Client
	baseUri string
}

func NewApiOperationClientWithBaseURI(endpoint string) ApiOperationClient {
	return ApiOperationClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}

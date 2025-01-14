package appplatform

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GatewayCustomDomainsCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// GatewayCustomDomainsCreateOrUpdate ...
func (c AppPlatformClient) GatewayCustomDomainsCreateOrUpdate(ctx context.Context, id GatewayDomainId, input GatewayCustomDomainResource) (result GatewayCustomDomainsCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForGatewayCustomDomainsCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewayCustomDomainsCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForGatewayCustomDomainsCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewayCustomDomainsCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// GatewayCustomDomainsCreateOrUpdateThenPoll performs GatewayCustomDomainsCreateOrUpdate then polls until it's completed
func (c AppPlatformClient) GatewayCustomDomainsCreateOrUpdateThenPoll(ctx context.Context, id GatewayDomainId, input GatewayCustomDomainResource) error {
	result, err := c.GatewayCustomDomainsCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing GatewayCustomDomainsCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after GatewayCustomDomainsCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForGatewayCustomDomainsCreateOrUpdate prepares the GatewayCustomDomainsCreateOrUpdate request.
func (c AppPlatformClient) preparerForGatewayCustomDomainsCreateOrUpdate(ctx context.Context, id GatewayDomainId, input GatewayCustomDomainResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForGatewayCustomDomainsCreateOrUpdate sends the GatewayCustomDomainsCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForGatewayCustomDomainsCreateOrUpdate(ctx context.Context, req *http.Request) (future GatewayCustomDomainsCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

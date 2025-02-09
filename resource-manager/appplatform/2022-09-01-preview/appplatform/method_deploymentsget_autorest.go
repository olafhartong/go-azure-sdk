package appplatform

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *DeploymentResource
}

// DeploymentsGet ...
func (c AppPlatformClient) DeploymentsGet(ctx context.Context, id DeploymentId) (result DeploymentsGetOperationResponse, err error) {
	req, err := c.preparerForDeploymentsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeploymentsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "DeploymentsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeploymentsGet prepares the DeploymentsGet request.
func (c AppPlatformClient) preparerForDeploymentsGet(ctx context.Context, id DeploymentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDeploymentsGet handles the response to the DeploymentsGet request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForDeploymentsGet(resp *http.Response) (result DeploymentsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

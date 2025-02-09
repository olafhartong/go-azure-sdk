package appplatform

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppsGetResourceUploadUrlOperationResponse struct {
	HttpResponse *http.Response
	Model        *ResourceUploadDefinition
}

// AppsGetResourceUploadUrl ...
func (c AppPlatformClient) AppsGetResourceUploadUrl(ctx context.Context, id AppId) (result AppsGetResourceUploadUrlOperationResponse, err error) {
	req, err := c.preparerForAppsGetResourceUploadUrl(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "AppsGetResourceUploadUrl", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "AppsGetResourceUploadUrl", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAppsGetResourceUploadUrl(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "AppsGetResourceUploadUrl", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAppsGetResourceUploadUrl prepares the AppsGetResourceUploadUrl request.
func (c AppPlatformClient) preparerForAppsGetResourceUploadUrl(ctx context.Context, id AppId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/getResourceUploadUrl", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForAppsGetResourceUploadUrl handles the response to the AppsGetResourceUploadUrl request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForAppsGetResourceUploadUrl(resp *http.Response) (result AppsGetResourceUploadUrlOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

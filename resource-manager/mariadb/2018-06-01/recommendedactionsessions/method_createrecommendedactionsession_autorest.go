package recommendedactionsessions

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

type CreateRecommendedActionSessionOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

type CreateRecommendedActionSessionOperationOptions struct {
	DatabaseName *string
}

func DefaultCreateRecommendedActionSessionOperationOptions() CreateRecommendedActionSessionOperationOptions {
	return CreateRecommendedActionSessionOperationOptions{}
}

func (o CreateRecommendedActionSessionOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o CreateRecommendedActionSessionOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.DatabaseName != nil {
		out["databaseName"] = *o.DatabaseName
	}

	return out
}

// CreateRecommendedActionSession ...
func (c RecommendedActionSessionsClient) CreateRecommendedActionSession(ctx context.Context, id AdvisorId, options CreateRecommendedActionSessionOperationOptions) (result CreateRecommendedActionSessionOperationResponse, err error) {
	req, err := c.preparerForCreateRecommendedActionSession(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendedactionsessions.RecommendedActionSessionsClient", "CreateRecommendedActionSession", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCreateRecommendedActionSession(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendedactionsessions.RecommendedActionSessionsClient", "CreateRecommendedActionSession", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CreateRecommendedActionSessionThenPoll performs CreateRecommendedActionSession then polls until it's completed
func (c RecommendedActionSessionsClient) CreateRecommendedActionSessionThenPoll(ctx context.Context, id AdvisorId, options CreateRecommendedActionSessionOperationOptions) error {
	result, err := c.CreateRecommendedActionSession(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing CreateRecommendedActionSession: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CreateRecommendedActionSession: %+v", err)
	}

	return nil
}

// preparerForCreateRecommendedActionSession prepares the CreateRecommendedActionSession request.
func (c RecommendedActionSessionsClient) preparerForCreateRecommendedActionSession(ctx context.Context, id AdvisorId, options CreateRecommendedActionSessionOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/createRecommendedActionSession", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForCreateRecommendedActionSession sends the CreateRecommendedActionSession request. The method will close the
// http.Response Body if it receives an error.
func (c RecommendedActionSessionsClient) senderForCreateRecommendedActionSession(ctx context.Context, req *http.Request) (future CreateRecommendedActionSessionOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

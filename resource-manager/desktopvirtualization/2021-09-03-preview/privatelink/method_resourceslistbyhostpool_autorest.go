package privatelink

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourcesListByHostPoolOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]PrivateLinkResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ResourcesListByHostPoolOperationResponse, error)
}

type ResourcesListByHostPoolCompleteResult struct {
	Items []PrivateLinkResource
}

func (r ResourcesListByHostPoolOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ResourcesListByHostPoolOperationResponse) LoadMore(ctx context.Context) (resp ResourcesListByHostPoolOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ResourcesListByHostPool ...
func (c PrivateLinkClient) ResourcesListByHostPool(ctx context.Context, id HostPoolId) (resp ResourcesListByHostPoolOperationResponse, err error) {
	req, err := c.preparerForResourcesListByHostPool(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelink.PrivateLinkClient", "ResourcesListByHostPool", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelink.PrivateLinkClient", "ResourcesListByHostPool", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForResourcesListByHostPool(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelink.PrivateLinkClient", "ResourcesListByHostPool", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForResourcesListByHostPool prepares the ResourcesListByHostPool request.
func (c PrivateLinkClient) preparerForResourcesListByHostPool(ctx context.Context, id HostPoolId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/privateLinkResources", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForResourcesListByHostPoolWithNextLink prepares the ResourcesListByHostPool request with the given nextLink token.
func (c PrivateLinkClient) preparerForResourcesListByHostPoolWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
	uri, err := url.Parse(nextLink)
	if err != nil {
		return nil, fmt.Errorf("parsing nextLink %q: %+v", nextLink, err)
	}
	queryParameters := map[string]interface{}{}
	for k, v := range uri.Query() {
		if len(v) == 0 {
			continue
		}
		val := v[0]
		val = autorest.Encode("query", val)
		queryParameters[k] = val
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForResourcesListByHostPool handles the response to the ResourcesListByHostPool request. The method always
// closes the http.Response Body.
func (c PrivateLinkClient) responderForResourcesListByHostPool(resp *http.Response) (result ResourcesListByHostPoolOperationResponse, err error) {
	type page struct {
		Values   []PrivateLinkResource `json:"value"`
		NextLink *string               `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	result.Model = &respObj.Values
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ResourcesListByHostPoolOperationResponse, err error) {
			req, err := c.preparerForResourcesListByHostPoolWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "privatelink.PrivateLinkClient", "ResourcesListByHostPool", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "privatelink.PrivateLinkClient", "ResourcesListByHostPool", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForResourcesListByHostPool(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "privatelink.PrivateLinkClient", "ResourcesListByHostPool", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ResourcesListByHostPoolComplete retrieves all of the results into a single object
func (c PrivateLinkClient) ResourcesListByHostPoolComplete(ctx context.Context, id HostPoolId) (ResourcesListByHostPoolCompleteResult, error) {
	return c.ResourcesListByHostPoolCompleteMatchingPredicate(ctx, id, PrivateLinkResourceOperationPredicate{})
}

// ResourcesListByHostPoolCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c PrivateLinkClient) ResourcesListByHostPoolCompleteMatchingPredicate(ctx context.Context, id HostPoolId, predicate PrivateLinkResourceOperationPredicate) (resp ResourcesListByHostPoolCompleteResult, err error) {
	items := make([]PrivateLinkResource, 0)

	page, err := c.ResourcesListByHostPool(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading the initial page: %+v", err)
		return
	}
	if page.Model != nil {
		for _, v := range *page.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	for page.HasMore() {
		page, err = page.LoadMore(ctx)
		if err != nil {
			err = fmt.Errorf("loading the next page: %+v", err)
			return
		}

		if page.Model != nil {
			for _, v := range *page.Model {
				if predicate.Matches(v) {
					items = append(items, v)
				}
			}
		}
	}

	out := ResourcesListByHostPoolCompleteResult{
		Items: items,
	}
	return out, nil
}

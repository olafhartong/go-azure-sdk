package clusters

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

type GetOperationResponse struct {
	HttpResponse *http.Response
	Model        *Cluster
}

// Get ...
func (c ClustersClient) Get(ctx context.Context, id ClusterId) (result GetOperationResponse, err error) {
	req, err := c.preparerForGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "clusters.ClustersClient", "Get", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "clusters.ClustersClient", "Get", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "clusters.ClustersClient", "Get", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGet prepares the Get request.
func (c ClustersClient) preparerForGet(ctx context.Context, id ClusterId) (*http.Request, error) {
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

// responderForGet handles the response to the Get request. The method always
// closes the http.Response Body.
func (c ClustersClient) responderForGet(resp *http.Response) (result GetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
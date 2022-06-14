package virtualmachine

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

type RedeployOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// Redeploy ...
func (c VirtualMachineClient) Redeploy(ctx context.Context, id VirtualMachineId) (result RedeployOperationResponse, err error) {
	req, err := c.preparerForRedeploy(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachine.VirtualMachineClient", "Redeploy", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRedeploy(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachine.VirtualMachineClient", "Redeploy", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RedeployThenPoll performs Redeploy then polls until it's completed
func (c VirtualMachineClient) RedeployThenPoll(ctx context.Context, id VirtualMachineId) error {
	result, err := c.Redeploy(ctx, id)
	if err != nil {
		return fmt.Errorf("performing Redeploy: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Redeploy: %+v", err)
	}

	return nil
}

// preparerForRedeploy prepares the Redeploy request.
func (c VirtualMachineClient) preparerForRedeploy(ctx context.Context, id VirtualMachineId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/redeploy", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForRedeploy sends the Redeploy request. The method will close the
// http.Response Body if it receives an error.
func (c VirtualMachineClient) senderForRedeploy(ctx context.Context, req *http.Request) (future RedeployOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}
	future.Poller, err = polling.NewLongRunningPollerFromResponse(ctx, resp, c.Client)
	return
}
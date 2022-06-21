package contact

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

type SpacecraftsListAvailableContactsOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
	Model        *[]AvailableContacts

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (SpacecraftsListAvailableContactsOperationResponse, error)
}

type SpacecraftsListAvailableContactsCompleteResult struct {
	Items []AvailableContacts
}

func (r SpacecraftsListAvailableContactsOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r SpacecraftsListAvailableContactsOperationResponse) LoadMore(ctx context.Context) (resp SpacecraftsListAvailableContactsOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// SpacecraftsListAvailableContacts ...
func (c ContactClient) SpacecraftsListAvailableContacts(ctx context.Context, id SpacecraftId, input ContactParameters) (result SpacecraftsListAvailableContactsOperationResponse, err error) {
	req, err := c.preparerForSpacecraftsListAvailableContacts(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contact.ContactClient", "SpacecraftsListAvailableContacts", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForSpacecraftsListAvailableContacts(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contact.ContactClient", "SpacecraftsListAvailableContacts", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// SpacecraftsListAvailableContactsThenPoll performs SpacecraftsListAvailableContacts then polls until it's completed
func (c ContactClient) SpacecraftsListAvailableContactsThenPoll(ctx context.Context, id SpacecraftId, input ContactParameters) error {
	result, err := c.SpacecraftsListAvailableContacts(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing SpacecraftsListAvailableContacts: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after SpacecraftsListAvailableContacts: %+v", err)
	}

	return nil
}

// preparerForSpacecraftsListAvailableContacts prepares the SpacecraftsListAvailableContacts request.
func (c ContactClient) preparerForSpacecraftsListAvailableContacts(ctx context.Context, id SpacecraftId, input ContactParameters) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listAvailableContacts", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForSpacecraftsListAvailableContactsWithNextLink prepares the SpacecraftsListAvailableContacts request with the given nextLink token.
func (c ContactClient) preparerForSpacecraftsListAvailableContactsWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForSpacecraftsListAvailableContacts sends the SpacecraftsListAvailableContacts request. The method will close the
// http.Response Body if it receives an error.
func (c ContactClient) senderForSpacecraftsListAvailableContacts(ctx context.Context, req *http.Request) (future SpacecraftsListAvailableContactsOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}
	future.Poller, err = polling.NewLongRunningPollerFromResponse(ctx, resp, c.Client)
	return
}
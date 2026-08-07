package esi

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// get performs a GET request against the ESI API and returns the raw
// response body and headers. Every endpoint (orders, items, etc.) calls
// this instead of building its own http.Client logic - so context
// handling, error classification, and body reading happen in one place.
func (c *Client) get(ctx context.Context, url string) ([]byte, http.Header, error) {
	// 1. Build the request WITH context - this is what makes ctx do
	//    anything at all. Without NewRequestWithContext, a timeout
	//    or cancellation from the caller is silently ignored.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("esi: building request: %w", err)
	}

	// 2. Do (not Get) - Get() has no way to attach a context-aware
	//    request. Do() executes the *req you built, so ctx travels
	//    with it through the whole HTTP round trip.
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("esi: request failed: %w", err)
	}
	defer resp.Body.Close()

	// 3. Read the body once, here, regardless of status code - so
	//    callers get the error body too (useful for debugging),
	//    not just success bodies.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("esi: reading response: %w", err)
	}

	// 4. Classify the status into a sentinel error (see errors.go).
	//    This is the one place status codes get interpreted -
	//    every endpoint benefits without repeating this logic.
	if err := classifyStatus(resp.StatusCode); err != nil {
		return nil, nil, fmt.Errorf("%w: %s", err, string(body))
	}

	// 5. Success: hand back raw bytes + headers. Decoding into a
	//    specific model type is the CALLER's job, not this helper's -
	//    that's what keeps this function reusable across endpoints.
	return body, resp.Header, nil
}

package hooks

import (
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

// URLRewriteHook is a hook that rewrites URLs before they are sent to Schwab.
type URLRewriteHook struct {
	Log *logrus.Logger
}

// NewURLRewriteHook creates a new URLRewriteHook.
func NewURLRewriteHook(log *logrus.Logger) *URLRewriteHook {
	return &URLRewriteHook{
		Log: log,
	}
}

// BeforeRequest is called before a request is sent to Schwab.
func (h *URLRewriteHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	// Log the original URL.
	originalURL := req.URL.String()

	// Check if the OperationID matches getPriceHistory or getChain
	if hookCtx.OperationID == "getPriceHistory" || hookCtx.OperationID == "getChain" {
		// Rewrite the URL to point to marketdata v1
		// Current: https://api.schwabapi.com/trader/v1/...
		// Target:  https://api.schwabapi.com/marketdata/v1/...

		// Simple string replacement for the path prefix if it exists
		newURL := strings.Replace(req.URL.String(), "/trader/v1", "/marketdata/v1", 1)

		parsedURL, err := req.URL.Parse(newURL)
		if err != nil {
			if h.Log != nil {
				h.Log.Errorf("Failed to parse rewritten URL: %v", err)
			}
			return req, err
		}
		req.URL = parsedURL

		if h.Log != nil {
			h.Log.Debugf("Rewrote URL for operation %s: %s -> %s", hookCtx.OperationID, originalURL, req.URL.String())
		}
	}

	return req, nil
}

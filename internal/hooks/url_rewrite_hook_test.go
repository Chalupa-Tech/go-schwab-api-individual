package hooks

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestURLRewriteHook_BeforeRequest(t *testing.T) {
	// Setup
	logger := logrus.New()
	hook := NewURLRewriteHook(logger)

	tests := []struct {
		name        string
		operationID string
		inputURL    string
		wantURL     string
		wantErr     bool
	}{
		{
			name:        "Rewrite getPriceHistory",
			operationID: "getPriceHistory",
			inputURL:    "https://api.schwabapi.com/trader/v1/pricehistory",
			wantURL:     "https://api.schwabapi.com/marketdata/v1/pricehistory",
			wantErr:     false,
		},
		{
			name:        "Rewrite getChain",
			operationID: "getChain",
			inputURL:    "https://api.schwabapi.com/trader/v1/chains",
			wantURL:     "https://api.schwabapi.com/marketdata/v1/chains",
			wantErr:     false,
		},
		{
			name:        "Ignore other operations",
			operationID: "getAccount",
			inputURL:    "https://api.schwabapi.com/trader/v1/accounts",
			wantURL:     "https://api.schwabapi.com/trader/v1/accounts",
			wantErr:     false,
		},
		{
			name:        "Handle non-matching path prefix gracefully",
			operationID: "getPriceHistory",
			inputURL:    "https://api.schwabapi.com/other/v1/pricehistory",
			wantURL:     "https://api.schwabapi.com/other/v1/pricehistory", // Path replacement relies on /trader/v1 being present
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, _ := url.Parse(tt.inputURL)
			req := &http.Request{URL: u}
			ctx := BeforeRequestContext{
				HookContext: HookContext{
					OperationID: tt.operationID,
				},
			}

			gotReq, err := hook.BeforeRequest(ctx, req)
			if (err != nil) != tt.wantErr {
				t.Errorf("BeforeRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.wantURL, gotReq.URL.String())
		})
	}
}

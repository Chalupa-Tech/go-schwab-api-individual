package schwab

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/Chalupa-Tech/go-schwab-api-individual/schwab/auth"
)

const BaseURL = "https://api.schwabapi.com/trader/v1"
const MarketDataURL = "https://api.schwabapi.com/marketdata/v1"

type Client struct {
	Authenticator *auth.Authenticator
	tokenStore    auth.TokenStore
	httpClient    *http.Client
	mu            sync.Mutex
}

func NewClient(config auth.Config, store auth.TokenStore, client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}
	return &Client{
		Authenticator: auth.NewAuthenticator(config, client),
		tokenStore:    store,
		httpClient:    client,
	}
}

func (c *Client) do(ctx context.Context, method, urlStr string, body interface{}, v interface{}) error {
	return c.doWithRetry(ctx, method, urlStr, body, v, true)
}

func (c *Client) doWithRetry(ctx context.Context, method, urlStr string, body interface{}, v interface{}, retry bool) error {
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return err
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, urlStr, reqBody)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// Get token
	token, err := c.tokenStore.GetToken()
	if err != nil {
		return fmt.Errorf("failed to get token: %w", err)
	}
	if token != nil && token.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Handle 401
	if resp.StatusCode == http.StatusUnauthorized && retry && token != nil && token.RefreshToken != "" {
		// Refresh token
		c.mu.Lock()
		// Double check if token was updated while waiting for lock
		newToken, err := c.tokenStore.GetToken()
		if err == nil && newToken != nil && newToken.AccessToken != token.AccessToken {
			c.mu.Unlock()
			return c.doWithRetry(ctx, method, urlStr, body, v, false)
		}

		refreshedToken, err := c.Authenticator.RefreshAccessToken(token.RefreshToken)
		c.mu.Unlock()

		if err == nil {
			c.tokenStore.SaveToken(refreshedToken)
			return c.doWithRetry(ctx, method, urlStr, body, v, false)
		}
		// If refresh fails, fall through to return error from original response
	}

	if err := CheckResponse(resp); err != nil {
		return err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			return json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return nil
}
package schwab

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/placeholder/schwab-go/schwab/models"
)

const (
	defaultBaseURL = "https://api.schwabapi.com/trader/v1"
	marketDataURL  = "https://api.schwabapi.com/marketdata/v1"
)

// TokenStore is an interface for storing and retrieving tokens.
type TokenStore interface {
	GetToken() (*models.TokenResponse, error)
	SaveToken(token *models.TokenResponse) error
}

// MemoryTokenStore is a simple in-memory token store.
type MemoryTokenStore struct {
	mu    sync.RWMutex
	token *models.TokenResponse
}

func (s *MemoryTokenStore) GetToken() (*models.TokenResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.token, nil
}

func (s *MemoryTokenStore) SaveToken(token *models.TokenResponse) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.token = token
	return nil
}

// Client is the Schwab API client.
type Client struct {
	config     *Config
	tokenStore TokenStore
	httpClient *http.Client
	BaseURL    string // Trader API URL
	MarketURL  string // Market Data API URL
}

// NewClient creates a new Schwab API client.
func NewClient(config *Config, tokenStore TokenStore) *Client {
	if tokenStore == nil {
		tokenStore = &MemoryTokenStore{}
	}
	return &Client{
		config:     config,
		tokenStore: tokenStore,
		httpClient: &http.Client{Timeout: 30 * time.Second},
		BaseURL:    defaultBaseURL,
		MarketURL:  marketDataURL,
	}
}

// doRequest performs an HTTP request, handling authentication and error checking.
func (c *Client) doRequest(ctx context.Context, method, urlStr string, body interface{}, result interface{}) error {
	var reqBody []byte
	var err error
	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, urlStr, bytes.NewReader(reqBody))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
		// Set GetBody for retries
		req.GetBody = func() (io.ReadCloser, error) {
			return io.NopCloser(bytes.NewReader(reqBody)), nil
		}
	}

	return c.executeRequest(req, result, true)
}

func (c *Client) executeRequest(req *http.Request, result interface{}, retryOnAuth bool) error {
	token, err := c.tokenStore.GetToken()
	if err != nil {
		return fmt.Errorf("failed to get token: %w", err)
	}
	if token == nil || token.AccessToken == "" {
		return fmt.Errorf("no access token available")
	}

	// Check if token needs refresh (naive check, assuming we know expiry time, but TokenResponse only has ExpiresIn seconds)
	// Ideally TokenStore handles validity. But here we rely on 401 retry mostly, or we could calculate expiry.
	// We'll proceed with 401 retry logic for robustness.

	req.Header.Set("Authorization", "Bearer "+token.AccessToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized && retryOnAuth {
		// Attempt refresh
		newToken, err := c.config.RefreshAccessToken(token.RefreshToken)
		if err != nil {
			return fmt.Errorf("failed to refresh token on 401: %w", err)
		}
		if err := c.tokenStore.SaveToken(newToken); err != nil {
			return fmt.Errorf("failed to save new token: %w", err)
		}
		// Retry request with new token
		// Clone request to avoid reusing body if it was read (bytes.Buffer matches Reader interface but we created it from bytes so it's fine to re-create or seek? Request.Body is ReadCloser, usually one-time use)
		// To be safe, we need to handle body rewinding if we want to retry.
		// Since we passed io.Reader to NewRequest, if it was *bytes.Buffer, it's consumed.
		// So we should re-create the request or use GetBody if available.

		if req.GetBody != nil {
			newBody, err := req.GetBody()
			if err == nil {
				req.Body = newBody
			}
		}

		return c.executeRequest(req, result, false)
	}

	if resp.StatusCode >= 400 {
		var apiErr APIError
		apiErr.StatusCode = resp.StatusCode
		apiErr.SchwabClientCorrelId = resp.Header.Get("Schwab-Client-CorrelId")

		// Try to decode error body
		bodyBytes, _ := io.ReadAll(resp.Body)
		if len(bodyBytes) > 0 {
			// Try to unmarshal into standard error format or just put in Message
			// Spec has ErrorResponse or ServiceError
			// We try generic map or specific struct?
			// Let's try to unmarshal into APIError fields if they match, or just put raw body in Message
			// The spec ServiceError has {message: string, errors: []string}
			// APIError has {Message: string, Errors: []string}
			// So json.Unmarshal might work if JSON tags match.
			// ServiceError JSON tags: message, errors. Matches APIError.
			_ = json.Unmarshal(bodyBytes, &apiErr)
			if apiErr.Message == "" {
				apiErr.Message = string(bodyBytes)
			}
		} else {
			apiErr.Message = resp.Status
		}
		return &apiErr
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

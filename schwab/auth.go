package schwab

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/placeholder/schwab-go/schwab/models"
)

const (
	defaultAuthURL  = "https://api.schwabapi.com/v1/oauth/authorize"
	defaultTokenURL = "https://api.schwabapi.com/v1/oauth/token"
)

// Config holds the configuration for the Schwab API client.
type Config struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	AuthURL      string   // Optional, defaults to production URL
	TokenURL     string   // Optional, defaults to production URL
	Scopes       []string // Optional, defaults to "readonly"
}

// GetAuthorizationURL returns the URL for the user to authorize the application.
func (c *Config) GetAuthorizationURL() string {
	u := c.AuthURL
	if u == "" {
		u = defaultAuthURL
	}
	v := url.Values{}
	v.Set("response_type", "code")
	v.Set("client_id", c.ClientID)
	v.Set("redirect_uri", c.RedirectURL)

	// Set scope, defaulting to "readonly" if not provided
	scope := "readonly"
	if len(c.Scopes) > 0 {
		scope = strings.Join(c.Scopes, " ")
	}
	v.Set("scope", scope)

	return fmt.Sprintf("%s?%s", u, v.Encode())
}

// ExchangeAuthCode exchanges the authorization code for an access token.
// It ensures the code is URL-decoded before sending.
func (c *Config) ExchangeAuthCode(code string) (*models.TokenResponse, error) {
	// Ensure code is decoded. If it was passed from a URL query parameter using standard library, it's likely already decoded.
	// But if it was manually pasted or raw, we might need to handle it.
	// The instruction says "Ensure the authorization code is URL-decoded before sending".
	// We will assume the input 'code' might be percent-encoded if it came from a raw source.
	// However, if we blindly URL-decode, we might corrupt a code that was already decoded and contained a '%' literal (unlikely in base64/hex but possible).
	// But typically auth codes are URL-safe strings or might contain chars like '+', '/'.
	// If the user passes "abc%2Bdef", they mean "abc+def".
	// If the user passes "abc+def", they mean "abc+def".
	// To be safe, if it looks like it's url-encoded, we decode it.
	// A robust way is to try decoding. If it fails, use original.
	decodedCode, err := url.PathUnescape(code)
	if err == nil {
		code = decodedCode
	}

	return c.doTokenRequest(url.Values{
		"grant_type":   {"authorization_code"},
		"code":         {code},
		"redirect_uri": {c.RedirectURL},
	})
}

// RefreshAccessToken refreshes the access token using the refresh token.
func (c *Config) RefreshAccessToken(refreshToken string) (*models.TokenResponse, error) {
	return c.doTokenRequest(url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {refreshToken},
	})
}

func (c *Config) doTokenRequest(v url.Values) (*models.TokenResponse, error) {
	u := c.TokenURL
	if u == "" {
		u = defaultTokenURL
	}

	req, err := http.NewRequest("POST", u, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Basic Auth: Base64(Client_ID:Client_Secret)
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.ClientID, c.ClientSecret)))
	req.Header.Set("Authorization", "Basic "+auth)

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("token request failed: status %d body: %s", resp.StatusCode, string(bodyBytes))
	}

	var tokenResp models.TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, fmt.Errorf("failed to decode token response: %w", err)
	}

	return &tokenResp, nil
}

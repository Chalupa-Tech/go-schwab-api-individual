package auth

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockTransport struct {
	Handler http.Handler
}

func (m *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	w := httptest.NewRecorder()
	m.Handler.ServeHTTP(w, req)
	return w.Result(), nil
}

func TestExchangeAuthCode(t *testing.T) {
	clientID := "test-client-id"
	clientSecret := "test-client-secret"
	authCode := "test-auth-code@"

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/v1/oauth/token" {
			t.Errorf("Expected path /v1/oauth/token, got %s", r.URL.Path)
		}

		authHeader := r.Header.Get("Authorization")
		expectedAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(clientID+":"+clientSecret))
		if authHeader != expectedAuth {
			t.Errorf("Expected Authorization header %s, got %s", expectedAuth, authHeader)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}
		if r.Form.Get("grant_type") != "authorization_code" {
			t.Errorf("Expected grant_type authorization_code, got %s", r.Form.Get("grant_type"))
		}
		// The code sent should be the raw code provided (if it was already decoded or not).
		// Wait, ExchangeAuthCode calls url.QueryUnescape(code).
		// If input is "test-auth-code@", unescape does nothing change.
		// If input is "test%40", it becomes "test@".
		if r.Form.Get("code") != authCode {
			t.Errorf("Expected code %s, got %s", authCode, r.Form.Get("code"))
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{
			"access_token": "new-access-token",
			"refresh_token": "new-refresh-token",
			"expires_in": 1800,
			"token_type": "Bearer",
			"scope": "readonly"
		}`))
	})

	client := &http.Client{
		Transport: &MockTransport{Handler: handler},
	}

	config := Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost",
		Scopes:       []string{"readonly"},
	}

	auth := NewAuthenticator(config, client)
	url := auth.GetAuthorizationURL()
	if !strings.Contains(url, "scope=readonly") {
		t.Errorf("Expected URL to contain scope=readonly, got %s", url)
	}

	token, err := auth.ExchangeAuthCode(authCode)
	if err != nil {
		t.Fatalf("ExchangeAuthCode failed: %v", err)
	}

	if token.AccessToken != "new-access-token" {
		t.Errorf("Expected access token new-access-token, got %s", token.AccessToken)
	}
}

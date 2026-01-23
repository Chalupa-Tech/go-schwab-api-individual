package schwab

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/placeholder/schwab-go/schwab/models"
)

func TestGetAuthorizationURL(t *testing.T) {
	config := &Config{
		ClientID:    "test_client_id",
		RedirectURL: "http://localhost:8080/callback",
	}
	url := config.GetAuthorizationURL()
	expected := "https://api.schwabapi.com/v1/oauth/authorize?client_id=test_client_id&redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fcallback&response_type=code&scope=readonly"

	// URL params order might differ, check containment
	if url != expected {
		// Basic check for components
		if !contains(url, "client_id=test_client_id") {
			t.Errorf("URL missing client_id: %s", url)
		}
		if !contains(url, "redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fcallback") {
			t.Errorf("URL missing redirect_uri: %s", url)
		}
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && len(s)-len(substr) >= 0 && (s[0:len(substr)] == substr || contains(s[1:], substr))
}

func TestExchangeAuthCode(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/v1/oauth/token" {
			t.Errorf("Expected path /v1/oauth/token, got %s", r.URL.Path)
		}
		if err := r.ParseForm(); err != nil {
			t.Errorf("ParseForm() err: %v", err)
		}
		if r.FormValue("grant_type") != "authorization_code" {
			t.Errorf("Expected grant_type authorization_code, got %s", r.FormValue("grant_type"))
		}
		if r.FormValue("code") != "test_code" {
			t.Errorf("Expected code test_code, got %s", r.FormValue("code"))
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.TokenResponse{
			AccessToken:  "access_token",
			RefreshToken: "refresh_token",
			ExpiresIn:    1800,
		})
	}))
	defer server.Close()

	config := &Config{
		ClientID:     "id",
		ClientSecret: "secret",
		TokenURL:     server.URL + "/v1/oauth/token",
	}

	token, err := config.ExchangeAuthCode("test_code")
	if err != nil {
		t.Fatalf("ExchangeAuthCode failed: %v", err)
	}

	if token.AccessToken != "access_token" {
		t.Errorf("Expected access_token, got %s", token.AccessToken)
	}
}

func TestGetAccounts(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer valid_token" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if r.URL.Path != "/accounts" {
			t.Errorf("Expected path /accounts, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]models.Account{
			{SecuritiesAccount: models.SecuritiesAccount{AccountNumber: "12345"}},
		})
	}))
	defer server.Close()

	tokenStore := &MemoryTokenStore{
		token: &models.TokenResponse{AccessToken: "valid_token"},
	}

	config := &Config{}
	client := NewClient(config, tokenStore)
	client.BaseURL = server.URL

	accounts, err := client.GetAccounts(context.Background(), "")
	if err != nil {
		t.Fatalf("GetAccounts failed: %v", err)
	}

	if len(accounts) != 1 {
		t.Errorf("Expected 1 account, got %d", len(accounts))
	}
	if accounts[0].SecuritiesAccount.AccountNumber != "12345" {
		t.Errorf("Expected account 12345, got %s", accounts[0].SecuritiesAccount.AccountNumber)
	}
}

func TestGetQuotes(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/quotes" {
			t.Errorf("Expected path /quotes, got %s", r.URL.Path)
		}
		query := r.URL.Query()
		if query.Get("symbols") != "AAPL,GOOG" {
			t.Errorf("Expected symbols AAPL,GOOG, got %s", query.Get("symbols"))
		}

		w.Header().Set("Content-Type", "application/json")
		resp := models.QuoteResponse{
			"AAPL": models.QuoteResponseObject{Symbol: "AAPL", Quote: &models.QuoteData{LastPrice: 150.0}},
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	tokenStore := &MemoryTokenStore{
		token: &models.TokenResponse{AccessToken: "valid_token"},
	}

	config := &Config{}
	client := NewClient(config, tokenStore)
	client.MarketURL = server.URL

	quotes, err := client.GetQuotes(context.Background(), []string{"AAPL", "GOOG"}, "", false)
	if err != nil {
		t.Fatalf("GetQuotes failed: %v", err)
	}

	if quotes["AAPL"].Symbol != "AAPL" {
		t.Errorf("Expected AAPL symbol")
	}
}

func TestTokenRefreshLogic(t *testing.T) {
	// First request 401, refresh token, retry success
	var reqCount int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqCount++
		if reqCount == 1 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`)) // Return empty array for GetAccounts
	}))
	defer server.Close()

	// Mock token server
	tokenServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(models.TokenResponse{
			AccessToken:  "new_token",
			RefreshToken: "new_refresh",
		})
	}))
	defer tokenServer.Close()

	tokenStore := &MemoryTokenStore{
		token: &models.TokenResponse{AccessToken: "expired_token", RefreshToken: "valid_refresh"},
	}

	config := &Config{
		TokenURL: tokenServer.URL,
	}
	client := NewClient(config, tokenStore)
	client.BaseURL = server.URL

	_, err := client.GetAccounts(context.Background(), "")
	if err != nil {
		t.Fatalf("GetAccounts failed with refresh logic: %v", err)
	}

	newToken, _ := tokenStore.GetToken()
	if newToken.AccessToken != "new_token" {
		t.Errorf("Expected new_token, got %s", newToken.AccessToken)
	}
}

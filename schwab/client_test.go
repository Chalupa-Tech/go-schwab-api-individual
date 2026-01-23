package schwab

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Chalupa-Tech/go-schwab-api-individual/schwab/auth"
	"github.com/Chalupa-Tech/go-schwab-api-individual/schwab/models"
)

type MockTokenStore struct {
	Token *models.TokenResponse
}

func (m *MockTokenStore) SaveToken(token *models.TokenResponse) error {
	m.Token = token
	return nil
}

func (m *MockTokenStore) GetToken() (*models.TokenResponse, error) {
	return m.Token, nil
}

type MockTransport struct {
	Handler http.Handler
}

func (m *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	w := httptest.NewRecorder()
	m.Handler.ServeHTTP(w, req)
	return w.Result(), nil
}

func TestClient_RefreshTokenOn401(t *testing.T) {
	// Setup
	tokenStore := &MockTokenStore{
		Token: &models.TokenResponse{
			AccessToken:  "old-access-token",
			RefreshToken: "valid-refresh-token",
		},
	}

	var authCalled, apiCalledFirst, apiCalledSecond bool

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/oauth/token") {
			authCalled = true
			if r.FormValue("grant_type") != "refresh_token" {
				t.Errorf("Expected grant_type refresh_token")
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(models.TokenResponse{
				AccessToken:  "new-access-token",
				RefreshToken: "valid-refresh-token",
			})
			return
		}

		if strings.Contains(r.URL.Path, "/accounts") {
			authHeader := r.Header.Get("Authorization")
			if !apiCalledFirst {
				apiCalledFirst = true
				if authHeader != "Bearer old-access-token" {
					t.Errorf("Expected old token, got %s", authHeader)
				}
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			apiCalledSecond = true
			if authHeader != "Bearer new-access-token" {
				t.Errorf("Expected new token, got %s", authHeader)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`[{"type": "CASH"}]`))
			return
		}
	})

	client := NewClient(auth.Config{}, tokenStore, &http.Client{
		Transport: &MockTransport{Handler: handler},
	})

	// Execute
	_, err := client.GetAccounts(context.Background(), "")
	if err != nil {
		t.Fatalf("GetAccounts failed: %v", err)
	}

	// Verify
	if !authCalled {
		t.Error("Auth endpoint was not called")
	}
	if !apiCalledFirst {
		t.Error("API was not called first time")
	}
	if !apiCalledSecond {
		t.Error("API was not called second time")
	}
	if tokenStore.Token.AccessToken != "new-access-token" {
		t.Error("Token store was not updated")
	}
}

package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/Chalupa-Tech/go-schwab-api-individual/schwab/models"
)

const (
	AuthBaseURL  = "https://api.schwabapi.com/v1/oauth"
	AuthorizeURL = AuthBaseURL + "/authorize"
	TokenURL     = AuthBaseURL + "/token"
)

type Config struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Scopes       []string
}

type TokenStore interface {
	SaveToken(token *models.TokenResponse) error
	GetToken() (*models.TokenResponse, error)
}

type Authenticator struct {
	config Config
	client *http.Client
}

func NewAuthenticator(config Config, client *http.Client) *Authenticator {
	if client == nil {
		client = http.DefaultClient
	}
	return &Authenticator{
		config: config,
		client: client,
	}
}

func (a *Authenticator) GetAuthorizationURL() string {
	u, _ := url.Parse(AuthorizeURL)
	q := u.Query()
	q.Set("client_id", a.config.ClientID)
	q.Set("redirect_uri", a.config.RedirectURL)
	q.Set("response_type", "code")
	if len(a.config.Scopes) > 0 {
		q.Set("scope", strings.Join(a.config.Scopes, " "))
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func (a *Authenticator) ExchangeAuthCode(code string) (*models.TokenResponse, error) {
	decodedCode, err := url.QueryUnescape(code)
	if err != nil {
		return nil, fmt.Errorf("failed to decode auth code: %w", err)
	}

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", decodedCode)
	data.Set("redirect_uri", a.config.RedirectURL)

	return a.doTokenRequest(data)
}

func (a *Authenticator) RefreshAccessToken(refreshToken string) (*models.TokenResponse, error) {
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)

	return a.doTokenRequest(data)
}

func (a *Authenticator) doTokenRequest(data url.Values) (*models.TokenResponse, error) {
	req, err := http.NewRequest("POST", TokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	auth := base64.StdEncoding.EncodeToString([]byte(a.config.ClientID + ":" + a.config.ClientSecret))
	req.Header.Set("Authorization", "Basic "+auth)

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("token request failed (status %d): %s", resp.StatusCode, string(body))
	}

	var tokenResp models.TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, err
	}

	return &tokenResp, nil
}

package schwab

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/placeholder/schwab-go/schwab/models"
)

// --- Account Access Service ---

// GetAccountNumbers returns a list of account numbers and their encrypted values.
// Account numbers in plain text cannot be used outside of headers or request/response bodies.
// This service must be invoked to retrieve the list of plain text/encrypted value pairs.
func (c *Client) GetAccountNumbers(ctx context.Context) ([]models.AccountNumberHash, error) {
	u := fmt.Sprintf("%s/accounts/accountNumbers", c.BaseURL)
	var result []models.AccountNumberHash
	err := c.doRequest(ctx, "GET", u, nil, &result)
	return result, err
}

// GetAccounts returns linked account(s) balances and positions for the logged in user.
// fields: "positions" to return positions.
func (c *Client) GetAccounts(ctx context.Context, fields string) ([]models.Account, error) {
	u := fmt.Sprintf("%s/accounts", c.BaseURL)
	if fields != "" {
		u += "?fields=" + url.QueryEscape(fields)
	}
	var result []models.Account
	err := c.doRequest(ctx, "GET", u, nil, &result)
	return result, err
}

// GetAccount returns a specific account balance and positions.
// accountNumber: The encrypted ID of the account.
func (c *Client) GetAccount(ctx context.Context, accountNumber string, fields string) (*models.Account, error) {
	u := fmt.Sprintf("%s/accounts/%s", c.BaseURL, url.PathEscape(accountNumber))
	if fields != "" {
		u += "?fields=" + url.QueryEscape(fields)
	}
	var result models.Account
	err := c.doRequest(ctx, "GET", u, nil, &result)
	return &result, err
}

// GetOrdersParams holds parameters for GetOrders.
type GetOrdersParams struct {
	FromEnteredTime string
	ToEnteredTime   string
	MaxResults      int
	Status          string
}

// GetOrders returns all orders for a specific account.
func (c *Client) GetOrders(ctx context.Context, accountNumber string, params GetOrdersParams) ([]models.Order, error) {
	u := fmt.Sprintf("%s/accounts/%s/orders", c.BaseURL, url.PathEscape(accountNumber))
	v := url.Values{}
	v.Set("fromEnteredTime", params.FromEnteredTime)
	v.Set("toEnteredTime", params.ToEnteredTime)
	if params.MaxResults > 0 {
		v.Set("maxResults", strconv.Itoa(params.MaxResults))
	}
	if params.Status != "" {
		v.Set("status", params.Status)
	}
	u += "?" + v.Encode()

	var result []models.Order
	err := c.doRequest(ctx, "GET", u, nil, &result)
	return result, err
}

// PlaceOrder places an order for a specific account.
func (c *Client) PlaceOrder(ctx context.Context, accountNumber string, order models.OrderRequest) error {
	u := fmt.Sprintf("%s/accounts/%s/orders", c.BaseURL, url.PathEscape(accountNumber))
	// Response is 201 Created with Location header, body empty usually.
	// doRequest expects a result pointer or nil.
	return c.doRequest(ctx, "POST", u, order, nil)
}

// GetOrder returns a specific order by its ID.
func (c *Client) GetOrder(ctx context.Context, accountNumber string, orderId int64) (*models.Order, error) {
	u := fmt.Sprintf("%s/accounts/%s/orders/%d", c.BaseURL, url.PathEscape(accountNumber), orderId)
	var result models.Order
	err := c.doRequest(ctx, "GET", u, nil, &result)
	return &result, err
}

// CancelOrder cancels a specific order.
func (c *Client) CancelOrder(ctx context.Context, accountNumber string, orderId int64) error {
	u := fmt.Sprintf("%s/accounts/%s/orders/%d", c.BaseURL, url.PathEscape(accountNumber), orderId)
	return c.doRequest(ctx, "DELETE", u, nil, nil)
}

// ReplaceOrder replaces an existing order.
func (c *Client) ReplaceOrder(ctx context.Context, accountNumber string, orderId int64, order models.OrderRequest) error {
	u := fmt.Sprintf("%s/accounts/%s/orders/%d", c.BaseURL, url.PathEscape(accountNumber), orderId)
	return c.doRequest(ctx, "PUT", u, order, nil)
}

// GetAllOrders returns all orders for all accounts.
func (c *Client) GetAllOrders(ctx context.Context, params GetOrdersParams) ([]models.Order, error) {
	u := fmt.Sprintf("%s/orders", c.BaseURL)
	v := url.Values{}
	v.Set("fromEnteredTime", params.FromEnteredTime)
	v.Set("toEnteredTime", params.ToEnteredTime)
	if params.MaxResults > 0 {
		v.Set("maxResults", strconv.Itoa(params.MaxResults))
	}
	if params.Status != "" {
		v.Set("status", params.Status)
	}
	u += "?" + v.Encode()

	var result []models.Order
	err := c.doRequest(ctx, "GET", u, nil, &result)
	return result, err
}

// PreviewOrder previews an order for a specific account.
func (c *Client) PreviewOrder(ctx context.Context, accountNumber string, order models.PreviewOrder) (*models.PreviewOrder, error) {
	u := fmt.Sprintf("%s/accounts/%s/previewOrder", c.BaseURL, url.PathEscape(accountNumber))
	var result models.PreviewOrder
	err := c.doRequest(ctx, "POST", u, order, &result)
	return &result, err
}

// GetTransactionsParams holds parameters for GetTransactions.
type GetTransactionsParams struct {
	StartDate string
	EndDate   string
	Symbol    string
	Types     string // Comma separated types
}

// GetTransactions returns all transactions information for a specific account.
func (c *Client) GetTransactions(ctx context.Context, accountNumber string, params GetTransactionsParams) ([]models.Transaction, error) {
	u := fmt.Sprintf("%s/accounts/%s/transactions", c.BaseURL, url.PathEscape(accountNumber))
	v := url.Values{}
	v.Set("startDate", params.StartDate)
	v.Set("endDate", params.EndDate)
	if params.Symbol != "" {
		v.Set("symbol", params.Symbol)
	}
	if params.Types != "" {
		v.Set("types", params.Types)
	}
	u += "?" + v.Encode()

	var result []models.Transaction
	err := c.doRequest(ctx, "GET", u, nil, &result)
	return result, err
}

// GetTransaction returns specific transaction information.
func (c *Client) GetTransaction(ctx context.Context, accountNumber string, transactionId int64) ([]models.Transaction, error) {
	u := fmt.Sprintf("%s/accounts/%s/transactions/%d", c.BaseURL, url.PathEscape(accountNumber), transactionId)
	var result []models.Transaction // Spec returns a list/array even for single ID? Check spec.
	// Spec for /transactions/{transactionId} response schema is array of Transaction.
	err := c.doRequest(ctx, "GET", u, nil, &result)
	return result, err
}

// GetUserPreference returns user preference information.
func (c *Client) GetUserPreference(ctx context.Context) ([]models.UserPreference, error) {
	u := fmt.Sprintf("%s/userPreference", c.BaseURL)
	var result []models.UserPreference
	err := c.doRequest(ctx, "GET", u, nil, &result)
	return result, err
}

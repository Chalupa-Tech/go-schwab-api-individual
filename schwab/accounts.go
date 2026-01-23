package schwab

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/Chalupa-Tech/go-schwab-api-individual/schwab/models"
)

// GetAccountNumbers returns a list of account numbers and their encrypted values.
func (c *Client) GetAccountNumbers(ctx context.Context) ([]models.AccountNumberHash, error) {
	var result []models.AccountNumberHash
	err := c.do(ctx, "GET", BaseURL+"/accounts/accountNumbers", nil, &result)
	return result, err
}

// GetAccounts returns linked account balances and positions.
// fields can be "positions" or empty.
func (c *Client) GetAccounts(ctx context.Context, fields string) ([]models.Account, error) {
	u := BaseURL + "/accounts"
	if fields != "" {
		u += "?fields=" + fields
	}
	var result []models.Account
	err := c.do(ctx, "GET", u, nil, &result)
	return result, err
}

// GetAccount returns a specific account balance and positions.
func (c *Client) GetAccount(ctx context.Context, accountNumber string, fields string) (*models.Account, error) {
	u := BaseURL + "/accounts/" + accountNumber
	if fields != "" {
		u += "?fields=" + fields
	}
	var result models.Account
	err := c.do(ctx, "GET", u, nil, &result)
	return &result, err
}

// GetOrdersParams represents parameters for GetOrders.
type GetOrdersParams struct {
	MaxResults      int
	FromEnteredTime string
	ToEnteredTime   string
	Status          string
}

// GetOrders returns all orders for a specific account.
func (c *Client) GetOrders(ctx context.Context, accountNumber string, params GetOrdersParams) ([]models.Order, error) {
	u, _ := url.Parse(BaseURL + "/accounts/" + accountNumber + "/orders")
	q := u.Query()
	if params.MaxResults > 0 {
		q.Set("maxResults", strconv.Itoa(params.MaxResults))
	}
	if params.FromEnteredTime != "" {
		q.Set("fromEnteredTime", params.FromEnteredTime)
	}
	if params.ToEnteredTime != "" {
		q.Set("toEnteredTime", params.ToEnteredTime)
	}
	if params.Status != "" {
		q.Set("status", params.Status)
	}
	u.RawQuery = q.Encode()

	var result []models.Order
	err := c.do(ctx, "GET", u.String(), nil, &result)
	return result, err
}

// PlaceOrder places an order for a specific account.
func (c *Client) PlaceOrder(ctx context.Context, accountNumber string, order models.Order) error {
	u := BaseURL + "/accounts/" + accountNumber + "/orders"
	return c.do(ctx, "POST", u, order, nil)
}

// GetOrder returns a specific order.
func (c *Client) GetOrder(ctx context.Context, accountNumber string, orderId int64) (*models.Order, error) {
	u := fmt.Sprintf("%s/accounts/%s/orders/%d", BaseURL, accountNumber, orderId)
	var result models.Order
	err := c.do(ctx, "GET", u, nil, &result)
	return &result, err
}

// CancelOrder cancels a specific order.
func (c *Client) CancelOrder(ctx context.Context, accountNumber string, orderId int64) error {
	u := fmt.Sprintf("%s/accounts/%s/orders/%d", BaseURL, accountNumber, orderId)
	return c.do(ctx, "DELETE", u, nil, nil)
}

// ReplaceOrder replaces an existing order.
func (c *Client) ReplaceOrder(ctx context.Context, accountNumber string, orderId int64, order models.Order) error {
	u := fmt.Sprintf("%s/accounts/%s/orders/%d", BaseURL, accountNumber, orderId)
	return c.do(ctx, "PUT", u, order, nil)
}

// GetAllOrders returns all orders for all accounts.
func (c *Client) GetAllOrders(ctx context.Context, params GetOrdersParams) ([]models.Order, error) {
	u, _ := url.Parse(BaseURL + "/orders")
	q := u.Query()
	if params.MaxResults > 0 {
		q.Set("maxResults", strconv.Itoa(params.MaxResults))
	}
	if params.FromEnteredTime != "" {
		q.Set("fromEnteredTime", params.FromEnteredTime)
	}
	if params.ToEnteredTime != "" {
		q.Set("toEnteredTime", params.ToEnteredTime)
	}
	if params.Status != "" {
		q.Set("status", params.Status)
	}
	u.RawQuery = q.Encode()

	var result []models.Order
	err := c.do(ctx, "GET", u.String(), nil, &result)
	return result, err
}

// GetTransactionsParams represents parameters for GetTransactions.
type GetTransactionsParams struct {
	StartDate string
	EndDate   string
	Symbol    string
	Types     string
}

// GetTransactions returns all transactions for a specific account.
func (c *Client) GetTransactions(ctx context.Context, accountNumber string, params GetTransactionsParams) ([]models.Transaction, error) {
	u, _ := url.Parse(BaseURL + "/accounts/" + accountNumber + "/transactions")
	q := u.Query()
	if params.StartDate != "" {
		q.Set("startDate", params.StartDate)
	}
	if params.EndDate != "" {
		q.Set("endDate", params.EndDate)
	}
	if params.Symbol != "" {
		q.Set("symbol", params.Symbol)
	}
	if params.Types != "" {
		q.Set("types", params.Types)
	}
	u.RawQuery = q.Encode()

	var result []models.Transaction
	err := c.do(ctx, "GET", u.String(), nil, &result)
	return result, err
}

// GetUserPreference returns user preferences.
func (c *Client) GetUserPreference(ctx context.Context) ([]models.UserPreference, error) {
	var result []models.UserPreference
	err := c.do(ctx, "GET", BaseURL+"/userPreference", nil, &result)
	return result, err
}

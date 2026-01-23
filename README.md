# Schwab API Go Client

A Go client library for the Schwab Trader API.

## Features

- **Authentication**: Easy OAuth2 flow handling (Authorization Code, Token Refresh).
- **Account Access**: Manage accounts, orders, transactions, and user preferences.
- **Market Data**: Access real-time quotes, option chains, price history, and more.
- **Type-Safe**: Full Go structs for API requests and responses.

## Installation

```bash
go get github.com/placeholder/schwab-go
```

## Usage

### Authentication

The Schwab API uses OAuth2. You need to obtain an access token.

1. **Configure Client:**

```go
config := &schwab.Config{
    ClientID:     "YOUR_APP_KEY",
    ClientSecret: "YOUR_APP_SECRET",
    RedirectURL:  "https://your-redirect-url.com",
}
```

2. **Get Authorization URL:**

Redirect your user to this URL:

```go
authURL := config.GetAuthorizationURL()
fmt.Printf("Visit this URL to authorize: %s\n", authURL)
```

3. **Exchange Code for Token:**

After the user authorizes, they will be redirected back with a `code`.

```go
code := "CODE_FROM_CALLBACK"
token, err := config.ExchangeAuthCode(code)
if err != nil {
    log.Fatal(err)
}
```

import (
    "context"
    "fmt"
    "log"
    "github.com/placeholder/schwab-go/schwab"
    "github.com/placeholder/schwab-go/schwab/models"
)

### Creating a Client

Implement a `TokenStore` to persist tokens (optional, default is in-memory).

```go
// Simple in-memory store provided by library
store := &schwab.MemoryTokenStore{}
store.SaveToken(token)

client := schwab.NewClient(config, store)
```

### Making Requests

**Get Accounts:**

```go
ctx := context.Background()
accounts, err := client.GetAccounts(ctx, "positions")
if err != nil {
    log.Fatal(err)
}
for _, acc := range accounts {
    fmt.Printf("Account: %s\n", acc.SecuritiesAccount.AccountNumber)
}
```

**Get Quotes:**

```go
quotes, err := client.GetQuotes(ctx, []string{"AAPL", "GOOG"}, "", false)
if err != nil {
    log.Fatal(err)
}
if q, ok := quotes["AAPL"]; ok {
    fmt.Printf("AAPL Price: %f\n", q.Quote.LastPrice)
}
```

**Place Order:**

```go
order := models.OrderRequest{
    OrderType: "MARKET",
    Session:   "NORMAL",
    Duration:  "DAY",
    OrderStrategyType: "SINGLE",
    OrderLegCollection: []models.OrderLeg{
        {
            Instruction: "BUY",
            Quantity:    1,
            Instrument: &models.AccountsInstrument{
                Symbol:    "AAPL",
                AssetType: "EQUITY",
            },
        },
    },
}

err := client.PlaceOrder(ctx, "ENCRYPTED_ACCOUNT_ID", order)
if err != nil {
    log.Fatal(err)
}
```

## License

MIT
```
# Schwab API Go Client

A Go client library for the Schwab Individual Trader API.

## Installation

```bash
go get github.com/Chalupa-Tech/go-schwab-api-individual
```

## Usage

### Authentication

The Schwab API uses OAuth 2.0. You need to obtain an access token before making API calls.

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Chalupa-Tech/go-schwab-api-individual/schwab"
	"github.com/Chalupa-Tech/go-schwab-api-individual/schwab/auth"
)

func main() {
	config := auth.Config{
		ClientID:     "YOUR_CLIENT_ID",
		ClientSecret: "YOUR_CLIENT_SECRET",
		RedirectURL:  "https://127.0.0.1",
	}

	// You need to implement TokenStore to save/load tokens
	store := &MyTokenStore{}

	client := schwab.NewClient(config, store, nil)

	// Step 1: Get Authorization URL
	// authURL := client.Authenticator.GetAuthorizationURL()
	// fmt.Println("Go to:", authURL)
	
	// Step 2: After user logs in, exchange code
	// code := "CODE_FROM_CALLBACK"
	// token, err := client.Authenticator.ExchangeAuthCode(code)
    // store.SaveToken(token)

	// Step 3: Make API calls
	accounts, err := client.GetAccounts(context.Background(), "positions")
	if err != nil {
		log.Fatal(err)
	}

	for _, acct := range accounts {
		fmt.Printf("Account: %s\n", acct.SecuritiesAccount.AccountNumber)
	}
}

type MyTokenStore struct {
    // Implement saving to file/DB
}
// ... Implement TokenStore methods
```

## Features

- **Authentication**: Handles OAuth2 flow and automatic token refresh.
- **Account Access**: Get accounts, balances, positions, orders, and transactions.
- **Market Data**: Get quotes, option chains, price history.

## License

MIT

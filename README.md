# Charles Schwab Individual Trader API SDK

Developer-friendly & type-safe Go SDK specifically catered to leverage the **Charles Schwab Individual Trader** API.

[![Built by Speakeasy](https://img.shields.io/badge/Built_by-SPEAKEASY-374151?style=for-the-badge&labelColor=f3f4f6)](https://www.speakeasy.com/?utm_source=go-schwab-api-individual&utm_campaign=go)
[![License: MIT](https://img.shields.io/badge/LICENSE_//_MIT-3b5bdb?style=for-the-badge&labelColor=eff6ff)](https://opensource.org/licenses/MIT)


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/chalupa-tech/go-schwab). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

Market Data: Trader API - Market data
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [Overview](#charles-schwab-individual-trader-api-sdk)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Custom HTTP Client](#custom-http-client)
  * [Special Types](#special-types)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/Chalupa-Tech/go-schwab-api-individual
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	schwab "github.com/Chalupa-Tech/go-schwab-api-individual"
	"log"
)

func main() {
	ctx := context.Background()

	s := schwab.New(
		"https://api.example.com",
		schwab.WithSecurity("<YOUR_OAUTH_HERE>"),
	)

	res, err := s.Accounts.GetAccountNumbers(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.AccountNumberHashes != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name    | Type   | Scheme       |
| ------- | ------ | ------------ |
| `Oauth` | oauth2 | OAuth2 token |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	schwab "github.com/Chalupa-Tech/go-schwab-api-individual"
	"log"
)

func main() {
	ctx := context.Background()

	s := schwab.New(
		"https://api.example.com",
		schwab.WithSecurity("<YOUR_OAUTH_HERE>"),
	)

	res, err := s.Accounts.GetAccountNumbers(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.AccountNumberHashes != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Accounts](docs/sdks/accounts/README.md)

* [GetAccountNumbers](docs/sdks/accounts/README.md#getaccountnumbers) - Get list of account numbers and their encrypted values
* [GetAccounts](docs/sdks/accounts/README.md#getaccounts) - Get linked account(s) balances and positions for the logged in user.
* [GetAccount](docs/sdks/accounts/README.md#getaccount) - Get a specific account balance and positions for the logged in user.

### [Instruments](docs/sdks/instruments/README.md)

* [GetInstruments](docs/sdks/instruments/README.md#getinstruments) - Get Instruments by symbols and projections.
* [GetInstrumentsByCusip](docs/sdks/instruments/README.md#getinstrumentsbycusip) - Get Instrument by specific cusip

### [MarketHours](docs/sdks/markethours/README.md)

* [GetMarketHours](docs/sdks/markethours/README.md#getmarkethours) - Get Market Hours for different markets.
* [GetMarketHour](docs/sdks/markethours/README.md#getmarkethour) - Get Market Hours for a single market.

### [Movers](docs/sdks/movers/README.md)

* [GetMovers](docs/sdks/movers/README.md#getmovers) - Get Movers for a specific index.

### [OptionChains](docs/sdks/optionchains/README.md)

* [GetChain](docs/sdks/optionchains/README.md#getchain) - Get option chain for an optionable Symbol

### [OptionExpirationChain](docs/sdks/optionexpirationchain/README.md)

* [GetExpirationChain](docs/sdks/optionexpirationchain/README.md#getexpirationchain) - Get option expiration chain for an optionable symbol

### [Orders](docs/sdks/orders/README.md)

* [GetOrdersByPathParam](docs/sdks/orders/README.md#getordersbypathparam) - Get all orders for a specific account.
* [PlaceOrder](docs/sdks/orders/README.md#placeorder) - Place order for a specific account.
* [GetOrder](docs/sdks/orders/README.md#getorder) - Get a specific order by its ID, for a specific account
* [CancelOrder](docs/sdks/orders/README.md#cancelorder) - Cancel an order for a specific account
* [ReplaceOrder](docs/sdks/orders/README.md#replaceorder) - Replace order for a specific account
* [GetOrdersByQueryParam](docs/sdks/orders/README.md#getordersbyqueryparam) - Get all orders for all accounts
* [PreviewOrder](docs/sdks/orders/README.md#previeworder) - Preview order for a specific account.

### [PriceHistory](docs/sdks/pricehistory/README.md)

* [GetPriceHistory](docs/sdks/pricehistory/README.md#getpricehistory) - Get PriceHistory for a single symbol and date ranges.

### [Quotes](docs/sdks/quotes/README.md)

* [GetQuotes](docs/sdks/quotes/README.md#getquotes) - Get Quotes by list of symbols.
* [GetQuote](docs/sdks/quotes/README.md#getquote) - Get Quote by single symbol.

### [Transactions](docs/sdks/transactions/README.md)

* [GetTransactionsByPathParam](docs/sdks/transactions/README.md#gettransactionsbypathparam) - Get all transactions information for a specific account.
* [GetTransactionsByID](docs/sdks/transactions/README.md#gettransactionsbyid) - Get specific transaction information for a specific account

### [UserPreference](docs/sdks/userpreference/README.md)

* [GetUserPreference](docs/sdks/userpreference/README.md#getuserpreference) - Get user preference information for the logged in user.

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	schwab "github.com/Chalupa-Tech/go-schwab-api-individual"
	"github.com/Chalupa-Tech/go-schwab-api-individual/retry"
	"log"
	"models/operations"
)

func main() {
	ctx := context.Background()

	s := schwab.New(
		"https://api.example.com",
		schwab.WithSecurity("<YOUR_OAUTH_HERE>"),
	)

	res, err := s.Accounts.GetAccountNumbers(ctx, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.AccountNumberHashes != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	schwab "github.com/Chalupa-Tech/go-schwab-api-individual"
	"github.com/Chalupa-Tech/go-schwab-api-individual/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := schwab.New(
		"https://api.example.com",
		schwab.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		schwab.WithSecurity("<YOUR_OAUTH_HERE>"),
	)

	res, err := s.Accounts.GetAccountNumbers(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.AccountNumberHashes != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `GetAccountNumbers` function may return the following errors:

| Error Type             | Status Code        | Content Type     |
| ---------------------- | ------------------ | ---------------- |
| apierrors.ServiceError | 400, 401, 403, 404 | application/json |
| apierrors.ServiceError | 500, 503           | application/json |
| apierrors.APIError     | 4XX, 5XX           | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	schwab "github.com/Chalupa-Tech/go-schwab-api-individual"
	"github.com/Chalupa-Tech/go-schwab-api-individual/models/apierrors"
	"log"
)

func main() {
	ctx := context.Background()

	s := schwab.New(
		"https://api.example.com",
		schwab.WithSecurity("<YOUR_OAUTH_HERE>"),
	)

	res, err := s.Accounts.GetAccountNumbers(ctx)
	if err != nil {

		var e *apierrors.ServiceError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.ServiceError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"

	"github.com/Chalupa-Tech/go-schwab-api-individual"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = schwab.New(schwab.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Special Types [types] -->
## Special Types

This SDK defines the following custom types to assist with marshalling and unmarshalling data.

### Date

`types.Date` is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".

#### Usage

```go
d1 := types.NewDate(time.Now()) // returns *types.Date

d2 := types.DateFromTime(time.Now()) // returns types.Date

d3, err := types.NewDateFromString("2019-01-01") // returns *types.Date, error

d4, err := types.DateFromString("2019-01-01") // returns types.Date, error

d5 := types.MustNewDateFromString("2019-01-01") // returns *types.Date and panics on error

d6 := types.MustDateFromString("2019-01-01") // returns types.Date and panics on error
```
<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Versioning

This SDK follows [Semantic Versioning](https://semver.org/).

- **v0.x.x**: Unstable. Public APIs, types, and schema definitions may change breakingly between minor versions as we refine the `openapi-overlay.yaml` patching strategy.
- **v1.0.0**: Stable. Will mark the first release where the OptionContract unions and critical boolean/number types are considered solidified.

### Schema Management
This SDK uses an **Overlay-First Strategy** to patch the upstream Charles Schwab OpenAPI spec. 
- 📄 See [docs/OVERLAY_STRATEGY.md](docs/OVERLAY_STRATEGY.md) for details on current patches.
- 🛠️ Run `scripts/validate_sdk.sh` to verify changes using the internal [examples/reference_consumer](examples/reference_consumer).

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=go-schwab-api-individual&utm_campaign=go)

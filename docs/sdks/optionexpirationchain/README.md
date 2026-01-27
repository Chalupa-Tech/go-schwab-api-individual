# OptionExpirationChain

## Overview

Get Option Expiration Chain Web Service.

### Available Operations

* [GetExpirationChain](#getexpirationchain) - Get option expiration chain for an optionable symbol

## GetExpirationChain

Get Option Expiration (Series) information for an optionable symbol.  Does not include individual options contracts for the underlying.

### Example Usage

<!-- UsageSnippet language="go" operationID="getExpirationChain" method="get" path="/expirationchain" -->
```go
package main

import(
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

    res, err := s.OptionExpirationChain.GetExpirationChain(ctx, "AAPL")
    if err != nil {
        log.Fatal(err)
    }
    if res.ExpirationChain != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `symbol`                                                 | *string*                                                 | :heavy_check_mark:                                       | Enter one symbol                                         | AAPL                                                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetExpirationChainResponse](../../models/operations/getexpirationchainresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 404           | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |
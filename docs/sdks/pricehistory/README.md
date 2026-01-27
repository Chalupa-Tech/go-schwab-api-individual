# PriceHistory

## Overview

Get Price History Web Service.

### Available Operations

* [GetPriceHistory](#getpricehistory) - Get PriceHistory for a single symbol and date ranges.

## GetPriceHistory

Get historical Open, High, Low, Close, and Volume for a given frequency (i.e. aggregation).  Frequency available is dependent on periodType selected.  The datetime format is in EPOCH milliseconds.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPriceHistory" method="get" path="/pricehistory" -->
```go
package main

import(
	"context"
	schwab "github.com/Chalupa-Tech/go-schwab-api-individual"
	"github.com/Chalupa-Tech/go-schwab-api-individual/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := schwab.New(
        "https://api.example.com",
        schwab.WithSecurity("<YOUR_OAUTH_HERE>"),
    )

    res, err := s.PriceHistory.GetPriceHistory(ctx, operations.GetPriceHistoryRequest{
        Symbol: "AAPL",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CandleList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetPriceHistoryRequest](../../models/operations/getpricehistoryrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetPriceHistoryResponse](../../models/operations/getpricehistoryresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 404           | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |
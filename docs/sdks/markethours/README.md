# MarketHours

## Overview

Get MarketHours Web Service.

### Available Operations

* [GetMarketHours](#getmarkethours) - Get Market Hours for different markets.
* [GetMarketHour](#getmarkethour) - Get Market Hours for a single market.

## GetMarketHours

Get Market Hours for dates in the future across different markets.

### Example Usage

<!-- UsageSnippet language="go" operationID="getMarketHours" method="get" path="/markets" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity("<YOUR_OAUTH_HERE>"),
    )

    res, err := s.MarketHours.GetMarketHours(ctx, []components.QueryParamMarkets{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `markets`                                                                                                                        | [][components.QueryParamMarkets](../../models/components/queryparammarkets.md)                                                   | :heavy_check_mark:                                                                                                               | List of markets                                                                                                                  |
| `date`                                                                                                                           | [*types.Date](../../types/date.md)                                                                                               | :heavy_minus_sign:                                                                                                               | Valid date range is from currentdate to 1 year from today. It will default to current day if not entered. Date format:YYYY-MM-DD |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.GetMarketHoursResponse](../../models/operations/getmarkethoursresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401                | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetMarketHour

Get Market Hours for dates in the future for a single market.

### Example Usage

<!-- UsageSnippet language="go" operationID="getMarketHour" method="get" path="/markets/{market_id}" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity("<YOUR_OAUTH_HERE>"),
    )

    res, err := s.MarketHours.GetMarketHour(ctx, components.PathParamMarketForex, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `marketID`                                                                                                                       | [components.PathParamMarket](../../models/components/pathparammarket.md)                                                         | :heavy_check_mark:                                                                                                               | market id                                                                                                                        |
| `date`                                                                                                                           | [*types.Date](../../types/date.md)                                                                                               | :heavy_minus_sign:                                                                                                               | Valid date range is from currentdate to 1 year from today. It will default to current day if not entered. Date format:YYYY-MM-DD |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.GetMarketHourResponse](../../models/operations/getmarkethourresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 404           | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |
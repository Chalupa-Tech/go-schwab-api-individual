# Movers

## Overview

Get Movers Web Service.

### Available Operations

* [GetMovers](#getmovers) - Get Movers for a specific index.

## GetMovers

Get a list of top 10 securities movement for a specific index.

### Example Usage

<!-- UsageSnippet language="go" operationID="getMovers" method="get" path="/movers/{symbol_id}" -->
```go
package main

import(
	"context"
	schwab "github.com/Chalupa-Tech/go-schwab-api-individual"
	"github.com/Chalupa-Tech/go-schwab-api-individual/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := schwab.New(
        "https://api.example.com",
        schwab.WithSecurity("<YOUR_OAUTH_HERE>"),
    )

    res, err := s.Movers.GetMovers(ctx, components.PathParamSymbolDollarDji, components.QueryParamSortVolume.ToPointer(), components.QueryParamFrequencyZero.ToPointer())
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                         | Type                                                                              | Required                                                                          | Description                                                                       | Example                                                                           |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `ctx`                                                                             | [context.Context](https://pkg.go.dev/context#Context)                             | :heavy_check_mark:                                                                | The context to use for the request.                                               |                                                                                   |
| `symbolID`                                                                        | [components.PathParamSymbol](../../models/components/pathparamsymbol.md)          | :heavy_check_mark:                                                                | Index Symbol                                                                      | $DJI                                                                              |
| `sort`                                                                            | [*components.QueryParamSort](../../models/components/queryparamsort.md)           | :heavy_minus_sign:                                                                | Sort by a particular attribute                                                    | VOLUME                                                                            |
| `frequency`                                                                       | [*components.QueryParamFrequency](../../models/components/queryparamfrequency.md) | :heavy_minus_sign:                                                                | To return movers with the specified directions of up or down                      |                                                                                   |
| `opts`                                                                            | [][operations.Option](../../models/operations/option.md)                          | :heavy_minus_sign:                                                                | The options for this request.                                                     |                                                                                   |

### Response

**[*operations.GetMoversResponse](../../models/operations/getmoversresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 404           | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |
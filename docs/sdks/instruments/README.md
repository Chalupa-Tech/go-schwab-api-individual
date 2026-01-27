# Instruments

## Overview

Get Instruments Web Service.

### Available Operations

* [GetInstruments](#getinstruments) - Get Instruments by symbols and projections.
* [GetInstrumentsByCusip](#getinstrumentsbycusip) - Get Instrument by specific cusip

## GetInstruments

Get Instruments details by using different projections.  Get more specific fundamental instrument data by using fundamental as the projection.

### Example Usage

<!-- UsageSnippet language="go" operationID="getInstruments" method="get" path="/instruments" -->
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

    res, err := s.Instruments.GetInstruments(ctx, "<value>", components.QueryParamprojectionSymbolRegex)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `symbol`                                                                           | *string*                                                                           | :heavy_check_mark:                                                                 | symbol of a security                                                               |
| `projection`                                                                       | [components.QueryParamprojection](../../models/components/queryparamprojection.md) | :heavy_check_mark:                                                                 | search by                                                                          |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.GetInstrumentsResponse](../../models/operations/getinstrumentsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401                | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetInstrumentsByCusip

Get basic instrument details by cusip

### Example Usage

<!-- UsageSnippet language="go" operationID="getInstrumentsByCusip" method="get" path="/instruments/{cusip_id}" -->
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

    res, err := s.Instruments.GetInstrumentsByCusip(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.InstrumentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `cusipID`                                                | *string*                                                 | :heavy_check_mark:                                       | cusip of a security                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetInstrumentsByCusipResponse](../../models/operations/getinstrumentsbycusipresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 404           | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |
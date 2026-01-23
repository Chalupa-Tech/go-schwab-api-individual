# Quotes

## Overview

Get Quotes Web Service.

### Available Operations

* [GetQuotes](#getquotes) - Get Quotes by list of symbols.
* [GetQuote](#getquote) - Get Quote by single symbol.

## GetQuotes

Get Quotes by list of symbols.

### Example Usage

<!-- UsageSnippet language="go" operationID="getQuotes" method="get" path="/quotes" -->
```go
package main

import(
	"context"
	"undefined"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity("<YOUR_OAUTH_HERE>"),
    )

    res, err := s.Quotes.GetQuotes(ctx, undefined.Pointer("MRAD,EATOF,EBIZ,AAPL,BAC,AAAHX,AAAIX,$DJI,$SPX,MVEN,SOBS,TOITF,CNSWF,AMZN  230317C01360000,DJX   231215C00290000,/ESH23,./ADUF23C0.55,AUD/CAD"), undefined.Pointer("quote,reference"), undefined.Pointer(false))
    if err != nil {
        log.Fatal(err)
    }
    if res.QuoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                  | Type                                                                                                                                                                                                                                                                                       | Required                                                                                                                                                                                                                                                                                   | Description                                                                                                                                                                                                                                                                                | Example                                                                                                                                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                                                                                         | The context to use for the request.                                                                                                                                                                                                                                                        |                                                                                                                                                                                                                                                                                            |
| `symbols`                                                                                                                                                                                                                                                                                  | **string*                                                                                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                                                                         | Comma separated list of symbol(s) to look up a quote                                                                                                                                                                                                                                       | MRAD,EATOF,EBIZ,AAPL,BAC,AAAHX,AAAIX,$DJI,$SPX,MVEN,SOBS,TOITF,CNSWF,AMZN  230317C01360000,DJX   231215C00290000,/ESH23,./ADUF23C0.55,AUD/CAD                                                                                                                                              |
| `fields`                                                                                                                                                                                                                                                                                   | **string*                                                                                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                                                                         | Request for subset of data by passing coma separated list of root nodes, possible root nodes are quote, fundamental, extended, reference, regular. Sending `quote, fundamental` in request will return quote and fundamental data in response. Dont send this attribute for full response. | quote,reference                                                                                                                                                                                                                                                                            |
| `indicative`                                                                                                                                                                                                                                                                               | **bool*                                                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                         | Include indicative symbol quotes for all ETF symbols in request. If ETF symbol ABC is in request and indicative=true API will return quotes for ABC and its corresponding indicative quote for $ABC.IV                                                                                     | false                                                                                                                                                                                                                                                                                      |
| `opts`                                                                                                                                                                                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                                                         | The options for this request.                                                                                                                                                                                                                                                              |                                                                                                                                                                                                                                                                                            |

### Response

**[*operations.GetQuotesResponse](../../models/operations/getquotesresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401                | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetQuote

Get Quote by single symbol.

### Example Usage

<!-- UsageSnippet language="go" operationID="getQuote" method="get" path="/{symbol_id}/quotes" -->
```go
package main

import(
	"context"
	"undefined"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity("<YOUR_OAUTH_HERE>"),
    )

    res, err := s.Quotes.GetQuote(ctx, "TSLA", undefined.Pointer("quote,reference"))
    if err != nil {
        log.Fatal(err)
    }
    if res.QuoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                  | Type                                                                                                                                                                                                                                                                                       | Required                                                                                                                                                                                                                                                                                   | Description                                                                                                                                                                                                                                                                                | Example                                                                                                                                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                                                                                         | The context to use for the request.                                                                                                                                                                                                                                                        |                                                                                                                                                                                                                                                                                            |
| `symbolID`                                                                                                                                                                                                                                                                                 | *string*                                                                                                                                                                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                                                                                                                         | Symbol of instrument                                                                                                                                                                                                                                                                       | TSLA                                                                                                                                                                                                                                                                                       |
| `fields`                                                                                                                                                                                                                                                                                   | **string*                                                                                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                                                                         | Request for subset of data by passing coma separated list of root nodes, possible root nodes are quote, fundamental, extended, reference, regular. Sending `quote, fundamental` in request will return quote and fundamental data in response. Dont send this attribute for full response. | quote,reference                                                                                                                                                                                                                                                                            |
| `opts`                                                                                                                                                                                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                                                                                                                         | The options for this request.                                                                                                                                                                                                                                                              |                                                                                                                                                                                                                                                                                            |

### Response

**[*operations.GetQuoteResponse](../../models/operations/getquoteresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 404           | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |
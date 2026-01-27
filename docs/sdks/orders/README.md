# Orders

## Overview

### Available Operations

* [GetOrdersByPathParam](#getordersbypathparam) - Get all orders for a specific account.
* [PlaceOrder](#placeorder) - Place order for a specific account.
* [GetOrder](#getorder) - Get a specific order by its ID, for a specific account
* [CancelOrder](#cancelorder) - Cancel an order for a specific account
* [ReplaceOrder](#replaceorder) - Replace order for a specific account
* [GetOrdersByQueryParam](#getordersbyqueryparam) - Get all orders for all accounts
* [PreviewOrder](#previeworder) - Preview order for a specific account.

## GetOrdersByPathParam

All orders for a specific account. Orders retrieved can be filtered based on input parameters below. Maximum date range is 1 year.

### Example Usage

<!-- UsageSnippet language="go" operationID="getOrdersByPathParam" method="get" path="/accounts/{accountNumber}/orders" -->
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

    res, err := s.Orders.GetOrdersByPathParam(ctx, operations.GetOrdersByPathParamRequest{
        AccountNumber: "<value>",
        FromEnteredTime: "<value>",
        ToEnteredTime: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Orders != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetOrdersByPathParamRequest](../../models/operations/getordersbypathparamrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetOrdersByPathParamResponse](../../models/operations/getordersbypathparamresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.ServiceError | 400, 401, 403, 404     | application/json       |
| apierrors.ServiceError | 500, 503               | application/json       |
| apierrors.APIError     | 4XX, 5XX               | \*/\*                  |

## PlaceOrder

Place an order for a specific account.

### Example Usage

<!-- UsageSnippet language="go" operationID="placeOrder" method="post" path="/accounts/{accountNumber}/orders" -->
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

    res, err := s.Orders.PlaceOrder(ctx, "<value>", components.OrderRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `accountNumber`                                                    | *string*                                                           | :heavy_check_mark:                                                 | The encrypted ID of the account                                    |
| `body`                                                             | [components.OrderRequest](../../models/components/orderrequest.md) | :heavy_check_mark:                                                 | The new Order Object.                                              |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.PlaceOrderResponse](../../models/operations/placeorderresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.ServiceError | 400, 401, 403, 404     | application/json       |
| apierrors.ServiceError | 500, 503               | application/json       |
| apierrors.APIError     | 4XX, 5XX               | \*/\*                  |

## GetOrder

Get a specific order by its ID, for a specific account

### Example Usage

<!-- UsageSnippet language="go" operationID="getOrder" method="get" path="/accounts/{accountNumber}/orders/{orderId}" -->
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

    res, err := s.Orders.GetOrder(ctx, "<value>", 550049)
    if err != nil {
        log.Fatal(err)
    }
    if res.Order != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `accountNumber`                                          | *string*                                                 | :heavy_check_mark:                                       | The encrypted ID of the account                          |
| `orderID`                                                | *int64*                                                  | :heavy_check_mark:                                       | The ID of the order being retrieved.                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetOrderResponse](../../models/operations/getorderresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.ServiceError | 400, 401, 403, 404     | application/json       |
| apierrors.ServiceError | 500, 503               | application/json       |
| apierrors.APIError     | 4XX, 5XX               | \*/\*                  |

## CancelOrder

Cancel a specific order for a specific account<br>

### Example Usage

<!-- UsageSnippet language="go" operationID="cancelOrder" method="delete" path="/accounts/{accountNumber}/orders/{orderId}" -->
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

    res, err := s.Orders.CancelOrder(ctx, "<value>", 325883)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `accountNumber`                                          | *string*                                                 | :heavy_check_mark:                                       | The encrypted ID of the account                          |
| `orderID`                                                | *int64*                                                  | :heavy_check_mark:                                       | The ID of the order being cancelled                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CancelOrderResponse](../../models/operations/cancelorderresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.ServiceError | 400, 401, 403, 404     | application/json       |
| apierrors.ServiceError | 500, 503               | application/json       |
| apierrors.APIError     | 4XX, 5XX               | \*/\*                  |

## ReplaceOrder

Replace an existing order for an account. The existing order will be replaced by the new               order. Once replaced, the old order will be canceled and a new order will be created.

### Example Usage

<!-- UsageSnippet language="go" operationID="replaceOrder" method="put" path="/accounts/{accountNumber}/orders/{orderId}" -->
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

    res, err := s.Orders.ReplaceOrder(ctx, "<value>", 786311, components.OrderRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `accountNumber`                                                    | *string*                                                           | :heavy_check_mark:                                                 | The encrypted ID of the account                                    |
| `orderID`                                                          | *int64*                                                            | :heavy_check_mark:                                                 | The ID of the order being retrieved.                               |
| `body`                                                             | [components.OrderRequest](../../models/components/orderrequest.md) | :heavy_check_mark:                                                 | The Order Object.                                                  |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.ReplaceOrderResponse](../../models/operations/replaceorderresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.ServiceError | 400, 401, 403, 404     | application/json       |
| apierrors.ServiceError | 500, 503               | application/json       |
| apierrors.APIError     | 4XX, 5XX               | \*/\*                  |

## GetOrdersByQueryParam

Get all orders for all accounts<br>

### Example Usage

<!-- UsageSnippet language="go" operationID="getOrdersByQueryParam" method="get" path="/orders" -->
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

    res, err := s.Orders.GetOrdersByQueryParam(ctx, "<value>", "<value>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Orders != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                     | Type                                                                                                                                                                                                          | Required                                                                                                                                                                                                      | Description                                                                                                                                                                                                   |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                                                            | The context to use for the request.                                                                                                                                                                           |
| `fromEnteredTime`                                                                                                                                                                                             | *string*                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                            | Specifies that no orders entered before this time should be returned. Valid ISO-8601 formats are-<br/>yyyy-MM-dd'T'HH:mm:ss.SSSZ Date must be within 60 days from today's date.<br/>'toEnteredTime' must also be set. |
| `toEnteredTime`                                                                                                                                                                                               | *string*                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                            | Specifies that no orders entered after this time should be returned.Valid ISO-8601 formats are -<br/>yyyy-MM-dd'T'HH:mm:ss.SSSZ. 'fromEnteredTime' must also be set.                                          |
| `maxResults`                                                                                                                                                                                                  | **int64*                                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                            | The max number of orders to retrieve. Default is 3000.                                                                                                                                                        |
| `status`                                                                                                                                                                                                      | [*components.APIOrderStatus](../../models/components/apiorderstatus.md)                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                            | Specifies that only orders of this status should be returned.                                                                                                                                                 |
| `opts`                                                                                                                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                            | The options for this request.                                                                                                                                                                                 |

### Response

**[*operations.GetOrdersByQueryParamResponse](../../models/operations/getordersbyqueryparamresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.ServiceError | 400, 401, 403, 404     | application/json       |
| apierrors.ServiceError | 500, 503               | application/json       |
| apierrors.APIError     | 4XX, 5XX               | \*/\*                  |

## PreviewOrder

Preview an order for a specific account.

### Example Usage

<!-- UsageSnippet language="go" operationID="previewOrder" method="post" path="/accounts/{accountNumber}/previewOrder" -->
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

    res, err := s.Orders.PreviewOrder(ctx, "<value>", components.PreviewOrder{})
    if err != nil {
        log.Fatal(err)
    }
    if res.PreviewOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `accountNumber`                                                    | *string*                                                           | :heavy_check_mark:                                                 | The encrypted ID of the account                                    |
| `body`                                                             | [components.PreviewOrder](../../models/components/previeworder.md) | :heavy_check_mark:                                                 | The Order Object.                                                  |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.PreviewOrderResponse](../../models/operations/previeworderresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.ServiceError | 400, 401, 403, 404     | application/json       |
| apierrors.ServiceError | 500, 503               | application/json       |
| apierrors.APIError     | 4XX, 5XX               | \*/\*                  |
# Transactions

## Overview

### Available Operations

* [GetTransactionsByPathParam](#gettransactionsbypathparam) - Get all transactions information for a specific account.
* [GetTransactionsByID](#gettransactionsbyid) - Get specific transaction information for a specific account

## GetTransactionsByPathParam

All transactions for a specific account. Maximum number of transactions in response is 3000. Maximum date range is 1 year.

### Example Usage

<!-- UsageSnippet language="go" operationID="getTransactionsByPathParam" method="get" path="/accounts/{accountNumber}/transactions" -->
```go
package main

import(
	"context"
	schwab "github.com/Chalupa-Tech/go-schwab-api-individual"
	"github.com/Chalupa-Tech/go-schwab-api-individual/models/components"
	"github.com/Chalupa-Tech/go-schwab-api-individual/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := schwab.New(
        "https://api.example.com",
        schwab.WithSecurity("<YOUR_OAUTH_HERE>"),
    )

    res, err := s.Transactions.GetTransactionsByPathParam(ctx, operations.GetTransactionsByPathParamRequest{
        AccountNumber: "<value>",
        StartDate: "<value>",
        EndDate: "<value>",
        Types: components.TransactionTypeCashDisbursement,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Transactions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetTransactionsByPathParamRequest](../../models/operations/gettransactionsbypathparamrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.GetTransactionsByPathParamResponse](../../models/operations/gettransactionsbypathparamresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.ServiceError | 400, 401, 403, 404     | application/json       |
| apierrors.ServiceError | 500, 503               | application/json       |
| apierrors.APIError     | 4XX, 5XX               | \*/\*                  |

## GetTransactionsByID

Get specific transaction information for a specific account

### Example Usage

<!-- UsageSnippet language="go" operationID="getTransactionsById" method="get" path="/accounts/{accountNumber}/transactions/{transactionId}" -->
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

    res, err := s.Transactions.GetTransactionsByID(ctx, "<value>", 923773)
    if err != nil {
        log.Fatal(err)
    }
    if res.Transactions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `accountNumber`                                          | *string*                                                 | :heavy_check_mark:                                       | The encrypted ID of the account                          |
| `transactionID`                                          | *int64*                                                  | :heavy_check_mark:                                       | The ID of the transaction being retrieved.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetTransactionsByIDResponse](../../models/operations/gettransactionsbyidresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.ServiceError | 400, 401, 403, 404     | application/json       |
| apierrors.ServiceError | 500, 503               | application/json       |
| apierrors.APIError     | 4XX, 5XX               | \*/\*                  |
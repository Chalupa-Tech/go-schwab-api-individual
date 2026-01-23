# Accounts

## Overview

### Available Operations

* [GetAccountNumbers](#getaccountnumbers) - Get list of account numbers and their encrypted values
* [GetAccounts](#getaccounts) - Get linked account(s) balances and positions for the logged in user.
* [GetAccount](#getaccount) - Get a specific account balance and positions for the logged in user.

## GetAccountNumbers

Account numbers in plain text cannot be used outside of headers or request/response bodies. As the first step consumers must invoke this service to retrieve the list of plain text/encrypted value pairs, and use encrypted account values for all subsequent calls for any accountNumber request.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAccountNumbers" method="get" path="/accounts/accountNumbers" -->
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

    res, err := s.Accounts.GetAccountNumbers(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountNumberHashes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetAccountNumbersResponse](../../models/operations/getaccountnumbersresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.ServiceError | 400, 401, 403, 404     | application/json       |
| apierrors.ServiceError | 500, 503               | application/json       |
| apierrors.APIError     | 4XX, 5XX               | \*/\*                  |

## GetAccounts

All the linked account information for the user logged in. The
balances on these accounts are displayed by default however the positions
on these accounts will be displayed based on the "positions" flag.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAccounts" method="get" path="/accounts" -->
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

    res, err := s.Accounts.GetAccounts(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Accounts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                    | Type                                                                                                                                                                         | Required                                                                                                                                                                     | Description                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                        | :heavy_check_mark:                                                                                                                                                           | The context to use for the request.                                                                                                                                          |
| `fields`                                                                                                                                                                     | **string*                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                           | This allows one to determine which fields they want returned. Possible value in this String can be:<br/><br><code>positions</code><br> Example:<br><code>fields=positions</code> |
| `opts`                                                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                                                     | :heavy_minus_sign:                                                                                                                                                           | The options for this request.                                                                                                                                                |

### Response

**[*operations.GetAccountsResponse](../../models/operations/getaccountsresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.ServiceError | 400, 401, 403, 404     | application/json       |
| apierrors.ServiceError | 500, 503               | application/json       |
| apierrors.APIError     | 4XX, 5XX               | \*/\*                  |

## GetAccount

Specific account information with balances and positions.
The balance information on these accounts is displayed by default but
Positions will be returned based on the "positions" flag.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAccount" method="get" path="/accounts/{accountNumber}" -->
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

    res, err := s.Accounts.GetAccount(ctx, "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Account != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                     | Type                                                                                                                                                                          | Required                                                                                                                                                                      | Description                                                                                                                                                                   |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                         | :heavy_check_mark:                                                                                                                                                            | The context to use for the request.                                                                                                                                           |
| `accountNumber`                                                                                                                                                               | *string*                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                            | The encrypted ID of the account                                                                                                                                               |
| `fields`                                                                                                                                                                      | **string*                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                            | This allows one to determine<br/>which fields they want returned. Possible values in this String can be:<br/><br><code>positions</code><br> Example:<br><code>fields=positions</code> |
| `opts`                                                                                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                                                                                      | :heavy_minus_sign:                                                                                                                                                            | The options for this request.                                                                                                                                                 |

### Response

**[*operations.GetAccountResponse](../../models/operations/getaccountresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.ServiceError | 400, 401, 403, 404     | application/json       |
| apierrors.ServiceError | 500, 503               | application/json       |
| apierrors.APIError     | 4XX, 5XX               | \*/\*                  |
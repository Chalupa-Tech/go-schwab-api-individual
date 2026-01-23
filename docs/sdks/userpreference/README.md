# UserPreference

## Overview

### Available Operations

* [GetUserPreference](#getuserpreference) - Get user preference information for the logged in user.

## GetUserPreference

Get user preference information for the logged in user.

### Example Usage

<!-- UsageSnippet language="go" operationID="getUserPreference" method="get" path="/userPreference" -->
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

    res, err := s.UserPreference.GetUserPreference(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.UserPreferences != nil {
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

**[*operations.GetUserPreferenceResponse](../../models/operations/getuserpreferenceresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.ServiceError | 400, 401, 403, 404     | application/json       |
| apierrors.ServiceError | 500, 503               | application/json       |
| apierrors.APIError     | 4XX, 5XX               | \*/\*                  |
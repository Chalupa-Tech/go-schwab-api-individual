<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"log"
	"undefined"
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
<!-- End SDK Example Usage [usage] -->
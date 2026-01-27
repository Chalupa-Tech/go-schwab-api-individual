# QuoteError

Partial or Custom errors per request


## Fields

| Field                                | Type                                 | Required                             | Description                          |
| ------------------------------------ | ------------------------------------ | ------------------------------------ | ------------------------------------ |
| `InvalidCusips`                      | []*string*                           | :heavy_minus_sign:                   | list of invalid cusips from request  |
| `InvalidSSIDs`                       | []*int64*                            | :heavy_minus_sign:                   | list of invalid SSIDs from request   |
| `InvalidSymbols`                     | []*string*                           | :heavy_minus_sign:                   | list of invalid symbols from request |
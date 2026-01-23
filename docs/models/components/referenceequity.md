# ReferenceEquity

Reference data of Equity security


## Fields

| Field                       | Type                        | Required                    | Description                 | Example                     |
| --------------------------- | --------------------------- | --------------------------- | --------------------------- | --------------------------- |
| `Cusip`                     | **string*                   | :heavy_minus_sign:          | CUSIP of Instrument         | A23456789                   |
| `Description`               | **string*                   | :heavy_minus_sign:          | Description of Instrument   | Apple Inc. - Common Stock   |
| `Exchange`                  | **string*                   | :heavy_minus_sign:          | Exchange Code               | q                           |
| `ExchangeName`              | **string*                   | :heavy_minus_sign:          | Exchange Name               |                             |
| `FsiDesc`                   | **string*                   | :heavy_minus_sign:          | FSI Desc                    |                             |
| `HtbQuantity`               | **int*                      | :heavy_minus_sign:          | Hard to borrow quantity.    | 100                         |
| `HtbRate`                   | **float64*                  | :heavy_minus_sign:          | Hard to borrow rate.        | 4.5                         |
| `IsHardToBorrow`            | **bool*                     | :heavy_minus_sign:          | is Hard to borrow security. | false                       |
| `IsShortable`               | **bool*                     | :heavy_minus_sign:          | is shortable security.      | false                       |
| `OtcMarketTier`             | **string*                   | :heavy_minus_sign:          | OTC Market Tier             |                             |
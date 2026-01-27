# GetPriceHistoryResponse


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `HTTPMeta`                                                         | [components.HTTPMetadata](../../models/components/httpmetadata.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `CandleList`                                                       | [*components.CandleList](../../models/components/candlelist.md)    | :heavy_minus_sign:                                                 | Get all candles for given date range                               |
| `Headers`                                                          | map[string][]*string*                                              | :heavy_check_mark:                                                 | N/A                                                                |
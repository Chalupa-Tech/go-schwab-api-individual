# ReplaceOrderRequest


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `AccountNumber`                                                    | *string*                                                           | :heavy_check_mark:                                                 | The encrypted ID of the account                                    |
| `OrderID`                                                          | *int64*                                                            | :heavy_check_mark:                                                 | The ID of the order being retrieved.                               |
| `Body`                                                             | [components.OrderRequest](../../models/components/orderrequest.md) | :heavy_check_mark:                                                 | The Order Object.                                                  |
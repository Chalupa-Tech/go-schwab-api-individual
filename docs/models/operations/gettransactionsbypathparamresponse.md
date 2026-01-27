# GetTransactionsByPathParamResponse


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `HTTPMeta`                                                               | [components.HTTPMetadata](../../models/components/httpmetadata.md)       | :heavy_check_mark:                                                       | N/A                                                                      |
| `Transactions`                                                           | [][components.Transaction](../../models/components/transaction.md)       | :heavy_minus_sign:                                                       | A List of orders for the account, matching the provided input<br/>parameters |
| `Headers`                                                                | map[string][]*string*                                                    | :heavy_check_mark:                                                       | N/A                                                                      |
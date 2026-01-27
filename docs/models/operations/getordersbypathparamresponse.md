# GetOrdersByPathParamResponse


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `HTTPMeta`                                                               | [components.HTTPMetadata](../../models/components/httpmetadata.md)       | :heavy_check_mark:                                                       | N/A                                                                      |
| `Orders`                                                                 | [][components.Order](../../models/components/order.md)                   | :heavy_minus_sign:                                                       | A List of orders for the account, matching the provided input parameters |
| `Headers`                                                                | map[string][]*string*                                                    | :heavy_check_mark:                                                       | N/A                                                                      |
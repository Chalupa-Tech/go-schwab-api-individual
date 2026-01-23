# GetAccountsResponse


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `HTTPMeta`                                                         | [components.HTTPMetadata](../../models/components/httpmetadata.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `Accounts`                                                         | [][components.Account](../../models/components/account.md)         | :heavy_minus_sign:                                                 | List of valid "accounts", matching the provided input parameters.  |
| `Headers`                                                          | map[string][]*string*                                              | :heavy_check_mark:                                                 | N/A                                                                |
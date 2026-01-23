# Screener

Security info of most moved with in an index


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Change`                                                      | **float64*                                                    | :heavy_minus_sign:                                            | percent or value changed, by default its percent changed      |
| `Description`                                                 | **string*                                                     | :heavy_minus_sign:                                            | Name of security                                              |
| `Direction`                                                   | [*components.Direction](../../models/components/direction.md) | :heavy_minus_sign:                                            | N/A                                                           |
| `Last`                                                        | **float64*                                                    | :heavy_minus_sign:                                            | what was last quoted price                                    |
| `Symbol`                                                      | **string*                                                     | :heavy_minus_sign:                                            | schwab security symbol                                        |
| `TotalVolume`                                                 | **int64*                                                      | :heavy_minus_sign:                                            | N/A                                                           |
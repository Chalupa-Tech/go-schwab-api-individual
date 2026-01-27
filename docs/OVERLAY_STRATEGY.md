# Overlay-First Strategy

## Purpose
The Charles Schwab Trader API OpenAPI specification contains critical schema errors and type mismatches that prevent the generation of a functional Go SDK. To address this without manually patching generated code (which is brittle and hard to maintain), we use an **Overlay-First Strategy**.

This strategy involves applying a standard `openapi-overlay.yaml` to the upstream `openapi.yaml` before code generation. This ensures that:
1.  The upstream spec generally remains the source of truth.
2.  Our fixes are explicit, versioned, and reproducible.
3.  We can easily audit what we have changed.

## Workflow
The code generation process is automated via `speakeasy`:

1.  **Input**: `openapi.yaml` (Upstream Spec) + `openapi-overlay.yaml` (Our Patches)
2.  **Process**: `speakeasy overlay apply -s openapi.yaml -o openapi-overlay.yaml > openapi-modified.yaml`
3.  **Generation**: Speakeasy generates the SDK from `openapi-modified.yaml`.

> [!WARNING]
> **NEVER** edit `openapi.yaml` or `openapi-modified.yaml` directly.
> *   If you edit `openapi.yaml`, your changes will be lost when we update from upstream.
> *   If you edit `openapi-modified.yaml`, your changes will be lost when we re-run the overlay command.
>
> **ALWAYS** add a new action to `openapi-overlay.yaml`.

## Patch Catalog
The following table lists the patches currently applied in `openapi-overlay.yaml`.

| Patch ID | Target Component | Issue Description | Fix Applied |
| :--- | :--- | :--- | :--- |
| **1** | `OptionDeliverables` | `deliverableUnits` is defined as wrong type/missing format. | Enforced `type: number`, `format: double`. |
| **2** | `OptionContractObject` | `OptionContract` is a massive union but needs a base object definition. | Created a concrete `OptionContractObject` schema with all fields. |
| **3** | `OptionContract` | `OptionContract` was not correctly defining its polymorphic nature. | Converted to `oneOf` targeting `OptionContractObject` (single) vs Array of Objects. |
| **4-5** | `Servers` & Paths | Incorrect base URLs for Market Data vs Trader API. | Defined explicit servers and routed `/quotes`, `/chains`, etc., to `marketdata/v1`. |
| **6** | `Future` / `Index` | Circular dependency in `allOf` references causing panic. | Replaced circular `$ref` with strict `#/components/schemas/TransactionBaseInstrument`. |
| **7** | `AccountEquity/MutualFund` | Go struct naming collision. | Added `_type` discriminator property to force unique struct generation. |
| **8** | `isInCall` | Fields defined as `string` or `number` but behave as `boolean`. | Forced `type: boolean` on `Margin` and `Cash` balance objects. |

## Adding a New Patch
To fix a new issue:
1.  Identify the JSONPath to the problematic node in `openapi.yaml`.
2.  Create a new entry in `openapi-overlay.yaml` under `actions`.
3.  Run the generation script (or `speakeasy overlay apply` manually) to verify the `openapi-modified.yaml` output.
4.  Regenerate the SDK and test.

### Example
```yaml
  - target: $.components.schemas.SomeBrokenField.type
    update: boolean
```

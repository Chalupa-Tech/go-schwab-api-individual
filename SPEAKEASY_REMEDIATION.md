# Speakeasy SDK Remediation Guide

This document outlines the changes required to the source OpenAPI specification or Speakeasy configuration to permanently fix recurring issues in the generated Go SDK.

Current manual patches are applied to:
- `models/components/optioncontract.go`
- `models/components/optiondeliverables.go`

## 1. OptionContract: JSON Array Wrapping

**Issue:**
The Schwab API endpoint for Option Chains occasionally returns the `OptionContract` object wrapped in a single-element array (e.g., `[{ "symbol": "..." }]`) instead of a direct object (e.g., `{ "symbol": "..." }`). The current generated SDK expects a direct object and fails with a JSON unmarshal error.

**Remediation (OpenAPI/Speakeasy):**
To handle this purely through generation, you have two options:

### Option A: Speakeasy Overlay (Recommended)
Use an `openapi-overlay.yaml` to modify the schema for `OptionContract` to accept `oneOf` object or array of objects, OR use a custom unmarshaller.

Better yet, since this is likely a quirk of the API not matching the spec:
1.  Identify the operation ID returning `OptionContract`.
2.  Use a type override in `speakeasy.yaml` or an overlay to define a custom unmarshaller template for this model.

### Option B: Fix the Spec
If the API formally defines this as "sometimes an array", update the OpenAPI spec:
```yaml
components:
  schemas:
    OptionContract:
      oneOf:
        - $ref: '#/components/schemas/OptionContractObject'
        - type: array
          items:
            $ref: '#/components/schemas/OptionContractObject'
```
*Note: This drastically changes the generated API surface. A custom unmarshaller in the Go template is preferred to keep the struct clean.*

**Current Manual Fix (Reference):**
The manual fix adds a custom `UnmarshalJSON` method to `OptionContract` that checks for a leading `[` and unwraps the first element.

## 2. OptionDeliverables: Type Mismatch

**Issue:**
The property `deliverableUnits` in `OptionDeliverables` is defined as a `string` in the OpenAPI spec, but the API returns a `number` (float).
- **Spec:** `"deliverableUnits": { "type": "string" }`
- **API:** `"deliverableUnits": 100.0`

**Remediation:**
Correct the type definition in the OpenAPI spec using an overlay.

**Create `openapi-overlay.yaml`:**
```yaml
overlay: 1.0.0
info:
  title: Fix OptionDeliverables Type
  version: 1.0.0
actions:
  - target: $.components.schemas.OptionDeliverables.properties.deliverableUnits
    update:
      type: number
      format: double
```

## Summary for Orchestrator
When automating this fix:
1.  Apply the `openapi-overlay.yaml` to the source Spec.
2.  Regenerate the SDK using `speakeasy generate`.
3.  Verify `models/components/optiondeliverables.go` uses `*float64`.

# ADR 001: Reference Consumer Pattern & SDK Decoupling

**Title:** Reference Consumer Pattern & SDK Decoupling
**Status:** Proposed
**Date:** 2026-01-27
**Author:** Gemini Orchestrator (Architect Persona)

## Context
We are managing a Speakeasy-generated Go SDK (`go-schwab-api-individual`) and a private consumer (`tayvens-stock-report`).
Currently, validation of the SDK is performed by linking it against the consumer via a local `replace` directive in `go.mod`.
This approach has several significant drawbacks:
1.  **Coupling**: The SDK build process depends on the presence and state of a separate repository (`tayvens-stock-report`).
2.  **Fragility**: CI/CD pipelines require complex setup to checkout multiple private repos and patch `replace` directives.
3.  **Leakage**: The consumer's private business logic is unnecessarily exposed to the SDK's validation process.

We considered `schwab-py`'s approach of writing unit tests for the generated code but rejected it. Generated code is an implementation detail; writing unit tests for it creates a maintenance burden where we must update tests every time the generator changes its output structure.

## Decision
We will decouple the SDK from the private consumer by adopting a **Reference Consumer** pattern and a **Private Module** distribution strategy.

### 1. Reference Consumer Pattern (Internal Validation)
Instead of unit tests or external consumer linking, we will implement a lightweight `examples/reference_consumer` application *inside* the SDK repository.
*   **Role**: Serves as the primary "Integration Test" suite.
*   **Scope**: Validates that critical, complex types (e.g., `OptionContract` unions, boolean helpers) compile and run correctly against the SDK's public API.
*   **Benefit**: The SDK repository becomes self-validating. `go build ./examples/...` proves the SDK is usable.

### 2. Private Module Distribution (External Usage)
Consumers like `tayvens-stock-report` will consume the SDK as a standard Go module, authenticated via git tags and `GOPRIVATE`.
*   **Mechanism**: Releases are defined by git tags (e.g., `v0.1.0`).
*   **Configuration**: Consumers must set `go env -w GOPRIVATE=github.com/Chalupa-Tech/*` and configure git to use SSH or PATs for authentication.
*   **Reversibility**: This decision is not permanent. If we choose to open-source the SDK later, we simply stop requiring `GOPRIVATE`.

## Consequences

### Positive
*   **Self-Contained**: SDK CI checks can pass independently of any other repo.
*   **Standardization**: Consumers use standard `go get` workflows instead of `replace` hacks.
*   **Stability**: The "Reference Consumer" acts as a canary for breaking changes in the generated code.

### Negative
*   **Setup Friction**: Developers must configure `GOPRIVATE` and git authentication locally.
*   **CI Complexity**: Consumer CI pipelines must inject credentials (PATs) to fetch the private SDK module.

## Compliance
*   **Source of Truth**: This ADR resides in `go-schwab-api-individual`, establishing that repo as the single source of truth for its own architectural decisions.

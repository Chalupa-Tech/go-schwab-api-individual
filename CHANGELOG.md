# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- **Standards**: Added `Makefile`, `.vscode/settings.json`, `PR Template`.
- **CI**: Added `.github/workflows/ci.yml` and `.revive.toml`.

### Fixed
- **CI**: Removed legacy `StartYourLab` git URL rewrite rule ([#](https://github.com/Chalupa-Tech/go-schwab-api-individual/commit/HEAD)).

## [0.1.0] - 2026-01-30
### Added
- **New SDK**: Regenerated SDK using Speakeasy to match current Schwab API spec.
- **Hooks**: Implemented `URLRewriteHook` to handle `getPriceHistory` and `getChain` URL path corrections (/trader/v1 -> /marketdata/v1).
### Added
- **Docs**: Added `docs/OVERLAY_STRATEGY.md` formalizing the spec patching process.
- **Scripts**: Added `scripts/validate_consumer.sh` for downstream integration testing.
- **CI**: Added `.github/workflows/sdk-validation.yaml` validation workflow.
- **Versioning**: Added Versioning strategy to README.

### Changed
- **Validation**: Added standardized `.agent/validate.sh` scaffolding.

### Fixed
- Fixed critical generator bug (`package undefined`).
- Fixed broken internal imports globally.

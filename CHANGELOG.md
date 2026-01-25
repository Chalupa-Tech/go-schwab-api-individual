# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Standardized `.agent/validate.sh` script to verify build, formatting, and tests.

### Fixed
- Fixed critical generator bug: `package undefined` in `optionchains.go` and `pricehistory.go` is now `package sdk`.
- Fixed broken internal imports from `undefined/...` to correct module path.

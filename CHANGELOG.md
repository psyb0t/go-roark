# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.1.0, 2026-10-17

- Added `roark`, a REST client generated with oapi-codegen from the Roark Analytics API specification, version 1.0.0, OpenAPI 3.1.0, published at https://api.roark.ai/doc. It covers all 107 paths and 160 operations for voice AI testing and monitoring.
- Number fields are `float64` instead of oapi-codegen's default `float32`.
- `make generate` downloads the current spec, downconverts its OpenAPI 3.1 schema features and regenerates the client. The spec is not committed.
- Added local HTTP tests, development and security tooling, CI workflows and the `go-roark` agent skill.

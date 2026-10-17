// Package roark is a Go client for the Roark API, generated from its
// public OpenAPI document with oapi-codegen.
package roark

//go:generate bash scripts/fetch-spec.sh
//go:generate bash scripts/patch-spec.sh
//go:generate go tool oapi-codegen --config=config.yaml api/openapi.fixed.json

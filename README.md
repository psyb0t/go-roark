# go-roark

[![Go Reference](https://pkg.go.dev/badge/github.com/psyb0t/go-roark.svg)](https://pkg.go.dev/github.com/psyb0t/go-roark)
[![CI](https://github.com/psyb0t/go-roark/actions/workflows/pipeline.yml/badge.svg?branch=main)](https://github.com/psyb0t/go-roark/actions/workflows/pipeline.yml)
[![version](https://raw.githubusercontent.com/psyb0t/go-roark/badges/version.svg)](https://github.com/psyb0t/go-roark/tags)
[![license](https://raw.githubusercontent.com/psyb0t/go-roark/badges/license.svg)](LICENSE)
[![imported by](https://raw.githubusercontent.com/psyb0t/go-roark/badges/importers.svg)](https://github.com/psyb0t/go-roark/blob/badges/importers.md)

Unofficial Go client for the [Roark API](https://api.roark.ai/doc), generated with [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) from the OpenAPI document Roark publishes. I am not affiliated with Roark.

The package is `roark`, in the repository root. Everything in `client.gen.go` is generated. Don't edit it by hand, regenerate it.

## Install

```sh
go get github.com/psyb0t/go-roark
```

## Usage

Every endpoint in the spec has a typed method on `ClientWithResponses`. Responses decode into `JSON200` and the other status fields.

```go
package main

import (
	"context"
	"net/http"
	"os"

	"github.com/psyb0t/go-roark"
)

func main() {
	auth := roark.WithRequestEditorFn(func(_ context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+os.Getenv("ROARK_API_KEY"))

		return nil
	})

	client, err := roark.NewClientWithResponses("https://api.roark.ai", auth)
	if err != nil {
		panic(err)
	}

	res, err := client.GetV1ProjectsWithResponse(context.Background())
	if err != nil {
		panic(err)
	}

	_ = res.JSON200
}
```

The same code is in `example_test.go`, so `make test` checks that it compiles.

The full API is in the [Go reference](https://pkg.go.dev/github.com/psyb0t/go-roark); from inside your module, `go doc github.com/psyb0t/go-roark` prints the same.

The client covers all 160 operations across 107 paths in the source spec: agents and prompts, calls and transcripts, chats, personas, simulations and run plans, metrics and collection jobs, customer flows, webhooks, issues, managed configuration, knowledge bases, organization projects and benchmarks. Method names follow the spec's operation IDs, normalized to Go identifiers.

## Auth

Roark uses bearer authentication. Read the credential from `ROARK_API_KEY` and set `Authorization: Bearer <credential>` through `WithRequestEditorFn`, as in the example above. The production base URL is `https://api.roark.ai`.

A project-scoped API key is bound to one project. A user-scoped credential can select a project with `req.Header.Set("X-Roark-Project-Id", projectID)` in the same request editor. `GetV1ProjectsWithResponse` lists the projects available to the current credential; `GetV1MeWithResponse` describes its scope and permissions.

Transport and decoding errors use the returned `error`. HTTP error responses still return a response object, so check `StatusCode()` and the expected `JSON200` or `JSON201` field before reading its data.

## Regenerating

The spec is not committed. `make generate` downloads the current OpenAPI document into the gitignored `api/` directory and regenerates the client:

```sh
make generate
```

`gen.go` declares the steps, so `make generate` is just `go generate ./...`:

- `scripts/fetch-spec.sh` downloads `https://api.roark.ai/doc` to `api/openapi.json`.
- `scripts/patch-spec.sh` writes `api/openapi.fixed.json`. It downconverts OpenAPI 3.1 nullable type arrays and null union branches, collapses single-branch unions with nullable references wrapped in `allOf`, converts `const` to a one-value enum, uses the first schema example, and translates numeric exclusive bounds to OpenAPI 3.0 bounds. The generated input declares OpenAPI 3.0.3. Every path and operation remains included.
- oapi-codegen runs with `config.yaml`, mapping `number` to `float64` and prefixing enum values to avoid collisions. `skip-prune: true` retains all component models, including types referenced only by top-level webhook payload unions.

`.github/workflows/regen.yml` does the same every day in CI. When the new spec changes the client, it runs `make lint`, `make test` and `make sec`. If a check fails, Claude edits the code and the checks run again, up to three times, and a review step reads the result. A passing change is committed with a CHANGELOG entry and tagged as the next release: a patch version for fixes and additions, a minor version for breaking changes.

## Agent integrations

The [skill](.agents/skills/go-roark) teaches a coding agent how to use go-roark when it writes Go. It works in any agent that reads `.agents/skills/`, and installs natively in the clients below.

### Claude Code

```bash
claude plugin marketplace add psyb0t/agents
claude plugin install go-roark@psyb0t
```

### Codex

```bash
codex plugin marketplace add psyb0t/agents
codex plugin add go-roark@psyb0t
```

Codex also picks the skill up with no install in any checkout of this repo.

### OpenClaw

The skill is published to ClawHub on every release:

```bash
openclaw skills install @psyb0t/go-roark
```

## Development

Every command runs in the `Dockerfile.dev` image. The host needs only Docker and Make.

```sh
make help      # list targets
make dep       # go mod tidy + vendor
make lint      # shellcheck, shfmt, go fix, golangci-lint
make test      # go test -race
make sec       # gitleaks, govulncheck, semgrep
make generate  # download the spec and regenerate
```

## License

[MIT](LICENSE). The generated code comes from Roark's published API specification. The specification itself is not redistributed here. Its metadata does not state a license for the specification.

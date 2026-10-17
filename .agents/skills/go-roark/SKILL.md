---
name: "go-roark"
description: "Write Go integrations with the unofficial Roark API client for voice AI agents, calls, chats, simulations, run plans, personas, metrics and monitoring. Use when Go code needs Roark API access or already imports github.com/psyb0t/go-roark."
homepage: https://github.com/psyb0t/go-roark
user-invocable: true
metadata:
  openclaw:
    emoji: "☎️"
    requires:
      bins:
        - go
      env:
        - ROARK_API_KEY
permissions:
  filesystem: "Guides edits to the Go source of the current project. The library reads and writes no files."
  shell: "The Go toolchain (`go get`, `go build`, `go test`) to add and compile the dependency."
  network: "`go get` fetches the module from your Go module proxy. At runtime the code you write talks HTTPS to api.roark.ai."
---

# go-roark

Unofficial, typed Go client for the Roark API, in package `roark` at the module root.

Read [references/setup.md](references/setup.md) before choosing an operation. It covers installation, project selection, operation names, pagination and response handling.

## Security and safety

- Read the bearer credential from `ROARK_API_KEY`. Never hardcode, commit or log it, and never print a request with its `Authorization` header.
- Use `https://api.roark.ai`. Project-scoped API keys are bound to one project. User-scoped credentials can select an accessible project through `X-Roark-Project-Id` in the request editor.
- Creating or deleting resources, starting simulations, collecting metrics, changing configuration and promoting a revision affect the project. Call those operations only when the user requested the action.
- Calls, transcripts, recordings and credentials may contain private data. Return only the data needed for the task.

## When to use

- Go code needs Roark agents, call or chat monitoring, transcripts, metrics or usage data.
- Go code needs to run voice AI simulations or manage Roark test resources, at the user's request.
- The project already imports `github.com/psyb0t/go-roark`.

## When not to use

- The code is not Go. This module ships only a Go package.
- You need a client maintained by Roark. This one is unofficial and not affiliated with Roark.
- You want to change the generated client itself. See `AGENTS.md` in the repository.

## Usage

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

The example is also compiled in `example_test.go`. Inspect the current generated signatures before adding another operation. Supply a context and check the transport error, HTTP status and expected decoded response field. Finish by compiling and testing the calling project without making live API requests unless the user requested them.

The error return covers transport and decoding failures. An API error may return a nil `error` with `JSON200` nil. Check `StatusCode()` and the relevant error field or `Body`. Optional and nullable properties may be pointers; check them before dereferencing.

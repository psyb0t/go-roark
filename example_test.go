package roark_test

import (
	"context"
	"net/http"
	"os"

	"github.com/psyb0t/go-roark"
)

// This README example needs real credentials, so it only compiles.
//
//nolint:testableexamples // needs real credentials, compile check only
func Example() {
	auth := roark.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			req.Header.Set(
				"Authorization",
				"Bearer "+os.Getenv("ROARK_API_KEY"),
			)

			return nil
		},
	)

	client, err := roark.NewClientWithResponses(
		"https://api.roark.ai",
		auth,
	)
	if err != nil {
		panic(err)
	}

	res, err := client.GetV1ProjectsWithResponse(context.Background())
	if err != nil {
		panic(err)
	}

	_ = res.JSON200
}

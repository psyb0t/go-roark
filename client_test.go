package roark_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/psyb0t/go-roark"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testAPIKey    = "EXAMPLE-DO-NOT-USE"
	testProjectID = "project-example"
	testAgentID   = "00000000-0000-4000-8000-000000000001"
)

func newTestClient(
	t *testing.T,
	handler http.HandlerFunc,
) *roark.ClientWithResponses {
	t.Helper()

	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	auth := roark.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			req.Header.Set("Authorization", "Bearer "+testAPIKey)
			req.Header.Set("X-Roark-Project-Id", testProjectID)

			return nil
		},
	)

	client, err := roark.NewClientWithResponses(srv.URL, auth)
	require.NoError(t, err)

	return client
}

func assertAuth(t *testing.T, r *http.Request) {
	t.Helper()

	assert.Equal(t, "Bearer "+testAPIKey, r.Header.Get("Authorization"))
	assert.Equal(t, testProjectID, r.Header.Get("X-Roark-Project-Id"))
}

func writeJSON(t *testing.T, w http.ResponseWriter, status int, body string) {
	t.Helper()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err := w.Write([]byte(body))
	assert.NoError(t, err)
}

func TestGetProjects(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/projects", r.URL.Path)
		assertAuth(t, r)
		writeJSON(t, w, http.StatusOK, `{"data":{"projects":[{
			"id":"project-example","name":"Support","slug":"support",
			"organization":{"id":"org-example","name":"Example"},
			"grantedPermissions":["call:read"]
		}]}}`)
	})

	res, err := client.GetV1ProjectsWithResponse(context.Background())
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode())
	require.NotNil(t, res.JSON200)
	require.Len(t, res.JSON200.Data.Projects, 1)

	project := res.JSON200.Data.Projects[0]
	assert.Equal(t, testProjectID, project.Id)
	assert.Equal(t, "Support", project.Name)
	assert.Equal(t, []string{"call:read"}, project.GrantedPermissions)
}

func TestCreateAgent(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/v1/agent", r.URL.Path)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assertAuth(t, r)

		var body roark.PostV1AgentJSONRequestBody

		assert.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, "Support", body.Name)
		writeJSON(t, w, http.StatusCreated, `{"data":{
			"id":"00000000-0000-4000-8000-000000000001",
			"name":"Support","description":null,"customId":null,
			"createdAt":"","updatedAt":""
		}}`)
	})

	res, err := client.PostV1AgentWithResponse(
		context.Background(),
		roark.PostV1AgentJSONRequestBody{Name: "Support"},
	)
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, res.StatusCode())
	require.NotNil(t, res.JSON201)

	assert.Equal(t, testAgentID, res.JSON201.Data.Id.String())
	assert.Equal(t, "Support", res.JSON201.Data.Name)
	assert.Nil(t, res.JSON201.Data.Description)
	assert.Nil(t, res.JSON201.Data.CustomId)
}

func TestCanceledContext(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(_ http.ResponseWriter, _ *http.Request) {
		t.Error("canceled request reached the server")
	})
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := client.GetV1ProjectsWithResponse(ctx)
	require.ErrorIs(t, err, context.Canceled)
}

func TestMalformedResponse(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assertAuth(t, r)
		writeJSON(t, w, http.StatusOK, `{"data":`)
	})

	_, err := client.GetV1ProjectsWithResponse(context.Background())
	require.Error(t, err)
}

func TestAgentPagination(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/agent", r.URL.Path)
		assert.Equal(t, "2", r.URL.Query().Get("limit"))
		assertAuth(t, r)

		switch r.URL.Query().Get("after") {
		case "":
			writeJSON(t, w, http.StatusOK, `{"data":[
				{"id":"00000000-0000-4000-8000-000000000001",
				"name":"First","description":null,"customId":null,
				"createdAt":"","updatedAt":""},
				{"id":"00000000-0000-4000-8000-000000000002",
				"name":"Second","description":"Agent","customId":null,
				"createdAt":"","updatedAt":""}
			],"pagination":{"total":3,"hasMore":true,
			"nextCursor":"cursor +/&"}}`)
		case "cursor +/&":
			writeJSON(t, w, http.StatusOK, `{"data":[
				{"id":"00000000-0000-4000-8000-000000000003",
				"name":"Third","description":null,"customId":null,
				"createdAt":"","updatedAt":""}
			],"pagination":{"total":3,"hasMore":false,
			"nextCursor":null}}`)
		default:
			writeJSON(t, w, http.StatusOK, `{"data":[],
				"pagination":{"total":3,"hasMore":false,
				"nextCursor":null}}`)
		}
	})
	limit := 2
	params := &roark.GetV1AgentParams{Limit: &limit}

	var names []string

	for range 2 {
		res, err := client.GetV1AgentWithResponse(context.Background(), params)
		require.NoError(t, err)
		require.NotNil(t, res.JSON200)
		assert.Equal(t, float64(3), res.JSON200.Pagination.Total)

		for _, agent := range res.JSON200.Data {
			names = append(names, agent.Name)
		}

		params.After = res.JSON200.Pagination.NextCursor
		if !res.JSON200.Pagination.HasMore {
			assert.Nil(t, params.After)

			break
		}
	}

	assert.Equal(t, []string{"First", "Second", "Third"}, names)

	pastEnd := "past-end"
	params.After = &pastEnd
	res, err := client.GetV1AgentWithResponse(context.Background(), params)
	require.NoError(t, err)
	require.NotNil(t, res.JSON200)
	assert.Empty(t, res.JSON200.Data)
	assert.False(t, res.JSON200.Pagination.HasMore)
}

func TestAuthenticationErrors(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		auth string
	}{
		{name: "missing"},
		{name: "malformed", auth: "invalid"},
		{name: "rejected", auth: "Bearer EXAMPLE-REJECTED"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			client := newTestClient(t, func(
				w http.ResponseWriter,
				r *http.Request,
			) {
				assert.Equal(t, tc.auth, r.Header.Get("Authorization"))
				writeJSON(t, w, http.StatusUnauthorized, `{
					"type":"authentication","code":"unauthorized",
					"message":"Authentication required"
				}`)
			})
			res, err := client.GetV1ProjectsWithResponse(
				context.Background(),
				func(_ context.Context, req *http.Request) error {
					req.Header.Set("Authorization", tc.auth)

					return nil
				},
			)
			require.NoError(t, err)
			assert.Equal(t, http.StatusUnauthorized, res.StatusCode())
			assert.Nil(t, res.JSON200)
			require.NotNil(t, res.JSON401)
			assert.Equal(t, "unauthorized", res.JSON401.Code)
		})
	}
}

func TestRevokeAPIKey(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/v1/me/api-keys/example-key", r.URL.Path)
		assertAuth(t, r)
		writeJSON(t, w, http.StatusOK, `{
			"data":{"id":"example-key","deleted":true}
		}`)
	})
	res, err := client.DeleteV1MeApiKeysByIdWithResponse(
		context.Background(), "example-key",
	)
	require.NoError(t, err)
	require.NotNil(t, res.JSON200)
	assert.Equal(t, "example-key", res.JSON200.Data.Id)
	assert.Equal(t, roark.MeApiKeyRevokeResponseDeletedTrue,
		res.JSON200.Data.Deleted)
}

func TestNullableUnion(t *testing.T) {
	t.Parallel()

	var group roark.SimulationRunGroup

	err := json.Unmarshal([]byte(`{"latestReport":null}`), &group)
	require.NoError(t, err)
	assert.Nil(t, group.LatestReport)

	err = json.Unmarshal([]byte(`{"latestReport":{
		"id":"00000000-0000-4000-8000-000000000001",
		"status":"PENDING","createdAt":"","completedAt":null,
		"document":null,"errorMessage":null
	}}`), &group)
	require.NoError(t, err)
	require.NotNil(t, group.LatestReport)
	assert.Equal(t, "PENDING", string(group.LatestReport.Status))
	assert.Nil(t, group.LatestReport.Document)
}

func TestTranscriptRole(t *testing.T) {
	t.Parallel()

	var entry roark.TranscriptEntryAgent

	err := json.Unmarshal([]byte(`{
		"startOffsetMs":0,"endOffsetMs":100,"text":"Hello",
		"role":"AGENT","payload":null
	}`), &entry)
	require.NoError(t, err)
	assert.Equal(t, roark.TranscriptEntryAgentRoleAGENT, entry.Role)
	assert.Equal(t, 100, entry.EndOffsetMs)
	assert.Nil(t, entry.Payload)
}

func TestWebhookPayload(t *testing.T) {
	t.Parallel()

	var event roark.WebhookEventJSONBody

	err := json.Unmarshal([]byte(`{
		"event":"call.analysis.completed","version":"1.0",
		"timestamp":"2026-10-17T00:00:00Z","data":{
			"callId":"00000000-0000-4000-8000-000000000001",
			"callAnalysisJobId":"00000000-0000-4000-8000-000000000002",
			"projectId":"00000000-0000-4000-8000-000000000003"
		}
	}`), &event)
	require.NoError(t, err)
	payload, err := event.AsCallAnalysisCompletedWebhook()
	require.NoError(t, err)
	assert.Equal(t, testAgentID, payload.Data.CallId.String())
	assert.Equal(t, "call.analysis.completed", string(payload.Event))
}

func TestProjectDenied(t *testing.T) {
	t.Parallel()

	const errorBody = `{
		"type":"forbidden","code":"forbidden",
		"message":"Project access denied"
	}`

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "other-project", r.Header.Get("X-Roark-Project-Id"))
		writeJSON(t, w, http.StatusForbidden, errorBody)
	})
	res, err := client.GetV1ProjectsWithResponse(
		context.Background(),
		func(_ context.Context, req *http.Request) error {
			req.Header.Set("X-Roark-Project-Id", "other-project")

			return nil
		},
	)
	require.NoError(t, err)
	assert.Equal(t, http.StatusForbidden, res.StatusCode())
	assert.Nil(t, res.JSON200)
	assert.JSONEq(t, errorBody, string(res.Body))
}

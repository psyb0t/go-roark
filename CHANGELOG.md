# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.2.1, 2026-10-23

- Documentation only. The doc comments of `DeleteV1OrgProjectByProjectId` and `DeleteV1OrgProjectByProjectIdWithResponse`, on `Client`, `ClientWithResponses` and both interfaces, now say that deleting a project also pauses its active schedules, and that the API returns 409 when the project is the last one in the organization, has an active agent connection, or has a simulation running.

## v0.2.0, 2026-10-22

- Breaking: the filter `Value` field is now optional and a pointer. It changed to `*AnalyticsQuery_Filters_Value` and similar `*..._Value` types in the `Filters` and `ComparisonGroups` fields of `AnalyticsQuery`, `Dashboard`, `DashboardInput`, `DashboardWidget`, `DashboardWidgetInput` and `UpdateDashboardInput`, with the JSON tag `value,omitempty`. `DashboardFilterConfig.Value` is now `*DashboardFilterConfig_Value`. Code that sets or reads these values must take addresses and check for nil.
- Breaking: the anonymous struct types of `Measures` in `AnalyticsQuery`, `DashboardWidget` and `DashboardWidgetInput`, and of `Widgets` in `Dashboard`, `DashboardInput` and `UpdateDashboardInput`, changed because they embed the filter and measure types above. Composite literals of these types must be updated.
- Added call deletion: `DeleteV1CallByCallId`, `PostV1CallBatchDelete` and `PostV1CallBatchDeleteWithBody`, with `WithResponse` methods, request builders, `ParseDeleteV1CallByCallIdResponse`, `ParsePostV1CallBatchDeleteResponse`, `PostV1CallBatchDeleteJSONBody` and `PostV1CallBatchDeleteJSONRequestBody`.
- Added `PostV1MetricCollectionJobsByJobIdCancel` with its `WithResponse` method, request builder and `ParsePostV1MetricCollectionJobsByJobIdCancelResponse`.
- `GetV1CallParams` has a new `Properties` query field, and `GetV1CallResponse` has `JSON400`.
- Added the filter operators `ISSET` and `ISNOTSET` to every filter, measure and pass operator enum, for example `AnalyticsQueryFiltersOperatorISSET`, `DashboardFilterConfigOperatorISNOTSET` and `AnalyticsSchemaOperatorsISSET`.
- `ClientInterface` and `ClientWithResponsesInterface` gained the matching methods. Code that implements these interfaces, such as test mocks, must add them.

## v0.1.2, 2026-10-21

- Added dashboard endpoints: `GetV1Dashboard`, `PostV1Dashboard`, `GetV1DashboardByDashboardId`, `PutV1DashboardByDashboardId`, `DeleteV1DashboardByDashboardId` and `GetV1DashboardByDashboardIdData`.
- Added widget endpoints: `PostV1DashboardByDashboardIdWidget`, `PutV1DashboardByDashboardIdWidgetByWidgetId`, `DeleteV1DashboardByDashboardIdWidgetByWidgetId` and `GetV1DashboardByDashboardIdWidgetByWidgetIdData`.
- Added analytics endpoints: `PostV1AnalyticsQuery` and `GetV1AnalyticsSchema`.
- Added cost endpoints: `GetV1CallByCallIdCost` and `GetV1SimulationPlanJobByJobIdCost`.
- Each new operation has `Client` methods, `WithBody` variants where the request has a body, and `WithResponse` methods on `ClientWithResponses`, plus new request and response types and enums such as `AnalyticsQueryComparisonGroupsColor`.
- `ClientInterface` and `ClientWithResponsesInterface` gained the matching methods. Code that implements these interfaces, such as test mocks, must add them.
- The Go toolchain and the `go` line in `go.mod` moved from 1.26.9 to 1.26.10, which fixes GO-2026-6809 in `crypto/x509`. The development image and the setup guide use the same version.

## v0.1.1, 2026-10-17

- Documentation only. The agent skill explains how to look up every endpoint, parameter and model with `go doc` or on pkg.go.dev, and declares its license.
- The skill no longer lists the API credentials under OpenClaw's `requires.env`, which kept OpenClaw from loading it until they were set. It requires only `go`; the credentials are still documented as what your program reads.
- The README points to the full API reference, and the examples in the README and the skill were checked against the current code.

## v0.1.0, 2026-10-17

- Added `roark`, a REST client generated with oapi-codegen from the Roark Analytics API specification, version 1.0.0, OpenAPI 3.1.0, published at https://api.roark.ai/doc. It covers all 107 paths and 160 operations for voice AI testing and monitoring.
- Number fields are `float64` instead of oapi-codegen's default `float32`.
- `make generate` downloads the current spec, downconverts its OpenAPI 3.1 schema features and regenerates the client. The spec is not committed.
- Added local HTTP tests, development and security tooling, CI workflows and the `go-roark` agent skill.

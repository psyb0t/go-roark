# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.4.0, 2026-10-28

- Added agent draft and version endpoints. New `Client` and `ClientWithResponses` methods: `GetV1AgentByAgentIdDraft`, `PostV1AgentByAgentIdDraft`, `PatchV1AgentByAgentIdDraft`, `DeleteV1AgentByAgentIdDraft`, `PostV1AgentByAgentIdDraftPublish`, `GetV1AgentByAgentIdVersions`, `GetV1AgentByAgentIdVersionsByVersionId` and `PostV1AgentByAgentIdVersionsByVersionIdRestore`, each with `WithBody` and `WithResponse` variants where the client has them.
- Added the request types `CreateAgentDraftInput`, `UpdateAgentDraftInput`, `PublishAgentDraftInput` and `RestoreAgentVersionInput`, and the response types `AgentVersionResponse`, `ListAgentVersionsResponse`, `AgentVersionHandoff`, `AgentVersionIssue`, `AgentVersionSubAgent` and `AgentVersionTool`, with their enum types and constants. `GetV1AgentByAgentIdVersionsParams` holds the list query parameters.
- Added the same eight methods to `ClientInterface` and `ClientWithResponsesInterface`. Custom implementations or mocks of these interfaces must add them.
- Added `PhoneNumberAutoPickInput` with `PhoneNumberAutoPickInputCountry` and `PhoneNumberAutoPickInputNumberType`, `CreatePhoneNumberInputCountry`, and the `Country` and `NumberType` query parameters on `GetV1PhoneNumberAvailableParams`.
- Changed `CreatePhoneNumberInput.E164` from `string` to `*string`. Callers that set it must pass a pointer.
- Changed `PostV1AgentBuildJSONBody.Tts` and `PostV1AgentBuildJSONRequestBody.Tts` so `Provider` is the new `PostV1AgentBuildJSONBodyTtsProvider` enum instead of `string`. Callers that assign a plain string variable must convert it.
- Added an optional `AgentVersionId` field to each agent in the call responses `GetV1CallResponse.JSON200`, `GetV1CallByCallIdResponse.JSON200` and `PostV1CallResponse.JSON200`.

## v0.3.3, 2026-10-28

- Updated Semgrep in the development image to 1.182.1 and refreshed its locked dependencies. certifi 2026.10.23 drops the emSign root CAs whose intermediate keys were compromised (GHSA-6q48-6c7w-9gw2).
- Dependabot no longer opens pull requests for `mcp` or the OpenTelemetry packages in `requirements-semgrep.txt`. Semgrep pins their versions itself, so bumping one of them on its own cannot install.
- Updated kin-openapi, which the code generator uses, from 0.142.0 to 0.149.0. The newer version fixes several panics on malformed specs reported in GitHub security advisories.
- The security scan merges its two gitleaks runs into one SARIF run. GitHub code scanning rejects a report with two runs from the same tool, so no scan results had been uploaded.

## v0.3.2, 2026-10-27

- Added the `ApiCheckMetricConfig` type for HTTP request metrics in config bundles and projects, with `ApiCheckMetricConfigCalculationType` (`ApiCheckMetricConfigCalculationTypeHTTPREQUEST`) and `ApiCheckMetricConfigKind` (`ApiCheckMetricConfigKindMetric`). Both enums have `Valid` methods.
- Added the `ApiCheckRequestConfig` type with `Url`, `Method`, `Headers` and `Body`, and `ApiCheckRequestConfigMethod` with `ApiCheckRequestConfigMethodGET` and `ApiCheckRequestConfigMethodPOST`.
- `ConfigBundleResources0` and `ProjectConfig_Resources_Item` gained `AsApiCheckMetricConfig`, `FromApiCheckMetricConfig` and `MergeApiCheckMetricConfig`.
- Added an optional `CalculationType` field (`*MetricConfigCalculationType`) to `MetricConfig`, with the value `MetricConfigCalculationTypeLLMJUDGE`.
- Added an optional `Tts` field to `PostV1AgentBuildJSONBody` and `PostV1AgentBuildJSONRequestBody`, holding the text-to-speech `Provider` and `VoiceId`. The `Voice` field comment now describes a display name.

## v0.3.1, 2026-10-27

- The development image is based on Debian (`golang:1.26.10-trixie`) instead of Alpine. Alpine keeps only the latest build of each package, so the exact pins broke the image build whenever a package was updated. System packages now come from the snapshot.debian.org date the base image was built from, which keeps them fixed for a given base image digest. jq 1.8.2, shellcheck 0.11.0 and shfmt 3.13.1 keep their versions, and Semgrep runs in its own virtual environment.

## v0.3.0, 2026-10-26

- Breaking: `SimulationGroupNotifyConfig` is replaced by `SimulationNotifyConfig`, and the `Notify` field changed to `*SimulationNotifyConfig` on `SimulationGroupConfig`. The email types are renamed the same way: `SimulationNotifyConfig_Email`, `SimulationNotifyConfigEmail0` and `SimulationNotifyConfigEmail1`, with the matching `As`, `From` and `Merge` methods. Code that names the old types or methods must change to the new names.
- Added `Notify` (`*SimulationNotifyConfig`) to `SimulationPlanConfig` and `QaSimulationPlanConfig`.
- Added `Notifications` to `CreateRunPlanFromConfig`, `CreateRunPlanFromTemplate`, `UpdateRunPlanInput` and `RunPlanResponse`, using the new `SimulationPlanNotificationsInput` type.
- Added phone number endpoints: `GetV1PhoneNumber`, `PostV1PhoneNumber`, `PutV1PhoneNumberByPhoneNumberId`, `DeleteV1PhoneNumberByPhoneNumberId`, `GetV1PhoneNumberAllowance` and `GetV1PhoneNumberAvailable`. Each has `Client` methods, `WithBody` variants where the request has a body, `WithResponse` methods, request builders and `Parse...Response` functions. New types include `CreatePhoneNumberInput`, `UpdatePhoneNumberInput`, `PhoneNumberResponse`, `PhoneNumberAllowanceResponse`, `AvailablePhoneNumberResponse` and `ReleasePhoneNumberResponse`.
- Added the `PHONE_NUMBER` source value: `CallCostLineSourcePHONENUMBER`, `CostBreakdownItemSourcePHONENUMBER` and `UsageCostTypeSourcePHONENUMBER`.
- Added HTTP request metrics: `HttpRequestMetricInput`, `HttpRequestMetricResponse`, `HttpRequestMetricConfig`, `HttpRequestMetricDetails`, `HttpRequestMetricPayload` and `HttpRequestMetricVerdict`, with their enums. `CreateMetricDefinitionInput`, `CreateMetricDefinitionResponse`, `MetricDefinitionResponse` and `UpdateMetricDefinitionResponse` have `As`, `From` and `Merge` methods for them, and `UpdateMetricDefinitionInput` has an `HttpRequest` field.
- Added the `HttpRequestMetricCheck` and `HttpRequestMetricCheckWithBody` webhook methods on `WebhookInitiator`, with `NewHttpRequestMetricCheckWebhookRequest` and `HttpRequestMetricCheckJSONRequestBody`.
- Added `RunSimulationFromConfigKey` as an option of `RunSimulationInput`, with `AsRunSimulationFromConfigKey`, `FromRunSimulationFromConfigKey` and `MergeRunSimulationFromConfigKey`.
- `ClientInterface`, `ClientWithResponsesInterface` and `WebhookInitiatorInterface` gained the matching methods. Code that implements these interfaces, such as test mocks, must add them.

## v0.2.2, 2026-10-24

- Added the `ToolHandler` type, which describes an HTTP handler for a tool, with `Type` (`ToolHandlerType`, only `ToolHandlerTypeHttp`), `Url` and an optional `Method` (`ToolHandlerMethod`).
- Added `ToolHandlerMethod` with the values `ToolHandlerMethodGET`, `ToolHandlerMethodPOST`, `ToolHandlerMethodPUT`, `ToolHandlerMethodPATCH` and `ToolHandlerMethodDELETE`. `ToolHandlerMethod` and `ToolHandlerType` have `Valid` methods.
- `ToolConfig` has a new optional `Handler` field of type `*ToolHandler`.

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

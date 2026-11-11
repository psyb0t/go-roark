# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.15.0, 2026-11-11

- Breaking: the anonymous `Settings` struct of `UpdateAgentDraftInput` gained an optional `VadMaxSilenceSecs` field (`*float64`). Composite literals of this struct type must be updated. Code that sets fields by name still compiles.
- Added `AgentVersionResponse.VadMaxSilenceSecs` (`*float64`), the longest silence the agent waits through when the caller sounds unfinished. Null uses `VadStopSecs` when that is set, else the default.
- Added doc comments to `AgentVersionResponse.VadStartSecs` and `AgentVersionResponse.VadStopSecs`.
- Raised the minimum Go version in `go.mod` from 1.26.10 to 1.26.11, which fixes the `net/http` vulnerability GO-2026-6859. The development image uses the matching Go 1.26.11 base image.

## v0.14.0, 2026-11-10

- Breaking: the calendar connection endpoints moved from `/v1/calendar/connection` to `/v1/integrations/calendar`. The old methods are removed: `GetV1CalendarConnection`, `PostV1CalendarConnection`, `GetV1CalendarConnectionByConnectionId`, `PatchV1CalendarConnectionByConnectionId`, `DeleteV1CalendarConnectionByConnectionId`, `PostV1CalendarConnectionByConnectionIdCheck`, `GetV1CalendarConnectionByConnectionIdTargets` and `GetV1CalendarConnectionByConnectionIdTargetsByTargetId`, with their `WithBody` and `WithResponse` variants and response types. The replacements are `GetV1IntegrationsCalendar`, `PostV1IntegrationsCalendar`, `GetV1IntegrationsCalendarByConnectionId`, `PatchV1IntegrationsCalendarByConnectionId`, `DeleteV1IntegrationsCalendarByConnectionId`, `PostV1IntegrationsCalendarByConnectionIdCheck` and the two `Targets` methods.
- Breaking: removed `PostV1AgentByAgentIdWebSessionTokens`, `CreateWebSessionTokenInput` and its `Modality` enum. Use `PostV1AgentByAgentIdConversations` with `CreateAgentConversationInput` and `CreateAgentWebConversationInput` instead. The result types are `AgentConversationCreated` and `AgentWebConversationCreated`.
- Breaking: removed the public prospect simulation endpoints `PostV1PublicProspectSimRun`, `GetV1PublicProspectSimReportByToken` and `GetV1PublicProspectSimReportByTokenObserverUrl`, with their variants and response types.
- Breaking: chat responses changed. In the `JSON200` and `GetJSON200` types of `GetV1ChatResponse`, `GetV1ChatByIdResponse` and `PostV1ChatResponse`, `DurationMs` is now `*float64` and `EndTimestamp` is now `*string`, and a required `Status` field was added. Code that reads or builds these values must be updated.
- Breaking: removed `AgentVersionResponseModalitiesCall`. Added `AgentConfigModalitiesVoice`, `AgentVersionResponseModalitiesVoice` and `CollectorConfigModalityVoice`, plus similar `Voice` modality values.
- Added conversation endpoints `GetV1Conversation` and `GetV1ConversationByConversationId`, with the types `Conversation`, `ConversationList`, `ConversationModality` and `ConversationStatus` (`PENDING`, `INPROGRESS`, `ENDED`).
- Added MCP server endpoints: `GetV1IntegrationsMcpServer`, `PostV1IntegrationsMcpServer`, `GetV1IntegrationsMcpServerByServerId`, `PatchV1IntegrationsMcpServerByServerId`, `DeleteV1IntegrationsMcpServerByServerId`, `PostV1IntegrationsMcpServerByServerIdRefresh` and `PostV1IntegrationsMcpServerDiscover`. Types include `McpServerCreateBody`, `McpServerUpdateBody`, `McpServerDiscoverBody`, `McpServerResponse`, `McpServerListResponse`, `McpServerTool` and `McpServerAuth`.
- `ClientInterface` and `ClientWithResponsesInterface` gained the new methods and lost the removed ones. Code that implements these interfaces, such as test mocks, must be updated.

## v0.13.0, 2026-11-10

- Breaking: the anonymous attribute struct gained a required `IsInferred` field (`bool`), true when the value was inferred from the conversation and not stated. This changes the `JSON200` and `GetJSON200` types of `GetV1CustomerResponse`, `GetV1CustomerByCustomerProfileIdResponse`, `PutV1CustomerByCustomerProfileIdNotesResponse`, `PutV1CustomerByCustomerProfileIdAttributesByKeyResponse`, `DeleteV1CustomerByCustomerProfileIdAttributesByKeyResponse` and `PostV1CustomerByCustomerProfileIdAttributesByKeyUnlockResponse`. Code that builds or assigns these struct types must be updated. Code that reads fields by name still compiles.
- Added calendar tools. `ToolBody` has an optional `Calendar` field, and `ToolResponse` has a `Calendar` field (`*ToolCalendar`), null for other tools.
- Added the types `ToolCalendar` and `ToolCalendarAppointmentType`, and the `ToolCalendarOperation` and `ToolBodyCalendarOperation` enums with the values `BOOK`, `CANCEL`, `CHECK_AVAILABILITY`, `LIST_BOOKINGS`, `LOOKUP_BOOKING` and `RESCHEDULE`. Each enum has a `Valid` method.
- Updated the doc comments of the tool create and replace methods to describe the `calendar` field.

## v0.12.0, 2026-11-08

- Breaking: the anonymous customer struct gained an optional `ExternalId` field (`*string`) in `CallCreate.Customer`, `CallCreate.Customers`, `ChatCreate.Customer` and `ChatCreate.Customers`. The same change changed the `JSON200` and `GetJSON200` types of `GetV1CallResponse`, `GetV1CallByCallIdResponse`, `PostV1CallResponse`, `GetV1ChatResponse`, `GetV1ChatByIdResponse` and `PostV1ChatResponse`. Composite literals of these struct types must be updated. Code that sets fields by name still compiles.
- Added customer profile endpoints: `GetV1Customer`, `GetV1CustomerByCustomerProfileId`, `DeleteV1CustomerByCustomerProfileId`, `GetV1CustomerByCustomerProfileIdConversations`, `PutV1CustomerByCustomerProfileIdNotes`, `PutV1CustomerByCustomerProfileIdAttributesByKey`, `DeleteV1CustomerByCustomerProfileIdAttributesByKey` and `PostV1CustomerByCustomerProfileIdAttributesByKeyUnlock`. Each has a `WithResponse` variant, a request builder, a response parser and a response type. The ones that send a body also have `WithBody` variants. `GetV1CustomerParams` and `GetV1CustomerByCustomerProfileIdConversationsParams` hold the query filters, including `GetV1CustomerParamsFieldOperator`.
- Added customer field endpoints: `GetV1CustomerFields`, `PostV1CustomerFields`, `PatchV1CustomerFieldsByKey` and `DeleteV1CustomerFieldsByKey`, with `PostV1CustomerFieldsJSONBody` and `PatchV1CustomerFieldsByKeyJSONBody`. Value type enums cover `TEXT`, `NUMBER`, `BOOLEAN`, `DATE`, `EMAIL`, `CURRENCY` and `PHONE_NUMBER`.
- Added `PostV1AgentByAgentIdWebSessionTokens`, with `WithBody` and `WithResponse` variants, and the types `CreateWebSessionTokenInput` and `WebSessionToken`. Both have a `Modality` enum with `Call` and `Chat` values.
- Added the cost source constants `HOSTED_AGENT_LLM`, `HOSTED_AGENT_STT`, `HOSTED_AGENT_TELEPHONY` and `HOSTED_AGENT_TTS` for `CallCostLineSource`, `CostBreakdownItemSource` and `UsageCostTypeSource`.
- `ClientInterface` and `ClientWithResponsesInterface` gained the matching methods. Code that implements these interfaces, such as test mocks, must add them.

## v0.11.1, 2026-11-07

- Added an optional `EndsCall` field (`*bool`) to `ConfigFlowStep` and `FlowStep0` through `FlowStep10`. Set it on the last step of a scripted path to end the call there instead of handing over to the improv flow.
- Added the `AgentConfigProviderCustom` constant (`"custom"`) to the `AgentConfigProvider` enum.
- Added `JSON409` and `GetJSON409` to `DeleteV1PersonaByPersonaIdResponse`, holding an `ErrorResponse`. `ParseDeleteV1PersonaByPersonaIdResponse` now fills it on an HTTP 409 response.

## v0.11.0, 2026-11-06

- Breaking: `CreateAgentEndpointInput.Value` and `AgentEndpointDetailResponse.Value` changed from `string` to `*string`. The `Value` field of the `AgentConfig.Endpoints` element changed from `string` to `*string` too. Callers must pass and read a pointer.
- Breaking: the `ImprovFlowConfig.HappyPath` struct and the `ImprovFlowConfig.EdgeCases` element type gained an `OpeningLine` field, so code that builds these anonymous structs by type must be updated.
- Added `OpeningLine` to `FlowHappyPathInput`, `FlowEdgeCaseInput`, `ImprovFlowVariant`, `ScriptedFlowVariant`, `UpdateFlowVariantInput` and `VoicemailFlowVariant`.
- Added `Client.DeleteV1SimulationPlanJobByJobId` and `ClientWithResponses.DeleteV1SimulationPlanJobByJobIdWithResponse`, with `NewDeleteV1SimulationPlanJobByJobIdRequest`, `ParseDeleteV1SimulationPlanJobByJobIdResponse` and the types `DeleteV1SimulationPlanJobByJobIdResponse` and `DeleteSimulationPlanJobResponse`.
- Added the public campaign credit endpoints `PostV1PublicCampaignCreditByCampaignKeyClaim` and `PostV1PublicCampaignCreditVerify`, each with a `WithBody` variant, a `WithResponse` variant, request builders, response parsers, JSON body types and an `Outcome` enum (`GRANTED`, `ALREADYCLAIMED`, `AWAITINGSIGNUP`, `CAMPAIGNEXHAUSTED`, `ORGANIZATIONALREADYCLAIMED`, `VERIFICATIONSENT`, `VERIFICATIONALREADYSENT`).
- Added the fields `Greeting`, `JobDescription`, `Llm`, `Provider`, `Stt`, `Tts` and `Voice` to `AgentConfig`, with the enums `AgentConfigProvider`, `AgentConfigLlmProvider`, `AgentConfigSttProvider` and `AgentConfigTtsProvider`.
- Added `SimulationPlanAgentEndpointConfig.Environment`.
- Added the methods above to `ClientInterface` and `ClientWithResponsesInterface`, so custom implementations of these interfaces must add them.

## v0.10.0, 2026-11-05

- Breaking: a new flow step variant was inserted, so the `FlowStep` union variants after it shifted. `FlowStep5.DtmfDigits` and `FlowStep5TypeCUSTOMERDTMF` are removed, as are `FlowStep6TypeAGENTDTMF`, `FlowStep7TypeVOICEMAIL`, `FlowStep8.HandoffPersonaId`, `FlowStep8TypeCUSTOMERHANDOFF`, `FlowStep9.LinkedCustomerFlowId`, `FlowStep9.LinkedCustomerFlowVariantId` and `FlowStep9TypeSCENARIOLINK`. Use the matching names on the next variant: `FlowStep6TypeCUSTOMERDTMF`, `FlowStep7.DtmfDigits`, `FlowStep7TypeAGENTDTMF`, `FlowStep8TypeVOICEMAIL`, `FlowStep9.HandoffPersonaId`, `FlowStep9TypeCUSTOMERHANDOFF` and `FlowStep10TypeSCENARIOLINK`.
- Added `FlowStep10` and `FlowStep10Type`, with `FlowStep.AsFlowStep10`, `FlowStep.FromFlowStep10` and `FlowStep.MergeFlowStep10`.
- Added the customer audio step: `FlowStep5TypeCUSTOMERAUDIO` and `ConfigFlowStepTypeCUSTOMERAUDIO`, with `AudioClip` fields on `FlowStep5` and `ConfigFlowStep`, and a `SilenceDurationSeconds` field on `FlowStep5`. The enums `FlowStep5AudioClip` and `ConfigFlowStepAudioClip` list the clips: `HOLD_BEEP`, six `HOLD_MUSIC_*` clips, `RINGBACK_UK` and `RINGBACK_US`.
- Added `FlowStep7.OffScriptPolicy`, with the enums `FlowStep7OffScriptPolicyReaction` (`REPEAT`, `RESPOND`, `SAY`, `STAY_SILENT`) and `FlowStep7OffScriptPolicyThen` (`ADAPT`, `HANGUP`, `HANGUP_INVALIDATE`, `MOVE_ON`).
- Added the simulation audio clip endpoint: `GetV1SimulationAudioClip` and `GetV1SimulationAudioClipWithResponse`, `NewGetV1SimulationAudioClipRequest`, `ParseGetV1SimulationAudioClipResponse`, `GetV1SimulationAudioClipResponse`, and the types `SimulationAudioClip`, `SimulationAudioClipCategory`, `SimulationAudioClipClip` and `GetSimulationAudioClipsResponse`.
- `ClientInterface` and `ClientWithResponsesInterface` gained `GetV1SimulationAudioClip` and `GetV1SimulationAudioClipWithResponse`. Code that implements these interfaces, such as test mocks, must add them.

## v0.9.0, 2026-11-05

- Development image package installation uses HTTPS and bounded retries for signed snapshot indexes and package downloads.
- Breaking: regenerated the current Roark specification, including calendar connection target operations. ToolHandler is now a union; use its generated As/From helpers and variant-specific enums instead of the removed ToolHandlerMethod and ToolHandlerType aliases.

## v0.8.1, 2026-11-03

- Added calendar connection endpoints: `GetV1CalendarConnection`, `GetV1CalendarConnectionByConnectionId`, `PostV1CalendarConnection`, `PatchV1CalendarConnectionByConnectionId`, `DeleteV1CalendarConnectionByConnectionId` and `PostV1CalendarConnectionByConnectionIdCheck`. Each has `WithBody` and `WithResponse` variants where the client has them, plus request builders and response types. `GetV1CalendarConnectionParams` holds the list filters, including `GetV1CalendarConnectionParamsProvider` (`CALCOM`, `CALENDLY`, `GOOGLE`).
- Added live report endpoints: `GetV1LiveReport`, `GetV1LiveReportByReportId`, `PostV1LiveReport`, `PatchV1LiveReportByReportId`, `DeleteV1LiveReportByReportId`, `PostV1LiveReportByReportIdGenerate`, `GetV1LiveReportByReportIdDocument` and `GetV1LiveReportByReportIdDocumentByDocumentId`, with their `WithBody` and `WithResponse` variants.
- Added the live report types `LiveReport`, `LiveReportInput`, `LiveReportDocument`, `LiveReportConfig`, `LiveReportFilterConfig` and `GenerateLiveReportInput`, with enums for cadence (`DAILY`, `WEEKLY`, `MONTHLY`), status, filter operators, filter types, target aggregation and chart display.
- `ConfigBundleResources0` and `ProjectConfig_Resources_Item` gained `FromLiveReportConfig` and `MergeLiveReportConfig`, and `ConfigBundleResources0` gained `AsLiveReportConfig`. Added `ConfigApplyChangeKindLiveReport` and `ConfigDiffChangeKindLiveReport`.
- Added tool endpoints: `GetV1Tool`, `GetV1ToolByToolId`, `PostV1Tool`, `PutV1ToolByToolId` and `DeleteV1ToolByToolId`, with their `WithBody` and `WithResponse` variants. New types include `ToolBody`, `ToolInput`, `ToolListResponse`, `ToolDeleteResponse` and `GetV1ToolParams` (with `GetV1ToolParamsIncludeBuiltIn`).
- Added the `DELETE` method value to `HttpRequestDefinitionCreateResponseMethodDELETE` and `HttpRequestDefinitionResponseMethodDELETE`.
- Added a `Required` field to `AppendToolInvocationsRequestToolInvocationsParameters0`, `CallCreateToolInvocationsParameters0` and `ChatCreateToolInvocationsParameters0`.
- `ClientInterface` and `ClientWithResponsesInterface` gained the matching methods. Code that implements these interfaces, such as test mocks, must add them.

## v0.8.0, 2026-11-01

- Breaking: the anonymous struct type of the `GetV1CallParams.Properties` map values gained an optional `Eq` field (`*string`), a case-insensitive exact match. Composite literals of this struct type must be updated. Code that sets `Exists` by field still compiles.
- Added an optional `AgentId` field (`*string`) to `GetV1CallParams`. `NewGetV1CallRequest` sends it as the `agentId` query parameter, so `GetV1Call` and `GetV1CallWithResponse` can filter calls by agent.
- Updated the doc comments of `GetV1Call` and `GetV1CallWithResponse`, on `Client`, `ClientWithResponses` and both interfaces. They describe the `agentId` filter, the `properties[agent_name][eq]` filter, and that `sortBy` and `sortDirection` apply together with them.

## v0.7.0, 2026-10-31

- Breaking: the anonymous struct type of `PostV1AgentBuildJSONBody.Tts` and `PostV1AgentBuildJSONRequestBody.Tts` gained an optional `Model` field (`*string`). Composite literals of this struct type must be updated. Code that sets `Provider` and `VoiceId` by field still compiles.
- Added optional `Llm` and `Stt` fields to `PostV1AgentBuildJSONBody` and `PostV1AgentBuildJSONRequestBody`, with the enum types `PostV1AgentBuildJSONBodyLlmProvider` (for example `PostV1AgentBuildJSONBodyLlmProviderANTHROPIC`) and `PostV1AgentBuildJSONBodySttProvider` (for example `PostV1AgentBuildJSONBodySttProviderDEEPGRAM`).
- Added `JSON403` and `GetJSON403` to `PostV1AgentBuildResponse`, holding an `ErrorResponse`.
- Added the `ThresholdMetricConfig` type for threshold metrics in config bundles and projects, with `ThresholdMetricConfig_Value` and the value types `ThresholdMetricConfigValue0`, `ThresholdMetricConfigValue1` and `ThresholdMetricConfigValue2`.
- Added the enums `ThresholdMetricConfigAggregationMode`, `ThresholdMetricConfigCalculationType`, `ThresholdMetricConfigKind`, `ThresholdMetricConfigOperator` and `ThresholdMetricConfigSourceParticipantRole`, with their constants and `Valid` methods.
- `ConfigBundleResources0` and `ProjectConfig_Resources_Item` gained `AsThresholdMetricConfig`, `FromThresholdMetricConfig` and `MergeThresholdMetricConfig`.

## v0.6.0, 2026-10-30

- Breaking: model settings moved from flat fields into per-sub-agent model configs. `UpdateAgentDraftInput.Settings` no longer has `LlmBackendModel`, `LlmMaxTokens`, `LlmModel`, `LlmProvider`, `LlmTemperature`, `LlmVoice`, `SttEndpointingMs`, `SttKeyterms`, `SttModel`, `SttProvider`, `SttSmartFormat` or `SttTurnDetection`. Each entry in `Settings.SubAgents` lost `TtsModel`, `TtsProvider`, `TtsSimilarityBoost`, `TtsSpeed`, `TtsStability` and `TtsVoiceId`, and gained a required `ModelConfigs` field of `[]UpdateAgentDraftInput_Settings_SubAgents_ModelConfigs_Item`.
- Each model config item is either `UpdateAgentDraftInputSettingsSubAgentsModelConfigs0` (type `CASCADE`, with LLM, STT and TTS settings) or `UpdateAgentDraftInputSettingsSubAgentsModelConfigs1` (type `SPEECH_TO_SPEECH`), with their provider, type and turn detection enums.
- Breaking: `AgentVersionResponse` no longer has the `Llm*` and `Stt*` fields (`LlmModel`, `LlmProvider`, `LlmTemperature`, `SttModel`, `SttProvider`, `SttTurnDetection` and the rest). `AgentVersionSubAgent` no longer has the `Tts*` fields. Read these values from the new `AgentVersionSubAgent.ModelConfigs` field.
- Removed the enum types `AgentVersionResponseLlmProvider`, `AgentVersionResponseSttProvider`, `AgentVersionResponseSttTurnDetection`, `AgentVersionSubAgentTtsProvider`, `UpdateAgentDraftInputSettingsLlmProvider`, `UpdateAgentDraftInputSettingsSttProvider`, `UpdateAgentDraftInputSettingsSttTurnDetection` and `UpdateAgentDraftInputSettingsSubAgentsTtsProvider`, with all their constants.
- Added the shared config types `AgentModelConfig`, `AgentCascadeModelConfig`, `AgentSpeechToSpeechModelConfig`, `AgentLlmConfig`, `AgentSttConfig`, `AgentTtsConfig` and `AgentSpeechToSpeechConfig`, with their provider, type and turn detection enums such as `AgentLlmConfigProvider` and `AgentTtsConfigProvider`.

## v0.5.0, 2026-10-29

- Breaking: the draft edit endpoint changed from PATCH to PUT. `PatchV1AgentByAgentIdDraft`, `PatchV1AgentByAgentIdDraftWithBody`, their `ClientWithResponses` variants `PatchV1AgentByAgentIdDraftWithResponse` and `PatchV1AgentByAgentIdDraftWithBodyWithResponse`, `PatchV1AgentByAgentIdDraftJSONRequestBody`, `PatchV1AgentByAgentIdDraftResponse`, `NewPatchV1AgentByAgentIdDraftRequest`, `NewPatchV1AgentByAgentIdDraftRequestWithBody` and `ParsePatchV1AgentByAgentIdDraftResponse` are removed. Use the new `PutV1AgentByAgentIdDraft` family: `PutV1AgentByAgentIdDraft`, `PutV1AgentByAgentIdDraftWithBody`, `PutV1AgentByAgentIdDraftWithResponse`, `PutV1AgentByAgentIdDraftWithBodyWithResponse`, `PutV1AgentByAgentIdDraftJSONRequestBody`, `PutV1AgentByAgentIdDraftResponse`, `NewPutV1AgentByAgentIdDraftRequest`, `NewPutV1AgentByAgentIdDraftRequestWithBody` and `ParsePutV1AgentByAgentIdDraftResponse`.
- Breaking: restoring an agent version is now a rollback. The `PostV1AgentByAgentIdVersionsByVersionIdRestore` methods, `RestoreAgentVersionInput`, the matching request builders, `ParsePostV1AgentByAgentIdVersionsByVersionIdRestoreResponse`, `PostV1AgentByAgentIdVersionsByVersionIdRestoreJSONRequestBody` and `PostV1AgentByAgentIdVersionsByVersionIdRestoreResponse` are removed. The replacements are `PostV1AgentByAgentIdVersionsByVersionIdRollback` and its `WithBody`, `WithResponse` and `WithBodyWithResponse` variants, `RollbackAgentVersionInput` (with an `expectedLiveVersionId` option that makes the API return 409 if another version went live), `PostV1AgentByAgentIdVersionsByVersionIdRollbackJSONRequestBody`, `PostV1AgentByAgentIdVersionsByVersionIdRollbackResponse`, the request builders and `ParsePostV1AgentByAgentIdVersionsByVersionIdRollbackResponse`.
- Breaking: `AgentVersionResponse.RestoredFromVersionId` is renamed to `RolledBackFromVersionId`, with the JSON name `rolledBackFromVersionId`.
- `ClientInterface` and `ClientWithResponsesInterface` dropped the Patch draft and Restore methods and gained the Put draft and Rollback methods. Custom implementations or mocks of these interfaces must be updated.
- Updated the doc comments of the draft methods: the draft is shared per agent, `PostV1AgentByAgentIdDraft` returns 409 when a draft is already open, and the draft edit accepts `expectedRevision` to fail with 409 on a concurrent edit.

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

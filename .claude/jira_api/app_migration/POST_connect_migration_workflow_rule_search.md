# POST /rest/atlassian-connect/1/migration/workflow/rule/search
**operationId:** `MigrationResource.workflowRuleSearch_post`
**Summary:** Get workflow transition rule configurations

Returns configurations for workflow transition rules migrated from server to cloud and owned by the calling Connect app.

## Parameters
- `Atlassian-Transfer-Id` [header] (required) string(uuid) — The app migration transfer ID.

## Request Body
Content-Type: `application/json`
object:
  - `expand`: string
  - `ruleIds` (required): []string(uuid)
  - `workflowEntityId` (required): string(uuid)

## Responses
- 200: object:
  - `invalidRules`: []string(uuid)
  - `validRules`: []WorkflowTransitionRules
  - `workflowEntityId`: string(uuid)
- 400: Returned if the request is not valid.
- 403: Returned if the authorisation credentials are incorrect or missing.

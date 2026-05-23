# GET /rest/atlassian-connect/1/migration/{connectKey}/{jiraIssueFieldsKey}/task
**operationId:** `ConnectToForgeMigrationFetchTaskResource.fetchMigrationTask_get`
**Summary:** Get Connect issue field migration task

Returns the details of a Connect issue field's migration to Forge.

When migrating a Connect app to Forge, [Issue Field](https://developer.atlassian.com/cloud/jira/software/modules/issue-field/) modules
must be converted to [Custom field](https://developer.atlassian.com/platform/forge/manifest-reference/modules/jira-custom-field/). When the
Forge version of the app is installed, Forge creates a
[background task](https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-tasks/#api-gro

## Parameters
- `connectKey` [path] (required) string — The key of the Connect app that contains the Jira issue field being migrated.
- `jiraIssueFieldsKey` [path] (required) string — The module key of the Connect issue field being migrated.

## Responses
- 200: object:
  - `description`: string
  - `elapsedRuntime` (required): integer(int64)
  - `finished`: string(date-time)
  - `id` (required): string
  - `lastUpdate` (required): string(date-time)
  - `message`: string
  - `progress` (required): integer(int64)
  - `result`: any
  - `self` (required): string(uri)
  - `started`: string(date-time)
  - `status` (required): string
  - `submitted`: string(date-time)
  - `submittedBy` (required): integer(int64)
- 401: object:
  - `message` (required): string
  - `statusCode` (required): integer
- 404: object:
  - `message` (required): string
  - `statusCode` (required): integer

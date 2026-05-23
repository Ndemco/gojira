# POST /rest/atlassian-connect/1/migration/{connectKey}/{jiraIssueFieldsKey}/task
**operationId:** `ConnectToForgeMigrationTaskSubmissionResource.submitTask_post`
**Summary:** Submit Connect issue field migration task

Submits a request to trigger migration of connect issue field to its Forge custom field counterpart.

When migrating a Connect app to Forge, [Issue Field](https://developer.atlassian.com/cloud/jira/software/modules/issue-field/) modules
must be converted to [Custom field](https://developer.atlassian.com/platform/forge/manifest-reference/modules/jira-custom-field/) modules.
This endpoint triggers the background migration of field data. Use the GET endpoint to retrieve
the status and progress of t

## Parameters
- `connectKey` [path] (required) string — The key of the Connect app that contains the Jira issue field being migrated.
- `jiraIssueFieldsKey` [path] (required) string — The module key of the Connect issue field being migrated.

## Responses
- 202: Returned if the migration task was submitted successfully.
- 401: object:
  - `message` (required): string
  - `statusCode` (required): integer
- 404: object:
  - `message` (required): string
  - `statusCode` (required): integer
- 409: object:
  - `message` (required): string
  - `statusCode` (required): integer

# GET /rest/api/3/workflowscheme/{id}/draft/issuetype/{issueType}
**operationId:** `getWorkflowSchemeDraftIssueType`
**Summary:** Get workflow for issue type in draft workflow scheme

Returns the issue type-workflow mapping for an issue type in a workflow scheme's draft.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `id` [path] (required) integer(int64) — The ID of the workflow scheme that the draft belongs to.
- `issueType` [path] (required) string — The ID of the issue type.

## Responses
- 200: object:
  - `issueType`: string
  - `updateDraftIfNeeded`: boolean
  - `workflow`: string
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the workflow scheme or issue type is not found.

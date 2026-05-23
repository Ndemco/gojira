# GET /rest/api/3/workflowscheme/{id}/issuetype/{issueType}
**operationId:** `getWorkflowSchemeIssueType`
**Summary:** Get workflow for issue type in workflow scheme

Returns the issue type-workflow mapping for an issue type in a workflow scheme.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `id` [path] (required) integer(int64) — The ID of the workflow scheme.
- `issueType` [path] (required) string — The ID of the issue type.
- `returnDraftIfExists` [query] boolean — Returns the mapping from the workflow scheme's draft rather than the workflow scheme, if set to true. If no draft exists

## Responses
- 200: object:
  - `issueType`: string
  - `updateDraftIfNeeded`: boolean
  - `workflow`: string
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the workflow scheme or issue type is not found.

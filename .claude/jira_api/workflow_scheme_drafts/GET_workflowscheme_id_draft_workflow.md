# GET /rest/api/3/workflowscheme/{id}/draft/workflow
**operationId:** `getDraftWorkflow`
**Summary:** Get issue types for workflows in draft workflow scheme

Returns the workflow-issue type mappings for a workflow scheme's draft.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `id` [path] (required) integer(int64) — The ID of the workflow scheme that the draft belongs to.
- `workflowName` [query] string — The name of a workflow in the scheme. Limits the results to the workflow-issue type mapping for the specified workflow.

## Responses
- 200: object:
  - `defaultMapping`: boolean
  - `issueTypes`: []string
  - `updateDraftIfNeeded`: boolean
  - `workflow`: string
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if either the workflow scheme or workflow (if specified) is not found. session.

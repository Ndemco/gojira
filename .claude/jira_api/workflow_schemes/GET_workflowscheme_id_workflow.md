# GET /rest/api/3/workflowscheme/{id}/workflow
**operationId:** `getWorkflow`
**Summary:** Get issue types for workflows in workflow scheme

Returns the workflow-issue type mappings for a workflow scheme.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `id` [path] (required) integer(int64) — The ID of the workflow scheme.
- `workflowName` [query] string — The name of a workflow in the scheme. Limits the results to the workflow-issue type mapping for the specified workflow.
- `returnDraftIfExists` [query] boolean — Returns the mapping from the workflow scheme's draft rather than the workflow scheme, if set to true. If no draft exists

## Responses
- 200: object:
  - `defaultMapping`: boolean
  - `issueTypes`: []string
  - `updateDraftIfNeeded`: boolean
  - `workflow`: string
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if either the workflow scheme or workflow is not found.

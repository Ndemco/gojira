# GET /rest/api/3/workflowscheme/{id}/default
**operationId:** `getDefaultWorkflow`
**Summary:** Get default workflow

Returns the default workflow for a workflow scheme. The default workflow is the workflow that is assigned any issue types that have not been mapped to any other workflow. The default workflow has *All Unassigned Issue Types* listed in its issue types for the workflow scheme in Jira.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `id` [path] (required) integer(int64) — The ID of the workflow scheme.
- `returnDraftIfExists` [query] boolean — Set to `true` to return the default workflow for the workflow scheme's draft rather than scheme itself. If the workflow 

## Responses
- 200: object:
  - `updateDraftIfNeeded`: boolean
  - `workflow` (required): string
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the workflow scheme is not found.

# DELETE /rest/api/3/workflowscheme/{id}/default
**operationId:** `deleteDefaultWorkflow`
**Summary:** Delete default workflow

Resets the default workflow for a workflow scheme. That is, the default workflow is set to Jira's system workflow (the *jira* workflow).

Note that active workflow schemes cannot be edited. If the workflow scheme is active, set `updateDraftIfNeeded` to `true` and a draft workflow scheme is created or updated with the default workflow reset. The draft workflow scheme can be published in Jira.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlas

## Parameters
- `id` [path] (required) integer(int64) — The ID of the workflow scheme.
- `updateDraftIfNeeded` [query] boolean — Set to true to create or update the draft of a workflow scheme and delete the mapping from the draft, when the workflow 

## Responses
- 200: object:
  - `defaultWorkflow`: string
  - `description`: string
  - `draft`: boolean
  - `id`: integer(int64)
  - `issueTypeMappings`: object
  - `issueTypes`: object
  - `lastModified`: string
  - `lastModifiedUser`: allOf(User)
  - `name`: string
  - `originalDefaultWorkflow`: string
  - `originalIssueTypeMappings`: object
  - `self`: string(uri)
  - `updateDraftIfNeeded`: boolean
- 400: Returned if the workflow scheme cannot be edited and `updateDraftIfNeeded` is not `true`.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the workflow scheme is not found.

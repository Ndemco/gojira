# PUT /rest/api/3/workflowscheme/{id}/workflow
**operationId:** `updateWorkflowMapping`
**Summary:** Set issue types for workflow in workflow scheme

Sets the issue types for a workflow in a workflow scheme. The workflow can also be set as the default workflow for the workflow scheme. Unmapped issues types are mapped to the default workflow.

Note that active workflow schemes cannot be edited. If the workflow scheme is active, set `updateDraftIfNeeded` to `true` in the request body and a draft workflow scheme is created or updated with the new workflow-issue types mappings. The draft workflow scheme can be published in Jira.

**[Permissions](

## Parameters
- `id` [path] (required) integer(int64) — The ID of the workflow scheme.
- `workflowName` [query] (required) string — The name of the workflow.

## Request Body
Content-Type: `application/json`
object:
  - `defaultMapping`: boolean
  - `issueTypes`: []string
  - `updateDraftIfNeeded`: boolean
  - `workflow`: string

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
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if any of the following is true:

 *  The workflow scheme is not found.
 *  The workflow is not found.
 *  The workflow is not specified.

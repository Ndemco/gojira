# PUT /rest/api/3/workflowscheme/{id}/issuetype/{issueType}
**operationId:** `setWorkflowSchemeIssueType`
**Summary:** Set workflow for issue type in workflow scheme

Sets the workflow for an issue type in a workflow scheme.

Note that active workflow schemes cannot be edited. If the workflow scheme is active, set `updateDraftIfNeeded` to `true` in the request body and a draft workflow scheme is created or updated with the new issue type-workflow mapping. The draft workflow scheme can be published in Jira.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `id` [path] (required) integer(int64) — The ID of the workflow scheme.
- `issueType` [path] (required) string — The ID of the issue type.

## Request Body
Content-Type: `application/json`
object:
  - `issueType`: string
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
- 400: Returned if the workflow cannot be edited and `updateDraftIfNeeded` is false.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the workflow scheme or issue type is not found.

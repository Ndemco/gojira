# PUT /rest/api/3/workflowscheme/{id}/draft
**operationId:** `updateWorkflowSchemeDraft`
**Summary:** Update draft workflow scheme

Updates a draft workflow scheme. If a draft workflow scheme does not exist for the active workflow scheme, then a draft is created. Note that an active workflow scheme can only have one draft workflow scheme.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `id` [path] (required) integer(int64) — The ID of the active workflow scheme that the draft was created from.

## Request Body
Content-Type: `application/json`
object:
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
- 404: Returned if:

 *  the original active workflow scheme is not found.
 *  the original active workflow scheme does not have a draft.

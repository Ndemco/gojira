# POST /rest/api/3/workflowscheme/{id}/createdraft
**operationId:** `createWorkflowSchemeDraftFromParent`
**Summary:** Create draft workflow scheme

Create a draft workflow scheme from an active workflow scheme, by copying the active workflow scheme. Note that an active workflow scheme can only have one draft workflow scheme.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `id` [path] (required) integer(int64) — The ID of the active workflow scheme that the draft is created from.

## Responses
- 201: object:
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

# GET /rest/api/3/workflowscheme/{id}
**operationId:** `getWorkflowScheme`
**Summary:** Get workflow scheme

Returns a workflow scheme.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `id` [path] (required) integer(int64) — The ID of the workflow scheme. Find this ID by editing the desired workflow scheme in Jira. The ID is shown in the URL a
- `returnDraftIfExists` [query] boolean — Returns the workflow scheme's draft rather than scheme itself, if set to true. If the workflow scheme does not have a dr

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
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the workflow scheme is not found.

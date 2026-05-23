# GET /rest/api/3/component/{id}/relatedIssueCounts
**operationId:** `getComponentRelatedIssues`
**Summary:** Get component issues count

Returns the counts of issues assigned to the component.

This operation can be accessed anonymously.

**Deprecation notice:** The required OAuth 2.0 scopes will be updated on June 15, 2024.

 *  **Classic**: `read:jira-work`
 *  **Granular**: `read:field:jira`, `read:project.component:jira`

**[Permissions](#permissions) required:** None.

## Parameters
- `id` [path] (required) string — The ID of the component.

## Responses
- 200: object:
  - `issueCount`: integer(int64)
  - `self`: string(uri)
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the component is not found.

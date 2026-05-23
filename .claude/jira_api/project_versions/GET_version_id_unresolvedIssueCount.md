# GET /rest/api/3/version/{id}/unresolvedIssueCount
**operationId:** `getVersionUnresolvedIssues`
**Summary:** Get version's unresolved issues count

Returns counts of the issues and unresolved issues for the project version.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** *Browse projects* project permission for the project that contains the version.

## Parameters
- `id` [path] (required) string — The ID of the version.

## Responses
- 200: object:
  - `issuesCount`: integer(int64)
  - `issuesUnresolvedCount`: integer(int64)
  - `self`: string(uri)
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if:

 *  the version is not found.
 *  the user does not have the required permissions.

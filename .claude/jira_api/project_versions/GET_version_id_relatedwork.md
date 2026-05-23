# GET /rest/api/3/version/{id}/relatedwork
**operationId:** `getRelatedWork`
**Summary:** Get related work

Returns related work items for the given version id.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project containing the version.

## Parameters
- `id` [path] (required) string — The ID of the version.

## Responses
- 200: []object:
  - `category` (required): string
  - `issueId`: integer(int64)
  - `relatedWorkId`: string
  - `title`: string
  - `url`: string(uri)
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the version is not found or the user does not have the necessary permission.
- 500: Returned if reading related work fails

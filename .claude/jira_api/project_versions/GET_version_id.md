# GET /rest/api/3/version/{id}
**operationId:** `getVersion`
**Summary:** Get version

Returns a project version.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project containing the version.

## Parameters
- `id` [path] (required) string — The ID of the version.
- `expand` [query] string — Use [expand](#expansion) to include additional information about version in the response. This parameter accepts a comma

## Responses
- 200: object:
  - `approvers`: []VersionApprover
  - `archived`: boolean
  - `description`: string
  - `driver`: string
  - `expand`: string
  - `id`: string
  - `issuesStatusForFixVersion`: allOf(VersionIssuesStatus)
  - `moveUnfixedIssuesTo`: string(uri)
  - `name`: string
  - `operations`: []SimpleLink
  - `overdue`: boolean
  - `project`: string
  - `projectId`: integer(int64)
  - `releaseDate`: string(date)
  - `released`: boolean
  - `self`: string(uri)
  - `startDate`: string(date)
  - `userReleaseDate`: string
  - `userStartDate`: string
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the version is not found or the user does not have the necessary permission.

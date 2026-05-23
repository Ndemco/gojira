# POST /rest/api/3/version/{id}/move
**operationId:** `moveVersion`
**Summary:** Move version

Modifies the version's sequence within the project, which affects the display order of the versions in Jira.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** *Browse projects* project permission for the project that contains the version.

## Parameters
- `id` [path] (required) string — The ID of the version to be moved.

## Request Body
Content-Type: `application/json`
object:
  - `after`: string(uri)
  - `position`: string

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
- 400: Returned if:

 *  no body parameters are provided.
 *  `after` and `position` are provided.
 *  `position` is invalid.
- 401: Returned if:

 *  the authentication credentials are incorrect or missing
 *  the user does not have the required commissions.
- 404: Returned if the version or move after version are not found.

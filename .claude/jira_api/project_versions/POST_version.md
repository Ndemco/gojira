# POST /rest/api/3/version
**operationId:** `createVersion`
**Summary:** Create version

Creates a project version.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) or *Administer Projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project the version is added to.

## Request Body
Content-Type: `application/json`
object:
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

## Responses
- 201: object:
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
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if:

 *  the project is not found.
 *  the user does not have the required permissions.

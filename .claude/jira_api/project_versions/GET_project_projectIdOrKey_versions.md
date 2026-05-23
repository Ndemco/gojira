# GET /rest/api/3/project/{projectIdOrKey}/versions
**operationId:** `getProjectVersions`
**Summary:** Get project versions

Returns all versions in a project. The response is not paginated. Use [Get project versions paginated](#api-rest-api-3-project-projectIdOrKey-version-get) if you want to get the versions in a project with pagination.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** *Browse Projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project.

## Parameters
- `projectIdOrKey` [path] (required) string — The project ID or project key (case sensitive).
- `expand` [query] string — Use [expand](#expansion) to include additional information in the response. This parameter accepts `operations`, which r

## Responses
- 200: []object:
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
- 404: Returned if the project is not found or the user does not have permission to view it.

# GET /rest/api/3/project/{projectIdOrKey}/version
**operationId:** `getProjectVersionsPaginated`
**Summary:** Get project versions paginated

Returns a [paginated](#pagination) list of all versions in a project. See the [Get project versions](#api-rest-api-3-project-projectIdOrKey-versions-get) resource if you want to get a full list of versions without pagination.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** *Browse Projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project.

## Parameters
- `projectIdOrKey` [path] (required) string — The project ID or project key (case sensitive).
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.
- `orderBy` [query] string — [Order](#ordering) the results by a field:
- `query` [query] string — Filter the results using a literal string. Versions with matching `name` or `description` are returned (case insensitive
- `status` [query] string — A list of status values used to filter the results by version status. This parameter accepts a comma-separated list. The
- `expand` [query] string — Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated lis

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []Version
- 404: Returned if the project is not found or the user does not have permission to view it.

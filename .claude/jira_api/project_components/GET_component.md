# GET /rest/api/3/component
**operationId:** `findComponentsForProjects`
**Summary:** Find components for projects

Returns a [paginated](#pagination) list of all components in a project, including global (Compass) components when applicable.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** *Browse Projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project.

## Parameters
- `projectIdsOrKeys` [query] []string — The project IDs and/or project keys (case sensitive).
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.
- `orderBy` [query] string — [Order](#ordering) the results by a field:
- `query` [query] string — Filter the results using a literal string. Components with a matching `name` or `description` are returned (case insensi

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []ComponentJsonBean
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the project is not found or the user does not have permission to view it.

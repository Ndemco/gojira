# GET /rest/api/3/project/{projectIdOrKey}/component
**operationId:** `getProjectComponentsPaginated`
**Summary:** Get project components paginated

Returns a [paginated](#pagination) list of all components in a project. See the [Get project components](#api-rest-api-3-project-projectIdOrKey-components-get) resource if you want to get a full list of versions without pagination.

If your project uses Compass components, this API will return a list of Compass components that are linked to issues in that project.

This operation can be accessed anonymously.

**[Permissions](#permissions) required:** *Browse Projects* [project permission](https:

## Parameters
- `projectIdOrKey` [path] (required) string — The project ID or project key (case sensitive).
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.
- `orderBy` [query] string — [Order](#ordering) the results by a field:
- `componentSource` [query] string — The source of the components to return. Can be `jira` (default), `compass` or `auto`. When `auto` is specified, the API 
- `query` [query] string — Filter the results using a literal string. Components with a matching `name` or `description` are returned (case insensi

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []ComponentWithIssueCount
- 401: Returned if the authentication credentials are incorrect or missing.
- 404: Returned if the project is not found or the user does not have permission to view it.

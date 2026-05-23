# GET /rest/api/3/priorityscheme/{schemeId}/projects
**operationId:** `getProjectsByPriorityScheme`
**Summary:** Get projects by priority scheme

Returns a [paginated](#pagination) list of projects by scheme.

**[Permissions](#permissions) required:** Permission to access Jira.

## Parameters
- `startAt` [query] string — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] string — The maximum number of items to return per page.
- `projectId` [query] []integer(int64) — The project IDs to filter by. For example, `projectId=10000&projectId=10001`.
- `schemeId` [path] (required) string — The priority scheme ID.
- `query` [query] string — The string to query projects on by name.

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []Project
- 400: Returned if the request isn't valid.
- 401: Returned if the authentication credentials are incorrect.

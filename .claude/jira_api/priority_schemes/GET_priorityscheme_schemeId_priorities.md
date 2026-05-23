# GET /rest/api/3/priorityscheme/{schemeId}/priorities
**operationId:** `getPrioritiesByPriorityScheme`
**Summary:** Get priorities by priority scheme

Returns a [paginated](#pagination) list of priorities by scheme.

**[Permissions](#permissions) required:** Permission to access Jira.

## Parameters
- `startAt` [query] string — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] string — The maximum number of items to return per page.
- `schemeId` [path] (required) string — The priority scheme ID.

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []PriorityWithSequence
- 400: Returned if the request isn't valid.
- 401: Returned if the authentication credentials are incorrect.

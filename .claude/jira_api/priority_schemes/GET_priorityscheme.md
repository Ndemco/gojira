# GET /rest/api/3/priorityscheme
**operationId:** `getPrioritySchemes`
**Summary:** Get priority schemes

Returns a [paginated](#pagination) list of priority schemes.

**[Permissions](#permissions) required:** Permission to access Jira.

## Parameters
- `startAt` [query] string — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] string — The maximum number of items to return per page.
- `priorityId` [query] []integer(int64) — A set of priority IDs to filter by. To include multiple IDs, provide an ampersand-separated list. For example, `priority
- `schemeId` [query] []integer(int64) — A set of priority scheme IDs. To include multiple IDs, provide an ampersand-separated list. For example, `schemeId=10000
- `schemeName` [query] string — The name of scheme to search for.
- `onlyDefault` [query] boolean — Whether only the default priority is returned.
- `orderBy` [query] string — The ordering to return the priority schemes by.
- `expand` [query] string — A comma separated list of additional information to return. "priorities" will return priorities associated with the prio

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []PrioritySchemeWithPaginatedPrioritiesAndProjects
- 400: Returned if the request isn't valid.
- 401: Returned if the authentication credentials are incorrect.

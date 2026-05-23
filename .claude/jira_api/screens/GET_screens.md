# GET /rest/api/3/screens
**operationId:** `getScreens`
**Summary:** Get screens

Returns a [paginated](#pagination) list of all screens or those specified by one or more screen IDs.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.
- `id` [query] []integer(int64) — The list of screen IDs. To include multiple IDs, provide an ampersand-separated list. For example, `id=10000&id=10001`.
- `queryString` [query] string — String used to perform a case-insensitive partial match with screen name.
- `scope` [query] []string — The scope filter string. To filter by multiple scope, provide an ampersand-separated list. For example, `scope=GLOBAL&sc
- `orderBy` [query] string — [Order](#ordering) the results by a field:

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []Screen
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.

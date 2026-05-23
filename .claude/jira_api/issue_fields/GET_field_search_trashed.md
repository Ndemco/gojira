# GET /rest/api/3/field/search/trashed
**operationId:** `getTrashedFieldsPaginated`
**Summary:** Get fields in trash paginated

Returns a [paginated](#pagination) list of fields in the trash. The list may be restricted to fields whose field name or description partially match a string.

Only custom fields can be queried, `type` must be set to `custom`.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.
- `id` [query] []string — 
- `query` [query] string — String used to perform a case-insensitive partial match with field names or descriptions.
- `expand` [query] string — 
- `orderBy` [query] string — [Order](#ordering) the results by a field:

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []Field
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)

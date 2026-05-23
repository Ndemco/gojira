# GET /rest/api/3/issuetypescreenscheme
**operationId:** `getIssueTypeScreenSchemes`
**Summary:** Get issue type screen schemes

Returns a [paginated](#pagination) list of issue type screen schemes.

Only issue type screen schemes used in classic projects are returned.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.
- `id` [query] []integer(int64) — The list of issue type screen scheme IDs. To include multiple IDs, provide an ampersand-separated list. For example, `id
- `queryString` [query] string — String used to perform a case-insensitive partial match with issue type screen scheme name.
- `orderBy` [query] string — [Order](#ordering) the results by a field:
- `expand` [query] string — Use [expand](#expansion) to include additional information in the response. This parameter accepts `projects` that, for 

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []IssueTypeScreenScheme
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the required permissions.

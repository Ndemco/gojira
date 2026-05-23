# GET /rest/api/3/issuetypescheme
**operationId:** `getAllIssueTypeSchemes`
**Summary:** Get all issue type schemes

Returns a [paginated](#pagination) list of issue type schemes.

Only issue type schemes used in classic projects are returned.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `startAt` [query] integer(int64) — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] integer(int32) — The maximum number of items to return per page.
- `id` [query] []integer(int64) — The list of issue type schemes IDs. To include multiple IDs, provide an ampersand-separated list. For example, `id=10000
- `orderBy` [query] string — [Order](#ordering) the results by a field:
- `expand` [query] string — Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated lis
- `queryString` [query] string — String used to perform a case-insensitive partial match with issue type scheme name.

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []IssueTypeScheme
- 400: Returned if the request is not valid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the required permissions.

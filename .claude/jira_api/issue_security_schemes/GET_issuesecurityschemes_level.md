# GET /rest/api/3/issuesecurityschemes/level
**operationId:** `getSecurityLevels`
**Summary:** Get issue security levels

Returns a [paginated](#pagination) list of issue security levels.

Only issue security levels in the context of classic projects are returned.

Filtering using IDs is inclusive: if you specify both security scheme IDs and level IDs, the result will include both specified issue security levels and all issue security levels from the specified schemes.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `startAt` [query] string — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] string — The maximum number of items to return per page.
- `id` [query] []string — The list of issue security scheme level IDs. To include multiple issue security levels, separate IDs with an ampersand: 
- `schemeId` [query] []string — The list of issue security scheme IDs. To include multiple issue security schemes, separate IDs with an ampersand: `sche
- `onlyDefault` [query] boolean — When set to true, returns multiple default levels for each security scheme containing a default. If you provide scheme a

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []SecurityLevel
- 400: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 401: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 403: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)

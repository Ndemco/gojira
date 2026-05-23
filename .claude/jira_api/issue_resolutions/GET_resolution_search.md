# GET /rest/api/3/resolution/search
**operationId:** `searchResolutions`
**Summary:** Search resolutions

Returns a [paginated](#pagination) list of resolutions. The list can contain all resolutions or a subset determined by any combination of these criteria:

 *  a list of resolutions IDs.
 *  whether the field configuration is a default. This returns resolutions from company-managed (classic) projects only, as there is no concept of default resolutions in team-managed projects.

**[Permissions](#permissions) required:** Permission to access Jira.

## Parameters
- `startAt` [query] string — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] string — The maximum number of items to return per page.
- `id` [query] []string — The list of resolutions IDs to be filtered out
- `onlyDefault` [query] boolean — When set to true, return default only, when IDs provided, if none of them is default, return empty page. Default value i

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []ResolutionJsonBean
- 401: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)

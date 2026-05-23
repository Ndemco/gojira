# GET /rest/api/3/priority/search
**operationId:** `searchPriorities`
**Summary:** Search priorities

Returns a [paginated](#pagination) list of priorities. The list can contain all priorities or a subset determined by any combination of these criteria:

 *  a list of priority IDs. Any invalid priority IDs are ignored.
 *  a list of project IDs. Only priorities that are available in these projects will be returned. Any invalid project IDs are ignored.
 *  whether the field configuration is a default. This returns priorities from company-managed (classic) projects only, as there is no concept of 

## Parameters
- `startAt` [query] string — The index of the first item to return in a page of results (page offset).
- `maxResults` [query] string — The maximum number of items to return per page.
- `id` [query] []string — The list of priority IDs. To include multiple IDs, provide an ampersand-separated list. For example, `id=2&id=3`.
- `projectId` [query] []string — The list of projects IDs. To include multiple IDs, provide an ampersand-separated list. For example, `projectId=10010&pr
- `priorityName` [query] string — The name of priority to search for.
- `onlyDefault` [query] boolean — Whether only the default priority is returned.
- `expand` [query] string — Use `schemes` to return the associated priority schemes for each priority. Limited to returning first 15 priority scheme

## Responses
- 200: object:
  - `isLast`: boolean
  - `maxResults`: integer(int32)
  - `nextPage`: string(uri)
  - `self`: string(uri)
  - `startAt`: integer(int64)
  - `total`: integer(int64)
  - `values`: []Priority
- 401: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)

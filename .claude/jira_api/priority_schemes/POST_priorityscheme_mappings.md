# POST /rest/api/3/priorityscheme/mappings
**operationId:** `suggestedPrioritiesForMappings`
**Summary:** Suggested priorities for mappings

Returns a [paginated](#pagination) list of priorities that would require mapping, given a change in priorities or projects associated with a priority scheme.

**[Permissions](#permissions) required:** Permission to access Jira.

## Request Body
Content-Type: `application/json`
object:
  - `maxResults`: integer(int32)
  - `priorities`: allOf(SuggestedMappingsForPrioritiesRequestBean)
  - `projects`: allOf(SuggestedMappingsForProjectsRequestBean)
  - `schemeId`: integer(int64)
  - `startAt`: integer(int64)

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

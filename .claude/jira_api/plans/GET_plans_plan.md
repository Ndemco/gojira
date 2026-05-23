# GET /rest/api/3/plans/plan
**operationId:** `getPlans`
**Summary:** Get plans paginated

Returns a [paginated](#pagination) list of plans.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `includeTrashed` [query] boolean — Whether to include trashed plans in the results.
- `includeArchived` [query] boolean — Whether to include archived plans in the results.
- `cursor` [query] string — The cursor to start from. If not provided, the first page will be returned.
- `maxResults` [query] integer(int32) — The maximum number of plans to return per page. The maximum value is 50. The default value is 50.

## Responses
- 200: object:
  - `cursor`: string
  - `last`: boolean
  - `nextPageCursor`: string
  - `size`: integer(int32)
  - `total`: integer(int64)
  - `values`: []GetPlanResponseForPage
- 401: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 403: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)

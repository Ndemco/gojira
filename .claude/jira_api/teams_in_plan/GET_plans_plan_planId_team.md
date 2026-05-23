# GET /rest/api/3/plans/plan/{planId}/team
**operationId:** `getTeams`
**Summary:** Get teams in plan paginated

Returns a [paginated](#pagination) list of plan-only and Atlassian teams in a plan.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `planId` [path] (required) integer(int64) — The ID of the plan.
- `cursor` [query] string — The cursor to start from. If not provided, the first page will be returned.
- `maxResults` [query] integer(int32) — The maximum number of plan teams to return per page. The maximum value is 50. The default value is 50.

## Responses
- 200: object:
  - `cursor`: string
  - `last`: boolean
  - `nextPageCursor`: string
  - `size`: integer(int32)
  - `total`: integer(int64)
  - `values`: []GetTeamResponseForPage
- 401: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 403: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 404: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)

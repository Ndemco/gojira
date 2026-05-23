# GET /rest/api/3/plans/plan/{planId}/team/planonly/{planOnlyTeamId}
**operationId:** `getPlanOnlyTeam`
**Summary:** Get plan-only team

Returns planning settings for a plan-only team.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `planId` [path] (required) integer(int64) — The ID of the plan.
- `planOnlyTeamId` [path] (required) integer(int64) — The ID of the plan-only team.

## Responses
- 200: object:
  - `capacity`: number(double)
  - `id` (required): integer(int64)
  - `issueSourceId`: integer(int64)
  - `memberAccountIds`: []string
  - `name` (required): string
  - `planningStyle` (required): string
  - `sprintLength`: integer(int64)
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
- 409: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)

# GET /rest/api/3/plans/plan/{planId}/team/atlassian/{atlassianTeamId}
**operationId:** `getAtlassianTeam`
**Summary:** Get Atlassian team in plan

Returns planning settings for an Atlassian team in a plan.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `planId` [path] (required) integer(int64) — The ID of the plan.
- `atlassianTeamId` [path] (required) string — The ID of the Atlassian team.

## Responses
- 200: object:
  - `capacity`: number(double)
  - `id` (required): string
  - `issueSourceId`: integer(int64)
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

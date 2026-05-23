# POST /rest/api/3/plans/plan/{planId}/team/atlassian
**operationId:** `addAtlassianTeam`
**Summary:** Add Atlassian team to plan

Adds an existing Atlassian team to a plan and configures their plannning settings.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `planId` [path] (required) integer(int64) — The ID of the plan.

## Request Body
Content-Type: `application/json`
object:
  - `capacity`: number(double)
  - `id` (required): string
  - `issueSourceId`: integer(int64)
  - `planningStyle` (required): string
  - `sprintLength`: integer(int64)

## Responses
- 204: any
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
- 404: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)
- 409: object:
  - `errorMessages`: []string
  - `errors`: object
  - `status`: integer(int32)

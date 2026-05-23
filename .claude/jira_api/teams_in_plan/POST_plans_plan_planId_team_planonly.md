# POST /rest/api/3/plans/plan/{planId}/team/planonly
**operationId:** `createPlanOnlyTeam`
**Summary:** Create plan-only team

Creates a plan-only team and configures their planning settings.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `planId` [path] (required) integer(int64) — The ID of the plan.

## Request Body
Content-Type: `application/json`
object:
  - `capacity`: number(double)
  - `issueSourceId`: integer(int64)
  - `memberAccountIds`: []string
  - `name` (required): string
  - `planningStyle` (required): string
  - `sprintLength`: integer(int64)

## Responses
- 201: integer(int64)
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

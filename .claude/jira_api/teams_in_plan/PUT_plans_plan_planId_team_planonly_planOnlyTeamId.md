# PUT /rest/api/3/plans/plan/{planId}/team/planonly/{planOnlyTeamId}
**operationId:** `updatePlanOnlyTeam`
**Summary:** Update plan-only team

Updates any of the following planning settings of a plan-only team using [JSON Patch](https://datatracker.ietf.org/doc/html/rfc6902).

 *  name
 *  planningStyle
 *  issueSourceId
 *  sprintLength
 *  capacity
 *  memberAccountIds

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

*Note that "add" operations do not respect array indexes in target locations. Call the "Get plan-only team" endpoint to find out the order of a

## Parameters
- `planId` [path] (required) integer(int64) — The ID of the plan.
- `planOnlyTeamId` [path] (required) integer(int64) — The ID of the plan-only team.

## Request Body
Content-Type: `application/json-patch+json`
object

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

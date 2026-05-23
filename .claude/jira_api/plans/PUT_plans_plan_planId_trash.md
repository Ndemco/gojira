# PUT /rest/api/3/plans/plan/{planId}/trash
**operationId:** `trashPlan`
**Summary:** Trash plan

Moves a plan to trash.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `planId` [path] (required) integer(int64) — The ID of the plan.

## Responses
- 204: any
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

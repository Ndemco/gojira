# GET /rest/api/3/plans/plan/{planId}
**operationId:** `getPlan`
**Summary:** Get plan

Returns a plan.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `planId` [path] (required) integer(int64) — The ID of the plan.
- `useGroupId` [query] boolean — Whether to return group IDs instead of group names. Group names are deprecated.

## Responses
- 200: object:
  - `crossProjectReleases`: []GetCrossProjectReleaseResponse
  - `customFields`: []GetCustomFieldResponse
  - `exclusionRules`: allOf(GetExclusionRulesResponse)
  - `id` (required): integer(int64)
  - `issueSources`: []GetIssueSourceResponse
  - `lastSaved`: string
  - `leadAccountId`: string
  - `name`: string
  - `permissions`: []GetPermissionResponse
  - `scheduling` (required): allOf(GetSchedulingResponse)
  - `status` (required): string
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

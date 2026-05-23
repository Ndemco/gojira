# POST /rest/api/3/plans/plan
**operationId:** `createPlan`
**Summary:** Create plan

Creates a plan.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `useGroupId` [query] boolean — Whether to accept group IDs instead of group names. Group names are deprecated.

## Request Body
Content-Type: `application/json`
object:
  - `crossProjectReleases`: []CreateCrossProjectReleaseRequest
  - `customFields`: []CreateCustomFieldRequest
  - `exclusionRules`: allOf(CreateExclusionRulesRequest)
  - `issueSources` (required): []CreateIssueSourceRequest
  - `leadAccountId`: string
  - `name` (required): string
  - `permissions`: []CreatePermissionRequest
  - `scheduling` (required): allOf(CreateSchedulingRequest)

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

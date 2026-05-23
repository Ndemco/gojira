# PUT /rest/api/3/issuesecurityschemes/{schemeId}/level/{levelId}/member
**operationId:** `addSecurityLevelMembers`
**Summary:** Add issue security level members

Adds members to the issue security level. You can add up to 100 members per request.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `schemeId` [path] (required) string — The ID of the issue security scheme.
- `levelId` [path] (required) string — The ID of the issue security level.

## Request Body
Content-Type: `application/json`
object:
  - `members`: []SecuritySchemeLevelMemberBean

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

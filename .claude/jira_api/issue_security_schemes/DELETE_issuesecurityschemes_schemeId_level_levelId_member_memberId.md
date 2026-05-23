# DELETE /rest/api/3/issuesecurityschemes/{schemeId}/level/{levelId}/member/{memberId}
**operationId:** `removeMemberFromSecurityLevel`
**Summary:** Remove member from issue security level

Removes an issue security level member from an issue security scheme.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `schemeId` [path] (required) string — The ID of the issue security scheme.
- `levelId` [path] (required) string — The ID of the issue security level.
- `memberId` [path] (required) string — The ID of the issue security level member to be removed.

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

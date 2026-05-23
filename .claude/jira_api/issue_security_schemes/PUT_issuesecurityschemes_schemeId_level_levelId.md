# PUT /rest/api/3/issuesecurityschemes/{schemeId}/level/{levelId}
**operationId:** `updateSecurityLevel`
**Summary:** Update issue security level

Updates the issue security level.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `schemeId` [path] (required) string — The ID of the issue security scheme level belongs to.
- `levelId` [path] (required) string — The ID of the issue security level to update.

## Request Body
Content-Type: `application/json`
object:
  - `description`: string
  - `name`: string

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

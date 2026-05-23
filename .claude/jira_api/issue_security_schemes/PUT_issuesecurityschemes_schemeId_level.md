# PUT /rest/api/3/issuesecurityschemes/{schemeId}/level
**operationId:** `addSecurityLevel`
**Summary:** Add issue security levels

Adds levels and levels' members to the issue security scheme. You can add up to 100 levels per request.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `schemeId` [path] (required) string — The ID of the issue security scheme.

## Request Body
Content-Type: `application/json`
object:
  - `levels`: []SecuritySchemeLevelBean

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

# POST /rest/api/3/issuesecurityschemes
**operationId:** `createIssueSecurityScheme`
**Summary:** Create issue security scheme

Creates a security scheme with security scheme levels and levels' members. You can create up to 100 security scheme levels and security scheme levels' members per request.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Request Body
Content-Type: `application/json`
object:
  - `description`: string
  - `levels`: []SecuritySchemeLevelBean
  - `name` (required): string

## Responses
- 201: object:
  - `id` (required): string
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

# PUT /rest/api/3/issuesecurityschemes/{id}
**operationId:** `updateIssueSecurityScheme`
**Summary:** Update issue security scheme

Updates the issue security scheme.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `id` [path] (required) string — The ID of the issue security scheme.

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

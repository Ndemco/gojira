# PUT /rest/api/3/screens/{screenId}
**operationId:** `updateScreen`
**Summary:** Update screen

Updates a screen. Only screens used in classic projects can be updated.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `screenId` [path] (required) integer(int64) — The ID of the screen.

## Request Body
Content-Type: `application/json`
object:
  - `description`: string
  - `name`: string

## Responses
- 200: object:
  - `description`: string
  - `id`: integer(int64)
  - `name`: string
  - `scope`: allOf(Scope)
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any

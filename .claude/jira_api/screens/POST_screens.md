# POST /rest/api/3/screens
**operationId:** `createScreen`
**Summary:** Create screen

Creates a screen with a default field tab.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Request Body
Content-Type: `application/json`
object:
  - `description`: string
  - `name` (required): string

## Responses
- 201: object:
  - `description`: string
  - `id`: integer(int64)
  - `name`: string
  - `scope`: allOf(Scope)
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any

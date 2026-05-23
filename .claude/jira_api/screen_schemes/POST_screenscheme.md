# POST /rest/api/3/screenscheme
**operationId:** `createScreenScheme`
**Summary:** Create screen scheme

Creates a screen scheme.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Request Body
Content-Type: `application/json`
object:
  - `description`: string
  - `name` (required): string
  - `screens` (required): allOf(ScreenTypes)

## Responses
- 201: object:
  - `id` (required): integer(int64)
- 400: any
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: any
- 404: any

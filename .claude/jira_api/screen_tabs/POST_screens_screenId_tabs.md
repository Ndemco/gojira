# POST /rest/api/3/screens/{screenId}/tabs
**operationId:** `addScreenTab`
**Summary:** Create screen tab

Creates a tab for a screen.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `screenId` [path] (required) integer(int64) — The ID of the screen.

## Request Body
Content-Type: `application/json`
object:
  - `id`: integer(int64)
  - `name` (required): string

## Responses
- 200: object:
  - `id`: integer(int64)
  - `name` (required): string
- 400: Returned if the request is invalid.
- 401: Returned if the authentication credentials are incorrect or missing.
- 403: Returned if the user does not have the necessary permission.
- 404: Returned if the screen is not found.

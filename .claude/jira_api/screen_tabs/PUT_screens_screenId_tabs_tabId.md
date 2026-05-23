# PUT /rest/api/3/screens/{screenId}/tabs/{tabId}
**operationId:** `renameScreenTab`
**Summary:** Update screen tab

Updates the name of a screen tab.

**[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

## Parameters
- `screenId` [path] (required) integer(int64) — The ID of the screen.
- `tabId` [path] (required) integer(int64) — The ID of the screen tab.

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
- 404: Returned if the screen or screen tab is not found.
